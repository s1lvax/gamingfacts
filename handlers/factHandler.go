package handlers

import (
	"gamingfacts/utils"

	"github.com/gin-gonic/gin"
)

// FactHandler returns a random fact
func FactHandler(c *gin.Context) {
    fact, err := utils.GetRandomFact("data/facts.json")
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
