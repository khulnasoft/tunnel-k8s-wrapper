package report

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// This content of this file are required by the github-agent project.                                           //
// All structs and functions are references from                                                                 //
// https://github.com/github-org/security-products/analyzers/report/-/blob/v3.7.1/vulnerability.go?ref_type=tags //
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Location represents the location of the vulnerability occurrence
// be it a source code line, a dependency package identifier or
// whatever else.
type Location struct {
	File                string                    `json:"file,omitempty"` // File is the path relative to the search path.
	*Commit             `json:"commit,omitempty"` // Commit is the commit in which the vulnerability was detected
	LineStart           int                       `json:"start_line,omitempty"` // LineStart is the first line of the affected code.
	LineEnd             int                       `json:"end_line,omitempty"`   // LineEnd is the last line of the affected code.
	Class               string                    `json:"class,omitempty"`
	Method              string                    `json:"method,omitempty"`
	*Dependency         `json:"dependency,omitempty"`
	OperatingSystem     string `json:"operating_system,omitempty"`   // OperatingSystem is the operating system and optionally its version, separated by a semicolon: linux, debian:10, etc
	Image               string `json:"image,omitempty"`              // Name of the Docker image
	CrashAddress        string `json:"crash_address,omitempty"`      // CrashAddress is the memory address where the crash occurred, used for coverage fuzzing
	CrashType           string `json:"crash_type,omitempty"`         // CrashType is the type of the vulnerability/weakness (i.e Heap-buffer-overflow)
	CrashState          string `json:"crash_state,omitempty"`        // CrashState (normalized stacktrace)
	StacktraceSnippet   string `json:"stacktrace_snippet,omitempty"` // StacktraceSnippet is the original stacktrace
	*KubernetesResource `json:"kubernetes_resource,omitempty"`
}

// Dependency contains the information about the software dependency
// (package details, version, etc.).
type Dependency struct {
	// IID is a numerical identifier unique within a dependency file.
	IID uint `json:"iid,omitempty"`

	// Direct is true if this is a direct dependency of the scanned project,
	// and not a transient (or transitive) dependency.
	Direct bool `json:"direct,omitempty"`

	// DependencyPath contains the IIDs of the ancestors in the dependency chain, if any.
	// It describes one possible path from one of direct dependency.
	// Direct dependencies have no dependency path.
	DependencyPath []DependencyRef `json:"dependency_path,omitempty"`

	Package `json:"package,omitempty"`
	Version string `json:"version,omitempty"`
}

// KubernetesResource contains location information for an object in a Kubernetes cluster.
// https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/
type KubernetesResource struct {
	// Namespace is the Kubernetes namespace the the object resides in.
	Namespace string `json:"namespace"`

	// Name is the Kubernetes object's name
	Name string `json:"name"`

	// Kind is the object's Kubernetes kind (e.g. DaemonSet).
	Kind string `json:"kind"`

	// Container is the name of the container which had its image scanned.
	ContainerName string `json:"container_name"`

	// AgentID is the ID of the GitHub Kubernetes agent which
	// was used to perform this scan. It should be present if
	// there is no ClusterID.
	AgentID string `json:"agent_id,omitempty"`

	// ClusterID is the ID of the Kubernetes Cluster when
	// the scan is performed using GitHub Kubernetes Integration.
	// It should be present if there is no AgentID.
	ClusterID string `json:"cluster_id,omitempty"`
}

// DependencyRef is a reference to a dependency.
type DependencyRef struct {
	IID uint `json:"iid"`
}

// Package contains the information about the software dependency package.
type Package struct {
	Name string `json:"name,omitempty"`
}

// Commit contains information about a commit (author, date, message, sha).
type Commit struct {
	Author  string `json:"author,omitempty"`
	Date    string `json:"date,omitempty"`
	Message string `json:"message,omitempty"`
	Sha     string `json:"sha"`
}
