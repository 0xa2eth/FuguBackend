///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package task_records

import (
	"fmt"

	"FuguBackend/app/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *TaskRecords {
	return new(TaskRecords)
}

func NewQueryBuilder() *taskRecordsQueryBuilder {
	return new(taskRecordsQueryBuilder)
}

func (t *TaskRecords) Create(db *gorm.DB) (id int64, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type taskRecordsQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *taskRecordsQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *taskRecordsQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&TaskRecords{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *taskRecordsQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&TaskRecords{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *taskRecordsQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&TaskRecords{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *taskRecordsQueryBuilder) First(db *gorm.DB) (*TaskRecords, error) {
	ret := &TaskRecords{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *taskRecordsQueryBuilder) QueryOne(db *gorm.DB) (*TaskRecords, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *taskRecordsQueryBuilder) QueryAll(db *gorm.DB) ([]*TaskRecords, error) {
	var ret []*TaskRecords
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *taskRecordsQueryBuilder) Limit(limit int) *taskRecordsQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *taskRecordsQueryBuilder) Offset(offset int) *taskRecordsQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereId(p mysql.Predicate, value int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereIdIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereIdNotIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderById(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCreatedAtIn(value []time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByCreatedAt(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUpdatedAtIn(value []time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByUpdatedAt(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereDeletedAt(p mysql.Predicate, value time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereDeletedAtIn(value []time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereDeletedAtNotIn(value []time.Time) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "deleted_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByDeletedAt(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "deleted_at "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCaveId(p mysql.Predicate, value string) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_id", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCaveIdIn(value []string) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_id", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCaveIdNotIn(value []string) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "cave_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByCaveId(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "cave_id "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUserId(p mysql.Predicate, value int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUserIdIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUserIdNotIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByUserId(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_id "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCaveid(p mysql.Predicate, value int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "caveid", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCaveidIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "caveid", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereCaveidNotIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "caveid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByCaveid(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "caveid "+order)
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUserid(p mysql.Predicate, value int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", p),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUseridIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) WhereUseridNotIn(value []int64) *taskRecordsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "userid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *taskRecordsQueryBuilder) OrderByUserid(asc bool) *taskRecordsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "userid "+order)
	return qb
}
