namespace go fcmservice

struct TKeyValue {
  1: string key;
  2: string value;
}

struct TNotificationPayload {
  1: string title;
  2: string body;
  3: optional string icon;
  4: optional list<TKeyValue> data;
  5: optional string click_action;
}

struct TDataPayload {
  1: list<TKeyValue> data;
}

struct TFCMMessage {
  1: TNotificationPayload notiPayload;
  2: TDataPayload dataPayload;
}

struct TResponse {
  1: i32 statusCode;
  2: string header;
  3: string body;
}

typedef string TDeviceToken
typedef list<TDeviceToken> TDeviceTokenList

service FCMService {
  bool addDeviceToken(1: string phone, 2: TDeviceToken deviceToken);

  bool addListDeviceToken(1: string phone, 2: TDeviceTokenList tokenList);

  TResponse notiToDeviceToken(1: TFCMMessage message, 2: TDeviceToken deviceToken);

  TResponse notiToPhone(1: TFCMMessage message, 2: string phone);

  TResponse notiToTopic(1: string topic, 3: string condition, 4: TFCMMessage message);
}