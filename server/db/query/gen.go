// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	CncChecker *cncChecker
	CncStatus  *cncStatus
	Employee   *employee
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CncChecker = &Q.CncChecker
	CncStatus = &Q.CncStatus
	Employee = &Q.Employee
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		CncChecker: newCncChecker(db, opts...),
		CncStatus:  newCncStatus(db, opts...),
		Employee:   newEmployee(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CncChecker cncChecker
	CncStatus  cncStatus
	Employee   employee
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		CncChecker: q.CncChecker.clone(db),
		CncStatus:  q.CncStatus.clone(db),
		Employee:   q.Employee.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		CncChecker: q.CncChecker.replaceDB(db),
		CncStatus:  q.CncStatus.replaceDB(db),
		Employee:   q.Employee.replaceDB(db),
	}
}

type queryCtx struct {
	CncChecker ICncCheckerDo
	CncStatus  ICncStatusDo
	Employee   IEmployeeDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CncChecker: q.CncChecker.WithContext(ctx),
		CncStatus:  q.CncStatus.WithContext(ctx),
		Employee:   q.Employee.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
