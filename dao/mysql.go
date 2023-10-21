package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //重要，数据库驱动！
	"log"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

// 数据库操作的优化 便于扩展方法
type MsDB struct {
	*sql.DB
}

var DB MsDB

func init() {
	//执行main之前 先执行init方法
	dataSourceName := fmt.Sprintf("root:root@tcp(localhost:3306)/dwoura?charset=utf8&loc=%s&parseTime=true", url.QueryEscape("Asia/Shanghai"))
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}
	DB = MsDB{db}
}

// 好好研究一下这个数据库查询优化
func (d *MsDB) QueryOne(model interface{}, sql string, args ...interface{}) error {
	rows, err := d.Query(sql, args...)
	if err != nil {
		return nil
	}
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	vals := make([][]byte, len(columns))       //byte是通用类型，这相当于是一张表
	scans := make([]interface{}, len(columns)) //表示一行，长度为列长度
	for k := range vals {
		//把结果数组转化成指针类型的才能在Scan里扫描进去
		//收集每行指针
		scans[k] = &vals[k]
	}
	if rows.Next() {
		//把指赋到各个指针指向的变量进去
		err = rows.Scan(scans...)
		if err != nil {
			return err
		}
	}
	var result = make(map[string]interface{})
	elem := reflect.ValueOf(model).Elem() //用反射，根据model拿到字段类型，使其与数据库取出的未知数据一一对应
	for index, val := range columns {
		result[val] = string(vals[index])
	}
	//遍历model中的字段
	for i := 0; i < elem.NumField(); i++ {
		structField := elem.Type().Field(i)
		fieldInfo := structField.Tag.Get("orm") //获取标签中的orm的值,这些值与数据库字段名对应
		v := result[fieldInfo]                  //根据数据库字段名获取map中的值
		t := structField.Type                   //获取model中第i个字段的类型
		//对类型判断,设置成对应的值
		switch t.String() {
		case "int":
			s := v.(string)
			vInt, _ := strconv.Atoi(s)
			elem.Field(i).Set(reflect.ValueOf(vInt))
		case "string":
			elem.Field(i).Set(reflect.ValueOf(v.(string)))
		case "int64":
			s := v.(string)
			vInt64, _ := strconv.ParseInt(s, 10, 64)
			elem.Field(i).Set(reflect.ValueOf(vInt64))
		case "int32":
			s := v.(string)
			vInt32, _ := strconv.ParseInt(s, 10, 32)
			elem.Field(i).Set(reflect.ValueOf(vInt32))
		case "time.Time":
			s := v.(string)
			t, _ := time.Parse(time.RFC3339, s)
			elem.Field(i).Set(reflect.ValueOf(t))
		}
	}
	return nil
}
