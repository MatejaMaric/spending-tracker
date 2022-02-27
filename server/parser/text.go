package parser

import (
	"golang.org/x/net/html"
)

func ExtractText(n *html.Node) [][]string {
	tbody := GetFirstTBody(n)
	if tbody == nil {
		return nil
	}

	var table [][]string
	for tr := GetNextTr(tbody.FirstChild); tr != nil; tr = GetNextTr(tr.NextSibling) {
		var row []string
		for td := GetNextTd(tr.FirstChild); td != nil; td = GetNextTd(td.NextSibling) {
			row = append(row, GetInnerText(td))
		}
		table = append(table, row)
	}

	return table
}
