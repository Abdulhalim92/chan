package basic

import (
	"fmt"
	"job/model"
	"time"
)

func Work(allData []model.SimpleData) {
	start := time.Now()
	for i, _ := range allData {
		Process(allData[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("Took =============> %s\n", elapsed)
}

func Process(data model.SimpleData) {
	fmt.Printf("Start processing %d\n", data.ID)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("finish processing %d\n", data.ID)
}
