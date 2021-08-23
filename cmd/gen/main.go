package main

import (
	"fmt"
	"os"

	"github.com/jwhittle933/profile/internal/badges"
	"github.com/jwhittle933/profile/internal/fs"
	"github.com/jwhittle933/profile/internal/markdown"
)

func main() {
	readme, err := fs.Open("README.md")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	defer readme.Close()

	markdown.H1(readme, "Jonathan Whittle")
	markdown.H3(readme, "Senior software engineer out of Louisville, KY")
	markdown.Anchor(
		readme,
		"https://www.linkedin.com/in/itsthewhittlethings",
		markdown.Img(
			badges.New("LinkedIn", badges.None, badges.None, "social", "linkedin"),
			"linkedin badge",
		),
	)
	markdown.Anchor(
		readme,
		"https://twitter.com/JonathanWhittle",
		markdown.Img(
			badges.New("Twitter", badges.None, badges.None, "social", "twitter"),
			"twitter badge",
		),
	)
	markdown.HR(readme)
	markdown.NewLine(readme)

	markdown.H3(readme, "Education")
	markdown.LI(readme, "B.A. Philosophy/Religion, Clemson University")
	markdown.LI(readme, "M.Div Ancient Near Eastern Language and Linguistics, The Southern Baptist Theological Seminary")
	markdown.LI(
		readme,
		fmt.Sprintf(
			"Th.M Hebrew Bible/Septuagint, The Southern Baptist Theological Seminary, %s: %s",
			markdown.Bold("Thesis"),
			markdown.Italics("Double Renderings in Greek Proverbs 1-9"),
		),
	)
	markdown.NewLine(readme)

	markdown.H3(readme, "Hobbies")
	markdown.LI(readme, "Musician (20+ years)")
	markdown.LI(readme, "Personal Trainer/Body Building (10+ years)")
	markdown.LI(readme, "Language/Linguistics (Classical Hebrew, Classical/Hellenistic Greek, Latin, Aramaic, Syriac, Ugaritic, German, French)")
	markdown.NewLine(readme)

	markdown.H3(readme, "Technologies")
	markdown.Link(readme, "Go", badges.New("Go", "Expert", badges.ColorPrimary, "", "Go"))
	markdown.Link(readme, "Rust", badges.New("Rust", "Enthusiast", badges.ColorPrimary, "", "rust"))
	markdown.Link(readme, "Elixir", badges.New("Elixir", "Enthusiast", badges.ColorPrimary, "", "elixir"))
	markdown.Link(readme, "Clojure", badges.New("Clojure", "Enthusiast", badges.ColorPrimary, "", "clojure"))
	markdown.Link(readme, "JavaScript", badges.New("JavaScript", "Expert", badges.ColorPrimary, "", "javascript"))
	markdown.BR(readme)
	markdown.Link(readme, "TypeScript", badges.New(badges.None, "TypeScript", badges.ColorSecondary, "", "typescript"))
	markdown.Link(readme, "Node", badges.New(badges.None, "Nodejs", badges.ColorSecondary, "", "node.js"))
	markdown.Link(readme, "React", badges.New(badges.None, "React", badges.ColorSecondary, "", "react"))
	markdown.Link(readme, "Vue", badges.New(badges.None, "Vue", badges.ColorSecondary, "", "vue.js"))
	markdown.Link(readme, "Python", badges.New(badges.None, "Python", badges.ColorSecondary, "", "python"))
	markdown.Link(readme, "PHP", badges.New(badges.None, "PHP", badges.ColorSecondary, "", "php"))
	markdown.Link(readme, "Android", badges.New(badges.None, "Android", badges.ColorSecondary, "", "android"))
	markdown.Link(readme, "Kotlin", badges.New(badges.None, "Kotlin", badges.ColorSecondary, "", "Kotlin"))
	markdown.Link(readme, "IOS", badges.New(badges.None, "iOS", badges.ColorSecondary, "", "ios"))
	markdown.Link(readme, "Swift", badges.New(badges.None, "Swift", badges.ColorSecondary, "", "swift"))
	markdown.Link(readme, ".NET", badges.New(badges.None, ".NET", badges.ColorSecondary, "", ".net"))
	markdown.Link(readme, "C#", badges.New(badges.None, "Csharp", badges.ColorSecondary, "", "csharp"))
	markdown.Link(readme, "C", badges.New(badges.None, "C", badges.ColorSecondary, "", "c"))
	markdown.BR(readme)

	markdown.Link(readme, "Postgres", badges.New(badges.None, "Postgres", badges.ColorSecondary, "", "postgresql"))
	markdown.Link(readme, "MySQL", badges.New(badges.None, "MySQL", badges.ColorSecondary, "", "mysql"))
	markdown.Link(readme, "MSSQL", badges.New(badges.None, "MSSQL", badges.ColorSecondary, "", "microsoftsqlserver"))
	markdown.Link(readme, "Mongo", badges.New(badges.None, "MongoDB", badges.ColorSecondary, "", "mongodb"))
	markdown.Link(readme, "ElasticSearch", badges.New(badges.None, "ElasticSearch", badges.ColorSecondary, "", "elasticsearch"))
	markdown.BR(readme)

	markdown.Link(readme, "AWS", badges.New(badges.None, "AWS", badges.ColorAWS, "", "amazonaws"))
	markdown.Link(readme, "Azure", badges.New(badges.None, "Azure", badges.ColorAzure, "", "microsoftazure"))
	markdown.Link(readme, "Heroku", badges.New(badges.None, "Heroku", badges.ColorHeroku, "", "heroku"))
	markdown.Link(readme, "CF", badges.New(badges.None, "CloudFoundry", badges.ColorSecondary, "", "cloudfoundry"))
	markdown.BR(readme)

	markdown.Link(readme, "Git", badges.New(badges.None, "Git", badges.ColorSecondary, "", "git"))
	markdown.Link(readme, "Docker", badges.New(badges.None, "Docker", badges.ColorSecondary, "", "docker"))
	markdown.Link(readme, "Kubernetes", badges.New(badges.None, "Kubernetes", badges.ColorSecondary, "", "kubernetes"))
	markdown.Link(readme, "Bash", badges.New(badges.None, "Bash", badges.ColorSecondary, "", "gnubash"))
	markdown.BR(readme)


	markdown.Link(readme, "GitHub Actions", badges.New(badges.None, "GitHubActions", badges.ColorSecondary, "", "githubactions"))
	markdown.Link(readme, "ADO", badges.New(badges.None, "AzureDevOps", badges.ColorSecondary, "", "azuredevops"))
	markdown.Link(readme, "Jenkins", badges.New(badges.None, "Jenkins", badges.ColorSecondary, "", "jenkins"))
	markdown.Link(readme, "TravisCI", badges.New(badges.None, "TravisCI", badges.ColorSecondary, "", "travisci"))
	markdown.BR(readme)

	markdown.Link(readme, "Spacemacs", badges.New(badges.None, "Spacemacs", badges.ColorSecondary, "", "spacemacs"))
	markdown.Link(readme, "Emacs", badges.New(badges.None, "Emacs", badges.ColorSecondary, "", "gnuemacs"))
	markdown.Link(readme, "Vim", badges.New(badges.None, "Vim", badges.ColorSecondary, "", "vim"))
	markdown.BR(readme)
}
