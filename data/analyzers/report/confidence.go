package report

import (
	"encoding/json"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// This content of this file are required by the tunnel-k8s-wrapper project                                       //
//                                                                                                               //
// All structs and functions are references from                                                                 //
// https://github.com/github-org/security-products/analyzers/report/-/blob/v3.7.1/vulnerability.go?ref_type=tags //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ConfidenceLevel is the vulnerability confidence level reported by scanner.
type ConfidenceLevel int

const (
	// ConfidenceLevelUndefined is a stub confidence value for the case when it was not reported by scanner.
	ConfidenceLevelUndefined ConfidenceLevel = iota
	// ConfidenceLevelIgnore represents the "ignore" confidence level.
	ConfidenceLevelIgnore
	// ConfidenceLevelUnknown represents the "unknown" confidence level.
	ConfidenceLevelUnknown
	// ConfidenceLevelExperimental represents the "experimental" confidence level.
	ConfidenceLevelExperimental
	// ConfidenceLevelLow represents the "low" confidence level.
	ConfidenceLevelLow
	// ConfidenceLevelMedium represents the "medium" confidence level.
	ConfidenceLevelMedium
	// ConfidenceLevelHigh represents the "high" confidence level.
	ConfidenceLevelHigh
	// ConfidenceLevelConfirmed represents the "critical" or "confirmed" confidence level.
	ConfidenceLevelConfirmed
)

// MarshalJSON converts a ConfidenceLevel value into the JSON representation
func (l ConfidenceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.String())
}

// UnmarshalJSON parses a ConfidenceLevel value from JSON representation
func (l *ConfidenceLevel) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*l = ParseConfidenceLevel(s)
	return nil
}

func (l ConfidenceLevel) String() string {
	switch l {
	case ConfidenceLevelConfirmed:
		return "Confirmed"
	case ConfidenceLevelHigh:
		return "High"
	case ConfidenceLevelMedium:
		return "Medium"
	case ConfidenceLevelLow:
		return "Low"
	case ConfidenceLevelExperimental:
		return "Experimental"
	case ConfidenceLevelUnknown:
		return "Unknown"
	case ConfidenceLevelIgnore:
		return "Ignore"
	}
	return ""
}

// ParseConfidenceLevel parses a ConfidenceLevel value from string
func ParseConfidenceLevel(s string) ConfidenceLevel {
	switch strings.ToLower(s) {
	case "critical", "confirmed":
		return ConfidenceLevelConfirmed
	case "high":
		return ConfidenceLevelHigh
	case "medium":
		return ConfidenceLevelMedium
	case "low":
		return ConfidenceLevelLow
	case "experimental":
		return ConfidenceLevelExperimental
	case "unknown":
		return ConfidenceLevelUnknown
	case "ignore":
		return ConfidenceLevelIgnore
	default:
		return ConfidenceLevelUnknown
	}
}
