package gresp

type JSON struct {
	Code    int
	Message string
	Errors  map[string]string
	Data    interface{}
}
