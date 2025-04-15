package v1

import (
	"billionmail-core/utility/types/api_v1"
	"github.com/gogf/gf/v2/frame/g"
)

type ContactGroup struct {
	Id          int    `json:"id"          dc:"Group ID"`
	Name        string `json:"name"        dc:"Group Name"`
	Description string `json:"description" dc:"Description"`
	CreateTime  int    `json:"create_time" dc:"Create Time"`
	UpdateTime  int    `json:"update_time" dc:"Update Time"`
}

type ContactGroupInfo struct {
	Id               int    `json:"id"          dc:"Group ID"`
	Name             string `json:"name"        dc:"Group Name"`
	Description      string `json:"description" dc:"Description"`
	CreateTime       int    `json:"create_time" dc:"Create Time"`
	UpdateTime       int    `json:"update_time" dc:"Update Time"`
	TotalCount       int    `json:"total_count" dc:"Total Contacts"`
	ActiveCount      int    `json:"active_count" dc:"Subscribed Contacts"`
	UnsubscribeCount int    `json:"unsubscribe_count" dc:"Unsubscribed Contacts"`
}

type Contact struct {
	Id         int    `json:"id"          dc:"Contact ID"`
	Email      string `json:"email"       dc:"Email Address"`
	GroupId    int    `json:"group_id"    dc:"Group ID"`
	Active     int    `json:"active"      dc:"Status(1:Subscribed 0:Unsubscribed)"`
	TaskId     int    `json:"task_id"     dc:"Bulk Mail Task ID"`
	CreateTime int    `json:"create_time" dc:"Create Time"`
}

type CreateGroupReq struct {
	g.Meta        `path:"/contact/group/create" method:"post" tags:"Contact" summary:"Create new contact group"`
	Authorization string   `json:"authorization" dc:"Authorization" in:"header"`
	Name          string   `json:"name" v:"required" dc:"Group Name"`
	Description   string   `json:"description" dc:"Description"`
	Contacts      []string `json:"contacts" dc:"Email List"`
	ImportFile    string   `json:"import_file" dc:"Import File"`
	FileType      string   `json:"file_type" dc:"File Type(csv/excel/txt)"`
	CreateType    int      `json:"create_type" v:"required|in:1,2,3" dc:"Create Type(1:Group Only 2:Group with Contacts 3:Group with Import)"`
}

type CreateGroupRes struct {
	api_v1.StandardRes
}

type ImportContactsReq struct {
	g.Meta        `path:"/contact/group/import" method:"post" tags:"Contact" summary:"Import contacts"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" dc:"Group ID(Not required for new group)"`
	Name          string `json:"name" dc:"Group Name(Required for new group)"`
	Description   string `json:"description" dc:"Description(Required for new group)"`
	File          string `json:"file" dc:"Contact File"`
	FileType      string `json:"file_type" v:"in:csv,excel,txt" dc:"File Type"`
	ImportType    int    `json:"import_type" v:"required|in:1,2,3" dc:"Import Type(1:New Group 2:Existing Group 3:Group Only)"`
	DefaultActive int    `json:"default_active" v:"in:0,1" dc:"Default Status(1:Subscribed 0:Unsubscribed) Default:1"`
}

type ImportContactsRes struct {
	api_v1.StandardRes
}

type ExportContactsReq struct {
	g.Meta          `path:"/contact/export" method:"post" tags:"Contact" summary:"Export contacts"`
	Authorization   string `json:"authorization" dc:"Authorization" in:"header"`
	GroupIds        []int  `json:"group_ids" v:"required" dc:"Group IDs"`
	ExportType      int    `json:"export_type" v:"required|in:1,2" dc:"Export Type(1:Merged 2:Separate)"`
	Format          string `json:"format" v:"required|in:csv,excel,txt" dc:"Export Format"`
	IncludeInactive bool   `json:"include_unsubscribe" dc:"Include Unsubscribed(Default:true)"`
}

type ExportContactsRes struct {
	api_v1.StandardRes
	Data struct {
		FileUrl string `json:"file_url" dc:"Download URL"`
	} `json:"data"`
}

type DeleteGroupReq struct {
	g.Meta        `path:"/contact/group/delete" method:"post" tags:"Contact" summary:"Delete contact group"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupIds      []int  `json:"group_ids" v:"required" dc:"Group IDs"`
	DeleteType    int    `json:"delete_type" v:"required|in:1,2" dc:"Delete Type(1:With Contacts 2:Keep Contacts)"`
}

type DeleteGroupRes struct {
	api_v1.StandardRes
	Data struct {
		SuccessCount int `json:"success_count" dc:"Successfully Deleted Groups"`
		FailedCount  int `json:"failed_count" dc:"Failed Deleted Groups"`
	} `json:"data"`
}

type UpdateGroupReq struct {
	g.Meta        `path:"/contact/group/update" method:"post" tags:"Contact" summary:"Update contact group"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	GroupId       int    `json:"group_id" v:"required" dc:"Group ID"`
	Name          string `json:"name" dc:"Group Name"`
	Description   string `json:"description" dc:"Description"`
}

type UpdateGroupRes struct {
	api_v1.StandardRes
}

type ListGroupsReq struct {
	g.Meta        `path:"/contact/group/list" method:"get" tags:"Contact" summary:"List contact groups"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `json:"page" v:"required|min:1" dc:"Page Number"`
	PageSize      int    `json:"page_size" v:"required|min:1" dc:"Page Size"`
	Keyword       string `json:"keyword" dc:"Search Keyword"`
}

type ListGroupsRes struct {
	api_v1.StandardRes
	Data struct {
		Total int                 `json:"total" dc:"Total Count"`
		List  []*ContactGroupInfo `json:"list" dc:"Group List"`
	} `json:"data"`
}

type MergeGroupsReq struct {
	g.Meta         `path:"/contact/group/merge" method:"post" tags:"Contact" summary:"Merge contact groups"`
	Authorization  string `json:"authorization" dc:"Authorization" in:"header"`
	SourceGroupIds []int  `json:"source_group_ids" v:"required" dc:"Source Group IDs"`
	TargetGroupId  int    `json:"target_group_id" v:"required" dc:"Target Group ID"`
}

type MergeGroupsRes struct {
	api_v1.StandardRes
	Data struct {
		TotalMerged int `json:"total_merged" dc:"Total Merged Contacts"`
		Duplicate   int `json:"duplicate" dc:"Duplicate Contacts"`
	} `json:"data"`
}

type ListContactsReq struct {
	g.Meta        `path:"/contact/list" method:"get" tags:"Contact" summary:"List all contacts"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Page          int    `json:"page" v:"required|min:1" dc:"Page Number"`
	PageSize      int    `json:"page_size" v:"required|min:1" dc:"Page Size"`
	GroupId       int    `json:"group_id" dc:"Group ID(Optional)"`
	Keyword       string `json:"keyword" dc:"Search Email"`
	Status        int    `json:"status" dc:"Status(1:Subscribed 0:Unsubscribed -1:All)" default:"1"`
}

type ListContactsRes struct {
	api_v1.StandardRes
	Data struct {
		Total int        `json:"total" dc:"Total Count"`
		List  []*Contact `json:"list" dc:"Contact List"`
	} `json:"data"`
}

type ListContactsGroupsReq struct {
	g.Meta        `path:"/contact/group/all" method:"get" tags:"Contact" summary:"List all contact groups"`
	Authorization string `json:"authorization" dc:"Authorization" in:"header"`
	Keyword       string `json:"keyword" dc:"Search Keyword"`
}

type ListContactsGroupsRes struct {
	api_v1.StandardRes
	Data struct {
		List []*ContactGroup `json:"list" dc:"Group List"`
	} `json:"data"`
}
