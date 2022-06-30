package scheduler

import (
	"github.com/amortaza/aceql/bsn/grpcclient"
	bsntime "github.com/amortaza/aceql/bsn/time"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"time"
)

var CADENCE_SECONDS time.Duration = 15

//todo in gui instead of date, tell me how long ago
func StartScheduler() {
	go func() {
		for {
			time.Sleep(CADENCE_SECONDS * time.Second)

			r := stdsql.NewRecord("x_jobs")
			r.Add("x_active", query.Equals, "true")
			r.Query()

			for {
				hasNext, _ := r.Next()

				if !hasNext {
					break
				}

				if !atCadence(r) {
					continue
				}

				v, _ := r.Get("x_script_name")
				grpcclient.GRPC_CallScript(v)

				setLastRun(r)
			}
		}
	}()
}

func atCadence(r *flux.Record) bool {
	//todo add .nil()
	//todo give error and return nil when field does not exist
	//todo is this prepending a space !??
	starting, _ := r.Get("x_starting_datetime")
	starting = bsntime.Normalize(starting)
	if starting == "" {
		return false
	}

	if bsntime.IsAfter(starting, bsntime.Now()) {
		return false
	}

	lastRun, _ := r.Get("x_last_run")
	lastRun = bsntime.Normalize(lastRun)
	if lastRun == "" {
		return true
	}

	seconds, _ := r.Get("x_seconds")
	if seconds == "" {
		return false
	}

	nextRun := bsntime.AddSeconds(lastRun, seconds)

	if bsntime.IsAfter(bsntime.Now(), nextRun) {
		return true
	}

	return false
}

func setLastRun(r *flux.Record) {
	r.Set("x_last_run", bsntime.Now())
	r.Update()
}
