package mocks

var (
	GetInitFunc func() error
	GetFunc func(key string) interface{}
)

type MockEnv struct {
	InitFunc func() error
	GetFunc func(key string) interface{}
}


func (m *MockEnv) Init() error {
	return GetInitFunc()
}

func (m *MockEnv) Get(key string) interface{} {
	return GetFunc(key)
}
