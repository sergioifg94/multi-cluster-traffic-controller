package dnspolicy

type DNSTree struct {
	GeoNodes []*GeoEndpointNode
}

type GeoEndpointNode struct {
	CountryCode string
}

type ManagedCNAMENode struct {
}

type ClusterAddressNode struct {
}
