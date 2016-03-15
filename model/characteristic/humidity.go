package characteristic

type Humidity struct {
	// *Characteristic
	*Float
	value float64
}

func NewHumidity(value float64) *Humidity {
	t := Humidity{NewFloat(value, PermsRead()), value}
	t.Type = TypeCurrentRelativeHumidity

	return &t
}

func (s *Humidity) SetHumidity(value float64) {
	s.SetFloat(value)
}

func (s *Humidity) Humidity() float64 {
	return s.FloatValue()
}
