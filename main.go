package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

const bufLen = 1024 * 1024 * 1024

func main() {
	go logStats()

	fmt.Println("Sleeping 10s...")
	time.Sleep(10 * time.Second)

	fmt.Println("Allocating 1GiB of nonsense...")
	fmt.Printf("Nonsense = %d\n", getNonsense())
	fmt.Println("Deallocated nonsense.")

	fmt.Println("Sleeping 1m...")
	time.Sleep(time.Minute)

	fmt.Println("runtime.GC()...")
	runtime.GC()

	fmt.Println("Sleeping 10s...")
	time.Sleep(10 * time.Second)

	fmt.Println("debug.FreeOSMemory()...")
	debug.FreeOSMemory()

	fmt.Println("Sleeping 1h...")
	time.Sleep(time.Hour)
}

func logStats() {
	for {
		time.Sleep(time.Second / 2)

		var mst runtime.MemStats
		runtime.ReadMemStats(&mst)

		fmt.Fprintf(os.Stderr, "MEM = %d @ %s\n", mst.Alloc, time.Now().Format(time.RFC3339))
	}
}

func getNonsense() (sum uint64) {
	buf := make([]byte, bufLen)

	rand.Read(buf)

	{
		var tmp byte

		for i := 0; i < bufLen; i += 2 {
			tmp = buf[i]
			buf[i] = buf[i+1]
			buf[i+1] = tmp
		}
	}

	for _, b := range buf {
		sum += uint64(b)
	}

	return
}
