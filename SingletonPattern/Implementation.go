package main

import "sync"

var once sync.Once
var instance *ComponentInstance

type ComponentInstance struct {
	InitCount int
	Count     int
}

func GetInstance() *ComponentInstance {
	once.Do(func() {
		instance = new(ComponentInstance)
		instance.InitCount = 1
	})
	instance.Count++
	return instance
}
