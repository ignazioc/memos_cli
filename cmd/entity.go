package cmd

type Property struct {
	Tags               []string `json:"tags"`
	HasLink            bool     `json:"hasLink"`
	HasTaskList        bool     `json:"hasTaskList"`
	HasCode            bool     `json:"hasCode"`
	HasIncompleteTasks bool     `json:"hasIncompleteTasks"`
}

type Memo struct {
	Name        string   `json:"name"`
	UID         string   `json:"uid"`
	RowStatus   string   `json:"rowStatus"`
	Creator     string   `json:"creator"`
	CreateTime  string   `json:"createTime"`
	UpdateTime  string   `json:"updateTime"`
	DisplayTime string   `json:"displayTime"`
	Content     string   `json:"content"`
	Nodes       []string `json:"nodes"`
	Visibility  string   `json:"visibility"`
	Tags        []string `json:"tags"`
	Pinned      bool     `json:"pinned"`
	Resources   []string `json:"resources"`
	Relations   []string `json:"relations"`
	Reactions   []string `json:"reactions"`
	Property    Property `json:"property"`
	Snippet     string   `json:"snippet"`
}

type MemosResponse struct {
	Memos         []Memo `json:"memos"`
	NextPageToken string `json:"nextPageToken"`
}
