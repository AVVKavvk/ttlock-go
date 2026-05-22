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

// LockInitialize godoc
//
//	@Summary		Initialize a TTLock
//	@Description	Initialize a TTLock device and register it in the system
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockInitializeRequest	true	"Lock initialize payload"
//	@Success		200		{object}	interface{}				"Lock initialized successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/locks/initialize [post]
func LockInitialize(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock initialize request received")

	var body ttt.LockInitializeRequest
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()

	lock := ttlock.Lock{}

	response, err := lock.Initialize(body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Initialized", response)
}

// LockList godoc
//
//	@Summary		Get lock list
//	@Description	Get a list of TTLocks
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockListRequestParams	true	"Lock list request params"
//	@Success		200		{object}	interface{}				"Lock list fetched successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/locks/list [post]
func LockList(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock list request received")

	var body ttt.LockListRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}

	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.List(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock List", response)
}

// LockDetail godoc
//
//	@Summary		Get lock detail
//	@Description	Get TTLock details
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockDetailsRequestParams	true	"Lock detail request params"
//	@Success		200		{object}	interface{}					"Lock detail fetched successfully"
//	@Failure		400		{object}	interface{}					"Bad request"
//	@Failure		500		{object}	interface{}					"Internal server error"
//	@Router			/api/locks/detail [post]
func LockDetail(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock detail request received")

	var body ttt.LockDetailsRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.Detail(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Detail", response)
}

// LockDelete godoc
//
//	@Summary		Delete lock
//	@Description	Delete a TTLock
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockDeleteRequestParams	true	"Lock delete request params"
//	@Success		200		{object}	interface{}				"Lock deleted successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/locks/delete [post]
func LockDelete(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock delete request received")

	var body ttt.LockDeleteRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.Delete(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Deleted", response)
}

// LockUpdate godoc
//
//	@Summary		Update lock
//	@Description	Update TTLock information
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockUpdateRequestParams	true	"Lock update request params"
//	@Success		200		{object}	interface{}				"Lock updated successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/locks/update [post]
func LockUpdate(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock update request received")

	var body ttt.LockUpdateRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.Update(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Updated", response)
}

// LockRename godoc
//
//	@Summary		Rename lock
//	@Description	Rename a TTLock
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockRenameRequestParams	true	"Lock rename request params"
//	@Success		200		{object}	interface{}				"Lock renamed successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/locks/rename [post]
func LockRename(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock rename request received")

	var body ttt.LockRenameRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.Rename(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Renamed", response)
}

// LockChangeAdminPasscode godoc
//
//	@Summary		Change lock admin passcode
//	@Description	Change TTLock admin passcode
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockChangeAdminPasscodeRequestParams	true	"Lock admin passcode request params"
//	@Success		200		{object}	interface{}								"Lock admin passcode changed successfully"
//	@Failure		400		{object}	interface{}								"Bad request"
//	@Failure		500		{object}	interface{}								"Internal server error"
//	@Router			/api/locks/change-admin-passcode [post]
func LockChangeAdminPasscode(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock change admin passcode request received")

	var body ttt.LockChangeAdminPasscodeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.ChangeAdminPasscode(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Admin Passcode Changed", response)
}

// LockAutoLockTime godoc
//
//	@Summary		Set lock auto lock time
//	@Description	Update TTLock auto lock time
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockAutoLockTimeRequestParams	true	"Lock auto lock time request params"
//	@Success		200		{object}	interface{}							"Lock auto lock time updated successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/locks/auto-lock-time [post]
func LockAutoLockTime(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock auto lock time request received")

	var body ttt.LockAutoLockTimeRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.AutoLockTime(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Auto Lock Time Changed", response)
}

// LockListEKeys godoc
//
//	@Summary		Get lock eKeys list
//	@Description	Get all eKeys associated with a TTLock
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.LockListEKeyRequestParams	true	"Lock eKeys list request params"
//	@Success		200		{object}	interface{}					"Lock eKeys fetched successfully"
//	@Failure		400		{object}	interface{}					"Bad request"
//	@Failure		500		{object}	interface{}					"Internal server error"
//	@Router			/api/locks/keys [post]
func LockListEKeys(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock list ekeys request received")

	var body ttt.LockListEKeyRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.ListEKeys(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock EKeys List", response)
}

// LockListPasscodes godoc
//
//	@Summary		Get lock passcodes list
//	@Description	Get all passcodes associated with a TTLock
//	@Tags			Locks
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.PasscodeListRequestParams	true	"Lock passcodes list request params"
//	@Success		200		{object}	interface{}					"Lock passcodes fetched successfully"
//	@Failure		400		{object}	interface{}					"Bad request"
//	@Failure		500		{object}	interface{}					"Internal server error"
//	@Router			/api/locks/passcodes [post]
func LockListPasscodes(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Lock list passcodes request received")

	var body ttt.PasscodeListRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	lock := ttlock.Lock{}

	response, err := lock.ListPasscodes(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Lock Passcodes List", response)
}
