package server

type CasePool struct {
	CaseData   map[string]func()
	BeforePool map[string]func()
	AfterPool  map[string]func()
}

func New() *CasePool {

	return &CasePool{}
}

func (cp *CasePool) Start() {

}

//执行所有用例
func (cp *CasePool) Execute() {
	for k, v := range cp.CaseData {
		cp.BeforePool[k]()
		v()
		cp.AfterPool[k]()
	}

}

//增加一个简单的 case
func (cp *CasePool) AddCase(funcName string, f func()) {
	cp.CaseData[funcName] = f
}

//注册一个 用例的case
func (cp *CasePool) RegisterCase(funcName string, caseFunc, Before, After func()) {
	cp.CaseData[funcName] = caseFunc
	cp.BeforePool[funcName] = Before
	cp.AfterPool[funcName] = After
}
