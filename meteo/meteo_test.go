package meteo

import (
	"testing"
)

// TestApiGetCall vérifie que la fonction apiGetCall retourne correctement la réponse de l'API.
func TestApiGetCall(t *testing.T) {
	err, resp := apiGetCall("https://api.open-meteo.com/v1/forecast?latitude=48.8567&longitude=2.3508&forecast_days=1&hourly=temperature_2m&past_days=1")
	if err != nil {
		t.Errorf("apiGetCall returned an error: %v", err)
	}
	if resp == "" {
		t.Error("apiGetCall returned an empty response")
	}
}

// TestGetMeteo vérifie que la fonction GetMeteo calcule correctement la température moyenne.
func TestGetMeteo(t *testing.T) {
	resp, err := GetMeteo("48.8567", "2.3508", "1")
	if err != nil {
		t.Errorf("GetMeteo returned an error: %v", err)
	}
	if resp.Temperature < 0 || resp.Temperature > 50 {
		t.Errorf("GetMeteo returned an invalid temperature: %f", resp.Temperature)
	}
}

// TestGetPrevisions vérifie que la fonction GetPrevisions calcule correctement la probabilité moyenne de précipitations et les plages horaires recommandées.
func TestGetPrevisions(t *testing.T) {
	resp, err := GetPrevisions("48.8567", "2.3508")
	if err != nil {
		t.Errorf("GetPrevisions returned an error: %v", err)
	}
	if resp.MoyennePrecipitations < 0 || resp.MoyennePrecipitations > 100 {
		t.Errorf("GetPrevisions returned an invalid average precipitation probability: %f", resp.MoyennePrecipitations)
	}
	if len(resp.Sort_ton_parapluie) == 0 {
		t.Error("GetPrevisions returned no recommended time slots")
	}
}
