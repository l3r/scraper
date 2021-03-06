package scraper

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var templateRe = regexp.MustCompile(`\{\{\s*(\w+)\s*(:(\w+))?\s*\}\}`)

func template(str string, vars url.Values) (out string, err error) {
	out = templateRe.ReplaceAllStringFunc(str, func(key string) string {
		m := templateRe.FindStringSubmatch(key)
		k := m[1]
		v := vars.Get(k)
		if v == "" {
			if m[3] != "" {
				v = m[3]
			} else {
				err = errors.New("Missing param: " + k)
			}
		}
		return v
	})
	return
}

func checkSelector(s string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	doc, _ := goquery.NewDocumentFromReader(bytes.NewBufferString(`<html>
		<body>
			<h3>foo bar</h3>
		</body>
	</html>`))
	doc.Find(s)
	return
}
