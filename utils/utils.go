package utils

import (
	"encoding/json"
	"gamingfacts/models"
	"math/rand"
	"os"
	"time"
)

// GetRandomFact reads a list of facts from a file and returns one at random.
func GetRandomFact(filePath string) (string, error) {
    var facts models.Facts
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
