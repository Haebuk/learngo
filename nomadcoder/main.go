package main

import (
	"os"
	"strings"

	"github.com/Haebuk/learngo/nomadcoder/scrapper"
	"github.com/labstack/echo/v4"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScape(c echo.Context) error {
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName , "job.csv")
} 

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScape)
	e.Logger.Fatal(e.Start(":1323"))
	
	
}