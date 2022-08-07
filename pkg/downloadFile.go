package pkg

import (
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"time"
)

type CreateFileInfo struct {
	Success      bool      `json:"success,omitempty"`
	Status       int16     `json:"status,omitempty"`
	Id           string    `json:"id,omitempty"`
	Key          string    `json:"key,omitempty"`
	Link         string    `json:"link,omitempty"`
	Private      bool      `json:"private,omitempty"`
	Downloads    int       `json:"downloads,omitempty"`
	MaxDownloads int       `json:"maxDownloads,omitempty"`
	Size         int64     `json:"size,omitempty"`
	ExpiryDate   time.Time `json:"expiry"`
	CreatedDate  time.Time `json:"created"`
	ModifiedDate time.Time `json:"modified"`
}

// DownloadFile Загрузка данных с ключем key, в file
func (c *ClientFileIO) DownloadFile(file io.Writer, key string) error {

	url := "https://file.io/" + key
	urls, err := urlpkg.Parse(url)
	if err != nil {
		return err
	}

	// Использовал именно такой способ, т.к. методом http.Get()
	// не получается установить HEADERS
	r := &http.Request{
		URL:        urls,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       urls.Host,
	}

	r.Method = "GET"
	r.Header = make(http.Header)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.keyAuth))
	r.Header.Add("accept", "application/json")

	resp, err := c.client.Do(r)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
