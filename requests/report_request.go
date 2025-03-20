package requests

type ReportRequest struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ReportByUserIdRequest struct {
	UserId uint `json:"user_id"`
}

type ReportById struct {
	IdReport uint `json:"id_report"`
}
