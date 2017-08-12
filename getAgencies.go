package main

import (
	"strconv"

	"github.com/OpenTransports/lib-go/models"
	"github.com/go-siris/siris/context"
)

// GetAgencies - /api/agencies?latitude=...&longitude=...
// Send the agencies aroud the passed position or all agencies if no position is passed
// @formParam latitude : optional, the latitude around where to search, default is 0
// @formParam longitude : optional, the longitude around where to search, default is 0
func getAgencies(ctx context.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	// Get position in params
	// Parse them to floats
	// Ignore errors because it default to 0
	latitude, _ := strconv.ParseFloat(ctx.FormValue("latitude"), 64)
	longitude, _ := strconv.ParseFloat(ctx.FormValue("longitude"), 64)
	// Return agencies that covers the position
	_, err := ctx.JSON(agenciesContaining(models.Position{
		Latitude:  latitude,
		Longitude: longitude,
	}))
	// Log the error if any
	if err != nil {
		ctx.Application().Log("Error writting answer in /api/agencies\n	==> %v", err)
	}
}

func agenciesContaining(position models.Position) []models.Agency {
	filteredNetworks := make([]models.Agency, 0)

	for _, a := range agencies {
		if a.ContainsPosition(position) {
			filteredNetworks = append(filteredNetworks, a)
		}
	}

	return filteredNetworks
}
