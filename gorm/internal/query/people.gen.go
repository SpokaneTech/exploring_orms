// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"github.com/SpokaneTech/exploring_orms/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newPerson(db *gorm.DB, opts ...gen.DOOption) person {
	_person := person{}

	_person.personDo.UseDB(db, opts...)
	_person.personDo.UseModel(&models.Person{})

	tableName := _person.personDo.TableName()
	_person.ALL = field.NewAsterisk(tableName)
	_person.ID = field.NewUint(tableName, "id")
	_person.CreatedAt = field.NewTime(tableName, "created_at")
	_person.UpdatedAt = field.NewTime(tableName, "updated_at")
	_person.DeletedAt = field.NewField(tableName, "deleted_at")
	_person.Name = field.NewString(tableName, "name")

	_person.fillFieldMap()

	return _person
}

type person struct {
	personDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String

	fieldMap map[string]field.Expr
}

func (p person) Table(newTableName string) *person {
	p.personDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p person) As(alias string) *person {
	p.personDo.DO = *(p.personDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *person) updateTableName(table string) *person {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewUint(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Name = field.NewString(table, "name")

	p.fillFieldMap()

	return p
}

func (p *person) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *person) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 5)
	p.fieldMap["id"] = p.ID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["name"] = p.Name
}

func (p person) clone(db *gorm.DB) person {
	p.personDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p person) replaceDB(db *gorm.DB) person {
	p.personDo.ReplaceDB(db)
	return p
}

type personDo struct{ gen.DO }

type IPersonDo interface {
	gen.SubQuery
	Debug() IPersonDo
	WithContext(ctx context.Context) IPersonDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPersonDo
	WriteDB() IPersonDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPersonDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPersonDo
	Not(conds ...gen.Condition) IPersonDo
	Or(conds ...gen.Condition) IPersonDo
	Select(conds ...field.Expr) IPersonDo
	Where(conds ...gen.Condition) IPersonDo
	Order(conds ...field.Expr) IPersonDo
	Distinct(cols ...field.Expr) IPersonDo
	Omit(cols ...field.Expr) IPersonDo
	Join(table schema.Tabler, on ...field.Expr) IPersonDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPersonDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPersonDo
	Group(cols ...field.Expr) IPersonDo
	Having(conds ...gen.Condition) IPersonDo
	Limit(limit int) IPersonDo
	Offset(offset int) IPersonDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPersonDo
	Unscoped() IPersonDo
	Create(values ...*models.Person) error
	CreateInBatches(values []*models.Person, batchSize int) error
	Save(values ...*models.Person) error
	First() (*models.Person, error)
	Take() (*models.Person, error)
	Last() (*models.Person, error)
	Find() ([]*models.Person, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Person, err error)
	FindInBatches(result *[]*models.Person, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Person) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPersonDo
	Assign(attrs ...field.AssignExpr) IPersonDo
	Joins(fields ...field.RelationField) IPersonDo
	Preload(fields ...field.RelationField) IPersonDo
	FirstOrInit() (*models.Person, error)
	FirstOrCreate() (*models.Person, error)
	FindByPage(offset int, limit int) (result []*models.Person, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPersonDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p personDo) Debug() IPersonDo {
	return p.withDO(p.DO.Debug())
}

func (p personDo) WithContext(ctx context.Context) IPersonDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p personDo) ReadDB() IPersonDo {
	return p.Clauses(dbresolver.Read)
}

func (p personDo) WriteDB() IPersonDo {
	return p.Clauses(dbresolver.Write)
}

func (p personDo) Session(config *gorm.Session) IPersonDo {
	return p.withDO(p.DO.Session(config))
}

func (p personDo) Clauses(conds ...clause.Expression) IPersonDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p personDo) Returning(value interface{}, columns ...string) IPersonDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p personDo) Not(conds ...gen.Condition) IPersonDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p personDo) Or(conds ...gen.Condition) IPersonDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p personDo) Select(conds ...field.Expr) IPersonDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p personDo) Where(conds ...gen.Condition) IPersonDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p personDo) Order(conds ...field.Expr) IPersonDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p personDo) Distinct(cols ...field.Expr) IPersonDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p personDo) Omit(cols ...field.Expr) IPersonDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p personDo) Join(table schema.Tabler, on ...field.Expr) IPersonDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p personDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPersonDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p personDo) RightJoin(table schema.Tabler, on ...field.Expr) IPersonDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p personDo) Group(cols ...field.Expr) IPersonDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p personDo) Having(conds ...gen.Condition) IPersonDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p personDo) Limit(limit int) IPersonDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p personDo) Offset(offset int) IPersonDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p personDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPersonDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p personDo) Unscoped() IPersonDo {
	return p.withDO(p.DO.Unscoped())
}

func (p personDo) Create(values ...*models.Person) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p personDo) CreateInBatches(values []*models.Person, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p personDo) Save(values ...*models.Person) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p personDo) First() (*models.Person, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Person), nil
	}
}

func (p personDo) Take() (*models.Person, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Person), nil
	}
}

func (p personDo) Last() (*models.Person, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Person), nil
	}
}

func (p personDo) Find() ([]*models.Person, error) {
	result, err := p.DO.Find()
	return result.([]*models.Person), err
}

func (p personDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Person, err error) {
	buf := make([]*models.Person, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p personDo) FindInBatches(result *[]*models.Person, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p personDo) Attrs(attrs ...field.AssignExpr) IPersonDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p personDo) Assign(attrs ...field.AssignExpr) IPersonDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p personDo) Joins(fields ...field.RelationField) IPersonDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p personDo) Preload(fields ...field.RelationField) IPersonDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p personDo) FirstOrInit() (*models.Person, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Person), nil
	}
}

func (p personDo) FirstOrCreate() (*models.Person, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Person), nil
	}
}

func (p personDo) FindByPage(offset int, limit int) (result []*models.Person, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p personDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p personDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p personDo) Delete(models ...*models.Person) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *personDo) withDO(do gen.Dao) *personDo {
	p.DO = *do.(*gen.DO)
	return p
}
