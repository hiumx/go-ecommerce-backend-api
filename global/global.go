package global

import (
	"github.com/hiumx/go-ecommerce-backend-api/pkg/logger"
	"github.com/hiumx/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
