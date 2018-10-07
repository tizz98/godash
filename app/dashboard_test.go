package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tizz98/godash/models"
	"testing"
)

func TestApp_ValidateDashboard(t *testing.T) {
	assert := assert.New(t)
	app := NewApp()
	ctx := app.Context

	dash := &models.Dashboard{
		Background:      "#fff000",
		Foreground:      "#ffffff",
		TemperatureUnit: "C",
		TimeUnit:        "12",
		Location:        "123 Mary Ln",
	}

	_, err := ctx.ValidateDashboard(dash)
	assert.NoError(err)

	dash = &models.Dashboard{
		Background:      "#fqf000",
		Foreground:      "#ffffff",
		TemperatureUnit: "C",
		TimeUnit:        "12",
		Location:        "123 Mary Ln",
	}
	_, err = ctx.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("invalid hex code"))
	}

	dash = &models.Dashboard{
		Background:      "#fff00",
		Foreground:      "#ffffff",
		TemperatureUnit: "C",
		TimeUnit:        "12",
		Location:        "123 Mary Ln",
	}
	_, err = ctx.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("color must be 6 characters long"))
	}

	dash = &models.Dashboard{
		Background:      "#fff000",
		Foreground:      "#ffffff",
		TemperatureUnit: "q",
		TimeUnit:        "12",
		Location:        "123 Mary Ln",
	}
	_, err = ctx.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("temperature unit must be either 'f' or 'c'"))
	}

	dash = &models.Dashboard{
		Background:      "#fff000",
		Foreground:      "#ffffff",
		TemperatureUnit: "F",
		TimeUnit:        "13",
		Location:        "123 Mary Ln",
	}
	_, err = ctx.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("time unit must be either '12' or '24'"))
	}

	dash, err = ctx.ValidateDashboard(&models.Dashboard{})
	assert.NoError(err)
	assert.NotEmpty(dash.Id)
	assert.NotEmpty(dash.Background)
	assert.NotEmpty(dash.Foreground)
	assert.NotEmpty(dash.TemperatureUnit)
	assert.NotEmpty(dash.TimeUnit)
	assert.NotEmpty(dash.Location)
}
