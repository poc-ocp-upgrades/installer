package asset

type FileFetcher interface {
	FetchByName(string) (*File, error)
	FetchByPattern(pattern string) ([]*File, error)
}
