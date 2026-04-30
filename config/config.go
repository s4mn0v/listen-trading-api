package config

import "os"

const BaseUrl = "https://api.bitget.com"

func GetApiKey() string     { return os.Getenv("BITGET_API_KEY") }
func GetSecretKey() string  { return os.Getenv("BITGET_SECRET_KEY") }
func GetPassphrase() string { return os.Getenv("BITGET_PASSPHRASE") }
