package lrucache

import (
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"
)

var (
	LRUCacheCreateFunc = NewSyncedLRUCache
	lruCache           = LRUCacheCreateFunc(time.Second * 1000)
)

func TestMain(m *testing.M) {
	flag.Parse()
	log.Println("NumCPU:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	os.Exit(m.Run())
}

func TestLruCacheBasicOperations(t *testing.T) {
	log.Println("Testing LRU cache basic operations ...")
	// t.Errorf("Counter is %d, but should be %d", counters[KEY], NGOROUTINES*N)
	lruCache := LRUCacheCreateFunc(time.Second)
	for i := 0; i < 100; i++ {
		lruCache.Put("abc", i)
		if lruCache.Get("abc").(int) != i {
			t.FailNow()
		}
	}
}

func TestLruCacheTimeout(t *testing.T) {
	log.Println("Testing LRU cache timeout ...")
	// t.Errorf("Counter is %d, but should be %d", counters[KEY], NGOROUTINES*N)
	lruCache := LRUCacheCreateFunc(time.Second)
	lruCache.Put("abc", "anything")
	time.Sleep(2 * time.Second)

	if lruCache.Get("abc") != nil {
		t.FailNow()
	}
}

func BenchmarkItoaOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(i)
	}
}

func BenchmarkLRUCachePut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := strconv.Itoa(i)
		lruCache.Put(k, i)
	}
}

func BenchmarkLRUCacheGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := strconv.Itoa(i)
		_ = lruCache.Get(k)
	}
}

func BenchmarkLRUCachePutAndGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := strconv.Itoa(i)
		lruCache.Put(k, i)
		_ = lruCache.Get(k)
	}
}
