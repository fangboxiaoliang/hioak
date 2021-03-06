// Copyright 2018 John Deng (hi.devops.io@gmail.com).
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openshift

import (
	"github.com/openshift/client-go/apps/clientset/versioned/fake"
	"github.com/stretchr/testify/assert"
	"hidevops.io/hiboot/pkg/log"
	"hidevops.io/hiboot/pkg/system"
	"hidevops.io/hioak/starter"
	"testing"
)

func TestDeploymentConfigCreation(t *testing.T) {
	log.Debug("TestDeploymentConfigCreation()")
	clientSet := fake.NewSimpleClientset().AppsV1()
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "demo-consumer"
	healthEndPoint := "http://localhost:8080/health"
	version := "v1"
	env := []system.Env{
		{
			Name:  "SPRING_PROFILES_ACTIVE",
			Value: profile,
		},
		{
			Name:  "APP_OPTIONS",
			Value: "-Xms128m -Xmx512m -Xss512k -XX:+ExitOnOutOfMemoryError",
		},
		{
			Name:  "TZ",
			Value: "Asia/Shanghai",
		},
	}

	ports := []orch.Ports{
		{
			ContainerPort: 8080,
			Protocol:      "TCP",
		},
		{
			ContainerPort: 7575,
			Protocol:      "TCP",
		},
	}

	// new dc instance
	dc := newDeploymentConfig(clientSet)
	// create dc
	fullName := app + "-" + version
	err := dc.Create(app, namespace, fullName, version, &env, map[string]string{}, &ports, 1, false, healthEndPoint, "")
	assert.Equal(t, nil, err)
}

func TestDeploymentConfigInstantiation(t *testing.T) {
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "demo-consumer"
	healthEndPoint := "http://localhost:8080/health"
	version := "v1"
	env := []system.Env{
		{
			Name:  "SPRING_PROFILES_ACTIVE",
			Value: profile,
		},
		{
			Name:  "APP_OPTIONS",
			Value: "-Xms128m -Xmx512m -Xss512k -XX:+ExitOnOutOfMemoryError",
		},
		{
			Name:  "TZ",
			Value: "Asia/Shanghai",
		},
	}

	ports := []orch.Ports{
		{
			ContainerPort: 8080,
			Protocol:      "TCP",
		},
		{
			ContainerPort: 7575,
			Protocol:      "TCP",
		},
	}
	log.Debug("TestDeploymentConfigInstantiation()")
	clientSet := fake.NewSimpleClientset().AppsV1()
	dc := newDeploymentConfig(clientSet)
	err := dc.Create(app, namespace, fullName, version, &env, map[string]string{}, &ports, 1, false, healthEndPoint, "")
	assert.Equal(t, nil, err)
}

func TestDeploymentConfig(t *testing.T) {
	projectName := "demo"
	profile := "dev"
	namespace := projectName + "-" + profile
	app := "demo-consumer"
	healthEndPoint := "http://localhost:8080/health"
	version := "v1"
	env := []system.Env{
		{
			Name:  "SPRING_PROFILES_ACTIVE",
			Value: profile,
		},
		{
			Name:  "APP_OPTIONS",
			Value: "-Xms128m -Xmx512m -Xss512k -XX:+ExitOnOutOfMemoryError",
		},
		{
			Name:  "TZ",
			Value: "Asia/Shanghai",
		},
	}

	ports := []orch.Ports{
		{
			ContainerPort: 8080,
			Protocol:      "TCP",
		},
		{
			ContainerPort: 7575,
			Protocol:      "TCP",
		},
	}
	log.Debug("TestDeploymentConfigDeletion()")
	clientSet := fake.NewSimpleClientset().AppsV1()
	dc := newDeploymentConfig(clientSet)
	err := dc.Create(app, namespace, fullName, version, &env, map[string]string{}, &ports, 1, false, healthEndPoint, "")
	assert.Equal(t, nil, err)
	err = dc.Delete(namespace, fullName)
	assert.Equal(t, nil, err)
}
