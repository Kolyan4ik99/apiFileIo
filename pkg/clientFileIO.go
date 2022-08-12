package pkg

import (
	"github.com/Kolyan4ik99/apiFileIo/intermal"
	"io"
	"log"
	"net/http"
	"time"
)

// ClientFileIO Клиент для работы с файлами с помощью
// сервиса - https://www.file.io/developers/
type ClientFileIO struct {
	client  http.Client
	keyAuth string
}

func NewClientFileIO(logDirection io.Writer, timeout time.Duration) *ClientFileIO {
	if timeout <= 0 {
		log.Fatal("timeout has to more than 0")
	}

	return &ClientFileIO{
		client: http.Client{
			Transport: intermal.LoggingRoundTripper{
				Logger: logDirection,
				Next:   http.DefaultTransport,
			},
			Timeout: timeout,
		},
		keyAuth: intermal.RandomString(10),
	}
}
