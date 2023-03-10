package main

import (
	"fmt"
	"job/workerpool"
	"math/rand"
	"time"
)

func main() {

	/*
		var allData []model.SimpleData
		for i := 0; i < 1000; i++ {
			data := model.SimpleData{ID: i}
			allData = append(allData, data)
		}
		fmt.Printf("Start processing all work \n")
	*/

	// basic.Work(allData)
	// worker.NotPooledWork(allData)
	// worker.PooledWork(allData)
	// worker.PooledWorkError(allData)

	var allTask []*workerpool.Task
	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 5)
	go func() {
		for {
			taskID := rand.Intn(100) + 20

			if taskID%7 == 0 {
				pool.Stop()
			}

			time.Sleep(time.Duration(rand.Intn(5)))
			task := workerpool.NewTask(func(data interface{}) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, taskID)
			pool.AddTask(task)
		}
	}()
	pool.RunBackground()
}
