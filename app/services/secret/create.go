package secret

import (
	"encoding/json"
	"fmt"
	"gorm.io/datatypes"
	"log"
	"net/http"
	"regexp"
	"time"

	"FuguBackend/app/code"
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/pkg/password"
	"FuguBackend/app/repository/mysql/secrets"
	"FuguBackend/app/repository/redis"
	"FuguBackend/config"
	"FuguBackend/pkg/snowflake"

	"go.uber.org/zap"
)

// CreateSecretData ...  一对多 has many
type CreateSecretData struct {
	SecretID int `json:"secretId,omitempty" gorm:"column:secretid;type:bigint"`
	AuthorID int `json:"authorId,omitempty" gorm:"column:authorid;type:bigint"`
	// ViewLevel 1 仅广场 2 仅洞穴 3 广场和洞穴
	ViewLevel int      `json:"viewLevel,omitempty" gorm:"column:viewlevel;type:int"`
	Timestamp int64    `json:"timestamp,omitempty" gorm:"column:timestamp;type:bigint"`
	Views     int64    `json:"views,omitempty" gorm:"column:views;type:bigint"`
	Content   string   `json:"content,omitempty" gorm:"column:content;type:varchar(255)"`
	Images    []string `json:"images,omitempty" gorm:"foreignKey:SecretID"`
	// Status 秘闻状态 平台可以将非法的帖子设置为不可见
	Status bool `json:"status,omitempty" gorm:"column:status;type:tinyint"`
}
type Image struct {
	SecretID int    `json:"SecretID,omitempty" gorm:"column:secretid;type:bigint"`
	ImageUrl string `json:"ImageUrl,omitempty" gorm:"column:imageurl;type:varchar(255)"`
}

func (s *service) Create(c core.Context, hashID string, data *CreateSecretData) (id int, err error) {

	// 提取被艾特的人
	screenName, _ := extractUsername(data.Content)
	//screenName = "Fergus_Hinn"
	s.logger.Info(fmt.Sprintf("=====people to be @ :%v=======", screenName))
	// 如果有 那么就推特私信
	if screenName != "" {
		inviteCode := password.GenInviteCode(hashID, config.DefaultInviteCodeLength)
		err = s.cache.Set(inviteCode, hashID, config.DefaultInviteCodeTTL, redis.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusInternalServerError,
				code.CacheSetError,
				code.Text(code.CacheSetError)).WithError(err),
			)
			return
		}
		err = s.twSvc.DirectMessage(c, screenName, buildMessage(screenName, inviteCode))
		if err != nil {
			s.logger.Error("send direct message failed ...", zap.Error(err))

		}
	}
	// 写库

	model := secrets.Secret{}
	genID, _ := snowflake.GenID()
	model.SecretID = genID
	model.AuthorID = data.AuthorID
	model.ViewLevel = data.ViewLevel
	model.Timestamp = time.Now().Unix()
	model.Content = data.Content
	model.Status = true
	jsonImages, _ := json.Marshal(data.Images)
	model.Images = datatypes.JSON(jsonImages)
	secretID, err := secrets.CreatSecret(model)
	if err != nil {
		return 0, err
	}

	return secretID, nil

}

// extractUsername 找内容中的被@的推特用户
func extractUsername(jsonStr string) (string, error) {
	// 解析 JSON 字符串到一个 map[string]interface{} 对象
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		return "", err
	}

	// 获取带有 @ 的字符串
	text, ok := data["text"].(string)
	if !ok {
		return "", fmt.Errorf("找不到文本字段或文本字段不是字符串")
	}

	// 使用正则表达式提取 @ 之后和空格之前的字符串
	re := regexp.MustCompile(`@(\S+)\s`)
	match := re.FindStringSubmatch(text)
	if len(match) < 2 {
		return "", fmt.Errorf("无法从字符串中提取用户名")
	}

	// 返回提取的用户名

	return match[1], nil
}

func testExtract() {
	jsonStr := `{"id": 123, "text": "Hello @john_doe, how are you?"}`

	username, err := extractUsername(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("用户名:", username)
}
func buildMessage(name, code string) string {
	defaultMessage := fmt.Sprintf(`Hey, @%v, 
	There's a secret admirer on Fugu Toxic! 
	Explore the mysterious world of anonymous affection, earn points, snag rewards. 
	Just log in with Twitter, grab a quirky username, and let the rewards flow. 
	Invitation code：%v，url：fugutoxic.com`, name, code)
	return defaultMessage
}
