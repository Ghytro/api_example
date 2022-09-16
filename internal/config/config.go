package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	Addr,
	DBUrl string

	BodyLimit,
	MaxIncomingConnections int

	RequestHandleTimeout,
	WriteTimeout,
	IdleTimeout,
	ReadTimeout time.Duration
)

func parseInt(val string) int {
	result, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getConfigKey(key string) string {
	return os.Getenv("TESLALABZ_" + key)
}

func getIntConfigKey(key string) int {
	return parseInt(getConfigKey(key))
}

func init() {
	Addr = getConfigKey("ADDR")
	DBUrl = getConfigKey("DB_URL")
	MaxIncomingConnections = getIntConfigKey("MAX_INCOMING_CONNS")
	BodyLimit = getIntConfigKey("BODY_LIMIT")
	RequestHandleTimeout = time.Duration(
		getIntConfigKey("REQUEST_HANDLE_TIMEOUT"),
	) * time.Second
	ReadTimeout = time.Duration(
		getIntConfigKey("READ_TIMEOUT"),
	) * time.Second
	WriteTimeout = time.Duration(
		getIntConfigKey("WRITE_TIMEOUT"),
	) * time.Second
	IdleTimeout = time.Duration(
		getIntConfigKey("IDLE_TIMEOUT"),
	) * time.Second
}
