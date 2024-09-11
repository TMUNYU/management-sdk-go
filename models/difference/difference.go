package difference

type Difference struct {
	Operation    string `json:"op"`
	PropertyName string `json:"property_name"`
	Value        any    `json:"value"`
}