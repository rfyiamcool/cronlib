package main

import (
	"log"

	"github.com/rfyiamcool/cronlib"
)

// start multi job
func main() {
	cron := cronlib.New()

	specList := map[string]string{
		"risk.scan.total.per.5s":  "*/5 * * * * *",
		"risk.scan.total.min.0s":  "0 * * * * *",
		"risk.scan.total.per.30s": "*/30 * * * * *",
	}

	for srv, spec := range specList {
		tspec := spec // copy
		ssrv := srv   // copy
		job, err := cronlib.NewJobModel(
			spec,
			func() {
				stdout(ssrv, tspec)
			},
		)
		if err != nil {
			panic(err.Error())
		}

		err = cron.Register(srv, job)
		if err != nil {
			panic(err.Error())
		}
	}

	cron.Start()
	log.Println("cron start")
	cron.Wait()
}

func stdout(srv, spec string) {
	log.Println(srv, spec)
}
