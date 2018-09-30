package app

import (
	"fmt"
	"github.com/tizz98/godash/models"
	"regexp"
	"strings"
)

func (a *App) ValidateDashboard(input *models.Dashboard) (*models.Dashboard, error) {
	dashboard := &models.Dashboard{}
	dashboard.Id = dashboard.GenerateId()

	if input.Background != nil {
		if background, err := validateColor(input.Background); err == nil {
			dashboard.Background = background
		} else {
			return nil, err
		}
	} else {
		dashboard.Background = a.String("000dff")
	}

	if input.Foreground != nil {
		if foreground, err := validateColor(input.Foreground); err == nil {
			dashboard.Foreground = foreground
		} else {
			return nil, err
		}
	} else {
		dashboard.Foreground = a.String("ffffff")
	}

	if input.TemperatureUnit != nil {
		if unit, err := validateTemperatureUnit(input.TemperatureUnit); err == nil {
			dashboard.TemperatureUnit = unit
		} else {
			return nil, err
		}
	} else {
		dashboard.TemperatureUnit = a.String("F")
	}

	if input.TimeUnit != nil {
		if unit, err := validateTimeUnit(input.TimeUnit); err == nil {
			dashboard.TimeUnit = unit
		} else {
			return nil, err
		}
	} else {
		dashboard.TimeUnit = a.String("12")
	}

	dashboard.Location = input.Location
	return dashboard, nil
}

var validHexColor = regexp.MustCompile("[A-Fa-f0-9]{6}")

func validateColor(color *string) (*string, error) {
	if color == nil {
		return nil, nil
	}
	trimmedColor := strings.TrimLeft(*color, "#")

	if len(trimmedColor) != 6 {
		return nil, fmt.Errorf("color must be 6 characters long")
	}

	if !validHexColor.MatchString(trimmedColor) {
		return nil, fmt.Errorf("invalid hex code")
	}

	lowerColor := strings.ToLower(trimmedColor)
	return &lowerColor, nil
}

func validateTemperatureUnit(unit *string) (*string, error) {
	if unit == nil {
		return nil, nil
	}

	lowerUnit := strings.ToLower(*unit)

	switch lowerUnit {
	case "f", "c":
	default:
		return nil, fmt.Errorf("temperature unit must be either 'f' or 'c'")
	}

	return &lowerUnit, nil
}

func validateTimeUnit(unit *string) (*string, error) {
	if unit == nil {
		return nil, nil
	}

	switch *unit {
	case "12", "24":
	default:
		return nil, fmt.Errorf("time unit must be either '12' or '24'")
	}

	return unit, nil
}
