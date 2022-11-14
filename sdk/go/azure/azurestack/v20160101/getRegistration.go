


package v20160101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRegistration(ctx *pulumi.Context, args *LookupRegistrationArgs, opts ...pulumi.InvokeOption) (*LookupRegistrationResult, error) {
	var rv LookupRegistrationResult
	err := ctx.Invoke("azure-native:azurestack/v20160101:getRegistration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegistrationArgs struct {
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type LookupRegistrationResult struct {
	BillingModel *string           `pulumi:"billingModel"`
	CloudId      *string           `pulumi:"cloudId"`
	Etag         *string           `pulumi:"etag"`
	Id           string            `pulumi:"id"`
	Location     string            `pulumi:"location"`
	Name         string            `pulumi:"name"`
	ObjectId     *string           `pulumi:"objectId"`
	Tags         map[string]string `pulumi:"tags"`
	Type         string            `pulumi:"type"`
}

func LookupRegistrationOutput(ctx *pulumi.Context, args LookupRegistrationOutputArgs, opts ...pulumi.InvokeOption) LookupRegistrationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRegistrationResult, error) {
			args := v.(LookupRegistrationArgs)
			r, err := LookupRegistration(ctx, &args, opts...)
			var s LookupRegistrationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRegistrationResultOutput)
}

type LookupRegistrationOutputArgs struct {
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup    pulumi.StringInput `pulumi:"resourceGroup"`
}

func (LookupRegistrationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationArgs)(nil)).Elem()
}


type LookupRegistrationResultOutput struct{ *pulumi.OutputState }

func (LookupRegistrationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRegistrationResult)(nil)).Elem()
}

func (o LookupRegistrationResultOutput) ToLookupRegistrationResultOutput() LookupRegistrationResultOutput {
	return o
}

func (o LookupRegistrationResultOutput) ToLookupRegistrationResultOutputWithContext(ctx context.Context) LookupRegistrationResultOutput {
	return o
}

func (o LookupRegistrationResultOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistrationResult) *string { return v.BillingModel }).(pulumi.StringPtrOutput)
}

func (o LookupRegistrationResultOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistrationResult) *string { return v.CloudId }).(pulumi.StringPtrOutput)
}

func (o LookupRegistrationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistrationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupRegistrationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRegistrationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupRegistrationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRegistrationResultOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRegistrationResult) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o LookupRegistrationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRegistrationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRegistrationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRegistrationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRegistrationResultOutput{})
}
