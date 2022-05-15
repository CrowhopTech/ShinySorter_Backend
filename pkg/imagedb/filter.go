package imagedb

type TagOperation int

const (
	Any TagOperation = 0
	All TagOperation = 1
)

type TagSearch struct {
	TagValues []string // Should always have at least one
	Invert    bool     // (e.g. do we want to see these, or not see these?)
}

type ImageFilter struct {
	Name   string
	Md5Sum *string

	// RequireTags is a list of required tags that must be
	// present on the image to return it in the search.
	RequireTags []int64

	// RequireTagOperation dictates, if RequireTags is set,
	// how multiple tags will be treated.
	RequireTagOperation TagOperation

	// ExcludeTags is a list of tags that will exclude a given
	// image from the search.
	ExcludeTags []int64

	// ExcludeTagOperation dictates, if ExcludeTags is set,
	// how multiple tags will be treated.
	ExcludeTagOperation TagOperation

	// Tagged dictates, if set, whether to return only
	// images that have been tagged or have not been tagged.
	// Note that this is not tied to the actual value of tags, it
	// is more intended for "A human has looked at this and verified"
	Tagged *bool
}
