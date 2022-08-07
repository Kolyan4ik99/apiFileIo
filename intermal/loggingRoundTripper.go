package intermal

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type LoggingRoundTripper struct {
	Logger io.Writer
	Next   http.RoundTripper
}

func (l LoggingRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	_, err := fmt.Fprintf(l.Logger, "[%s] %s '%s' [%s]\n",
		time.Now().Format(time.RFC3339), request.Method, request.URL, printMap(request.Header))
	if err != nil {
		return nil, err
	}
	return l.Next.RoundTrip(request)
}
