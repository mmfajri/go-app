package requests

import "time"

type BaseRequest struct {
	IsDelete    bool      `json:"is_deleted"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
	DeletedTime time.Time `json:"deleted_time"`
}
