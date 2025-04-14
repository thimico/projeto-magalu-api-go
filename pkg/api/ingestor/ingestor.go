package ingestor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projeto-magalu-api-go/pkg/utl/config"
	"time"

	"projeto-magalu-api-go/pkg/api/service"
)

var telegramAPIURLFormat = "https://api.telegram.org/bot%s/sendMessage"

func StartPulseIngestor(pulseService service.Pulse, cfg *config.Configuration) {
	ticker := time.NewTicker(1 * time.Hour)
	monthlyTicker := time.NewTicker(30 * 24 * time.Hour)

	go func() {
		for {
			select {
			case <-ticker.C:
				ProcessPulses(pulseService)
			case <-monthlyTicker.C:
				SendMonthlyConsumptionToContractAndCatalog(pulseService, cfg)
			}
		}
	}()
}

func ProcessPulses(pulseService service.Pulse) {
	ctx := context.Background()
	err := pulseService.AggregateAndStore(ctx)
	if err != nil {
		log.Printf("Error processing pulses: %v", err)
	}
}

func SendMonthlyConsumptionToContractAndCatalog(pulseService service.Pulse, cfg *config.Configuration) {
	ctx := context.Background()
	tenant := "tenant_xpto"

	consumption, err := pulseService.GetAllResourcesConsumption(ctx, tenant)
	if err != nil {
		log.Printf("Error getting consumption: %v", err)
		return
	}

	message := fmt.Sprintf("📊 Consumo mensal do tenant [%s]:\n", tenant)
	for sku, amount := range consumption {
		message += fmt.Sprintf("- Produto: %s | Total: %.2f\n", sku, amount)
	}

	err = SendToTelegram(cfg.TelegramBotToken.TelegramBotToken, cfg.TelegramChatID.TelegramChatID, message)
	if err != nil {
		log.Printf("Error sending message to Telegram: %v", err)
	}
}

func OverrideTelegramAPIURL(format string) {
	telegramAPIURLFormat = format
}

func SendToTelegram(botToken, chatID, message string) error {
	url := fmt.Sprintf(telegramAPIURLFormat, botToken)

	payload := map[string]string{
		"chat_id": chatID,
		"text":    message,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status: %s", resp.Status)
	}

	return nil
}
