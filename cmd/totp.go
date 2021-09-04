package main

import (
	"os"
	"time"

	"github.com/pquerna/otp/totp"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewDevelopment()
	logger := log.Sugar()
	logger.Info("Generating TOTP secret")

	if len(os.Args) <= 1 {
		logger.Fatalf("no totp secret provided")
	}
	code := os.Args[1]
	logger.Infof("Start:%s:end", code)
	totpCode, err := totp.GenerateCode(code, time.Now())
	if err != nil {
		logger.Fatalf("failed to generate code %v", err)
	}
	logger.Infof("Code %s", totpCode)
}
