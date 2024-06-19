package module

import (
	"YudoleChatServer/packages/tab"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"strings"
)

//github.com/WebmiRU/YudolePlatformPackages/module

type Resource struct {
	Path string `json:"path"`
	Type string `json:"type"`
}

type Module struct {
	//Key      string                                  `json:"key"`
	Autostart bool               `json:"autostart"`
	Name      string             `json:"name"`
	Service   string             `json:"service"`
	Type      string             `json:"type"`
	Command   string             `json:"command"`
	Tabs      map[string]tab.Tab `json:"tabs"`
	Exec      *exec.Cmd          `json:"-"`
	State     string             `json:"proc_state"`
	Resources []Resource         `json:"resources"`

	dir        string `json:"dir"`
	isRunning  bool
	configPath string
}

func (m *Module) Load(configPath string) error {
	m.Command = m.Command
	m.configPath = configPath + string(os.PathSeparator) + "module.json"
	configBytes, _ := os.ReadFile(m.configPath)
	m.dir = configPath

	if err := json.Unmarshal(configBytes, &m); err != nil {
		return err
	}

	//if m.Autostart && m.Exec != nil {
	//	// Run module
	//}

	return nil
}

func (m *Module) Save() error {
	if data, err := json.MarshalIndent(m, "", "  "); err != nil {
		return err
	} else {
		if err := os.WriteFile(m.configPath, data, 0666); err != nil {
			return err
		}
	}

	return nil
}

func (m *Module) Start() error {
	if len(m.Command) > 0 && !m.isRunning {
		command := m.Command

		if len(m.Command) >= 2 && m.Command[0:2] == "./" {
			command = m.dir + string(os.PathSeparator) + strings.Replace(m.Command, "./", "", 1)
		}

		m.Exec = exec.Command(command)
		m.Exec.Dir = m.dir
	} else {
		return nil
	}

	m.isRunning = true
	if err := m.Exec.Start(); err != nil {
		m.State = "failed"
		return err
	}

	m.State = "run"

	go func(cmd *exec.Cmd) {
		if err := cmd.Wait(); err == nil {
			m.State = "stopped"
			log.Printf("Module %s stopped", cmd.Path)
		} else {
			m.State = "failed"
			log.Printf("Module %s failed: %s", cmd.Path, err)
		}

		m.isRunning = false
	}(m.Exec)

	return nil
}

func (m *Module) Stop() error {
	if err := m.Exec.Process.Kill(); err != nil {
		return err
	}

	return nil
}

func (m *Module) StopWait() (int, error) {
	if err := m.Exec.Process.Kill(); err != nil {
		return 0, err
	}

	wait, err := m.Exec.Process.Wait()
	m.isRunning = false

	if err != nil {
		return 0, err
	}

	return wait.ExitCode(), nil
}

func (m *Module) RestartWait() (int, error) {
	exitCode := 0

	if exCode, err := m.StopWait(); err != nil {
		exitCode = exCode
		return exCode, err
	}

	if err := m.Start(); err != nil {
		return 0, err
	}

	return exitCode, nil
}

func (m *Module) Restart() error {
	if err := m.Stop(); err != nil {
		return err
	}

	if err := m.Start(); err != nil {
		return err
	}

	return nil
}
