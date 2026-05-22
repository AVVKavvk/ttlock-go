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

// UsersList godoc
//
//	@Summary		Get users list
//	@Description	Get a list of TTLock users
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.UserListRequestParams	true	"Users list request params"
//	@Success		202		{object}	interface{}				"Users list fetched successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/users/list [post]
func UsersList(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "Users list request received")

	var body ttt.UserListRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	users := ttlock.Users{}

	response, err := users.List(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Accepted(ctx, "Users list fetched successfully", response)
}

// UserRegister godoc
//
//	@Summary		Register user
//	@Description	Register a new TTLock user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.UserRegisterRequestParams	true	"User register request params"
//	@Success		202		{object}	interface{}					"User registered successfully"
//	@Failure		400		{object}	interface{}					"Bad request"
//	@Failure		500		{object}	interface{}					"Internal server error"
//	@Router			/api/users/register [post]
func UserRegister(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "User register request received")

	var body ttt.UserRegisterRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	users := ttlock.Users{}

	response, err := users.Register(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Accepted(ctx, "User registered successfully", response)
}

// UserPasswordReset godoc
//
//	@Summary		Reset user password
//	@Description	Reset TTLock user password
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.UserPasswordResetRequestParams	true	"User password reset request params"
//	@Success		202		{object}	interface{}						"User password reset successfully"
//	@Failure		400		{object}	interface{}						"Bad request"
//	@Failure		500		{object}	interface{}						"Internal server error"
//	@Router			/api/users/password-reset [post]
func UserPasswordReset(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "User password reset request received")

	var body ttt.UserPasswordResetRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()

	users := ttlock.Users{}

	response, err := users.PasswordReset(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Accepted(ctx, "User password reset successfully", response)
}

// UserDelete godoc
//
//	@Summary		Delete user
//	@Description	Delete a TTLock user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.UserDeleteRequestParams	true	"User delete request params"
//	@Success		202		{object}	interface{}				"User deleted successfully"
//	@Failure		400		{object}	interface{}				"Bad request"
//	@Failure		500		{object}	interface{}				"Internal server error"
//	@Router			/api/users/delete [post]
func UserDelete(ctx echo.Context) error {
	_, rid := utils.GetRequestContextAndIdFromEchoContext(ctx)

	log.Infoxf(&logger.XFields{"rid": rid}, "User delete request received")

	var body ttt.UserDeleteRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}
	body.Date = time.Now().UnixMilli()
	users := ttlock.Users{}

	response, err := users.Delete(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Accepted(ctx, "User deleted successfully", response)
}
