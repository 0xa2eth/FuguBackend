package secrets

import (
	"context"
	"time"

	"FuguBackend/app/pkg/model"
	"FuguBackend/pkg/errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IExpandSecretModel = interface {
	Insert(data *Mission) error
	InsertWithClauses(data *Mission, clauses ...clause.Expression) error
	InsertBatch(dataset []*Mission) error
	InsertBatchWithClauses(dataset []*Mission, clauses ...clause.Expression) (int64, error)
	Update(data *Mission, selects ...string) error
	UpdateByCondition(data *Mission, condition func(*gorm.DB) *gorm.DB) error
	UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error
	FindOne(id int64) (*Mission, error)
	FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*Mission, error)
	FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*Mission, error)
	FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error)
	Paginate(pq *model.PaginationQuery) ([]*Mission, int64, error)
	PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error)
	PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error)
	CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error)
	DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error
	Init() error
}

type Mission struct {
	Id          int64          `json:"id" gorm:"primaryKey;column:id;"`
	Created_at  time.Time      `json:"created_at" gorm:"column:created_at;datetime(3)"`
	Updated_at  time.Time      `json:"updated_at" gorm:"column:updated_at;datetime(3)"`
	Deleted_at  gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;datetime(3)"`
	Projectid   int64          `json:"projectid" gorm:"column:projectid;bigint"`
	Missionid   int64          `json:"missionid" gorm:"column:missionid;bigint"`
	Title       string         `json:"title" gorm:"column:title;varchar(255)"`
	Description string         `json:"description" gorm:"column:description;varchar(255)"`
	Starttime   time.Time      `json:"starttime" gorm:"column:starttime;timestamp"`
	Endtime     time.Time      `json:"endtime" gorm:"column:endtime;timestamp"`
	Type        string         `json:"type" gorm:"column:type;varchar(255)"`
	Viewable    bool           `json:"viewable" gorm:"column:viewable;tinyint(1)"`
	Chainid     int            `json:"chainid" gorm:"column:chainid;int"`
	Contract    string         `json:"contract" gorm:"column:contract;varchar(255)"`
	Banner      string         `json:"banner" gorm:"column:banner;varchar(255)"`
	Status      string         `json:"status" gorm:"column:status;enum('','Ended','Active','Not Start')"`
	Reward_type int            `json:"reward_type" gorm:"column:reward_type;tinyint(1)"`
}

// TableName 返回用户表信息的数据库表名
func (o *Mission) TableName() string {
	return "missions"
}

// GetId 获取id
func (o *Mission) GetId() int64 {
	return o.Id
}

// NewMintModel 新建用户表模型
func NewMissionModel(ctx context.Context, db *gorm.DB) IExpandSecretModel {
	return &defaultMissionModel{
		BaseModel: model.NewBaseModel(ctx, db),
	}
}

// defaultMintModel 默认用户表模型
type defaultMissionModel struct {
	*model.BaseModel
}

// Insert 插入Mission表信息
func (o *defaultMissionModel) Insert(data *Mission) error {
	return o.DB.Create(data).Error
}

// InsertWithClauses 使用子句插入Mission表信息
func (o *defaultMissionModel) InsertWithClauses(data *Mission, clauses ...clause.Expression) error {
	return o.DB.Clauses(clauses...).Create(data).Error
}

// InsertBatch 批量插入Mission表信息
func (o *defaultMissionModel) InsertBatch(dataset []*Mission) error {
	return o.DB.Create(&dataset).Error
}

// InsertBatchWithClauses 使用子句批量插入Mission表信息
func (o *defaultMissionModel) InsertBatchWithClauses(dataset []*Mission, clauses ...clause.Expression) (int64, error) {
	db := o.DB.Clauses(clauses...).Create(&dataset)

	return db.RowsAffected, db.Error
}

// Update 更新mint表信息
func (o *defaultMissionModel) Update(data *Mission, selects ...string) error {
	return o.BaseModel.Update(data, selects)
}

// UpdateByCondition 通过动态条件更新mint表信息
func (o *defaultMissionModel) UpdateByCondition(data *Mission, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateByCondition(data, condition)
}

// UpdateWithMapByCondition 使用map通过动态条件更新Mission表信息
func (o *defaultMissionModel) UpdateWithMapByCondition(data map[string]interface{}, condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.UpdateWithMapByCondition(&Mission{}, data, condition)
}

// FindOne 通过用户表id查找一个Mission表信息
func (m *defaultMissionModel) FindOne(id int64) (*Mission, error) {
	if id < 1 {
		return nil, errors.New("ErrInvalidParams")
	}

	var o Mission
	err := m.DB.Where("`id` = ?", id).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("Mission not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindOneByCondition 通过动态条件查找一个Mission表信息
func (m *defaultMissionModel) FindOneByCondition(condition func(*gorm.DB) *gorm.DB) (*Mission, error) {
	var o Mission
	err := m.DB.Scopes(condition).First(&o).Error
	if err != nil {
		if model.IsRecordNotFound(err) {
			return nil, errors.New("Mission not exist")
		}
		return nil, err
	}

	return &o, nil
}

// FindByCondition 通过动态条件查找Mission表信息
func (o *defaultMissionModel) FindByCondition(condition func(*gorm.DB) *gorm.DB) ([]*Mission, error) {
	var op []*Mission
	err := o.DB.Scopes(condition).Find(&op).Error
	if err != nil {
		return nil, err
	}

	return op, nil
}

func (o *defaultMissionModel) FindMapByCondition(db *gorm.DB) ([]map[string]interface{}, error) {
	var op []map[string]interface{}
	err := db.Find(&op).Error
	if err != nil {
		return op, err
	}
	return op, nil
}

// Paginate 分页查找Mission表信息
func (o *defaultMissionModel) Paginate(pq *model.PaginationQuery) ([]*Mission, int64, error) {
	var total int64
	err := o.DB.Scopes(pq.Queries()).Model(&Mission{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var op []*Mission
	err = o.DB.Scopes(pq.Paginate()).Find(&op).Error
	if err != nil {
		return nil, 0, err
	}

	return op, total, nil
}

// PluckInt64s 通过动态条件查找单个列并将结果扫描到int64切片中
func (o *defaultMissionModel) PluckInt64s(column string, condition func(*gorm.DB) *gorm.DB) ([]int64, error) {
	return o.BaseModel.PluckInt64s(&Mission{}, column, condition)
}

// PluckStrings 通过动态条件查找单个列并将结果扫描到string切片中
func (o *defaultMissionModel) PluckStrings(column string, condition func(*gorm.DB) *gorm.DB) ([]string, error) {
	return o.BaseModel.PluckStrings(&Mission{}, column, condition)
}

// CountByCondition 通过动态条件计数用户表信息
func (o *defaultMissionModel) CountByCondition(condition func(*gorm.DB) *gorm.DB) (int64, error) {
	return o.BaseModel.CountByCondition(&Mission{}, condition)
}

// DeleteByCondition 通过动态条件删除用户表信息
func (o *defaultMissionModel) DeleteByCondition(condition func(*gorm.DB) *gorm.DB) error {
	return o.BaseModel.DeleteByCondition(&Mission{}, condition)
}

// Init 初始化Mission表信息
func (o *defaultMissionModel) Init() error {
	var num int64
	err := o.DB.Model(&Mission{}).Count(&num).Error
	if err != nil {
		return err
	}

	if num == 0 {
		// 执行初始化方法
		return nil
	}
	return nil
}
