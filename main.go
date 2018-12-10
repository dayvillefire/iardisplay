package main // import "github.com/dayvillefire/iardisplay"

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"

	"github.com/dayvillefire/iardisplay/config"
	"github.com/elastic/apm-agent-go/module/apmgin"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	cadmonitor "github.com/jbuchbinder/cadmonitor/monitor"
	"github.com/jbuchbinder/iarapi"
	"github.com/natefinch/lumberjack"
)

var (
	Apm        = flag.Bool("apm", false, "Use apm")
	ConfigFile = flag.String("config-file", "./display.yml", "App configuration file")
	Debug      = flag.Bool("debug", false, "Enable debugging (overrides config)")
	Daemonize  = flag.Bool("daemon", false, "Run as daemon")
	LogFile    = flag.String("log", "./display.log", "Log file (when run as daemon)")
	CPUProfile = flag.String("cpu-profile", "", "Write cpu profile to file")

	cacheStatusChan     = make(chan bool)
	cacheStatusQuitChan = make(chan bool)
	shutdownChannel     = make(chan os.Signal, 1)
	iar                 iarapi.IamRespondingAPI
	cad                 cadmonitor.CadMonitor
	cadStatusCache      CadCallStatusCache
	hostname            string
	Version             string
	VERSION             string
)

func main() {
	flag.Parse()

	Version = VERSION

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	if *CPUProfile != "" {
		f, err := os.Create(*CPUProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *Daemonize && *LogFile != "" {
		// Fix logging
		log.SetOutput(LockedWriter{&lumberjack.Logger{
			Filename:   *LogFile,
			MaxSize:    1024, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
		}, &sync.Mutex{}})
	}

	c, err := config.LoadConfigWithDefaults(*ConfigFile)
	if err != nil {
		panic(err)
	}
	if c == nil {
		panic("UNABLE TO LOAD CONFIG")
	}
	config.Config = c

	if *Debug {
		log.Print("Overriding existing debug configuration")
		config.Config.Debug = true
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	log.Print("Logging into IAR")
	iar.Debug = c.Debug
	iar.Login(c.Login.Iar.Agency, c.Login.Iar.Username, c.Login.Iar.Password)
	if err != nil {
		panic(err)
	}

	log.Print("Configuring CAD interface")
	cad, err = cadmonitor.GetCadMonitor(c.Login.Cad.Monitor)
	if err != nil {
		panic(err)
	}
	err = cad.ConfigureFromValues(map[string]string{
		"fdid":    c.Login.Cad.FDID,
		"baseUrl": c.Login.Cad.BaseURL,
	})
	if err != nil {
		panic(err)
	}

	log.Print("Logging into CAD system")
	err = cad.Login(c.Login.Cad.Username, c.Login.Cad.Password)
	if err != nil {
		panic(err)
	}

	log.Print("Initializing CAD status cache")
	cadStatusCache = CadCallStatusCache{
		Monitor:        cad,
		ExpiryDuration: time.Duration(c.Login.Cad.CacheDuration) * time.Second,
	}

	application()
}

func application() {
	//hostname, _ := os.Hostname()

	log.Printf("Initializing web services")
	m := gin.New()
	m.Use(gin.Logger())
	if *Apm {
		m.Use(apmgin.Middleware(m))
	} else {
		m.Use(gin.Recovery())
	}

	// Enable gzip compression
	m.Use(gzip.Gzip(gzip.DefaultCompression))

	initApi(m)

	log.Print("[static] Initializing with local resources")
	m.Use(static.Serve("/", static.LocalFile(config.Config.Paths.BasePath+string(os.PathSeparator)+"ui", false)))
	m.StaticFile("/", config.Config.Paths.BasePath+string(os.PathSeparator)+"ui"+string(os.PathSeparator)+"index.html")

	go func() {
		log.Printf("Initializing display on :%d", config.Config.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), m); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		for {
			select {
			default:
				cad.GetActiveCalls()
				iar.GetNowRespondingWithSort()
				time.Sleep(time.Duration(5) * time.Second)
			case <-shutdownChannel:
				// stop
				return
			}
		}
	}()

	// Catch signals and termination
	signal.Notify(shutdownChannel, os.Interrupt)
	signal.Notify(shutdownChannel, syscall.SIGTERM)
	log.Println(<-shutdownChannel)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		//c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
