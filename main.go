package main

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	addLevelMockAPI(router)

	router.Run(":8888")
}

func cleanJSONString(s string) string {
	result := strings.ReplaceAll(s, "\n", "")
	result = strings.ReplaceAll(result, "\t", "")
	result = strings.ReplaceAll(result, "\\", "")
	return result
}

func addLevelMockAPI(router *gin.Engine) {
	url := "/v1/mock/level"
	data := `{
		"error": null,
		"data": {
			"levels": [{
				"id": 1,
				"name": "Bronze",
				"is_active": true,
				"image_url": "http://example.com/image-bronze.jpg",
				"colour": "#000000",
				"minimum_experience": 1000,
				"detail": "<p>Bla bla</p>"
			}, {
				"id":2,
				"name":"Gold",
				"is_active": true,
				"image_url": "http://example.com/image-gold.jpg",
				"colour": "#ffffff",
				"minimum_experience":1500,
				"detail": "<p>Bla bla</p>"
			}],
			"total_data": 2
		},
		"message":"success get all level",
		"status":"success"
	}`

	router.GET(url, func(c *gin.Context) {
		var objmap map[string]*json.RawMessage
		err := json.Unmarshal([]byte(data), &objmap)
		if err != nil {
			c.JSON(400, "The provided json mock data is invalid")
		} else {
			c.JSON(200, objmap)
		}
	})
}
