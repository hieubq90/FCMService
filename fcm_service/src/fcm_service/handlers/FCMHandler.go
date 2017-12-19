package handlers

import "fcmservice"

type FCMHandler struct {}

func NewFCMHandler() *FCMHandler {
	handler := new(FCMHandler)
	return handler
}

func (h *FCMHandler) AddDeviceToken(phone string, deviceToken fcmservice.TDeviceToken) (r bool, err error) {
	return
}

func (h *FCMHandler) AddListDeviceToken(phone string, tokenList fcmservice.TDeviceTokenList) (r bool, err error) {
	return
}

func (h *FCMHandler) NotiToDeviceToken(message *fcmservice.TFCMMessage, deviceToken fcmservice.TDeviceToken) (r *fcmservice.TResponse, err error) {
	return
}

func (h *FCMHandler) NotiToPhone(message *fcmservice.TFCMMessage, phone string) (r *fcmservice.TResponse, err error) {
	return
}

func (h *FCMHandler) NotiToTopic(topic string, condition string, message *fcmservice.TFCMMessage) (r *fcmservice.TResponse, err error) {
	return
}