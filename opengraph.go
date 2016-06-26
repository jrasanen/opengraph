package main

import (
	"net/http"

	"github.com/lestrrat/go-libxml2"
	"github.com/lestrrat/go-libxml2/xpath"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type OpenGraphTag struct {
	property string
	content  string
}

type TagList struct {
	tags []OpenGraphTag
}

// @TODO: map by name to avoid iteration
func (tl TagList) GetTagsByName(name string) []OpenGraphTag {
	var list []OpenGraphTag
	for _, tag := range tl.tags {
		if tag.property == name {
			list = append(list, tag)
		}
	}
	return list
}

func GetTags(url string) TagList {
	res, err := http.Get(url)
	if err != nil {
		panic("failed to get golang.org: " + err.Error())
	}

	doc, err := libxml2.ParseHTMLReader(res.Body)
	if err != nil {
		panic("failed to parse HTML: " + err.Error())
	}
	defer doc.Free()

	nodeList := xpath.NodeList(doc.Find(`//*/meta[starts-with(@property, 'og:')]`))

	if err != nil {
		panic("failed to evaluate xpath: " + err.Error())
	}

	var list TagList

	for i := 0; i < len(nodeList); i++ {
		property := xpath.String(nodeList[i].Find("@property"))
		content := xpath.String(nodeList[i].Find("@content"))
		list.tags = append(list.tags, OpenGraphTag{
			property: property,
			content:  content,
		})
	}

	return list
}
