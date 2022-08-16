package usecase

import (
	"time"

	domain "github.com/getsticker/domain"

	logg "github.com/sirupsen/logrus"
)

type stickerUsecase struct {
	stickerRepo    domain.StickerRepository
	contextTimeout time.Duration
}

// NewStickerUsecase inject the dependency for stickerusecase
func NewStickerUsecase(repo domain.StickerRepository) domain.StickerUsecase {
	return &stickerUsecase{
		stickerRepo: repo,
	}
}

func (service *stickerUsecase) GetStickers() (*[]domain.Sticker, error) {

	stickers, err := service.stickerRepo.GetStickers()
	if err != nil {

		logg.Error(err)
		return nil, err

	}

	return stickers, nil
}
