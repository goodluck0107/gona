package boost

import (
	"log"
	"net/http"
	"runtime/pprof"

	// pprof hooks on localhost:6060
	_ "net/http/pprof"
)

// Profiling used for pro
func Profiling(path string, proc func()) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	f, err := OpenFile(NewFile, path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	proc()
}
