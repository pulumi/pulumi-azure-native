


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetP2sVpnGatewayP2sVpnConnectionHealthDetailed(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult, error) {
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult
	err := ctx.Invoke("azure-native:network/v20200801:getP2sVpnGatewayP2sVpnConnectionHealthDetailed", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs struct {
	GatewayName        string   `pulumi:"gatewayName"`
	OutputBlobSasUrl   *string  `pulumi:"outputBlobSasUrl"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	VpnUserNamesFilter []string `pulumi:"vpnUserNamesFilter"`
}


type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult struct {
	SasUrl *string `pulumi:"sasUrl"`
}

func GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutput(ctx *pulumi.Context, args GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutputArgs, opts ...pulumi.InvokeOption) GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult, error) {
			args := v.(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs)
			r, err := GetP2sVpnGatewayP2sVpnConnectionHealthDetailed(ctx, &args, opts...)
			var s GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput)
}

type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutputArgs struct {
	GatewayName        pulumi.StringInput      `pulumi:"gatewayName"`
	OutputBlobSasUrl   pulumi.StringPtrInput   `pulumi:"outputBlobSasUrl"`
	ResourceGroupName  pulumi.StringInput      `pulumi:"resourceGroupName"`
	VpnUserNamesFilter pulumi.StringArrayInput `pulumi:"vpnUserNamesFilter"`
}

func (GetP2sVpnGatewayP2sVpnConnectionHealthDetailedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthDetailedArgs)(nil)).Elem()
}


type GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput struct{ *pulumi.OutputState }

func (GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult)(nil)).Elem()
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput() GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) ToGetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutputWithContext(ctx context.Context) GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput {
	return o
}

func (o GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResult) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetP2sVpnGatewayP2sVpnConnectionHealthDetailedResultOutput{})
}
