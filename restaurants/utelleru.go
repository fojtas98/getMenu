package restaurants

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func uTelleru(c chan string, wg *sync.WaitGroup) {
	r := Restarant{}
	r.xPath = "/html/body/section/div[" + strconv.Itoa(day) + "]/div/div"
	r.url = "https://www.utelleru.cz/obedy/"
	r.GetMenu()
	r.result = "U Telleru\n"

	n := regexp.MustCompile(`(([0-9]+)(,|\n))|([0-9]\.) `)

	for _, v := range r.response {
		v = n.ReplaceAllString(v, "")
		v = strings.ToLower(v)
		v = whitespaces.ReplaceAllString(v, "")
		v = strings.Replace(v, "kč", "kč\n", -1)
		v = numlet.ReplaceAllStringFunc(v, splitNumberAndLetter)
		r.result += v
	}
	c <- r.result
	wg.Done()
}
