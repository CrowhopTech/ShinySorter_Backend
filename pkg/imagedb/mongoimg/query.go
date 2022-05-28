package mongoimg

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"go.mongodb.org/mongo-driver/bson"
)

const tagsField = "tags"

// tagOperationToKey simply converts a tag operation (imagedb.Any or imagedb.All)
// into the corresponding MongoDB and/or string. Panics if an invalid value is given.
func tagOperationToKey(operation imagedb.TagOperation) string {
	switch operation {
	case imagedb.All:
		return "$and"
	case imagedb.Any:
		return "$or"
	}
	panic(fmt.Errorf("unknown tag operation %d", operation))
}

// getComponentQuery returns the full either Require or Exclude query tree. Root is an and/or, with
// all the tag queries below it.
func getComponentQuery(tags []int64, op imagedb.TagOperation, exclude bool) bson.M {

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
func getQueriesForFilter(filter *imagedb.ImageFilter) bson.M {
	if filter == nil {
		return bson.M{}
	}

	andComponents := []bson.M{
		{"hasContent": !filter.MissingContent}, // By default only includes images with content, but can be inverted
	}

	if filter.Name != "" {
		andComponents = append(andComponents, bson.M{
			"_id": filter.Name,
		})
	}

	if filter.Md5Sum != nil { // Empty string is still valid here: search for all images without an md5sum set is useful for db populator
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
