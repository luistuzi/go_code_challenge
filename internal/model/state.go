package model

type State int

const (
	Unknown OrderState = iota
	Available
	In_Use
	Inactive
)
