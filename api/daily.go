package api

import (
	"encoding/json"
	"horoscope-api/scraper"
	"log"
	"net/http"
	"strings"
	"time"
)

func DailyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	sign := strings.ToLower(r.URL.Query().Get("sign"))
	category := strings.ToLower(r.URL.Query().Get("category"))
	if category == "" {
		category = "general"
	}
	date := strings.ToLower(r.URL.Query().Get("date"))
	if date != "" {
		if _, err := time.Parse("2006-01-02", date); err != nil {
			log.Printf("Error parsing date %s: %v", date, err)
			http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
			return
		}
	}

	horoscope, err := scraper.ScrapeDaily(sign, category, date)
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
