package scheduler

import (
	"github.com/amortaza/aceql/bsn/grpc_client"
	bsntime "github.com/amortaza/aceql/bsn/time"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
	"time"
)

var CADENCE_SECONDS time.Duration = 15

//todo in gui instead of date, tell me how long ago
func StartScheduler() {
	go func() {
		for {
			time.Sleep(CADENCE_SECONDS * time.Second)

			r, err := stdsql.NewRecord("x_jobs")
			if err != nil {
				logger.Info("Stopping Scheduler, see logs", "Scheduler")
				return
			}

			if err := r.Add("x_active", query.Equals, "true"); err != nil {
				logger.Info("Stopping Scheduler, see logs", "Scheduler")
				return
			}

			if _, err := r.Query(); err != nil {
				logger.Info("Stopping Scheduler, see logs", "Scheduler")
				return
			}

			for {
				hasNext, err := r.Next()

				if err != nil {
					logger.Info("Stopping Scheduler. r.Next(), see logs", "Scheduler")
					break
				}

				if !hasNext {
					break
				}

				at, err := atCadence(r)
				if err != nil {
					logger.Info("Stopping Scheduler. atCadence(), see logs", "Scheduler")
					break
				}

				if !at {
					continue
				}

				v, err := r.Get("x_script_name")
				if err != nil {
					logger.Info("Stopping Scheduler. r.Get(script_name), see logs", "Scheduler")
					break
				}

				if err := grpc_client.GRPC_CallScript("../js/scheduled_jobs", v, nil); err != nil {
					logger.Err(err, "???")
					continue
				}

				if err := setLastRun(r); err != nil {
					logger.Err(err, "???")
					continue
				}
			}
		}
	}()
}

func atCadence(r *flux.Record) (bool, error) {
	//todo add .nil()
	//todo give error and return nil when field does not exist
	//todo is this prepending a space !??
	starting, err := r.Get("x_starting_datetime")
	if err != nil {
		return false, err
	}

	starting = bsntime.Normalize(starting)
	if starting == "" {
		return false, nil
	}

	if bsntime.IsAfter(starting, bsntime.Now()) {
		return false, nil
	}

	lastRun, err := r.Get("x_last_run")
	if err != nil {
		return false, err
	}

	lastRun = bsntime.Normalize(lastRun)
	if lastRun == "" {
		return true, nil
	}

	seconds, err := r.Get("x_seconds")
	if err != nil {
		return false, err
	}

	if seconds == "" {
		return false, nil
	}

	nextRun, err := bsntime.AddSeconds(lastRun, seconds)
	if err != nil {
		return false, err
	}

	if bsntime.IsAfter(bsntime.Now(), nextRun) {
		return true, nil
	}

	return false, nil
}

func setLastRun(r *flux.Record) error {
	if err := r.Set("x_last_run", bsntime.Now()); err != nil {
		return err
	}

	if err := r.Update(); err != nil {
		return err
	}

	return nil
}
