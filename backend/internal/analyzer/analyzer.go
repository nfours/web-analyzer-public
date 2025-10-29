package analyzer

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

type Analyzer interface {
	Analyze(url string) (*AnalysisResult, error)
}

type htmlAnalyzer struct{}

func NewHTMLAnalyzer() Analyzer {
	return &htmlAnalyzer{}
}

func (a *htmlAnalyzer) Analyze(targetURL string) (*AnalysisResult, error) {
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &AnalysisResult{
		URL:      targetURL,
		Headings: make(map[string]int),
	}

	result.HTMLVersion = detectHTMLVersion(doc)

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			switch n.Data {
			case "title":
				if n.FirstChild != nil {
					result.PageTitle = n.FirstChild.Data
				}
			case "h1", "h2", "h3", "h4", "h5", "h6":
				result.Headings[n.Data]++
			case "a":
				processLink(n, targetURL, result)
			case "input":
				for _, attr := range n.Attr {
					if attr.Key == "type" && attr.Val == "password" {
						result.LoginFormExists = true
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)
	return result, nil
}

func detectHTMLVersion(doc *html.Node) string {
	if doc.Type == html.DoctypeNode && strings.Contains(strings.ToLower(doc.Data), "html") {
		return "HTML5"
	}
	return "Unknown"
}

func processLink(n *html.Node, baseURL string, result *AnalysisResult) {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			link := attr.Val
			u, err := url.Parse(link)
			if err != nil {
				continue
			}
			if u.Host == "" || strings.Contains(u.Host, baseURL) {
				result.InternalLinks++
			} else {
				result.ExternalLinks++
			}
		}
	}
}
