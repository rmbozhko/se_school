package api

import (
	"context"
	"fmt"
	"se_school/util"
	"time"
)

func (s *Server) SendRatesEmail() {
	currentRate, err := util.GetCurrentRate()
	if err != nil {
		fmt.Println("Error getting current rate:", err)
		return
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	subscriptions, err := s.store.GetSubscriptions(context.Background())
	if err != nil {
		fmt.Println("Error getting subscribers:", err)
		return
	}

	for _, subscriber := range subscriptions {
		emailBody := fmt.Sprintf("<h2>Актуальний курс USD/UAH: %f</h2><br><br><h4>Час оновлення: %s</h4>", currentRate, currentTime)

		err := s.dialer.SendEmail(subscriber.Email.(string), "Актуальний курс USD/UAH", emailBody)
		if err != nil {
			fmt.Println("Error sending an email:", err)
			return
		}
	}
}
