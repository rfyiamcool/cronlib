package main

import (
	"log"
	"time"

	"github.com/rfyiamcool/cronlib"
)

var (
	cron = cronlib.New()
)

func main() {
	handleClean()

	time.AfterFunc(time.Duration(2*time.Second), func() {
		// dynamic add
		handleBackup()
	})

	time.AfterFunc(
		12*time.Second,
		func() {
			log.Println("stop clean")
			cron.StopService("clean")

			log.Println("stop backup")
			cron.StopService("backup")
		},
	)

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
