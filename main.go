package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/matoous/go-nanoid/v2"
	"log"
	"os"
	"strconv"
)

type CreateShortLinkReq struct {
	Url string `json:"url"`
}

type ShortLinkResp struct {
	Link string `json:"link"`
}

type LinkStatResp struct {
	Visit int `json:"visit"`
}

var ctx = context.Background()

const FullKey = "full"
const CountKey = "count"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get env var
	AppPort := os.Getenv("PORT")
	ServerName := os.Getenv("SERVER_NAME")
	BaseUrl := os.Getenv("BASE_URL")
	RedisConnect := os.Getenv("REDIS_CONNECT")

	// Create new fiber app
	app := fiber.New()

	// Connect to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisConnect,
		Password: "",
		DB:       0,
	})
	log.Println("Redis Connection: ", rdb.Ping(ctx))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"stack": "golang (fiber), redis",
			"server": ServerName,
			"version": 4,
		})
	})

	app.Post("/link", func(c *fiber.Ctx) error {
		input := new(CreateShortLinkReq)
		if err := c.BodyParser(input); err != nil {
			return err
		}

		fullUrl := input.Url

		fullUrlInRedis, _ := rdb.Get(ctx, fullUrl).Result()
		if fullUrlInRedis != "" {
			// If full url is in redis
			link := fmt.Sprintf("%s/l/%s", BaseUrl, fullUrlInRedis)
			short := ShortLinkResp{
				Link: link,
			}
			return c.JSON(short)
		}

		// Full url not exist it will generate short url and save to redis
		id, _ := gonanoid.New(6)
		rdb.HMSet(ctx, id, FullKey, fullUrl)
		rdb.HMSet(ctx, id, CountKey, 0)

		// Save full url to redis
		rdb.Set(ctx, fullUrl, id, 0)

		link := fmt.Sprintf("%s/l/%s", BaseUrl, id)
		short := ShortLinkResp{
			Link: link,
		}
		return c.JSON(short)
	})

	app.Get("/l/:short", func(c *fiber.Ctx) error {
		shortParam := c.Params("short")

		// Get full url from short url
		link, _ := rdb.HMGet(ctx, shortParam, FullKey).Result()
		if link != nil && link[0] != nil {
			// Increment count with plus 1
			rdb.HIncrBy(ctx, shortParam, CountKey, 1)

			fullUrl := link[0].(string)
			return c.Redirect(fullUrl)
		}

		return c.SendStatus(fiber.StatusNotFound)
	})

	app.Get("/l/:short/stats", func(c *fiber.Ctx) error {
		shortParam := c.Params("short")

		// Get total count from short url
		link, _ := rdb.HMGet(ctx, shortParam, CountKey).Result()
		if link != nil && link[0] != nil {
			countRaw := link[0].(string)
			count, _ := strconv.Atoi(countRaw)
			visit := LinkStatResp{
				Visit: count,
			}
			return c.JSON(visit)
		}

		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", AppPort)))
}
