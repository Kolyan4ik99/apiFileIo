package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// UploadFile [POST /] Загрузка файла с именем fileName,
// файл должен существовать в рабочей дериктории
func (c *ClientFileIO) UploadFile(fileName string) (*CreateFileInfo, error) {
	body, err := os.Open(fileName)
	if err != nil {
		log.Fatal("bad file")
	}
	values := map[string]io.Reader{
		"file": body,
	}

	// Нужно разобраться в этой части кода, использовал для загрузки файла
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}

	}
	w.Close()

	r, err := http.NewRequest("POST", "https://file.io/", &b)
	if err != nil {
		return nil, err
	}

	r.Header.Set("accept", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.keyAuth))
	r.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.client.Do(r)
	if err != nil {
		return nil, err
	}

	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fileInfo := CreateFileInfo{}
	err = json.Unmarshal(output, &fileInfo)
	if err != nil {
		return nil, err
	}
	return &fileInfo, nil
}
