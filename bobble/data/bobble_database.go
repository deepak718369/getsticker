package data

import (
	"context"
	"fmt"

	bobble_interface "test.com/test/bobble/bobble_interface"
	"test.com/test/bobble/model"
)

type BobbleDataBaseLayer struct {
	DBService DBService
}

func NewBobbleDatabaseLayer() bobble_interface.Datalayer {

	return &BobbleDataBaseLayer{
		DBService: DBService{},
	}
}
func (database *BobbleDataBaseLayer) GetSticker(ctx context.Context) (error, *[]model.Sticker) {
	fmt.Println("database layer")
	var sticker []model.Sticker
	query := "SELECT * FROM sticker order by time DESC"
	result := database.DBService.GetDB().Raw(query).Scan(&sticker)
	if result.Error != nil {
		return err, nil
	}
	return nil, &sticker

}
