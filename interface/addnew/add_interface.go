package addnew

type Add struct {
}

func (add *Add) AddSum(a int, b int) int {
	return a + b
}

func NewAdd() *Add { return &Add{} }
