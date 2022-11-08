


package azure

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:azure-native" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourcePackage(
		"azure-native",
		&pkg{version},
	)
	fmt.Println("\033[33m" + "warning:" + "\033[0m" + " module github.com/pulumi/pulumi-azure-native/sdk is deprecated; " +
		"update references to github.com/pulumi/pulumi-azure-native-sdk. " +
		"Go to https://github.com/pulumi/pulumi-azure-native/discussions/000 for more information.")
}
