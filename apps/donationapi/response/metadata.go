package response

import domain "shareddomain"

type Metadata struct {
	Id   string            `json:"id"`
	Type domain.EntityType `json:"entity_type"`
}

func CreateMetadata(object domain.Entity) Metadata {
	return Metadata{Id: object.Id(), Type: object.Type()}
}
