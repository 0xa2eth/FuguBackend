///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package users

import (
	"FuguBackend/app/repository/mysql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Users {
	return new(Users)
}

func NewQueryBuilder() *usersQueryBuilder {
	return new(usersQueryBuilder)
}

func (t *Users) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Userid, nil
}

type usersQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *usersQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *usersQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Users{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *usersQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Users{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *usersQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Users{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *usersQueryBuilder) First(db *gorm.DB) (*Users, error) {
	ret := &Users{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *usersQueryBuilder) QueryOne(db *gorm.DB) (*Users, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *usersQueryBuilder) QueryAll(db *gorm.DB) ([]*Users, error) {
	var ret []*Users
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *usersQueryBuilder) Limit(limit int) *usersQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *usersQueryBuilder) Offset(offset int) *usersQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *usersQueryBuilder) WhereId(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereIdIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereIdNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderById(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCreatedAtIn(value []time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByCreatedAt(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereUpdatedAtIn(value []time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByUpdatedAt(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereDeletedAt(p mysql.Predicate, value time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereDeletedAtIn(value []time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereDeletedAtNotIn(value []time.Time) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByDeletedAt(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "deleted_at "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereUserid(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereUseridIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereUseridNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByUserid(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "userid "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTicketnum(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ticketnum", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTicketnumIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ticketnum", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTicketnumNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ticketnum", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTicketnum(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ticketnum "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereCavefans(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cavefans", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCavefansIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cavefans", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCavefansNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cavefans", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByCavefans(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "cavefans "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterfans(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitterfans", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterfansIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitterfans", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterfansNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitterfans", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitterfans(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitterfans "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereLastlogin(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lastlogin", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereLastloginIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lastlogin", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereLastloginNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lastlogin", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByLastlogin(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "lastlogin "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereRegistime(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "registime", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereRegistimeIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "registime", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereRegistimeNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "registime", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByRegistime(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "registime "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereEarnedpoint(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "earnedpoint", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereEarnedpointIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "earnedpoint", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereEarnedpointNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "earnedpoint", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByEarnedpoint(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "earnedpoint "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereCavepoint(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cavepoint", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCavepointIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cavepoint", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCavepointNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cavepoint", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByCavepoint(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "cavepoint "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereViews(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "views", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereViewsIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "views", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereViewsNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "views", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByViews(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "views "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereNickName(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nick_name", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereNickNameIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nick_name", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereNickNameNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nick_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByNickName(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nick_name "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereBios(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "bios", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereBiosIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "bios", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereBiosNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "bios", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByBios(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "bios "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereAvatar(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereAvatarIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereAvatarNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByAvatar(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "avatar "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereAddress(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereAddressIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereAddressNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByAddress(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "address "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterId(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_id", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterIdIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_id", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterIdNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitterId(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitter_id "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterAvatar(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_avatar", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterAvatarIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_avatar", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterAvatarNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_avatar", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitterAvatar(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitter_avatar "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterName(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_name", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterNameIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_name", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterNameNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitterName(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitter_name "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereCaveretweeturl(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "caveretweeturl", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCaveretweeturlIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "caveretweeturl", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCaveretweeturlNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "caveretweeturl", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByCaveretweeturl(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "caveretweeturl "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereMintcave(p mysql.Predicate, value int32) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mintcave", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereMintcaveIn(value []int32) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mintcave", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereMintcaveNotIn(value []int32) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mintcave", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByMintcave(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "mintcave "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereNickname(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereNicknameIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereNicknameNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByNickname(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nickname "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterid(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitterid", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitteridIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitterid", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitteridNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitterid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitterid(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitterid "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitteravatar(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitteravatar", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitteravatarIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitteravatar", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitteravatarNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitteravatar", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitteravatar(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitteravatar "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwittername(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twittername", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitternameIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twittername", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitternameNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twittername", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwittername(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twittername "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTag(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "tag", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTagIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "tag", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTagNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "tag", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTag(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "tag "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereEarnedPoint(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "earned_point", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereEarnedPointIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "earned_point", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereEarnedPointNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "earned_point", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByEarnedPoint(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "earned_point "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereCavePoint(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_point", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCavePointIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_point", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCavePointNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_point", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByCavePoint(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "cave_point "+order)
	return qb
}
