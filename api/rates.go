package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"se_school/util"
)

// @Summary Отримати поточний курс USD до UAH
// @Tags rate
// @Description Запит має повертати поточний курс USD до UAH використовуючи будь-який third party сервіс з публічним АРІ
// @ID rate
// @Produce application/json
// @Success 200 {number} float "Повертається актуальний курс USD до UAH"
// @Failure 400 "Invalid status value"
// @Router /rate [get]
func (s *Server) getCurrentRate(context *gin.Context) {
	currentRate, err := util.GetCurrentRate()
	if err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
	} else {
		context.JSON(http.StatusOK, currentRate)
	}
}
