package difference

type IModel interface {
	GetDifferences(newModel interface{}) ([]Difference, error)
}
