package kas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/kas"
	"github.com/khulnasoft/tunnel-k8s-wrapper/prototool"
)

func TestToProto(t *testing.T) {
	tcs := []struct {
		name     string
		report   kas.Report
		expected *prototool.Payload
	}{
		{
			name:     "OS vuln report",
			report:   osVulnReport,
			expected: &osVulnPayload,
		},
		{
			name:     "Language vuln report",
			report:   langVulnReport,
			expected: &langVulnPayload,
		},
		{
			name:     "Multi container report",
			report:   multiContainerReport,
			expected: &multiContainerPayload,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			protoPayload := tc.report.ToProto("1234")
			assert.Equal(t, tc.expected.Vulnerabilities, protoPayload.Vulnerabilities)
		})
	}
}

func generateReportWithIdentifier(identifier string) kas.Report {
	return kas.Report{
		Resources: []kas.Resource{
			{
				Results: []kas.Result{
					{
						Vulnerabilities: []kas.DetectedVulnerability{
							{
								VulnerabilityID: identifier,
							},
						},
					},
				},
			},
		},
	}
}

func TestConvertIdentifiers(t *testing.T) {
	tcs := []struct {
		name       string
		identifier string
		report     kas.Report
		expected   *prototool.Identifier
	}{
		{
			name:       "CVE Identifier",
			identifier: "CVE-123",
			report:     generateReportWithIdentifier("CVE-123"),
			expected: &prototool.Identifier{
				Type:  "cve",
				Name:  "CVE-123",
				Value: "CVE-123",
				Url:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-123",
			},
		},
		{
			name:       "CWE Identifier",
			identifier: "CWE-123",
			report:     generateReportWithIdentifier("CWE-123"),
			expected: &prototool.Identifier{
				Type:  "cwe",
				Name:  "CWE-123",
				Value: "123",
				Url:   "https://cwe.mitre.org/data/definitions/123.html",
			},
		},
		{
			name:       "OSVDB Identifier",
			identifier: "OSVDB-123",
			report:     generateReportWithIdentifier("OSVDB-123"),
			expected: &prototool.Identifier{
				Type:  "osvdb",
				Name:  "OSVDB-123",
				Value: "OSVDB-123",
				Url:   "https://cve.mitre.org/data/refs/refmap/source-OSVDB.html",
			},
		},
		{
			name:       "USN Identifier",
			identifier: "USN-123",
			report:     generateReportWithIdentifier("USN-123"),
			expected: &prototool.Identifier{
				Type:  "usn",
				Name:  "USN-123",
				Value: "USN-123",
				Url:   "https://usn.ubuntu.com/123/",
			},
		},
		{
			name:       "RHSA Identifier",
			identifier: "RHSA-2019:3892",
			report:     generateReportWithIdentifier("RHSA-2019:3892"),
			expected: &prototool.Identifier{
				Type:  "rhsa",
				Name:  "RHSA-2019:3892",
				Value: "RHSA-2019:3892",
				Url:   "https://access.redhat.com/errata/RHSA-2019:3892",
			},
		},
		{
			name:       "GHSA Identifier",
			identifier: "GHSA-w64w-qqph-5gxm",
			report:     generateReportWithIdentifier("GHSA-w64w-qqph-5gxm"),
			expected: &prototool.Identifier{
				Type:  "ghsa",
				Name:  "GHSA-w64w-qqph-5gxm",
				Value: "GHSA-w64w-qqph-5gxm",
				Url:   "https://github.com/advisories/GHSA-w64w-qqph-5gxm",
			},
		},
		{
			name:       "ELSA Identifier",
			identifier: "ELSA-2017-1101",
			report:     generateReportWithIdentifier("ELSA-2017-1101"),
			expected: &prototool.Identifier{
				Type:  "elsa",
				Name:  "ELSA-2017-1101",
				Value: "ELSA-2017-1101",
				Url:   "https://linux.oracle.com/errata/ELSA-2017-1101.html",
			},
		},
		{
			name:       "H1 Identifier",
			identifier: "HACKERONE-350401",
			report:     generateReportWithIdentifier("HACKERONE-350401"),
			expected: &prototool.Identifier{
				Type:  "hackerone",
				Name:  "HACKERONE-350401",
				Value: "350401",
				Url:   "https://hackerone.com/reports/350401",
			},
		},
		{
			name:       "Not recognised identifier fallsback to CVE",
			identifier: "aaaaaaaaa-2024:4533",
			report:     generateReportWithIdentifier("aaaaaaaaa-2024:4533"),
			expected: &prototool.Identifier{
				Type:  "cve",
				Name:  "aaaaaaaaa-2024:4533",
				Value: "aaaaaaaaa-2024:4533",
				Url:   "https://cve.mitre.org/cgi-bin/cvename.cgi?name=aaaaaaaaa-2024%3A4533",
			},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			protoPayload := tc.report.ToProto("1234")
			require.Equal(t, 1, len(protoPayload.Vulnerabilities))
			require.Equal(t, 1, len(protoPayload.Vulnerabilities[0].Identifiers))
			assert.Equal(t, tc.expected.Name, protoPayload.Vulnerabilities[0].Identifiers[0].GetName())
			assert.Equal(t, tc.expected.Type, protoPayload.Vulnerabilities[0].Identifiers[0].GetType())
			assert.Equal(t, tc.expected.Url, protoPayload.Vulnerabilities[0].Identifiers[0].GetUrl())
			assert.Equal(t, tc.expected.Value, protoPayload.Vulnerabilities[0].Identifiers[0].GetValue())
		})
	}

}
