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

// PasscodeTypeList godoc
//
//	@Summary		Get passcode types
//	@Description	Get all TTLock passcode types
//	@Tags			Passcodes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	interface{}	"Passcode types fetched successfully"
//	@Failure		500	{object}	interface{}	"Internal server error"
//	@Router			/api/passcode/types [post]
func PasscodeTypeList(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Passcode type list request received")

	passcode := ttlock.Passcode{}

	response, err := passcode.Types()

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)
}

// PasscodeGenerateRandom godoc
//
//	@Summary		Generate random passcode
//	@Description	Generate a random TTLock passcode
//	@Tags			Passcodes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.PasscodeRandomRequestParams	true	"Random passcode request params"
//	@Success		200		{object}	interface{}						"Random passcode generated successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/passcode/generate-random [post]
func PasscodeGenerateRandom(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Passcode generate random request received")

	passcode := ttlock.Passcode{}

	var body ttt.PasscodeRandomRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	response, err := passcode.Random(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)
}

// PasscodeGenerateCustom godoc
//
//	@Summary		Generate custom passcode
//	@Description	Generate a custom TTLock passcode
//	@Tags			Passcodes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.PasscodeCustomRequestParams	true	"Custom passcode request params"
//	@Success		200		{object}	interface{}						"Custom passcode generated successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/passcode/generate-custom [post]
func PasscodeGenerateCustom(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Passcode generate request received")

	passcode := ttlock.Passcode{}

	var body ttt.PasscodeCustomRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	response, err := passcode.Custom(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)
}

// PasscodeUpdate godoc
//
//	@Summary		Update passcode
//	@Description	Update a TTLock passcode
//	@Tags			Passcodes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.PasscodeUpdateRequestParams	true	"Passcode update request params"
//	@Success		200		{object}	interface{}						"Passcode updated successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/passcode/update [post]
func PasscodeUpdate(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Passcode update request received")

	passcode := ttlock.Passcode{}

	var body ttt.PasscodeUpdateRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	response, err := passcode.Update(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)
}

// PasscodeDelete godoc
//
//	@Summary		Delete passcode
//	@Description	Delete a TTLock passcode
//	@Tags			Passcodes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.PasscodeDeleteRequestParams	true	"Passcode delete request params"
//	@Success		200		{object}	interface{}						"Passcode deleted successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/passcode/delete [post]
func PasscodeDelete(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)
	log.Infoxf(&logger.XFields{"rid": rid}, "Passcode delete request received")

	passcode := ttlock.Passcode{}

	var body ttt.PasscodeDeleteRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()

	response, err := passcode.Delete(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)
}
