package mongoimg

import (
	"fmt"

	"github.com/CrowhopTech/shinysorter/backend/pkg/imagedb"
	"go.mongodb.org/mongo-driver/bson"
)

const tagsFieldPrefix = "tags."

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

// getQueryForTagSearch takes in a tag search (list of tag values and an inversion value),
// as well as whether this is for a Require or Exclude, and generates the proper query.
// This is what effectively makes the "end leaves" of the query tree
func getQueryForTagSearch(search *imagedb.TagSearch, exclude bool) bson.M {
	if search == nil {
		return bson.M{"$exists": !exclude} //  Normally true, if excluding should be false
	}

	if exclude != search.Invert { // Effectively XORing the two
		return bson.M{"$nin": search.TagValues}
	}

	return bson.M{"$in": search.TagValues}
}

// getComponentQuery returns the full either Require or Exclude query tree. Root is an and/or, with
// all the tag queries below it.
func getComponentQuery(tags map[string]*imagedb.TagSearch, op imagedb.TagOperation, exclude bool) bson.M {
	comps := []bson.M{}
	for key, search := range tags {
		query := bson.M{
			tagsFieldPrefix + key: getQueryForTagSearch(search, exclude),
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

	andComponents := []bson.M{}

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
