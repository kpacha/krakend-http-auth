package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	authgin "test/gin"

	"github.com/devopsfaith/krakend/proxy"
	"github.com/devopsfaith/krakend/router"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	"github.com/luraproject/lura/logging"
	"github.com/spf13/viper"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")

	configFile := flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	flag.Parse()

	parser := viper.New()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	var routerFactory router.Factory

	routerFactory = krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		ProxyFactory:   proxy.DefaultFactory(logger),
		Middlewares:    []gin.HandlerFunc{},
		Logger:         logger,
		HandlerFactory: authgin.HandlerFactory(krakendgin.EndpointHandler),
	})

	routerFactory.New().Run(serviceConfig)
}
