//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Benchmark struct {
	ID                    string  `sql:"primary_key"` // @UUID("v4")(DC2Type:guid)
	PlanningMaterialityID *string // (DC2Type:guid)
	RevisedMaterialityID  *string // (DC2Type:guid)
	BenchmarkLovID        *string
	OtherBenchmarkName    *string
	OtherBenchmarkFloor   *float64
	OtherBenchmarkCeiling *float64
	Quantum               *float64
	State                 *string
}
