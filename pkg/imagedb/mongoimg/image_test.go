package mongoimg

import (
	"testing"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_mongoConnection_getUpdateParameter(t *testing.T) {
	trueRef := true
	type args struct {
		i *imagedb.Image
	}
	tests := []struct {
		name    string
		args    args
		want    bson.M
		wantErr bool
	}{
		{
			name: "nil image",
			args: args{
				i: nil,
			},
			want:    bson.M{},
			wantErr: false,
		},
		{
			name: "empty image",
			args: args{
				i: &imagedb.Image{},
			},
			want:    bson.M{"$set": bson.M{}},
			wantErr: false,
		},
		{
			name: "full image",
			args: args{
				i: &imagedb.Image{
					FileMetadata: imagedb.FileMetadata{
						Name:     "test",
						Md5Sum:   "testsum",
						MIMEType: "test/mime",
					},
					Tags:          &[]int64{1, 2, 3},
					HasBeenTagged: &trueRef,
					HasContent:    &trueRef,
				},
			},
			want: bson.M{"$set": bson.M{
				"md5sum":        "testsum",
				"mimeType":      "test/mime",
				"tags":          &[]int64{1, 2, 3},
				"hasBeenTagged": true,
				"hasContent":    true,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &mongoConnection{}
			got, err := mc.getUpdateParameter(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("mongoConnection.getUpdateParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
