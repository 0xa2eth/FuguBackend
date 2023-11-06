///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package secret

import (
	"fmt"

	"FuguBackend/app/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Secret {
	return new(Secret)
}

func NewQueryBuilder() *secretQueryBuilder {
	return new(secretQueryBuilder)
}

func (t *Secret) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type secretQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *secretQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *secretQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Secret{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *secretQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Secret{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *secretQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Secret{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *secretQueryBuilder) First(db *gorm.DB) (*Secret, error) {
	ret := &Secret{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *secretQueryBuilder) QueryOne(db *gorm.DB) (*Secret, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *secretQueryBuilder) QueryAll(db *gorm.DB) ([]*Secret, error) {
	var ret []*Secret
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *secretQueryBuilder) Limit(limit int) *secretQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *secretQueryBuilder) Offset(offset int) *secretQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *secretQueryBuilder) WhereId(p mysql.Predicate, value int32) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereIdIn(value []int32) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereIdNotIn(value []int32) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) OrderById(asc bool) *secretQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *secretQueryBuilder) WhereAuthorid(p mysql.Predicate, value int32) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "authorid", p),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereAuthoridIn(value []int32) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "authorid", "IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereAuthoridNotIn(value []int32) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "authorid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) OrderByAuthorid(asc bool) *secretQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "authorid "+order)
	return qb
}

func (qb *secretQueryBuilder) WhereContent(p mysql.Predicate, value string) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereContentIn(value []string) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereContentNotIn(value []string) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) OrderByContent(asc bool) *secretQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *secretQueryBuilder) WhereTimestamp(p mysql.Predicate, value int64) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "timestamp", p),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereTimestampIn(value []int64) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "timestamp", "IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereTimestampNotIn(value []int64) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "timestamp", "NOT IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) OrderByTimestamp(asc bool) *secretQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "timestamp "+order)
	return qb
}

func (qb *secretQueryBuilder) WhereViews(p mysql.Predicate, value int64) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "views", p),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereViewsIn(value []int64) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "views", "IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) WhereViewsNotIn(value []int64) *secretQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "views", "NOT IN"),
		value,
	})
	return qb
}

func (qb *secretQueryBuilder) OrderByViews(asc bool) *secretQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "views "+order)
	return qb
}