


package connectedvmwarevsphere

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHybridIdentityMetadatum(ctx *pulumi.Context, args *LookupHybridIdentityMetadatumArgs, opts ...pulumi.InvokeOption) (*LookupHybridIdentityMetadatumResult, error) {
	var rv LookupHybridIdentityMetadatumResult
	err := ctx.Invoke("azure-native:connectedvmwarevsphere:getHybridIdentityMetadatum", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHybridIdentityMetadatumArgs struct {
	MetadataName       string `pulumi:"metadataName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	VirtualMachineName string `pulumi:"virtualMachineName"`
}


type LookupHybridIdentityMetadatumResult struct {
	Id                string             `pulumi:"id"`
	Identity          IdentityResponse   `pulumi:"identity"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	PublicKey         *string            `pulumi:"publicKey"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
	VmId              *string            `pulumi:"vmId"`
}
