package kas_test

import "github.com/khulnasoft/tunnel-k8s-wrapper/prototool"

var (
	osVulnPayload = prototool.Payload{
		Vulnerabilities: []*prototool.Vulnerability{
			{
				Name:        "CVE-2023-2253",
				Message:     "CVE-2023-2253 in github.com/docker/distribution",
				Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
				Solution:    "Upgrade github.com/docker/distribution from v2.8.1+incompatible to 2.8.2-beta.1",
				Severity:    "HIGH",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "cve", Name: "CVE-2023-2253", Value: "CVE-2023-2253", Url: "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-2253",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "kube-system",
						Kind:          "Pod",
						Name:          "kube-apiserver-kind-control-plane",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "registry.k8s.io/kube-apiserver:v1.25.3",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/docker/distribution",
						},
						Version: "v2.8.1+incompatible",
					},
					OperatingSystem: "debian 11.5",
				},
				Links: []*prototool.Link{
					{
						Url: "https://avd.aquasec.com/nvd/cve-2023-2253",
					},
					{
						Url: "https://access.redhat.com/security/cve/CVE-2023-2253",
					},
					{
						Url: "https://bugzilla.redhat.com/show_bug.cgi?id=2189886",
					},
				},
			},
			{
				Name:        "RHSA-2024:4533",
				Message:     "RHSA-2024:4533 in github.com/something",
				Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
				Solution:    "Upgrade github.com/something from v2.8.1+incompatible to 2.8.2-beta.1",
				Severity:    "HIGH",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "rhsa", Name: "RHSA-2024:4533", Value: "RHSA-2024:4533", Url: "https://access.redhat.com/errata/RHSA-2024:4533",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "kube-system",
						Kind:          "Pod",
						Name:          "kube-apiserver-kind-control-plane",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "registry.k8s.io/kube-apiserver:v1.25.3",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/something",
						},
						Version: "v2.8.1+incompatible",
					},
					OperatingSystem: "debian 11.5",
				},
				Links: []*prototool.Link{
					{
						Url: "https://access.redhat.com/errata/RHSA-2024:4533",
					},
					{
						Url: "https://access.redhat.com/security/updates/classification/#important",
					},
				},
			},
		},
	}
	langVulnPayload = prototool.Payload{
		Vulnerabilities: []*prototool.Vulnerability{
			{
				Name:        "CVE-2022-23471",
				Message:     "CVE-2022-23471 in github.com/containerd/containerd",
				Description: "containerd is an open source container runtime. A bug was found in containerd's CRI implementation where a user can exhaust memory on the host. In the CRI stream server, a goroutine is launched to handle terminal resize events if a TTY is requested. If the user's process fails to launch due to, for example, a faulty command, the goroutine will be stuck waiting to send without a receiver, resulting in a memory leak. Kubernetes and crictl can both be configured to use containerd's CRI implementation and the stream server is used for handling container IO. This bug has been fixed in containerd 1.6.12 and 1.5.16.  Users should update to these versions to resolve the issue. Users unable to upgrade should ensure that only trusted images and commands are used and that only trusted users have permissions to execute commands in running containers. ",
				Solution:    "Upgrade github.com/containerd/containerd from v1.6.8 to 1.5.16, 1.6.12",
				Severity:    "MEDIUM",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "cve", Name: "CVE-2022-23471", Value: "CVE-2022-23471", Url: "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2022-23471",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "default",
						Kind:          "Pod",
						Name:          "lang-vuln-pod",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "public.ecr.aws/cloudwatch-agent/cloudwatch-agent:1.247360.0b252689",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/containerd/containerd",
						},
						Version: "v1.6.8",
					},
					OperatingSystem: "",
				},
				Links: []*prototool.Link{
					{Url: "https://avd.aquasec.com/nvd/cve-2022-23471"},
					{Url: "https://github.com/containerd/containerd"},
					{Url: "https://github.com/containerd/containerd/commit/241563be06a3de8b6a849414c4e805b68d3bb295"},
					{Url: "https://github.com/containerd/containerd/commit/a05d175400b1145e5e6a735a6710579d181e7fb0"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.5.16"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.6.12"},
					{Url: "https://github.com/containerd/containerd/security/advisories/GHSA-2qjp-425j-52j9"},
					{Url: "https://nvd.nist.gov/vuln/detail/CVE-2022-23471"},
					{Url: "https://security.gentoo.org/glsa/202401-31"},
					{Url: "https://ubuntu.com/security/notices/USN-5776-1"},
					{Url: "https://www.cve.org/CVERecord?id=CVE-2022-23471"},
				},
			},
			{
				Name:        "CVE-2023-25153",
				Message:     "CVE-2023-25153 in github.com/containerd/containerd",
				Description: "containerd is an open source container runtime. Before versions 1.6.18 and 1.5.18, when importing an OCI image, there was no limit on the number of bytes read for certain files. A maliciously crafted image with a large file where a limit was not applied could cause a denial of service. This bug has been fixed in containerd 1.6.18 and 1.5.18.  Users should update to these versions to resolve the issue. As a workaround, ensure that only trusted images are used and that only trusted users have permissions to import images.",
				Solution:    "Upgrade github.com/containerd/containerd from v1.6.8 to 1.5.18, 1.6.18",
				Severity:    "MEDIUM",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "cve", Name: "CVE-2023-25153", Value: "CVE-2023-25153", Url: "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-25153",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "default",
						Kind:          "Pod",
						Name:          "lang-vuln-pod",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "public.ecr.aws/cloudwatch-agent/cloudwatch-agent:1.247360.0b252689",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/containerd/containerd",
						},
						Version: "v1.6.8",
					},
					OperatingSystem: "",
				},
				Links: []*prototool.Link{
					{Url: "https://avd.aquasec.com/nvd/cve-2023-25153"},
					{Url: "https://access.redhat.com/security/cve/CVE-2023-25153"},
					{Url: "https://github.com/containerd/containerd"},
					{Url: "https://github.com/containerd/containerd/commit/0c314901076a74a7b797a545d2f462285fdbb8c4"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.5.18"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.6.18"},
					{Url: "https://github.com/containerd/containerd/security/advisories/GHSA-259w-8hf6-59c2"},
					{Url: "https://nvd.nist.gov/vuln/detail/CVE-2023-25153"},
					{Url: "https://pkg.go.dev/vuln/GO-2023-1573"},
					{Url: "https://ubuntu.com/security/notices/USN-6202-1"},
					{Url: "https://www.cve.org/CVERecord?id=CVE-2023-25153"},
				},
			},
		},
	}
	multiContainerPayload = prototool.Payload{
		Vulnerabilities: []*prototool.Vulnerability{
			{
				Name:        "CVE-2023-2253",
				Message:     "CVE-2023-2253 in github.com/docker/distribution",
				Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
				Solution:    "Upgrade github.com/docker/distribution from v2.8.1+incompatible to 2.8.2-beta.1",
				Severity:    "HIGH",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "cve", Name: "CVE-2023-2253", Value: "CVE-2023-2253", Url: "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-2253",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "kube-system",
						Kind:          "Pod",
						Name:          "kube-apiserver-kind-control-plane",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "registry.k8s.io/kube-apiserver:v1.25.3",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/docker/distribution",
						},
						Version: "v2.8.1+incompatible",
					},
					OperatingSystem: "debian 11.5",
				},
				Links: []*prototool.Link{
					{
						Url: "https://avd.aquasec.com/nvd/cve-2023-2253",
					},
					{
						Url: "https://access.redhat.com/security/cve/CVE-2023-2253",
					},
					{
						Url: "https://bugzilla.redhat.com/show_bug.cgi?id=2189886",
					},
				},
			},
			{
				Name:        "RHSA-2024:4533",
				Message:     "RHSA-2024:4533 in github.com/something",
				Description: "A flaw was found in the `/v2/_catalog` endpoint in distribution/distribution, which accepts a parameter to control the maximum number of records returned (query string: `n`). This vulnerability allows a malicious user to submit an unreasonably large value for `n,` causing the allocation of a massive string array, possibly causing a denial of service through excessive use of memory.",
				Solution:    "Upgrade github.com/something from v2.8.1+incompatible to 2.8.2-beta.1",
				Severity:    "HIGH",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "rhsa", Name: "RHSA-2024:4533", Value: "RHSA-2024:4533", Url: "https://access.redhat.com/errata/RHSA-2024:4533",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "kube-system",
						Kind:          "Pod",
						Name:          "kube-apiserver-kind-control-plane",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "registry.k8s.io/kube-apiserver:v1.25.3",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/something",
						},
						Version: "v2.8.1+incompatible",
					},
					OperatingSystem: "debian 11.5",
				},
				Links: []*prototool.Link{
					{
						Url: "https://access.redhat.com/errata/RHSA-2024:4533",
					},
					{
						Url: "https://access.redhat.com/security/updates/classification/#important",
					},
				},
			},
			{
				Name:        "CVE-2022-23471",
				Message:     "CVE-2022-23471 in github.com/containerd/containerd",
				Description: "containerd is an open source container runtime. A bug was found in containerd's CRI implementation where a user can exhaust memory on the host. In the CRI stream server, a goroutine is launched to handle terminal resize events if a TTY is requested. If the user's process fails to launch due to, for example, a faulty command, the goroutine will be stuck waiting to send without a receiver, resulting in a memory leak. Kubernetes and crictl can both be configured to use containerd's CRI implementation and the stream server is used for handling container IO. This bug has been fixed in containerd 1.6.12 and 1.5.16.  Users should update to these versions to resolve the issue. Users unable to upgrade should ensure that only trusted images and commands are used and that only trusted users have permissions to execute commands in running containers. ",
				Solution:    "Upgrade github.com/containerd/containerd from v1.6.8 to 1.5.16, 1.6.12",
				Severity:    "MEDIUM",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "cve", Name: "CVE-2022-23471", Value: "CVE-2022-23471", Url: "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2022-23471",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "default",
						Kind:          "Pod",
						Name:          "lang-vuln-pod",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "public.ecr.aws/cloudwatch-agent/cloudwatch-agent:1.247360.0b252689",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/containerd/containerd",
						},
						Version: "v1.6.8",
					},
					OperatingSystem: "",
				},
				Links: []*prototool.Link{
					{Url: "https://avd.aquasec.com/nvd/cve-2022-23471"},
					{Url: "https://github.com/containerd/containerd"},
					{Url: "https://github.com/containerd/containerd/commit/241563be06a3de8b6a849414c4e805b68d3bb295"},
					{Url: "https://github.com/containerd/containerd/commit/a05d175400b1145e5e6a735a6710579d181e7fb0"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.5.16"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.6.12"},
					{Url: "https://github.com/containerd/containerd/security/advisories/GHSA-2qjp-425j-52j9"},
					{Url: "https://nvd.nist.gov/vuln/detail/CVE-2022-23471"},
					{Url: "https://security.gentoo.org/glsa/202401-31"},
					{Url: "https://ubuntu.com/security/notices/USN-5776-1"},
					{Url: "https://www.cve.org/CVERecord?id=CVE-2022-23471"},
				},
			},
			{
				Name:        "CVE-2023-25153",
				Message:     "CVE-2023-25153 in github.com/containerd/containerd",
				Description: "containerd is an open source container runtime. Before versions 1.6.18 and 1.5.18, when importing an OCI image, there was no limit on the number of bytes read for certain files. A maliciously crafted image with a large file where a limit was not applied could cause a denial of service. This bug has been fixed in containerd 1.6.18 and 1.5.18.  Users should update to these versions to resolve the issue. As a workaround, ensure that only trusted images are used and that only trusted users have permissions to import images.",
				Solution:    "Upgrade github.com/containerd/containerd from v1.6.8 to 1.5.18, 1.6.18",
				Severity:    "MEDIUM",
				Confidence:  "Unknown",
				Identifiers: []*prototool.Identifier{
					{
						Type: "cve", Name: "CVE-2023-25153", Value: "CVE-2023-25153", Url: "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-25153",
					},
				},
				Location: &prototool.Location{
					KubernetesResource: &prototool.KubernetesResource{
						Namespace:     "default",
						Kind:          "Pod",
						Name:          "lang-vuln-pod",
						ContainerName: "",
						AgentId:       "1234",
					},
					Image: "public.ecr.aws/cloudwatch-agent/cloudwatch-agent:1.247360.0b252689",
					Dependency: &prototool.Dependency{
						Package: &prototool.Package{
							Name: "github.com/containerd/containerd",
						},
						Version: "v1.6.8",
					},
					OperatingSystem: "",
				},
				Links: []*prototool.Link{
					{Url: "https://avd.aquasec.com/nvd/cve-2023-25153"},
					{Url: "https://access.redhat.com/security/cve/CVE-2023-25153"},
					{Url: "https://github.com/containerd/containerd"},
					{Url: "https://github.com/containerd/containerd/commit/0c314901076a74a7b797a545d2f462285fdbb8c4"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.5.18"},
					{Url: "https://github.com/containerd/containerd/releases/tag/v1.6.18"},
					{Url: "https://github.com/containerd/containerd/security/advisories/GHSA-259w-8hf6-59c2"},
					{Url: "https://nvd.nist.gov/vuln/detail/CVE-2023-25153"},
					{Url: "https://pkg.go.dev/vuln/GO-2023-1573"},
					{Url: "https://ubuntu.com/security/notices/USN-6202-1"},
					{Url: "https://www.cve.org/CVERecord?id=CVE-2023-25153"},
				},
			},
		},
	}
)
