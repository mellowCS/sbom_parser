package internal

type Scope string

const (
	required Scope = "required"
	optional Scope = "optional"
	excluded Scope = "excluded"
)
