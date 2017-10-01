package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/devopsfaith/krakend/config/viper"
	"github.com/devopsfaith/krakend/logging/gologging"
	"github.com/devopsfaith/krakend/proxy"
	krakendgin "github.com/devopsfaith/krakend/router/gin"
	auth "github.com/kpacha/krakend-http-auth/gin"
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

	logger, err := gologging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         gin.Default(),
		ProxyFactory:   proxy.DefaultFactory(logger),
		Middlewares:    []gin.HandlerFunc{},
		Logger:         logger,
		HandlerFactory: auth.HandlerFactory(krakendgin.EndpointHandler),
	})

	routerFactory.New().Run(serviceConfig)
}
