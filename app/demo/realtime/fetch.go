package realtime

import "fmt"

type fetchInterface interface {
	model() string
	clearTimeout() bool
	caller() ([]byte, error)
	updateData(content []byte)
	forceStart() bool
	fetchLine() (last, current int)
	isUpdate() bool
	readData() (content interface{})
	isStartUp() bool
}

type fetchControl struct {
	model        func() string
	clearTimeout func() bool
	caller       func() ([]byte, error)
	updateData   func(content []byte)
	forceStart   func() bool
	fetchLine    func() (last, current int)
	isUpdate     func() bool
	readData     func() (content interface{})
	isStartUp    func() bool
}

func newRealtimeBase(rt fetchInterface) *fetchControl {
	return &fetchControl{
		model:        rt.model,
		clearTimeout: rt.clearTimeout,
		caller:       rt.caller,
		updateData:   rt.updateData,
		forceStart:   rt.forceStart,
	}
}

func (r *fetchControl) update() {
	err := r.fetch(1)
	switch {
	case err != nil && !r.isStartUp():
		r.forceStart()
		return
	case err != nil:
		fmt.Println("warning fetch error")
		return
	default:
		r.snapshot()
	}
}

func (r *fetchControl) fetch(n int) error {
	content, err := r.caller()
	if err != nil {
		return err
	}
	r.dump(content)
	r.updateData(content)
	return nil
}

func (r *fetchControl) dump(content []byte) {
	if len(content) == 0 {
		return
	}
	//todo write to file
}

func (r *fetchControl) snapshot() {
	if !r.isUpdate() {
		return
	}
	r.readData()
	//todo write file
}

func (r *fetchControl) Run() {
	r.update()
	if r.clearTimeout() {
		// todo snapshot to file
	}
}
