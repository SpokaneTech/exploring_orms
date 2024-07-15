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

func newVehicle(db *gorm.DB, opts ...gen.DOOption) vehicle {
	_vehicle := vehicle{}

	_vehicle.vehicleDo.UseDB(db, opts...)
	_vehicle.vehicleDo.UseModel(&models.Vehicle{})

	tableName := _vehicle.vehicleDo.TableName()
	_vehicle.ALL = field.NewAsterisk(tableName)
	_vehicle.ID = field.NewUint(tableName, "id")
	_vehicle.CreatedAt = field.NewTime(tableName, "created_at")
	_vehicle.UpdatedAt = field.NewTime(tableName, "updated_at")
	_vehicle.DeletedAt = field.NewField(tableName, "deleted_at")
	_vehicle.Vin = field.NewString(tableName, "vin")
	_vehicle.VehicleModelID = field.NewUint(tableName, "vehicle_model_id")
	_vehicle.PersonID = field.NewInt(tableName, "person_id")
	_vehicle.VehicleModel = vehicleBelongsToVehicleModel{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("VehicleModel", "models.Model"),
		Manufacturer: struct {
			field.RelationField
			Vehicles struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("VehicleModel.Manufacturer", "models.Manufacturer"),
			Vehicles: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("VehicleModel.Manufacturer.Vehicles", "models.Model"),
			},
		},
		Parts: struct {
			field.RelationField
			Models struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("VehicleModel.Parts", "models.Part"),
			Models: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("VehicleModel.Parts.Models", "models.Model"),
			},
		},
	}

	_vehicle.Person = vehicleBelongsToPerson{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Person", "models.Person"),
	}

	_vehicle.fillFieldMap()

	return _vehicle
}

type vehicle struct {
	vehicleDo

	ALL            field.Asterisk
	ID             field.Uint
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Vin            field.String
	VehicleModelID field.Uint
	PersonID       field.Int
	VehicleModel   vehicleBelongsToVehicleModel

	Person vehicleBelongsToPerson

	fieldMap map[string]field.Expr
}

func (v vehicle) Table(newTableName string) *vehicle {
	v.vehicleDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vehicle) As(alias string) *vehicle {
	v.vehicleDo.DO = *(v.vehicleDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vehicle) updateTableName(table string) *vehicle {
	v.ALL = field.NewAsterisk(table)
	v.ID = field.NewUint(table, "id")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.DeletedAt = field.NewField(table, "deleted_at")
	v.Vin = field.NewString(table, "vin")
	v.VehicleModelID = field.NewUint(table, "vehicle_model_id")
	v.PersonID = field.NewInt(table, "person_id")

	v.fillFieldMap()

	return v
}

func (v *vehicle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vehicle) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 9)
	v.fieldMap["id"] = v.ID
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["deleted_at"] = v.DeletedAt
	v.fieldMap["vin"] = v.Vin
	v.fieldMap["vehicle_model_id"] = v.VehicleModelID
	v.fieldMap["person_id"] = v.PersonID

}

func (v vehicle) clone(db *gorm.DB) vehicle {
	v.vehicleDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vehicle) replaceDB(db *gorm.DB) vehicle {
	v.vehicleDo.ReplaceDB(db)
	return v
}

type vehicleBelongsToVehicleModel struct {
	db *gorm.DB

	field.RelationField

	Manufacturer struct {
		field.RelationField
		Vehicles struct {
			field.RelationField
		}
	}
	Parts struct {
		field.RelationField
		Models struct {
			field.RelationField
		}
	}
}

func (a vehicleBelongsToVehicleModel) Where(conds ...field.Expr) *vehicleBelongsToVehicleModel {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a vehicleBelongsToVehicleModel) WithContext(ctx context.Context) *vehicleBelongsToVehicleModel {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vehicleBelongsToVehicleModel) Session(session *gorm.Session) *vehicleBelongsToVehicleModel {
	a.db = a.db.Session(session)
	return &a
}

func (a vehicleBelongsToVehicleModel) Model(m *models.Vehicle) *vehicleBelongsToVehicleModelTx {
	return &vehicleBelongsToVehicleModelTx{a.db.Model(m).Association(a.Name())}
}

type vehicleBelongsToVehicleModelTx struct{ tx *gorm.Association }

func (a vehicleBelongsToVehicleModelTx) Find() (result *models.Model, err error) {
	return result, a.tx.Find(&result)
}

func (a vehicleBelongsToVehicleModelTx) Append(values ...*models.Model) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vehicleBelongsToVehicleModelTx) Replace(values ...*models.Model) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vehicleBelongsToVehicleModelTx) Delete(values ...*models.Model) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vehicleBelongsToVehicleModelTx) Clear() error {
	return a.tx.Clear()
}

func (a vehicleBelongsToVehicleModelTx) Count() int64 {
	return a.tx.Count()
}

type vehicleBelongsToPerson struct {
	db *gorm.DB

	field.RelationField
}

func (a vehicleBelongsToPerson) Where(conds ...field.Expr) *vehicleBelongsToPerson {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a vehicleBelongsToPerson) WithContext(ctx context.Context) *vehicleBelongsToPerson {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vehicleBelongsToPerson) Session(session *gorm.Session) *vehicleBelongsToPerson {
	a.db = a.db.Session(session)
	return &a
}

func (a vehicleBelongsToPerson) Model(m *models.Vehicle) *vehicleBelongsToPersonTx {
	return &vehicleBelongsToPersonTx{a.db.Model(m).Association(a.Name())}
}

type vehicleBelongsToPersonTx struct{ tx *gorm.Association }

func (a vehicleBelongsToPersonTx) Find() (result *models.Person, err error) {
	return result, a.tx.Find(&result)
}

func (a vehicleBelongsToPersonTx) Append(values ...*models.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vehicleBelongsToPersonTx) Replace(values ...*models.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vehicleBelongsToPersonTx) Delete(values ...*models.Person) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vehicleBelongsToPersonTx) Clear() error {
	return a.tx.Clear()
}

func (a vehicleBelongsToPersonTx) Count() int64 {
	return a.tx.Count()
}

type vehicleDo struct{ gen.DO }

type IVehicleDo interface {
	gen.SubQuery
	Debug() IVehicleDo
	WithContext(ctx context.Context) IVehicleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVehicleDo
	WriteDB() IVehicleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVehicleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVehicleDo
	Not(conds ...gen.Condition) IVehicleDo
	Or(conds ...gen.Condition) IVehicleDo
	Select(conds ...field.Expr) IVehicleDo
	Where(conds ...gen.Condition) IVehicleDo
	Order(conds ...field.Expr) IVehicleDo
	Distinct(cols ...field.Expr) IVehicleDo
	Omit(cols ...field.Expr) IVehicleDo
	Join(table schema.Tabler, on ...field.Expr) IVehicleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVehicleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVehicleDo
	Group(cols ...field.Expr) IVehicleDo
	Having(conds ...gen.Condition) IVehicleDo
	Limit(limit int) IVehicleDo
	Offset(offset int) IVehicleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVehicleDo
	Unscoped() IVehicleDo
	Create(values ...*models.Vehicle) error
	CreateInBatches(values []*models.Vehicle, batchSize int) error
	Save(values ...*models.Vehicle) error
	First() (*models.Vehicle, error)
	Take() (*models.Vehicle, error)
	Last() (*models.Vehicle, error)
	Find() ([]*models.Vehicle, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Vehicle, err error)
	FindInBatches(result *[]*models.Vehicle, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Vehicle) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVehicleDo
	Assign(attrs ...field.AssignExpr) IVehicleDo
	Joins(fields ...field.RelationField) IVehicleDo
	Preload(fields ...field.RelationField) IVehicleDo
	FirstOrInit() (*models.Vehicle, error)
	FirstOrCreate() (*models.Vehicle, error)
	FindByPage(offset int, limit int) (result []*models.Vehicle, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVehicleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v vehicleDo) Debug() IVehicleDo {
	return v.withDO(v.DO.Debug())
}

func (v vehicleDo) WithContext(ctx context.Context) IVehicleDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vehicleDo) ReadDB() IVehicleDo {
	return v.Clauses(dbresolver.Read)
}

func (v vehicleDo) WriteDB() IVehicleDo {
	return v.Clauses(dbresolver.Write)
}

func (v vehicleDo) Session(config *gorm.Session) IVehicleDo {
	return v.withDO(v.DO.Session(config))
}

func (v vehicleDo) Clauses(conds ...clause.Expression) IVehicleDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vehicleDo) Returning(value interface{}, columns ...string) IVehicleDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vehicleDo) Not(conds ...gen.Condition) IVehicleDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vehicleDo) Or(conds ...gen.Condition) IVehicleDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vehicleDo) Select(conds ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vehicleDo) Where(conds ...gen.Condition) IVehicleDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vehicleDo) Order(conds ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vehicleDo) Distinct(cols ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vehicleDo) Omit(cols ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vehicleDo) Join(table schema.Tabler, on ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vehicleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vehicleDo) RightJoin(table schema.Tabler, on ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vehicleDo) Group(cols ...field.Expr) IVehicleDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vehicleDo) Having(conds ...gen.Condition) IVehicleDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vehicleDo) Limit(limit int) IVehicleDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vehicleDo) Offset(offset int) IVehicleDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vehicleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVehicleDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vehicleDo) Unscoped() IVehicleDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vehicleDo) Create(values ...*models.Vehicle) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vehicleDo) CreateInBatches(values []*models.Vehicle, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vehicleDo) Save(values ...*models.Vehicle) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vehicleDo) First() (*models.Vehicle, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vehicle), nil
	}
}

func (v vehicleDo) Take() (*models.Vehicle, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vehicle), nil
	}
}

func (v vehicleDo) Last() (*models.Vehicle, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vehicle), nil
	}
}

func (v vehicleDo) Find() ([]*models.Vehicle, error) {
	result, err := v.DO.Find()
	return result.([]*models.Vehicle), err
}

func (v vehicleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Vehicle, err error) {
	buf := make([]*models.Vehicle, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vehicleDo) FindInBatches(result *[]*models.Vehicle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vehicleDo) Attrs(attrs ...field.AssignExpr) IVehicleDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vehicleDo) Assign(attrs ...field.AssignExpr) IVehicleDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vehicleDo) Joins(fields ...field.RelationField) IVehicleDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vehicleDo) Preload(fields ...field.RelationField) IVehicleDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vehicleDo) FirstOrInit() (*models.Vehicle, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vehicle), nil
	}
}

func (v vehicleDo) FirstOrCreate() (*models.Vehicle, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vehicle), nil
	}
}

func (v vehicleDo) FindByPage(offset int, limit int) (result []*models.Vehicle, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vehicleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vehicleDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vehicleDo) Delete(models ...*models.Vehicle) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vehicleDo) withDO(do gen.Dao) *vehicleDo {
	v.DO = *do.(*gen.DO)
	return v
}
