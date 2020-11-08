package main

// Category 一言数据分类
type Category struct {
	CreatedAt string `json:"created_at"`
	Desc      string `json:"desc"`
	ID        int64  `json:"id"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	UpdatedAt string `json:"updated_at"`
}

// Sentence 单条一言数据
type Sentence struct {
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Creator    string `json:"creator"`
	CreatorUID int64  `json:"creator_uid"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Hitokoto   string `json:"hitokoto"`
	ID         int64  `json:"id"`
	Length     int64  `json:"length"`
	Reviewer   int64  `json:"reviewer"`
	Type       string `json:"type"`
	UUID       string `json:"uuid"`
}
