package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"new-filterization/internal/app/model/domain"
)

type RegistrationRepository struct {
	db *mongo.Database
}

func NewRegistrationRepository(db *mongo.Database) *RegistrationRepository {
	return &RegistrationRepository{db}
}

func (rr RegistrationRepository) GetByFirstChoiceLevel(ctx context.Context, optType string, schoolOption []*domain.PpdbOption, firstChoice primitive.ObjectID) []domain.PpdbRegistration {
	//TODO implement me

	var showInfoCursor *mongo.Cursor
	var err error
	registrationsCollection := rr.db.Collection("ppdb_registrations")
	matchStage := bson.D{{"$match", bson.M{
		"first_option_id": firstChoice,

		"$and": []bson.M{bson.M{
			"status": "fit",
		},
		},
	}}}

	if optType == "perpindahan" {

		sortByScore := bson.D{{"$sort", bson.D{{"priority", -1}}}}
		sortByDistance := bson.D{{"$sort", bson.D{{"distance1", 1}}}}
		sortByAge := bson.D{{"$sort", bson.D{{"birth_date", 1}}}}

		showInfoCursor, err = registrationsCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, sortByScore, sortByDistance, sortByAge,
		}, options.Aggregate().SetAllowDiskUse(true))

	} else if optType == "prestasi" || optType == "prestasi-rapor" || optType == "rapor" {

		sortByScore := bson.D{{"$sort", bson.D{{"score", -1}}}}
		sortByAge := bson.D{{"$sort", bson.D{{"birth_date", 1}}}}

		showInfoCursor, err = registrationsCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, sortByScore, sortByAge,
		}, options.Aggregate().SetAllowDiskUse(true))

	} else if optType == "rapor-unggulan" {

		sortByScore := bson.D{{"$sort", bson.D{{"score_a1", -1}}}}
		sortByAge := bson.D{{"$sort", bson.D{{"birth_date", 1}}}}

		showInfoCursor, err = registrationsCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, sortByScore, sortByAge,
		}, options.Aggregate().SetAllowDiskUse(true))

	} else {
		sortByScore := bson.D{{"$sort", bson.D{{"distance1", 1}}}}
		sortByAge := bson.D{{"$sort", bson.D{{"birth_date", 1}}}}

		showInfoCursor, err = registrationsCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, sortByScore, sortByAge,
		}, options.Aggregate().SetAllowDiskUse(true))
	}

	var showsWithInfo []domain.PpdbRegistration

	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	defer showInfoCursor.Close(ctx)

	/*
		db.getCollection("ppdb_registrations").aggregate(
		   [
		     { $match :
		        {
		            first_choice_option : ObjectId('60b026fa2ddd2fcf61dc4346'),
		            registration_level : 'smk',
		            status : 'registered'
		        }
		     },
		     {
		     $addFields: {
		           totalKejuaraanUjikom: { $sum: ["$score_kejuaraan","$score_ujikom"] } ,
		         }
		     },
		     { $sort: { totalKejuaraanUjikom: -1 } }
		   ]
		)
	*/

	result := make([]domain.PpdbRegistration, 0)

	for _, row := range showsWithInfo {
		tmp := domain.PpdbRegistration{
			Id:                row.Id,
			Name:              row.Name,
			OptionType:        row.OptionType,
			RegistrationLevel: row.RegistrationLevel,

			FirstChoiceOption: row.FirstChoiceOption,
			FirstChoiceSchool: row.FirstChoiceSchool,
			FirstOption:       row.FirstOption,
			FirstSchool:       row.FirstSchool,

			SecondChoiceOption: row.SecondChoiceOption,
			SecondChoiceSchool: row.SecondChoiceSchool,
			SecondOption:       row.SecondOption,
			SecondSchool:       row.SecondSchool,

			ThirdChoiceOption: row.ThirdChoiceOption,
			ThirdChoiceSchool: row.ThirdChoiceSchool,
			ThirdOption:       row.ThirdOption,
			ThirdSchool:       row.ThirdSchool,

			Score:     row.Score,
			Priority:  row.Priority,
			Distance:  row.Distance1,
			Distance1: row.Distance1,
			Distance2: row.Distance2,
			Distance3: row.Distance3,

			ScoreA:               row.ScoreA1,
			ScoreA1:              row.ScoreA1,
			ScoreA2:              row.ScoreA2,
			ScoreA3:              row.ScoreA3,
			ScoreKejuaraanUjikom: row.ScoreKejuaraan + row.ScoreUjikom,
			ScoreKejuaraan:       row.ScoreKejuaraan,
			ScoreUjikom:          row.ScoreUjikom,

			BirthDate:      row.BirthDate,
			AcceptedStatus: 0,
			//AcceptedIndex:  0, //perlu di update idx berapa untuk firstchoice
			AcceptedIndex: domain.FindIndex(row.FirstChoiceOption, schoolOption),

			AcceptedChoiceId:     row.FirstChoiceOption,
			AcceptedSchoolId:     row.FirstChoiceSchool,
			AcceptedChoiceOption: row.FirstOption,
			AcceptedChoiceSchool: row.FirstSchool,
		}

		result = append(result, tmp)
	}
	return result
}
