package process

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
)

type darwin platform

var (
	defaultConfDir = path.Join(os.Getenv("HOME"), "/Library/LaunchAgents")
)

func init() {
	exeDir = path.Join(os.Getenv("HOME"), "/tmp/hailo/bin")
}

func newPlatform() *darwin {
	return &darwin{
		InitCmd: "launchctl",
		Config: config{
			Directory: getConfDir(),
			Extension: ".plist",
		},
	}
}

func (env *darwin) Install(serviceName string, serviceVersion, noFileSoftLimit, noFileHardLimit uint64) error {
	templateText := `
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">

<!-- Auto-generated by the provisioning service at {{.GeneratedAt}} -->
<!-- Description: {{.Description}} -->
<!-- Author:      {{.Author}} -->

<plist version="1.0">
<dict>
    <key>Label</key>
    <string>{{.Description}}</string>

    <key>RunAtLoad</key>
    <true/>

    {{if .Environment}}
    <key>EnvironmentVariables</key>
      <dict>
      {{range $key, $val := .Environment}}
        <key>{{$key}}</key>
        <string>{{$val}}</string>
      {{end}}
      </dict>
    {{end}}

    <key>UserName</key>
    <string>{{.RunAsUser}}</string>

    <key>GroupName</key>
    <string>{{.RunAsGroup}}</string>

    <key>Program</key>
    <string>{{.ProcessName}}</string>

    <key>StandardOutPath</key>
    <string>/tmp/{{.Description}}-console.log</string>

    <key>StandardErrorPath</key>
    <string>/tmp/{{.Description}}-error.log</string>
</dict>
</plist>
`
	tmpl, err := template.New("launchd").Parse(templateText)
	if err != nil {
		return err
	}

	return install(serviceName, serviceVersion, noFileSoftLimit, noFileHardLimit, env.Config, tmpl)
}

func (env *darwin) List(matching string) ([]string, error) {
	cmd := exec.Command(env.InitCmd, "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(out.Bytes()))
	processes := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if parts[0] != "-" && (len(matching) == 0 || strings.Contains(parts[2], matching)) {
			processes = append(processes, parts[2])
		}
	}

	return processes, nil
}

func (env *darwin) Start(serviceName string, serviceVersion, noFileSoftLimit, noFileHardLimit uint64) error {
	if err := env.Install(serviceName, serviceVersion, noFileSoftLimit, noFileHardLimit); err != nil {
		return err
	}
	cmdName := combineNameVersion(serviceName, serviceVersion)
	confPath := getConfPath(serviceName, serviceVersion, env.Config)
	if err := run(env.InitCmd, "load", confPath); err != nil {
		return fmt.Errorf("Tried to load %s: %v", cmdName, err)
	}

	if err := run(env.InitCmd, "start", cmdName); err != nil {
		return fmt.Errorf("Tried to start %s: %v", cmdName, err)
	}
	return nil
}

func (env *darwin) Stop(serviceName string, serviceVersion uint64) error {
	cmdName := combineNameVersion(serviceName, serviceVersion)
	confPath := getConfPath(serviceName, serviceVersion, env.Config)
	if err := run(env.InitCmd, "stop", cmdName); err != nil {
		return fmt.Errorf("Tried to stop %s: %v", cmdName, err)
	}

	if err := run(env.InitCmd, "unload", confPath); err != nil {
		return fmt.Errorf("Tried to unload %s: %v", cmdName, err)
	}

	if err := env.Uninstall(serviceName, serviceVersion); err != nil {
		return err
	}

	return nil
}

func (env *darwin) Restart(serviceName string, serviceVersion uint64) error {
	// launchctl does not support restart so stop and start.
	if err := env.Stop(serviceName, serviceVersion); err != nil {
		return err
	}

	if err := env.Start(serviceName, serviceVersion, 1024, 1024); err != nil {
		return err
	}
	return nil
}

func (env *darwin) Uninstall(serviceName string, serviceVersion uint64) error {
	return uninstall(serviceName, serviceVersion, env.Config)
}
