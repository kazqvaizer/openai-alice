package server

import (
	"fmt"
	"net/http"

	"github.com/kazqvaizer/openai-alice/dialog"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Host  string
	Port  string
	Token string
}

type AliceRequestScheme struct {
	OriginalUtterance string `json:"original_utterance"`
}

type AliceScheme struct {
	Request AliceRequestScheme `json:"request"`
}

func StartServer(config ServerConfig, dconf dialog.DialogConfig) {

	router := gin.Default()

	router.POST("/alice-webhook/:token/", WebhookHandler(config.Token, dconf))

	address := config.Host + ":" + config.Port
	router.Run(address)

}

func WebhookHandler(token string, dconf dialog.DialogConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		GotToken := ctx.Param("token")
		if GotToken == "" || token != GotToken {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid token."})
			return
		}

		var aliceMsg AliceScheme

		if err := ctx.BindJSON(&aliceMsg); err != nil {
			return
		}

		answer, err := dialog.AskAlice(aliceMsg.Request.OriginalUtterance, dconf)

		if err != nil {
			fmt.Printf("main error: %v\n", err)
			return
		}

		fmt.Println(answer)

		ctx.IndentedJSON(http.StatusOK, gin.H{
			"response": gin.H{
				"text": answer,
			},
			"version": "1.0",
		})
	}
}
