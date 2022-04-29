package domain

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/rhino/request"

	"github.com/gin-gonic/gin"
)

//DailyMessage
func DailyMessage(c *gin.Context) {

	daily := request.NewDaily()
	text := c.Param("text")

	if text == "Metaphorpsum" {
		if err := daily.GetSentence(request.MetaphorpsumMethod); err != nil {

			fmt.Println("get Metaphorpsum error : ", err.Error())
			return
		}

		c.JSON(http.StatusOK, daily)
		return
	}

	if text == "Itsthisforthat" {
		if err := daily.GetSentence(request.ItsthisforthatMethod); err != nil {

			fmt.Println("get Metaphorpsum error : ", err.Error())
			return
		}

		c.JSON(http.StatusOK, daily)
		return
	}

	c.JSON(http.StatusOK, daily)
}
