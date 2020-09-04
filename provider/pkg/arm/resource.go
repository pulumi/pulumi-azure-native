package arm

type resource struct {
	name string
	properties map[string]interface{}
}

type dag struct {
	vertices map[string]interface{}
}