package config

import (
	"time"

	p2pforge "github.com/ipshipyard/p2p-forge/client"
)

// AutoTLS includes optional configuration of p2p-forge client of service
// for obtaining a domain and TLS certificate to improve connectivity for web
// browser clients. More: https://github.com/ipshipyard/p2p-forge#readme
type AutoTLS struct {
	// Enables the p2p-forge feature and all related features.
	Enabled Flag `json:",omitempty"`

	// Optional, controls if Kubo should add /tls/sni/.../ws listener to every /tcp port if no explicit /ws is defined in Addresses.Swarm
	AutoWSS Flag `json:",omitempty"`

	// Optional override of the parent domain that will be used
	DomainSuffix *OptionalString `json:",omitempty"`

	// Optional override of HTTP API that acts as ACME DNS-01 Challenge broker
	RegistrationEndpoint *OptionalString `json:",omitempty"`

	// Optional Authorization token, used with private/test instances of p2p-forge
	RegistrationToken *OptionalString `json:",omitempty"`

	// Optional registration delay used when AutoTLS.Enabled is not explicitly set to true in config
	RegistrationDelay *OptionalDuration `json:",omitempty"`

	// Optional override of CA ACME API used by p2p-forge system
	CAEndpoint *OptionalString `json:",omitempty"`

	// Optional, controls if features like AutoWSS should generate shorter /dnsX instead of /ipX/../sni/..
	ShortAddrs Flag `json:",omitempty"`
}

const (
	DefaultAutoTLSEnabled           = true // with DefaultAutoTLSRegistrationDelay, unless explicitly enabled  in config
	DefaultDomainSuffix             = p2pforge.DefaultForgeDomain
	DefaultRegistrationEndpoint     = p2pforge.DefaultForgeEndpoint
	DefaultCAEndpoint               = p2pforge.DefaultCAEndpoint
	DefaultAutoWSS                  = true // requires AutoTLS.Enabled
	DefaultAutoTLSShortAddrs        = true // requires AutoTLS.Enabled
	DefaultAutoTLSRegistrationDelay = 1 * time.Hour
)
