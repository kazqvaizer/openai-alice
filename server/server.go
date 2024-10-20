package server

import (
	"fmt"
	"github.com/kazqvaizer/openai-alice/dialog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Host  string
	Port  string
	Token string
}

type AliceScheme struct {
	Message string `json:"message"`
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

		answer, err := dialog.AskAlice(aliceMsg.Message, dconf)

		if err != nil {
			fmt.Printf("main error: %v\n", err)
			return
		}

		fmt.Println(answer)

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": answer})

	}
}
