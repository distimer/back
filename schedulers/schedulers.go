package schedulers

import "github.com/robfig/cron/v3"

func GenerateSchedularObj() *cron.Cron {
	c := cron.New()
	deleteUserPurgeSchedule(c)

	return c
}
