package app

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/tizz98/godash/models"
)

func (ctx *Context) ValidateDashboard(input *models.Dashboard) (*models.Dashboard, error) {
	dashboard := &models.Dashboard{}
	dashboard.Id = models.GenerateId()

	if input.Background != "" {
		if background, err := validateColor(input.Background); err == nil {
			dashboard.Background = background
		} else {
			return nil, err
		}
	} else {
		dashboard.Background = "000dff"
	}

	if input.Foreground != "" {
		if foreground, err := validateColor(input.Foreground); err == nil {
			dashboard.Foreground = foreground
		} else {
			return nil, err
		}
	} else {
		dashboard.Foreground = "ffffff"
	}

	if input.TemperatureUnit != "" {
		if unit, err := validateTemperatureUnit(input.TemperatureUnit); err == nil {
			dashboard.TemperatureUnit = unit
		} else {
			return nil, err
		}
	} else {
		dashboard.TemperatureUnit = "F"
	}

	if input.TimeUnit != "" {
		if unit, err := validateTimeUnit(input.TimeUnit); err == nil {
			dashboard.TimeUnit = unit
		} else {
			return nil, err
		}
	} else {
		dashboard.TimeUnit = "12"
	}

	if input.Location != "" {
		dashboard.Location = input.Location
	} else {
		dashboard.Location = "San Francisco, CA"
	}

	return dashboard, nil
}

var validHexColor = regexp.MustCompile("[A-Fa-f0-9]{6}")

func validateColor(color string) (string, error) {
	if color == "" {
		return "", nil
	}
	trimmedColor := strings.TrimLeft(color, "#")

	if len(trimmedColor) != 6 {
		return "", fmt.Errorf("color must be 6 characters long")
	}

	if !validHexColor.MatchString(trimmedColor) {
		return "", fmt.Errorf("invalid hex code")
	}

	return strings.ToLower(trimmedColor), nil
}

func validateTemperatureUnit(unit string) (string, error) {
	if unit == "" {
		return "", nil
	}

	lowerUnit := strings.ToLower(unit)

	switch lowerUnit {
	case "f", "c":
	default:
		return "", fmt.Errorf("temperature unit must be either 'f' or 'c'")
	}

	return lowerUnit, nil
}

func validateTimeUnit(unit string) (string, error) {
	if unit == "" {
		return "", nil
	}

	switch unit {
	case "12", "24":
	default:
		return "", fmt.Errorf("time unit must be either '12' or '24'")
	}

	return unit, nil
}
