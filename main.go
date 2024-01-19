package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const (
	rediskey = "product_key"
	mysql    = "user:password@tcp(localhost:3306)/syncsb"
)

var (
	redisClient *redis.Client
	mysqldb     *sqlx.DB
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})

	var err error
	mysqldb, err = sqlx.Connect("mysql", mysql)
	if err != nil {
		log.Fatal(err)
	}
}

func readTest(c *fiber.Ctx) error {
	val, err := redisClient.Get(c.Context(), rediskey).Result()
	if err == redis.Nil {
		return c.Status(404).SendString("Data not found")
	} else if err != nil {
		return c.Status(500).SendString("Internal server error")
	}
	return c.SendString(val)
}

func writeTest(c *fiber.Ctx) error {
	user := User{
		Username: "Ballalala",
		Email:    "balal@gmail.com",
	}

	_, err := mysqldb.Exec(
		`INSERT INTO users (username, email) VALUES (?, ?)`,
		user.Username, user.Email,
	)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	err = redisClient.Set(c.Context(), rediskey, fmt.Sprintf("%v", user), 10*time.Second).Err()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendString("Data has been written")
}

func main() {
	app := fiber.New()

	app.Get("/read", readTest)
	app.Post("/write", writeTest)
	log.Fatal(app.Listen(":3000"))
}
