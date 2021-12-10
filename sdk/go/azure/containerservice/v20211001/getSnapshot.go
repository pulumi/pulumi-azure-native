


package v20211001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:containerservice/v20211001:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ResourceName      string `pulumi:"resourceName"`
}


type LookupSnapshotResult struct {
	CreationData      *CreationDataResponse `pulumi:"creationData"`
	EnableFIPS        bool                  `pulumi:"enableFIPS"`
	Id                string                `pulumi:"id"`
	KubernetesVersion string                `pulumi:"kubernetesVersion"`
	Location          string                `pulumi:"location"`
	Name              string                `pulumi:"name"`
	NodeImageVersion  string                `pulumi:"nodeImageVersion"`
	OsSku             string                `pulumi:"osSku"`
	OsType            string                `pulumi:"osType"`
	SnapshotType      *string               `pulumi:"snapshotType"`
	SystemData        SystemDataResponse    `pulumi:"systemData"`
	Tags              map[string]string     `pulumi:"tags"`
	Type              string                `pulumi:"type"`
	VmSize            string                `pulumi:"vmSize"`
}
