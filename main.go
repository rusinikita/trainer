package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/rusinikita/trainer/choose"
)

func main() {
	p := tea.NewProgram(choose.New())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

//
// package main
//
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
//
// func callSomething(caller string) string {
// 	seconds := 1
//
// 	time.Sleep(time.Duration(seconds) * time.Second)
//
// 	return fmt.Sprintf("%s: %d", caller, seconds)
// }
//
// func main() {
// 	start := time.Now()
// 	var callResults []any
//
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func() {
// 		defer wg.Done()
// 		callResults = append(callResults, callSomething("first"))
// 	}()
//
// 	go func() {
// 		defer wg.Done()
// 		callResults = append(callResults, callSomething("second"))
// 	}()
//
// 	wg.Wait()
//
// 	fmt.Println("finished", int(time.Since(start)/time.Millisecond))
// 	fmt.Println(callResults...)
// }
