package main

import (
	"fmt"
	"github.com/tomas-mraz/android-go/android"
	"github.com/tomas-mraz/android-go/app"
	"simple-vulkan-app/vulkandraw"
)

func init() {
	app.SetLogTag("VulkanDraw")
}

func main() {
	nativeWindowEvents := make(chan app.NativeWindowEvent)
	inputQueueEvents := make(chan app.InputQueueEvent, 1)
	inputQueueChan := make(chan *android.InputQueue, 1)

	app.Main(func(a app.NativeActivity) {

		a.HandleNativeWindowEvents(nativeWindowEvents)
		a.HandleInputQueueEvents(inputQueueEvents)
		// just skip input events (so app won't be dead on touch input)
		go app.HandleInputQueues(inputQueueChan, func() {
			a.InputQueueHandled()
		}, app.SkipInputEvents)

		var vkDevice *vulkandraw.VulkanDeviceInfo

		a.InitDone()

		for {
			select {
			case <-a.LifecycleEvents():
				// ignore
			case event := <-inputQueueEvents:
				switch event.Kind {
				case app.QueueCreated:
					fmt.Println("INPUT Created")
					inputQueueChan <- event.Queue
				case app.QueueDestroyed:
					fmt.Println("INPUT Destroyed")
					inputQueueChan <- nil
				}
			case event := <-nativeWindowEvents:
				switch event.Kind {
				case app.NativeWindowCreated:
					fmt.Println("C R E A T E D")
					//vk.Init()
				case app.NativeWindowDestroyed:
					fmt.Println("D E S T R O Y E D")
					vkDevice.Destroy()
				case app.NativeWindowRedrawNeeded:
					fmt.Println("R E D R A W")
					a.NativeWindowRedrawDone()
				}
			}
		}
	})
}
