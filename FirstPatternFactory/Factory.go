package FirstPatternFactory

import (
	"log"
	"sync"
)

/*
	伴随着巡检业务的迭代开始对接越来越多的南向服务再使用流式处理
越来越多的巡检项以及业务节点，使得业务代码变得越来越难以维护。
*/

/*
	工厂类型的设计模式刚好使得再面对多个子系统时变的易于维护
*/

// Inspecter 巡检接口
type Inspecter interface {
	//Do  各个对接子系统需要实现的巡检方法
	Do(InspectReq) error
}

// FirstSvc 第一个巡检南向系统
type FirstSvc struct {
	IP string
}

func NewFirstSvc(ip string) Inspecter {
	return &FirstSvc{
		IP: ip,
	}
}

func (s *FirstSvc) Do(req InspectReq) error {
	// 第一个系统的业务逻辑处理
	log.Printf("first service do inspect.req:%+v", req)
	return nil
}

// SecondSvc 第二个巡检南向系统
type SecondSvc struct {
	IP string
}

func NewSecondSvc(ip string) Inspecter {
	return &SecondSvc{
		IP: ip,
	}
}

func (s *SecondSvc) Do(req InspectReq) error {
	log.Printf("second service do inspect.req:%+v", req)
	return nil
}

type Impl struct {
	ImplOBJ map[string]Inspecter
	m       sync.Mutex
}

func NewImpl() *Impl {
	return &Impl{
		ImplOBJ: make(map[string]Inspecter),
		m:       sync.Mutex{},
	}
}

// AddSvc 添加微服务到处理器中
func (i *Impl) AddSvc(svcType string, inspect Inspecter) {
	i.m.Lock()
	defer i.m.Unlock()
	i.ImplOBJ[svcType] = inspect
}

// Dispatcher 服务调度执行所有子系统
func (i *Impl) Dispatcher(req InspectReq) {
	for s := range i.ImplOBJ {
		err := i.ImplOBJ[s].Do(req)
		if err != nil {
			log.Printf("Failed to do inpect Type:[%s]", s)
			return
		}
	}
}
