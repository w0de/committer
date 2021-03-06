package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidYamlString(t *testing.T) {
	var data = []byte(`
tasks:
  - name: 'SCSS Lint'
    command: yarn scss-lint
    files: ".scss"
`)

	config, err := NewConfig(data)

	assert.Nil(t, err, "The error should not be present.")
	assert.Equal(t, len(config.Tasks), 1, "There should only be one task.")
	assert.Equal(t, config.Tasks[0].Name, "SCSS Lint", "The name should be set correctly")
	assert.Equal(t, config.Tasks[0].Command, "yarn scss-lint", "The command should be set correctly")
}

func TestInvalidYamlString(t *testing.T) {
	var invalidData = []byte(`
tasks:
  name: 'SCSS Lint'
    command: yarn scss-lint
`)

	config, err := NewConfig(invalidData)

	assert.Nil(t, config, "The config should be nil.")
	assert.NotNil(t, err, "The error should be present.")
}

func TestValidYamlFile(t *testing.T) {
	config, err := NewConfigFromFile("../test/fixtures/valid.yml")

	assert.Equal(t, len(config.Tasks), 1, "There should only be one task.")
	assert.Equal(t, config.Tasks[0].Name, "SCSS Lint", "The name should be set correctly")
	assert.Equal(t, config.Tasks[0].Command, "yarn scss-lint", "The command should be set correctly")
	assert.Nil(t, err, "The error should not be present.")
}

func TestInvalidYamlFile(t *testing.T) {
	config, err := NewConfigFromFile("../test/fixtures/invalid.yml")

	assert.Nil(t, config, "The config should be nil.")
	assert.NotNil(t, err, "The error should be present.")
}

func TestMissingYamlFile(t *testing.T) {
	config, err := NewConfigFromFile("../test/fixtures/doesnotexist.yml")

	assert.Nil(t, config, "The config should be nil.")
	assert.NotNil(t, err, "The error should be present.")
}

func TestCompleteYaml(t *testing.T) {
	var data = []byte(`
tasks:
  - name: 'SCSS Lint'
    command: yarn scss-lint
    files: '\.scss|\.css'
    fix:
      command: yarn scss-lint-fix
      output: Completed
      autostage: true
`)

	config, err := NewConfig(data)

	assert.Equal(t, len(config.Tasks), 1, "There should only be one task.")
	assert.Equal(t, config.Tasks[0].Name, "SCSS Lint", "The name should be set correctly")
	assert.Equal(t, config.Tasks[0].Command, "yarn scss-lint", "The command should be set correctly")
	assert.Equal(t, config.Tasks[0].Fix.Command, "yarn scss-lint-fix", "The fix command should be set correctly")
	assert.Equal(t, config.Tasks[0].Fix.Output, "Completed", "The fix output should be set correctly")
	assert.True(t, config.Tasks[0].Fix.Autostage, "The fix autostage should be set correctly")
	assert.Nil(t, err, "The error should not be present.")

}

func TestTasklessCommandYaml(t *testing.T) {
	var data = []byte(`
tasks:
`)

	_, err := NewConfig(data)

	assert.NotNil(t, err, "The error should be present.")
}

func TestNamelessCommandYaml(t *testing.T) {
	var data = []byte(`
tasks:
  -
    command: yarn scss-lint
    files: '\.scss|\.css'
    fix:
      command: yarn scss-lint-fix
      output: Completed
`)

	_, err := NewConfig(data)

	assert.NotNil(t, err, "The error should be present.")
}

func TestCommandlessYaml(t *testing.T) {
	var data = []byte(`
tasks:
  -
    name: 'SCSS Lint'
    files: '\.scss|\.css'
    fix:
      command: yarn scss-lint-fix
      output: Completed
`)

	_, err := NewConfig(data)

	assert.NotNil(t, err, "The error should be present.")
}

func TestWithoutFixOutputYaml(t *testing.T) {
	var data = []byte(`
tasks:
  -
    name: 'SCSS Lint'
    command: yarn scss-lint
    files: '\.scss|\.css'
    fix:
      command: yarn scss-lint-fix
`)

	_, err := NewConfig(data)

	assert.NotNil(t, err, "The error should be present.")
}

func TestWithoutFixFilesYaml(t *testing.T) {
	var data = []byte(`
tasks:
  -
    name: 'SCSS Lint'
    command: yarn scss-lint
    fix:
      command: yarn scss-lint-fix
      output: Completed
`)

	_, err := NewConfig(data)

	assert.NotNil(t, err, "The error should be present.")
}
