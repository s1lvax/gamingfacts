package main

import (
	"gamingfacts/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func main() {
    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, 
        AllowMethods:     []string{"GET"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }))

    // Rate limiting middleware
    rate, err := limiter.NewRateFromFormatted("20-M")
    if err != nil {
        panic(err)
    }
    store := memory.NewStore()
    r.Use(mgin.NewMiddleware(limiter.New(store, rate)))

    r.GET("/", handlers.FactHandler)

    r.Run() 
}
