package restaurants

import (
	"strconv"
	"sync"
)

func letsmeat(c chan string, wg *sync.WaitGroup) {
	lm := Restarant{}
	p := strconv.Itoa(2 + day)
	lm.xPath = "//*[@id=\"post-335\"]/div/div/p[" + p + "]/text()"

	lm.url = "https://www.letsmeat.cz/poledni-menu/"
	lm.GetMenu()
	lm.result = "LetsMeat" + "\n"
	for _, v := range lm.response {
		lm.result += v + "\n"
	}
	c <- lm.result
	wg.Done()
}
