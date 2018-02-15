package main

import (
	"path/filepath"
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"time"
	"sync"
	"log"
)

type fileInfo struct {
	dirname  string
	fileSize int64
}

var verbose = flag.Bool("v", false, "show verbose progress messages")
var done = make(chan struct{})

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan fileInfo)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, root, &wg, fileSizes)
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

	sizeMap := make(map[string]int64)
loop:
	for {
		select {
		case <-done:
			log.Println("main's done")
			//既存のゴルーチンが完了できるようにfileSizesを空にする
			for range fileSizes {
				//何もしない
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			sizeMap[size.dirname] += size.fileSize
		case <-tick:
			printDiskUsage(sizeMap)
		}
	}
	printDiskUsage(sizeMap)
}

func printDiskUsage(sizemap map[string]int64) {
	for k, v := range sizemap {
		fmt.Printf("%v : %.1f GB", k, float64(v)/1e9)
	}
	fmt.Printf("\n")
}

func walkDir(rootdir string, dir string, wg *sync.WaitGroup, fileSize chan<- fileInfo) {
	defer wg.Done()
	if canceled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			wg.Add(1)
			go walkDir(rootdir, subdir, wg, fileSize)
		} else {
			fileSize <- fileInfo{rootdir, entry.Size()}
		}
	}
}

//係数セマフォを使用しないとtoo many filesで怒られる
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: //トークン取得
	case <-done:
		log.Println("dirents's done")
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}
