package internal

type DynamicContent struct {
	PageTitle   string
	PageContent PageContent
}

type PageContent struct {
	Header             string
	Date               string
	Description        string
	AchievementsTitle  string
	Achievements       []Achievement
	LatestUpdatesTitle string
	LatestBlogs        []LatestBlog
	LatestButtonTitle  string
	SenderName         string
	SenderOccupation   string
	CurrentYear        string
}

type LatestBlog struct {
	LatestBlogImage       string
	LatestBlogTitle       string
	LatestBlogDate        string
	LatestBlogDescription string
}

type Achievement struct {
	AchievementImage       string
	AchievementTitle       string
	AchievementDescription string
	AchievementButtonText  string
}
