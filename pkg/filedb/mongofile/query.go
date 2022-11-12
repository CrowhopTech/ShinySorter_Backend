package mongofile

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const tagsField = "tags"

// tagOperationToKey simply converts a tag operation (filedb.Any or filedb.All)
// into the corresponding MongoDB and/or string. Panics if an invalid value is given.
func tagOperationToKey(operation filedb.TagOperation) string {
	switch operation {
	case filedb.All:
		return "$and"
	case filedb.Any:
		return "$or"
	}
	panic(fmt.Errorf("unknown tag operation %d", operation))
}

// getComponentQuery returns the full either Require or Exclude query tree. Root is an and/or, with
// all the tag queries below it.
func getComponentQuery(tags []int64, op filedb.TagOperation, exclude bool) bson.M {

	comps := []bson.M{}

	for _, tag := range tags {
		var query bson.M
		if exclude {
			query = bson.M{
				tagsField: bson.M{
					"$ne": tag,
				},
			}
		} else {
			query = bson.M{
				tagsField: tag,
			}
		}

		comps = append(comps, query)
	}

	if len(comps) == 0 {
		return bson.M{}
	}

	if len(comps) == 1 {
		return comps[0]
	}

	return bson.M{
		tagOperationToKey(op): comps,
	}
}

// getQueriesForFilter takes the given filter and constructs
// all the various filter components, joining them together with
// an "and" at the end if there are multiple.
func getQueriesForFilter(filter *filedb.FileFilter) bson.M {
	hasContentQuery := bson.M{"hasContent": true}

	if filter == nil {
		return hasContentQuery
	}

	// By default only includes files with content, but can be inverted
	hasContentQuery["hasContent"] = !filter.MissingContent
	andComponents := []bson.M{
		hasContentQuery,
	}

	if filter.ID != primitive.NilObjectID {
		andComponents = append(andComponents, bson.M{
			"_id": filter.ID,
		})
	}

	if filter.Name != "" {
		andComponents = append(andComponents, bson.M{
			"name": filter.Name,
		})
	}

	if filter.Md5Sum != nil { // Empty string is still valid here: search for all files without an md5sum set is useful for db populator
		andComponents = append(andComponents, bson.M{
			"md5sum": *filter.Md5Sum,
		})
	}

	if filter.Tagged != nil {
		andComponents = append(andComponents, bson.M{
			"hasBeenTagged": *filter.Tagged,
		})
	}

	requiredQuery := getComponentQuery(filter.RequireTags, filter.RequireTagOperation, false)
	if len(requiredQuery) > 0 {
		andComponents = append(andComponents, requiredQuery)
	}

	excludedQuery := getComponentQuery(filter.ExcludeTags, filter.ExcludeTagOperation, true)
	if len(excludedQuery) > 0 {
		andComponents = append(andComponents, excludedQuery)
	}

	if len(andComponents) == 0 {
		return bson.M{}
	}

	if len(andComponents) == 1 {
		return andComponents[0]
	}

	return bson.M{
		"$and": andComponents,
	}
}
