package compoundapi

import (
	"github.com/mycryptohq/DeFiReserveMapper/pkg"
	"github.com/postables/go-compound/client"
)


func MakeCompoundApiClient () (*client.Client) {
	client := client.NewClient(root.CompoundEndpoint)
	return client
}
