package entities

import "time"

type Area struct {
	ID             int64      `json:"id"`
	ProprietarioID int64      `json:"proprietario_id"`
	Descricao      string     `json:"descricao"`
	AreaTotal      float64    `json:"area_total"`
	GeoJSON        string     `json:"geo_json"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated-at"`
}
