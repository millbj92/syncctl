package connections

type ConntrackStatRequest struct {
	PerCPU bool `json:"per_cpu"`
}

type ConnectionsRequest struct {
	Kind string `json:"kind"`
}
