package postgres

import (
	domain "github.com/getsticker/domain"

	"github.com/jinzhu/gorm"
	logg "github.com/sirupsen/logrus"
)

type postgresStickerRepository struct {
	Conn *gorm.DB
}
// NewpostgresStickerRepository inject dependency to postgresStickerRepository
func NewpostgresStickerRepository(Conn *gorm.DB) domain.StickerRepository {
	return &postgresStickerRepository{Conn}
}
// BobbleDataBaseLayer stores connection 
type BobbleDataBaseLayer struct {
	DBService *gorm.DB
}
// GetStickers perform db operation 
func (database *postgresStickerRepository) GetStickers() (*[]domain.Sticker,error) {

	var sticker *[]domain.Sticker
	query := "SELECT * FROM sticker order by time DESC"
	err := database.Conn.Raw(query).Scan(&sticker).Error
	if err != nil {
		logg.Error(err)
		return  nil, err

	}
	return sticker,nil
}
