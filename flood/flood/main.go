package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(mysql-master:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(iteration int) {
				defer wg.Done()
				for j := 0; j < 20; j++ {
					_, err := db.Exec(
						"INSERT INTO user_cart (user_id, product_id, quantity, item_price_cents, total_price_cents, original_product_id, category_id, warehouse_id, promotion_id, product_name, product_attributes, session_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
						15432, // user_id
						88765, // product_id
						2,     // quantity
						29999, // item_price_cents (299.99)
						59998, // total_price_cents (599.98)
						88760, // original_product_id
						15,    // category_id (электроника)
						3,     // warehouse_id
						nil,   // promotion_id (NULL)
						"Смартфон iPhone 15 Pro 128GB Синий", // product_name
						"color:blue|storage:128gb|model:pro", // product_attributes
						"a1b2c3d4e5f6g7h8i9j0",               // session_id
					)

					if err != nil {
						log.Printf("Ошибка при вставке записи %d: %v", iteration, err)
					}
				}
			}(i)
		}

	wg.Wait()
	fmt.Println("Все операции завершены")
}
