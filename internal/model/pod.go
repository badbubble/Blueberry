package model

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	"strings"
)

type ProbeType int

const (
	LivenessProbe = iota
	StartupProbe
	ReadinessProbe
)

type ProbeHandler string

const (
	HTTPProbe = "http"
	TCPProbe  = "tcp"
	EXECProbe = "exec"
)

const (
	VOLUME_EMPTYDIR = "emptyDir"
)

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Base struct {
	Name          string        `json:"name"`
	Labels        []ListMapItem `json:"labels"`
	Namespace     string        `json:"namespace"`
	RestartPolicy string        `json:"restartPolicy"`
}

type Volume struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type DnsConfig struct {
	Nameservers []string `json:"nameservers"`
}
type NetWorking struct {
	HostNetwork bool          `json:"hostNetwork"`
	HostName    string        `json:"hostName"`
	DnsPolicy   string        `json:"dnsPolicy"`
	DnsConfig   DnsConfig     `json:"dnsConfig"`
	HostAliases []ListMapItem `json:"hostAliases"`
}

type Resources struct {
	Enable     bool  `json:"enable"`
	MemRequest int32 `json:"memRequest"`
	MemLimit   int32 `json:"memLimit"`
	CpuRequest int32 `json:"cpuRequest"`
	CpuLimit   int32 `json:"cpuLimit"`
}

type VolumeMount struct {
	MountName string `json:"mountName"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}

type ProbeHttpGet struct {
	Scheme      string        `json:"scheme"`
	Host        string        `json:"host"`
	Path        string        `json:"path"`
	Port        int32         `json:"port"`
	HttpHeaders []ListMapItem `json:"httpHeaders"`
}
type ProbeCommand struct {
	Command []string `json:"command"`
}

type ProbeTcpSocket struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

type ProbeTime struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds"`
	PeriodSeconds       int32 `json:"periodSeconds"`
	TimeoutSeconds      int32 `json:"timeoutSeconds"`
	SuccessThreshold    int32 `json:"successThreshold"`
	FailureThreshold    int32 `json:"failureThreshold"`
}

type ContainerProbe struct {
	Enable    bool           `json:"enable"`
	Type      string         `json:"type"`
	HttpGet   ProbeHttpGet   `json:"httpGet"`
	Exec      ProbeCommand   `json:"exec"`
	TcpSocket ProbeTcpSocket `json:"tcpSocket"`
	ProbeTime
}

type ContainerPort struct {
	Name          string `json:"name"`
	ContainerPort int32  `json:"containerPort"`
	HostPort      int32  `json:"hostPort"`
}

type EnvVar struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	RefName string `json:"refName"`
	Value   string `json:"value"`
	Type    string `json:"type"`
}
type EnvVarFromResource struct {
	Name    string `json:"name"`
	RefType string `json:"refType"`
	Prefix  string `json:"prefix"`
}

type Container struct {
	Name            string               `json:"name"`
	Image           string               `json:"image"`
	ImagePullPolicy string               `json:"imagePullPolicy"`
	Tty             bool                 `json:"tty"`
	Ports           []ContainerPort      `json:"ports"`
	WorkingDir      string               `json:"workingDir"`
	Command         []string             `json:"command"`
	Args            []string             `json:"args"`
	Envs            []EnvVar             `json:"envs"`
	EnvsFrom        []EnvVarFromResource `json:"envsFrom"`
	Privileged      bool                 `json:"privileged"`
	Resources       Resources            `json:"resources"`
	VolumeMounts    []VolumeMount        `json:"volumeMounts"`
	StartupProbe    ContainerProbe       `json:"startupProbe"`
	LivenessProbe   ContainerProbe       `json:"livenessProbe"`
	ReadinessProbe  ContainerProbe       `json:"readinessProbe"`
}

type Pod struct {
	Base           Base        `json:"base"`
	Volumes        []Volume    `json:"volumes"`
	NetWorking     NetWorking  `json:"netWorking"`
	InitContainers []Container `json:"initContainers"`
	Containers     []Container `json:"containers"`
}

func (p *Pod) ConvertToK8s() *corev1.Pod {
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      p.Base.Name,
			Namespace: p.Base.Namespace,
			Labels:    p.GetLabels(),
		},
		Spec: corev1.PodSpec{
			Volumes:        p.GetVolumes(),
			InitContainers: p.GetContainers(true),
			Containers:     p.GetContainers(false),
			DNSConfig: &corev1.PodDNSConfig{
				Nameservers: p.NetWorking.DnsConfig.Nameservers,
			},
			DNSPolicy:     corev1.DNSPolicy(p.NetWorking.DnsPolicy),
			HostAliases:   p.GetHostAliases(),
			Hostname:      p.NetWorking.HostName,
			RestartPolicy: corev1.RestartPolicy(p.Base.RestartPolicy),
		},
		Status: corev1.PodStatus{},
	}
}

func (p *Pod) GetHostAliases() []corev1.HostAlias {
	hostAliases := make([]corev1.HostAlias, 0)
	for _, alias := range p.NetWorking.HostAliases {
		hostAliases = append(hostAliases, corev1.HostAlias{
			IP:        alias.Key,
			Hostnames: strings.Split(alias.Value, ","),
		})
	}
	return hostAliases
}

func (p *Pod) GetVolumes() []corev1.Volume {
	volumes := make([]corev1.Volume, 0)
	for _, volume := range p.Volumes {
		if volume.Type != VOLUME_EMPTYDIR {
			continue
		}
		source := corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		}
		volumes = append(volumes, corev1.Volume{
			Name:         volume.Name,
			VolumeSource: source,
		})
	}
	return volumes
}

func (p *Pod) GetLabels() map[string]string {
	labels := make(map[string]string)
	for _, l := range p.Base.Labels {
		labels[l.Key] = l.Value
	}
	return labels
}

func (p *Pod) GetContainers(init bool) []corev1.Container {
	var containers []Container
	if init {
		containers = p.InitContainers
	} else {
		containers = p.Containers
	}
	k8sContainers := make([]corev1.Container, 0)
	for _, container := range containers {
		k8sContainers = append(k8sContainers, corev1.Container{
			Name:            container.Name,
			Image:           container.Image,
			Command:         container.Command,
			Args:            container.Args,
			WorkingDir:      container.WorkingDir,
			Ports:           container.GetPorts(),
			Env:             container.GetEnv(),
			Resources:       container.GetResources(),
			ImagePullPolicy: corev1.PullPolicy(container.ImagePullPolicy),
			SecurityContext: &corev1.SecurityContext{Privileged: &container.Privileged},
			TTY:             container.Tty,
			VolumeMounts:    container.GetVolumeMounts(),
			StartupProbe:    container.GetProbe(StartupProbe),
			LivenessProbe:   container.GetProbe(LivenessProbe),
			ReadinessProbe:  container.GetProbe(ReadinessProbe),
		})
	}
	return k8sContainers
}

func (c *Container) GetEnv() []corev1.EnvVar {
	envs := make([]corev1.EnvVar, 0)
	for _, e := range c.Envs {
		envs = append(envs, corev1.EnvVar{
			Name:      e.Key,
			Value:     e.Value,
			ValueFrom: nil,
		})
	}
	return envs
}

func (c *Container) GetVolumeMounts() []corev1.VolumeMount {
	volumeMounts := make([]corev1.VolumeMount, 0)
	for _, mount := range c.VolumeMounts {
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      mount.MountName,
			ReadOnly:  mount.ReadOnly,
			MountPath: mount.MountPath,
		})
	}
	return volumeMounts
}

func (c *Container) GetProbe(probeType ProbeType) *corev1.Probe {
	var containerProbe ContainerProbe

	switch probeType {
	case LivenessProbe:
		if c.LivenessProbe.Enable {
			containerProbe = c.LivenessProbe
		} else {
			return nil
		}
	case ReadinessProbe:
		if c.ReadinessProbe.Enable {
			containerProbe = c.ReadinessProbe
		} else {
			return nil
		}
	case StartupProbe:
		if c.StartupProbe.Enable {
			containerProbe = c.StartupProbe
		} else {
			return nil
		}
	}

	result := &corev1.Probe{}

	result.InitialDelaySeconds = containerProbe.InitialDelaySeconds
	result.TimeoutSeconds = containerProbe.TimeoutSeconds
	result.PeriodSeconds = containerProbe.PeriodSeconds
	result.SuccessThreshold = containerProbe.SuccessThreshold
	result.FailureThreshold = containerProbe.FailureThreshold

	switch containerProbe.Type {
	case HTTPProbe:
		headers := make([]corev1.HTTPHeader, 0)
		for _, header := range containerProbe.HttpGet.HttpHeaders {
			headers = append(headers, corev1.HTTPHeader{
				Name:  header.Key,
				Value: header.Value,
			})
		}
		result.ProbeHandler.HTTPGet = &corev1.HTTPGetAction{
			Path:        containerProbe.HttpGet.Path,
			Port:        intstr.FromInt32(containerProbe.HttpGet.Port),
			Host:        containerProbe.HttpGet.Host,
			Scheme:      corev1.URIScheme(containerProbe.HttpGet.Scheme),
			HTTPHeaders: headers,
		}
		break
	case TCPProbe:
		result.ProbeHandler.TCPSocket = &corev1.TCPSocketAction{
			Port: intstr.FromInt32(containerProbe.TcpSocket.Port),
			Host: containerProbe.TcpSocket.Host,
		}
		break
	case EXECProbe:
		result.ProbeHandler.Exec = &corev1.ExecAction{Command: containerProbe.Exec.Command}
	}
	return result
}

func (c *Container) GetResources() corev1.ResourceRequirements {
	resources := corev1.ResourceRequirements{
		Limits:   nil,
		Requests: nil,
		Claims:   nil,
	}
	if c.Resources.Enable {
		resources.Requests = corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(c.Resources.CpuRequest)) + "m"),
			corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(c.Resources.MemRequest)) + "Mi"),
		}
		resources.Limits = corev1.ResourceList{
			corev1.ResourceCPU:    resource.MustParse(strconv.Itoa(int(c.Resources.CpuLimit)) + "m"),
			corev1.ResourceMemory: resource.MustParse(strconv.Itoa(int(c.Resources.MemLimit)) + "Mi"),
		}
	}
	return resources
}

func (c *Container) GetPorts() []corev1.ContainerPort {
	containerPorts := make([]corev1.ContainerPort, 0)
	for _, port := range c.Ports {
		containerPorts = append(containerPorts, corev1.ContainerPort{
			Name:          port.Name,
			HostPort:      port.HostPort,
			ContainerPort: port.ContainerPort,
		})
	}
	return containerPorts
}
