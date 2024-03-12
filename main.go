package main

import (
	"exercice_meteo/meteo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Default port, sera rerouté par Traefik sur le port interne
const PORT = "22222"

func main() {

	r := gin.Default()

	// Route pour l'obtention des temperatures
	r.GET("/temperature", func(c *gin.Context) {

		latitude := c.Query("latitude")
		longitude := c.Query("longitude")
		previsionsJours := c.Query("previsions_jours")

		result, err := meteo.GetMeteo(latitude, longitude, previsionsJours)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	// Route pour l'obtention des prévisions
	r.GET("/parapluie", func(c *gin.Context) {
		latitude := c.Query("latitude")
		longitude := c.Query("longitude")

		result, err := meteo.GetPrevisions(latitude, longitude)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	//On lance le serveur
	r.Run(":" + PORT)

}
