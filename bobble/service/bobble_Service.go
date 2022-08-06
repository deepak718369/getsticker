package service

import (
	"context"
	"fmt"

	bobble_interface "test.com/test/bobble/bobble_interface"
	"test.com/test/bobble/model"
)

type BoobleServiceImpl struct {
	bobble_data_layer bobble_interface.Datalayer
}

func NewBobbleService(bobble_data_layer bobble_interface.Datalayer) bobble_interface.BobbleService {

	return &BoobleServiceImpl{
		bobble_data_layer: bobble_data_layer,
	}
}
func (service *BoobleServiceImpl) GetSticker(ctx context.Context) (error, *[]model.Sticker) {

	fmt.Println("service layer ")
	err, result := service.bobble_data_layer.GetSticker(ctx)
	if err != nil {
		return err, nil

	}
	return err, result
}
