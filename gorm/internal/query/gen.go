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
	Q            = new(Query)
	Manufacturer *manufacturer
	Model        *model
	Part         *part
	Person       *person
	Vehicle      *vehicle
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Manufacturer = &Q.Manufacturer
	Model = &Q.Model
	Part = &Q.Part
	Person = &Q.Person
	Vehicle = &Q.Vehicle
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:           db,
		Manufacturer: newManufacturer(db, opts...),
		Model:        newModel(db, opts...),
		Part:         newPart(db, opts...),
		Person:       newPerson(db, opts...),
		Vehicle:      newVehicle(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Manufacturer manufacturer
	Model        model
	Part         part
	Person       person
	Vehicle      vehicle
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:           db,
		Manufacturer: q.Manufacturer.clone(db),
		Model:        q.Model.clone(db),
		Part:         q.Part.clone(db),
		Person:       q.Person.clone(db),
		Vehicle:      q.Vehicle.clone(db),
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
		db:           db,
		Manufacturer: q.Manufacturer.replaceDB(db),
		Model:        q.Model.replaceDB(db),
		Part:         q.Part.replaceDB(db),
		Person:       q.Person.replaceDB(db),
		Vehicle:      q.Vehicle.replaceDB(db),
	}
}

type queryCtx struct {
	Manufacturer IManufacturerDo
	Model        IModelDo
	Part         IPartDo
	Person       IPersonDo
	Vehicle      IVehicleDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Manufacturer: q.Manufacturer.WithContext(ctx),
		Model:        q.Model.WithContext(ctx),
		Part:         q.Part.WithContext(ctx),
		Person:       q.Person.WithContext(ctx),
		Vehicle:      q.Vehicle.WithContext(ctx),
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
