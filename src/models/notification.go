package models 

type Notification struct{
	Type string `json:"type" binding:"required"`
	Recipient string `json:"ricipient" binding:"required"`
	From string `json:"from" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Body string `json:"body" binding:"required"`
}
