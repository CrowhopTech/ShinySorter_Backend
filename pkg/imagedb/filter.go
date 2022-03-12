package imagedb

type ImageFilter struct {
	Name   *string
	Md5Sum *string

	// RequireTags is a map of required tags on the image
	// If you pass a key with a nil value, it will just
	// require the key to exist on the item. If you pass
	// a value, the tag must have the proper value. Note
	// that an empty string *is* a valid value and will
	// filter specifically for those with empty string values
	// (e.g. the it's marked as "yes, this is relevant" but the
	// exact value isn't known yet or isn't relevant)
	RequireTags map[string]*string
	// ExcludeTags is a map of tags on the image that will
	// exclude it from the result. If you pass a key with a
	// nil value, any image with that key in its tags will
	// be excluded. If any value is passed, it will only exclude
	// images matching that exact value (including empty string).
	ExcludeTags map[string]*string
}
