


package v20221201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContact(ctx *pulumi.Context, args *LookupContactArgs, opts ...pulumi.InvokeOption) (*LookupContactResult, error) {
	var rv LookupContactResult
	err := ctx.Invoke("azure-native:voiceservices/v20221201preview:getContact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactArgs struct {
	CommunicationsGatewayName string `pulumi:"communicationsGatewayName"`
	ContactName               string `pulumi:"contactName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
}


type LookupContactResult struct {
	ContactName       string             `pulumi:"contactName"`
	Email             string             `pulumi:"email"`
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	PhoneNumber       string             `pulumi:"phoneNumber"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Role              string             `pulumi:"role"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
}

func LookupContactOutput(ctx *pulumi.Context, args LookupContactOutputArgs, opts ...pulumi.InvokeOption) LookupContactResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactResult, error) {
			args := v.(LookupContactArgs)
			r, err := LookupContact(ctx, &args, opts...)
			var s LookupContactResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactResultOutput)
}

type LookupContactOutputArgs struct {
	CommunicationsGatewayName pulumi.StringInput `pulumi:"communicationsGatewayName"`
	ContactName               pulumi.StringInput `pulumi:"contactName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupContactOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactArgs)(nil)).Elem()
}


type LookupContactResultOutput struct{ *pulumi.OutputState }

func (LookupContactResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactResult)(nil)).Elem()
}

func (o LookupContactResultOutput) ToLookupContactResultOutput() LookupContactResultOutput {
	return o
}

func (o LookupContactResultOutput) ToLookupContactResultOutputWithContext(ctx context.Context) LookupContactResultOutput {
	return o
}

func (o LookupContactResultOutput) ContactName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.ContactName }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) PhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.PhoneNumber }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupContactResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContactResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContactResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContactResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupContactResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactResultOutput{})
}
