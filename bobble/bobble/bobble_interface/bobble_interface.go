package bobble_interface

import (
	"context"

	"test.com/test/bobble/model"
)

type BobbleService interface {
	GetSticker(ctx context.Context) (error, *[]model.Sticker)
}
type Datalayer interface {
	GetSticker(ctx context.Context) (error, *[]model.Sticker)
}
