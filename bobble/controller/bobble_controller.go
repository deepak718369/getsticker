package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	bobble_interface "test.com/test/bobble/bobble_interface"
)

type BobbleController struct {
	BobbleService bobble_interface.BobbleService
}

func NewBobbleController(echoCtx *echo.Echo, bobble_service_object bobble_interface.BobbleService) {

	Bobble_Controller_object := &BobbleController{
		BobbleService: bobble_service_object,
	}
	echoCtx.GET("/v1/trendingStickers", Bobble_Controller_object.GetSticker)

}

func (bobble BobbleController) GetSticker(e echo.Context) error {
	fmt.Println("***controller***")
	requestContext := e.Request().Context()
	err, result := bobble.BobbleService.GetSticker(requestContext)
	if err != nil {
		return err
	}
	e.JSON(http.StatusOK, result)

	return nil

}
