package notifier

import (
	"hhscaner/configuration"
	"hhscaner/test/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSendNotificationToTelegram(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := configuration.Config{
		Telegram: configuration.Telegram{
			ApiUrl:   "https://mock.mock/bot",
			BotToken: "12345:ABCDE",
			ChatId:   "12345",
		},
	}

	expectedUrl := config.Telegram.ApiUrl + config.Telegram.BotToken + "/sendMessage"
	expectedParams := []configuration.UrlParameter{
		{Key: "chat_id", Value: config.Telegram.ChatId},
		{Key: "text", Value: "test message"},
	}

	mockHttpClient := mock.NewMockHttpClient(ctrl)
	mockHttpClient.
		EXPECT().
		Get(
			gomock.Eq(expectedUrl),
			gomock.Eq(&expectedParams),
		).Return([]byte(`{"ok":true}`), nil)

	err := SendNotificationToTelegram(&config, mockHttpClient, "test message")
	if err != nil {
		t.Errorf("SendNotificationToTelegram() error = %s, want nil", err)
		return
	}
}
