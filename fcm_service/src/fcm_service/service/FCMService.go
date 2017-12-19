package service

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"fcmservice"
	"fcm_service/handlers"
)

func RunFCMService(transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory,
	address string) error {

	var transport thrift.TServerTransport
	var err error


	transport, err = thrift.NewTServerSocket(address)


	if err != nil {
		return err
	}

	notifyHandler := handlers.NewFCMHandler()
	processor := fcmservice.NewFCMServiceProcessor(notifyHandler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("START SERVER ON " + address)

	return server.Serve()

}
