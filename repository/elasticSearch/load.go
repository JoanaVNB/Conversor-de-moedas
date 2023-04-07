package elasticSearch

import (
	"exchange/presenter"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

func LoadDatasFromFile(ctx context.Context) context.Context {
const (
	concurrency = 5
	usersFile  = "dados.csv"
)

var (
	datas    []presenter.Presenter
	waitGroup = new(sync.WaitGroup)
	workQueue = make(chan string)
	mutex     = &sync.Mutex{}
)

go func() {
	usersFile, err := os.Open(usersFile)
	if err != nil {
		panic(err)
	}
	defer usersFile.Close()
	scanner := bufio.NewScanner(usersFile)
	for scanner.Scan() {
		workQueue <- scanner.Text()
	}
	close(workQueue)
}()

var data presenter.Presenter
for i := 0; i < concurrency; i++ {
	waitGroup.Add(1)
	go func(workQueue chan string, waitGroup *sync.WaitGroup) {
		for entry := range workQueue {
			err := json.Unmarshal([]byte(entry), &data)
			if err == nil {
				mutex.Lock()
				datas = append(datas, data)
				mutex.Unlock()
			}
		}
		waitGroup.Done()
	}(workQueue, waitGroup)
}

waitGroup.Wait()

fmt.Printf("✅ Datas loaded from the file: %d \n", len(datas))
return context.WithValue(ctx, DataKey, datas)
}