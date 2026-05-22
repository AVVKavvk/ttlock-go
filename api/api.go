package api

import (
	"github.com/AVVKavvk/ttlock/constants"
	logging "github.com/AVVKavvk/ttlock/utils/logger"
)

var (
	log *logging.Logger
)

func init() {
	// Register this package's logger initializer
	logging.RegisterPackageLogger(func() {
		log = logging.DefaultV1Context.GetLogger(
			constants.SERVICE_NAME_FOR_LOGS+".api",
			logging.LevelDebug,
		)
	})
}
