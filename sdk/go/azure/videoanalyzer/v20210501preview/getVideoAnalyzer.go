


package v20210501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVideoAnalyzer(ctx *pulumi.Context, args *LookupVideoAnalyzerArgs, opts ...pulumi.InvokeOption) (*LookupVideoAnalyzerResult, error) {
	var rv LookupVideoAnalyzerResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getVideoAnalyzer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVideoAnalyzerArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVideoAnalyzerResult struct {
	Encryption      AccountEncryptionResponse      `pulumi:"encryption"`
	Endpoints       []EndpointResponse             `pulumi:"endpoints"`
	Id              string                         `pulumi:"id"`
	Identity        *VideoAnalyzerIdentityResponse `pulumi:"identity"`
	Location        string                         `pulumi:"location"`
	Name            string                         `pulumi:"name"`
	StorageAccounts []StorageAccountResponse       `pulumi:"storageAccounts"`
	SystemData      SystemDataResponse             `pulumi:"systemData"`
	Tags            map[string]string              `pulumi:"tags"`
	Type            string                         `pulumi:"type"`
}

func LookupVideoAnalyzerOutput(ctx *pulumi.Context, args LookupVideoAnalyzerOutputArgs, opts ...pulumi.InvokeOption) LookupVideoAnalyzerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVideoAnalyzerResult, error) {
			args := v.(LookupVideoAnalyzerArgs)
			r, err := LookupVideoAnalyzer(ctx, &args, opts...)
			var s LookupVideoAnalyzerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVideoAnalyzerResultOutput)
}

type LookupVideoAnalyzerOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVideoAnalyzerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoAnalyzerArgs)(nil)).Elem()
}


type LookupVideoAnalyzerResultOutput struct{ *pulumi.OutputState }

func (LookupVideoAnalyzerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVideoAnalyzerResult)(nil)).Elem()
}

func (o LookupVideoAnalyzerResultOutput) ToLookupVideoAnalyzerResultOutput() LookupVideoAnalyzerResultOutput {
	return o
}

func (o LookupVideoAnalyzerResultOutput) ToLookupVideoAnalyzerResultOutputWithContext(ctx context.Context) LookupVideoAnalyzerResultOutput {
	return o
}

func (o LookupVideoAnalyzerResultOutput) Encryption() AccountEncryptionResponseOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) AccountEncryptionResponse { return v.Encryption }).(AccountEncryptionResponseOutput)
}

func (o LookupVideoAnalyzerResultOutput) Endpoints() EndpointResponseArrayOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) []EndpointResponse { return v.Endpoints }).(EndpointResponseArrayOutput)
}

func (o LookupVideoAnalyzerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVideoAnalyzerResultOutput) Identity() VideoAnalyzerIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) *VideoAnalyzerIdentityResponse { return v.Identity }).(VideoAnalyzerIdentityResponsePtrOutput)
}

func (o LookupVideoAnalyzerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVideoAnalyzerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVideoAnalyzerResultOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) []StorageAccountResponse { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o LookupVideoAnalyzerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVideoAnalyzerResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVideoAnalyzerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVideoAnalyzerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVideoAnalyzerResultOutput{})
}
