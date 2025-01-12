package tunnel_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/kas"
	"github.com/khulnasoft/tunnel-k8s-wrapper/tunnel"
)

const (
	maxReportSize uint64 = 100000000
)

func TestNewReport(t *testing.T) {

	tcs := []struct {
		name       string
		before     func()
		after      func()
		reportPath string
		expErr     string
		expReport  *tunnel.Report
	}{
		{
			name:       "Fails tor read the report",
			before:     func() {},
			after:      func() {},
			reportPath: "something/wrong",
			expErr:     "could not read tunnel report: open something/wrong: no such file or directory",
			expReport:  nil,
		},
		{
			name: "Report is bigger than the limit",
			before: func() {
				buffer := make([]byte, 101*1024*1024)
				assert.Nil(t, os.WriteFile("file", buffer, 0400))
			},
			after:      func() { assert.Nil(t, os.Remove("file")) },
			reportPath: "file",
			expErr:     tunnel.ErrSizeLimit.Error(),
			expReport:  nil,
		},
		{
			name:       "Succeeds",
			before:     func() {},
			after:      func() {},
			reportPath: "fixtures/tunnel_report.json",
			expErr:     "",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			tc.before()
			actualReport, err := tunnel.NewReport(tc.reportPath, maxReportSize)
			if tc.expErr != "" {
				assert.Error(t, err, tc.expErr)
				assert.Nil(t, actualReport)
			} else {
				assert.Nil(t, err)
			}
			tc.after()
		})
	}
}

func TestToReport(t *testing.T) {
	expected := kas.Report{
		Resources: []kas.Resource{
			{
				Namespace: "kube-system",
				Kind:      "Pod",
				Name:      "kube-apiserver-kind-control-plane",
				Results: []kas.Result{
					{
						Target: "registry.k8s.io/kube-apiserver:v1.25.3 (debian 11.5)",
						Class:  "os-pkgs",
						Type:   "debian",
					},
				},
			},
		},
	}
	tunnelReport, err := tunnel.NewReport("fixtures/tunnel_report.json", maxReportSize)
	assert.Nil(t, err)
	actual, err := tunnelReport.ToReport()
	assert.Nil(t, err)
	assert.Equal(t, expected, *actual)
}
