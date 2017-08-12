package main

type network struct {
	Company  []string  `json:"company"`
	Href     string    `json:"href"`
	Name     string    `json:"name"`
	ID       string    `json:"id"`
	Location location  `json:"location"`
	Stations []station `json:"stations"`
}

type location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
}

type station struct {
	Name       string  `json:"name"`
	Timestamp  string  `json:"timestamp"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	FreeBikes  int     `json:"free_bikes"`
	EmptySlots int     `json:"empty_slots"`
	ID         string  `json:"id"`
}
