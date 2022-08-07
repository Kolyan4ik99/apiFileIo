
# api-file.io

ApiClient for upload and download file from file.io
https://www.file.io/


## Installation

Write commands in your go project

```bash
  go get -u github.com/Kolyan4ik99/api-file.io
  go mod download
```

## Example

	client := pkg.NewClientFileIO(time.Second * 5)
	cl, err := client.UploadFile("log-file") // Загружаем файл с названием log-file в удалёный сервер
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cl.Info()) // Через ссылку в поле Link можно скачать файл через браузер

	file, err := os.Create("tmp") // Создаём файл в который сохраним информацию
	if err != nil {
		log.Fatal(err)
	}
	
	err = client.DownloadFile(file, cl.Key) // Сохраняем файл из удалёного сервера
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.Create("tmp2") // Создаём новый файл для повторного сохранения
	if err != nil {
		log.Fatal(err)
	}

	err = client.DownloadFile(file, cl.Key) // Не получится сохранить файл, т.к. в бесплатном плане файл живёт ОДНО скачивание
	if err != nil {
		log.Fatal(err)
	}
    