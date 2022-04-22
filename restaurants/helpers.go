package restaurants

import (
	"regexp"
	"time"

	"github.com/gocolly/colly"
)

type Restarant struct {
	// name     string
	url      string
	response []string
	xPath    string
	result   string
}

var day = int(time.Now().Weekday()) - 1
var c = colly.NewCollector()
var whitespaces = regexp.MustCompile(`(\s\s\s+)`)
var numlet = regexp.MustCompile(`[ a-z][0-9]`)

func (r *Restarant) GetMenu() {
	c.OnXML(r.xPath, func(x *colly.XMLElement) {
		r.response = append(r.response, x.Text)
	})
	c.Visit(r.url)
}

func (r *Restarant) GetOneDayMenu() {
	c.OnXML(r.xPath, func(x *colly.XMLElement) {
		r.response = append(r.response, x.Text)
	})
	c.Visit(r.url)
}
