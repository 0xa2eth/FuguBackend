///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package fts

import (
	"fmt"
	"time"

	"FuguBackend/app/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Fts {
	return new(Fts)
}

func NewQueryBuilder() *ftsQueryBuilder {
	return new(ftsQueryBuilder)
}

func (t *Fts) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type ftsQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *ftsQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *ftsQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Fts{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *ftsQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Fts{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *ftsQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Fts{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *ftsQueryBuilder) First(db *gorm.DB) (*Fts, error) {
	ret := &Fts{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *ftsQueryBuilder) QueryOne(db *gorm.DB) (*Fts, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *ftsQueryBuilder) QueryAll(db *gorm.DB) ([]*Fts, error) {
	var ret []*Fts
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *ftsQueryBuilder) Limit(limit int) *ftsQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *ftsQueryBuilder) Offset(offset int) *ftsQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *ftsQueryBuilder) WhereId(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereIdIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereIdNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderById(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereCreatedAtIn(value []time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByCreatedAt(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereUpdatedAtIn(value []time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByUpdatedAt(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereDeletedAt(p mysql.Predicate, value time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereDeletedAtIn(value []time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereDeletedAtNotIn(value []time.Time) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByDeletedAt(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "deleted_at "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereUserId(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereUserIdIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereUserIdNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByUserId(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_id "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereChainId(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "chain_id", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereChainIdIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "chain_id", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereChainIdNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "chain_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByChainId(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "chain_id "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereTokenNum(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "token_num", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereTokenNumIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "token_num", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereTokenNumNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "token_num", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByTokenNum(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "token_num "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereAddress(p mysql.Predicate, value string) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereAddressIn(value []string) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereAddressNotIn(value []string) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByAddress(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "address "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereUserid(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereUseridIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereUseridNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByUserid(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "userid "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereChainid(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "chainid", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereChainidIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "chainid", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereChainidNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "chainid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByChainid(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "chainid "+order)
	return qb
}

func (qb *ftsQueryBuilder) WhereTokennum(p mysql.Predicate, value int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "tokennum", p),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereTokennumIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "tokennum", "IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) WhereTokennumNotIn(value []int64) *ftsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "tokennum", "NOT IN"),
		value,
	})
	return qb
}

func (qb *ftsQueryBuilder) OrderByTokennum(asc bool) *ftsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "tokennum "+order)
	return qb
}
