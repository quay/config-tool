package buildmanager

import (
	"fmt"

	"github.com/creasty/defaults"
	"github.com/quay/config-tool/pkg/lib/shared"
	"gopkg.in/yaml.v3"
)

// BuildManagerFieldGroup represents the BuildManager config fields
type BuildManagerFieldGroup struct {
	FeatureBuildSupport bool                    `default:""  json:"FEATURE_BUILD_SUPPORT" yaml:"FEATURE_BUILD_SUPPORT"`
	BuildManagerConfig  *BuildManagerDefinition `default:""  json:"BUILD_MANAGER,omitempty" yaml:"BUILD_MANAGER,omitempty"`
}

// BuildManagerDefinition represents a single storage configuration as a tuple (Name, Arguments)
type BuildManagerDefinition struct {
	Name string            `default:""  json:",inline" yaml:",inline"`
	Args *BuildManagerArgs `default:""  json:",inline" yaml:",inline"`
}

// BuildManagerArgs represents the arguments in the second value of a definition tuple
type BuildManagerArgs struct {
	// Args for ephemeral
	AllowedWorkerCount *shared.IntOrString `default:""  json:"ALLOWED_WORKER_COUNT,omitempty" yaml:"ALLOWED_WORKER_COUNT,omitempty"`
	OrchestratorPrefix string              `default:""  json:"ORCHESTRATOR_PREFIX,omitempty" yaml:"ORCHESTRATOR_PREFIX,omitempty"`
	Orchestrator       *OrchestratorArgs   `default:""  json:"ORCHESTRATOR,omitempty" yaml:"ORCHESTRATOR,omitempty"`
	Executors          []*ExecutorArgs     `default:""  json:"EXECUTORS,omitempty" yaml:"EXECUTORS,omitempty"`
}

// OrchestratorArgs represents the arguments in the orchestrator object
type OrchestratorArgs struct {
	RedisHost                   string `default:""  json:"REDIS_HOST,omitempty" yaml:"REDIS_HOST,omitempty"`
	RedisPassword               string `default:""  json:"REDIS_PASSWORD,omitempty" yaml:"REDIS_PASSWORD,omitempty"`
	RedisSSL                    bool   `default:""  json:"REDIS_SSL" yaml:"REDIS_SSL"`
	RedisSkipKeyspaceEventSetup bool   `default:""  json:"REDIS_SKIP_KEYSPACE_EVENT_SETUP" yaml:"REDIS_SKIP_KEYSPACE_EVENT_SETUP"`
}

// ExecutorArgs represents the arguments in an executor object
type ExecutorArgs struct {
	Executor                string              `default:""  json:"EXECUTOR,omitempty" yaml:"EXECUTOR,omitempty"`
	BuilderNamespace        string              `default:""  json:"BUILDER_NAMESPACE,omitempty" yaml:"BUILDER_NAMESPACE,omitempty"`
	K8sAPIServer            string              `default:""  json:"K8S_API_SERVER,omitempty" yaml:"K8S_API_SERVER,omitempty"`
	VolumeSize              string              `default:""  json:"VOLUME_SIZE,omitempty" yaml:"VOLUME_SIZE,omitempty"`
	KubernetesDistribution  string              `default:""  json:"KUBERNETES_DISTRIBUTION,omitempty" yaml:"KUBERNETES_DISTRIBUTION,omitempty"`
	ContainerMemoryLimits   string              `default:""  json:"CONTAINER_MEMORY_LIMITS,omitempty" yaml:"CONTAINER_MEMORY_LIMITS,omitempty"`
	ContainerCPULimits      string              `default:""  json:"CONTAINER_CPU_LIMITS,omitempty" yaml:"CONTAINER_CPU_LIMITS,omitempty"`
	ContainerMemoryRequest  string              `default:""  json:"CONTAINER_MEMORY_REQUEST,omitempty" yaml:"CONTAINER_MEMORY_REQUEST,omitempty"`
	ContainerCPURequest     string              `default:""  json:"CONTAINER_CPU_REQUEST,omitempty" yaml:"CONTAINER_CPU_REQUEST,omitempty"`
	NodeSelectorLabelKey    string              `default:""  json:"NODE_SELECTOR_LABEL_KEY,omitempty" yaml:"NODE_SELECTOR_LABEL_KEY,omitempty"`
	NodeSelectorLabelValue  string              `default:""  json:"NODE_SELECTOR_LABEL_VALUE,omitempty" yaml:"NODE_SELECTOR_LABEL_VALUE,omitempty"`
	ContainerRuntime        string              `default:""  json:"CONTAINER_RUNTIME,omitempty" yaml:"CONTAINER_RUNTIME,omitempty"`
	ServiceAccountName      string              `default:""  json:"SERVICE_ACCOUNT_NAME,omitempty" yaml:"SERVICE_ACCOUNT_NAME,omitempty"`
	ServiceAccountToken     string              `default:""  json:"SERVICE_ACCOUNT_TOKEN,omitempty" yaml:"SERVICE_ACCOUNT_TOKEN,omitempty"`
	QuayUsername            string              `default:""  json:"QUAY_USERNAME,omitempty" yaml:"QUAY_USERNAME,omitempty"`
	QuayPassword            string              `default:""  json:"QUAY_PASSWORD,omitempty" yaml:"QUAY_PASSWORD,omitempty"`
	WorkerImage             string              `default:""  json:"WORKER_IMAGE,omitempty" yaml:"WORKER_IMAGE,omitempty"`
	WorkerTag               string              `default:""  json:"WORKER_TAG,omitempty" yaml:"WORKER_TAG,omitempty"`
	BuilderVMContainerImage string              `default:""  json:"BUILDER_VM_CONTAINER_IMAGE,omitempty" yaml:"BUILDER_VM_CONTAINER_IMAGE,omitempty"`
	SetupTime               *shared.IntOrString `default:""  json:"SETUP_TIME,omitempty" yaml:"SETUP_TIME,omitempty"`
	MinimumRetryThreshold   *shared.IntOrString `default:""  json:"MINIMUM_RETRY_THRESHOLD,omitempty" yaml:"MINIMUM_RETRY_THRESHOLD,omitempty"`
	SSHAuthorizedKeys       []interface{}       `default:""  json:"SSH_AUTHORIZED_KEYS,omitempty" yaml:"SSH_AUTHORIZED_KEYS,omitempty"`
	// ec2 fields
	EC2Region           string              `default:""  json:"EC2_REGION,omitempty" yaml:"EC2_REGION,omitempty"`
	CoreOSAMI           string              `default:""  json:"COREOS_AMI,omitempty" yaml:"COREOS_AMI,omitempty"`
	AwsAccessKey        string              `default:""  json:"AWS_ACCESS_KEY,omitempty" yaml:"AWS_ACCESS_KEY,omitempty"`
	AwsSecretKey        string              `default:""  json:"AWS_SECRET_KEY,omitempty" yaml:"AWS_SECRET_KEY,omitempty"`
	EC2InstanceType     string              `default:""  json:"EC2_INSTANCE_TYPE,omitempty" yaml:"EC2_INSTANCE_TYPE,omitempty"`
	EC2VPCSubnetID      string              `default:""  json:"EC2_VPC_SUBNET_ID,omitempty" yaml:"EC2_VPC_SUBNET_ID,omitempty"`
	EC2SecurityGroupIDs []interface{}       `default:""  json:"EC2_SECURITY_GROUP_IDS,omitempty" yaml:"EC2_SECURITY_GROUP_IDS,omitempty"`
	EC2KeyName          string              `default:""  json:"EC2_KEY_NAME,omitempty" yaml:"EC2_KEY_NAME,omitempty"`
	BlockDeviceSize     *shared.IntOrString `default:""  json:"BLOCK_DEVICE_SIZE,omitempty" yaml:"BLOCK_DEVICE_SIZE,omitempty"`
}

// NewBuildManagerFieldGroup creates a new BitbucketBuildTriggerFieldGroup
func NewBuildManagerFieldGroup(fullConfig map[string]interface{}) (*BuildManagerFieldGroup, error) {
	newBuildManagerFieldGroup := &BuildManagerFieldGroup{}
	defaults.Set(newBuildManagerFieldGroup)

	bytes, err := yaml.Marshal(fullConfig)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, newBuildManagerFieldGroup)
	if err != nil {
		return nil, err
	}

	return newBuildManagerFieldGroup, nil
}

func (bm *BuildManagerDefinition) UnmarshalYAML(value *yaml.Node) error {

	// Ensure correct shape
	if len(value.Content) != 2 || value.Content[0].Tag != "!!str" || value.Content[1].Tag != "!!map" {
		return fmt.Errorf("Incorrect format for value BUILD_MANAGER")
	}

	bm.Name = value.Content[0].Value
	err := value.Content[1].Decode(&bm.Args)
	if err != nil {
		return err
	}

	return nil

}

func (bm *BuildManagerDefinition) MarshalYAML() (interface{}, error) {

	name := &yaml.Node{
		Kind:  yaml.ScalarNode,
		Value: bm.Name,
	}

	args := &yaml.Node{}
	err := args.Encode(bm.Args)
	if err != nil {
		return nil, err
	}

	node := &yaml.Node{
		Kind:    yaml.SequenceNode,
		Content: []*yaml.Node{name, args},
	}

	return node, nil

}
