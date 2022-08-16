package http

import (
	domain "github.com/getsticker/domain"
	"github.com/labstack/echo"
	logg "github.com/sirupsen/logrus"
)

// StickerHandler takes usecase interface and implements it.
type StickerHandler struct {
	StickerUcase domain.StickerUsecase
}

// NewstickerHandler route to end point
func NewstickerHandler(echoCtx *echo.Echo, stickerUsecase domain.StickerUsecase) {

	handler := &StickerHandler{

		StickerUcase: stickerUsecase,
	}

	echoCtx.GET("/v1/trendingStickers", handler.GetStickers)

}
// GetStickers service layer for getsticker 
func (handler StickerHandler) GetStickers(e echo.Context) error {

	stickers, err := handler.StickerUcase.GetStickers()
	if err != nil {

		logg.Error(err)
		return e.JSON(domain.GetStatusCode(err), domain.ResponseError{Message: err.Error()})
	}

	return e.JSON(domain.GetStatusCode(err), stickers)

}
