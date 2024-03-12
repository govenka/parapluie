package meteo

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// Permets de gérer la réponse de l'API
func apiGetCall(url string) (error, string) {
	var apiResp ApiResponse

	client := resty.New()

	if client == nil {
		fmt.Println("Erreur lors de la création du client Resty")
		return errors.New("Erreur lors de la création du client Resty"), ""
	}

	resp, err := client.R().
		SetResult(&apiResp).
		Get(url)

	if err != nil {
		return err, ""
	}

	if resp.StatusCode() == http.StatusOK {

		return nil, resp.String()
	}

	return errors.New("Erreur lors de la requête"), ""

}

// Permet de récupérer les données de l'API au niveau de la météo
// les arguments sont la latitude, la longitude et le nombre de jours de prévisions à récupérer
func GetMeteo(latitude string, longitude string, previsions_jours string) (*MeteoResponse, error) {
	urlMeteo := "https://api.open-meteo.com/v1/forecast?latitude=" + latitude + "&longitude=" + longitude + "&forecast_days=1&hourly=temperature_2m&past_days=" + previsions_jours

	err, reponse := apiGetCall(urlMeteo)
	if err != nil {
		return &MeteoResponse{}, errors.New("Erreur lors de la requête à l'API Météo")
	} else {
		//On déserialize la réponse
		var apiResp ApiResponse

		err = json.Unmarshal([]byte(reponse), &apiResp)
		if err != nil {
			return &MeteoResponse{}, errors.New("Erreur lors de la désérialisation de la réponse")
		} else {
			//On calcule la moyenne des températures
			var somme float64
			for i := 0; i < len(apiResp.Hourly.Temperature2m); i++ {
				somme += apiResp.Hourly.Temperature2m[i]
			}
			moyenne := somme / float64(len(apiResp.Hourly.Temperature2m))
			moyenne = math.Ceil(moyenne*100) / 100

			return &MeteoResponse{
				Latitude:    apiResp.Latitude,
				Longitude:   apiResp.Longitude,
				Temperature: moyenne,
			}, nil

		}
	}
	return &MeteoResponse{}, errors.New("Erreur lors de la récupération des données")
}

// Permet de récupérer les prévisions de précipitations
// les arguments sont la latitude et la longitude de l'endroit
func GetPrevisions(latitude string, longitude string) (*PrevisionsResponse, error) {
	urlPrecipitation := "https://api.open-meteo.com/v1/forecast?forecast_days=2&latitude=" + latitude + "&longitude=" + longitude + "&hourly=precipitation_probability"

	err, reponse := apiGetCall(urlPrecipitation)
	if err != nil {
		return &PrevisionsResponse{}, errors.New("Erreur lors de la requête à l'API Météo")
	} else {
		var apiResp ApiResponsePrecipitation
		err = json.Unmarshal([]byte(reponse), &apiResp)
		if err != nil {
			return &PrevisionsResponse{}, errors.New("Erreur lors de la désérialisation de la réponse")
		} else {
			//On calcule la moyenne des précipitations
			var somme float64
			for _, prob := range apiResp.Hourly.PrecipitationProbability {
				somme += prob
			}
			moyenne := somme / float64(len(apiResp.Hourly.PrecipitationProbability))
			moyenne = math.Ceil(moyenne*100) / 100

			var plagesHorairesRecommandees []SortTonParapluieDetails
			for i, prob := range apiResp.Hourly.PrecipitationProbability {
				if prob >= 50 {
					plageHoraire := SortTonParapluieDetails{
						Time:                     apiResp.Hourly.Time[i],
						PrecipitationProbability: prob,
					}
					plagesHorairesRecommandees = append(plagesHorairesRecommandees, plageHoraire)
				}
			}
			return &PrevisionsResponse{
				Latitude:              apiResp.Latitude,
				Longitude:             apiResp.Longitude,
				MoyennePrecipitations: moyenne,
				Sort_ton_parapluie:    plagesHorairesRecommandees,
			}, nil
		}
	}
	return &PrevisionsResponse{}, errors.New("Erreur lors de la récupération des données")
}
