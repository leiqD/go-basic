package realtime

type realTimeIns struct {
	popRealTime *popRealTime
}

type manager struct {
	rtIns       realTimeIns
	popRealTime *fetchControl
}

func NewManager() *manager {
	ins := &manager{}
	ins.rtIns.popRealTime = newPopRealTime()
	ins.popRealTime = newRealtimeBase(ins.rtIns.popRealTime)
	return ins
}

func (m *manager) PopRealTime() PopRealTime {
	return m.rtIns.popRealTime
}

func (m *manager) Run() {
	m.popRealTime.Run()
}
