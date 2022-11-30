package supafile

import (
	"fmt"
	"strconv"

	"github.com/CrowhopTech/shinysorter/backend/pkg/filedb"
	postgrest_go "github.com/supabase/postgrest-go"
)

// getQueriesForFilter takes the given filter and constructs
// all the various filter components, joining them together with
// an "and" at the end if there are multiple.
func getQueriesForFilter(requestBuilder *postgrest_go.FilterBuilder, filter *filedb.FileFilter) error {

	// Always include hasContent as a query (default only with content)
	*requestBuilder = *requestBuilder.Eq("hasContent", strconv.FormatBool(!filter.MissingContent))
	*requestBuilder = *requestBuilder.Eq("hasBeenTagged", strconv.FormatBool(*filter.Tagged))

	if filter.Name != "" {
		*requestBuilder = *requestBuilder.Eq("filename", filter.Name)
	}
	if filter.ID != "" {
		*requestBuilder = *requestBuilder.Eq("id", filter.ID)
	}
	if filter.Continue != "" {
		*requestBuilder = *requestBuilder.Gt("id", filter.Continue)
	}
	if filter.Md5Sum != nil {
		*requestBuilder = *requestBuilder.Eq("md5sum", *filter.Md5Sum)
	}

	return fmt.Errorf("not implemented")

	// hasContentQuery := bson.M{}

	// if filter == nil {
	// 	return nil
	// }

	// // By default only includes files with content, but can be inverted
	// hasContentQuery["hasContent"] = !filter.MissingContent
	// andComponents := []bson.M{
	// 	hasContentQuery,
	// }

	// if filter.ID != "" {
	// 	parsedID, err := primitive.ObjectIDFromHex(filter.ID)
	// 	if err != nil {
	// 		return fmt.Errorf("invalid bson object ID '%s'", filter.ID)
	// 	}
	// 	andComponents = append(andComponents, bson.M{
	// 		"_id": parsedID,
	// 	})
	// }

	// if filter.Name != "" {
	// 	andComponents = append(andComponents, bson.M{
	// 		"name": filter.Name,
	// 	})
	// }

	// if filter.Md5Sum != nil { // Empty string is still valid here: search for all files without an md5sum set is useful for db populator
	// 	andComponents = append(andComponents, bson.M{
	// 		"md5sum": *filter.Md5Sum,
	// 	})
	// }

	// if filter.Tagged != nil {
	// 	andComponents = append(andComponents, bson.M{
	// 		"hasBeenTagged": *filter.Tagged,
	// 	})
	// }

	// if filter.Continue != primitive.NilObjectID {
	// 	andComponents = append(andComponents, bson.M{
	// 		"_id": bson.M{"$gt": filter.Continue},
	// 	})
	// }

	// requiredQuery := getComponentQuery(filter.RequireTags, filter.RequireTagOperation, false)
	// if len(requiredQuery) > 0 {
	// 	andComponents = append(andComponents, requiredQuery)
	// }

	// excludedQuery := getComponentQuery(filter.ExcludeTags, filter.ExcludeTagOperation, true)
	// if len(excludedQuery) > 0 {
	// 	andComponents = append(andComponents, excludedQuery)
	// }

	// if len(andComponents) == 0 {
	// 	return bson.M{}, nil
	// }

	// if len(andComponents) == 1 {
	// 	return andComponents[0], nil
	// }

	// return bson.M{
	// 	"$and": andComponents,
	// }, nil
}
