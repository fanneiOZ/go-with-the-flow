package response

import (
	"shareddomain/entity"
)

type Metadata struct {
	Id      string      `json:"id"`
	Type    entity.Type `json:"entity_type"`
	Version uint        `json:"version"`
}

func CreateMetadata(object entity.Entity) Metadata {
	return Metadata{Id: object.Id(), Type: object.Type(), Version: object.EntityVersion()}
}
