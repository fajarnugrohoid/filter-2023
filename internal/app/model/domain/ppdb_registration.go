package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PpdbRegistration struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name,omitempty"`
	OptionType        string             `bson:"option_type,omitempty"`
	RegistrationLevel string             `bson:"registration_level,omitempty"`

	FirstChoiceOption primitive.ObjectID `bson:"first_option_id,omitempty"`
	FirstChoiceSchool primitive.ObjectID `bson:"first_school_id,omitempty"`
	FirstOption       PpdbChoiceOption   `json:"first_option" bson:"first_option"`
	FirstSchool       PpdbChoiceSchool   `json:"first_school" bson:"first_school,omitempty"`

	SecondChoiceOption primitive.ObjectID `bson:"second_option_id,omitempty"`
	SecondChoiceSchool primitive.ObjectID `bson:"second_school_id,omitempty"`
	SecondOption       PpdbChoiceOption   `json:"second_option" bson:"second_option,omitempty"`
	SecondSchool       PpdbChoiceSchool   `json:"second_school" bson:"second_school,omitempty"`

	ThirdChoiceOption primitive.ObjectID `bson:"third_option_id,omitempty"`
	ThirdChoiceSchool primitive.ObjectID `bson:"third_school_id,omitempty"`
	ThirdOption       PpdbChoiceOption   `json:"third_option" bson:"third_option,omitempty"`
	ThirdSchool       PpdbChoiceSchool   `json:"third_school" bson:"third_school,omitempty"`

	Score                float64 `bson:"score,omitempty"`
	Priority             int     `bson:"priority,omitempty"`
	Distance             float64 `bson:"distance"`
	Distance1            float64 `bson:"distance1,omitempty"`
	Distance2            float64 `bson:"distance2,omitempty"`
	Distance3            float64 `bson:"distance3,omitempty"`
	ScoreA               float64 `bson:"score_a"`
	ScoreA1              float64 `bson:"score_a1,omitempty"`
	ScoreA2              float64 `bson:"score_a2,omitempty"`
	ScoreA3              float64 `bson:"score_a3,omitempty"`
	ScoreKejuaraanUjikom float64 `bson:"total_kejuaraan_ujikom"`
	ScoreKejuaraan       float64 `bson:"score_kejuaraan,omitempty"`
	ScoreUjikom          float64 `bson:"score_ujikom,omitempty"`

	BirthDate primitive.DateTime `bson:"birth_date,omitempty"`

	AcceptedStatus         int                `bson:"selection"`
	AcceptedIndex          int                `bson:"accepted_index"`
	AcceptedChoicePosition int                `bson:"accepted_choice_position"`
	AcceptedChoiceId       primitive.ObjectID `bson:"accepted_choice_id,omitempty"`
	AcceptedSchoolId       primitive.ObjectID `bson:"accepted_school_id,omitempty"`
	AcceptedChoiceOption   PpdbChoiceOption   `bson:"accepted_choice_option,omitempty"`
	AcceptedChoiceSchool   PpdbChoiceSchool   `bson:"accepted_choice_school,omitempty"`
}
