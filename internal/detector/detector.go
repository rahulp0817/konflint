package detector

import (
	"bytes"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type FileType string

const (
	KubernetesDeployment FileType = "kubernetes-deployment"
	KubernetesService    FileType = "kubernetes-service"
	KubernetesIngress    FileType = "kubernetes-ingress"
	KubernetesConfigMap  FileType = "kubernetes-configmap"
	DockerCompose        FileType = "docker-compose"
	GitHubActions        FileType = "github-actions"
	GenericYAML          FileType = "generic-yaml"
	Unknown              FileType = "unknown"
)

// Labels shown in CLI output
var Labels = map[FileType]string{
	KubernetesDeployment: "Kubernetes Deployment (apps/v1)",
	KubernetesService:    "Kubernetes Service",
	KubernetesIngress:    "Kubernetes Ingress",
	KubernetesConfigMap:  "Kubernetes ConfigMap",
	DockerCompose:        "Docker Compose",
	GitHubActions:        "GitHub Actions Workflow",
	GenericYAML:          "Generic YAML",
}

// Kubernetes minimal struct
type K8sObject struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
}

// Detect figures out the file type using path + content
func Detect(filePath string, content []byte) FileType {

	// Layer 1 — Path based detection
	if isGitHubActions(filePath) {
		return GitHubActions
	}

	if isDockerCompose(filePath) {
		return DockerCompose
	}

	// Layer 2 — Kubernetes detection via YAML parsing
	if ft := detectKubernetes(content); ft != Unknown {
		return ft
	}

	// Layer 3 — fallback generic YAML
	ext := strings.ToLower(filepath.Ext(filePath))
	if ext == ".yaml" || ext == ".yml" {
		return GenericYAML
	}

	return Unknown
}

// Detect Kubernetes resources using YAML parser
func detectKubernetes(content []byte) FileType {

	decoder := yaml.NewDecoder(bytes.NewReader(content))

	for {
		var obj K8sObject
		err := decoder.Decode(&obj)

		if err != nil {
			break
		}

		switch obj.Kind {

		case "Deployment":
			return KubernetesDeployment

		case "Service":
			return KubernetesService

		case "Ingress":
			return KubernetesIngress

		case "ConfigMap":
			return KubernetesConfigMap
		}
	}

	return Unknown
}

func isGitHubActions(path string) bool {
	return strings.Contains(path, ".github/workflows")
}

func isDockerCompose(path string) bool {
	base := strings.ToLower(filepath.Base(path))

	return base == "docker-compose.yml" ||
		base == "docker-compose.yaml" ||
		strings.HasPrefix(base, "docker-compose.")
}

// Label returns human-readable name
func Label(ft FileType) string {
	if label, ok := Labels[ft]; ok {
		return label
	}
	return "Unknown"
}
