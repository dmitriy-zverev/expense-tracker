package budget

type Budget struct {
	Month    int     `json:"month"`
	Year     int     `json:"year"`
	Category string  `json:"category"`
	Limit    float64 `json:"limit"`
}
