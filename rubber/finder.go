package rubber

type RubberFinder interface {
	FindRubbers() ([]*Rubber, error)
}

