package notifier

import (
	"hhscaner/test/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSendNotificationToTelegram_Valid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	baseUrl := "https://mock.mock/bot/message"
	message := "test message"

	mockHttpClient := mock.NewMockHttpClient(ctrl)
	mockHttpClient.
		EXPECT().
		Get(gomock.Eq(baseUrl)).
		Return([]byte(`{"ok":true}`), nil)

	err := SendNotificationToTelegram(mockHttpClient, baseUrl, message)
	if err != nil {
		t.Errorf("SendNotificationToTelegram() error = %s, want nil", err)
		return
	}
}
