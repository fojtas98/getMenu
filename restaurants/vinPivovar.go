package restaurants

import (
	"strings"
	"sync"
	"time"
)

func vinPivovar(c chan string, wg *sync.WaitGroup) {
	vin := Restarant{}
	vin.xPath = "//*[@id=\"block-system-main\"]/div/div/div/div/div[2]/div"
	vin.url = "https://www.vinohradskypivovar.cz/menu/" + time.Now().Format("2006-01-02")
	vin.GetOneDayMenu()
	vin.result = "Vinohradsky pivovar" + "\n"

	for _, v := range vin.response {
		v = strings.Replace(v, "Specialita týdne", "", -1)
		v = whitespaces.ReplaceAllString(v, "")
		v = numlet.ReplaceAllStringFunc(v, splitNumberAndLetter)
		v = strings.Replace(v, "Kč", "Kč\n", -1)
		if len(v) > 0 {
			vin.result += v
		}
	}
	c <- vin.result
	wg.Done()
}

func splitNumberAndLetter(s string) string {
	t := strings.Split(s, "")
	return t[0] + " " + t[1]
}
