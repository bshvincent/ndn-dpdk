package pingserver

import (
	"ndn-dpdk/core/nnduration"
	"ndn-dpdk/ndn"
)

// Server config.
type Config struct {
	Patterns []Pattern              // traffic patterns
	Nack     bool                   // whether to respond Nacks to unmatched Interests
	Delay    nnduration.Nanoseconds // minimum processing delay
}

// Server pattern definition.
type Pattern struct {
	Prefix  *ndn.Name // name prefix
	Replies []Reply   // reply settings
}

// Server reply definition.
type Reply struct {
	Weight int // weight of random choice, minimum is 1

	Suffix          *ndn.Name               // suffix to append to Interest name
	FreshnessPeriod nnduration.Milliseconds // FreshnessPeriod value
	PayloadLen      int                     // Content payload length

	Nack ndn.NackReason // if not NackReason_None, reply with Nack instead of Data

	Timeout bool // if true, drop the Interest instead of sending Data
}
