package main

import (
	"beyond_effective/chapter_03/pprof/cpu_profiling/game"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	router.HandleFunc("/", CardShuffler)

	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	server := &http.Server{
		Addr:    ":8888",
		Handler: router,
	}

	go func() {
		_ = server.ListenAndServe()
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh
	_ = server.Close()
}

func CardShuffler(resp http.ResponseWriter, req *http.Request) {
	cards := game.NewDeck()
	for x := 0; x < 100; x++ {
		game.Shuffle(cards)
	}

	for index, card := range cards {
		if index > 0 {
			_, _ = resp.Write([]byte(", "))
		}
		_, _ = resp.Write([]byte(card.Face))
		_, _ = resp.Write([]byte(card.Suit))
	}
}