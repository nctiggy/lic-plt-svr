package main

import (
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type License struct {
	Id    int    `json:"id"`
	State string `json:"state"`
	Found bool   `json:"found"`
	Img   string `json:"img"`
}

type Licenses struct {
	licenses []License
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (l *Licenses) init() *Licenses {
	data, err := ioutil.ReadFile("plates.yaml")
	check(err)
	err = yaml.Unmarshal(data, &l.licenses)
	check(err)
	return l
}

func (l *Licenses) reset() *Licenses {
	l.init()
	return l
}

func (l *Licenses) toggleFound(id int) *Licenses {
	for i, lic := range l.licenses {
		if lic.Id == id {
			lic.Found = !lic.Found
			l.licenses[i] = lic
		}
	}
	return l
}

func (l *Licenses) apiToggleFound(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	check(err)
	l.toggleFound(id)
	c.Status(204)
}

func (l *Licenses) getCurrentStatus(c *gin.Context) {
	//r, err := json.Marshal(l.licenses)
	//check(err)
	//c.JSON(200, string(r))
	c.JSON(200, l.licenses)
}

func main() {
	l := Licenses{}
	l.init()
	r := gin.Default()
	r.GET("/", l.getCurrentStatus)
	//r.POST("/reset", l.resetGame)
	r.PUT("/:id", l.apiToggleFound)
	r.Run() // listen and serve on 0.0.0.0:8080
}
