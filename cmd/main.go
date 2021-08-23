package main

import (
	"fmt"
	"os"

	"github.com/jwhittle933/profile/internal/badges"
	"github.com/jwhittle933/profile/internal/markdown"
	"github.com/jwhittle933/profile/internal/readme"
)

func main() {
	rm, err := readme.Open("README.md")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		os.Exit(1)
	}
	defer rm.Close()

	markdown.H1(rm, "Jonathan Whittle")
	markdown.H3(rm, "Senior software engineer out of Louisville, KY")
	markdown.Anchor(
		rm,
		"https://www.linkedin.com/in/itsthewhittlethings",
		markdown.Img(
			badges.New("LinkedIn", badges.None, badges.None, "social", "linkedin"),
			"linkedin badge",
		),
	)
	markdown.Anchor(
		rm,
		"https://twitter.com/JonathanWhittle",
		markdown.Img(
			badges.New("Twitter", badges.None, badges.None, "social", "twitter"),
			"twitter badge",
		),
	)
	markdown.HR(rm)
	markdown.NewLine(rm)

	markdown.H3(rm, "Education")
	markdown.LI(rm, "B.A. Philosophy/Religion, Clemson University")
	markdown.LI(rm, "M.Div Ancient Near Eastern Language and Linguistics, The Southern Baptist Theological Seminary")
	markdown.LI(
		rm,
		fmt.Sprintf(
			"Th.M Hebrew Bible/Septuagint, The Southern Baptist Theological Seminary, %s: %s",
			markdown.Bold("Thesis"),
			markdown.Italics("Double Renderings in Greek Proverbs 1-9"),
		),
	)
	markdown.NewLine(rm)


	markdown.H3(rm, "Hobbies")
	markdown.LI(rm, "Musician (20+ years)")
	markdown.LI(rm, "Personal Trainer/Body Building (10+ years)")
	markdown.LI(rm, "Language/Linguistics (Classical Hebrew, Classical/Hellenistic Greek, Latin, Aramaic, Syriac, Ugaritic, German, French)")
	markdown.NewLine(rm)


	markdown.H3(rm, "Technologies")
	markdown.Link(rm, "Go", badges.New("Go", "Expert", badges.ColorPrimary, "", "Go"))
	markdown.Link(rm, "Rust", badges.New("Rust", "Enthusiast", badges.ColorPrimary, "", "rust"))
}
