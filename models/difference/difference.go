package difference

type Difference struct {
	Operation    string `json:"op,"` // enum: ["move", "remove", "replace", "addInto"]
	PropertyName string `json:"property_name,omitempty"`
	Value        any    `json:"value"`
}