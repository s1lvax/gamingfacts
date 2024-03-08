package main

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// Fact structure
type Facts struct {
	VideoGameFacts []string `json:"video_game_facts"`
}

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
	rate, err := limiter.NewRateFromFormatted("10-M")
	if err != nil {
		panic(err)
	}
	store := memory.NewStore()
	r.Use(mgin.NewMiddleware(limiter.New(store, rate)))

	r.GET("/", FactHandler)

	r.Run() 
}

// FactHandler returns a random fact
func FactHandler(c *gin.Context) {
	fact, err := getRandomFact("data/facts.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"fact": fact,
	})
}

// getRandomFact reads a list of facts from a file and returns one at random.
func getRandomFact(filePath string) (string, error) {
	var facts Facts
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(file, &facts)
	if err != nil {
		return "", err
	}
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	index := rng.Intn(len(facts.VideoGameFacts))
	return facts.VideoGameFacts[index], nil
}
