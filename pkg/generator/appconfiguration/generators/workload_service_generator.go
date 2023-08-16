package generators

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kusionstack.io/kusion/pkg/models"
	"kusionstack.io/kusion/pkg/models/appconfiguration/workload"
)

// workloadServiceGenerator is a struct for generating service
// workload resources.
type workloadServiceGenerator struct {
	projectName string
	appName     string
	service     *workload.Service
}

// NewWorkloadServiceGenerator returns a new workloadServiceGenerator
// instance.
func NewWorkloadServiceGenerator(
	projectName string,
	appName string,
	service *workload.Service,
) (Generator, error) {
	if len(projectName) == 0 {
		return nil, fmt.Errorf("project name must not be empty")
	}

	if len(appName) == 0 {
		return nil, fmt.Errorf("app name must not be empty")
	}

	if service == nil {
		return nil, fmt.Errorf("service workload must not be nil")
	}

	return &workloadServiceGenerator{
		projectName: projectName,
		appName:     appName,
		service:     service,
	}, nil
}

// NewWorkloadServiceGeneratorFunc returns a new NewGeneratorFunc that
// returns a workloadServiceGenerator instance.
func NewWorkloadServiceGeneratorFunc(
	projectName string,
	appName string,
	service *workload.Service,
) NewGeneratorFunc {
	return func() (Generator, error) {
		return NewWorkloadServiceGenerator(projectName, appName, service)
	}
}

// Generate generates a service workload resource to the given spec.
func (g *workloadServiceGenerator) Generate(spec *models.Spec) error {
	lrs := g.service
	if lrs == nil {
		return nil
	}

	// Create an empty resource slice if it doesn't exist yet.
	if spec.Resources == nil {
		spec.Resources = make(models.Resources, 0)
	}

	// Create a slice of containers based on the app's
	// containers.
	containers, err := toOrderedContainers(lrs.Containers)
	if err != nil {
		return err
	}

	// Create a Deployment object based on the app's
	// configuration.
	resource := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: appsv1.SchemeGroupVersion.String(),
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Labels: mergeMaps(
				uniqueAppLabels(g.projectName, g.appName),
				g.service.Labels,
			),
			Annotations: mergeMaps(
				g.service.Annotations,
			),
			Name:      uniqueAppName(g.projectName, g.appName),
			Namespace: g.projectName,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: IntPtr(int32(lrs.Replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: uniqueAppLabels(g.projectName, g.appName),
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: mergeMaps(
						uniqueAppLabels(g.projectName, g.appName),
						g.service.Labels,
					),
					Annotations: mergeMaps(
						g.service.Annotations,
					),
				},
				Spec: v1.PodSpec{
					Containers: containers,
				},
			},
		},
	}

	// Add the Deployment resource to the spec.
	return appendToSpec(
		kubernetesResourceID(resource.TypeMeta, resource.ObjectMeta),
		resource,
		spec,
	)
}