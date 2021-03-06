module github.com/dayvillefire/iardisplay

go 1.17

replace (
	github.com/dayvillefire/iardisplay/config => ./config
	github.com/jbuchbinder/cadmonitor => ../../jbuchbinder/cadmonitor
	github.com/jbuchbinder/cadmonitor/monitor => ../../jbuchbinder/cadmonitor/monitor
	github.com/jbuchbinder/iarapi => ../../jbuchbinder/iarapi
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dayvillefire/iardisplay/config v0.0.0-00010101000000-000000000000
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/contrib v0.0.0-20181101072842-54170a7b0b4b
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/jbuchbinder/cadmonitor/monitor v0.0.0-00010101000000-000000000000
	github.com/jbuchbinder/iarapi v0.0.0-00010101000000-000000000000
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.2.2 // indirect
	github.com/ugorji/go/codec v0.0.0-20181209151446-772ced7fd4c2 // indirect
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
