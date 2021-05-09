package realtime

import "sync"

type basic struct {
	start     bool
	modelName string
	current   int
	last      int
	update    bool
	sync.Mutex
}

func (p *basic) model() string {
	return p.modelName
}

func (p *basic) fetchLine() (last, current int) {
	return last, current
}

func (p *basic) updateLine(current int) {
	p.last = p.current
	p.current = current
}

func (p *basic) isUpdate() bool {
	return p.update
}

func (p *basic) isStartUp() bool {
	return p.start
}

func (p *basic) reloadFromLocal() ([]byte, error) {
	return nil, nil
}
