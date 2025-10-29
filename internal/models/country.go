package models

import "time"

type Country struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"unique;not null" json:"name"`
	Capital         string    `json:"capital"`
	Region          string    `json:"region"`
	Population      int64     `json:"population"`
	CurrencyCode    string    `json:"currency_code"`
	ExchangeRate    float64   `json:"exchange_rate"`
	EstimatedGDP    float64   `json:"estimated_gdp"`
	FlagURL         string    `json:"flag_url"`
	LastRefreshedAt time.Time `json:"last_refreshed_at"`
}
