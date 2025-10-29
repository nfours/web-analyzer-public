package service

import "web-analyzer/internal/analyzer"

type AnalyzerService interface {
	AnalyzeURL(url string) (*analyzer.AnalysisResult, error)
}

type analyzerService struct {
	analyzer analyzer.Analyzer
}

func NewAnalyzerService(an analyzer.Analyzer) AnalyzerService {
	return &analyzerService{analyzer: an}
}

func (s *analyzerService) AnalyzeURL(url string) (*analyzer.AnalysisResult, error) {
	return s.analyzer.Analyze(url)
}
