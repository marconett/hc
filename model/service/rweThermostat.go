package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
)

// RWEThermostat is service to represent a thermostat.
type RWEThermostat struct {
	*TemperatureSensor

	TargetTemp *characteristic.TemperatureCharacteristic
	Mode       *characteristic.HeatingCoolingMode
	TargetMode *characteristic.HeatingCoolingMode
	Humidity   *characteristic.Humidity

	targetTempChange func(float64)
}

// NewRWEThermostat returns a thermostat service.
func NewRWEThermostat(name string, temperature, min, max, steps, humidity float64) *RWEThermostat {

	svc := NewTemperatureSensor(name, temperature, min, max, steps)

	tempUnit := svc.Unit.Unit()
	targetTemp := characteristic.NewTargetTemperatureCharacteristic(temperature, min, max, steps, string(tempUnit))
	mode := characteristic.NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	targetMode := characteristic.NewTargetHeatingCoolingMode(model.HeatCoolModeOff)
	hum := characteristic.NewHumidity(humidity)

	svc.Type = TypeThermostat
	svc.AddCharacteristic(mode.Characteristic)
	svc.AddCharacteristic(targetMode.Characteristic)
	svc.AddCharacteristic(targetTemp.Characteristic)
	svc.AddCharacteristic(hum.Characteristic)

	t := RWEThermostat{svc, targetTemp, mode, targetMode, hum, nil}

	return &t
}
