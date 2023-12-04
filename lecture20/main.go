package main

import (
	"fmt"
	"sync"
	"time"
)

type City struct {
	Name       string
	Population int
}

type App struct {
	data   map[string]City
	mutex  sync.RWMutex
	ticker *time.Ticker
	quit   chan struct{}
}

func NewApp() *App {
	return &App{
		data:   make(map[string]City),
		ticker: time.NewTicker(time.Minute),
		quit:   make(chan struct{}),
	}
}

func (a *App) Start() {
	a.loadData()

	go func() {
		for {
			select {
			case <-a.ticker.C:
				a.updateData()
			case <-a.quit:
				return
			}
		}
	}()
}

func (a *App) Stop() {
	close(a.quit)
	a.ticker.Stop()
}

func (a *App) loadData() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	cities := map[string]City{
		"City1": {Name: "Almaty", Population: 1000000},
		"City2": {Name: "Astana", Population: 500000},
	}

	for key, value := range cities {
		a.data[key] = value
	}
}

func (a *App) updateData() {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for key, city := range a.data {
		city.Population += 100
		a.data[key] = city
	}
}

func (a *App) getData() map[string]City {
	a.mutex.RLock()
	defer a.mutex.RUnlock()

	result := make(map[string]City)
	for key, value := range a.data {
		result[key] = value
	}

	return result
}

func main() {
	app := NewApp()
	app.Start()
	defer app.Stop()

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 10)
		data := app.getData()
		fmt.Println("Current Data:", data)
	}
}
