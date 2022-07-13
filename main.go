package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Chapter struct {
	Id          int          `json:"id"`
	ChapterName string       `json:"chapter_name"`
	Invocations []Invocation `json:"invocations"`
}

type Invocation struct {
	Id        int    `json:"id"`
	AudioPath string `json:"audio"`
	Arabic    string `json:"arabic"`
	Latin     string `json:"latin"`
	Albanian  string `json:"albanian"`
	Reference string `json:"reference"`
}

func readData() []Chapter {
	content, err := ioutil.ReadFile("./invocations.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload []Chapter
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func getChapter(c *gin.Context) {
	data := readData()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	var chapter Chapter
	for _, item := range data {
		if item.Id == id {
			chapter = item
		}
	}

	if chapter.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Chapter not found."})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

func getChapters(c *gin.Context) {
	data := readData()

	var list []Chapter
	for _, chapter := range data {
		list = append(list, Chapter{Id: chapter.Id, ChapterName: chapter.ChapterName})
	}

	c.JSON(http.StatusOK, list)
}

func getChaptersWithInvocations(c *gin.Context) {
	data := readData()
	c.JSON(http.StatusOK, data)
}

func getChapterInvocations(c *gin.Context) {
	data := readData()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// ... handle error
		panic(err)
	}

	var invocations []Invocation
	for _, item := range data {
		if item.Id == id {
			invocations = item.Invocations
		}
	}
	c.JSON(http.StatusOK, invocations)
}

func getSingleInvocation(c *gin.Context) {
	data := readData()
	id, _ := strconv.Atoi(c.Param("id"))
	invocationId, _ := strconv.Atoi(c.Param("invocationId"))

	var invocations []Invocation
	for _, item := range data {
		if item.Id == id {
			invocations = item.Invocations
		}
	}

	if len(invocations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Chapter not found."})
		return
	}

	var invocation Invocation
	for _, item := range invocations {
		if item.Id == invocationId {
			invocation = item
		}
	}

	if (Invocation{}) == invocation {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invocation not found."})
		return
	}

	c.JSON(http.StatusOK, invocation)
}

func main() {
	router := gin.Default()
	router.Static("/audios", "./audios")

	v1 := router.Group("/api/v1")
	{
		chapters := v1.Group("/chapters")
		{
			chapters.GET("", getChapters)
			chapters.GET(":id", getChapter)
			chapters.GET(":id/invocations", getChapterInvocations)
			chapters.GET(":id/invocations/:invocationId", getSingleInvocation)
		}

		invocations := v1.Group("/invocations")
		{
			invocations.GET("", getChaptersWithInvocations)
		}
	}

	router.Run(":8888")
}
