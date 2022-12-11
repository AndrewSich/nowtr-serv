package users

import (
	"net/http"

  "nowtr/configs"

	"github.com/gin-gonic/gin"
)

func usersIndex(c *gin.Context) {
  addr := "0xD79389938c5A2aA33E665B54932D87718c63c11B"
  data := configs.GetBalance(configs.GOERLI, addr)
	c.JSON(http.StatusOK, gin.H{
    // for [TESTING]
    "address": addr,
    "balance": data,
		//"status":  http.StatusOK,
		//"message": "success",
		//"data":    "Users Index",
	})
}
