package pkg

import (
	"api-file.io/intermal"
	"log"
	"net/http"
	"os"
	"time"
)

// ClientFileIO Клиент для работы с файлами с помощью
// сервиса - https://www.file.io/developers/
type ClientFileIO struct {
	client  http.Client
	keyAuth string
}

func NewClientFileIO(timeout time.Duration) *ClientFileIO {
	if timeout <= 0 {
		log.Fatal("timeout has to more than 0")
	}
	loggerFile, err := os.OpenFile("log-file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("can't create logger file", err)
	}

	return &ClientFileIO{
		client: http.Client{
			Transport: intermal.LoggingRoundTripper{
				Logger: loggerFile,
				Next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
		keyAuth: intermal.RandomString(10),
	}
}
