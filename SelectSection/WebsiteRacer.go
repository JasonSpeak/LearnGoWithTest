package SelectSection

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeOut = 10 * time.Second

func Racer(url1, url2 string) {
	ConfigurableRacer(url1, url2, tenSecondTimeOut)
}

func ConfigurableRacer(url1, url2 string, timeOut time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeOut):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)

	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
