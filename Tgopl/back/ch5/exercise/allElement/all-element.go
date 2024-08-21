package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	outline(os.Args[1])
}

func outline(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	doc, _ := html.Parse(resp.Body)

	doc2 := ElementByID(doc, "navs", startElement2)
	forEachNode(doc2, startElement, endElement)
	resp.Body.Close()

	s := expand("footest", expand2)
	fmt.Println(s)
	return "", nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attr := ""
		for _, a := range n.Attr {
			attr += " " + a.Key + "=" + "\"" + a.Val + "\""
		}
		fmt.Printf("%*s%s%s", depth*2, "", n.Data, attr)
		depth++
	}
	if n.Type == html.ElementNode && n.FirstChild == nil && n.Data != "script" {
		fmt.Printf("/>\n")
	} else if n.Type == html.ElementNode {
		fmt.Printf(">\n")
	}

	if n.Type == html.TextNode {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.TextNode {
		depth--
		fmt.Printf("\n")
		return
	}
	if n.Type == html.ElementNode {
		depth--

		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
func ElementByID(n *html.Node, idStr string, pre func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if pre(n, idStr) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(c, idStr, pre)
	}
	return n
}

func startElement2(n *html.Node, idStr string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == idStr {
				fmt.Println(a.Val)
				break
				return true
			}
		}
	}
	return false
}

func expand(s string, f func(string) string) string {
	str := f("foo")
	s = strings.Replace(s, "foo", str, -1)
	return s
}
func expand2(s string) string {
	return s + "-test"
}
