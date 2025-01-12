package errorcodes

const (
	// KubeClient is the exit code when tunnel-k8s-wrapper fails to initialize a kube client
	KubeClient = 3
	// FlagsValidation is the exit code when tunnel-k8s-wrapper fails to validate the input flags
	FlagsValidation = 4
	// TunnelVersion is the exit code when tunnel-k8s-wrapper fails to get the Tunnel version
	TunnelVersion = 5
	// TunnelScan is the exit code when tunnel-k8s-wrapper fails to make a Tunnel scan
	TunnelScan = 6
	// SizeLimit is the exit code when tunnel-k8s-wrapper fails because the Tunnel report is bigger than the limit
	SizeLimit = 7
	// DataConvertion is the exit code when tunnel-k8s-wrapper fails to get a data converter
	DataConvertion = 8
	// ToReport is the exit code when tunnel-k8s-wrapper fails to unmarshal the Tunnel report
	ToReport = 9
	// PrepareData is the exit code when tunnel-k8s-wrapper fails to prepare the data for the configmaps
	PrepareData = 10
	// ChainedConfigmaps is the exit code when tunnel-k8s-wrapper fails to create the chained configmaps
	ChainedConfigmaps = 11
)
