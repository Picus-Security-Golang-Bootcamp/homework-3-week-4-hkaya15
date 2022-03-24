package db

import "gorm.io/gorm"

type DBSelector interface {
	Create() (*gorm.DB, error)
}

type DBBase struct {
	DbType DBSelector
}

func New(db DBSelector) *DBBase {
	return &DBBase{DbType: db}
}

func (base *DBBase) Connect() (*gorm.DB, error){
	res,err:=base.DbType.Create()
	if err!=nil{
		return nil,err
	}
	return res,nil
}
