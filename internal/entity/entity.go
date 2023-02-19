package entity

import "github.com/google/uuid"

type Exgauster struct {
	Guid uuid.UUID `json:"guid"`
	Name string    `json:"name"`
	Work float64   `json:"work"`
}

type Bearing struct {
	Guid          uuid.UUID `json:"guid"`
	Number        int       `json:"number"`
	ExgausterGuid uuid.UUID `json:"guid_exgauster"`
}

type Temperature struct {
	Guid        uuid.UUID `json:"guid"`
	Temperature float64   `json:"temperature"`
	BearingGuid uuid.UUID `json:"guid_bearing"`
}

type Vibration struct {
	Guid        uuid.UUID `json:"guid"`
	Type        string    `json:"type"`
	BearingGuid uuid.UUID `json:"guid_bearing"`
}

type AlarmAndWarning struct {
	Guid       uuid.UUID `json:"guid"`
	AlarmMax   float64   `json:"alarm_max"`
	AlarmMin   float64   `json:"alarm_min"`
	WarninMax  float64   `json:"warnin_max"`
	WarningMin float64   `json:"warning_min"`
	MainGuid   uuid.UUID `json:"main_guid"`
}

type Cooler struct {
	Guid          uuid.UUID `json:"guid"`
	Type          string    `json:"type"`
	ExgausterGuid uuid.UUID `json:"exgauster_guid"`
}

type CoolerTemperature struct {
	Guid              uuid.UUID `json:"guid"`
	TemperatureBefore float64   `json:"temperature_before"`
	TemperatureAfter  float64   `json:"temperature_after"`
	CoolerGuid        uuid.UUID `json:"guid_cooler"`
}

type GasMainfold struct {
	Guid                uuid.UUID `json:"guid"`
	TemperatureBefore   float64   `json:"temperature_before"`
	UnderpressureBefore float64   `json:"underpressure_before"`
	ExgausterGuid       uuid.UUID `json:"exgauster_guid"`
}

type GateValve struct {
	Guid          uuid.UUID `json:"guid"`
	Closed        float64   `json:"closed"`
	Open          float64   `json:"open"`
	Position      float64   `json:"position"`
	ExgausterGuid uuid.UUID
}

type MainDrive struct {
	Guid          uuid.UUID `json:"guid"`
	RotorCurrent  float64   `json:"rotor_current"`
	RotorVoltage  float64   `json:"rotor_voltage"`
	StatorCurrent float64   `json:"stator_current"`
	StatorVoltage float64   `json:"stator_voltage"`
	ExgusterGuid  uuid.UUID `json:"exguster_guid"`
}

type OilSystem struct {
	Guid          uuid.UUID `json:"guid"`
	OilLevel      float64   `json:"oil_level"`
	OilPressure   float64   `json:"oil_pressure"`
	ExgausterGuid uuid.UUID `json:"exgauster_guid"`
}
