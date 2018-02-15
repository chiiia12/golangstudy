package main

import (
	"path/filepath"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"time"
	"sync"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	sizeMap := make(map[string]int64)
	for _, root := range roots {
		var fileSizeSum int64
		sizeMap[root] = fileSizeSum
		wg.Add(1)
		go walkDir(root, &wg, fileSizes, &fileSizeSum)
	}
	go func() {
		wg.Wait()
		//すべてのgoroutineが終わってから閉じる
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes, &sizeMap)
		}
	}
	printDiskUsage(nfiles, nbytes, &sizeMap)
}

func printDiskUsage(nfiles, nbytes int64, sizeMap *map[string]int64) {
	//fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
	for k, v := range *sizeMap {
		fmt.Printf("%v :%v GB ", k, v)
	}
	fmt.Printf("\n")
}

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64, sum *int64) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(subdir, wg, fileSizes, sum)
		} else {
			*sum = *sum + entry.Size()
			fmt.Println("fileSize is ", *sum)
			fileSizes <- entry.Size()
		}
	}
}

//係数セマフォを使用しないとtoo many filesで怒られる
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
