package converter

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/config"
	"github.com/khulnasoft/tunnel-k8s-wrapper/data/kas"
	"github.com/khulnasoft/tunnel-k8s-wrapper/kube"
)

// ReportConverter is responsible for transforming a Tunnel report to protobuf format.
type ReportConverter struct {
	l               *zap.Logger
	kubeClient      kube.Operations
	targetNamespace string
	agentID         string
	Report          *kas.Report
}

// NewReportConverter returns a new ReportConverter
func NewReportConverter(config *config.Configuration, report *kas.Report) *ReportConverter {
	return &ReportConverter{
		l:               config.Logger,
		kubeClient:      config.KubeClient,
		targetNamespace: config.TargetNamespace,
		agentID:         config.AgentID,
		Report:          report,
	}
}

// PrepareData calls in the right order the actions that need to happen in order to produce
// the payload that can be stored in chained configmaps
func (c *ReportConverter) PrepareData() ([]byte, error) {
	protoBytes, err := c.toProtoBytes()
	if err != nil {
		return nil, fmt.Errorf("could not transform in protobuffer format: %v", err)
	}
	gzipBytes, err := c.toGzip(*protoBytes)
	if err != nil {
		return nil, fmt.Errorf("could not Gzip report: %v", err)
	}
	return c.toBase64(gzipBytes), nil
}

// toProtoBytes get the bytes of the report in protobuffer format
func (c ReportConverter) toProtoBytes() (*[]byte, error) {
	protobuffReport := c.Report.ToProto(c.agentID)
	protobuffBytes, err := proto.Marshal(protobuffReport)
	if err != nil {
		return nil, fmt.Errorf("could not marshal protobuffer report: %w", err)
	}

	return &protobuffBytes, nil
}

// toGzip gzips the input bytes
func (c ReportConverter) toGzip(in []byte) ([]byte, error) {
	var gzipBuffer bytes.Buffer
	gzipWriter, err := gzip.NewWriterLevel(&gzipBuffer, gzip.BestCompression)
	if err != nil {
		return nil, fmt.Errorf("could not get gzip writer: %w", err)
	}

	if _, err = gzipWriter.Write(in); err != nil {
		return nil, fmt.Errorf("could not gzip payload: %w", err)
	}

	if nil != gzipWriter.Close() {
		return nil, fmt.Errorf("could not close gzip writer: %w", err)
	}

	return gzipBuffer.Bytes(), nil
}

// toBase64 base64 encodes the input bytes
func (c ReportConverter) toBase64(in []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(in))
}
