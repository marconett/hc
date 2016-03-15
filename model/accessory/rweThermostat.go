package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type RWEThermostat struct {
	*Accessory

	RWEThermostat *service.RWEThermostat
}

// NewRWEThermostat returns a RWEThermostat which implements model.RWEThermostat.
func NewRWEThermostat(info model.Info, temp, min, max, steps, humidity float64) *RWEThermostat {
	accessory := New(info, TypeThermostat)
	t := service.NewRWEThermostat(info.Name, temp, min, max, steps, humidity)

	accessory.AddService(t.Service)

	return &RWEThermostat{accessory, t}
}

func (t *RWEThermostat) Temperature() float64 {
	return t.RWEThermostat.Temp.Temperature()
}

func (t *RWEThermostat) SetTemperature(value float64) {
	t.RWEThermostat.Temp.SetTemperature(value)
}

func (t *RWEThermostat) Unit() model.TempUnit {
	return t.RWEThermostat.Unit.Unit()
}

func (t *RWEThermostat) SetTargetTemperature(value float64) {
	t.RWEThermostat.TargetTemp.SetTemperature(value)
}

func (t *RWEThermostat) TargetTemperature() float64 {
	return t.RWEThermostat.TargetTemp.Temperature()
}

func (t *RWEThermostat) SetMode(value model.HeatCoolModeType) {
	if value != model.HeatCoolModeAuto {
		t.RWEThermostat.Mode.SetHeatingCoolingMode(value)
	}
}

func (t *RWEThermostat) Mode() model.HeatCoolModeType {
	return t.RWEThermostat.Mode.HeatingCoolingMode()
}

func (t *RWEThermostat) SetTargetMode(value model.HeatCoolModeType) {
	t.RWEThermostat.TargetMode.SetHeatingCoolingMode(value)
}

func (t *RWEThermostat) TargetMode() model.HeatCoolModeType {
	return t.RWEThermostat.TargetMode.HeatingCoolingMode()
}

func (t *RWEThermostat) Humidity() float64 {
	return t.RWEThermostat.Humidity.Humidity()
}

func (t *RWEThermostat) SetHumidity(value float64) {
	t.RWEThermostat.Humidity.SetHumidity(value)
}

func (t *RWEThermostat) OnTargetTempChange(fn func(float64)) {
	t.RWEThermostat.TargetTemp.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}

func (t *RWEThermostat) OnTargetModeChange(fn func(model.HeatCoolModeType)) {
	t.RWEThermostat.TargetMode.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(model.HeatCoolModeType(new.(byte)))
	})
}
