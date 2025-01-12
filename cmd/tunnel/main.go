package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"

	"github.com/khulnasoft/tunnel-k8s-wrapper/configmap"
	"github.com/khulnasoft/tunnel-k8s-wrapper/converter"
	"github.com/khulnasoft/tunnel-k8s-wrapper/data/config"
	"github.com/khulnasoft/tunnel-k8s-wrapper/data/errorcodes"
	"github.com/khulnasoft/tunnel-k8s-wrapper/kube"
	"github.com/khulnasoft/tunnel-k8s-wrapper/logging"
	"github.com/khulnasoft/tunnel-k8s-wrapper/tunnel"
	"github.com/khulnasoft/tunnel-k8s-wrapper/validator"
)

var (
	githubAgentNS = &cli.StringFlag{
		Name:     "github-agent-ns",
		Usage:    "The namespace of the github agent",
		EnvVars:  []string{"GITHUB_AGENT_NS"},
		Required: true,
	}
	workloads = &cli.StringFlag{
		Name:     "workloads",
		Usage:    "The K8S workload kinds to be scanned for vulnerabilities",
		EnvVars:  []string{"WORKLOADS"},
		Required: true,
	}
	namespace = &cli.StringFlag{
		Name:     "namespace",
		Usage:    "The namespace to scan",
		EnvVars:  []string{"NAMESPACE"},
		Required: true,
	}
	githubAgentID = &cli.StringFlag{
		Name:     "github-agent-id",
		Usage:    "The github agent id",
		EnvVars:  []string{"GITHUB_AGENT_ID"},
		Required: true,
	}
	logLevel = &cli.StringFlag{
		Name:    "log-level",
		Usage:   "The log level (trace, debug, info, warn, error)",
		Value:   "info",
		EnvVars: []string{"LOG_LEVEL"},
	}
	timeout = &cli.DurationFlag{
		Name:     "timeout",
		Usage:    "The timeout of the Tunnel scan",
		Value:    time.Minute * 15, // 15 Minute default timeout
		EnvVars:  []string{"TIMEOUT"},
		Required: false,
	}
	reportMaxSize = &cli.IntFlag{
		Name:     "report-max-size",
		Usage:    "The maximum file size of the Tunnel Report in bytes",
		Value:    100000000, // 100MB
		EnvVars:  []string{"REPORT_MAX_SIZE"},
		Required: false,
	}
	tunnelJavaDB = &cli.StringSliceFlag{
		Name:     "java-db-repository",
		Usage:    "OCI repository(ies) to retrieve tunnel-java-db in order of priority",
		Value:    cli.NewStringSlice("registry.github.com/github-org/security-products/dependencies/tunnel-java-db:1", "ghcr.io/khulnasoft/tunnel-java-db:1"),
		EnvVars:  []string{"TUNNEL_JAVA_DB_REPOSITORY"},
		Required: false,
	}
)

func main() {
	app := &cli.App{
		Name:  "Tunnel K8S Wrapper",
		Usage: "Uses Tunnel scanner to scan K8S workloads",
		Commands: []*cli.Command{
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "scan workloads",
				Action: func(cCtx *cli.Context) error {
					githubAgentNS := cCtx.String("github-agent-ns")
					workloads := strings.ReplaceAll(cCtx.String("workloads"), " ", "")
					namespace := cCtx.String("namespace")
					githubAgentID := cCtx.String("github-agent-id")
					logLevel := cCtx.String("log-level")
					timeout := cCtx.Duration("timeout")
					reportMaxSize := cCtx.Uint64("report-max-size")
					logger := logging.NewLogger(logLevel)
					tunnelJavaDB := strings.Join(cCtx.StringSlice("java-db-repository"), ",")
					defer logger.Sync()

					ctx, cancel := context.WithTimeout(context.Background(), timeout)
					defer cancel()

					logger.Info("Tunnel wrapper image initialized with",
						zap.String("github agent namespace", githubAgentNS),
						zap.String("workloads", workloads),
						zap.String("namespace", namespace),
						zap.String("github agend id", githubAgentID),
						zap.String("timeout", timeout.String()),
						zap.Uint64("report-max-size", reportMaxSize),
						zap.String("java-db-repository", tunnelJavaDB))

					kubeClient, err := kube.NewClient()
					if err != nil {
						return logAndExit(logger, "Kube client", err, errorcodes.KubeClient)
					}

					// Validate flags
					validator := validator.Flags{
						KubeClient:           kubeClient,
						GithubAgentNamespace: githubAgentNS,
						GithubAgentID:        githubAgentID,
						Workloads:            workloads,
						Namespace:            namespace,
					}
					if err := validator.Check(ctx); err != nil {
						return logAndExit(logger, "Flags validation", err, errorcodes.FlagsValidation)
					}

					// Get tunnel version
					scanner := tunnel.New(logger)
					version, stdErr, err := scanner.Version()
					if err != nil {
						return logStdAndExit(logger, "Tunnel scanner version", err, errorcodes.TunnelVersion, stdErr)
					}
					logger.Info("Tunnel version information", zap.String("version", version))

					// Perform scan
					if stdErr, err = scanner.Scan(workloads, namespace, timeout, tunnelJavaDB); err != nil {
						return logStdAndExit(logger, "Tunnel scan", err, errorcodes.TunnelScan, stdErr)
					}

					// Read the Tunnel report
					tunnelReport, err := tunnel.NewReport(tunnel.ReportFileName, reportMaxSize)
					if err != nil {
						if errors.Is(err, tunnel.ErrSizeLimit) {
							return logAndExit(logger, "Tunnel report size limit exceeded", err, errorcodes.SizeLimit)
						}

						return logAndExit(logger, "Tunnel report", err, errorcodes.DataConvertion)
					}

					report, err := tunnelReport.ToReport()
					if err != nil {
						return logAndExit(logger, "Unmarshaling Tunnel report", err, errorcodes.ToReport)
					}

					config := config.Configuration{
						Logger:              logger,
						KubeClient:          kubeClient,
						TargetNamespace:     namespace,
						AgentID:             githubAgentID,
						AgentNamespace:      githubAgentNS,
						TunnelScannerVersion: version,
					}
					c := converter.NewReportConverter(&config, report)

					base64Bytes, err := c.PrepareData()
					if err != nil {
						return logAndExit(logger, "Preparing data", err, errorcodes.PrepareData)
					}

					// Store payload in a chained config map
					maxConfigSizeBytes := 990000 // 990KB. Configmap can have up to 1MB
					cmCreator := configmap.NewChainCreator(&config, maxConfigSizeBytes)
					if err = cmCreator.CreateChainedConfigMaps(ctx, base64Bytes); err != nil {
						return logAndExit(logger, "Chained configmaps", err, errorcodes.ChainedConfigmaps)
					}

					return nil
				},
				Flags: []cli.Flag{
					githubAgentNS,
					workloads,
					namespace,
					githubAgentID,
					logLevel,
					timeout,
					reportMaxSize,
					tunnelJavaDB,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func logAndExit(logger *zap.Logger, msg string, err error, code int) error {
	logger.Error(msg, zap.Error(err))
	return cli.Exit(fmt.Sprintf("%v: %v", msg, err), code)
}

func logStdAndExit(logger *zap.Logger, msg string, err error, code int, stdErr string) error {
	logger.Error(msg, zap.Error(err))
	if len(stdErr) != 0 {
		logger.Error("", zap.String("stderr", stdErr), zap.Error(err))
	}
	return cli.Exit(fmt.Sprintf("%v: %v", msg, err), code)
}
