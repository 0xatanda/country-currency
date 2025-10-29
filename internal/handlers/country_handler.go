package handlers

import (
	"net/http"

	"github.com/0xatanda/country-currency/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CountryHandler struct {
	Service *services.CountryService
}

func RegisterCountryRoutes(r *gin.Engine, db *gorm.DB) {
	h := &CountryHandler{Service: services.NewCountryService(db)}

	api := r.Group("/countries")
	api.POST("/refresh", h.RefreshCountries)
	api.GET("", h.GetAllCountries)
	api.GET("/:name", h.GetCountry)
	api.DELETE("/:name", h.DeleteCountry)
	api.GET("/image", h.GetSummaryImage)

	r.GET("/status", h.GetStatus)
}

func (h *CountryHandler) RefreshCountries(c *gin.Context) {
	if err := h.Service.RefreshCountries(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error":   "External data source unavailable",
			"details": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Countries refreshed sucessfully"})
}

func (h *CountryHandler) GetAllCountries(c *gin.Context) {
	countries, err := h.Service.GetAllCountries(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, countries)
}

func (h *CountryHandler) GetCountry(c *gin.Context) {
	name := c.Param("name")
	country, err := h.Service.GetCountry(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}
	c.JSON(http.StatusOK, country)
}

func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	name := c.Param("name")
	if err := h.Service.DeleteCountry(name); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Country deleted"})
}

func (h *CountryHandler) GetStatus(c *gin.Context) {
	status := h.Service.GetStatus()
	c.JSON(http.StatusOK, status)
}

func (h *CountryHandler) GetSummaryImage(c *gin.Context) {
	imgPath := "cache/summary.png"
	c.File(imgPath)
}
