package main

import (
	"log"

	"github.com/rfyiamcool/cronlib"
)

var (
	cron = cronlib.New()
)

func main() {
	handleClean()
	go start()

	// dynamic add
	handleBackup()

	select {}
}

func start() {
	cron.Start()
	cron.Wait()
}

func handleClean() {
	job, err := cronlib.NewJobModel(
		"*/5 * * * * *",
		func() {
			pstdout("do clean gc action")
		},
	)
	if err != nil {
		panic(err.Error())
	}

	err = cron.Register("clean", job)
	if err != nil {
		panic(err.Error())
	}
}

func handleBackup() {
	job, err := cronlib.NewJobModel(
		"*/5 * * * * *",
		func() {
			pstdout("do backup action")
		},
	)
	if err != nil {
		panic(err.Error())
	}

	err = cron.DynamicRegister("backup", job)
	if err != nil {
		panic(err.Error())
	}
}

func pstdout(srv string) {
	log.Println(srv)
}
