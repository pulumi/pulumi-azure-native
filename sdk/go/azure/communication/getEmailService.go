


package communication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEmailService(ctx *pulumi.Context, args *LookupEmailServiceArgs, opts ...pulumi.InvokeOption) (*LookupEmailServiceResult, error) {
	var rv LookupEmailServiceResult
	err := ctx.Invoke("azure-native:communication:getEmailService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEmailServiceArgs struct {
	EmailServiceName  string `pulumi:"emailServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEmailServiceResult struct {
	DataLocation      string             `pulumi:"dataLocation"`
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupEmailServiceOutput(ctx *pulumi.Context, args LookupEmailServiceOutputArgs, opts ...pulumi.InvokeOption) LookupEmailServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEmailServiceResult, error) {
			args := v.(LookupEmailServiceArgs)
			r, err := LookupEmailService(ctx, &args, opts...)
			var s LookupEmailServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEmailServiceResultOutput)
}

type LookupEmailServiceOutputArgs struct {
	EmailServiceName  pulumi.StringInput `pulumi:"emailServiceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupEmailServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailServiceArgs)(nil)).Elem()
}


type LookupEmailServiceResultOutput struct{ *pulumi.OutputState }

func (LookupEmailServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailServiceResult)(nil)).Elem()
}

func (o LookupEmailServiceResultOutput) ToLookupEmailServiceResultOutput() LookupEmailServiceResultOutput {
	return o
}

func (o LookupEmailServiceResultOutput) ToLookupEmailServiceResultOutputWithContext(ctx context.Context) LookupEmailServiceResultOutput {
	return o
}

func (o LookupEmailServiceResultOutput) DataLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) string { return v.DataLocation }).(pulumi.StringOutput)
}

func (o LookupEmailServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEmailServiceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupEmailServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupEmailServiceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupEmailServiceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupEmailServiceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupEmailServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEmailServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEmailServiceResultOutput{})
}
