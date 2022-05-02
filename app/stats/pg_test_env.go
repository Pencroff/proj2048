package stats

import (
	"github.com/pkg/errors"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type ComposeTestEnv struct {
	composeFilePaths []string
	identifier       string
	env              map[string]string
	serviceName      string
	strategy         wait.Strategy
	Cmps             *tc.LocalDockerCompose
}

func NewComposeTestEnv(composeFilePaths []string) *ComposeTestEnv {
	identifier := "compose_test_env"
	env := map[string]string{}
	return &ComposeTestEnv{composeFilePaths, identifier, env, "", nil, nil}
}

func (e *ComposeTestEnv) SetEnv(m map[string]string) {
	e.env = make(map[string]string)
	for k, v := range m {
		e.env[k] = v
	}
}

func (e *ComposeTestEnv) SetStrategy(name string, strategy wait.Strategy) {
	e.serviceName = name
	e.strategy = strategy
}

func (e *ComposeTestEnv) Up() error {
	e.Cmps = tc.NewLocalDockerCompose(e.composeFilePaths, e.identifier)
	c := e.Cmps.WithCommand([]string{"up", "-d"})
	if len(e.env) > 0 {
		c = c.WithEnv(e.env)
	}
	if e.serviceName != "" {
		c = c.WaitForService(e.serviceName, e.strategy)
	}
	execError := c.Invoke()

	err := execError.Error
	if err != nil {
		return errors.Errorf("Could not UP cmps file: %v - %v", e.composeFilePaths, err)
	}
	return nil
}

func (e *ComposeTestEnv) Down() error {
	execError := e.Cmps.Down()
	err := execError.Error
	if err != nil {
		return errors.Errorf("Could not DOWN cmps file: %v - %v", e.composeFilePaths, err)
	}
	return nil
}
