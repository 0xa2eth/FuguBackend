///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package invite_codes

import (
	"fmt"
	"time"

	"FuguBackend/app/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *InviteCodes {
	return new(InviteCodes)
}

func NewQueryBuilder() *inviteCodesQueryBuilder {
	return new(inviteCodesQueryBuilder)
}

func (t *InviteCodes) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type inviteCodesQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *inviteCodesQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *inviteCodesQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&InviteCodes{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *inviteCodesQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&InviteCodes{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *inviteCodesQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&InviteCodes{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *inviteCodesQueryBuilder) First(db *gorm.DB) (*InviteCodes, error) {
	ret := &InviteCodes{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *inviteCodesQueryBuilder) QueryOne(db *gorm.DB) (*InviteCodes, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *inviteCodesQueryBuilder) QueryAll(db *gorm.DB) ([]*InviteCodes, error) {
	var ret []*InviteCodes
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *inviteCodesQueryBuilder) Limit(limit int) *inviteCodesQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *inviteCodesQueryBuilder) Offset(offset int) *inviteCodesQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereId(p mysql.Predicate, value int64) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereIdIn(value []int64) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereIdNotIn(value []int64) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) OrderById(asc bool) *inviteCodesQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereCreatedAtIn(value []time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) OrderByCreatedAt(asc bool) *inviteCodesQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereUpdatedAtIn(value []time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) OrderByUpdatedAt(asc bool) *inviteCodesQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereDeletedAt(p mysql.Predicate, value time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereDeletedAtIn(value []time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereDeletedAtNotIn(value []time.Time) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) OrderByDeletedAt(asc bool) *inviteCodesQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "deleted_at "+order)
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereUserid(p mysql.Predicate, value int64) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", p),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereUseridIn(value []int64) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereUseridNotIn(value []int64) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) OrderByUserid(asc bool) *inviteCodesQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "userid "+order)
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereCode(p mysql.Predicate, value string) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "code", p),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereCodeIn(value []string) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "code", "IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) WhereCodeNotIn(value []string) *inviteCodesQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "code", "NOT IN"),
		value,
	})
	return qb
}

func (qb *inviteCodesQueryBuilder) OrderByCode(asc bool) *inviteCodesQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "code "+order)
	return qb
}
