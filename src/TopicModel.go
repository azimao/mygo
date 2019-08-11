package src

type Topic struct {
	TopicId    int    `json:"id"`
	TopicTitle string `json:"title"`
}

func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}

type TopicQuery struct {
	Username string `json:"username"`
	Page     int    `json:"page"`
	PageSize int    `json:"pagesize"`
}
