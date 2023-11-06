package src

// User ...内存对齐 省空间
type User struct {
	ID            int
	NftNum        int
	FtNum         int
	TicketNum     int
	NickName      string
	Bios          string
	Avatar        string
	Address       string
	TwitterID     string
	TwitterAvatar string
	TwitterName   string
	LastLogin     int
	registerTime  int
}
type Cave struct {
}
type Secret struct {
	ID        int      `json:"ID"` //autoIncrement
	AuthorID  int      `json:"author_id"`
	Content   string   `json:"content"`
	Image     []string `json:"image"`
	Timestamp int64    `json:"timestamp"`
	Views     int64    `json:"views"`
}
type Viewable struct {
	SecretID int
	Users    []int
}

type SiteFd struct {
	Base   int
	Friend []int
}
type SiteFd_table struct {
	Base   int
	Friend int
	Status int
}
