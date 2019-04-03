/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	// ensure libs have a chance to globally register their flags
	"fmt"
	"io"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
	_ "k8s.io/klog"
)

// Run runs it
func Run() error {

	// read in yaml config file
	config, err := ParseFile()

	if err != nil {
		fmt.Printf("this error: %s\n", err)
		return nil
	}

	fmt.Printf("config: {%s}\n\n", config)

	// set up the filesystem
	// TODO: do this

	// build the args
	args := []string{
		fmt.Sprintf("--version='%s'", config.Version),
		fmt.Sprintf("--iteration='%s'", "0"),
		fmt.Sprintf("--url='%s'", "https://kubernetes.io"),
		fmt.Sprintf("--name='%s'", config.Name),
		fmt.Sprintf("--description='%s'", config.Description),
	}

	for _, d := range config.Depends {
		args = append(args, fmt.Sprintf("--depends='%s'", d))
	}

	//TODO: dir might not be best choice.
	//TODO: needs to support debs too
	args = append(args,
		"-s dir",
		"-t rpm",
		".", //path
	)

	fmt.Printf("args: %s\n\n", args)

	// do the things
	cmd := exec.Command("fpm", args...)

	cmd.Dir = "."
	_, err = cmd.CombinedOutput()
	// profit

	return err
}

// ParseFile decodes YAML data from a file in current directory a configuration struct
func ParseFile() (config Config, err error) {
	var file *os.File
	file, err = os.Open("config.yaml")
	if err != nil {
		return
	}
	defer file.Close()
	return Parse(file)
}

// Parse decodes YAML data from an io.Reader into a configuration struct
func Parse(in io.Reader) (config Config, err error) {
	dec := yaml.NewDecoder(in)
	dec.SetStrict(true)
	if err = dec.Decode(&config); err != nil {
		return
	}

	config.Info.Version = os.ExpandEnv(config.Info.Version)
	return config, nil
}
