package models

const (
	WorkoutTypeCycle           = 1
	WorkoutTypeCompositeParent = 2
	WorkoutTypeComposite       = 3
	WorkoutTypeGymnastic       = 4
	WorkoutTypeSingleCombat    = 5
	WorkoutTypeGame            = 6
)

type WorkoutTypes []WorkoutType

type WorkoutType struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Type     uint16 `json:"type"`
}
