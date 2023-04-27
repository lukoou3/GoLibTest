package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sort"
	"strings"
	"testing"
	"time"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jdbc_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return nil
}

type People struct {
	id       int32
	code     int32
	name     string
	age      int32
	birthday string
}

func queryOne() {
	s := "select id,code,name,age,birthday from people where id = ?"
	r, err := db.Query(s, 122)
	var p People
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		if r.Next() {
			r.Scan(&p.id, &p.code, &p.name, &p.age, &p.birthday)
			fmt.Printf("u: %v\n", p)
		}
	}
}

func queryOne2() {
	s := "select id,code,name,age,birthday from people where id = ?"
	r, err := db.Query(s, 1)
	var p People
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		if r.Next() {
			r.Scan(&p.id, &p.code, &p.name, &p.age, &p.birthday)
			fmt.Printf("u: %v\n", p)
		}
	}
}

func queryOneSome() {
	s := "select id,code,name,age,birthday from people"
	r, err := db.Query(s)
	defer r.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		var datas []People
		for r.Next() {
			var p People
			r.Scan(&p.id, &p.code, &p.name, &p.age, &p.birthday)
			datas = append(datas, p)
			fmt.Printf("u: %v\n", p)
		}
		fmt.Println(datas)
		sort.SliceStable(datas, func(i, j int) bool {
			return datas[i].code <= datas[j].code
		})
		fmt.Println(datas)
	}
}

func queryUpdate() {
	s := "update people set name=?,age=? where id=?"
	r, err := db.Exec(s, "青丝", 134, "90")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}
}

func queryDelete() {
	s := "delete from people where id=?"
	r, err := db.Exec(s, 90)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}

}

func insert() {
	s := "insert into people (id,code,name,age,birthday) values(?,?,?,?,?)"
	r, err := db.Exec(s, 1, 1111, "青丝", 11, "2022-12-17 14:14:02")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.LastInsertId()
		fmt.Printf("i: %v\n", i)
	}
}

func TestInsert(t *testing.T) {
	initDB()
	insert()
	queryOneSome()
}

func TestQuery(t *testing.T) {
	initDB()
	queryOne()
	fmt.Println(strings.Repeat("-", 60) + "\n")
	queryOne2()
	fmt.Println(strings.Repeat("-", 60) + "\n")
	queryOneSome()
}

func TestUpdate(t *testing.T) {
	initDB()
	queryUpdate()
	queryOneSome()
}

func TestDelete(t *testing.T) {
	initDB()
	queryDelete()
	queryOneSome()
}
