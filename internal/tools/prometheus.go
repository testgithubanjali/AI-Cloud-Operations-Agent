package tools

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	promapi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func QueryPrometheus(query string) (string, error) {

	_ = godotenv.Load()

	url := os.Getenv("PROMETHEUS_URL")

	client, err := promapi.NewClient(promapi.Config{
		Address: url,
	})

	if err != nil {
		return "", err
	}

	api := v1.NewAPI(client)

	result, warnings, err := api.Query(
		context.Background(),
		query,
		time.Now(),
	)

	if err != nil {
		return "", err
	}

	if len(warnings) > 0 {
		fmt.Println("Warnings:", warnings)
	}

	return result.String(), nil
}
