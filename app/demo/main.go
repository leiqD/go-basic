package main

import "github.com/leiqD/go-socket5/app/demo/realtime"

func main() {
	manager := realtime.NewManager()
	manager.Run()
	manager.PopRealTime().Query()
}
