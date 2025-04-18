package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type EmailTask struct {
	Id             int    `json:"id"              dc:"task id"`
	TaskName       string `json:"task_name"       dc:"task name"`
	Addresser      string `json:"addresser"       dc:"addresser"`
	Subject        string `json:"subject"         dc:"subject"`
	FullName       string `json:"full_name"       dc:"full name"`
	RecipientCount int    `json:"recipient_count" dc:"recipient count"`
	TaskProcess    int    `json:"task_process"    dc:"task process"`
	Pause          int    `json:"pause"           dc:"pause"`
	TemplateId     int    `json:"template_id"     dc:"template id"`
	IsRecord       int    `json:"is_record"       dc:"is record"`
	Unsubscribe    int    `json:"unsubscribe"     dc:"unsubscribe"`
	Threads        int    `json:"threads"         dc:"threads"`
	Etypes         string `json:"etypes"          dc:"etypes"`
	TrackOpen      int    `json:"track_open"      dc:"track open"`
	TrackClick     int    `json:"track_click"     dc:"track click"`
	StartTime      int    `json:"start_time"      dc:"start time"`
	CreateTime     int    `json:"create_time"     dc:"create time"`
	UpdateTime     int    `json:"update_time"     dc:"update time"`
	Remark         string `json:"remark"          dc:"remark"`
	Active         int    `json:"active"          dc:"status"`
	AddType        int    `json:"add_type"        dc:"add type"`
}

type GroupInfo struct {
	Id          int    `json:"id"    dc:"group id"`
	Name        string `json:"name"  dc:"group name"`
	Description string `json:"description" dc:"group description"`
	Count       int    `json:"count" dc:"contact count"`
}

type TaskDetail struct {
	EmailTask
	Groups       []*GroupInfo `json:"groups"         dc:"groups"`
	SentCount    int          `json:"sent_count"     dc:"sent count"`
	UnsentCount  int          `json:"unsent_count"   dc:"unsent count"`
	Progress     int          `json:"progress"      dc:"progress"`
	TemplateName string       `json:"template_name"  dc:"template name"`
	SuccessCount int          `json:"success_count"  dc:"success count"`
	ErrorCount   int          `json:"error_count"    dc:"error count"`
}

type ListTasksReq struct {
	g.Meta        `path:"/batch_mail/task/list" method:"get" tags:"BatchMail" summary:"get task list"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `json:"page" v:"required|min:1" dc:"page"`
	PageSize      int    `json:"page_size" v:"required|min:1" dc:"page size"`
	Keyword       string `json:"keyword" dc:"search keyword"`
	Status        int    `json:"status" dc:"task status" default:"-1"`
}

type ListTasksRes struct {
	api_v1.StandardRes
	Data struct {
		Total int           `json:"total" dc:"total"`
		List  []*TaskDetail `json:"list"  dc:"task list"`
	} `json:"data"`
}

type DeleteTaskReq struct {
	g.Meta        `path:"/batch_mail/task/delete" method:"post" tags:"BatchMail" summary:"delete task"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Id            int    `json:"id" v:"required" dc:"task id"`
}

type DeleteTaskRes struct {
	api_v1.StandardRes
}

type CreateTaskReq struct {
	g.Meta        `path:"/batch_mail/task/create" method:"post" tags:"BatchMail" summary:"create task"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Addresser     string `json:"addresser" v:"required" dc:"addresser"`
	Subject       string `json:"subject" v:"required" dc:"subject"`
	FullName      string `json:"full_name" dc:"full name"`
	TemplateId    int    `json:"template_id" v:"required" dc:"template id"`
	GroupIds      []int  `json:"group_ids" v:"required" dc:"group ids"`
	IsRecord      int    `json:"is_record" v:"in:0,1" dc:"is record" default:"1"`
	Unsubscribe   int    `json:"unsubscribe" v:"in:0,1" dc:"unsubscribe" default:"1"`
	Threads       int    `json:"threads" v:"min:1" dc:"threads" default:"1"`
	TrackOpen     int    `json:"track_open" v:"in:0,1" dc:"track open" default:"1"`
	TrackClick    int    `json:"track_click" v:"in:0,1" dc:"track click" default:"1"`
	StartTime     int    `json:"start_time" v:"required" dc:"start time"`
	Remark        string `json:"remark" dc:"remark"`
}

type CreateTaskRes struct {
	api_v1.StandardRes
	Data struct {
		Id int `json:"id" dc:"task id"`
	} `json:"data"`
}
