package analyzer

type AnalysisResult struct {
	URL               string            `json:"url"`
	HTMLVersion       string            `json:"htmlVersion"`
	PageTitle         string            `json:"pageTitle"`
	Headings          map[string]int    `json:"headings"`
	InternalLinks     int               `json:"internalLinks"`
	ExternalLinks     int               `json:"externalLinks"`
	InaccessibleLinks int               `json:"inaccessibleLinks"`
	LoginFormExists   bool              `json:"loginFormExists"`
}
