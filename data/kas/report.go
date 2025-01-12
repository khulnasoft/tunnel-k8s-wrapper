package kas

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/analyzers/report"
	"github.com/khulnasoft/tunnel-k8s-wrapper/prototool"
)

// Report Type referenced from Tunnel https://github.com/github-org/security-products/dependencies/tunnel/-/blob/v0.54.0/pkg/k8s/report/report.go#L40
type Report struct {
	Resources []Resource `json:",omitempty"`
}

// Resource Type referenced from Tunnel https://github.com/khulnasoft/tunnel/blob/v0.0.3/pkg/k8s/report/report.go#L60
type Resource struct {
	Namespace string     `json:",omitempty"`
	Kind      string     `json:",omitempty"`
	Name      string     `json:",omitempty"`
	Metadata  []Metadata `json:",omitempty"`
	Results   []Result   `json:",omitempty"`
}

// Metadata Type referenced from Tunnel https://github.com/github-org/security-products/dependencies/tunnel/-/blob/v0.54.0/pkg/types/report.go#L27
type Metadata struct {
	OS       OS       `json:",omitempty"`
	RepoTags []string `json:",omitempty"`
}

// OS Type referenced from Tunnel https://github.com/github-org/security-products/dependencies/tunnel/-/blob/v0.54.0/pkg/fanal/types/artifact.go#L9
type OS struct {
	Family string
	Name   string
}

// Result Type referenced from Tunnel https://github.com/github-org/security-products/dependencies/tunnel/-/blob/v0.54.0/pkg/types/report.go#L109
type Result struct {
	Target          string                  `json:"Target"`
	Class           string                  `json:"Class,omitempty"`
	Type            string                  `json:"Type,omitempty"`
	Vulnerabilities []DetectedVulnerability `json:"Vulnerabilities,omitempty"`
}

// DetectedVulnerability Type referenced from Tunnel https://github.com/github-org/security-products/dependencies/tunnel/-/blob/v0.54.0/pkg/types/vulnerability.go#L9
type DetectedVulnerability struct {
	VulnerabilityID  string `json:",omitempty"`
	PkgName          string `json:",omitempty"`
	InstalledVersion string `json:",omitempty"`
	FixedVersion     string `json:",omitempty"`
	PrimaryURL       string `json:",omitempty"`

	// Embed vulnerability details
	Vulnerability
}

// Vulnerability Type referenced from Tunnel-db https://github.com/github-org/security-products/dependencies/tunnel-db/-/blob/d23a6ca8ba04f8acaeac9b1d2e1c52c5242b2814/pkg/types/types.go#L177 referenced by Tunnel v0.54.0
type Vulnerability struct {
	Title            string     `json:",omitempty"`
	Description      string     `json:",omitempty"`
	Severity         string     `json:",omitempty"` // Selected from VendorSeverity, depending on a scan target
	References       []string   `json:",omitempty"`
	PublishedDate    *time.Time `json:",omitempty"` // Take from NVD
	LastModifiedDate *time.Time `json:",omitempty"` // Take from NVD
}

// ToProto converts a report into its protobuffer format
func (c *Report) ToProto(agentID string) *prototool.Payload {
	payload := prototool.Payload{}
	for _, resource := range c.Resources {
		image, operatingSystem := c.findImageAndOS(resource)
		kubernetesResource := c.convertKubernetesResource(resource, agentID)
		results := resource.Results
		for _, result := range results {
			vulns := result.Vulnerabilities
			for _, vuln := range vulns {
				payload.Vulnerabilities = append(payload.Vulnerabilities, &prototool.Vulnerability{
					Name:        vuln.VulnerabilityID,
					Message:     fmt.Sprintf("%s in %s", vuln.VulnerabilityID, vuln.PkgName),
					Description: vuln.Description,
					Solution:    fmt.Sprintf("Upgrade %s from %s to %s", vuln.PkgName, vuln.InstalledVersion, vuln.FixedVersion),
					Severity:    strings.ToUpper(c.convertSeverity(vuln.Severity).String()), // severity constants are in upper case
					Confidence:  report.ConfidenceLevelUnknown.String(),
					Identifiers: c.convertIdentifiers(vuln),
					Links:       c.convertLinks(vuln),
					Location:    c.convertLocation(image, operatingSystem, kubernetesResource, vuln),
				})
			}
		}
	}
	return &payload
}

// findImageAndOS identifies the image and OS associated with the resource.
// When transmitting report to Github the image name is required for users to identify the source of the vulnerability.
func (c *Report) findImageAndOS(resource Resource) (image string, operatingSystem string) {
	// When using tunnel k8s with --report=all we know that the Metadata array will have exactly one field
	// For more info see https://github.com/github-org/github/-/issues/480838#note_2288599658
	if len(resource.Metadata) == 0 {
		return
	}
	repoTags := resource.Metadata[0].RepoTags
	// I've only encountered resources with one repo tag so far, which is why I'm only returning the first repo tag as the image.
	if len(repoTags) > 0 {
		image = repoTags[0]
	}

	if resource.Metadata[0].OS.Family != "" {
		operatingSystem = fmt.Sprintf("%s %s", resource.Metadata[0].OS.Family, resource.Metadata[0].OS.Name)
	}
	return
}

// Location is used to fingerprint(uniquely identify) the resource in github.
// The fields used for fingerprinting are: agentID, k8sresource.namespace,
// k8sresource.kind, k8sresource.name, k8sresource.container, PkgName
// As defined here:
// https://github.com/github-org/github/-/blob/f50075762cf33d3841b88bb191770776b07ede77/ee/app/services/vulnerabilities/starboard_vulnerability_create_service.rb#L62
// WARNING! Be extra careful when changing these fields as it could cause new resources to be
// flagged to the user when they might have been previously addressed.
func (c *Report) convertKubernetesResource(resource Resource, agentID string) *prototool.KubernetesResource {
	return &prototool.KubernetesResource{
		Namespace:     resource.Namespace,
		Name:          resource.Name,
		Kind:          resource.Kind,
		AgentId:       agentID,
		ContainerName: "", //NOTE In Tunnel k8s, the ContainerName is not provided. https://github.com/github-org/security-products/dependencies/tunnel/-/blob/v0.38.3/pkg/k8s/report/report.go#L58-L69.
		// Leaving ContainerName as an empty string as such.
		// This does not affect the fingerprint as the field referenced in github is `container` while the one defined in KubernetesResource is `container_name`.
	}
}

// Adapted from severityNames in Tunnel-db
// https://github.com/github-org/security-products/dependencies/tunnel-db/-/blob/2bd1364579ec652f8f595c4a61595fd9575e8496/pkg/types/types.go#L35
const (
	TunnelSeverityCritical = "CRITICAL"
	TunnelSeverityHigh     = "HIGH"
	TunnelSeverityMedium   = "MEDIUM"
	TunnelSeverityLow      = "LOW"

	TunnelSeverityNone    = "NONE" // Kept for legacy reasons since starboard contains this severity level
	TunnelSeverityUnknown = "UNKNOWN"
)

var severityMapping = map[string]report.SeverityLevel{
	TunnelSeverityCritical: report.SeverityLevelCritical,
	TunnelSeverityHigh:     report.SeverityLevelHigh,
	TunnelSeverityMedium:   report.SeverityLevelMedium,
	TunnelSeverityLow:      report.SeverityLevelLow,
	TunnelSeverityNone:     report.SeverityLevelInfo,
	TunnelSeverityUnknown:  report.SeverityLevelUnknown,
}

func (c *Report) convertSeverity(severity string) report.SeverityLevel {
	sev, ok := severityMapping[severity]
	if !ok {
		return report.SeverityLevelUnknown
	}
	return sev
}

func (c *Report) convertIdentifiers(vuln DetectedVulnerability) []*prototool.Identifier {
	id := vuln.VulnerabilityID
	identifier, ok := report.ParseIdentifierID(id)
	if !ok {
		// fallback to CVE
		return []*prototool.Identifier{
			{
				Type:  string(report.IdentifierTypeCVE),
				Name:  id,
				Value: id,
				Url:   fmt.Sprintf("https://cve.mitre.org/cgi-bin/cvename.cgi?name=%s", url.QueryEscape(id)),
			},
		}
	}
	return []*prototool.Identifier{
		{
			Type:  string(identifier.Type),
			Name:  identifier.Name,
			Value: identifier.Value,
			Url:   identifier.URL,
		},
	}
}

func (c *Report) convertLinks(vuln DetectedVulnerability) []*prototool.Link {
	var links []*prototool.Link // nolint:prealloc
	if vuln.PrimaryURL != "" {
		links = append(links, &prototool.Link{Url: vuln.PrimaryURL})
	}

	for _, r := range vuln.References {
		links = append(links, &prototool.Link{Url: r})
	}
	return links
}

func (c *Report) convertLocation(image string,
	operatingSystem string, kubernetesResource *prototool.KubernetesResource,
	vuln DetectedVulnerability) *prototool.Location {
	return &prototool.Location{
		Dependency: &prototool.Dependency{
			Package: &prototool.Package{Name: vuln.PkgName},
			Version: vuln.InstalledVersion,
		},
		KubernetesResource: kubernetesResource,
		Image:              image,
		OperatingSystem:    operatingSystem,
	}
}
