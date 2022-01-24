module personFrame

go 1.16

require github.com/spf13/viper v1.10.1

replace (
	./personFrame/model/common => ./commonModel
	./personFrame/service/common => ./commonService
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-redis/redis/v8 v8.11.4
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/songzhibin97/gkit v1.1.4
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.20.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)
