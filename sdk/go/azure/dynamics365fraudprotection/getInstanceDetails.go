


package dynamics365fraudprotection

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInstanceDetails(ctx *pulumi.Context, args *LookupInstanceDetailsArgs, opts ...pulumi.InvokeOption) (*LookupInstanceDetailsResult, error) {
	var rv LookupInstanceDetailsResult
	err := ctx.Invoke("azure-native:dynamics365fraudprotection:getInstanceDetails", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceDetailsArgs struct {
	InstanceName      string `pulumi:"instanceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupInstanceDetailsResult struct {
	Administration    *DFPInstanceAdministratorsResponse `pulumi:"administration"`
	Id                string                             `pulumi:"id"`
	Location          string                             `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse                 `pulumi:"systemData"`
	Tags              map[string]string                  `pulumi:"tags"`
	Type              string                             `pulumi:"type"`
}

func LookupInstanceDetailsOutput(ctx *pulumi.Context, args LookupInstanceDetailsOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceDetailsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceDetailsResult, error) {
			args := v.(LookupInstanceDetailsArgs)
			r, err := LookupInstanceDetails(ctx, &args, opts...)
			var s LookupInstanceDetailsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceDetailsResultOutput)
}

type LookupInstanceDetailsOutputArgs struct {
	InstanceName      pulumi.StringInput `pulumi:"instanceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupInstanceDetailsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceDetailsArgs)(nil)).Elem()
}


type LookupInstanceDetailsResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceDetailsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceDetailsResult)(nil)).Elem()
}

func (o LookupInstanceDetailsResultOutput) ToLookupInstanceDetailsResultOutput() LookupInstanceDetailsResultOutput {
	return o
}

func (o LookupInstanceDetailsResultOutput) ToLookupInstanceDetailsResultOutputWithContext(ctx context.Context) LookupInstanceDetailsResultOutput {
	return o
}

func (o LookupInstanceDetailsResultOutput) Administration() DFPInstanceAdministratorsResponsePtrOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) *DFPInstanceAdministratorsResponse { return v.Administration }).(DFPInstanceAdministratorsResponsePtrOutput)
}

func (o LookupInstanceDetailsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupInstanceDetailsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupInstanceDetailsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupInstanceDetailsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupInstanceDetailsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupInstanceDetailsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupInstanceDetailsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceDetailsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceDetailsResultOutput{})
}
