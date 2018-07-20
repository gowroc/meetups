type Error interface {
	Code() ErrorCode // identifies a valid error type
	Msg() string     // free-form human-readable message

	WithMeta(key string, val string) Error // set metadata
	Meta(key string) string                // get metadata value
	MetaMap() map[string]string            // see all metadata

	Error() string // as an error returns "twirp error <Code>: <Msg>"
}