package models

type Item struct {
	Rid         string  `json:"rid" db:"rid"`
	OrderUId    string  `json:"order_uid" db:"order_uid"`
	ChrtId      int64   `json:"chrt_id" db:"chrt_id"`
	TrackNumber string  `json:"track_number" db:"track_number"`
	Price       float64 `json:"price" db:"price"`
	Name        string  `json:"name" db:"name"`
	Sale        float64 `json:"sale" db:"sale"`
	Size        string  `json:"size" db:"size"`
	TotalPrice  float64 `json:"total_price" db:"total_price"`
	NmId        int64   `json:"nm_id" db:"nm_id"`
	Brand       string  `json:"brand" db:"brand"`
	Status      int     `json:"status" db:"status"`
}
