package dao

import (
	"database/sql"
	"fmt"
)

type IMysqlOperation interface {
	Query(sql string) []map[string]interface{}
	Execute(sql string) (rowsAffected int64)
}

type mysqlObject struct {
	db *sql.DB
}

func NewMysqlObject(db *sql.DB) IMysqlOperation {
	return &mysqlObject{
		db: db,
	}
}

// Query Function parameters only support select queries. The return parameter is an array of maps.
func (d *mysqlObject) Query(sql string) []map[string]interface{} {
	rows, err := d.db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	// columns 表的字段名称
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	// ret 列表map,存放返回有多少行且每一行的字段值
	ret := make([]map[string]interface{}, 0)
	for rows.Next() {
		//m 存放每一行的字段值
		m := make(map[string]interface{})
		// columns 表字段名称的一个list ["id","delete_at"]
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Printf("rows.Scan failed,error:%v\n", err)
		}
		for i, col := range columns {
			//i 代表数据下标，col 代表数组值
			val := values[i]
			b, ok := val.([]byte)

			var v interface{}
			if ok {
				v = string(b)
			} else {
				v = val
			}
			m[col] = v
			fmt.Println(col, v)
		}
		ret = append(ret, m)
	}
	return ret
}

//Execute function support update ,delete,insert and select operation.but no return value.
func (d *mysqlObject) Execute(sql string) (rowsAffected int64) {
	result, err := d.db.Exec(sql)

	if err != nil {
		panic(err)
	}
	rowsAffectedNum, _ := result.RowsAffected()
	return rowsAffectedNum
}
