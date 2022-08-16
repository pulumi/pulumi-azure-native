


package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFactory(ctx *pulumi.Context, args *LookupFactoryArgs, opts ...pulumi.InvokeOption) (*LookupFactoryResult, error) {
	var rv LookupFactoryResult
	err := ctx.Invoke("azure-native:datafactory:getFactory", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFactoryArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFactoryResult struct {
	CreateTime           string                                          `pulumi:"createTime"`
	ETag                 string                                          `pulumi:"eTag"`
	Encryption           *EncryptionConfigurationResponse                `pulumi:"encryption"`
	GlobalParameters     map[string]GlobalParameterSpecificationResponse `pulumi:"globalParameters"`
	Id                   string                                          `pulumi:"id"`
	Identity             *FactoryIdentityResponse                        `pulumi:"identity"`
	Location             *string                                         `pulumi:"location"`
	Name                 string                                          `pulumi:"name"`
	ProvisioningState    string                                          `pulumi:"provisioningState"`
	PublicNetworkAccess  *string                                         `pulumi:"publicNetworkAccess"`
	PurviewConfiguration *PurviewConfigurationResponse                   `pulumi:"purviewConfiguration"`
	RepoConfiguration    interface{}                                     `pulumi:"repoConfiguration"`
	Tags                 map[string]string                               `pulumi:"tags"`
	Type                 string                                          `pulumi:"type"`
	Version              string                                          `pulumi:"version"`
}

func LookupFactoryOutput(ctx *pulumi.Context, args LookupFactoryOutputArgs, opts ...pulumi.InvokeOption) LookupFactoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFactoryResult, error) {
			args := v.(LookupFactoryArgs)
			r, err := LookupFactory(ctx, &args, opts...)
			var s LookupFactoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFactoryResultOutput)
}

type LookupFactoryOutputArgs struct {
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFactoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFactoryArgs)(nil)).Elem()
}


type LookupFactoryResultOutput struct{ *pulumi.OutputState }

func (LookupFactoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFactoryResult)(nil)).Elem()
}

func (o LookupFactoryResultOutput) ToLookupFactoryResultOutput() LookupFactoryResultOutput {
	return o
}

func (o LookupFactoryResultOutput) ToLookupFactoryResultOutputWithContext(ctx context.Context) LookupFactoryResultOutput {
	return o
}

func (o LookupFactoryResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupFactoryResultOutput) ETag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.ETag }).(pulumi.StringOutput)
}

func (o LookupFactoryResultOutput) Encryption() EncryptionConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *EncryptionConfigurationResponse { return v.Encryption }).(EncryptionConfigurationResponsePtrOutput)
}

func (o LookupFactoryResultOutput) GlobalParameters() GlobalParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupFactoryResult) map[string]GlobalParameterSpecificationResponse { return v.GlobalParameters }).(GlobalParameterSpecificationResponseMapOutput)
}

func (o LookupFactoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFactoryResultOutput) Identity() FactoryIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *FactoryIdentityResponse { return v.Identity }).(FactoryIdentityResponsePtrOutput)
}

func (o LookupFactoryResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupFactoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFactoryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFactoryResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o LookupFactoryResultOutput) PurviewConfiguration() PurviewConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupFactoryResult) *PurviewConfigurationResponse { return v.PurviewConfiguration }).(PurviewConfigurationResponsePtrOutput)
}

func (o LookupFactoryResultOutput) RepoConfiguration() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupFactoryResult) interface{} { return v.RepoConfiguration }).(pulumi.AnyOutput)
}

func (o LookupFactoryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFactoryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupFactoryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupFactoryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFactoryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFactoryResultOutput{})
}
