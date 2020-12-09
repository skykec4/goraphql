package main

import (
	"fmt"

	// "goraphql/graph"
	// "goraphql/graph/generated"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/skykec4/goraphql/graph/generated"
	"github.com/skykec4/goraphql/graph/resolver"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	//  AllowedOrigins: []string{"http://localhost"},
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET, PATCH, PUT, POST, DELETE, OPTIONS"},
		AllowCredentials: true,
	}))

	pg := playground.Handler("GraphQL", "/query")
	config := generated.Config{
		Resolvers: &resolver.Resolver{}}
	h := handler.GraphQL(generated.NewExecutableSchema(config))

	e.GET("/", func(c echo.Context) error {
		pg.ServeHTTP(c.Response(), c.Request())

		fmt.Println("play ground!!")

		return nil
		// return c.String(http.StatusOK, "df")
		// return c.NoContent(http.StatusOK)
	})

	e.POST("/query", func(c echo.Context) error {

		// h.AddTransport(&transport.Websocket{
		// 	Upgrader: websocket.Upgrader{
		// 		CheckOrigin: func(r *http.Request) bool {
		// 			// Check against your desired domains here
		// 			return r.Host == "example.org"
		// 		},
		// 		ReadBufferSize:  1024,
		// 		WriteBufferSize: 1024,
		// 	},
		// })
		h.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	// e.HideBanner = true
	e.Logger.Fatal(e.Start(":3000"))

}
