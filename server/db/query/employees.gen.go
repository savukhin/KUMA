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

func newEmployee(db *gorm.DB, opts ...gen.DOOption) employee {
	_employee := employee{}

	_employee.employeeDo.UseDB(db, opts...)
	_employee.employeeDo.UseModel(&models.Employee{})

	tableName := _employee.employeeDo.TableName()
	_employee.ALL = field.NewAsterisk(tableName)
	_employee.ID = field.NewInt(tableName, "id")
	_employee.Name = field.NewString(tableName, "name")
	_employee.TelegramUserName = field.NewString(tableName, "telegram_user_name")
	_employee.CheckedIN = field.NewBool(tableName, "checked_in")
	_employee.Username = field.NewString(tableName, "username")
	_employee.PasswordHash = field.NewString(tableName, "password_hash")
	_employee.CreatedAt = field.NewUint64(tableName, "created_at")
	_employee.UpdatedAt = field.NewUint64(tableName, "updated_at")
	_employee.DeletedAt = field.NewField(tableName, "deleted_at")

	_employee.fillFieldMap()

	return _employee
}

type employee struct {
	employeeDo

	ALL              field.Asterisk
	ID               field.Int
	Name             field.String
	TelegramUserName field.String
	CheckedIN        field.Bool
	Username         field.String
	PasswordHash     field.String
	CreatedAt        field.Uint64
	UpdatedAt        field.Uint64
	DeletedAt        field.Field

	fieldMap map[string]field.Expr
}

func (e employee) Table(newTableName string) *employee {
	e.employeeDo.UseTable(newTableName)
	return e.updateTableName(newTableName)
}

func (e employee) As(alias string) *employee {
	e.employeeDo.DO = *(e.employeeDo.As(alias).(*gen.DO))
	return e.updateTableName(alias)
}

func (e *employee) updateTableName(table string) *employee {
	e.ALL = field.NewAsterisk(table)
	e.ID = field.NewInt(table, "id")
	e.Name = field.NewString(table, "name")
	e.TelegramUserName = field.NewString(table, "telegram_user_name")
	e.CheckedIN = field.NewBool(table, "checked_in")
	e.Username = field.NewString(table, "username")
	e.PasswordHash = field.NewString(table, "password_hash")
	e.CreatedAt = field.NewUint64(table, "created_at")
	e.UpdatedAt = field.NewUint64(table, "updated_at")
	e.DeletedAt = field.NewField(table, "deleted_at")

	e.fillFieldMap()

	return e
}

func (e *employee) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := e.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (e *employee) fillFieldMap() {
	e.fieldMap = make(map[string]field.Expr, 9)
	e.fieldMap["id"] = e.ID
	e.fieldMap["name"] = e.Name
	e.fieldMap["telegram_user_name"] = e.TelegramUserName
	e.fieldMap["checked_in"] = e.CheckedIN
	e.fieldMap["username"] = e.Username
	e.fieldMap["password_hash"] = e.PasswordHash
	e.fieldMap["created_at"] = e.CreatedAt
	e.fieldMap["updated_at"] = e.UpdatedAt
	e.fieldMap["deleted_at"] = e.DeletedAt
}

func (e employee) clone(db *gorm.DB) employee {
	e.employeeDo.ReplaceConnPool(db.Statement.ConnPool)
	return e
}

func (e employee) replaceDB(db *gorm.DB) employee {
	e.employeeDo.ReplaceDB(db)
	return e
}

type employeeDo struct{ gen.DO }

type IEmployeeDo interface {
	gen.SubQuery
	Debug() IEmployeeDo
	WithContext(ctx context.Context) IEmployeeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IEmployeeDo
	WriteDB() IEmployeeDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IEmployeeDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IEmployeeDo
	Not(conds ...gen.Condition) IEmployeeDo
	Or(conds ...gen.Condition) IEmployeeDo
	Select(conds ...field.Expr) IEmployeeDo
	Where(conds ...gen.Condition) IEmployeeDo
	Order(conds ...field.Expr) IEmployeeDo
	Distinct(cols ...field.Expr) IEmployeeDo
	Omit(cols ...field.Expr) IEmployeeDo
	Join(table schema.Tabler, on ...field.Expr) IEmployeeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo
	Group(cols ...field.Expr) IEmployeeDo
	Having(conds ...gen.Condition) IEmployeeDo
	Limit(limit int) IEmployeeDo
	Offset(offset int) IEmployeeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeDo
	Unscoped() IEmployeeDo
	Create(values ...*models.Employee) error
	CreateInBatches(values []*models.Employee, batchSize int) error
	Save(values ...*models.Employee) error
	First() (*models.Employee, error)
	Take() (*models.Employee, error)
	Last() (*models.Employee, error)
	Find() ([]*models.Employee, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Employee, err error)
	FindInBatches(result *[]*models.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Employee) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IEmployeeDo
	Assign(attrs ...field.AssignExpr) IEmployeeDo
	Joins(fields ...field.RelationField) IEmployeeDo
	Preload(fields ...field.RelationField) IEmployeeDo
	FirstOrInit() (*models.Employee, error)
	FirstOrCreate() (*models.Employee, error)
	FindByPage(offset int, limit int) (result []*models.Employee, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IEmployeeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (e employeeDo) Debug() IEmployeeDo {
	return e.withDO(e.DO.Debug())
}

func (e employeeDo) WithContext(ctx context.Context) IEmployeeDo {
	return e.withDO(e.DO.WithContext(ctx))
}

func (e employeeDo) ReadDB() IEmployeeDo {
	return e.Clauses(dbresolver.Read)
}

func (e employeeDo) WriteDB() IEmployeeDo {
	return e.Clauses(dbresolver.Write)
}

func (e employeeDo) Session(config *gorm.Session) IEmployeeDo {
	return e.withDO(e.DO.Session(config))
}

func (e employeeDo) Clauses(conds ...clause.Expression) IEmployeeDo {
	return e.withDO(e.DO.Clauses(conds...))
}

func (e employeeDo) Returning(value interface{}, columns ...string) IEmployeeDo {
	return e.withDO(e.DO.Returning(value, columns...))
}

func (e employeeDo) Not(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Not(conds...))
}

func (e employeeDo) Or(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Or(conds...))
}

func (e employeeDo) Select(conds ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Select(conds...))
}

func (e employeeDo) Where(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Where(conds...))
}

func (e employeeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IEmployeeDo {
	return e.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (e employeeDo) Order(conds ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Order(conds...))
}

func (e employeeDo) Distinct(cols ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Distinct(cols...))
}

func (e employeeDo) Omit(cols ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Omit(cols...))
}

func (e employeeDo) Join(table schema.Tabler, on ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Join(table, on...))
}

func (e employeeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.LeftJoin(table, on...))
}

func (e employeeDo) RightJoin(table schema.Tabler, on ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.RightJoin(table, on...))
}

func (e employeeDo) Group(cols ...field.Expr) IEmployeeDo {
	return e.withDO(e.DO.Group(cols...))
}

func (e employeeDo) Having(conds ...gen.Condition) IEmployeeDo {
	return e.withDO(e.DO.Having(conds...))
}

func (e employeeDo) Limit(limit int) IEmployeeDo {
	return e.withDO(e.DO.Limit(limit))
}

func (e employeeDo) Offset(offset int) IEmployeeDo {
	return e.withDO(e.DO.Offset(offset))
}

func (e employeeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IEmployeeDo {
	return e.withDO(e.DO.Scopes(funcs...))
}

func (e employeeDo) Unscoped() IEmployeeDo {
	return e.withDO(e.DO.Unscoped())
}

func (e employeeDo) Create(values ...*models.Employee) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Create(values)
}

func (e employeeDo) CreateInBatches(values []*models.Employee, batchSize int) error {
	return e.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (e employeeDo) Save(values ...*models.Employee) error {
	if len(values) == 0 {
		return nil
	}
	return e.DO.Save(values)
}

func (e employeeDo) First() (*models.Employee, error) {
	if result, err := e.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Employee), nil
	}
}

func (e employeeDo) Take() (*models.Employee, error) {
	if result, err := e.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Employee), nil
	}
}

func (e employeeDo) Last() (*models.Employee, error) {
	if result, err := e.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Employee), nil
	}
}

func (e employeeDo) Find() ([]*models.Employee, error) {
	result, err := e.DO.Find()
	return result.([]*models.Employee), err
}

func (e employeeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Employee, err error) {
	buf := make([]*models.Employee, 0, batchSize)
	err = e.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (e employeeDo) FindInBatches(result *[]*models.Employee, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return e.DO.FindInBatches(result, batchSize, fc)
}

func (e employeeDo) Attrs(attrs ...field.AssignExpr) IEmployeeDo {
	return e.withDO(e.DO.Attrs(attrs...))
}

func (e employeeDo) Assign(attrs ...field.AssignExpr) IEmployeeDo {
	return e.withDO(e.DO.Assign(attrs...))
}

func (e employeeDo) Joins(fields ...field.RelationField) IEmployeeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Joins(_f))
	}
	return &e
}

func (e employeeDo) Preload(fields ...field.RelationField) IEmployeeDo {
	for _, _f := range fields {
		e = *e.withDO(e.DO.Preload(_f))
	}
	return &e
}

func (e employeeDo) FirstOrInit() (*models.Employee, error) {
	if result, err := e.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Employee), nil
	}
}

func (e employeeDo) FirstOrCreate() (*models.Employee, error) {
	if result, err := e.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Employee), nil
	}
}

func (e employeeDo) FindByPage(offset int, limit int) (result []*models.Employee, count int64, err error) {
	result, err = e.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = e.Offset(-1).Limit(-1).Count()
	return
}

func (e employeeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = e.Count()
	if err != nil {
		return
	}

	err = e.Offset(offset).Limit(limit).Scan(result)
	return
}

func (e employeeDo) Scan(result interface{}) (err error) {
	return e.DO.Scan(result)
}

func (e employeeDo) Delete(models ...*models.Employee) (result gen.ResultInfo, err error) {
	return e.DO.Delete(models)
}

func (e *employeeDo) withDO(do gen.Dao) *employeeDo {
	e.DO = *do.(*gen.DO)
	return e
}
