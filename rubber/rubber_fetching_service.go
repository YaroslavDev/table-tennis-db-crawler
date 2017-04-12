package rubber

type RubberFetchingService interface {
	FetchRubbers() ([]*Rubber, error)
}

