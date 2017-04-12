package rubber

type RubberFetcher interface {
	FetchRubbers() ([]*Rubber, error)
}

