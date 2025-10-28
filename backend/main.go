package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "net/url"
    "strings"

    "golang.org/x/net/html"
)

type AnalysisResult struct {
    URL               string
    HTMLVersion       string
    PageTitle         string
    Headings          map[string]int
    InternalLinks     int
    ExternalLinks     int
    InaccessibleLinks int
    LoginFormExists   bool
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/analyze", analyzeHandler)
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("../index.html"))
    tmpl.Execute(w, nil)
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
    urlStr := r.FormValue("url")
    if urlStr == "" {
        http.Error(w, "URL is required", http.StatusBadRequest)
        return
    }

    result, err := analyzeURL(urlStr)
    if err != nil {
        http.Error(w, "Error analyzing URL: "+err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, result)
}

func analyzeURL(urlStr string) (*AnalysisResult, error) {
    res, err := http.Get(urlStr)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    doc, err := html.Parse(res.Body)
    if err != nil {
        return nil, err
    }

    result := &AnalysisResult{
        URL:      urlStr,
        Headings: make(map[string]int),
    }

    // HTML version (simplified check)
    if doc.Type == html.DoctypeNode && strings.Contains(strings.ToLower(doc.Data), "html") {
        result.HTMLVersion = "HTML5"
    } else {
        result.HTMLVersion = "Unknown"
    }

    // Traverse HTML
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode {
            switch n.Data {
            case "title":
                if n.FirstChild != nil {
                    result.PageTitle = n.FirstChild.Data
                }
            case "h1", "h2", "h3", "h4", "h5", "h6":
                result.Headings[n.Data]++
            case "a":
                href := ""
                for _, attr := range n.Attr {
                    if attr.Key == "href" {
                        href = attr.Val
                    }
                }
                if href != "" {
                    u, err := url.Parse(href)
                    if err == nil {
                        if u.Host == "" || strings.Contains(u.Host, urlStr) {
                            result.InternalLinks++
                        } else {
                            result.ExternalLinks++
                        }
                        //check accessibility
                        resp, err := http.Head(href)
                        if err != nil || resp.StatusCode >= 400 {
                            result.InaccessibleLinks++
                        }
                    }
                }
            case "form":
                for _, attr := range n.Attr {
                    if attr.Key == "type" && attr.Val == "login" {
                        result.LoginFormExists = true
                    }
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }

    f(doc)
    return result, nil
}
