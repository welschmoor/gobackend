package models

import "time"

type Listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Category    string    `json:"cat"`
	PicUrl      string    `json:"pu"`
	Description string    `json:"descr"`
	Price       int       `json:"price"` //price in cent
	CreatedAt   time.Time `json:"ca"`    //dont include it in json
	UpdatedAt   time.Time `json:"ua"`    //dont include it in json
}
