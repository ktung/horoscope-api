package main

import (
	"horoscope-api/api"
	"net/http"
)

func main() {
	http.HandleFunc("/daily", api.DailyHandler)
	http.HandleFunc("/weekly", api.WeeklyHandler)
	http.HandleFunc("/monthly", api.MonthlyHandler)

	http.ListenAndServe(":8080", nil)
}
