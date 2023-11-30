package FirstPatternFactory

// InspectReq 巡检请求体
type InspectReq struct {
	// 巡检名称
	Name string
	// 执行策略
	Policy string
	// 巡检项ID
	Items []int
	// 巡检对象
	Object []InspectOBJ
}

// InspectOBJ 巡检对象
type InspectOBJ struct {
	// 巡检类型
	Type string
	// 巡检对象ID
	ID string
}
