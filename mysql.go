package main

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("test")
	db, err := sql.Open("mysql", "root:devOps!23@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("db.Begin")
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback()

	reader := bufio.NewReader(os.Stdin)
	start := time.Now()

	fmt.Print("tx.Prepare: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	stmt, err := tx.Prepare("SELECT no FROM test WHERE no=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	//

	fmt.Print("stmt.Exec: ")
	text, _ = reader.ReadString('\n')
	fmt.Println(text)

	_, err = stmt.Exec(2)
	if err != nil {
		panic(err.Error())
	}

	fmt.Print("tx.Commit: ")
	text, _ = reader.ReadString('\n')
	fmt.Println(text)

	err = tx.Commit()
	if err != nil {
		panic(err.Error())
	}

	end := time.Now()
	fmt.Println(end.Sub(start))

}
