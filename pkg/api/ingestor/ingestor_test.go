package ingestor_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"projeto-magalu-api-go/pkg/utl/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"projeto-magalu-api-go/pkg/api/ingestor"
	"projeto-magalu-api-go/pkg/api/ingestor/mocks"
)

var (
	mockPulseService *mocks.Pulse
	cfg              *config.Configuration
	mockTelegramSrv  *httptest.Server
	telegramMessage  string
)

func setup() {
	os.Setenv("TELEGRAM_BOT_TOKEN", "test-token")
	os.Setenv("TELEGRAM_CHAT_ID", "test-chat")

	cfg = &config.Configuration{
		TelegramBotToken: &config.Telegram{TelegramBotToken: "test-token"},
		TelegramChatID:   &config.Telegram{TelegramChatID: "test-chat"},
	}

	mockPulseService = new(mocks.Pulse)

	mockTelegramSrv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]string
		_ = json.NewDecoder(r.Body).Decode(&payload)
		telegramMessage = payload["text"]
		w.WriteHeader(http.StatusOK)
	}))

	ingestorOverrideTelegramURL(mockTelegramSrv.URL + "/bot%s/sendMessage")
}

func teardown() {
	mockTelegramSrv.Close()
	telegramMessage = ""
	os.Unsetenv("TELEGRAM_BOT_TOKEN")
	os.Unsetenv("TELEGRAM_CHAT_ID")
}

func TestProcessPulses_Success(t *testing.T) {
	setup()
	defer teardown()

	mockPulseService.On("AggregateAndStore", mock.Anything).Return(nil)
	ingestor.ProcessPulses(mockPulseService)
	mockPulseService.AssertExpectations(t)
}

func TestProcessPulses_Error(t *testing.T) {
	setup()
	defer teardown()

	mockPulseService.On("AggregateAndStore", mock.Anything).Return(errors.New("some error"))
	ingestor.ProcessPulses(mockPulseService)
	mockPulseService.AssertExpectations(t)
}

func TestSendMonthlyConsumptionToContractAndCatalog_Success(t *testing.T) {
	setup()
	defer teardown()

	mockPulseService.On("GetAllResourcesConsumption", mock.Anything, "tenant_xpto").
		Return(map[string]float64{
			"SKU-001": 123.45,
			"SKU-002": 67.89,
		}, nil)

	ingestor.SendMonthlyConsumptionToContractAndCatalog(mockPulseService, cfg)
	mockPulseService.AssertExpectations(t)
	assert.Contains(t, telegramMessage, "SKU-001")
	assert.Contains(t, telegramMessage, "123.45")
}

func TestSendMonthlyConsumptionToContractAndCatalog_Error(t *testing.T) {
	setup()
	defer teardown()

	mockPulseService.On("GetAllResourcesConsumption", mock.Anything, "tenant_xpto").
		Return(nil, errors.New("fail"))

	ingestor.SendMonthlyConsumptionToContractAndCatalog(mockPulseService, cfg)
	mockPulseService.AssertExpectations(t)
	assert.Empty(t, telegramMessage)
}

func TestSendToTelegram_ErrorMarshal(t *testing.T) {
	err := ingestor.SendToTelegram("token", "chat", string([]byte{0xff}))
	assert.Error(t, err)
}

func TestSendToTelegram_ErrorRequest(t *testing.T) {
	err := ingestor.SendToTelegram("token", "chat", "msg")
	assert.Error(t, err)
}

func TestSendToTelegram_InvalidStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "fail", http.StatusUnauthorized)
	}))
	defer srv.Close()

	ingestorOverrideTelegramURL(srv.URL + "/bot%s/sendMessage")
	err := ingestor.SendToTelegram("token", "chat", "msg")
	assert.Error(t, err)
}

func TestSendToTelegram_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	ingestorOverrideTelegramURL(srv.URL + "/bot%s/sendMessage")
	err := ingestor.SendToTelegram("token", "chat", "ok")
	assert.NoError(t, err)
}

func ingestorOverrideTelegramURL(urlFormat string) {
	ingestor.OverrideTelegramAPIURL(urlFormat)
}
