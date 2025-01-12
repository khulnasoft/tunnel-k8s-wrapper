package kas_test

import (
	"time"

	"github.com/khulnasoft/tunnel-k8s-wrapper/data/kas"
)

var (
	someTimestamp = time.Now()
	osVulnReport  = kas.Report{
		Resources: []kas.Resource{
			{
				Namespace: "kube-system",
				Kind:      "Pod",
				Name:      "kube-apiserver-kind-control-plane",
				Metadata: []kas.Metadata{
					{
						RepoTags: []string{
							"registry.k8s.io/kube-apiserver:v1.25.3",
						},
						OS: kas.OS{
							Family: "debian",
							Name:   "11.5",
						},
					},
				},
				Results: []kas.Result{
					{
						Target: "registry.k8s.io/kube-apiserver:v1.25.3 (debian 11.5)",
						Class:  "os-pkgs",
						Type:   "debian",
						Vulnerabilities: []kas.DetectedVulnerability{
							{
								VulnerabilityID:  "CVE-2023-2253",
								PkgName:          "github.com/docker/distribution",
								InstalledVersion: "v2.8.1+incompatible",
								FixedVersion:     "2.8.2-beta.1",
								PrimaryURL:       "https://avd.aquasec.com/nvd/cve-2023-2253",
								Vulnerability: kas.Vulnerability{
									Title:       "DoS from malicious API request",
									Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
									Severity:    "HIGH",
									References: []string{
										"https://access.redhat.com/security/cve/CVE-2023-2253",
										"https://bugzilla.redhat.com/show_bug.cgi?id=2189886",
									},
									PublishedDate:    &someTimestamp,
									LastModifiedDate: &someTimestamp,
								},
							},
							{
								VulnerabilityID:  "RHSA-2024:4533",
								PkgName:          "github.com/something",
								InstalledVersion: "v2.8.1+incompatible",
								FixedVersion:     "2.8.2-beta.1",
								PrimaryURL:       "https://access.redhat.com/errata/RHSA-2024:4533",
								Vulnerability: kas.Vulnerability{
									Title:       "DoS from malicious API request",
									Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
									Severity:    "HIGH",
									References: []string{
										"https://access.redhat.com/security/updates/classification/#important",
									},
									PublishedDate:    &someTimestamp,
									LastModifiedDate: &someTimestamp,
								},
							},
						},
					},
				},
			},
		},
	}
	langVulnReport = kas.Report{
		Resources: []kas.Resource{
			{
				Namespace: "default",
				Kind:      "Pod",
				Name:      "lang-vuln-pod",
				Metadata: []kas.Metadata{
					{
						RepoTags: []string{
							"public.ecr.aws/cloudwatch-agent/cloudwatch-agent:1.247360.0b252689",
						},
					},
				},
				Results: []kas.Result{
					{
						Target: "opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent",
						Class:  "lang-pkgs",
						Type:   "gobinary",
						Vulnerabilities: []kas.DetectedVulnerability{
							{
								VulnerabilityID:  "CVE-2022-23471",
								PkgName:          "github.com/containerd/containerd",
								InstalledVersion: "v1.6.8",
								FixedVersion:     "1.5.16, 1.6.12",
								PrimaryURL:       "https://avd.aquasec.com/nvd/cve-2022-23471",
								Vulnerability: kas.Vulnerability{
									Title:       "containerd is an open source container runtime. A bug was found in con ...",
									Description: "containerd is an open source container runtime. A bug was found in containerd's CRI implementation where a user can exhaust memory on the host. In the CRI stream server, a goroutine is launched to handle terminal resize events if a TTY is requested. If the user's process fails to launch due to, for example, a faulty command, the goroutine will be stuck waiting to send without a receiver, resulting in a memory leak. Kubernetes and crictl can both be configured to use containerd's CRI implementation and the stream server is used for handling container IO. This bug has been fixed in containerd 1.6.12 and 1.5.16.  Users should update to these versions to resolve the issue. Users unable to upgrade should ensure that only trusted images and commands are used and that only trusted users have permissions to execute commands in running containers. ",
									Severity:    "MEDIUM",
									References: []string{
										"https://github.com/containerd/containerd",
										"https://github.com/containerd/containerd/commit/241563be06a3de8b6a849414c4e805b68d3bb295",
										"https://github.com/containerd/containerd/commit/a05d175400b1145e5e6a735a6710579d181e7fb0",
										"https://github.com/containerd/containerd/releases/tag/v1.5.16",
										"https://github.com/containerd/containerd/releases/tag/v1.6.12",
										"https://github.com/containerd/containerd/security/advisories/GHSA-2qjp-425j-52j9",
										"https://nvd.nist.gov/vuln/detail/CVE-2022-23471",
										"https://security.gentoo.org/glsa/202401-31",
										"https://ubuntu.com/security/notices/USN-5776-1",
										"https://www.cve.org/CVERecord?id=CVE-2022-23471",
									},
								},
							},
							{
								VulnerabilityID:  "CVE-2023-25153",
								PkgName:          "github.com/containerd/containerd",
								InstalledVersion: "v1.6.8",
								FixedVersion:     "1.5.18, 1.6.18",
								PrimaryURL:       "https://avd.aquasec.com/nvd/cve-2023-25153",
								Vulnerability: kas.Vulnerability{
									Title:       "containerd: OCI image importer memory exhaustion",
									Description: "containerd is an open source container runtime. Before versions 1.6.18 and 1.5.18, when importing an OCI image, there was no limit on the number of bytes read for certain files. A maliciously crafted image with a large file where a limit was not applied could cause a denial of service. This bug has been fixed in containerd 1.6.18 and 1.5.18.  Users should update to these versions to resolve the issue. As a workaround, ensure that only trusted images are used and that only trusted users have permissions to import images.",
									Severity:    "MEDIUM",
									References: []string{
										"https://access.redhat.com/security/cve/CVE-2023-25153",
										"https://github.com/containerd/containerd",
										"https://github.com/containerd/containerd/commit/0c314901076a74a7b797a545d2f462285fdbb8c4",
										"https://github.com/containerd/containerd/releases/tag/v1.5.18",
										"https://github.com/containerd/containerd/releases/tag/v1.6.18",
										"https://github.com/containerd/containerd/security/advisories/GHSA-259w-8hf6-59c2",
										"https://nvd.nist.gov/vuln/detail/CVE-2023-25153",
										"https://pkg.go.dev/vuln/GO-2023-1573",
										"https://ubuntu.com/security/notices/USN-6202-1",
										"https://www.cve.org/CVERecord?id=CVE-2023-25153",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	multiContainerReport = kas.Report{
		Resources: []kas.Resource{
			{
				Namespace: "kube-system",
				Kind:      "Pod",
				Name:      "kube-apiserver-kind-control-plane",
				Metadata: []kas.Metadata{
					{
						RepoTags: []string{
							"registry.k8s.io/kube-apiserver:v1.25.3",
						},
						OS: kas.OS{
							Family: "debian",
							Name:   "11.5",
						},
					},
				},
				Results: []kas.Result{
					{
						Target: "registry.k8s.io/kube-apiserver:v1.25.3 (debian 11.5)",
						Class:  "os-pkgs",
						Type:   "debian",
						Vulnerabilities: []kas.DetectedVulnerability{
							{
								VulnerabilityID:  "CVE-2023-2253",
								PkgName:          "github.com/docker/distribution",
								InstalledVersion: "v2.8.1+incompatible",
								FixedVersion:     "2.8.2-beta.1",
								PrimaryURL:       "https://avd.aquasec.com/nvd/cve-2023-2253",
								Vulnerability: kas.Vulnerability{
									Title:       "DoS from malicious API request",
									Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
									Severity:    "HIGH",
									References: []string{
										"https://access.redhat.com/security/cve/CVE-2023-2253",
										"https://bugzilla.redhat.com/show_bug.cgi?id=2189886",
									},
									PublishedDate:    &someTimestamp,
									LastModifiedDate: &someTimestamp,
								},
							},
							{
								VulnerabilityID:  "RHSA-2024:4533",
								PkgName:          "github.com/something",
								InstalledVersion: "v2.8.1+incompatible",
								FixedVersion:     "2.8.2-beta.1",
								PrimaryURL:       "https://access.redhat.com/errata/RHSA-2024:4533",
								Vulnerability: kas.Vulnerability{
									Title:       "DoS from malicious API request",
									Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
									Severity:    "HIGH",
									References: []string{
										"https://access.redhat.com/security/updates/classification/#important",
									},
									PublishedDate:    &someTimestamp,
									LastModifiedDate: &someTimestamp,
								},
							},
						},
					},
				},
			},
			{
				Namespace: "default",
				Kind:      "Pod",
				Name:      "lang-vuln-pod",
				Metadata: []kas.Metadata{
					{
						RepoTags: []string{
							"public.ecr.aws/cloudwatch-agent/cloudwatch-agent:1.247360.0b252689",
						},
					},
				},
				Results: []kas.Result{
					{
						Target: "opt/aws/amazon-cloudwatch-agent/bin/amazon-cloudwatch-agent",
						Class:  "lang-pkgs",
						Type:   "gobinary",
						Vulnerabilities: []kas.DetectedVulnerability{
							{
								VulnerabilityID:  "CVE-2022-23471",
								PkgName:          "github.com/containerd/containerd",
								InstalledVersion: "v1.6.8",
								FixedVersion:     "1.5.16, 1.6.12",
								PrimaryURL:       "https://avd.aquasec.com/nvd/cve-2022-23471",
								Vulnerability: kas.Vulnerability{
									Title:       "containerd is an open source container runtime. A bug was found in con ...",
									Description: "containerd is an open source container runtime. A bug was found in containerd's CRI implementation where a user can exhaust memory on the host. In the CRI stream server, a goroutine is launched to handle terminal resize events if a TTY is requested. If the user's process fails to launch due to, for example, a faulty command, the goroutine will be stuck waiting to send without a receiver, resulting in a memory leak. Kubernetes and crictl can both be configured to use containerd's CRI implementation and the stream server is used for handling container IO. This bug has been fixed in containerd 1.6.12 and 1.5.16.  Users should update to these versions to resolve the issue. Users unable to upgrade should ensure that only trusted images and commands are used and that only trusted users have permissions to execute commands in running containers. ",
									Severity:    "MEDIUM",
									References: []string{
										"https://github.com/containerd/containerd",
										"https://github.com/containerd/containerd/commit/241563be06a3de8b6a849414c4e805b68d3bb295",
										"https://github.com/containerd/containerd/commit/a05d175400b1145e5e6a735a6710579d181e7fb0",
										"https://github.com/containerd/containerd/releases/tag/v1.5.16",
										"https://github.com/containerd/containerd/releases/tag/v1.6.12",
										"https://github.com/containerd/containerd/security/advisories/GHSA-2qjp-425j-52j9",
										"https://nvd.nist.gov/vuln/detail/CVE-2022-23471",
										"https://security.gentoo.org/glsa/202401-31",
										"https://ubuntu.com/security/notices/USN-5776-1",
										"https://www.cve.org/CVERecord?id=CVE-2022-23471",
									},
								},
							},
							{
								VulnerabilityID:  "CVE-2023-25153",
								PkgName:          "github.com/containerd/containerd",
								InstalledVersion: "v1.6.8",
								FixedVersion:     "1.5.18, 1.6.18",
								PrimaryURL:       "https://avd.aquasec.com/nvd/cve-2023-25153",
								Vulnerability: kas.Vulnerability{
									Title:       "containerd: OCI image importer memory exhaustion",
									Description: "containerd is an open source container runtime. Before versions 1.6.18 and 1.5.18, when importing an OCI image, there was no limit on the number of bytes read for certain files. A maliciously crafted image with a large file where a limit was not applied could cause a denial of service. This bug has been fixed in containerd 1.6.18 and 1.5.18.  Users should update to these versions to resolve the issue. As a workaround, ensure that only trusted images are used and that only trusted users have permissions to import images.",
									Severity:    "MEDIUM",
									References: []string{
										"https://access.redhat.com/security/cve/CVE-2023-25153",
										"https://github.com/containerd/containerd",
										"https://github.com/containerd/containerd/commit/0c314901076a74a7b797a545d2f462285fdbb8c4",
										"https://github.com/containerd/containerd/releases/tag/v1.5.18",
										"https://github.com/containerd/containerd/releases/tag/v1.6.18",
										"https://github.com/containerd/containerd/security/advisories/GHSA-259w-8hf6-59c2",
										"https://nvd.nist.gov/vuln/detail/CVE-2023-25153",
										"https://pkg.go.dev/vuln/GO-2023-1573",
										"https://ubuntu.com/security/notices/USN-6202-1",
										"https://www.cve.org/CVERecord?id=CVE-2023-25153",
									},
								},
							},
						},
					},
				},
			},
		},
	}
)
