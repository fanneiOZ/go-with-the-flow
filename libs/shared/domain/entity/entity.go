package entity

type Type string

type Entity interface {
	Id() string
	Type() Type
	EntityVersion() uint
}
