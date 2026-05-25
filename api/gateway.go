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

// GatewayList godoc
//
//	@Summary		Get gateway list
//	@Description	Get all TTLock gateways
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayListRequestParams	true	"Gateway list request params"
//	@Success		200		{object}	interface{}					"Gateway list fetched successfully"
//	@Failure		400		{object}	interface{}					"Bad request"
//	@Failure		500		{object}	interface{}					"Internal server error"
//	@Router			/api/gateway/list [post]
func GatewayList(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway list request received")

	var body ttt.GatewayListRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}

	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.List(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway list fetched successfully", response)
}

// GatewayDelete godoc
//
//	@Summary		Delete gateway
//	@Description	Delete a TTLock gateway
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayDeleteRequestParams	true	"Gateway delete request params"
//	@Success		200		{object}	interface{}						"Gateway deleted successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/gateway/delete [post]
func GatewayDelete(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway delete request received")

	var body ttt.GatewayDeleteRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.Delete(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway deleted successfully", response)
}

// GatewayRename godoc
//
//	@Summary		Rename gateway
//	@Description	Rename a TTLock gateway
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayRenameRequestParams	true	"Gateway rename request params"
//	@Success		200		{object}	interface{}						"Gateway renamed successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/gateway/rename [post]
func GatewayRename(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway rename request received")

	var body ttt.GatewayRenameRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.Rename(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway renamed successfully", response)
}

// GatewayListByLock godoc
//
//	@Summary		Get gateways by lock
//	@Description	Get gateways associated with a TTLock
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayListByLockRequestParams	true	"Gateway list by lock request params"
//	@Success		200		{object}	interface{}							"Gateway list by lock fetched successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/gateway/list-by-lock [post]
func GatewayListByLock(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway list by lock request received")

	var body ttt.GatewayListByLockRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.ListByLock(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway list by lock fetched successfully", response)
}

// GatewayListLocks godoc
//
//	@Summary		Get locks by gateway
//	@Description	Get locks connected to a gateway
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayListLockRequestParams	true	"Gateway list locks request params"
//	@Success		200		{object}	interface{}						"Gateway locks fetched successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/gateway/list-locks [post]
func GatewayListLocks(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway list locks request received")

	var body ttt.GatewayListLockRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.ListLocks(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway list locks fetched successfully", response)
}

// GatewayListDevices godoc
//
//	@Summary		Get gateway devices
//	@Description	Get devices connected to a gateway
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayListDeviceRequestParams	true	"Gateway list devices request params"
//	@Success		200		{object}	interface{}							"Gateway devices fetched successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/gateway/list-devices [post]
func GatewayListDevices(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway list devices request received")

	var body ttt.GatewayListDeviceRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.ListDevices(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway list devices fetched successfully", response)
}

// GatewayDetail godoc
//
//	@Summary		Get gateway details
//	@Description	Get TTLock gateway details
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayDetailsRequestParams	true	"Gateway details request params"
//	@Success		200		{object}	interface{}						"Gateway details fetched successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/gateway/detail [post]
func GatewayDetail(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway detail request received")

	var body ttt.GatewayDetailsRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.Detail(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway detail fetched successfully", response)
}

// GatewayUploadDetails godoc
//
//	@Summary		Upload gateway details
//	@Description	Upload TTLock gateway details
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayUploadDetailRequestParams	true	"Gateway upload details request params"
//	@Success		200		{object}	interface{}							"Gateway details uploaded successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/gateway/upload-details [post]
func GatewayUploadDetails(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway upload details request received")

	var body ttt.GatewayUploadDetailRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.UploadDetail(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway details uploaded successfully", response)
}

// GatewayCheckUpgrade godoc
//
//	@Summary		Check gateway upgrade
//	@Description	Check TTLock gateway firmware upgrade
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewayCheckUpgradeRequestParams	true	"Gateway check upgrade request params"
//	@Success		200		{object}	interface{}								"Gateway upgrade check fetched successfully"
//	@Failure		400		{object}	interface{}								"Bad request"
//	@Failure		500		{object}	interface{}								"Internal server error"
//	@Router			/api/gateway/check-upgrade [post]
func GatewayCheckUpgrade(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway check upgrade request received")

	var body ttt.GatewayCheckUpgradeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.CheckUpgrade(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway check upgrade fetched successfully", response)
}

// GatewaySetUpgradeMode godoc
//
//	@Summary		Set gateway upgrade mode
//	@Description	Set TTLock gateway firmware upgrade mode
//	@Tags			Gateway
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.GatewaySetUpgradeModeRequestParams	true	"Gateway set upgrade mode request params"
//	@Success		200		{object}	interface{}								"Gateway upgrade mode set successfully"
//	@Failure		400		{object}	interface{}								"Bad request"
//	@Failure		500		{object}	interface{}								"Internal server error"
//	@Router			/api/gateway/set-upgrade-mode [post]
func GatewaySetUpgradeMode(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Gateway set upgrade mode request received")

	var body ttt.GatewaySetUpgradeModeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	gateway := ttlock.Gateway{}
	response, err := gateway.SetUpgradeMode(&body)
	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}
	return lib.ContextV1Success(ctx, "Gateway set upgrade mode successfully", response)
}
