package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/openshift/installer/pkg/terraform/exec/plugins"
)

var (
	rootOpts struct {
		dir		string
		logLevel	string
	}
)

func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.CommandLine.Parse([]string{})
	flag.CommandLine.Set("stderrthreshold", "4")
	if len(os.Args) > 0 {
		base := filepath.Base(os.Args[0])
		cname := strings.TrimSuffix(base, filepath.Ext(base))
		if pluginRunner, ok := plugins.KnownPlugins[cname]; ok {
			pluginRunner()
			return
		}
	}
	installerMain()
}
func installerMain() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	rootCmd := newRootCmd()
	for _, subCmd := range []*cobra.Command{newCreateCmd(), newDestroyCmd(), newWaitForCmd(), newGatherCmd(), newVersionCmd(), newGraphCmd(), newCompletionCmd()} {
		rootCmd.AddCommand(subCmd)
	}
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Error executing openshift-install: %v", err)
	}
}
func newRootCmd() *cobra.Command {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cmd := &cobra.Command{Use: "openshift-install", Short: "Creates OpenShift clusters", Long: "", PersistentPreRun: runRootCmd, SilenceErrors: true, SilenceUsage: true}
	cmd.PersistentFlags().StringVar(&rootOpts.dir, "dir", ".", "assets directory")
	cmd.PersistentFlags().StringVar(&rootOpts.logLevel, "log-level", "info", "log level (e.g. \"debug | info | warn | error\")")
	return cmd
}
func runRootCmd(cmd *cobra.Command, args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	logrus.SetOutput(ioutil.Discard)
	logrus.SetLevel(logrus.TraceLevel)
	level, err := logrus.ParseLevel(rootOpts.logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.AddHook(newFileHook(os.Stderr, level, &logrus.TextFormatter{ForceColors: terminal.IsTerminal(int(os.Stderr.Fd())), DisableTimestamp: true, DisableLevelTruncation: true}))
	if err != nil {
		logrus.Fatal(errors.Wrap(err, "invalid log-level"))
	}
}
