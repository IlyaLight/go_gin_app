package models

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type queryModel struct {
	TableName   string
	Fields      []string
	CsvFields   string
	CsvсFields  string
	InsertQuery string
}

type IModel interface {
	TableName() string
}

var buffer = make(map[reflect.Type]*queryModel)

func Insert(db sqlx.Ext, arg interface{}) (sql.Result, error) {
	argType := reflect.TypeOf(arg)
	if _, ok := buffer[argType]; !ok {
		buffer[argType] = newQueryModel(arg)
	}
	queryM, _ := buffer[argType]
	return sqlx.NamedExec(db, queryM.InsertQuery, &arg)
}

func newQueryModel(arg interface{}) *queryModel {
	currentQm := queryModel{
		//TableName:strings.ToLower(reflect.TypeOf(arg).Elem().Name()),
		TableName: arg.(IModel).TableName(),
		Fields:    dbTagFields(arg),
	}
	currentQm.CsvFields = strings.Join(currentQm.Fields, ", ") // название полей должно браться из структуры а не из тегов!
	currentQm.CsvсFields = ":" + strings.Join(currentQm.Fields, ", :")
	currentQm.InsertQuery = "INSERT INTO " + currentQm.TableName + " (" + currentQm.CsvFields + ") " +
		"VALUES (" + currentQm.CsvсFields + ")"
	return &currentQm
}

// DBFields reflects on a struct and returns the values of fields with `db` tags,
// --deprecated--or a map[string]interface{} and returns the keys.
func dbTagFields(values interface{}) []string {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var fields []string
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" && field != "id" {
				fields = append(fields, field)
			}
		}
		return fields
	}
	//if v.Kind() == reflect.Map {
	//	for _, keyv := range v.MapKeys() {
	//		fields = append(fields, keyv.String())
	//	}
	//	return fields
	//}
	panic(fmt.Errorf("DBFields requires a struct or a map, found: %s", v.Kind().String()))
}
