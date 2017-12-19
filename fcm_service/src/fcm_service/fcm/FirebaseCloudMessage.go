package fcm

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest.git"
	"github.com/pkg/errors"

	"fcm_service/app_config"
	"fcmservice"
)

func SendNotifyToOne(appServerKey string, noti *fcmservice.TFCMMessage, deviceToken fcmservice.TDeviceToken, isProxy bool, proxyHost string, proxyPort int) (resp *fcmservice.TResponse, err error) {

	var data string
	data = "{\"to\":\"" + string(deviceToken) + "\","

	isDataValid := false

	if noti.NotiPayload != nil {
		data += "\"notification\":{" +
			"\"title\":\"" + noti.NotiPayload.Title + "\"," +
			"\"body\":\"" + noti.NotiPayload.Body + "\""

		if strings.Compare(noti.NotiPayload.Icon, "") != 0 {
			data += ",\"icon\":\"" + noti.NotiPayload.Icon + "\""
		}

		if strings.Compare(noti.NotiPayload.ClickAction, "") != 0 {
			data += ",\"click_action\":\"" + noti.NotiPayload.ClickAction + "\""
		}

		if noti.NotiPayload.Data != nil {
			for _, kv := range noti.NotiPayload.Data {
				data += ",\"" + kv.Key + "\":\"" + kv.Value + "\""
			}
		}

		data += "}"

		isDataValid = true
	}

	if noti.DataPayload != nil {
		if noti.NotiPayload != nil {
			data += ",\"data\":{"
		} else {
			data += "\"data\":{"
		}

		if noti.DataPayload.Data != nil {
			for i, kv := range noti.DataPayload.Data {
				if i < len(noti.DataPayload.Data)-1 {
					data += "\"" + kv.Key + "\":\"" + kv.Value + "\","
				} else {
					data += "\"" + kv.Key + "\":\"" + kv.Value + "\""
				}
			}
		}

		data += "}"

		isDataValid = true
	}

	data += "}"

	if !isDataValid {
		return nil, errors.New("Invalid Parameters")
	}
	fmt.Println("SendNotifyToOne : ", data)

	var request *gorequest.SuperAgent
	if isProxy {
		request = gorequest.New().Proxy(proxyHost + ":" + strconv.Itoa(proxyPort))
	} else {
		request = gorequest.New()
	}

	httpResp, httpBody, errs := request.Post(app_config.AppConfig.FCM_URL).
		Set("Content-Type", "application/json").
		Set("Authorization", "key="+appServerKey).
		Send(data).
		End()

	if len(errs) > 0 {
		//fmt.Println("Unexpected errors: %s", errs)
		return nil, errs[0]
	}

	var buf bytes.Buffer
	httpResp.Header.WriteSubset(&buf, nil)
	//
	resp = new(fcmservice.TResponse)
	resp.StatusCode = int32(httpResp.StatusCode)
	resp.Header = buf.String()
	resp.Body = httpBody

	request.ClearSuperAgent()

	return resp, nil
}

func SendNotiToMulti(appServerKey string, noti *fcmservice.TFCMMessage, tokenList []string, isProxy bool, proxyHost string, proxyPort int) (resp *fcmservice.TResponse, err error) {

	var listDiviceToken string = ""

	for i, token := range tokenList {
		if i > 0 {
			listDiviceToken += ",\"" + string(token) + "\""
		} else {
			listDiviceToken += "\"" + string(token) + "\""
		}
	}

	var data string
	data = "{\"registration_ids\":[" + listDiviceToken + "],"

	isDataValid := false

	if noti.NotiPayload != nil {
		data += "\"notification\":{" +
			"\"title\":\"" + noti.NotiPayload.Title + "\"," +
			"\"body\":\"" + noti.NotiPayload.Body + "\""

		if strings.Compare(noti.NotiPayload.Icon, "") != 0 {
			data += ",\"icon\":\"" + noti.NotiPayload.Icon + "\""
		}

		if strings.Compare(noti.NotiPayload.ClickAction, "") != 0 {
			data += ",\"click_action\":\"" + noti.NotiPayload.ClickAction + "\""
		}

		if noti.NotiPayload.Data != nil {
			for _, kv := range noti.NotiPayload.Data {
				data += ",\"" + kv.Key + "\":\"" + kv.Value + "\""
			}
		}

		data += "}"

		isDataValid = true
	}

	if noti.DataPayload != nil {
		if noti.NotiPayload != nil {
			data += ",\"data\":{"
		} else {
			data += "\"data\":{"
		}

		if noti.DataPayload.Data != nil {
			for i, kv := range noti.DataPayload.Data {
				if i < len(noti.DataPayload.Data)-1 {
					data += "\"" + kv.Key + "\":\"" + kv.Value + "\","
				} else {
					data += "\"" + kv.Key + "\":\"" + kv.Value + "\""
				}
			}
		}

		data += "}"

		isDataValid = true
	}

	data += "}"

	if !isDataValid {
		return nil, errors.New("Invalid Parameters")
	}
	fmt.Println("SendNotiToMulti : ", data)

	var request *gorequest.SuperAgent
	if isProxy {
		request = gorequest.New().Proxy(proxyHost + ":" + strconv.Itoa(proxyPort))
	} else {
		request = gorequest.New()
	}

	httpResp, httpBody, errs := request.Post(app_config.AppConfig.FCM_URL).
		Set("Content-Type", "application/json").
		Set("Authorization", "key="+appServerKey).
		Send(data).
		End()

	if len(errs) > 0 {
		fmt.Println("Unexpected errors: %s", errs)
		return nil, errs[0]
	}

	var buf bytes.Buffer
	httpResp.Header.WriteSubset(&buf, nil)

	resp = new(fcmservice.TResponse)
	resp.StatusCode = int32(httpResp.StatusCode)
	resp.Header = buf.String()
	resp.Body = httpBody

	request.ClearSuperAgent()

	return resp, nil
}

func SendNotiToTopic(appServerKey string, topic string, condition string, noti *fcmservice.TFCMMessage, isProxy bool, proxyHost string, proxyPort int) (resp *fcmservice.TResponse, err error) {
	var data string

	if strings.Compare(topic, "") != 0 {
		data = "{\"to\":\"" + topic + "\","
	} else if strings.Compare(condition, "") != 0 {
		data = "{\"condition\":\"" + condition + "\","
	} else {
		return nil, errors.New("Invalid Parameters")
	}

	isDataValid := false

	if noti.NotiPayload != nil {
		data += "\"notification\":{" +
			"\"title\":\"" + noti.NotiPayload.Title + "\"," +
			"\"body\":\"" + noti.NotiPayload.Body + "\""

		if strings.Compare(noti.NotiPayload.Icon, "") != 0 {
			data += ",\"icon\":\"" + noti.NotiPayload.Icon + "\""
		}

		if strings.Compare(noti.NotiPayload.ClickAction, "") != 0 {
			data += ",\"click_action\":\"" + noti.NotiPayload.ClickAction + "\""
		}

		if noti.NotiPayload.Data != nil {
			for _, kv := range noti.NotiPayload.Data {
				data += ",\"" + kv.Key + "\":\"" + kv.Value + "\""
			}
		}

		data += "}"

		isDataValid = true
	}

	if noti.DataPayload != nil {
		if noti.NotiPayload != nil {
			data += ",\"data\":{"
		} else {
			data += "\"data\":{"
		}

		if noti.DataPayload.Data != nil {
			for i, kv := range noti.DataPayload.Data {
				if i < len(noti.DataPayload.Data)-1 {
					data += "\"" + kv.Key + "\":\"" + kv.Value + "\","
				} else {
					data += "\"" + kv.Key + "\":\"" + kv.Value + "\""
				}
			}
		}

		data += "}"

		isDataValid = true
	}

	data += "}"

	fmt.Println("SendNotiToTopic : ", data)

	if !isDataValid {
		return nil, errors.New("Invalid Parameters")
	}

	var request *gorequest.SuperAgent
	if isProxy {
		request = gorequest.New().Proxy(proxyHost + ":" + strconv.Itoa(proxyPort))
	} else {
		request = gorequest.New()
	}

	httpResp, httpBody, errs := request.Post(app_config.AppConfig.FCM_URL).
		Set("Content-Type", "application/json").
		Set("Authorization", "key="+appServerKey).
		Send(data).
		End()

	if len(errs) > 0 {
		fmt.Println("Unexpected errors: %s", errs)
		return nil, errs[0]
	}

	var buf bytes.Buffer
	httpResp.Header.WriteSubset(&buf, nil)

	resp = new(fcmservice.TResponse)
	resp.StatusCode = int32(httpResp.StatusCode)
	resp.Header = buf.String()
	resp.Body = httpBody

	request.ClearSuperAgent()

	return resp, nil
}
