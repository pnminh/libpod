package entities

import (
	"errors"
	"net"

	"github.com/containers/buildah/imagebuildah"
	"github.com/containers/libpod/libpod/events"
	"github.com/containers/libpod/pkg/specgen"
	"github.com/containers/storage/pkg/archive"
	"github.com/cri-o/ocicni/pkg/ocicni"
)

type Container struct {
	IdOrNamed
}

type Volume struct {
	Identifier
}

type Report struct {
	Id  []string
	Err map[string]error
}

type PodDeleteReport struct{ Report }

type VolumeDeleteOptions struct{}
type VolumeDeleteReport struct{ Report }

// NetOptions reflect the shared network options between
// pods and containers
type NetOptions struct {
	AddHosts           []string
	CNINetworks        []string
	UseImageResolvConf bool
	DNSOptions         []string
	DNSSearch          []string
	DNSServers         []net.IP
	Network            specgen.Namespace
	NoHosts            bool
	PublishPorts       []ocicni.PortMapping
	StaticIP           *net.IP
	StaticMAC          *net.HardwareAddr
}

// All CLI inspect commands and inspect sub-commands use the same options
type InspectOptions struct {
	Format string `json:",omitempty"`
	Latest bool   `json:",omitempty"`
	Size   bool   `json:",omitempty"`
	Type   string `json:",omitempty"`
}

// All API and CLI diff commands and diff sub-commands use the same options
type DiffOptions struct {
	Format  string `json:",omitempty"` // CLI only
	Latest  bool   `json:",omitempty"` // API and CLI, only supported by containers
	Archive bool   `json:",omitempty"` // CLI only
}

// DiffReport provides changes for object
type DiffReport struct {
	Changes []archive.Change
}

type EventsOptions struct {
	FromStart bool
	EventChan chan *events.Event
	Filter    []string
	Stream    bool
	Since     string
	Until     string
}

// ContainerCreateResponse is the response struct for creating a container
type ContainerCreateResponse struct {
	// ID of the container created
	ID string `json:"Id"`
	// Warnings during container creation
	Warnings []string `json:"Warnings"`
}

type ErrorModel struct {
	// API root cause formatted for automated parsing
	// example: API root cause
	Because string `json:"cause"`
	// human error message, formatted for a human to read
	// example: human error message
	Message string `json:"message"`
	// http response code
	ResponseCode int `json:"response"`
}

func (e ErrorModel) Error() string {
	return e.Message
}

func (e ErrorModel) Cause() error {
	return errors.New(e.Because)
}

func (e ErrorModel) Code() int {
	return e.ResponseCode
}

// BuildOptions describe the options for building container images.
type BuildOptions struct {
	imagebuildah.BuildOptions
}

// BuildReport is the image-build report.
type BuildReport struct {
	// ID of the image.
	ID string
}
