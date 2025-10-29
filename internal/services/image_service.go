package services

import (
	"fmt"
	"image/color"
	"os"
	"time"

	"github.com/0xatanda/country-currency/internal/models"
	"github.com/fogleman/gg"
	"gorm.io/gorm"
)

func GeneraeteSummaryImage(db *gorm.DB, t time.Time) {
	const imgPath = "cache/summary.png"
	os.MkdirAll("cache", 0755)

	var total int64
	db.Model(&models.Country{}).Count(&total)

	var tops []models.Country
	db.Order("estimated_gdp DESC").Limit(5).Find(&tops)

	dc := gg.NewContext(800, 400)
	dc.SetColor(color.White)
	dc.Clear()

	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(fmt.Sprintf("Total Countries: %d", total), 400, 50, 0.5, 0.5)

	y := 100.0
	for i, c := range tops {
		text := fmt.Sprintf("%d. %s - GDP: $%.2f", i+1, c.Name, c.EstimatedGDP)
		dc.DrawStringAnchored(text, 400, y, 0.5, 0.5)
		y += 40
	}

	dc.DrawStringAnchored(fmt.Sprintf("Last Refreshed: %s", t.Format(time.RFC3339)), 400, 350, 0.5, 0.5)

	file, _ := os.Create(imgPath)
	defer file.Close()
	dc.EncodePNG(file)
}
