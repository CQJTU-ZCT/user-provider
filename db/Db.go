/**
 * @Desc
 * @author zjhfyq
 * @data 2018/3/29 14:24.
 */
package db

import (
	"database/sql"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	_"github.com/go-sql-driver/mysql"
	"strings"
	"user-provider/models"
)

var db *sql.DB
var dbConfig models.DbConfig

func init() {
	dbConfig.DataSourceName =  beego.AppConfig.String("dataSourceName")
	dbConfig.DriverName = beego.AppConfig.String("driverName")
	open  , err :=  strconv.Atoi(beego.AppConfig.String("maxOpenConns"))
	if err != nil {
		log.Println(err)
		open = 0
	}else {
		dbConfig.MaxOpenConns = open
	}
	idel , err :=  strconv.Atoi(beego.AppConfig.String("maxIdelConns"))
	if err != nil {
		log.Println(err)
		idel  = 1
	}else {
		dbConfig.MaxIdleConns = idel
	}
	log.Println("db:",dbConfig)
	db = GetDB()
}

func GetDB()(dbp *sql.DB) {
	if db != nil {
		dbp =  db
	}else {
		dbTemp , err := sql.Open(dbConfig.DriverName, dbConfig.DataSourceName)
		if err != nil {
			log.Println(err)
			dbp = nil
		}else {
			dbTemp.SetMaxOpenConns(dbConfig.MaxOpenConns)
			dbTemp.SetMaxIdleConns(dbConfig.MaxIdleConns)
			db = dbTemp
			dbp = db
		}
	}
	return
}

func Select(sqlStr string, args ...interface{}) (result []map[string]string , err error) {
	for index,value := range args {
		if value , ok :=value.(string);ok {
			args[index] = strings.TrimSpace(value)
		}
	}
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Println(err)
		result = nil
	} else {
		rows,err := stmt.Query(args...)
		if err != nil {
			result = nil
		} else {
			columns, err := rows.Columns()
			len := len(columns)
			if err != nil {
				result = nil
			} else {
				for rows.Next() {
					values := make([][]byte, len)
					scans := make([]interface{}, len)
					for i := 0; i < len; i++ {
						scans[i] = &values[i]
					}
					err := rows.Scan(scans...)
					if err != nil {
						result = nil
					} else {
						temp := make(map[string]string)
						for i := 0; i < len; i++ {
							temp[columns[i]] = string(values[i])
						}
						result = append(result, temp)
					}
				}
			}
		}
	}
	return result,err
}



func Delete(sqlStr string, args ...interface{}) (result sql.Result,err error) {
	stmt , err :=db.Prepare(sqlStr)
	if err != nil {
		result =  nil
	}else {
		result , err =stmt.Exec(args)
		if err != nil {
			result =  nil
		}
	}
	return result,err
}



func Insert(sqlStr string,args ...interface{}) (result sql.Result,err error) {
	stmt , err :=db.Prepare(sqlStr)
	if err != nil {
		result = nil
	}else {
		result , err =stmt.Exec(args...)
		if err  != nil {
			result = nil
		}
	}
	return result,err
}

func Update(sqlStr string , args ...interface{}) (result sql.Result,err error) {
	stmt , err :=db.Prepare(sqlStr)
	if err != nil {
		result =  nil
	}else {
		result , err =stmt.Exec(args...)
		if err != nil {
			result =  nil
		}
	}
	return  result ,err
}



