package shim

import (
	"github.com/datafy-io/terraform-provider-datafy/internal/provider"
	"github.com/datafy-io/terraform-provider-datafy/version"
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New(version.ProviderVersion)()
}
