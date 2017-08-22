package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/OpenTransports/lib-go/models"
)

func fetchAgencies() ([]models.Agency, error) {
	response, err := http.Get(citybikesServerURL + "/v2/networks")

	if err != nil {
		return nil, fmt.Errorf("Error fetching networks\n	==> %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("Error reading body\n	==> %v", err)
	}

	content := struct {
		Networks []network `json:"networks"`
	}{}

	_ = json.Unmarshal(body, &content)
	// if err != nil {
	// 	return nil, fmt.Errorf("Error parsing body\n	==> %v", err)
	// }

	agencies := make([]models.Agency, len(content.Networks))

	for i, net := range content.Networks {
		agencies[i] = models.Agency{
			Name:   net.Name,
			ID:     net.ID,
			Radius: 20000,
			URL:    "https://api.citybik.es/v2",
			Types: map[int]models.TypeInfo{
				models.Bike: models.TypeInfo{
					Name: net.Name,
					Icon: serverURL + "/medias/bicycle.png",
				},
			},
			Center: models.Position{
				Latitude:  net.Location.Latitude,
				Longitude: net.Location.Longitude,
			},
		}
	}

	return agencies, nil
}
