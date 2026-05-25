package server

import (
	"strings"

	"github.com/AVVKavvk/ttlock/api"
	"github.com/AVVKavvk/ttlock/constants"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	logging "github.com/AVVKavvk/ttlock/utils/logger"
)

var log *logging.Logger

// customRemoveTrailingSlash is an Echo middleware designed for pprof that
// conditionally removes trailing slashes from the request URL path. It
// excludes paths starting with "/debug/" to accommodate pprof endpoints.
func customRemoveTrailingSlash() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check if the request URL path starts with "/debug/"
			if strings.HasPrefix(c.Request().URL.Path, "/debug/") {
				// Skip trailing slash removal for URLs containing "/debug/"
				return next(c)
			}

			// Apply the standard RemoveTrailingSlash middleware for other URLs
			return echoMiddleware.RemoveTrailingSlash()(next)(c)
		}
	}
}

func AddRoutes(router *echo.Echo) {
	router.Pre(customRemoveTrailingSlash())
	router.Use(echoMiddleware.RequestID())

	users := router.Group("/api/users")
	{
		users.POST("/register", api.UserRegister)
		users.POST("/list", api.UsersList)
		users.POST("/delete", api.UserDelete)
		users.POST("/password-reset", api.UserPasswordReset)
	}

	locks := router.Group("/api/locks")
	{
		locks.POST("/initialize", api.LockInitialize)
		locks.POST("/list", api.LockList)
		locks.PATCH("details", api.LockDetail)
		locks.POST("/delete", api.LockDelete)
		locks.POST("/update", api.LockUpdate)
		locks.POST("/rename", api.LockRename)
		locks.POST("/change-admin-passcode", api.LockChangeAdminPasscode)
		locks.POST("/auto-lock-time", api.LockAutoLockTime)
		locks.POST("/keys", api.LockListEKeys)
		locks.POST("/passcodes", api.LockListPasscodes)
		locks.POST("/lock", api.LockLock)
		locks.POST("/unlock", api.LockUnlock)
		locks.POST("/query-open-state", api.LockQueryOpenState)
		locks.POST("/time", api.LockTime)
		locks.POST("/update-time", api.LockUpdateTime)
		locks.POST("/battery-status", api.LockBatteryStatus)

	}

	ekey := router.Group("/api/ekey")
	{
		ekey.POST("/send", api.EKeySend)
		ekey.POST("/list", api.EKeyList)
		ekey.POST("/get-one", api.EKeyGetOne)
		ekey.POST("/delete", api.EKeyDelete)
		ekey.POST("/freeze", api.EKeyFreeze)
		ekey.POST("/unfreeze", api.EKeyUnfreeze)
		ekey.POST("/update", api.EKeyUpdate)
		ekey.POST("/change-valid-time", api.EKeyChangeValidTime)
	}

	passcode := router.Group("/api/passcode")
	{
		passcode.POST("/types", api.PasscodeTypeList)
		passcode.POST("/generate-random", api.PasscodeGenerateRandom)
		passcode.POST("/generate-custom", api.PasscodeGenerateCustom)
		passcode.POST("/update", api.PasscodeUpdate)
		passcode.POST("/delete", api.PasscodeDelete)
	}
	gateway := router.Group("/api/gateway")
	{
		gateway.POST("/list", api.GatewayList)
		gateway.POST("/delete", api.GatewayDelete)
		gateway.POST("/rename", api.GatewayRename)
		gateway.POST("/list-by-lock", api.GatewayListByLock)
		gateway.POST("/list-locks", api.GatewayListLocks)
		gateway.POST("/list-devices", api.GatewayListDevices)
		gateway.POST("/detail", api.GatewayDetail)
		gateway.POST("/upload-details", api.GatewayUploadDetails)
		gateway.POST("/check-upgrade", api.GatewayCheckUpgrade)
		gateway.POST("/set-upgrade-mode", api.GatewaySetUpgradeMode)
	}

	auth := router.Group("/api/auth")
	{
		auth.POST("/access-token", api.AccessToken)
		auth.POST("/refresh-access-token", api.RefreshAccessToken)
	}

}

func init() {
	// Register this package's logger initializer
	logging.RegisterPackageLogger(func() {
		log = logging.DefaultV1Context.GetLogger(
			constants.SERVICE_NAME_FOR_LOGS+".server",
			logging.LevelDebug,
		)
	})
}
