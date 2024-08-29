package notify

import (
	"pentag.kr/distimer/ent"
)

func SendTimerCreate(userID string, timerObj *ent.Timer, subjectObj *ent.Subject) {

	// onesignalUser, err := GetOneSignalUser(userID)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	// androidSubs, _ := SplitAndroidIOSSubscriptions(onesignalUser.Subscriptions)

	// data := AndroidTimerNotificationReq{
	// 	AppID:            configs.Env.OneSignalAppID,
	// 	TargetChannel:    "push",
	// 	ContentAvailable: true,
	// 	Priority:         10,
	// 	Data: AndroidTimerNotificationData{
	// 		Type:    "start",
	// 		Subject: subjectObj.Name,
	// 		StartAt: timerObj.StartAt.Format(time.RFC3339),
	// 		Content: timerObj.Content,
	// 		Color:   subjectObj.Color,
	// 	},
	// 	IncludeSubsriptionIDs: androidSubs,
	// }

	// fmt.Println("---------------------")
	// fmt.Println(androidSubs)
	// fmt.Println(data)

	// err = AndroidTimerNotification(data)
	// if err != nil {
	// 	logger.Error(err)
	// }
}

func SendTimerUpdate(userID string, timerObj *ent.Timer, subjectObj *ent.Subject) {

	// onesignalUser, err := GetOneSignalUser(userID)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// androidSubs, _ := SplitAndroidIOSSubscriptions(onesignalUser.Subscriptions)

	// data := AndroidTimerNotificationReq{
	// 	AppID:            configs.Env.OneSignalAppID,
	// 	TargetChannel:    "push",
	// 	ContentAvailable: true,
	// 	Priority:         10,
	// 	Data: AndroidTimerNotificationData{
	// 		Type:    "update",
	// 		Subject: subjectObj.Name,
	// 		StartAt: timerObj.StartAt.Format(time.RFC3339),
	// 		Content: timerObj.Content,
	// 		Color:   subjectObj.Color,
	// 	},
	// 	IncludeSubsriptionIDs: androidSubs,
	// }

	// err = AndroidTimerNotification(data)
	// if err != nil {
	// 	logger.Error(err)
	// }
}

func SendTimerDelete(userID string) {

	// onesignalUser, err := GetOneSignalUser(userID)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }

	// androidSubs, _ := SplitAndroidIOSSubscriptions(onesignalUser.Subscriptions)

	// data := AndroidTimerNotificationReq{
	// 	AppID:            configs.Env.OneSignalAppID,
	// 	TargetChannel:    "push",
	// 	ContentAvailable: true,
	// 	Priority:         10,
	// 	Data: AndroidTimerNotificationData{
	// 		Type: "delete",
	// 	},
	// 	IncludeSubsriptionIDs: androidSubs,
	// }

	// err = AndroidTimerNotification(data)
	// if err != nil {
	// 	logger.Error(err)
	// }
}
