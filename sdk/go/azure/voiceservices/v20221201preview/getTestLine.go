


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTestLine(ctx *pulumi.Context, args *LookupTestLineArgs, opts ...pulumi.InvokeOption) (*LookupTestLineResult, error) {
	var rv LookupTestLineResult
	err := ctx.Invoke("azure-native:voiceservices/v20221201preview:getTestLine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestLineArgs struct {
	CommunicationsGatewayName string `pulumi:"communicationsGatewayName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	TestLineName              string `pulumi:"testLineName"`
}


type LookupTestLineResult struct {
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	PhoneNumber       string             `pulumi:"phoneNumber"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Purpose           string             `pulumi:"purpose"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupTestLineOutput(ctx *pulumi.Context, args LookupTestLineOutputArgs, opts ...pulumi.InvokeOption) LookupTestLineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTestLineResult, error) {
			args := v.(LookupTestLineArgs)
			r, err := LookupTestLine(ctx, &args, opts...)
			var s LookupTestLineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTestLineResultOutput)
}

type LookupTestLineOutputArgs struct {
	CommunicationsGatewayName pulumi.StringInput `pulumi:"communicationsGatewayName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	TestLineName              pulumi.StringInput `pulumi:"testLineName"`
}

func (LookupTestLineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestLineArgs)(nil)).Elem()
}


type LookupTestLineResultOutput struct{ *pulumi.OutputState }

func (LookupTestLineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestLineResult)(nil)).Elem()
}

func (o LookupTestLineResultOutput) ToLookupTestLineResultOutput() LookupTestLineResultOutput {
	return o
}

func (o LookupTestLineResultOutput) ToLookupTestLineResultOutputWithContext(ctx context.Context) LookupTestLineResultOutput {
	return o
}

func (o LookupTestLineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTestLineResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupTestLineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTestLineResultOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

func (o LookupTestLineResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupTestLineResultOutput) Purpose() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Purpose }).(pulumi.StringOutput)
}

func (o LookupTestLineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTestLineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTestLineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTestLineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupTestLineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestLineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestLineResultOutput{})
}
