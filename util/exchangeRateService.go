package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func GetCurrentRate() (float64, error) {
	var err error
	var uahRate float64

	apiKey := os.Getenv("API_TOKEN")
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/USD", apiKey)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return uahRate, err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return uahRate, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		return uahRate, err
	}

	if data["result"] != "success" {
		fmt.Println("Error:", data["error"])
		return uahRate, errors.New(fmt.Sprintf("third party services responded: %s", data["result"]))
	} else {
		conversionRates := data["conversion_rates"].(map[string]interface{})
		uahRate := conversionRates["UAH"].(float64)

		fmt.Printf("Current exchange rate for USD to UAH: %.4f\n", uahRate)
		return uahRate, nil
	}
}
