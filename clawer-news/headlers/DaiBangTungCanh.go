package headlers

import (
	"clawer-news/models"
	"clawer-news/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func GetNews(c *gin.Context) {
	body, err := utils.ConnectAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var rawNews struct {
		Articles []struct {
			Title       string `json:"title"`
			URL         string `json:"url"`
			PublishedAt string `json:"publishedAt"`
			Source      struct {
				Name string `json:"name"`
			} `json:"source"`
		} `json:"articles"`
	}

	if err := json.Unmarshal(body, &rawNews); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unmarshal api: " + err.Error()})
		return
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"127.0.0.1:9092"},
		Topic:    "news.raw",
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	result := make([]models.News, 0, len(rawNews.Articles))
	msgs := make([]kafka.Message, 0, len(rawNews.Articles))

	for _, a := range rawNews.Articles {
		n := models.News{
			Title:   a.Title,
			Link:    a.URL,
			Src:     a.Source.Name,
			PubDate: a.PublishedAt,
		}
		result = append(result, n)

		b, err := json.Marshal(n)
		if err == nil {
			msgs = append(msgs, kafka.Message{Value: b})
		}
	}

	if len(msgs) > 0 {
		if err := writer.WriteMessages(c.Request.Context(), msgs...); err != nil {
			log.Println("kafka write failed:", err)
		}
	}

	c.JSON(200, result)
}
