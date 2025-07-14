package api

import (
	"encoding/json"
	"horoscope-api/scraper"
	"log"
	"net/http"
	"strings"
)

func MonthlyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	sign := strings.ToLower(r.URL.Query().Get("sign"))
	category := strings.ToLower(r.URL.Query().Get("category"))
	if category == "" {
		category = "general"
	}

	horoscope, err := scraper.ScrapeMonthly(sign, category)
	if err != nil {
		log.Printf("Error scraping horoscope for sign %s: %v", sign, err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(horoscope); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}
}
