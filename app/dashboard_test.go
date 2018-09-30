package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tizz98/godash/models"
	"testing"
)

func TestApp_ValidateDashboard(t *testing.T) {
	assert := assert.New(t)
	app := New()

	dash := &models.Dashboard{
		Background:      app.String("#fff000"),
		Foreground:      app.String("#ffffff"),
		TemperatureUnit: app.String("C"),
		TimeUnit:        app.String("12"),
		Location:        app.String("123 Mary Ln"),
	}

	_, err := app.ValidateDashboard(dash)
	assert.NoError(err)

	dash = &models.Dashboard{
		Background:      app.String("#fqf000"),
		Foreground:      app.String("#ffffff"),
		TemperatureUnit: app.String("C"),
		TimeUnit:        app.String("12"),
		Location:        app.String("123 Mary Ln"),
	}
	_, err = app.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("invalid hex code"))
	}

	dash = &models.Dashboard{
		Background:      app.String("#fff00"),
		Foreground:      app.String("#ffffff"),
		TemperatureUnit: app.String("C"),
		TimeUnit:        app.String("12"),
		Location:        app.String("123 Mary Ln"),
	}
	_, err = app.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("color must be 6 characters long"))
	}

	dash = &models.Dashboard{
		Background:      app.String("#fff000"),
		Foreground:      app.String("#ffffff"),
		TemperatureUnit: app.String("q"),
		TimeUnit:        app.String("12"),
		Location:        app.String("123 Mary Ln"),
	}
	_, err = app.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("temperature unit must be either 'f' or 'c'"))
	}

	dash = &models.Dashboard{
		Background:      app.String("#fff000"),
		Foreground:      app.String("#ffffff"),
		TemperatureUnit: app.String("F"),
		TimeUnit:        app.String("13"),
		Location:        app.String("123 Mary Ln"),
	}
	_, err = app.ValidateDashboard(dash)
	if assert.Error(err) {
		assert.Equal(err, fmt.Errorf("time unit must be either '12' or '24'"))
	}

	dash, err = app.ValidateDashboard(&models.Dashboard{})
	assert.NoError(err)
	assert.NotNil(dash.Id)
	assert.NotNil(dash.Background)
	assert.NotNil(dash.Foreground)
	assert.NotNil(dash.TemperatureUnit)
	assert.NotNil(dash.TimeUnit)
	assert.NotNil(dash.Location)
}
