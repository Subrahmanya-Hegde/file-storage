package p2p

// Any node in the peer-2-peer connection
type Peer interface{}

// Anything that handles rhe connection between the Peers. (TCP, UDP etc.)
type Transport interface{}
