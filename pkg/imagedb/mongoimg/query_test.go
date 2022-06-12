package mongoimg

import (
	"reflect"
	"testing"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_getQueriesForFilter(t *testing.T) {
	var (
		testSum = "testsum"
		trueRef = true
	)

	type args struct {
		filter *imagedb.ImageFilter
	}
	tests := []struct {
		name string
		args args
		want bson.M
	}{
		{
			// Nil filter
			name: "nil filter",
			args: args{
				filter: nil,
			},
			want: bson.M{"hasContent": true},
		},
		{
			// Empty filter
			name: "Empty filter",
			args: args{
				filter: &imagedb.ImageFilter{},
			},
			want: bson.M{"hasContent": true},
		},
		{
			// Complex filter
			name: "Complex filter",
			args: args{
				filter: &imagedb.ImageFilter{
					Name:                "test",
					Md5Sum:              &testSum,
					MissingContent:      true,
					Tagged:              &trueRef,
					RequireTags:         []int64{1, 2, 3},
					RequireTagOperation: imagedb.All,
					ExcludeTags:         []int64{4, 5, 6},
					ExcludeTagOperation: imagedb.Any,
				},
			},
			want: bson.M{
				"$and": []bson.M{
					{"hasContent": false},
					{"_id": "test"},
					{"md5sum": "testsum"},
					{"hasBeenTagged": true},
					{"$and": []bson.M{
						{"tags": int64(1)},
						{"tags": int64(2)},
						{"tags": int64(3)},
					}},
					{"$or": []bson.M{
						{"tags": bson.M{"$ne": int64(4)}},
						{"tags": bson.M{"$ne": int64(5)}},
						{"tags": bson.M{"$ne": int64(6)}},
					}},
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
