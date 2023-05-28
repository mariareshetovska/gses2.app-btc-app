package handler

import (
	"encoding/json"
	"net/http"

	"gses2.app-btc/services"
)

// @Summary GetRate
// @Tags Rate
// @Description Get the current BTC to UAH rate
// @Produce json
// @Success 200
// @Failure 400
// @Router /rate [get]
func GetRateHandler(w http.ResponseWriter, r *http.Request) {
	price, err := services.GetBitcoinRate()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := struct {
		Price float64 `json:"price"`
	}{
		Price: price,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
