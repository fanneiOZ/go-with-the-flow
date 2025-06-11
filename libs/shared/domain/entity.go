package domain

type EntityType string

type Entity interface {
	Id() string
	Type() EntityType
}
