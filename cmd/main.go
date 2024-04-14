package main

import (
	"html/template"
	"os"

	"github.com/siddhant-vij/Dynamic-Newsletter-Tool/pkg/common"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*.gohtml"))
}

func main() {
	achievement1 := common.Achievement{
		AchievementImage:       "achievement-1.png",
		AchievementTitle:       "Infographics for December",
		AchievementDescription: "Take a look at the progress we've made over the past month. This is a valuable source of new experience.",
		AchievementButtonText:  "View report",
	}

	achievement2 := common.Achievement{
		AchievementImage:       "achievement-2.png",
		AchievementTitle:       "Employees of the month",
		AchievementDescription: "Find out the best employees and strive to reach the same heights and charge your motivation.",
		AchievementButtonText:  "View listings",
	}

	achievements := []common.Achievement{achievement1, achievement2}

	latestBlog1 := common.LatestBlog{
		LatestBlogImage:       "latest-1.jpg",
		LatestBlogTitle:       "Design that gains profit. An ultimate guide.",
		LatestBlogDate:        "10.04.23",
		LatestBlogDescription: "Find out the top design trends that will help skyrocket profits and stand out from the competition.",
	}

	latestBlog2 := common.LatestBlog{
		LatestBlogImage:       "latest-2.jpg",
		LatestBlogTitle:       "How much does good design cost?",
		LatestBlogDate:        "13.05.23",
		LatestBlogDescription: "We carried out our calculations in order to tell you the cost of good design.",
	}

	latestBlogs := []common.LatestBlog{latestBlog1, latestBlog2}

	pageContent := common.PageContent{
		Header:             "Monthly Digest",
		Date:               "January 01, 2023",
		Description:        "December has come to an end, so we have prepared a small digest that closes the last month of 2022.",
		AchievementsTitle:  "Our achievements",
		Achievements:       achievements,
		LatestUpdatesTitle: "Latest from blog",
		LatestBlogs:        latestBlogs,
		LatestButtonTitle:  "Read more",
		SenderName:         "Tammy Guerrero",
		SenderOccupation:   "Hiring Manager",
		CurrentYear:        "2023",
	}

	dynamicContent := common.DynamicContent{
		PageTitle:   "HTML Template",
		PageContent: pageContent,
	}

	file, err := os.Create("../public/index.html")
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(file, "index.gohtml", dynamicContent)
	if err != nil {
		panic(err)
	}
}
