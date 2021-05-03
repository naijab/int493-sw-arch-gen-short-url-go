package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
	"github.com/matoous/go-nanoid/v2"
)

type CreateShortLinkReq struct {
	Url     string `json:"url"`
}

type ShortLinkResp struct {
	Link     string `json:"link"`
}

type LinkStatResp struct {
	Visit     int `json:"visit"`
}

type MessageResp struct {
	Message     int `json:"message"`
}

var ctx          = context.Background()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env var
	AppPort := os.Getenv("PORT")
	BaseUrl      := os.Getenv("BASE_URL")
	RedisConnect := os.Getenv("REDIS_CONNECT")

	// Create new fiber app
	app := fiber.New()

	// Connect to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisConnect,
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	log.Println("Redis Connection: ", rdb.Ping(ctx))

	app.Post("/link", func(c *fiber.Ctx) error {
		input := new(CreateShortLinkReq)
		if err := c.BodyParser(input); err != nil {
			return err
		}

		// TODO: Implement generate unique short link

		// Hash full url for check is in redis if exist return short url from redis
		// If not exist generate short from hash full url

		//rdb.HMSet(ctx, )

		link := fmt.Sprintf("%s/l/%x", BaseUrl, bs)
		short := ShortLinkResp {
			Link: link,
		}
		return c.JSON(short)
	})

	app.Get("/l/:short", func(c *fiber.Ctx) error {
		shortParam := c.Params("short")
		linkRaw, err := rdb.Get(ctx, shortParam).Result()

		// TODO: Implement get full url and redirect from short key

		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.Redirect(linkRaw)
	})

	app.Get("/l/:short/stats", func(c *fiber.Ctx) error {
		// TODO: Implement link stats
		visit := LinkStatResp {
			Visit: 30,
		}
		return c.JSON(visit)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", AppPort)))
}
