module github.com/dayvillefire/iardisplay

go 1.17

replace (
	github.com/dayvillefire/iardisplay/config => ./config
	github.com/jbuchbinder/cadmonitor => ../../jbuchbinder/cadmonitor
	github.com/jbuchbinder/cadmonitor/monitor => ../../jbuchbinder/cadmonitor/monitor
	github.com/jbuchbinder/iarapi => ../../jbuchbinder/iarapi
)

require (
	github.com/dayvillefire/iardisplay/config v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/contrib v0.0.0-20181101072842-54170a7b0b4b
	github.com/gin-gonic/gin v1.7.0
	github.com/jbuchbinder/cadmonitor/monitor v0.0.0-00010101000000-000000000000
	github.com/jbuchbinder/iarapi v0.0.0-00010101000000-000000000000
	github.com/natefinch/lumberjack v2.0.0+incompatible
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
