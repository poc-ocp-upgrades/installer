package exec

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/logutils"
	"github.com/hashicorp/terraform/command"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/mitchellh/cli"
)

type cmdFunc func(command.Meta) cli.Command

var commands = map[string]cmdFunc{"apply": func(meta command.Meta) cli.Command {
	return &command.ApplyCommand{Meta: meta}
}, "destroy": func(meta command.Meta) cli.Command {
	return &command.ApplyCommand{Meta: meta, Destroy: true}
}, "init": func(meta command.Meta) cli.Command {
	return &command.InitCommand{Meta: meta}
}}

func runner(cmd string, dir string, args []string, stdout, stderr io.Writer) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lf := ioutil.Discard
	if level := logging.LogLevel(); level != "" {
		lf = &logutils.LevelFilter{Levels: logging.ValidLevels, MinLevel: logutils.LogLevel(level), Writer: stdout}
	}
	log.SetOutput(lf)
	defer log.SetOutput(os.Stderr)
	defer plugin.CleanupClients()
	sdCh, cancel := makeShutdownCh()
	defer cancel()
	pluginDirs, err := globalPluginDirs(dir)
	if err != nil {
		fmt.Fprintf(stderr, "Error discovering plugin directories for Terraform: %v", err)
		return 1
	}
	meta := command.Meta{Color: false, GlobalPluginDirs: pluginDirs, Ui: &cli.BasicUi{Writer: stdout, ErrorWriter: stderr}, OverrideDataDir: dir, ShutdownCh: sdCh}
	f := commands[cmd]
	oldStderr := os.Stderr
	outR, outW, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(stderr, "error creating Pipe: %v", err)
		return 1
	}
	os.Stderr = outW
	go func() {
		scanner := bufio.NewScanner(outR)
		for scanner.Scan() {
			fmt.Fprintf(lf, "%s\n", scanner.Bytes())
		}
	}()
	defer func() {
		outW.Close()
		os.Stderr = oldStderr
	}()
	return f(meta).Run(args)
}
func Apply(datadir string, args []string, stdout, stderr io.Writer) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return runner("apply", datadir, args, stdout, stderr)
}
func Destroy(datadir string, args []string, stdout, stderr io.Writer) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return runner("destroy", datadir, args, stdout, stderr)
}
func Init(datadir string, args []string, stdout, stderr io.Writer) int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return runner("init", datadir, args, stdout, stderr)
}
func makeShutdownCh() (<-chan struct{}, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	resultCh := make(chan struct{})
	signalCh := make(chan os.Signal, 4)
	handle := []os.Signal{}
	handle = append(handle, ignoreSignals...)
	handle = append(handle, forwardSignals...)
	signal.Notify(signalCh, handle...)
	go func() {
		for {
			<-signalCh
			resultCh <- struct{}{}
		}
	}()
	return resultCh, func() {
		signal.Reset(handle...)
	}
}
