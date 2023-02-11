package domain

type PpdbChoiceSchool struct {
	Id            string `bson:"id,omitempty"`
	Type          string `bson:"type,omitempty"`
	Level         string `bson:"level,omitempty"`
	Name          string `bson:"name,omitempty"`
	Npsn          string `bson:"npsn,omitempty"`
	AddressCity   string `bson:"address_city,omitempty"`
	CoordinateLat string `bson:"coordinate_lat,omitempty"`
	CoordinateLng string `bson:"coordinate_lng,omitempty"`
}
