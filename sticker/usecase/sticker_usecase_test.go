package usecase_test

import (
	"testing"

	"github.com/bxcodec/faker"
	domain "github.com/getsticker/domain"
	usecase "github.com/getsticker/sticker/usecase"

	mocks "github.com/getsticker/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTrendingStickers(t *testing.T) {

	mockStickerRepository := new(mocks.StickerRepository)

	var mockSticker domain.Sticker
	err := faker.FakeData(&mockSticker)
	assert.NoError(t, err)

	mockListSticker := make([]domain.Sticker, 0)

	mockListSticker = append(mockListSticker, mockSticker)

	t.Run("success", func(t *testing.T) {
		mockStickerRepository.On("GetStickers", mock.Anything, mock.Anything).Return(&mockListSticker,nil)
		
		stickerUseCase := usecase.NewStickerUsecase(mockStickerRepository)
		
		stickers,err := stickerUseCase.GetStickers()
		
		assert.NoError(t, err)
		assert.NotEmpty(t, stickers)
		mockStickerRepository.AssertExpectations(t)
	})

}
