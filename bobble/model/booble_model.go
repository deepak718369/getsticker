package model

import (
	"time"
)

type Sticker struct {
	Id         string    `json:"id"`
	Sticker    string    `json:"sticker"`
	Time       time.Time `json:"timing"`
	Priorities string    `json:"priority"`
}
