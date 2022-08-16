package domain

// //go:generate mockgen -destination=mocks/ucasemock.go -package=mocks github.com/getsticker/domain StickerUsecase
//go:generate mockgen -destination=mocks/repomocks.go -package=mocks github.com/getsticker/domain StickerRepository
import (
	"time"
)
// Sticker model for storing sticker information 
type Sticker struct {
	ID      *int       `json:"id" gorm:"primary key"`
	Sticker   *string    `json:"sticker"`
	Time      *time.Time `json:"timing"`
	Priority  *string    `json:"priority"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
// StickerUsecase usecase interface 
type StickerUsecase interface {
	GetStickers() (*[]Sticker,error)
}
// StickerRepository repository interface 
type StickerRepository interface {
	GetStickers() (*[]Sticker,error)
}
