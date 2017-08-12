package main

import (
	"fmt"
	"strconv"

	"github.com/OpenTransports/lib-go/models"
	"github.com/go-siris/siris/context"
)

// GetTransports - /api/transports?latitude=...&longitude=...&radius=...
// Send the transports aroud the passed position
// @formParam latitude : optional, the latitude around where to search, default is 0
// @formParam longitude : optional, the longitude around where to search, default is 0
// @formParam radius : optional, default is 200m
func getTransports(ctx context.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	// Get position in params
	// Parse them to floats
	// Ignore errors because it default to 0
	latitude, _ := strconv.ParseFloat(ctx.FormValue("latitude"), 64)
	longitude, _ := strconv.ParseFloat(ctx.FormValue("longitude"), 64)
	radius, _ := strconv.ParseInt(ctx.FormValue("radius"), 10, 32)
	// Create a Position object
	position := models.Position{
		Latitude:  latitude,
		Longitude: longitude,
	}
	// Set the radius to its default value if none is passed
	if radius == 0 {
		radius = 200
	}
	nearest, err := nearestBikeStand(position, int(radius))
	if err != nil {
		ctx.Application().Log("Error getting nearest transports in /api/transports\n	==> %v", err)
	}
	// Return the result
	_, err = ctx.JSON(nearest)
	// Log the error if any
	if err != nil {
		ctx.Application().Log("Error writting answer in /api/transports\n	==> %v", err)
	}
}

func nearestBikeStand(position models.Position, radius int) ([]models.Transport, error) {
	// Create transports array
	nearestTransports := []models.Transport{}
	// Get all transports near the passed position
	// Only aske agencies that cover the passed position
	for _, a := range agenciesContaining(position) {
		transports, err := fetchTransports(a)
		if err != nil {
			return nil, fmt.Errorf("Error fetching transports for %v\n	==> %v", a, err)
		}
		nearestTransports = append(
			nearestTransports,
			a.TransportsNearPosition(transports, position, radius)...,
		)
	}
	return nearestTransports, nil
}
