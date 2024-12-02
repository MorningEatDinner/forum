package jobtype

// 延迟发送邮件通知更新用户信息
const DeferEmailNotifyJob = "defer:User:Notify"

// 周期任务
// 周期生成热点帖子信息发送给用户邮箱
const ScheduleHotPostPushing = "schedule:Post:Pushing"

// 周期删除在窗口之外的数据
const ScheduleDeletePost = "schedule:Post:Delete"
