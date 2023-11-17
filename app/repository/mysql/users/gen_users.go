///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package users

import (
	"fmt"
	"time"

	"FuguBackend/app/repository/mysql"

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
	return t.Id, nil
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

func (qb *usersQueryBuilder) WhereUserId(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereUserIdIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereUserIdNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByUserId(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_id "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTicketNum(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ticket_num", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTicketNumIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ticket_num", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTicketNumNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ticket_num", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTicketNum(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ticket_num "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereCaveFans(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_fans", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCaveFansIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_fans", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereCaveFansNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_fans", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByCaveFans(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "cave_fans "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterFans(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_fans", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterFansIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_fans", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereTwitterFansNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "twitter_fans", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByTwitterFans(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "twitter_fans "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereLastLogin(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "last_login", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereLastLoginIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "last_login", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereLastLoginNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "last_login", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByLastLogin(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "last_login "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereRegisTime(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "regis_time", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereRegisTimeIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "regis_time", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereRegisTimeNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "regis_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByRegisTime(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "regis_time "+order)
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

func (qb *usersQueryBuilder) WhereVipLevel(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "vip_level", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereVipLevelIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "vip_level", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereVipLevelNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "vip_level", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByVipLevel(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "vip_level "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereInvitedBy(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "invited_by", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereInvitedByIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "invited_by", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereInvitedByNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "invited_by", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByInvitedBy(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "invited_by "+order)
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

func (qb *usersQueryBuilder) WhereRetweetUrl(p mysql.Predicate, value string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retweet_url", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereRetweetUrlIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retweet_url", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereRetweetUrlNotIn(value []string) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "retweet_url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByRetweetUrl(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "retweet_url "+order)
	return qb
}

func (qb *usersQueryBuilder) WhereMintCave(p mysql.Predicate, value int32) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mint_cave", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereMintCaveIn(value []int32) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mint_cave", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WhereMintCaveNotIn(value []int32) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "mint_cave", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByMintCave(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "mint_cave "+order)
	return qb
}

func (qb *usersQueryBuilder) WherePosts(p mysql.Predicate, value int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "posts", p),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WherePostsIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "posts", "IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) WherePostsNotIn(value []int64) *usersQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "posts", "NOT IN"),
		value,
	})
	return qb
}

func (qb *usersQueryBuilder) OrderByPosts(asc bool) *usersQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "posts "+order)
	return qb
}
