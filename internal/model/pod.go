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
	VolumeEmptyDir = "emptyDir"
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

func (p *Pod) ConvertToPod(k8sPod *corev1.Pod) {
	podLabels := make([]ListMapItem, 0)
	for k, v := range k8sPod.Labels {
		podLabels = append(podLabels, ListMapItem{
			Key:   k,
			Value: v,
		})
	}

	p.Base = Base{
		Name:          k8sPod.Name,
		Labels:        podLabels,
		Namespace:     k8sPod.Namespace,
		RestartPolicy: string(k8sPod.Spec.RestartPolicy),
	}

	hostAliases := make([]ListMapItem, 0)
	for _, alias := range k8sPod.Spec.HostAliases {
		hostAliases = append(hostAliases, ListMapItem{
			Key:   alias.IP,
			Value: strings.Join(alias.Hostnames, ","),
		})
	}

	var dnsConfig DnsConfig
	if k8sPod.Spec.DNSConfig != nil {
		dnsConfig.Nameservers = k8sPod.Spec.DNSConfig.Nameservers
	}

	p.NetWorking = NetWorking{
		HostNetwork: k8sPod.Spec.HostNetwork,
		HostName:    k8sPod.Spec.Hostname,
		DnsPolicy:   string(k8sPod.Spec.DNSPolicy),
		DnsConfig:   dnsConfig,
		HostAliases: hostAliases,
	}

	p.Volumes = p.ConvertVolumes(k8sPod)
	p.Containers = p.ConvertContainers(k8sPod, false)
	p.InitContainers = p.ConvertContainers(k8sPod, true)

}

func (p *Pod) ConvertVolumes(k8sPod *corev1.Pod) []Volume {
	volumeList := make([]Volume, 0)
	for _, v := range k8sPod.Spec.Volumes {
		if v.EmptyDir == nil {
			continue
		}
		volumeList = append(volumeList, Volume{
			Name: v.Name,
			Type: VolumeEmptyDir,
		})
	}
	return volumeList
}

func (p *Pod) ConvertContainerPorts(containerPortList []corev1.ContainerPort) []ContainerPort {
	portList := make([]ContainerPort, 0)
	for _, port := range containerPortList {
		portList = append(portList, ContainerPort{
			Name:          port.Name,
			ContainerPort: port.ContainerPort,
			HostPort:      port.HostPort,
		})
	}
	return portList
}

func (p *Pod) ConvertContainers(k8sPod *corev1.Pod, init bool) []Container {
	containerList := make([]Container, 0)
	var k8sContainers []corev1.Container
	if init {
		k8sContainers = k8sPod.Spec.InitContainers
	} else {
		k8sContainers = k8sPod.Spec.Containers
	}
	for _, container := range k8sContainers {
		containerList = append(containerList, Container{
			Name:            container.Name,
			Image:           container.Image,
			ImagePullPolicy: string(container.ImagePullPolicy),
			Tty:             container.TTY,
			Ports:           p.ConvertContainerPorts(container.Ports),
			WorkingDir:      container.WorkingDir,
			Command:         container.Command,
			Args:            container.Args,
			Envs:            p.ConvertContainerEnv(container.Env),
			Privileged:      p.ConvertContainerPrivileged(container.SecurityContext),
			Resources:       p.ConvertContainerResources(container.Resources),
			VolumeMounts:    p.ConvertContainerVolumeMounts(container.VolumeMounts),
			StartupProbe:    p.ConvertContainerProbe(container.StartupProbe),
			LivenessProbe:   p.ConvertContainerProbe(container.LivenessProbe),
			ReadinessProbe:  p.ConvertContainerProbe(container.ReadinessProbe),
		})
	}
	return containerList
}

func (p *Pod) ConvertContainerProbe(k8sProbe *corev1.Probe) ContainerProbe {
	probe := ContainerProbe{
		Enable:    false,
		Type:      "",
		HttpGet:   ProbeHttpGet{},
		Exec:      ProbeCommand{},
		TcpSocket: ProbeTcpSocket{},
		ProbeTime: ProbeTime{},
	}
	if k8sProbe == nil {
		return probe
	} else {
		probe.Enable = true
	}
	if k8sProbe.Exec != nil {
		probe.Type = EXECProbe
		probe.Exec.Command = k8sProbe.Exec.Command
	} else if k8sProbe.HTTPGet != nil {
		probe.Type = HTTPProbe
		httpGet := k8sProbe.HTTPGet
		headersReq := make([]ListMapItem, 0)
		for _, headerK8s := range httpGet.HTTPHeaders {
			headersReq = append(headersReq, ListMapItem{
				Key:   headerK8s.Name,
				Value: headerK8s.Value,
			})
		}
		probe.HttpGet = ProbeHttpGet{
			Host:        httpGet.Host,
			Port:        httpGet.Port.IntVal,
			Scheme:      string(httpGet.Scheme),
			Path:        httpGet.Path,
			HttpHeaders: headersReq,
		}
	} else if k8sProbe.TCPSocket != nil {
		probe.Type = TCPProbe
		probe.TcpSocket = ProbeTcpSocket{
			Host: k8sProbe.TCPSocket.Host,
			Port: k8sProbe.TCPSocket.Port.IntVal,
		}
	} else {
		probe.Type = HTTPProbe
	}

	probe.InitialDelaySeconds = k8sProbe.InitialDelaySeconds
	probe.PeriodSeconds = k8sProbe.PeriodSeconds
	probe.TimeoutSeconds = k8sProbe.TimeoutSeconds
	probe.SuccessThreshold = k8sProbe.SuccessThreshold
	probe.FailureThreshold = k8sProbe.FailureThreshold

	return probe
}

func (p *Pod) ConvertContainerVolumeMounts(volumeMountsK8s []corev1.VolumeMount) []VolumeMount {
	volumesReq := make([]VolumeMount, 0)
	for _, item := range volumeMountsK8s {

		volumesReq = append(volumesReq, VolumeMount{
			MountName: item.Name,
			MountPath: item.MountPath,
			ReadOnly:  item.ReadOnly,
		})

	}
	return volumesReq
}

func (p *Pod) ConvertContainerPrivileged(pr *corev1.SecurityContext) (privileged bool) {
	if pr != nil {
		privileged = *pr.Privileged
	}
	return
}

func (p *Pod) ConvertContainerEnv(containerEnvList []corev1.EnvVar) []EnvVar {
	envList := make([]EnvVar, 0)
	for _, envVar := range containerEnvList {
		envList = append(envList, EnvVar{
			Key:  envVar.Name,
			Name: envVar.Value,
		})
	}
	return envList
}

func (p *Pod) ConvertContainerResources(requirements corev1.ResourceRequirements) Resources {
	reqResources := Resources{
		Enable: false,
	}
	requests := requirements.Requests
	limits := requirements.Limits
	if requests != nil {
		reqResources.Enable = true
		reqResources.CpuRequest = int32(requests.Cpu().MilliValue())
		reqResources.MemRequest = int32(requests.Memory().Value() / (1024 * 1024))
	}
	if limits != nil {
		reqResources.Enable = true
		reqResources.CpuLimit = int32(limits.Cpu().MilliValue())
		reqResources.MemLimit = int32(limits.Memory().Value() / (1024 * 1024))
	}
	return reqResources
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
		if volume.Type != VolumeEmptyDir {
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
