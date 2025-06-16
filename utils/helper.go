package utils

import (
	models "github.com/abdanhafidz/ai-visual-multi-modal-backend/models"
	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) models.AccountData {
	cParam, _ := c.Get("accountData")
	return cParam.(models.AccountData)
}
