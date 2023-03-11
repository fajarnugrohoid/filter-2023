package repository

import (
	"context"
	"new-filterization/internal/app/model/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VocationalSchoolOptionRepository struct {
	db *mongo.Database
}

func NewVocationalSchoolOptionRepository(db *mongo.Database) *VocationalSchoolOptionRepository {
	return &VocationalSchoolOptionRepository{db}
}

func (rr VocationalSchoolOptionRepository) GetVocationalSchoolOptionByOptType(ctx context.Context, level string, optType string) []*domain.PpdbOption {
	//TODO implement me

	registrationsCollection := rr.db.Collection("ppdb_options")
	/*
		objectId1, err := primitive.ObjectIDFromHex("62861fc0563cbe29e2af33ed")
		objectId2, err := primitive.ObjectIDFromHex("62861ef6563cbe29e2af2997")
		objectId3, err := primitive.ObjectIDFromHex("62861eef563cbe29e2af28d1")
		objectId4, err := primitive.ObjectIDFromHex("62861fb9563cbe29e2af3345")
		objectId5, err := primitive.ObjectIDFromHex("62861fb8563cbe29e2af3336")
		objectId6, err := primitive.ObjectIDFromHex("62861fc1563cbe29e2af3411")
		objectId7, err := primitive.ObjectIDFromHex("62861fba563cbe29e2af3363")
		objectId8, err := primitive.ObjectIDFromHex("62861fbe563cbe29e2af33ba")
		objectId9, err := primitive.ObjectIDFromHex("62861fb5563cbe29e2af32e2")
		objectId10, err := primitive.ObjectIDFromHex("62861fb3563cbe29e2af32ac")
		objectId11, err := primitive.ObjectIDFromHex("62861fb2563cbe29e2af32a6")
		objectId12, err := primitive.ObjectIDFromHex("62861fb8563cbe29e2af3333")
		objectId13, err := primitive.ObjectIDFromHex("62861fb9563cbe29e2af333c")
		objectId14, err := primitive.ObjectIDFromHex("62861fb1563cbe29e2af327f")
		objectId15, err := primitive.ObjectIDFromHex("62861fbc563cbe29e2af3384")
		objectId16, err := primitive.ObjectIDFromHex("62861fba563cbe29e2af334e")
		objectId17, err := primitive.ObjectIDFromHex("62861fc0563cbe29e2af33f6")
		objectId18, err := primitive.ObjectIDFromHex("62861fb3563cbe29e2af32be")
		objectId19, err := primitive.ObjectIDFromHex("62861f02563cbe29e2af2a1e")
		objectId20, err := primitive.ObjectIDFromHex("62861fbc563cbe29e2af3384")
		objectId21, err := primitive.ObjectIDFromHex("62861fb9563cbe29e2af333c")
		objectId22, err := primitive.ObjectIDFromHex("62861eef563cbe29e2af28d1")
		objectId23, err := primitive.ObjectIDFromHex("62861fb3563cbe29e2af32ac")
		objectId24, err := primitive.ObjectIDFromHex("62861ef2563cbe29e2af2922")
		objectId25, err := primitive.ObjectIDFromHex("62861fb8563cbe29e2af3336")
		if err != nil {
			log.Println("Invalid id")
		}

		var schoolIds = [25]primitive.ObjectID{
			objectId1, objectId2, objectId3, objectId4, objectId5, objectId6, objectId7,
			objectId8, objectId9, objectId10, objectId11, objectId12, objectId13, objectId14,
			objectId15, objectId16, objectId17, objectId18, objectId19,
			objectId20, objectId21,
			objectId22, objectId23, objectId24, objectId25,
		} */

	matchStage := bson.D{{"$match", bson.M{
		"type": optType,
		/*
			"$and": []bson.M{bson.M{
				"school_id": bson.M{"$in": schoolIds},
			}}, */
		/*
			"$and": []bson.M{bson.M{
				"name": primitive.Regex{Pattern: "TANJUNGSARI", Options: ""},
			}}, */
	}}}
	//groupStage := bson.M{{"$group", bson.M{{"_id", "$podcast"}, {"total", bson.M{{"$sum", "$duration"}}}}}}

	pipeline := []bson.M{
		bson.M{"$match": bson.M{
			"$expr": bson.M{
				"$and": []bson.M{
					{"$eq": []string{"$_id", "$$school_id"}},
					{"$eq": []string{"$level", level}},
					//	{"$eq": []string{"$type", "negeri"}},
				},
			},
		}},
	}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "ppdb_schools"},
		{"let", bson.D{{"school_id", "$school_id"}}},
		{"pipeline", pipeline},
		{"as", "ppdb_schools"}}}}
	unwindStage := bson.D{{"$unwind", "$ppdb_schools"}}
	sortByName := bson.D{{"$sort", bson.D{{"name", 1}}}}
	sortByType := bson.D{{"$sort", bson.D{{"type", 1}}}}
	//allowDisk := bson.D{{"allow", true}}
	//fields := bson.D{{"$project", bson.D{{"name", 1}}}}

	showInfoCursor, err := registrationsCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, unwindStage, sortByName, sortByType,
	}, options.Aggregate().SetAllowDiskUse(true))

	if err != nil {
		panic(err)
	}

	//var showsWithInfo []bson.M
	var showsWithInfo []*domain.PpdbOption

	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	defer showInfoCursor.Close(ctx)
	return showsWithInfo
}

func (rr VocationalSchoolOptionRepository) GetVocationalSchoolOptionTemporary(ctx context.Context, level string, optType string) []*domain.PpdbOption {
	//TODO implement me

	registrationsCollection := rr.db.Collection("ppdb_options")

	matchStage := bson.D{{"$match", bson.M{
		"type": optType,
		/*
			"$and": []bson.M{bson.M{
				"school_id": bson.M{"$in": schoolIds},
			}}, */
		/*
			"$and": []bson.M{bson.M{
				"name": primitive.Regex{Pattern: "TANJUNGSARI", Options: ""},
			}}, */
	}}}
	//groupStage := bson.M{{"$group", bson.M{{"_id", "$podcast"}, {"total", bson.M{{"$sum", "$duration"}}}}}}

	pipeline := []bson.M{
		bson.M{"$match": bson.M{
			"$expr": bson.M{
				"$and": []bson.M{
					{"$eq": []string{"$_id", "$$school_id"}},
					{"$eq": []string{"$level", level}},
					//	{"$eq": []string{"$type", "negeri"}},
				},
			},
		}},
	}
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "ppdb_schools"},
		{"let", bson.D{{"school_id", "$school_id"}}},
		{"pipeline", pipeline},
		{"as", "ppdb_schools"}}}}
	unwindStage := bson.D{{"$unwind", "$ppdb_schools"}}
	sortByName := bson.D{{"$sort", bson.D{{"name", 1}}}}
	sortByType := bson.D{{"$sort", bson.D{{"type", 1}}}}
	//allowDisk := bson.D{{"allow", true}}
	//fields := bson.D{{"$project", bson.D{{"name", 1}}}}

	showInfoCursor, err := registrationsCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, lookupStage, unwindStage, sortByName, sortByType,
	}, options.Aggregate().SetAllowDiskUse(true))

	if err != nil {
		panic(err)
	}

	//var showsWithInfo []bson.M
	var showsWithInfo []*domain.PpdbOption

	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	defer showInfoCursor.Close(ctx)
	return showsWithInfo
}
