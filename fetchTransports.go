package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OpenTransports/lib-go/models"
)

func fetchTransports(agency models.Agency) ([]models.Transport, error) {
	response, err := http.Get(citybikesServerURL + "/v2/networks/" + agency.ID)

	if err != nil {
		return nil, fmt.Errorf("Error fetching %v\n	==> %v", agency, err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading body from %v\n	==> %v", agency, err)
	}

	content := struct {
		Network network `json:"network"`
	}{}
	err = json.Unmarshal(body, &content)
	if err != nil {
		return nil, fmt.Errorf("Error parsing body from %v\n	==> %v", agency, err)
	}

	transports := make([]models.Transport, len(content.Network.Stations))

	for i, station := range content.Network.Stations {
		transports[i] = models.Transport{
			ID:       station.ID,
			Name:     station.Name,
			AgencyID: agency.ID,
			Line:     agency.ID,
			Type:     models.Bike,
			Informations: []models.Information{
				models.Information{
					Content: []string{fmt.Sprintf("%v", station.FreeBikes)},
				},
				models.Information{
					Content: []string{fmt.Sprintf("%v", station.EmptySlots)},
				},
			},
			Position: models.Position{
				Latitude:  station.Latitude,
				Longitude: station.Longitude,
			},
		}
	}

	return transports, nil
}
