package jobtype

type DeferNotifyUserPayload struct {
	Email    string
	UserName string
	UserId   int64
}

type SchedulerWriteBackMysqlPostPayload struct {
	PostId int64
}
