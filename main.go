package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	var resp Response
	err := godotenv.Load("file.env")
	if err != nil {
		ErrorsToDocker(err, "Error loading .env file:")
		return
	}
	postURL := "https://api.hamsterkombatgame.io/clicker/tap"
	authkey := os.Getenv("AUTH_KEY")
	unixTime := time.Now().Unix()
	maxTaps := MaxTaps(resp.ClickerUser.Level)

	//смысла в проверке нет
	if maxTaps == 0 {
		maxTaps = 1000
	}
	payload := Payload{
		Count:         int64(maxTaps),
		AvailableTaps: 0,
		Timestamp:     unixTime,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		ErrorsToDocker(err, "Error marshalling JSON:")
		return
	}

	request, err := http.NewRequest("POST", postURL, bytes.NewBuffer(body))
	if err != nil {
		ErrorsToDocker(err, "Failed to send POST")
	}
	request.Header.Add("Authorization", authkey)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		ErrorsToDocker(err, "Error reading response body: ")
		return
	}

	//мб надо было сохранить просто в ClickerUser?
	err = json.Unmarshal(body, &resp)
	if err != nil {
		ErrorsToDocker(err, "Error unmarshalling JSON: ")
		return
	}

	switch response.StatusCode {
	case 200:
		timer := resp.ClickerUser.Level * 340
		time.Sleep(time.Duration(timer) * time.Second)
		main()
	default:
		fmt.Println(response.StatusCode)
		time.Sleep(10 * time.Second)
		main()

	}

}

//использовалось для просмотра json, который дает сервис
/*
func SaveDataToJSON(data []byte) error {
	outPath := "/responseData"
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		description := fmt.Sprintf(("Path file does not exist: %s"), outPath)
		ErrorsToDocker(err, description)
	}
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		if err := os.Mkdir(outPath, os.ModePerm); err != nil {
			description := fmt.Sprintf(("Failed to create output directory: %s"), err)
			ErrorsToDocker(err, description)
		}
	}
	now := time.Now()
	timestamp := now.Format("02012006150405")
	filename := fmt.Sprintf("output_%s.json", timestamp)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil

}*/

// формула расчета тапов/кликов
// 1 уровень - 1000 тапов
// 2 уровень - 1500
// 3 уровень - 2000
// i := 1000, i = 1000+(n-1)*500, так как первый уровень - нулевой 0
// если уровень игрока 3, то макс тапов по формуле 1000+(3-1)*500 = 2000.
func MaxTaps(level int) int {
	taps := 1000 + (level-1)*500
	return taps
}

// Возвращает ошибки в docker logs
func ErrorsToDocker(err error, description string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, description, err)
		return
	}
}
