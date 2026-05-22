package api

import (
	"net/http"
	"time"

	"github.com/AVVKavvk/ttlock/lib"
	"github.com/AVVKavvk/ttlock/ttlock"
	ttt "github.com/AVVKavvk/ttlock/ttlock/types"
	"github.com/AVVKavvk/ttlock/utils"
	"github.com/AVVKavvk/ttlock/utils/logger"
	"github.com/labstack/echo/v4"
)

// EKeySend godoc
//
//	@Summary		Send eKey
//	@Description	Send an eKey to a user
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeySendRequestParams	true	"eKey send request params"
//	@Success		200		{object}	interface{}				"eKey sent successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/ekey/send [post]
func EKeySend(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey send request received")

	var body ttt.EKeySendRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.Send(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyList godoc
//
//	@Summary		Get eKey list
//	@Description	Get all eKeys
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyListRequestParams	true	"eKey list request params"
//	@Success		200		{object}	interface{}				"eKeys fetched successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/ekey/list [post]
func EKeyList(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey list request received")

	var body ttt.EKeyListRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.List(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyGetOne godoc
//
//	@Summary		Get one eKey
//	@Description	Get eKey details by ID
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyGetOneRequestParams	true	"eKey get one request params"
//	@Success		200		{object}	interface{}				"eKey fetched successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/ekey/get-one [post]
func EKeyGetOne(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey get one request received")

	var body ttt.EKeyGetOneRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.GetOne(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyDelete godoc
//
//	@Summary		Delete eKey
//	@Description	Delete an eKey
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyDeleteRequestParams	true	"eKey delete request params"
//	@Success		200		{object}	interface{}				"eKey deleted successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/ekey/delete [post]
func EKeyDelete(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey delete request received")

	var body ttt.EKeyDeleteRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.Delete(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyFreeze godoc
//
//	@Summary		Freeze eKey
//	@Description	Freeze an eKey
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyFreezeUnfreezeRequestParams	true	"eKey freeze request params"
//	@Success		200		{object}	interface{}							"eKey frozen successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/ekey/freeze [post]
func EKeyFreeze(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey freeze request received")

	var body ttt.EKeyFreezeUnfreezeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.Freeze(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyUnfreeze godoc
//
//	@Summary		Unfreeze eKey
//	@Description	Unfreeze an eKey
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyFreezeUnfreezeRequestParams	true	"eKey unfreeze request params"
//	@Success		200		{object}	interface{}							"eKey unfrozen successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/ekey/unfreeze [post]
func EKeyUnfreeze(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey unfreeze request received")

	var body ttt.EKeyFreezeUnfreezeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.Unfreeze(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyUpdate godoc
//
//	@Summary		Update eKey
//	@Description	Update eKey information
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyUpdateRequestParams	true	"eKey update request params"
//	@Success		200		{object}	interface{}				"eKey updated successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/ekey/update [post]
func EKeyUpdate(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey update request received")

	var body ttt.EKeyUpdateRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.Update(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}

// EKeyChangeValidTime godoc
//
//	@Summary		Change eKey valid time
//	@Description	Change eKey validity period
//	@Tags			EKeys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.EKeyChangeValidTimeRequestParams	true	"eKey valid time request params"
//	@Success		200		{object}	interface{}							"eKey valid time changed successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/ekey/change-valid-time [post]
func EKeyChangeValidTime(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "EKey change valid time request received")

	var body ttt.EKeyChangeValidTimeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	ekey := ttlock.EKeys{}

	response, err := ekey.ChangeValidTime(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return ctx.JSON(http.StatusOK, response)
}
