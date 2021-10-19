package db

import "github.com/go-xorm/xorm"

type DataBaseOperator struct {
	sqlCli *xorm.Engine
}

func NewDataBaseOperator(sqlCli *xorm.Engine) DataBaseDao {
	return &DataBaseOperator{sqlCli: sqlCli}
}


