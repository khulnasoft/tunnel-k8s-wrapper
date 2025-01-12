package tunnel

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/dustin/go-humanize"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/kas"
)

var (
	// ErrSizeLimit is an error that is thrown in case the Tunnel report is larger than the limit
	ErrSizeLimit = errors.New("tunnel report size limit exceeded")
)

// Report resembles the bytes of a Tunnel report
type Report struct {
	bytes []byte
}

// NewReport returns a new tunnelReport struct by reading a Tunnel report by the given path argument
func NewReport(reportPath string, maxTunnelReportSizeBytes uint64) (*Report, error) {
	bytes, err := os.ReadFile(reportPath)
	if err != nil {
		return nil, fmt.Errorf("could not read tunnel report: %w", err)
	}

	reportSize := uint64(len(bytes))

	if reportSize >= maxTunnelReportSizeBytes {
		return nil, fmt.Errorf("tunnel report is bigger (%s) than the max allowed size of %s: %w", humanize.Bytes(reportSize), humanize.Bytes(maxTunnelReportSizeBytes), ErrSizeLimit)
	}
	return &Report{bytes: bytes}, nil
}

// ToReport converts the raw bytes of the report into a structured kas.Report
func (tr *Report) ToReport() (*kas.Report, error) {
	var report kas.Report
	if err := json.Unmarshal(tr.bytes, &report); err != nil {
		return nil, fmt.Errorf("could not unmarshal report: %w", err)
	}

	return &report, nil
}
