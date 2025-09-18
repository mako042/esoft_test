package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Data struct {
	Time         string
	Url          string
	State        bool
	Status       string
	Descriptions string
}

type StatusHistory struct {
	LastState bool
	LastTime  time.Time
}

func main() {
	history := &StatusHistory{
		LastState: false,
		LastTime:  time.Now(),
	}

	d := &Data{}
	url := "https://google.com/"

	filename := fmt.Sprintf("logs_%s.json", time.Now().Format("02_Jan_2006_15-04-05"))
	file, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	for {
		req, err := http.Get(url)
		if err != nil {
			log.Printf("Error making request: %v", err)
			d.State = false
			d.Status = "ERROR"
			d.Descriptions = err.Error()
		} else {
			d.Status, d.Descriptions = d.fetch(req)
			req.Body.Close()
		}

		d.Time = time.Now().Format("02_Jan_2006_15-04-05")
		d.Url = url

		jsonData, err := json.Marshal(d)
		if err != nil {
			log.Printf("Error marshaling JSON: %v", err)
			continue
		}

		jsonData = append(jsonData, '\n')
		if _, err := file.Write(jsonData); err != nil {
			log.Printf("Error writing to log file: %v", err)
		}

		if history.LastState != d.State {
			e := time.Since(history.LastTime)

			fmt.Printf("=== NOTIFICATION ===\n")
			fmt.Printf("Время: %s\n", d.Time)
			fmt.Printf("Информация: Изменение статуса сайта. Текущее состояние: %s %s\n", d.Status, d.Descriptions)
			fmt.Printf("Прошло времени с последнего изменения статуса: %v\n", formatDuration(e))
			fmt.Printf("URL: %s\n", d.Url)
			fmt.Printf("================================\n\n")

			history.LastState = d.State
			history.LastTime = time.Now()
		}
		time.Sleep(15 * time.Second)
	}
}

func (d *Data) fetch(req *http.Response) (string, string) {
	if req == nil || req.StatusCode >= 400 {
		d.State = false
		if req != nil {
			return req.Status, http.StatusText(req.StatusCode)
		}
		return "ERROR", "Request failed"
	}
	d.State = true
	return "UP", http.StatusText(req.StatusCode)
}

func formatDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

	if days > 0 {
		return fmt.Sprintf("%dд %dч %dм %dс", days, hours, minutes, seconds)
	}
	if hours > 0 {
		return fmt.Sprintf("%dч %dм %dс", hours, minutes, seconds)
	}
	if minutes > 0 {
		return fmt.Sprintf("%dм %dс", minutes, seconds)
	}
	return fmt.Sprintf("%dс", seconds)
}
