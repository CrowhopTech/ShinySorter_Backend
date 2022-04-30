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
					RequireTags: []int64{1},
				},
			},
			want: bson.M{
				"tags": int64(1),
			},
		},
		{
			name: "multiple tags exist",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags:         []int64{1, 2},
					RequireTagOperation: imagedb.All,
				},
			},
			want: bson.M{
				"$and": []bson.M{
					{"tags": int64(1)},
					{"tags": int64(2)},
				},
			},
		},
		{
			name: "any of tags exist",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags:         []int64{1, 2},
					RequireTagOperation: imagedb.Any,
				},
			},
			want: bson.M{
				"$or": []bson.M{
					{"tags": int64(1)},
					{"tags": int64(2)},
				},
			},
		},
		{
			name: "really fkkin complicated query",
			args: args{
				filter: &imagedb.ImageFilter{
					RequireTags:         []int64{1, 2, 3},
					RequireTagOperation: imagedb.Any,
					ExcludeTags:         []int64{4, 5},
					ExcludeTagOperation: imagedb.All,
				},
			},
			want: bson.M{
				"$and": []bson.M{
					{
						"$or": []bson.M{
							{"tags": int64(1)},
							{"tags": int64(2)},
							{"tags": int64(3)},
						},
					},
					{
						"$and": []bson.M{
							{"$not": bson.M{"tags": int64(4)}},
							{"$not": bson.M{"tags": int64(5)}},
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
		})
	}
}
