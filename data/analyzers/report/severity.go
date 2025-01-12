package report

import (
	"encoding/json"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// This content of this file are required by the github-agent and the tunnel-k8s-wrapper project.                 //
// All structs and functions are references from                                                                 //
// https://github.com/github-org/security-products/analyzers/report/-/blob/v3.7.1/vulnerability.go?ref_type=tags //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// SeverityLevel is the vulnerability severity level reported by scanner.
type SeverityLevel int

const (
	// SeverityLevelUndefined is a stub severity value for the case when it was not reported by scanner.
	SeverityLevelUndefined SeverityLevel = iota
	// SeverityLevelInfo represents the "info" or "ignore" severity level.
	SeverityLevelInfo
	// SeverityLevelUnknown represents the "experimental" or "unknown" severity level.
	SeverityLevelUnknown
	// SeverityLevelLow represents the "low" severity level.
	SeverityLevelLow
	// SeverityLevelMedium represents the "medium" severity level.
	SeverityLevelMedium
	// SeverityLevelHigh represents the "high" severity level.
	SeverityLevelHigh
	// SeverityLevelCritical represents the "critical" severity level.
	SeverityLevelCritical
)

// MarshalJSON converts a SeverityLevel value into the JSON representation
func (l SeverityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

// UnmarshalJSON parses a SeverityLevel value from JSON representation
func (l *SeverityLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*l = ParseSeverityLevel(s)
	return nil
}

// ParseSeverityLevel transforms the SeverityLevel into a string
func (l SeverityLevel) String() string {
	switch l {
	case SeverityLevelCritical:
		return "Critical"
	case SeverityLevelHigh:
		return "High"
	case SeverityLevelMedium:
		return "Medium"
	case SeverityLevelLow:
		return "Low"
	case SeverityLevelUnknown:
		return "Unknown"
	case SeverityLevelInfo:
		return "Info"
	}
	return ""
}

// ParseSeverityLevel parses a SeverityLevel value from string
func ParseSeverityLevel(s string) SeverityLevel {
	switch strings.ToLower(s) {
	case "critical":
		return SeverityLevelCritical
	case "high":
		return SeverityLevelHigh
	case "medium":
		return SeverityLevelMedium
	case "low":
		return SeverityLevelLow
	case "experimental", "unknown":
		return SeverityLevelUnknown
	case "ignore", "info":
		return SeverityLevelInfo
	default:
		return SeverityLevelUnknown
	}
}
