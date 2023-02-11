package domain

type PpdbChoiceOption struct {
	Id           string `bson:"id,omitempty"`
	Type         string `bson:"type,omitempty"`
	Name         string `bson:"name,omitempty"`
	MajorId      int    `bson:"major_id,omitempty"`
	NoColorBlind bool   `bson:"no_color_blind,omitempty"`
}
