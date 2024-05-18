package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createSubscriptionRequest struct {
	Email string `form:"email" binding:"required"`
}

// @Summary Підписати емейл на отримання поточного курсу
// @Tags subscription
// @Description Запит має перевірити, чи немає данної електронної адреси в поточній базі даних і, в разі її відсутності, записувати її.
// @ID subscribe
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Param email formData createSubscriptionRequest true "Електронна адреса, яку потрібно підписати"
// @Success 200 "E-mail додано"
// @Failure 409 "Повертати, якщо e-mail вже є в базі даних"
// @Router /subscribe [post]
func (s *Server) createSubscription(context *gin.Context) {
	var request createSubscriptionRequest

	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := s.store.GetSubscriptionByEmail(context, request.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			_, err = s.store.CreateSubscription(context, request.Email)

			if err != nil {
				context.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}

			context.Status(http.StatusOK)
			return
		} else {
			context.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}
	context.JSON(http.StatusConflict, errorResponse(errors.New("provided email is already subscribed")))
}
