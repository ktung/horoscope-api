package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Horoscope struct {
	Horoscope string `json:"horoscope"`
}

func ScrapeDaily(sign string, category string, date string) (*Horoscope, error) {
	zodiacNumber, ok := ZodiacMap[sign]
	if !ok {
		return nil, fmt.Errorf("invalid sign: %s", sign)
	}
	if _, ok := allowedCategories["daily"][category]; !ok {
		return nil, fmt.Errorf("invalid category: %s", category)
	}

	var url string
	if date != "" {
		date = strings.ReplaceAll(date, "-", "")
		url = fmt.Sprintf("https://www.horoscope.com/us/horoscopes/%s/horoscope-archive.aspx?sign=%d&laDate=%s", category, zodiacNumber, date)
	} else {
		url = fmt.Sprintf("https://www.horoscope.com/us/horoscopes/%s/horoscope-%s-daily-today.aspx?sign=%d", category, category, zodiacNumber)
	}
	return scrapeHoroscope(url)
}

func ScrapeWeekly(sign string, category string) (*Horoscope, error) {
	zodiacNumber, ok := ZodiacMap[sign]
	if !ok {
		return nil, fmt.Errorf("invalid sign: %s", sign)
	}
	if _, ok := allowedCategories["weekly"][category]; !ok {
		return nil, fmt.Errorf("invalid category: %s", category)
	}

	url := fmt.Sprintf("https://www.horoscope.com/us/horoscopes/%s/horoscope-%s-weekly.aspx?sign=%d", category, category, zodiacNumber)
	return scrapeHoroscope(url)
}

func ScrapeMonthly(sign string, category string) (*Horoscope, error) {
	zodiacNumber, ok := ZodiacMap[sign]
	if !ok {
		return nil, fmt.Errorf("invalid sign: %s", sign)
	}
	if _, ok := allowedCategories["monthly"][category]; !ok {
		return nil, fmt.Errorf("invalid category: %s", category)
	}

	url := fmt.Sprintf("https://www.horoscope.com/us/horoscopes/%s/horoscope-%s-monthly.aspx?sign=%d", category, category, zodiacNumber)
	return scrapeHoroscope(url)
}

func scrapeHoroscope(url string) (*Horoscope, error) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.horoscope.com"),
	)

	horoscope := &Horoscope{}

	found := false
	c.OnHTML("div.main-horoscope p:first-of-type", func(e *colly.HTMLElement) {
		if !found {
			horoscope.Horoscope = strings.TrimSpace(e.Text)
			found = true
		}
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return horoscope, nil
}

var ZodiacMap = map[string]int{
	"aries":       1,
	"taurus":      2,
	"gemini":      3,
	"cancer":      4,
	"leo":         5,
	"virgo":       6,
	"libra":       7,
	"scorpio":     8,
	"sagittarius": 9,
	"capricorn":   10,
	"aquarius":    11,
	"pisces":      12,
}

// allow money only for weekly
var allowedCategories = map[string]map[string]bool{
	"daily": {
		"general": true,
		"love": true,
		"career": true,
	},
	"weekly": {
		"general": true,
		"love": true,
		"career": true,
		"money": true,
	},
	"monthly": {
		"general": true,
		"love": true,
		"career": true,
	},
}
