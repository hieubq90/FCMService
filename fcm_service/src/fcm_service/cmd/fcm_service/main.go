package main

import (
	"fcm_service/app_config"
	"os"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"fcm_service/service"
	"strconv"
)

func main() {
	if app_config.InitFromYAML() {
		// Init Protocol
		var protocolFactory thrift.TProtocolFactory
		switch app_config.AppConfig.Protocol {
		case "compact":
			protocolFactory = thrift.NewTCompactProtocolFactory()
		case "simplejson":
			protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		case "json":
			protocolFactory = thrift.NewTJSONProtocolFactory()
		case "binary", "":
			protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		default:
			fmt.Fprint(os.Stderr, "Invalid protocol specified", app_config.AppConfig.Protocol, "\n")
			//Usage()
			os.Exit(1)
		}

		// Init Transport Factory
		var transportFactory thrift.TTransportFactory
		if app_config.AppConfig.Buffered {
			transportFactory = thrift.NewTBufferedTransportFactory(8192)
		} else {
			transportFactory = thrift.NewTTransportFactory()
		}
		if app_config.AppConfig.Framed {
			transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
		}

		if err := service.RunFCMService(transportFactory, protocolFactory, app_config.AppConfig.Host + ":"+ strconv.Itoa(app_config.AppConfig.Port)); err != nil {
			fmt.Println("error running server:", err)
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}
}
