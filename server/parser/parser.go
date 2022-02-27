package parser

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func GetFirstTBody(n *html.Node) *html.Node {
	if n == nil {
		return nil
	}

	if n.Type == html.ElementNode && strings.ToLower(n.Data) == "tbody" {
		return n
	}

	if tbody := GetFirstTBody(n.FirstChild); tbody != nil {
		return tbody
	}

	if tbody := GetFirstTBody(n.NextSibling); tbody != nil {
		return tbody
	}

	return nil
}

func GetNextTr(n *html.Node) *html.Node {
	return getNextElementNode(n, "tr")
}

func GetNextTd(n *html.Node) *html.Node {
	return getNextElementNode(n, "td")
}

func GetInnerText(n *html.Node) string {
	if n == nil {
		return ""
	}

	if n.Type == html.TextNode {
		return n.Data
	}

	result := ""
	for el := n.FirstChild; el != nil; el = el.NextSibling {
		result += GetInnerText(el)
	}
	return result
}

func ParseCommaFloat(s string) (float64, error) {
	withoutDots := strings.ReplaceAll(s, ".", "")
	replacedCommas := strings.ReplaceAll(withoutDots, ",", ".")
	return strconv.ParseFloat(replacedCommas, 64)
}

func getNextElementNode(n *html.Node, typeOfNode string) *html.Node {
	for el := n; el != nil; el = el.NextSibling {
		if el.Type == html.ElementNode && strings.ToLower(el.Data) == typeOfNode {
			return el
		}
	}
	return nil
}
