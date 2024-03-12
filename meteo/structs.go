package meteo

////////////////////////////////////////////////////////
// ---Partie pour récupérer les données de l'API appelée
////////////////////////////////////////////////////////

// La structure pour gérer les températures
type ApiResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    struct {
		Time          []string  `json:"time"`
		Temperature2m []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

// La struct pOur gérer les prévisions
type ApiResponsePrecipitation struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     float64 `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time                     string `json:"time"`
		PrecipitationProbability string `json:"precipitation_probability"`
	} `json:"hourly_units"`
	Hourly struct {
		Time                     []string  `json:"time"`
		PrecipitationProbability []float64 `json:"precipitation_probability"`
	} `json:"hourly"`
}

// /////////////////////////////////////////
// ---Partie pour les envoies à notre l'API
// /////////////////////////////////////////
type MeteoResponse struct {
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Temperature float64 `json:"temperature"`
}

// La structure pour gérer les prévisions
type PrevisionsResponse struct {
	Latitude              float64                   `json:"latitude"`
	Longitude             float64                   `json:"longitude"`
	MoyennePrecipitations float64                   `json:"moyenne_precicipitations"`
	Sort_ton_parapluie    []SortTonParapluieDetails `json:"sort_ton_parapluie"`
}

// La structure pour gérer les détails des prévisions
type SortTonParapluieDetails struct {
	Time                     string  `json:"time"`
	PrecipitationProbability float64 `json:"precipitations_probability"`
}
