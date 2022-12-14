


package v20221201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLoadTest(ctx *pulumi.Context, args *LookupLoadTestArgs, opts ...pulumi.InvokeOption) (*LookupLoadTestResult, error) {
	var rv LookupLoadTestResult
	err := ctx.Invoke("azure-native:loadtestservice/v20221201:getLoadTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLoadTestArgs struct {
	LoadTestName      string `pulumi:"loadTestName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLoadTestResult struct {
	DataPlaneURI      string                          `pulumi:"dataPlaneURI"`
	Description       *string                         `pulumi:"description"`
	Encryption        *EncryptionPropertiesResponse   `pulumi:"encryption"`
	Id                string                          `pulumi:"id"`
	Identity          *ManagedServiceIdentityResponse `pulumi:"identity"`
	Location          string                          `pulumi:"location"`
	Name              string                          `pulumi:"name"`
	ProvisioningState string                          `pulumi:"provisioningState"`
	SystemData        SystemDataResponse              `pulumi:"systemData"`
	Tags              map[string]string               `pulumi:"tags"`
	Type              string                          `pulumi:"type"`
}

func LookupLoadTestOutput(ctx *pulumi.Context, args LookupLoadTestOutputArgs, opts ...pulumi.InvokeOption) LookupLoadTestResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLoadTestResult, error) {
			args := v.(LookupLoadTestArgs)
			r, err := LookupLoadTest(ctx, &args, opts...)
			var s LookupLoadTestResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLoadTestResultOutput)
}

type LookupLoadTestOutputArgs struct {
	LoadTestName      pulumi.StringInput `pulumi:"loadTestName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLoadTestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadTestArgs)(nil)).Elem()
}


type LookupLoadTestResultOutput struct{ *pulumi.OutputState }

func (LookupLoadTestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLoadTestResult)(nil)).Elem()
}

func (o LookupLoadTestResultOutput) ToLookupLoadTestResultOutput() LookupLoadTestResultOutput {
	return o
}

func (o LookupLoadTestResultOutput) ToLookupLoadTestResultOutputWithContext(ctx context.Context) LookupLoadTestResultOutput {
	return o
}

func (o LookupLoadTestResultOutput) DataPlaneURI() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadTestResult) string { return v.DataPlaneURI }).(pulumi.StringOutput)
}

func (o LookupLoadTestResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLoadTestResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLoadTestResultOutput) Encryption() EncryptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadTestResult) *EncryptionPropertiesResponse { return v.Encryption }).(EncryptionPropertiesResponsePtrOutput)
}

func (o LookupLoadTestResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadTestResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLoadTestResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupLoadTestResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupLoadTestResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadTestResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupLoadTestResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadTestResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLoadTestResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadTestResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupLoadTestResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLoadTestResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLoadTestResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupLoadTestResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupLoadTestResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLoadTestResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLoadTestResultOutput{})
}
