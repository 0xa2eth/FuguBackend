package secret

import (
	"FuguBackend/app/pkg/core"
	"FuguBackend/app/repository/mysql/secrets"
	"FuguBackend/pkg/snowflake"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
	"regexp"
	"time"
)

// Secret ...  一对多 has many
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
type SecretImage struct {
	SecretID int    `json:"SecretID,omitempty" gorm:"column:secretid;type:bigint"`
	ImageUrl string `json:"ImageUrl,omitempty" gorm:"column:imageurl;type:varchar(255)"`
}

func (s *service) Create(c core.Context, data *CreateSecretData) (id int64, err error) {

	// 提取被艾特的人
	screenName, _ := extractUsername(data.Content)
	// 如果有 那么就推特私信
	if screenName != "" {
		err = s.twSvc.DirectMessage(c, screenName, buildMessage())
		if err != nil {
			s.logger.Error("send direct message failed ...", zap.Error(err))

		}
	}
	// 写库
	model := secrets.NewModel()
	genID, _ := snowflake.GenID()
	model.SecretId = int64(genID)
	model.Authorid = int64(data.AuthorID)
	model.ViewLevel = int64(data.ViewLevel)
	model.Timestamp = time.Now().Unix()
	model.Status = 1

	id, err = model.Create(s.db.GetDbW().WithContext(c.RequestContext()))
	if err != nil {
		return 0, err
	}

	//for i, image := range data.Images {
	//
	//}
	//newModel := secret_images.NewModel()
	//newModel.ImageUrl =

	return id, nil

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
func buildMessage() string {
	defaultMessage := "One of your twitter friends in FuGu-Toxic has mentioned you in his latest secret,come and see by clicking the link https://metahome.tech "
	return defaultMessage
}
