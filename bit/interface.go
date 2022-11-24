package bit

type StbitInterface interface {
	On(id interface{}) error
	Off(id interface{}) error
	IsOn(id interface{}) bool
	GetValue() interface{}
	OnesCount() int
}
