package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gsom95/go-web-dev/models"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("connected")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT NOT NULL
	  );
	  
	  CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
	  );`)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Tables created.")

	us := models.UserService{
		DB: db,
	}
	user, err := us.Create("bob@bobmail.kek", "secret password")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("User:", user)

	// 	name := "New User"
	// 	email := "new@calhoun.io"
	// 	row := db.QueryRow(`INSERT INTO users (name, email)
	//   VALUES ($1, $2) RETURNING id;`, name, email)

	// no need to call row.Err() because as noted in docs:
	// "If this error is not nil, this error will also be returned from Scan."
	// if row.Err() != nil {
	// 	log.Fatalln(err)
	// }
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("User created. id =", id)

	// 	id := 1
	// 	row := db.QueryRow(`
	//   SELECT name, email
	//   FROM users
	//   WHERE id=$1;`, id)
	// 	var name, email string
	// 	err = row.Scan(&name, &email)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	log.Printf("User information: name=%s, email=%s\n", name, email)

	// userID := id
	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 	  INSERT INTO orders(user_id, amount, description)
	// 	  VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	// log.Println("Created fake orders.")

	// type Order struct {
	// 	ID          int
	// 	UserID      int
	// 	Amount      int
	// 	Description string
	// }

	// var orders []Order
	// rows, err := db.Query(`
	// SELECT id, amount, description
	// FROM orders
	// WHERE user_id = $1`, id)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer func() {
	// 	_ = rows.Close()
	// }()

	// for rows.Next() {
	// 	order := Order{UserID: id}
	// 	err = rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(orders)
}