package rule

import "strings"

type Rule struct {
	Name             string
	SegmentationRule string
	FilteringRules   string
	DivisionIndex    int
}

func NewRule(ruleName string, segmentationRule string, filteringRules string, divisionIndex int) *Rule {
	rule := &Rule{
		Name:             ruleName,
		SegmentationRule: segmentationRule,
		FilteringRules:   filteringRules,
		DivisionIndex:    divisionIndex,
	}
	return rule
}

func (r *Rule) Segmentation(content string) string {
	strs := strings.Split(strings.TrimSpace(content), r.SegmentationRule)
	if len(strs) == 0 || r.DivisionIndex >= len(strs) {
		return ""
	}
	return strings.TrimSpace(strs[r.DivisionIndex])
}

func (r *Rule) HasPrefix(content string) bool {
	if len(content) == 0 {
		return false
	}
	if strings.HasPrefix(content, r.FilteringRules) ||
		strings.HasPrefix(strings.ToLower(content), strings.ToLower(r.FilteringRules)) ||
		strings.HasPrefix(strings.ToUpper(content), strings.ToUpper(r.FilteringRules)) {
		return true
	}
	return false
}
