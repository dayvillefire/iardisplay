module github.com/dayvillefire/iardisplay

go 1.23.0

toolchain go1.24.0

replace (
	github.com/dayvillefire/iardisplay/config => ./config
	github.com/jbuchbinder/cadmonitor => ../../jbuchbinder/cadmonitor
	github.com/jbuchbinder/cadmonitor/monitor => ../../jbuchbinder/cadmonitor/monitor
	github.com/jbuchbinder/iarapi => ../../jbuchbinder/iarapi
)

require (
	github.com/dayvillefire/iardisplay/config v0.0.0-20241102015704-760d5321c7cf
	github.com/gin-gonic/contrib v0.0.0-20250113154928-93b827325fec
	github.com/gin-gonic/gin v1.10.0
	github.com/jbuchbinder/cadmonitor/monitor v0.0.0-20250327183818-e07f6f02833a
	github.com/jbuchbinder/iarapi v0.0.0-20250221144523-57c4f48da04e
	github.com/natefinch/lumberjack v2.0.0+incompatible
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/PuerkitoBio/goquery v1.10.3 // indirect
	github.com/andybalholm/cascadia v1.3.3 // indirect
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.26.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/headzoo/surf v1.0.1 // indirect
	github.com/jbuchbinder/shims v0.0.0-20250315180801-ea13cafaf717 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.16.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
