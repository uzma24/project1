package internal

import (
	"log"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/project1/internal/utils"
	"github.com/project1/models"
)

func GetFrequency(c *gin.Context) ([]models.OutputFrequency, error) {
	textInputRequest := models.InputText{}
	err := c.BindJSON(&textInputRequest)
	if err != nil {
		log.Printf("Error while c.BindJSON(&textInputRequest): %v", err)
		return nil, &utils.BadRequest{ErrMessage: err.Error()}
	}
	words := strings.Split(textInputRequest.Text, " ") //Split return slice of strings

	freqMap := make(map[string]int)
	for _, value := range words {
		freqMap[value]++
	}

	mostUsedWordsFrquency := sortMapByValue(freqMap)
	if len(mostUsedWordsFrquency) > 10 {
		return mostUsedWordsFrquency[:10], nil
	}
	return mostUsedWordsFrquency, nil
}

func sortMapByValue(freqMap map[string]int) []models.OutputFrequency {
	keys := make([]string, 0, len(freqMap))
	for key := range freqMap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] > freqMap[keys[j]]
	})
	var outputFrequency []models.OutputFrequency

	for _, k := range keys {
		freq := models.OutputFrequency{
			Word: k,
			Freq: freqMap[k],
		}
		outputFrequency = append(outputFrequency, freq)
	}
	return outputFrequency
}
