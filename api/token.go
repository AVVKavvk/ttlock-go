package api

import (
	"net/http"

	"github.com/AVVKavvk/ttlock/lib"
	"github.com/AVVKavvk/ttlock/ttlock"
	ttt "github.com/AVVKavvk/ttlock/ttlock/types"
	"github.com/labstack/echo/v4"
)

// AccessToken godoc
//
//	@Summary		Get access token
//	@Description	Generate TTLock access token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.AccessTokenRequestParams	true	"Access token request params"
//	@Success		200		{object}	interface{}					"Access token generated successfully"
//	@Failure		400		{object}	interface{}					"Bad request"
//	@Failure		500		{object}	interface{}					"Internal server error"
//	@Router			/api/auth/access-token [post]
func AccessToken(ctx echo.Context) error {
	var body ttt.AccessTokenRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}

	tokens := ttlock.Token{}

	response, err := tokens.GetAccessToken(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)

}

// RefreshAccessToken godoc
//
//	@Summary		Refresh access token
//	@Description	Refresh TTLock access token using refresh token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ttt.RefreshAccessTokenRequestParams	true	"Refresh access token request params"
//	@Success		200		{object}	interface{}							"Access token refreshed successfully"
//	@Failure		400		{object}	interface{}							"Bad request"
//	@Failure		500		{object}	interface{}							"Internal server error"
//	@Router			/api/auth/refresh-access-token [post]
func RefreshAccessToken(ctx echo.Context) error {
	var body ttt.RefreshAccessTokenRequestParams
	if err := ctx.Bind(&body); err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusBadRequest, err.Error(), ""))
	}

	tokens := ttlock.Token{}

	response, err := tokens.RefreshAccessToken(&body)

	if err != nil {
		return lib.ContextV1Err(ctx, lib.NewErrorCustom(http.StatusInternalServerError, err.Error(), ""))
	}

	return lib.ContextV1Success(ctx, "Success", response)
}
