package main

import (
	"database/sql"
	"log"
	"math/rand"
	"sync/atomic"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var counter uint64

func main() {
	dbconf := "root:@tcp(mysql-master:3306)/test_db"
	targetRPS := 800
	duration := 20 * time.Minute

	db, err := sql.Open("mysql", dbconf)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer db.Close()

	start := time.Now()
	ticker := time.NewTicker(time.Second / time.Duration(targetRPS))
	defer ticker.Stop()

	for i := 0; time.Since(start) < duration; i++ {
		<-ticker.C
		go insert(db)
	}
}

    func insert(db *sql.DB) {
	query := `INSERT INTO test_data (
		int_field1, int_field2, int_field3, int_field4, int_field5,
		int_field6, int_field7, int_field8, int_field9, int_field10,
		varchar_field1, varchar_field2, varchar_field3
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	args := make([]interface{}, 13)
	for i := 0; i < 10; i++ {
		args[i] = rand.Int31()
	}
	for i := 10; i < 13; i++ {
		args[i] = randString(255)
	}

	if _, err := db.Exec(query, args...); err != nil {
		log.Println(err)
		return
	}
	atomic.AddUint64(&counter, 1)
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
