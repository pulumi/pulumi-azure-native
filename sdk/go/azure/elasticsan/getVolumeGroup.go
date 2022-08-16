


package elasticsan

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolumeGroup(ctx *pulumi.Context, args *LookupVolumeGroupArgs, opts ...pulumi.InvokeOption) (*LookupVolumeGroupResult, error) {
	var rv LookupVolumeGroupResult
	err := ctx.Invoke("azure-native:elasticsan:getVolumeGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeGroupArgs struct {
	ElasticSanName    string `pulumi:"elasticSanName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeGroupName   string `pulumi:"volumeGroupName"`
}


type LookupVolumeGroupResult struct {
	Encryption        *string                 `pulumi:"encryption"`
	Id                string                  `pulumi:"id"`
	Name              string                  `pulumi:"name"`
	NetworkAcls       *NetworkRuleSetResponse `pulumi:"networkAcls"`
	ProtocolType      *string                 `pulumi:"protocolType"`
	ProvisioningState string                  `pulumi:"provisioningState"`
	SystemData        SystemDataResponse      `pulumi:"systemData"`
	Tags              map[string]string       `pulumi:"tags"`
	Type              string                  `pulumi:"type"`
}

func LookupVolumeGroupOutput(ctx *pulumi.Context, args LookupVolumeGroupOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeGroupResult, error) {
			args := v.(LookupVolumeGroupArgs)
			r, err := LookupVolumeGroup(ctx, &args, opts...)
			var s LookupVolumeGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeGroupResultOutput)
}

type LookupVolumeGroupOutputArgs struct {
	ElasticSanName    pulumi.StringInput `pulumi:"elasticSanName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeGroupName   pulumi.StringInput `pulumi:"volumeGroupName"`
}

func (LookupVolumeGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupArgs)(nil)).Elem()
}


type LookupVolumeGroupResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeGroupResult)(nil)).Elem()
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutput() LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) ToLookupVolumeGroupResultOutputWithContext(ctx context.Context) LookupVolumeGroupResultOutput {
	return o
}

func (o LookupVolumeGroupResultOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *string { return v.Encryption }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *NetworkRuleSetResponse { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

func (o LookupVolumeGroupResultOutput) ProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) *string { return v.ProtocolType }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVolumeGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeGroupResultOutput{})
}
