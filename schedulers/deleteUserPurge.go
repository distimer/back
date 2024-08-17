package schedulers

import (
	"context"
	"time"

	"github.com/robfig/cron/v3"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent/deleteduser"
	"pentag.kr/distimer/utils/logger"
)

func deleteUserPurgeSchedule(c *cron.Cron) {
	_, err := c.AddFunc("0 * * * *", deleteUserPurge)
	if err != nil {
		logger.Fatal(err)
	}
}

func deleteUserPurge() {
	now := time.Now()

	// Delete users who have been deleted for more than 7 days
	dbConn := db.GetDBClient()
	_, err := dbConn.DeletedUser.Delete().Where(deleteduser.DeletedAtLT(now.AddDate(0, 0, -configs.ReRegisterDay))).Exec(context.Background())
	if err != nil {
		logger.Error(err)
	}
}
