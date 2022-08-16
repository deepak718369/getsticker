package http_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	domain "github.com/getsticker/domain"
	mocks "github.com/getsticker/domain/mocks"
	stickerhttp "github.com/getsticker/sticker/delivery/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSticker(t *testing.T) {

	var mockSticker domain.Sticker
	err := faker.FakeData(&mockSticker)
	assert.NoError(t, err)

	mockUCase := new(mocks.StickerUsecase)
	mockListSticker := make([]domain.Sticker, 0)
	mockListSticker = append(mockListSticker, mockSticker)
	mockUCase.On("GetStickers").Return(&mockListSticker,nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/v1/trendingSticker", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	handler := stickerhttp.StickerHandler{
		StickerUcase: mockUCase,
	}
	err = handler.GetStickers(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
}
