


package servicelinker

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinker(ctx *pulumi.Context, args *LookupLinkerArgs, opts ...pulumi.InvokeOption) (*LookupLinkerResult, error) {
	var rv LookupLinkerResult
	err := ctx.Invoke("azure-native:servicelinker:getLinker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkerArgs struct {
	LinkerName  string `pulumi:"linkerName"`
	ResourceUri string `pulumi:"resourceUri"`
}


type LookupLinkerResult struct {
	AuthInfo          interface{}        `pulumi:"authInfo"`
	ClientType        *string            `pulumi:"clientType"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	TargetId          *string            `pulumi:"targetId"`
	Type              string             `pulumi:"type"`
}
