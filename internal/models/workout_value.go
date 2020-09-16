package models

const (
	UnitMeters      = 1
	UnitSeconds     = 2
	UnitKilograms   = 3
	UnitReiteration = 4
	UnitWatt        = 5
	UnitTemp        = 6
	UnitPulse       = 7
)

type WorkoutValues []WorkoutValue

type WorkoutValue struct {
	ID      int     `json:"id"`
	Workout Workout `json:"workout"`
	Value   float64 `json:"value"`
	Unit    int8    `json:"unit"`
}
