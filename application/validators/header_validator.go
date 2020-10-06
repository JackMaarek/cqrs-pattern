package validators

import (
	"errors"
	"github.com/JackMaarek/cqrsPattern/application/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateJsonHeader(c *gin.Context) error {
	header := structs.Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusBadRequest,"")
		return err
	}
	if header.ContentType != "application/json" {
		c.JSON(http.StatusBadRequest,"")
		return errors.New("Wrong content type.")
	}
	return nil
}
