package main

import (
	"sync"
	"time"
)

type SiteMonitor struct {
	statuses sync.Map

	sites chan string
}

func (s *SiteMonitor) updater(stoprCh chan struct{}) {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <- ticker.C:
			url := <- s.sites

			result := update(url)

			s.statuses.Store(url, result)

			s.sites <- url

		case <- stoprCh:
			return
		}
	}
}

func update(url string) int {
	return 0
}