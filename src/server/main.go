package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	prefix = "todo-"
	JSON   = "application/json"
)

var (
	ctx        = context.Background()
	client     *redis.Client
	dbAddr     = "localhost:6379"
	pathPrefix = "/api"
	port       = "1323"
	dbPass     = ""
)

type (
	todo struct {
		Id    string `json:"id" validate:"required"`
		Value string `json:"value" validate:"required"`
	}
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func initRedis() error {
	fmt.Printf("Testing Golang Redis ADDR=%s, PASS=%s\n", dbAddr, dbPass)

	client = redis.NewClient(&redis.Options{
		Addr:     dbAddr,
		Password: dbPass,
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)
	return err
}

func ping(c echo.Context) error {
	log.Println("Called function [ping]")
	return c.String(http.StatusOK, "pong")
}

func pingDB(c echo.Context) error {
	log.Println("Called function [pingDB]")
	pong, err := client.Ping(ctx).Result()
	log.Println(pong, err)
	if err != nil {
		return c.String(http.StatusInternalServerError, pong)
	}
	return c.String(http.StatusOK, pong)
}

func list(c echo.Context) error {
	log.Println("Called function [list]")

	keys, err := client.Keys(ctx, prefix+"*").Result()
	log.Println(keys, err)

	if len(keys) == 0 {
		return c.JSON(http.StatusNoContent, keys)
	}

	todoList := []todo{}

	for _, k := range keys {
		value, _ := client.Get(ctx, k).Result()
		todoList = append(todoList, todo{k, value})
	}

	return c.JSON(http.StatusOK, todoList)
}

func create(c echo.Context) error {
	log.Println("Called function [create]")

	t := new(todo)

	if err := c.Bind(t); err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Println(t)
	if err := c.Validate(t); err != nil {
		log.Println(err.Error())
		return err
	}

	result, err := client.Set(ctx, t.Id, t.Value, 0).Result()
	log.Println(result, err)

	return c.String(http.StatusOK, "created")
}

func delete(c echo.Context) error {
	log.Println("Called function [delete]")

	t := new(todo)
	if err := c.Bind(t); err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(t); err != nil {
		log.Println(err.Error())
		return err
	}

	result, err := client.Del(ctx, t.Id, t.Value).Result()
	log.Println(result, err)

	return c.String(http.StatusOK, "deleted")
}

func initServer() error {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	g := e.Group(pathPrefix)

	g.GET("/ping", ping)
	g.GET("/pingdb", pingDB)
	g.GET("/list", list)
	g.POST("/create", create)
	g.POST("/delete", delete)

	g.Use(middleware.CORS())

	err := e.Start(":" + port)
	if err != nil {
		e.Logger.Fatal(e.Start(":" + port))
	}
	return err
}

func main() {

	if value, ok := os.LookupEnv("DB_ADDR"); ok {
		dbAddr = value
	}

	if value, ok := os.LookupEnv("PATH_PREFIX"); ok {
		pathPrefix = value
	}

	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}

	if value, ok := os.LookupEnv("DB_PASS"); ok {
		dbPass = value
	}

	if err := initRedis(); err != nil {
		log.Println(err)
	}

	if err := initServer(); err != nil {
		log.Fatalln(err)
		return
	}
}
