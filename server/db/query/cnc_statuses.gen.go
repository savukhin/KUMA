// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"server/db/models"
)

func newCncStatus(db *gorm.DB, opts ...gen.DOOption) cncStatus {
	_cncStatus := cncStatus{}

	_cncStatus.cncStatusDo.UseDB(db, opts...)
	_cncStatus.cncStatusDo.UseModel(&models.CncStatus{})

	tableName := _cncStatus.cncStatusDo.TableName()
	_cncStatus.ALL = field.NewAsterisk(tableName)
	_cncStatus.ID = field.NewInt(tableName, "id")
	_cncStatus.StatusName = field.NewString(tableName, "status_name")

	_cncStatus.fillFieldMap()

	return _cncStatus
}

type cncStatus struct {
	cncStatusDo

	ALL        field.Asterisk
	ID         field.Int
	StatusName field.String

	fieldMap map[string]field.Expr
}

func (c cncStatus) Table(newTableName string) *cncStatus {
	c.cncStatusDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cncStatus) As(alias string) *cncStatus {
	c.cncStatusDo.DO = *(c.cncStatusDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cncStatus) updateTableName(table string) *cncStatus {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt(table, "id")
	c.StatusName = field.NewString(table, "status_name")

	c.fillFieldMap()

	return c
}

func (c *cncStatus) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cncStatus) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 2)
	c.fieldMap["id"] = c.ID
	c.fieldMap["status_name"] = c.StatusName
}

func (c cncStatus) clone(db *gorm.DB) cncStatus {
	c.cncStatusDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cncStatus) replaceDB(db *gorm.DB) cncStatus {
	c.cncStatusDo.ReplaceDB(db)
	return c
}

type cncStatusDo struct{ gen.DO }

type ICncStatusDo interface {
	gen.SubQuery
	Debug() ICncStatusDo
	WithContext(ctx context.Context) ICncStatusDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICncStatusDo
	WriteDB() ICncStatusDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICncStatusDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICncStatusDo
	Not(conds ...gen.Condition) ICncStatusDo
	Or(conds ...gen.Condition) ICncStatusDo
	Select(conds ...field.Expr) ICncStatusDo
	Where(conds ...gen.Condition) ICncStatusDo
	Order(conds ...field.Expr) ICncStatusDo
	Distinct(cols ...field.Expr) ICncStatusDo
	Omit(cols ...field.Expr) ICncStatusDo
	Join(table schema.Tabler, on ...field.Expr) ICncStatusDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICncStatusDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICncStatusDo
	Group(cols ...field.Expr) ICncStatusDo
	Having(conds ...gen.Condition) ICncStatusDo
	Limit(limit int) ICncStatusDo
	Offset(offset int) ICncStatusDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICncStatusDo
	Unscoped() ICncStatusDo
	Create(values ...*models.CncStatus) error
	CreateInBatches(values []*models.CncStatus, batchSize int) error
	Save(values ...*models.CncStatus) error
	First() (*models.CncStatus, error)
	Take() (*models.CncStatus, error)
	Last() (*models.CncStatus, error)
	Find() ([]*models.CncStatus, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.CncStatus, err error)
	FindInBatches(result *[]*models.CncStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.CncStatus) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICncStatusDo
	Assign(attrs ...field.AssignExpr) ICncStatusDo
	Joins(fields ...field.RelationField) ICncStatusDo
	Preload(fields ...field.RelationField) ICncStatusDo
	FirstOrInit() (*models.CncStatus, error)
	FirstOrCreate() (*models.CncStatus, error)
	FindByPage(offset int, limit int) (result []*models.CncStatus, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICncStatusDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cncStatusDo) Debug() ICncStatusDo {
	return c.withDO(c.DO.Debug())
}

func (c cncStatusDo) WithContext(ctx context.Context) ICncStatusDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cncStatusDo) ReadDB() ICncStatusDo {
	return c.Clauses(dbresolver.Read)
}

func (c cncStatusDo) WriteDB() ICncStatusDo {
	return c.Clauses(dbresolver.Write)
}

func (c cncStatusDo) Session(config *gorm.Session) ICncStatusDo {
	return c.withDO(c.DO.Session(config))
}

func (c cncStatusDo) Clauses(conds ...clause.Expression) ICncStatusDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cncStatusDo) Returning(value interface{}, columns ...string) ICncStatusDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cncStatusDo) Not(conds ...gen.Condition) ICncStatusDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cncStatusDo) Or(conds ...gen.Condition) ICncStatusDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cncStatusDo) Select(conds ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cncStatusDo) Where(conds ...gen.Condition) ICncStatusDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cncStatusDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICncStatusDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cncStatusDo) Order(conds ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cncStatusDo) Distinct(cols ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cncStatusDo) Omit(cols ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cncStatusDo) Join(table schema.Tabler, on ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cncStatusDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cncStatusDo) RightJoin(table schema.Tabler, on ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cncStatusDo) Group(cols ...field.Expr) ICncStatusDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cncStatusDo) Having(conds ...gen.Condition) ICncStatusDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cncStatusDo) Limit(limit int) ICncStatusDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cncStatusDo) Offset(offset int) ICncStatusDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cncStatusDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICncStatusDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cncStatusDo) Unscoped() ICncStatusDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cncStatusDo) Create(values ...*models.CncStatus) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cncStatusDo) CreateInBatches(values []*models.CncStatus, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cncStatusDo) Save(values ...*models.CncStatus) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cncStatusDo) First() (*models.CncStatus, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.CncStatus), nil
	}
}

func (c cncStatusDo) Take() (*models.CncStatus, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.CncStatus), nil
	}
}

func (c cncStatusDo) Last() (*models.CncStatus, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.CncStatus), nil
	}
}

func (c cncStatusDo) Find() ([]*models.CncStatus, error) {
	result, err := c.DO.Find()
	return result.([]*models.CncStatus), err
}

func (c cncStatusDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.CncStatus, err error) {
	buf := make([]*models.CncStatus, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cncStatusDo) FindInBatches(result *[]*models.CncStatus, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cncStatusDo) Attrs(attrs ...field.AssignExpr) ICncStatusDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cncStatusDo) Assign(attrs ...field.AssignExpr) ICncStatusDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cncStatusDo) Joins(fields ...field.RelationField) ICncStatusDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cncStatusDo) Preload(fields ...field.RelationField) ICncStatusDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cncStatusDo) FirstOrInit() (*models.CncStatus, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.CncStatus), nil
	}
}

func (c cncStatusDo) FirstOrCreate() (*models.CncStatus, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.CncStatus), nil
	}
}

func (c cncStatusDo) FindByPage(offset int, limit int) (result []*models.CncStatus, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cncStatusDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cncStatusDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cncStatusDo) Delete(models ...*models.CncStatus) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cncStatusDo) withDO(do gen.Dao) *cncStatusDo {
	c.DO = *do.(*gen.DO)
	return c
}
