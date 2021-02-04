package entities

// Servicio de impuestos internos indicador sii.cl

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

// Get on board public API

// Attributes return attributes of company job posting
type Attributes struct {
	Title               string  `json:"title"`
	DescriptionHeadline string  `json:"description_headline"`
	FunctionsHeadline   string  `json:"functions_headline"`
	Functions           string  `json:"functions"`
	Remote              bool    `json:"remote"`
	RemoteModality      string  `json:"remote_modality"`
	RemoteZone          string  `json:"remote_zone"`
	Country             string  `json:"country"`
	MinSalary           float32 `json:"min_salary"`
	MaxSalary           float32 `json:"max_salary"`
	Modality            string  `json:"modality"`
	Seniority           string  `json:"seniority"`
	PublishedAt         int32   `json:"published_at"`
}

// Links that contains object of url to lookup detailed.
type Links struct {
	PublicURL string `json:"public_url"`
}

// Data that contains detailed array of atributes of companies
type Data struct {
	ID         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
	Links      Links      `json:"links"`
}

// GetOnBoard response data
type GetOnBoard struct {
	Data []Data `json:"data"`
}
