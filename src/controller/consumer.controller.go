package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"unisun/api/article-listening/src/constants"
	"unisun/api/article-listening/src/model"
	"unisun/api/article-listening/src/service"

	"github.com/gin-gonic/gin"
)

func ArticleAll(c *gin.Context) {
	var articles model.Articles
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE)
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetArticles(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": articles})
}
