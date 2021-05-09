package realtime

import "encoding/json"

type PopRealTime interface {
	Query() bool
}

type popRealTime struct {
	basic
	data map[int64]int
}

func newPopRealTime() *popRealTime {
	return &popRealTime{}
}

func (p *popRealTime) updateData(content []byte) {

}

func (p *popRealTime) readData() (content interface{}) {
	return p.data
}

func (p *popRealTime) clearTimeout() bool {
	return false
}

func (p *popRealTime) caller() ([]byte, error) {
	return nil, nil
}

func (p *popRealTime) forceStart() bool {
	content, err := p.reloadFromLocal()
	if err != nil {
		return false
	}
	if err := json.Unmarshal(content, p.data); err != nil {
		return false
	}
	return true
}

func (p *popRealTime) Query() bool {
	return true
}
