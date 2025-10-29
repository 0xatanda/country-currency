package services

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/0xatanda/country-currency/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CountryService struct {
	DB *gorm.DB
}

func NewCountryService(db *gorm.DB) *CountryService {
	return &CountryService{DB: db}
}

func (s *CountryService) RefreshCountries() error {
	countriesURL := "https://restcountries.com/v2/all?fields=name,capital,region,population,flag,currencies"
	exchanegURL := "https://open.er-api.com/v6/latest/USD"

	cResp, err := http.Get(countriesURL)
	if err != nil {
		return fmt.Errorf("could not fetch countries: %v", err)
	}
	defer cResp.Body.Close()

	var countries []map[string]interface{}
	body, _ := io.ReadAll(cResp.Body)
	json.Unmarshal(body, &countries)

	eResp, err := http.Get(exchanegURL)
	if err != nil {
		return fmt.Errorf("could not fetch exchange rates: %v", err)
	}
	defer eResp.Body.Close()

	var exchangeData map[string]interface{}
	eBody, _ := io.ReadAll(eResp.Body)
	json.Unmarshal(eBody, &exchangeData)

	rates := exchangeData["rates"].(map[string]interface{})

	now := time.Now()

	for _, c := range countries {
		name := c["name"].(string)
		pop := int64(c["population"].(float64))
		flag := c["flag"]
		region := ""
		if v, ok := c["region"].(string); ok {
			region = v
		}
		capital := ""
		if v, ok := c["capital"].(string); ok {
			capital = v
		}

		var currencyCode string
		if curVal, ok := c["currencies"]; ok && curVal != nil {
			if currencies, ok := curVal.([]interface{}); ok && len(currencies) > 0 {
				if currency, ok := currencies[0].(map[string]interface{}); ok {
					if code, ok := currency["code"].(string); ok {
						currencyCode = code
					}
				}
			}
		}

		var exchangeRate, estimatedGDP float64
		if val, ok := rates[currencyCode]; ok {
			exchangeRate = val.(float64)
			randomMultiplier := float64(rand.Intn(1001) + 1000)
			estimatedGDP = float64(pop) * exchangeRate * randomMultiplier
		}

		country := models.Country{
			Name:            name,
			Capital:         fmt.Sprintf("%v", capital),
			Region:          region,
			Population:      pop,
			CurrencyCode:    currencyCode,
			ExchangeRate:    exchangeRate,
			EstimatedGDP:    estimatedGDP,
			FlagURL:         fmt.Sprintf("%v", flag),
			LastRefreshedAt: now,
		}

		var existing models.Country
		if err := s.DB.Where("LOWER(name) = ?", strings.ToLower(name)).First(&existing).Error; err == nil {
			s.DB.Model(&existing).Updates(country)
		} else {
			s.DB.Create(&country)
		}
	}

	GeneraeteSummaryImage(s.DB, now)
	return nil
}

func (s *CountryService) GetAllCountries(c *gin.Context) ([]models.Country, error) {
	var countries []models.Country
	query := s.DB

	if region := c.Query("region"); region != "" {
		query = query.Where("region = ?", region)
	}

	if currency := c.Query("currency"); currency != "" {
		query = query.Where("currency_code = ?", currency)
	}

	sort := c.Query("sort")
	if sort == "gdp_asc" {
		query = query.Order("estimated_gdp ASC")
	}

	if err := query.Find(&countries).Error; err != nil {
		return nil, err
	}

	return countries, nil
}

func (s *CountryService) GetCountry(name string) (*models.Country, error) {
	var country models.Country
	if err := s.DB.Where("LOWER(name) = ?", strings.ToLower(name)).First(&country).Error; err != nil {
		return nil, err
	}
	return &country, nil
}

func (s *CountryService) DeleteCountry(name string) error {
	return s.DB.Where("LOWER(name) = ?", strings.ToLower(name)).Delete(&models.Country{}).Error
}

func (s *CountryService) GetStatus() map[string]interface{} {
	var count int64
	s.DB.Model(&models.Country{}).Count(&count)

	var last models.Country
	s.DB.Order("last_refreshed_at DESC").First(&last)

	return map[string]interface{}{
		"total_countries":   count,
		"last_refreshed_at": last.LastRefreshedAt,
	}
}
