package handlers

import (
	"fmt"

	"fcmservice"
	"fcm_service/database"
	"fcm_service/fcm"
	"fcm_service/app_config"
)

type FCMHandler struct {}

func NewFCMHandler() *FCMHandler {
	handler := new(FCMHandler)
	return handler
}

func (h *FCMHandler) AddDeviceToken(phone string, deviceToken fcmservice.TDeviceToken) (r bool, err error) {
	err = database.AddNewToken(phone, string(deviceToken))
	if err == nil {
		r = true
	} else {
		r = false
	}
	return
}

func (h *FCMHandler) AddListDeviceToken(phone string, tokenList fcmservice.TDeviceTokenList) (r bool, err error) {
	return
}

func (h *FCMHandler) NotiToDeviceToken(message *fcmservice.TFCMMessage, deviceToken fcmservice.TDeviceToken) (err error) {
	go func() {
		fmt.Println("[=====> NotiToDeviceToken...]")

		resp, err := fcm.SendNotifyToOne(app_config.AppConfig.FCM_API_KEY, message, deviceToken, app_config.AppConfig.Proxy, app_config.AppConfig.ProxyHost, app_config.AppConfig.ProxyPort)

		/*fmt.Println("Response status: ", resp.StatusCode)
		fmt.Println("Response Header: ", resp.Header)
		fmt.Println("Response body: ", resp.Body)
		*/

		if err != nil {
			fmt.Println("func=NotiToDeviceToken | Notify=", message, " | DeviceToken=", deviceToken, " | Error=", err)
			fmt.Println("======>>>>>> NotiToDeviceToken complete, err=", err, "\n")
		} else if resp != nil {
			fmt.Println("func=NotiToDeviceToken | Notify=", message, " | DeviceToken=", deviceToken, " | StatusCode=", resp.StatusCode, " | ResponseBody=", resp.Body)
			fmt.Println("======>>>>>> NotiToDeviceToken complete, status code=", resp.StatusCode, "\n")
		} else {
			fmt.Println("KHONG TRA VE RESPONSE")
		}
	}()
	return
}

func (h *FCMHandler) NotiToPhone(message *fcmservice.TFCMMessage, phone string) (err error) {
	go func() {
		user, err := database.GetUserByPhone(phone)
		if err == nil && user != nil && len(user.Tokens) > 0 {
			fmt.Println("[=====> NotiToMultiDeviceToken...]")
			resp, err := fcm.SendNotiToMulti(app_config.AppConfig.FCM_API_KEY, message, user.Tokens, app_config.AppConfig.Proxy, app_config.AppConfig.ProxyHost, app_config.AppConfig.ProxyPort)

			if err != nil {
				fmt.Println("func=NotiToMultiDeviceToken | Notify=", message, " | DeviceTokenList=", user.Tokens, " | Error=", err)
				fmt.Println("======>>>>>> NotiToMultiDeviceToken complete, err=", err, "\n")
			} else if resp != nil {
				fmt.Println("func=NotiToMultiDeviceToken | Notify=", message, " | DeviceTokenList=", user.Tokens, " | StatusCode=", resp.StatusCode, " | ResponseBody=", resp.Body)
				fmt.Println("======>>>>>> NotiToMultiDeviceToken complete, status code=", resp.StatusCode, "\n")
			} else {
				fmt.Println("KHONG TRA VE RESPONSE")
			}
		}
	}()
	return
}

func (h *FCMHandler) NotiToTopic(topic string, condition string, message *fcmservice.TFCMMessage) (err error) {
	go func() {
		fmt.Println("[=====> NotiToTopic...]")
		resp, err := fcm.SendNotiToTopic(app_config.AppConfig.FCM_API_KEY, topic, condition, message, app_config.AppConfig.Proxy, app_config.AppConfig.ProxyHost, app_config.AppConfig.ProxyPort)

		/*fmt.Println("Response status: ", resp.StatusCode)
		fmt.Println("Response Header: ", resp.Header)
		fmt.Println("Response body: ", resp.Body)
		*/
		if err != nil {
			fmt.Println("func=NotiToTopic | Notify=", message, " | topic=", topic, " | condition=", condition, " | Error=", err)
			fmt.Println("======>>>>>> NotiToTopic complete, err=", err, "\n")
		} else if resp != nil {
			fmt.Println("func=NotiToTopic | Notify=", message, " | topic=", topic, " | condition=", condition, " | StatusCode=", resp.StatusCode, " | ResponseBody=", resp.Body)
			fmt.Println("======>>>>>> NotiToTopic complete, status code=", resp.StatusCode, "\n")
		} else {
			fmt.Println("KHONG TRA VE RESPONSE")
		}
	}()
	return
}