package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/leryn1122/terraform-provider-vmware-fusion/pkg/terraform"
)

const (
	TerraformProviderName = "registry.terraform.io/leryn1122/vmware-fusion"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "Enable debug mode if run with debugger like `delve`")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: TerraformProviderName,
		Debug:   debugMode,
	}

	err := providerserver.Serve(context.Background(), terraform.NewProvider, opts)

	if err != nil {
		fmt.Println(err)
	}
}
