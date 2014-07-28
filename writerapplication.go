package gost

import (
    "fmt"
)

type WriterApplication struct {
    workers     []Worker        // the instanciated worker
}

func NewWriterApplication() *WriterApplication {
    return &WriterApplication{}
}

func (a *WriterApplication) Start(gost Gost) {
    a.initWorkers(gost)
}

func (a *WriterApplication) Stop() {
    for i := 0; i < len(a.workers); i++ {
        a.workers[i].Stop()
    }
}

// Inits the workers
func (a *WriterApplication) initWorkers(gost Gost) {
    a.workers = make([]Worker, 1)
    printer := NewPrinterWorker("printer_application", "writer")
    err := printer.Start(gost)
    if err != nil {
        fmt.Println("[APPLICATION] [writer] ERROR - while starting the workers of WriterApplication : ")
        fmt.Printf("[APPLICATION] [writer] ERROR - %s\n", err)
    } else {
        fmt.Println("[APPLICATION] [writer] Writer application started.")
        a.workers = append(a.workers, printer)
    }
}
