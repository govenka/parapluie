package main

import (
	"encoding/json"
	"exercice_meteo/meteo"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTemperature(t *testing.T) {
	// Créer un enregistreur de test Gin
	w := httptest.NewRecorder()

	// Créer une nouvelle instance de Gin
	r := gin.Default()

	// Définir la route pour obtenir la température
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

	// Envoyer une requête GET pour obtenir la température
	req, _ := http.NewRequest("GET", "/temperature?latitude=48.8567&longitude=2.3508&previsions_jours=1", nil)
	r.ServeHTTP(w, req)

	// Vérifier que le code de statut est OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserialiser la réponse JSON
	var resp meteo.MeteoResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Vérifier que la température est dans une plage raisonnable
	assert.InDelta(t, 10.0, resp.Temperature, 20.0)
}

func TestGetPrevisions(t *testing.T) {
	// Créer un enregistreUr de test Gin
	w := httptest.NewRecorder()

	// Créer une nouvelle instance de Gin
	r := gin.Default()

	// Définir la route pour obtenir les prévisions de précipitations
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

	// Envoyer une requête GET pour obtenir les prévisions de précipitations
	req, _ := http.NewRequest("GET", "/parapluie?latitude=48.8567&longitude=2.3508", nil)
	r.ServeHTTP(w, req)

	// Vérifier que le code de statut est OK
	assert.Equal(t, http.StatusOK, w.Code)

	// Deserialiser la réponse JSON
	var resp meteo.PrevisionsResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	// Vérifier que la probabilité moyenne de précipitations est dans une plage valide
	assert.InDelta(t, 50.0, resp.MoyennePrecipitations, 50.0)

	// Vérifier qu'au moins une plage horaire recommandée a été retournée
	assert.GreaterOrEqual(t, len(resp.Sort_ton_parapluie), 1)
}
