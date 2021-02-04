package entities

// Values return values of indicator
type Values struct {
	Fecha string  `json:"fecha"`
	Valor float64 `json:"valor"`
}

// Indicator shows values arrays
type Indicator struct {
	Codigo       string   `json:"codigo"`
	Nombre       string   `json:"nombre"`
	UnidadMedida string   `json:"unidad_medida"`
	Serie        []Values `json:"serie"`
}
