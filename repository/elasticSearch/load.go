package elasticSearch

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"strings"
)

func LoadDatasFromFile(ctx context.Context) context.Context {
	const (
			concurrency = 5
			usersFile   = "dados.csv"
	)

	type GetPresenter struct{
		ValorConvertido	string `json:"valorConvertido"`
		SimboloMoeda	string	`json:"simboloMoeda"`
	}

	var (
			datas     []GetPresenter
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
					line := scanner.Text()
					if strings.HasSuffix(line, ",$") {
							line = strings.TrimSuffix(line, ",$")
					}
					workQueue <- line
			}
			close(workQueue)
	}()



	for i := 0; i < concurrency; i++ {
			waitGroup.Add(1)
			go func(workQueue chan string, waitGroup *sync.WaitGroup) {
					for entry := range workQueue {
							parts := strings.Split(entry, ",")
							data := GetPresenter{
								ValorConvertido:  parts[0],
								SimboloMoeda: parts[1],
							}
							mutex.Lock()
							datas = append(datas, data)
							mutex.Unlock()
					}
					waitGroup.Done()
			}(workQueue, waitGroup)
	}

	waitGroup.Wait()

	fmt.Printf("✅ Datas loaded from the file: %d \n", len(datas))
	return context.WithValue(ctx, DataKey, datas)
}
