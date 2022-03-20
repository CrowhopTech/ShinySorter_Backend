package mongoimg

import (
	"reflect"
	"testing"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_getQueriesForFilter(t *testing.T) {
	type args struct {
		filter *imagedb.ImageFilter
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			name: "empty",
			args: args{
				filter: nil,
			},
			want: bson.M{},
		},
		{
			name: "empty",
			args: args{
				filter: &imagedb.ImageFilter{},
			},
			want: bson.M{},
		},
		{
			name: "tag exists",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags: map[string]*imagedb.TagSearch{
						"tag1": nil,
					},
				},
			},
			want: bson.M{
				"tag1": bson.M{"$exists": true},
			},
		},
		{
			name: "multiple tags exist",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags: map[string]*imagedb.TagSearch{
						"tag1": nil,
						"tag2": nil,
					},
					RequireTagOperation: imagedb.All,
				},
			},
			want: bson.M{
				"$and": []bson.M{
					{"tag1": bson.M{"$exists": true}},
					{"tag2": bson.M{"$exists": true}},
				},
			},
		},
		{
			name: "any of tags exist",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags: map[string]*imagedb.TagSearch{
						"tag1": nil,
						"tag2": nil,
					},
					RequireTagOperation: imagedb.Any,
				},
			},
			want: bson.M{
				"$or": []bson.M{
					{"tag1": bson.M{"$exists": true}},
					{"tag2": bson.M{"$exists": true}},
				},
			},
		},
		{
			name: "tag value in",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags: map[string]*imagedb.TagSearch{
						"tag1": {TagValues: []string{"asdf"}},
						"tag2": {TagValues: []string{"fdsa"}},
					},
					RequireTagOperation: imagedb.Any,
				},
			},
			want: bson.M{
				"$or": []bson.M{
					{"tag1": bson.M{"$in": []string{"asdf"}}},
					{"tag2": bson.M{"$in": []string{"fdsa"}}},
				},
			},
		},
		{
			name: "really fkkin complicated query",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags: map[string]*imagedb.TagSearch{
						"tag1": {TagValues: []string{"asdf"}},
						"tag2": {TagValues: []string{"fdsa"}, Invert: true},
					},
					RequireTagOperation: imagedb.Any,
					ExcludeTags: map[string]*imagedb.TagSearch{
						"tag3": nil,
						"tag4": {TagValues: []string{"nope"}},
					},
					ExcludeTagOperation: imagedb.All,
				},
			},
			want: bson.M{
				"$and": []bson.M{
					{
						"$or": []bson.M{
							{
								"tag1": bson.M{
									"$in": []string{"asdf"},
								},
							},
							{
								"tag2": bson.M{
									"$nin": []string{"fdsa"},
								},
							},
						},
					},
					{
						"$and": []bson.M{
							{
								"tag3": bson.M{
									"$exists": false,
								},
							},
							{
								"tag4": bson.M{
									"$nin": []string{"nope"},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getQueriesForFilter(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getQueriesForFilter() = %v, want %v", got, tt.want)
			}

			// Got  map[$and:[map[$or:[map[tag1:map[$in:[asdf]]] map[tag2:map[$nin:[fdsa]]]]] map[$and:[map[tag3:map[$exists:false]] map[tag4:map[$nin:[nope]]]]]]]
			// Want map[$and:[map[$or:[map[tag1:map[$in:[asdf]]] map[tag2:map[$nin:[fdsa]]]]] map[$and:[map[tag3:map[$exists:true]] map[tag4:map[$in:[nope]]]]]]]
		})
	}
}
