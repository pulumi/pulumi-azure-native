


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AggregateFunctionProperties struct {
	Binding interface{}         `pulumi:"binding"`
	Inputs  []FunctionInputType `pulumi:"inputs"`
	Output  *FunctionOutputType `pulumi:"output"`
	Type    string              `pulumi:"type"`
}





type AggregateFunctionPropertiesInput interface {
	pulumi.Input

	ToAggregateFunctionPropertiesOutput() AggregateFunctionPropertiesOutput
	ToAggregateFunctionPropertiesOutputWithContext(context.Context) AggregateFunctionPropertiesOutput
}

type AggregateFunctionPropertiesArgs struct {
	Binding pulumi.Input                `pulumi:"binding"`
	Inputs  FunctionInputTypeArrayInput `pulumi:"inputs"`
	Output  FunctionOutputTypePtrInput  `pulumi:"output"`
	Type    pulumi.StringInput          `pulumi:"type"`
}

func (AggregateFunctionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregateFunctionProperties)(nil)).Elem()
}

func (i AggregateFunctionPropertiesArgs) ToAggregateFunctionPropertiesOutput() AggregateFunctionPropertiesOutput {
	return i.ToAggregateFunctionPropertiesOutputWithContext(context.Background())
}

func (i AggregateFunctionPropertiesArgs) ToAggregateFunctionPropertiesOutputWithContext(ctx context.Context) AggregateFunctionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateFunctionPropertiesOutput)
}

type AggregateFunctionPropertiesOutput struct{ *pulumi.OutputState }

func (AggregateFunctionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregateFunctionProperties)(nil)).Elem()
}

func (o AggregateFunctionPropertiesOutput) ToAggregateFunctionPropertiesOutput() AggregateFunctionPropertiesOutput {
	return o
}

func (o AggregateFunctionPropertiesOutput) ToAggregateFunctionPropertiesOutputWithContext(ctx context.Context) AggregateFunctionPropertiesOutput {
	return o
}

func (o AggregateFunctionPropertiesOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v AggregateFunctionProperties) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

func (o AggregateFunctionPropertiesOutput) Inputs() FunctionInputTypeArrayOutput {
	return o.ApplyT(func(v AggregateFunctionProperties) []FunctionInputType { return v.Inputs }).(FunctionInputTypeArrayOutput)
}

func (o AggregateFunctionPropertiesOutput) Output() FunctionOutputTypePtrOutput {
	return o.ApplyT(func(v AggregateFunctionProperties) *FunctionOutputType { return v.Output }).(FunctionOutputTypePtrOutput)
}

func (o AggregateFunctionPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AggregateFunctionProperties) string { return v.Type }).(pulumi.StringOutput)
}

type AggregateFunctionPropertiesResponse struct {
	Binding interface{}             `pulumi:"binding"`
	Etag    string                  `pulumi:"etag"`
	Inputs  []FunctionInputResponse `pulumi:"inputs"`
	Output  *FunctionOutputResponse `pulumi:"output"`
	Type    string                  `pulumi:"type"`
}





type AggregateFunctionPropertiesResponseInput interface {
	pulumi.Input

	ToAggregateFunctionPropertiesResponseOutput() AggregateFunctionPropertiesResponseOutput
	ToAggregateFunctionPropertiesResponseOutputWithContext(context.Context) AggregateFunctionPropertiesResponseOutput
}

type AggregateFunctionPropertiesResponseArgs struct {
	Binding pulumi.Input                    `pulumi:"binding"`
	Etag    pulumi.StringInput              `pulumi:"etag"`
	Inputs  FunctionInputResponseArrayInput `pulumi:"inputs"`
	Output  FunctionOutputResponsePtrInput  `pulumi:"output"`
	Type    pulumi.StringInput              `pulumi:"type"`
}

func (AggregateFunctionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregateFunctionPropertiesResponse)(nil)).Elem()
}

func (i AggregateFunctionPropertiesResponseArgs) ToAggregateFunctionPropertiesResponseOutput() AggregateFunctionPropertiesResponseOutput {
	return i.ToAggregateFunctionPropertiesResponseOutputWithContext(context.Background())
}

func (i AggregateFunctionPropertiesResponseArgs) ToAggregateFunctionPropertiesResponseOutputWithContext(ctx context.Context) AggregateFunctionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AggregateFunctionPropertiesResponseOutput)
}

type AggregateFunctionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AggregateFunctionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AggregateFunctionPropertiesResponse)(nil)).Elem()
}

func (o AggregateFunctionPropertiesResponseOutput) ToAggregateFunctionPropertiesResponseOutput() AggregateFunctionPropertiesResponseOutput {
	return o
}

func (o AggregateFunctionPropertiesResponseOutput) ToAggregateFunctionPropertiesResponseOutputWithContext(ctx context.Context) AggregateFunctionPropertiesResponseOutput {
	return o
}

func (o AggregateFunctionPropertiesResponseOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v AggregateFunctionPropertiesResponse) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

func (o AggregateFunctionPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v AggregateFunctionPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o AggregateFunctionPropertiesResponseOutput) Inputs() FunctionInputResponseArrayOutput {
	return o.ApplyT(func(v AggregateFunctionPropertiesResponse) []FunctionInputResponse { return v.Inputs }).(FunctionInputResponseArrayOutput)
}

func (o AggregateFunctionPropertiesResponseOutput) Output() FunctionOutputResponsePtrOutput {
	return o.ApplyT(func(v AggregateFunctionPropertiesResponse) *FunctionOutputResponse { return v.Output }).(FunctionOutputResponsePtrOutput)
}

func (o AggregateFunctionPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AggregateFunctionPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AvroSerialization struct {
	Type string `pulumi:"type"`
}





type AvroSerializationInput interface {
	pulumi.Input

	ToAvroSerializationOutput() AvroSerializationOutput
	ToAvroSerializationOutputWithContext(context.Context) AvroSerializationOutput
}

type AvroSerializationArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (AvroSerializationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvroSerialization)(nil)).Elem()
}

func (i AvroSerializationArgs) ToAvroSerializationOutput() AvroSerializationOutput {
	return i.ToAvroSerializationOutputWithContext(context.Background())
}

func (i AvroSerializationArgs) ToAvroSerializationOutputWithContext(ctx context.Context) AvroSerializationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvroSerializationOutput)
}

type AvroSerializationOutput struct{ *pulumi.OutputState }

func (AvroSerializationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvroSerialization)(nil)).Elem()
}

func (o AvroSerializationOutput) ToAvroSerializationOutput() AvroSerializationOutput {
	return o
}

func (o AvroSerializationOutput) ToAvroSerializationOutputWithContext(ctx context.Context) AvroSerializationOutput {
	return o
}

func (o AvroSerializationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AvroSerialization) string { return v.Type }).(pulumi.StringOutput)
}

type AvroSerializationResponse struct {
	Type string `pulumi:"type"`
}





type AvroSerializationResponseInput interface {
	pulumi.Input

	ToAvroSerializationResponseOutput() AvroSerializationResponseOutput
	ToAvroSerializationResponseOutputWithContext(context.Context) AvroSerializationResponseOutput
}

type AvroSerializationResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (AvroSerializationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvroSerializationResponse)(nil)).Elem()
}

func (i AvroSerializationResponseArgs) ToAvroSerializationResponseOutput() AvroSerializationResponseOutput {
	return i.ToAvroSerializationResponseOutputWithContext(context.Background())
}

func (i AvroSerializationResponseArgs) ToAvroSerializationResponseOutputWithContext(ctx context.Context) AvroSerializationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvroSerializationResponseOutput)
}

type AvroSerializationResponseOutput struct{ *pulumi.OutputState }

func (AvroSerializationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvroSerializationResponse)(nil)).Elem()
}

func (o AvroSerializationResponseOutput) ToAvroSerializationResponseOutput() AvroSerializationResponseOutput {
	return o
}

func (o AvroSerializationResponseOutput) ToAvroSerializationResponseOutputWithContext(ctx context.Context) AvroSerializationResponseOutput {
	return o
}

func (o AvroSerializationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AvroSerializationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureDataLakeStoreOutputDataSource struct {
	AccountName            *string `pulumi:"accountName"`
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	DateFormat             *string `pulumi:"dateFormat"`
	FilePathPrefix         *string `pulumi:"filePathPrefix"`
	RefreshToken           *string `pulumi:"refreshToken"`
	TenantId               *string `pulumi:"tenantId"`
	TimeFormat             *string `pulumi:"timeFormat"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}





type AzureDataLakeStoreOutputDataSourceInput interface {
	pulumi.Input

	ToAzureDataLakeStoreOutputDataSourceOutput() AzureDataLakeStoreOutputDataSourceOutput
	ToAzureDataLakeStoreOutputDataSourceOutputWithContext(context.Context) AzureDataLakeStoreOutputDataSourceOutput
}

type AzureDataLakeStoreOutputDataSourceArgs struct {
	AccountName            pulumi.StringPtrInput `pulumi:"accountName"`
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	DateFormat             pulumi.StringPtrInput `pulumi:"dateFormat"`
	FilePathPrefix         pulumi.StringPtrInput `pulumi:"filePathPrefix"`
	RefreshToken           pulumi.StringPtrInput `pulumi:"refreshToken"`
	TenantId               pulumi.StringPtrInput `pulumi:"tenantId"`
	TimeFormat             pulumi.StringPtrInput `pulumi:"timeFormat"`
	TokenUserDisplayName   pulumi.StringPtrInput `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName pulumi.StringPtrInput `pulumi:"tokenUserPrincipalName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (AzureDataLakeStoreOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataLakeStoreOutputDataSource)(nil)).Elem()
}

func (i AzureDataLakeStoreOutputDataSourceArgs) ToAzureDataLakeStoreOutputDataSourceOutput() AzureDataLakeStoreOutputDataSourceOutput {
	return i.ToAzureDataLakeStoreOutputDataSourceOutputWithContext(context.Background())
}

func (i AzureDataLakeStoreOutputDataSourceArgs) ToAzureDataLakeStoreOutputDataSourceOutputWithContext(ctx context.Context) AzureDataLakeStoreOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDataLakeStoreOutputDataSourceOutput)
}

type AzureDataLakeStoreOutputDataSourceOutput struct{ *pulumi.OutputState }

func (AzureDataLakeStoreOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataLakeStoreOutputDataSource)(nil)).Elem()
}

func (o AzureDataLakeStoreOutputDataSourceOutput) ToAzureDataLakeStoreOutputDataSourceOutput() AzureDataLakeStoreOutputDataSourceOutput {
	return o
}

func (o AzureDataLakeStoreOutputDataSourceOutput) ToAzureDataLakeStoreOutputDataSourceOutputWithContext(ctx context.Context) AzureDataLakeStoreOutputDataSourceOutput {
	return o
}

func (o AzureDataLakeStoreOutputDataSourceOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) FilePathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.FilePathPrefix }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) TokenUserDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.TokenUserDisplayName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) TokenUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) *string { return v.TokenUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type AzureDataLakeStoreOutputDataSourceResponse struct {
	AccountName            *string `pulumi:"accountName"`
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	DateFormat             *string `pulumi:"dateFormat"`
	FilePathPrefix         *string `pulumi:"filePathPrefix"`
	RefreshToken           *string `pulumi:"refreshToken"`
	TenantId               *string `pulumi:"tenantId"`
	TimeFormat             *string `pulumi:"timeFormat"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}





type AzureDataLakeStoreOutputDataSourceResponseInput interface {
	pulumi.Input

	ToAzureDataLakeStoreOutputDataSourceResponseOutput() AzureDataLakeStoreOutputDataSourceResponseOutput
	ToAzureDataLakeStoreOutputDataSourceResponseOutputWithContext(context.Context) AzureDataLakeStoreOutputDataSourceResponseOutput
}

type AzureDataLakeStoreOutputDataSourceResponseArgs struct {
	AccountName            pulumi.StringPtrInput `pulumi:"accountName"`
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	DateFormat             pulumi.StringPtrInput `pulumi:"dateFormat"`
	FilePathPrefix         pulumi.StringPtrInput `pulumi:"filePathPrefix"`
	RefreshToken           pulumi.StringPtrInput `pulumi:"refreshToken"`
	TenantId               pulumi.StringPtrInput `pulumi:"tenantId"`
	TimeFormat             pulumi.StringPtrInput `pulumi:"timeFormat"`
	TokenUserDisplayName   pulumi.StringPtrInput `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName pulumi.StringPtrInput `pulumi:"tokenUserPrincipalName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (AzureDataLakeStoreOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataLakeStoreOutputDataSourceResponse)(nil)).Elem()
}

func (i AzureDataLakeStoreOutputDataSourceResponseArgs) ToAzureDataLakeStoreOutputDataSourceResponseOutput() AzureDataLakeStoreOutputDataSourceResponseOutput {
	return i.ToAzureDataLakeStoreOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i AzureDataLakeStoreOutputDataSourceResponseArgs) ToAzureDataLakeStoreOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureDataLakeStoreOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDataLakeStoreOutputDataSourceResponseOutput)
}

type AzureDataLakeStoreOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (AzureDataLakeStoreOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDataLakeStoreOutputDataSourceResponse)(nil)).Elem()
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) ToAzureDataLakeStoreOutputDataSourceResponseOutput() AzureDataLakeStoreOutputDataSourceResponseOutput {
	return o
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) ToAzureDataLakeStoreOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureDataLakeStoreOutputDataSourceResponseOutput {
	return o
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) FilePathPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.FilePathPrefix }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) TokenUserDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.TokenUserDisplayName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) TokenUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) *string { return v.TokenUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o AzureDataLakeStoreOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureDataLakeStoreOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureFunctionOutputDataSource struct {
	ApiKey          *string  `pulumi:"apiKey"`
	FunctionAppName *string  `pulumi:"functionAppName"`
	FunctionName    *string  `pulumi:"functionName"`
	MaxBatchCount   *float64 `pulumi:"maxBatchCount"`
	MaxBatchSize    *float64 `pulumi:"maxBatchSize"`
	Type            string   `pulumi:"type"`
}





type AzureFunctionOutputDataSourceInput interface {
	pulumi.Input

	ToAzureFunctionOutputDataSourceOutput() AzureFunctionOutputDataSourceOutput
	ToAzureFunctionOutputDataSourceOutputWithContext(context.Context) AzureFunctionOutputDataSourceOutput
}

type AzureFunctionOutputDataSourceArgs struct {
	ApiKey          pulumi.StringPtrInput  `pulumi:"apiKey"`
	FunctionAppName pulumi.StringPtrInput  `pulumi:"functionAppName"`
	FunctionName    pulumi.StringPtrInput  `pulumi:"functionName"`
	MaxBatchCount   pulumi.Float64PtrInput `pulumi:"maxBatchCount"`
	MaxBatchSize    pulumi.Float64PtrInput `pulumi:"maxBatchSize"`
	Type            pulumi.StringInput     `pulumi:"type"`
}

func (AzureFunctionOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionOutputDataSource)(nil)).Elem()
}

func (i AzureFunctionOutputDataSourceArgs) ToAzureFunctionOutputDataSourceOutput() AzureFunctionOutputDataSourceOutput {
	return i.ToAzureFunctionOutputDataSourceOutputWithContext(context.Background())
}

func (i AzureFunctionOutputDataSourceArgs) ToAzureFunctionOutputDataSourceOutputWithContext(ctx context.Context) AzureFunctionOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionOutputDataSourceOutput)
}

type AzureFunctionOutputDataSourceOutput struct{ *pulumi.OutputState }

func (AzureFunctionOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionOutputDataSource)(nil)).Elem()
}

func (o AzureFunctionOutputDataSourceOutput) ToAzureFunctionOutputDataSourceOutput() AzureFunctionOutputDataSourceOutput {
	return o
}

func (o AzureFunctionOutputDataSourceOutput) ToAzureFunctionOutputDataSourceOutputWithContext(ctx context.Context) AzureFunctionOutputDataSourceOutput {
	return o
}

func (o AzureFunctionOutputDataSourceOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSource) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AzureFunctionOutputDataSourceOutput) FunctionAppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSource) *string { return v.FunctionAppName }).(pulumi.StringPtrOutput)
}

func (o AzureFunctionOutputDataSourceOutput) FunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSource) *string { return v.FunctionName }).(pulumi.StringPtrOutput)
}

func (o AzureFunctionOutputDataSourceOutput) MaxBatchCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSource) *float64 { return v.MaxBatchCount }).(pulumi.Float64PtrOutput)
}

func (o AzureFunctionOutputDataSourceOutput) MaxBatchSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSource) *float64 { return v.MaxBatchSize }).(pulumi.Float64PtrOutput)
}

func (o AzureFunctionOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type AzureFunctionOutputDataSourceResponse struct {
	ApiKey          *string  `pulumi:"apiKey"`
	FunctionAppName *string  `pulumi:"functionAppName"`
	FunctionName    *string  `pulumi:"functionName"`
	MaxBatchCount   *float64 `pulumi:"maxBatchCount"`
	MaxBatchSize    *float64 `pulumi:"maxBatchSize"`
	Type            string   `pulumi:"type"`
}





type AzureFunctionOutputDataSourceResponseInput interface {
	pulumi.Input

	ToAzureFunctionOutputDataSourceResponseOutput() AzureFunctionOutputDataSourceResponseOutput
	ToAzureFunctionOutputDataSourceResponseOutputWithContext(context.Context) AzureFunctionOutputDataSourceResponseOutput
}

type AzureFunctionOutputDataSourceResponseArgs struct {
	ApiKey          pulumi.StringPtrInput  `pulumi:"apiKey"`
	FunctionAppName pulumi.StringPtrInput  `pulumi:"functionAppName"`
	FunctionName    pulumi.StringPtrInput  `pulumi:"functionName"`
	MaxBatchCount   pulumi.Float64PtrInput `pulumi:"maxBatchCount"`
	MaxBatchSize    pulumi.Float64PtrInput `pulumi:"maxBatchSize"`
	Type            pulumi.StringInput     `pulumi:"type"`
}

func (AzureFunctionOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionOutputDataSourceResponse)(nil)).Elem()
}

func (i AzureFunctionOutputDataSourceResponseArgs) ToAzureFunctionOutputDataSourceResponseOutput() AzureFunctionOutputDataSourceResponseOutput {
	return i.ToAzureFunctionOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i AzureFunctionOutputDataSourceResponseArgs) ToAzureFunctionOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureFunctionOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFunctionOutputDataSourceResponseOutput)
}

type AzureFunctionOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (AzureFunctionOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFunctionOutputDataSourceResponse)(nil)).Elem()
}

func (o AzureFunctionOutputDataSourceResponseOutput) ToAzureFunctionOutputDataSourceResponseOutput() AzureFunctionOutputDataSourceResponseOutput {
	return o
}

func (o AzureFunctionOutputDataSourceResponseOutput) ToAzureFunctionOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureFunctionOutputDataSourceResponseOutput {
	return o
}

func (o AzureFunctionOutputDataSourceResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSourceResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AzureFunctionOutputDataSourceResponseOutput) FunctionAppName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSourceResponse) *string { return v.FunctionAppName }).(pulumi.StringPtrOutput)
}

func (o AzureFunctionOutputDataSourceResponseOutput) FunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSourceResponse) *string { return v.FunctionName }).(pulumi.StringPtrOutput)
}

func (o AzureFunctionOutputDataSourceResponseOutput) MaxBatchCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSourceResponse) *float64 { return v.MaxBatchCount }).(pulumi.Float64PtrOutput)
}

func (o AzureFunctionOutputDataSourceResponseOutput) MaxBatchSize() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSourceResponse) *float64 { return v.MaxBatchSize }).(pulumi.Float64PtrOutput)
}

func (o AzureFunctionOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFunctionOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureMachineLearningServiceFunctionBinding struct {
	ApiKey                   *string                                   `pulumi:"apiKey"`
	BatchSize                *int                                      `pulumi:"batchSize"`
	Endpoint                 *string                                   `pulumi:"endpoint"`
	Inputs                   []AzureMachineLearningServiceInputColumn  `pulumi:"inputs"`
	NumberOfParallelRequests *int                                      `pulumi:"numberOfParallelRequests"`
	Outputs                  []AzureMachineLearningServiceOutputColumn `pulumi:"outputs"`
	Type                     string                                    `pulumi:"type"`
}





type AzureMachineLearningServiceFunctionBindingInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceFunctionBindingOutput() AzureMachineLearningServiceFunctionBindingOutput
	ToAzureMachineLearningServiceFunctionBindingOutputWithContext(context.Context) AzureMachineLearningServiceFunctionBindingOutput
}

type AzureMachineLearningServiceFunctionBindingArgs struct {
	ApiKey                   pulumi.StringPtrInput                             `pulumi:"apiKey"`
	BatchSize                pulumi.IntPtrInput                                `pulumi:"batchSize"`
	Endpoint                 pulumi.StringPtrInput                             `pulumi:"endpoint"`
	Inputs                   AzureMachineLearningServiceInputColumnArrayInput  `pulumi:"inputs"`
	NumberOfParallelRequests pulumi.IntPtrInput                                `pulumi:"numberOfParallelRequests"`
	Outputs                  AzureMachineLearningServiceOutputColumnArrayInput `pulumi:"outputs"`
	Type                     pulumi.StringInput                                `pulumi:"type"`
}

func (AzureMachineLearningServiceFunctionBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceFunctionBinding)(nil)).Elem()
}

func (i AzureMachineLearningServiceFunctionBindingArgs) ToAzureMachineLearningServiceFunctionBindingOutput() AzureMachineLearningServiceFunctionBindingOutput {
	return i.ToAzureMachineLearningServiceFunctionBindingOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceFunctionBindingArgs) ToAzureMachineLearningServiceFunctionBindingOutputWithContext(ctx context.Context) AzureMachineLearningServiceFunctionBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceFunctionBindingOutput)
}

type AzureMachineLearningServiceFunctionBindingOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceFunctionBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceFunctionBinding)(nil)).Elem()
}

func (o AzureMachineLearningServiceFunctionBindingOutput) ToAzureMachineLearningServiceFunctionBindingOutput() AzureMachineLearningServiceFunctionBindingOutput {
	return o
}

func (o AzureMachineLearningServiceFunctionBindingOutput) ToAzureMachineLearningServiceFunctionBindingOutputWithContext(ctx context.Context) AzureMachineLearningServiceFunctionBindingOutput {
	return o
}

func (o AzureMachineLearningServiceFunctionBindingOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingOutput) Inputs() AzureMachineLearningServiceInputColumnArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) []AzureMachineLearningServiceInputColumn {
		return v.Inputs
	}).(AzureMachineLearningServiceInputColumnArrayOutput)
}

func (o AzureMachineLearningServiceFunctionBindingOutput) NumberOfParallelRequests() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) *int { return v.NumberOfParallelRequests }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingOutput) Outputs() AzureMachineLearningServiceOutputColumnArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) []AzureMachineLearningServiceOutputColumn {
		return v.Outputs
	}).(AzureMachineLearningServiceOutputColumnArrayOutput)
}

func (o AzureMachineLearningServiceFunctionBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBinding) string { return v.Type }).(pulumi.StringOutput)
}

type AzureMachineLearningServiceFunctionBindingResponse struct {
	ApiKey                   *string                                           `pulumi:"apiKey"`
	BatchSize                *int                                              `pulumi:"batchSize"`
	Endpoint                 *string                                           `pulumi:"endpoint"`
	Inputs                   []AzureMachineLearningServiceInputColumnResponse  `pulumi:"inputs"`
	NumberOfParallelRequests *int                                              `pulumi:"numberOfParallelRequests"`
	Outputs                  []AzureMachineLearningServiceOutputColumnResponse `pulumi:"outputs"`
	Type                     string                                            `pulumi:"type"`
}





type AzureMachineLearningServiceFunctionBindingResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceFunctionBindingResponseOutput() AzureMachineLearningServiceFunctionBindingResponseOutput
	ToAzureMachineLearningServiceFunctionBindingResponseOutputWithContext(context.Context) AzureMachineLearningServiceFunctionBindingResponseOutput
}

type AzureMachineLearningServiceFunctionBindingResponseArgs struct {
	ApiKey                   pulumi.StringPtrInput                                     `pulumi:"apiKey"`
	BatchSize                pulumi.IntPtrInput                                        `pulumi:"batchSize"`
	Endpoint                 pulumi.StringPtrInput                                     `pulumi:"endpoint"`
	Inputs                   AzureMachineLearningServiceInputColumnResponseArrayInput  `pulumi:"inputs"`
	NumberOfParallelRequests pulumi.IntPtrInput                                        `pulumi:"numberOfParallelRequests"`
	Outputs                  AzureMachineLearningServiceOutputColumnResponseArrayInput `pulumi:"outputs"`
	Type                     pulumi.StringInput                                        `pulumi:"type"`
}

func (AzureMachineLearningServiceFunctionBindingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceFunctionBindingResponse)(nil)).Elem()
}

func (i AzureMachineLearningServiceFunctionBindingResponseArgs) ToAzureMachineLearningServiceFunctionBindingResponseOutput() AzureMachineLearningServiceFunctionBindingResponseOutput {
	return i.ToAzureMachineLearningServiceFunctionBindingResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceFunctionBindingResponseArgs) ToAzureMachineLearningServiceFunctionBindingResponseOutputWithContext(ctx context.Context) AzureMachineLearningServiceFunctionBindingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceFunctionBindingResponseOutput)
}

type AzureMachineLearningServiceFunctionBindingResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceFunctionBindingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceFunctionBindingResponse)(nil)).Elem()
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) ToAzureMachineLearningServiceFunctionBindingResponseOutput() AzureMachineLearningServiceFunctionBindingResponseOutput {
	return o
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) ToAzureMachineLearningServiceFunctionBindingResponseOutputWithContext(ctx context.Context) AzureMachineLearningServiceFunctionBindingResponseOutput {
	return o
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) Inputs() AzureMachineLearningServiceInputColumnResponseArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) []AzureMachineLearningServiceInputColumnResponse {
		return v.Inputs
	}).(AzureMachineLearningServiceInputColumnResponseArrayOutput)
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) NumberOfParallelRequests() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) *int { return v.NumberOfParallelRequests }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) Outputs() AzureMachineLearningServiceOutputColumnResponseArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) []AzureMachineLearningServiceOutputColumnResponse {
		return v.Outputs
	}).(AzureMachineLearningServiceOutputColumnResponseArrayOutput)
}

func (o AzureMachineLearningServiceFunctionBindingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceFunctionBindingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureMachineLearningServiceInputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningServiceInputColumnInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceInputColumnOutput() AzureMachineLearningServiceInputColumnOutput
	ToAzureMachineLearningServiceInputColumnOutputWithContext(context.Context) AzureMachineLearningServiceInputColumnOutput
}

type AzureMachineLearningServiceInputColumnArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	MapTo    pulumi.IntPtrInput    `pulumi:"mapTo"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningServiceInputColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceInputColumn)(nil)).Elem()
}

func (i AzureMachineLearningServiceInputColumnArgs) ToAzureMachineLearningServiceInputColumnOutput() AzureMachineLearningServiceInputColumnOutput {
	return i.ToAzureMachineLearningServiceInputColumnOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceInputColumnArgs) ToAzureMachineLearningServiceInputColumnOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceInputColumnOutput)
}





type AzureMachineLearningServiceInputColumnArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceInputColumnArrayOutput() AzureMachineLearningServiceInputColumnArrayOutput
	ToAzureMachineLearningServiceInputColumnArrayOutputWithContext(context.Context) AzureMachineLearningServiceInputColumnArrayOutput
}

type AzureMachineLearningServiceInputColumnArray []AzureMachineLearningServiceInputColumnInput

func (AzureMachineLearningServiceInputColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceInputColumn)(nil)).Elem()
}

func (i AzureMachineLearningServiceInputColumnArray) ToAzureMachineLearningServiceInputColumnArrayOutput() AzureMachineLearningServiceInputColumnArrayOutput {
	return i.ToAzureMachineLearningServiceInputColumnArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceInputColumnArray) ToAzureMachineLearningServiceInputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceInputColumnArrayOutput)
}

type AzureMachineLearningServiceInputColumnOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceInputColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceInputColumn)(nil)).Elem()
}

func (o AzureMachineLearningServiceInputColumnOutput) ToAzureMachineLearningServiceInputColumnOutput() AzureMachineLearningServiceInputColumnOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnOutput) ToAzureMachineLearningServiceInputColumnOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceInputColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceInputColumnOutput) MapTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceInputColumn) *int { return v.MapTo }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceInputColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceInputColumn) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningServiceInputColumnArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceInputColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceInputColumn)(nil)).Elem()
}

func (o AzureMachineLearningServiceInputColumnArrayOutput) ToAzureMachineLearningServiceInputColumnArrayOutput() AzureMachineLearningServiceInputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnArrayOutput) ToAzureMachineLearningServiceInputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningServiceInputColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningServiceInputColumn {
		return vs[0].([]AzureMachineLearningServiceInputColumn)[vs[1].(int)]
	}).(AzureMachineLearningServiceInputColumnOutput)
}

type AzureMachineLearningServiceInputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningServiceInputColumnResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceInputColumnResponseOutput() AzureMachineLearningServiceInputColumnResponseOutput
	ToAzureMachineLearningServiceInputColumnResponseOutputWithContext(context.Context) AzureMachineLearningServiceInputColumnResponseOutput
}

type AzureMachineLearningServiceInputColumnResponseArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	MapTo    pulumi.IntPtrInput    `pulumi:"mapTo"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningServiceInputColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceInputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningServiceInputColumnResponseArgs) ToAzureMachineLearningServiceInputColumnResponseOutput() AzureMachineLearningServiceInputColumnResponseOutput {
	return i.ToAzureMachineLearningServiceInputColumnResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceInputColumnResponseArgs) ToAzureMachineLearningServiceInputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceInputColumnResponseOutput)
}





type AzureMachineLearningServiceInputColumnResponseArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceInputColumnResponseArrayOutput() AzureMachineLearningServiceInputColumnResponseArrayOutput
	ToAzureMachineLearningServiceInputColumnResponseArrayOutputWithContext(context.Context) AzureMachineLearningServiceInputColumnResponseArrayOutput
}

type AzureMachineLearningServiceInputColumnResponseArray []AzureMachineLearningServiceInputColumnResponseInput

func (AzureMachineLearningServiceInputColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceInputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningServiceInputColumnResponseArray) ToAzureMachineLearningServiceInputColumnResponseArrayOutput() AzureMachineLearningServiceInputColumnResponseArrayOutput {
	return i.ToAzureMachineLearningServiceInputColumnResponseArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceInputColumnResponseArray) ToAzureMachineLearningServiceInputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceInputColumnResponseArrayOutput)
}

type AzureMachineLearningServiceInputColumnResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceInputColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceInputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningServiceInputColumnResponseOutput) ToAzureMachineLearningServiceInputColumnResponseOutput() AzureMachineLearningServiceInputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnResponseOutput) ToAzureMachineLearningServiceInputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceInputColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceInputColumnResponseOutput) MapTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceInputColumnResponse) *int { return v.MapTo }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceInputColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceInputColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningServiceInputColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceInputColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceInputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningServiceInputColumnResponseArrayOutput) ToAzureMachineLearningServiceInputColumnResponseArrayOutput() AzureMachineLearningServiceInputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnResponseArrayOutput) ToAzureMachineLearningServiceInputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceInputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningServiceInputColumnResponseArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningServiceInputColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningServiceInputColumnResponse {
		return vs[0].([]AzureMachineLearningServiceInputColumnResponse)[vs[1].(int)]
	}).(AzureMachineLearningServiceInputColumnResponseOutput)
}

type AzureMachineLearningServiceOutputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningServiceOutputColumnInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceOutputColumnOutput() AzureMachineLearningServiceOutputColumnOutput
	ToAzureMachineLearningServiceOutputColumnOutputWithContext(context.Context) AzureMachineLearningServiceOutputColumnOutput
}

type AzureMachineLearningServiceOutputColumnArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	MapTo    pulumi.IntPtrInput    `pulumi:"mapTo"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningServiceOutputColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceOutputColumn)(nil)).Elem()
}

func (i AzureMachineLearningServiceOutputColumnArgs) ToAzureMachineLearningServiceOutputColumnOutput() AzureMachineLearningServiceOutputColumnOutput {
	return i.ToAzureMachineLearningServiceOutputColumnOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceOutputColumnArgs) ToAzureMachineLearningServiceOutputColumnOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceOutputColumnOutput)
}





type AzureMachineLearningServiceOutputColumnArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceOutputColumnArrayOutput() AzureMachineLearningServiceOutputColumnArrayOutput
	ToAzureMachineLearningServiceOutputColumnArrayOutputWithContext(context.Context) AzureMachineLearningServiceOutputColumnArrayOutput
}

type AzureMachineLearningServiceOutputColumnArray []AzureMachineLearningServiceOutputColumnInput

func (AzureMachineLearningServiceOutputColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceOutputColumn)(nil)).Elem()
}

func (i AzureMachineLearningServiceOutputColumnArray) ToAzureMachineLearningServiceOutputColumnArrayOutput() AzureMachineLearningServiceOutputColumnArrayOutput {
	return i.ToAzureMachineLearningServiceOutputColumnArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceOutputColumnArray) ToAzureMachineLearningServiceOutputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceOutputColumnArrayOutput)
}

type AzureMachineLearningServiceOutputColumnOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceOutputColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceOutputColumn)(nil)).Elem()
}

func (o AzureMachineLearningServiceOutputColumnOutput) ToAzureMachineLearningServiceOutputColumnOutput() AzureMachineLearningServiceOutputColumnOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnOutput) ToAzureMachineLearningServiceOutputColumnOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceOutputColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceOutputColumnOutput) MapTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceOutputColumn) *int { return v.MapTo }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceOutputColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceOutputColumn) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningServiceOutputColumnArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceOutputColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceOutputColumn)(nil)).Elem()
}

func (o AzureMachineLearningServiceOutputColumnArrayOutput) ToAzureMachineLearningServiceOutputColumnArrayOutput() AzureMachineLearningServiceOutputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnArrayOutput) ToAzureMachineLearningServiceOutputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningServiceOutputColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningServiceOutputColumn {
		return vs[0].([]AzureMachineLearningServiceOutputColumn)[vs[1].(int)]
	}).(AzureMachineLearningServiceOutputColumnOutput)
}

type AzureMachineLearningServiceOutputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningServiceOutputColumnResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceOutputColumnResponseOutput() AzureMachineLearningServiceOutputColumnResponseOutput
	ToAzureMachineLearningServiceOutputColumnResponseOutputWithContext(context.Context) AzureMachineLearningServiceOutputColumnResponseOutput
}

type AzureMachineLearningServiceOutputColumnResponseArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	MapTo    pulumi.IntPtrInput    `pulumi:"mapTo"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningServiceOutputColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceOutputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningServiceOutputColumnResponseArgs) ToAzureMachineLearningServiceOutputColumnResponseOutput() AzureMachineLearningServiceOutputColumnResponseOutput {
	return i.ToAzureMachineLearningServiceOutputColumnResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceOutputColumnResponseArgs) ToAzureMachineLearningServiceOutputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceOutputColumnResponseOutput)
}





type AzureMachineLearningServiceOutputColumnResponseArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningServiceOutputColumnResponseArrayOutput() AzureMachineLearningServiceOutputColumnResponseArrayOutput
	ToAzureMachineLearningServiceOutputColumnResponseArrayOutputWithContext(context.Context) AzureMachineLearningServiceOutputColumnResponseArrayOutput
}

type AzureMachineLearningServiceOutputColumnResponseArray []AzureMachineLearningServiceOutputColumnResponseInput

func (AzureMachineLearningServiceOutputColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceOutputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningServiceOutputColumnResponseArray) ToAzureMachineLearningServiceOutputColumnResponseArrayOutput() AzureMachineLearningServiceOutputColumnResponseArrayOutput {
	return i.ToAzureMachineLearningServiceOutputColumnResponseArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningServiceOutputColumnResponseArray) ToAzureMachineLearningServiceOutputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningServiceOutputColumnResponseArrayOutput)
}

type AzureMachineLearningServiceOutputColumnResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceOutputColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningServiceOutputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningServiceOutputColumnResponseOutput) ToAzureMachineLearningServiceOutputColumnResponseOutput() AzureMachineLearningServiceOutputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnResponseOutput) ToAzureMachineLearningServiceOutputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceOutputColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningServiceOutputColumnResponseOutput) MapTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceOutputColumnResponse) *int { return v.MapTo }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningServiceOutputColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningServiceOutputColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningServiceOutputColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningServiceOutputColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningServiceOutputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningServiceOutputColumnResponseArrayOutput) ToAzureMachineLearningServiceOutputColumnResponseArrayOutput() AzureMachineLearningServiceOutputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnResponseArrayOutput) ToAzureMachineLearningServiceOutputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningServiceOutputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningServiceOutputColumnResponseArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningServiceOutputColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningServiceOutputColumnResponse {
		return vs[0].([]AzureMachineLearningServiceOutputColumnResponse)[vs[1].(int)]
	}).(AzureMachineLearningServiceOutputColumnResponseOutput)
}

type AzureMachineLearningStudioFunctionBinding struct {
	ApiKey    *string                                  `pulumi:"apiKey"`
	BatchSize *int                                     `pulumi:"batchSize"`
	Endpoint  *string                                  `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningStudioInputs        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningStudioOutputColumn `pulumi:"outputs"`
	Type      string                                   `pulumi:"type"`
}





type AzureMachineLearningStudioFunctionBindingInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioFunctionBindingOutput() AzureMachineLearningStudioFunctionBindingOutput
	ToAzureMachineLearningStudioFunctionBindingOutputWithContext(context.Context) AzureMachineLearningStudioFunctionBindingOutput
}

type AzureMachineLearningStudioFunctionBindingArgs struct {
	ApiKey    pulumi.StringPtrInput                            `pulumi:"apiKey"`
	BatchSize pulumi.IntPtrInput                               `pulumi:"batchSize"`
	Endpoint  pulumi.StringPtrInput                            `pulumi:"endpoint"`
	Inputs    AzureMachineLearningStudioInputsPtrInput         `pulumi:"inputs"`
	Outputs   AzureMachineLearningStudioOutputColumnArrayInput `pulumi:"outputs"`
	Type      pulumi.StringInput                               `pulumi:"type"`
}

func (AzureMachineLearningStudioFunctionBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioFunctionBinding)(nil)).Elem()
}

func (i AzureMachineLearningStudioFunctionBindingArgs) ToAzureMachineLearningStudioFunctionBindingOutput() AzureMachineLearningStudioFunctionBindingOutput {
	return i.ToAzureMachineLearningStudioFunctionBindingOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioFunctionBindingArgs) ToAzureMachineLearningStudioFunctionBindingOutputWithContext(ctx context.Context) AzureMachineLearningStudioFunctionBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioFunctionBindingOutput)
}

type AzureMachineLearningStudioFunctionBindingOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioFunctionBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioFunctionBinding)(nil)).Elem()
}

func (o AzureMachineLearningStudioFunctionBindingOutput) ToAzureMachineLearningStudioFunctionBindingOutput() AzureMachineLearningStudioFunctionBindingOutput {
	return o
}

func (o AzureMachineLearningStudioFunctionBindingOutput) ToAzureMachineLearningStudioFunctionBindingOutputWithContext(ctx context.Context) AzureMachineLearningStudioFunctionBindingOutput {
	return o
}

func (o AzureMachineLearningStudioFunctionBindingOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBinding) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBinding) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBinding) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingOutput) Inputs() AzureMachineLearningStudioInputsPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBinding) *AzureMachineLearningStudioInputs { return v.Inputs }).(AzureMachineLearningStudioInputsPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingOutput) Outputs() AzureMachineLearningStudioOutputColumnArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBinding) []AzureMachineLearningStudioOutputColumn {
		return v.Outputs
	}).(AzureMachineLearningStudioOutputColumnArrayOutput)
}

func (o AzureMachineLearningStudioFunctionBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBinding) string { return v.Type }).(pulumi.StringOutput)
}

type AzureMachineLearningStudioFunctionBindingResponse struct {
	ApiKey    *string                                          `pulumi:"apiKey"`
	BatchSize *int                                             `pulumi:"batchSize"`
	Endpoint  *string                                          `pulumi:"endpoint"`
	Inputs    *AzureMachineLearningStudioInputsResponse        `pulumi:"inputs"`
	Outputs   []AzureMachineLearningStudioOutputColumnResponse `pulumi:"outputs"`
	Type      string                                           `pulumi:"type"`
}





type AzureMachineLearningStudioFunctionBindingResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioFunctionBindingResponseOutput() AzureMachineLearningStudioFunctionBindingResponseOutput
	ToAzureMachineLearningStudioFunctionBindingResponseOutputWithContext(context.Context) AzureMachineLearningStudioFunctionBindingResponseOutput
}

type AzureMachineLearningStudioFunctionBindingResponseArgs struct {
	ApiKey    pulumi.StringPtrInput                                    `pulumi:"apiKey"`
	BatchSize pulumi.IntPtrInput                                       `pulumi:"batchSize"`
	Endpoint  pulumi.StringPtrInput                                    `pulumi:"endpoint"`
	Inputs    AzureMachineLearningStudioInputsResponsePtrInput         `pulumi:"inputs"`
	Outputs   AzureMachineLearningStudioOutputColumnResponseArrayInput `pulumi:"outputs"`
	Type      pulumi.StringInput                                       `pulumi:"type"`
}

func (AzureMachineLearningStudioFunctionBindingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioFunctionBindingResponse)(nil)).Elem()
}

func (i AzureMachineLearningStudioFunctionBindingResponseArgs) ToAzureMachineLearningStudioFunctionBindingResponseOutput() AzureMachineLearningStudioFunctionBindingResponseOutput {
	return i.ToAzureMachineLearningStudioFunctionBindingResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioFunctionBindingResponseArgs) ToAzureMachineLearningStudioFunctionBindingResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioFunctionBindingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioFunctionBindingResponseOutput)
}

type AzureMachineLearningStudioFunctionBindingResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioFunctionBindingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioFunctionBindingResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) ToAzureMachineLearningStudioFunctionBindingResponseOutput() AzureMachineLearningStudioFunctionBindingResponseOutput {
	return o
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) ToAzureMachineLearningStudioFunctionBindingResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioFunctionBindingResponseOutput {
	return o
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBindingResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBindingResponse) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBindingResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) Inputs() AzureMachineLearningStudioInputsResponsePtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBindingResponse) *AzureMachineLearningStudioInputsResponse {
		return v.Inputs
	}).(AzureMachineLearningStudioInputsResponsePtrOutput)
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) Outputs() AzureMachineLearningStudioOutputColumnResponseArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBindingResponse) []AzureMachineLearningStudioOutputColumnResponse {
		return v.Outputs
	}).(AzureMachineLearningStudioOutputColumnResponseArrayOutput)
}

func (o AzureMachineLearningStudioFunctionBindingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioFunctionBindingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type AzureMachineLearningStudioInputColumn struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningStudioInputColumnInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputColumnOutput() AzureMachineLearningStudioInputColumnOutput
	ToAzureMachineLearningStudioInputColumnOutputWithContext(context.Context) AzureMachineLearningStudioInputColumnOutput
}

type AzureMachineLearningStudioInputColumnArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	MapTo    pulumi.IntPtrInput    `pulumi:"mapTo"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningStudioInputColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputColumn)(nil)).Elem()
}

func (i AzureMachineLearningStudioInputColumnArgs) ToAzureMachineLearningStudioInputColumnOutput() AzureMachineLearningStudioInputColumnOutput {
	return i.ToAzureMachineLearningStudioInputColumnOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputColumnArgs) ToAzureMachineLearningStudioInputColumnOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputColumnOutput)
}





type AzureMachineLearningStudioInputColumnArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputColumnArrayOutput() AzureMachineLearningStudioInputColumnArrayOutput
	ToAzureMachineLearningStudioInputColumnArrayOutputWithContext(context.Context) AzureMachineLearningStudioInputColumnArrayOutput
}

type AzureMachineLearningStudioInputColumnArray []AzureMachineLearningStudioInputColumnInput

func (AzureMachineLearningStudioInputColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioInputColumn)(nil)).Elem()
}

func (i AzureMachineLearningStudioInputColumnArray) ToAzureMachineLearningStudioInputColumnArrayOutput() AzureMachineLearningStudioInputColumnArrayOutput {
	return i.ToAzureMachineLearningStudioInputColumnArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputColumnArray) ToAzureMachineLearningStudioInputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputColumnArrayOutput)
}

type AzureMachineLearningStudioInputColumnOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputColumn)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputColumnOutput) ToAzureMachineLearningStudioInputColumnOutput() AzureMachineLearningStudioInputColumnOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnOutput) ToAzureMachineLearningStudioInputColumnOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioInputColumnOutput) MapTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputColumn) *int { return v.MapTo }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningStudioInputColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputColumn) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioInputColumnArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioInputColumn)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputColumnArrayOutput) ToAzureMachineLearningStudioInputColumnArrayOutput() AzureMachineLearningStudioInputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnArrayOutput) ToAzureMachineLearningStudioInputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningStudioInputColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningStudioInputColumn {
		return vs[0].([]AzureMachineLearningStudioInputColumn)[vs[1].(int)]
	}).(AzureMachineLearningStudioInputColumnOutput)
}

type AzureMachineLearningStudioInputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	MapTo    *int    `pulumi:"mapTo"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningStudioInputColumnResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputColumnResponseOutput() AzureMachineLearningStudioInputColumnResponseOutput
	ToAzureMachineLearningStudioInputColumnResponseOutputWithContext(context.Context) AzureMachineLearningStudioInputColumnResponseOutput
}

type AzureMachineLearningStudioInputColumnResponseArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	MapTo    pulumi.IntPtrInput    `pulumi:"mapTo"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningStudioInputColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningStudioInputColumnResponseArgs) ToAzureMachineLearningStudioInputColumnResponseOutput() AzureMachineLearningStudioInputColumnResponseOutput {
	return i.ToAzureMachineLearningStudioInputColumnResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputColumnResponseArgs) ToAzureMachineLearningStudioInputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputColumnResponseOutput)
}





type AzureMachineLearningStudioInputColumnResponseArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputColumnResponseArrayOutput() AzureMachineLearningStudioInputColumnResponseArrayOutput
	ToAzureMachineLearningStudioInputColumnResponseArrayOutputWithContext(context.Context) AzureMachineLearningStudioInputColumnResponseArrayOutput
}

type AzureMachineLearningStudioInputColumnResponseArray []AzureMachineLearningStudioInputColumnResponseInput

func (AzureMachineLearningStudioInputColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioInputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningStudioInputColumnResponseArray) ToAzureMachineLearningStudioInputColumnResponseArrayOutput() AzureMachineLearningStudioInputColumnResponseArrayOutput {
	return i.ToAzureMachineLearningStudioInputColumnResponseArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputColumnResponseArray) ToAzureMachineLearningStudioInputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputColumnResponseArrayOutput)
}

type AzureMachineLearningStudioInputColumnResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputColumnResponseOutput) ToAzureMachineLearningStudioInputColumnResponseOutput() AzureMachineLearningStudioInputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnResponseOutput) ToAzureMachineLearningStudioInputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioInputColumnResponseOutput) MapTo() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputColumnResponse) *int { return v.MapTo }).(pulumi.IntPtrOutput)
}

func (o AzureMachineLearningStudioInputColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioInputColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioInputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputColumnResponseArrayOutput) ToAzureMachineLearningStudioInputColumnResponseArrayOutput() AzureMachineLearningStudioInputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnResponseArrayOutput) ToAzureMachineLearningStudioInputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningStudioInputColumnResponseArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningStudioInputColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningStudioInputColumnResponse {
		return vs[0].([]AzureMachineLearningStudioInputColumnResponse)[vs[1].(int)]
	}).(AzureMachineLearningStudioInputColumnResponseOutput)
}

type AzureMachineLearningStudioInputs struct {
	ColumnNames []AzureMachineLearningStudioInputColumn `pulumi:"columnNames"`
	Name        *string                                 `pulumi:"name"`
}





type AzureMachineLearningStudioInputsInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputsOutput() AzureMachineLearningStudioInputsOutput
	ToAzureMachineLearningStudioInputsOutputWithContext(context.Context) AzureMachineLearningStudioInputsOutput
}

type AzureMachineLearningStudioInputsArgs struct {
	ColumnNames AzureMachineLearningStudioInputColumnArrayInput `pulumi:"columnNames"`
	Name        pulumi.StringPtrInput                           `pulumi:"name"`
}

func (AzureMachineLearningStudioInputsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputs)(nil)).Elem()
}

func (i AzureMachineLearningStudioInputsArgs) ToAzureMachineLearningStudioInputsOutput() AzureMachineLearningStudioInputsOutput {
	return i.ToAzureMachineLearningStudioInputsOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputsArgs) ToAzureMachineLearningStudioInputsOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputsOutput)
}

func (i AzureMachineLearningStudioInputsArgs) ToAzureMachineLearningStudioInputsPtrOutput() AzureMachineLearningStudioInputsPtrOutput {
	return i.ToAzureMachineLearningStudioInputsPtrOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputsArgs) ToAzureMachineLearningStudioInputsPtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputsOutput).ToAzureMachineLearningStudioInputsPtrOutputWithContext(ctx)
}









type AzureMachineLearningStudioInputsPtrInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputsPtrOutput() AzureMachineLearningStudioInputsPtrOutput
	ToAzureMachineLearningStudioInputsPtrOutputWithContext(context.Context) AzureMachineLearningStudioInputsPtrOutput
}

type azureMachineLearningStudioInputsPtrType AzureMachineLearningStudioInputsArgs

func AzureMachineLearningStudioInputsPtr(v *AzureMachineLearningStudioInputsArgs) AzureMachineLearningStudioInputsPtrInput {
	return (*azureMachineLearningStudioInputsPtrType)(v)
}

func (*azureMachineLearningStudioInputsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMachineLearningStudioInputs)(nil)).Elem()
}

func (i *azureMachineLearningStudioInputsPtrType) ToAzureMachineLearningStudioInputsPtrOutput() AzureMachineLearningStudioInputsPtrOutput {
	return i.ToAzureMachineLearningStudioInputsPtrOutputWithContext(context.Background())
}

func (i *azureMachineLearningStudioInputsPtrType) ToAzureMachineLearningStudioInputsPtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputsPtrOutput)
}

type AzureMachineLearningStudioInputsOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputs)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputsOutput) ToAzureMachineLearningStudioInputsOutput() AzureMachineLearningStudioInputsOutput {
	return o
}

func (o AzureMachineLearningStudioInputsOutput) ToAzureMachineLearningStudioInputsOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsOutput {
	return o
}

func (o AzureMachineLearningStudioInputsOutput) ToAzureMachineLearningStudioInputsPtrOutput() AzureMachineLearningStudioInputsPtrOutput {
	return o.ToAzureMachineLearningStudioInputsPtrOutputWithContext(context.Background())
}

func (o AzureMachineLearningStudioInputsOutput) ToAzureMachineLearningStudioInputsPtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureMachineLearningStudioInputs) *AzureMachineLearningStudioInputs {
		return &v
	}).(AzureMachineLearningStudioInputsPtrOutput)
}

func (o AzureMachineLearningStudioInputsOutput) ColumnNames() AzureMachineLearningStudioInputColumnArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputs) []AzureMachineLearningStudioInputColumn { return v.ColumnNames }).(AzureMachineLearningStudioInputColumnArrayOutput)
}

func (o AzureMachineLearningStudioInputsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputs) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioInputsPtrOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMachineLearningStudioInputs)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputsPtrOutput) ToAzureMachineLearningStudioInputsPtrOutput() AzureMachineLearningStudioInputsPtrOutput {
	return o
}

func (o AzureMachineLearningStudioInputsPtrOutput) ToAzureMachineLearningStudioInputsPtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsPtrOutput {
	return o
}

func (o AzureMachineLearningStudioInputsPtrOutput) Elem() AzureMachineLearningStudioInputsOutput {
	return o.ApplyT(func(v *AzureMachineLearningStudioInputs) AzureMachineLearningStudioInputs {
		if v != nil {
			return *v
		}
		var ret AzureMachineLearningStudioInputs
		return ret
	}).(AzureMachineLearningStudioInputsOutput)
}

func (o AzureMachineLearningStudioInputsPtrOutput) ColumnNames() AzureMachineLearningStudioInputColumnArrayOutput {
	return o.ApplyT(func(v *AzureMachineLearningStudioInputs) []AzureMachineLearningStudioInputColumn {
		if v == nil {
			return nil
		}
		return v.ColumnNames
	}).(AzureMachineLearningStudioInputColumnArrayOutput)
}

func (o AzureMachineLearningStudioInputsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMachineLearningStudioInputs) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioInputsResponse struct {
	ColumnNames []AzureMachineLearningStudioInputColumnResponse `pulumi:"columnNames"`
	Name        *string                                         `pulumi:"name"`
}





type AzureMachineLearningStudioInputsResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputsResponseOutput() AzureMachineLearningStudioInputsResponseOutput
	ToAzureMachineLearningStudioInputsResponseOutputWithContext(context.Context) AzureMachineLearningStudioInputsResponseOutput
}

type AzureMachineLearningStudioInputsResponseArgs struct {
	ColumnNames AzureMachineLearningStudioInputColumnResponseArrayInput `pulumi:"columnNames"`
	Name        pulumi.StringPtrInput                                   `pulumi:"name"`
}

func (AzureMachineLearningStudioInputsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputsResponse)(nil)).Elem()
}

func (i AzureMachineLearningStudioInputsResponseArgs) ToAzureMachineLearningStudioInputsResponseOutput() AzureMachineLearningStudioInputsResponseOutput {
	return i.ToAzureMachineLearningStudioInputsResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputsResponseArgs) ToAzureMachineLearningStudioInputsResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputsResponseOutput)
}

func (i AzureMachineLearningStudioInputsResponseArgs) ToAzureMachineLearningStudioInputsResponsePtrOutput() AzureMachineLearningStudioInputsResponsePtrOutput {
	return i.ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioInputsResponseArgs) ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputsResponseOutput).ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(ctx)
}









type AzureMachineLearningStudioInputsResponsePtrInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioInputsResponsePtrOutput() AzureMachineLearningStudioInputsResponsePtrOutput
	ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(context.Context) AzureMachineLearningStudioInputsResponsePtrOutput
}

type azureMachineLearningStudioInputsResponsePtrType AzureMachineLearningStudioInputsResponseArgs

func AzureMachineLearningStudioInputsResponsePtr(v *AzureMachineLearningStudioInputsResponseArgs) AzureMachineLearningStudioInputsResponsePtrInput {
	return (*azureMachineLearningStudioInputsResponsePtrType)(v)
}

func (*azureMachineLearningStudioInputsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMachineLearningStudioInputsResponse)(nil)).Elem()
}

func (i *azureMachineLearningStudioInputsResponsePtrType) ToAzureMachineLearningStudioInputsResponsePtrOutput() AzureMachineLearningStudioInputsResponsePtrOutput {
	return i.ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(context.Background())
}

func (i *azureMachineLearningStudioInputsResponsePtrType) ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioInputsResponsePtrOutput)
}

type AzureMachineLearningStudioInputsResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioInputsResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputsResponseOutput) ToAzureMachineLearningStudioInputsResponseOutput() AzureMachineLearningStudioInputsResponseOutput {
	return o
}

func (o AzureMachineLearningStudioInputsResponseOutput) ToAzureMachineLearningStudioInputsResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsResponseOutput {
	return o
}

func (o AzureMachineLearningStudioInputsResponseOutput) ToAzureMachineLearningStudioInputsResponsePtrOutput() AzureMachineLearningStudioInputsResponsePtrOutput {
	return o.ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(context.Background())
}

func (o AzureMachineLearningStudioInputsResponseOutput) ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureMachineLearningStudioInputsResponse) *AzureMachineLearningStudioInputsResponse {
		return &v
	}).(AzureMachineLearningStudioInputsResponsePtrOutput)
}

func (o AzureMachineLearningStudioInputsResponseOutput) ColumnNames() AzureMachineLearningStudioInputColumnResponseArrayOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputsResponse) []AzureMachineLearningStudioInputColumnResponse {
		return v.ColumnNames
	}).(AzureMachineLearningStudioInputColumnResponseArrayOutput)
}

func (o AzureMachineLearningStudioInputsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioInputsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioInputsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioInputsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMachineLearningStudioInputsResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioInputsResponsePtrOutput) ToAzureMachineLearningStudioInputsResponsePtrOutput() AzureMachineLearningStudioInputsResponsePtrOutput {
	return o
}

func (o AzureMachineLearningStudioInputsResponsePtrOutput) ToAzureMachineLearningStudioInputsResponsePtrOutputWithContext(ctx context.Context) AzureMachineLearningStudioInputsResponsePtrOutput {
	return o
}

func (o AzureMachineLearningStudioInputsResponsePtrOutput) Elem() AzureMachineLearningStudioInputsResponseOutput {
	return o.ApplyT(func(v *AzureMachineLearningStudioInputsResponse) AzureMachineLearningStudioInputsResponse {
		if v != nil {
			return *v
		}
		var ret AzureMachineLearningStudioInputsResponse
		return ret
	}).(AzureMachineLearningStudioInputsResponseOutput)
}

func (o AzureMachineLearningStudioInputsResponsePtrOutput) ColumnNames() AzureMachineLearningStudioInputColumnResponseArrayOutput {
	return o.ApplyT(func(v *AzureMachineLearningStudioInputsResponse) []AzureMachineLearningStudioInputColumnResponse {
		if v == nil {
			return nil
		}
		return v.ColumnNames
	}).(AzureMachineLearningStudioInputColumnResponseArrayOutput)
}

func (o AzureMachineLearningStudioInputsResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMachineLearningStudioInputsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioOutputColumn struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningStudioOutputColumnInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioOutputColumnOutput() AzureMachineLearningStudioOutputColumnOutput
	ToAzureMachineLearningStudioOutputColumnOutputWithContext(context.Context) AzureMachineLearningStudioOutputColumnOutput
}

type AzureMachineLearningStudioOutputColumnArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningStudioOutputColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioOutputColumn)(nil)).Elem()
}

func (i AzureMachineLearningStudioOutputColumnArgs) ToAzureMachineLearningStudioOutputColumnOutput() AzureMachineLearningStudioOutputColumnOutput {
	return i.ToAzureMachineLearningStudioOutputColumnOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioOutputColumnArgs) ToAzureMachineLearningStudioOutputColumnOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioOutputColumnOutput)
}





type AzureMachineLearningStudioOutputColumnArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioOutputColumnArrayOutput() AzureMachineLearningStudioOutputColumnArrayOutput
	ToAzureMachineLearningStudioOutputColumnArrayOutputWithContext(context.Context) AzureMachineLearningStudioOutputColumnArrayOutput
}

type AzureMachineLearningStudioOutputColumnArray []AzureMachineLearningStudioOutputColumnInput

func (AzureMachineLearningStudioOutputColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioOutputColumn)(nil)).Elem()
}

func (i AzureMachineLearningStudioOutputColumnArray) ToAzureMachineLearningStudioOutputColumnArrayOutput() AzureMachineLearningStudioOutputColumnArrayOutput {
	return i.ToAzureMachineLearningStudioOutputColumnArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioOutputColumnArray) ToAzureMachineLearningStudioOutputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioOutputColumnArrayOutput)
}

type AzureMachineLearningStudioOutputColumnOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioOutputColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioOutputColumn)(nil)).Elem()
}

func (o AzureMachineLearningStudioOutputColumnOutput) ToAzureMachineLearningStudioOutputColumnOutput() AzureMachineLearningStudioOutputColumnOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnOutput) ToAzureMachineLearningStudioOutputColumnOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioOutputColumn) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioOutputColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioOutputColumn) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioOutputColumnArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioOutputColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioOutputColumn)(nil)).Elem()
}

func (o AzureMachineLearningStudioOutputColumnArrayOutput) ToAzureMachineLearningStudioOutputColumnArrayOutput() AzureMachineLearningStudioOutputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnArrayOutput) ToAzureMachineLearningStudioOutputColumnArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnArrayOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningStudioOutputColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningStudioOutputColumn {
		return vs[0].([]AzureMachineLearningStudioOutputColumn)[vs[1].(int)]
	}).(AzureMachineLearningStudioOutputColumnOutput)
}

type AzureMachineLearningStudioOutputColumnResponse struct {
	DataType *string `pulumi:"dataType"`
	Name     *string `pulumi:"name"`
}





type AzureMachineLearningStudioOutputColumnResponseInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioOutputColumnResponseOutput() AzureMachineLearningStudioOutputColumnResponseOutput
	ToAzureMachineLearningStudioOutputColumnResponseOutputWithContext(context.Context) AzureMachineLearningStudioOutputColumnResponseOutput
}

type AzureMachineLearningStudioOutputColumnResponseArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (AzureMachineLearningStudioOutputColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioOutputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningStudioOutputColumnResponseArgs) ToAzureMachineLearningStudioOutputColumnResponseOutput() AzureMachineLearningStudioOutputColumnResponseOutput {
	return i.ToAzureMachineLearningStudioOutputColumnResponseOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioOutputColumnResponseArgs) ToAzureMachineLearningStudioOutputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioOutputColumnResponseOutput)
}





type AzureMachineLearningStudioOutputColumnResponseArrayInput interface {
	pulumi.Input

	ToAzureMachineLearningStudioOutputColumnResponseArrayOutput() AzureMachineLearningStudioOutputColumnResponseArrayOutput
	ToAzureMachineLearningStudioOutputColumnResponseArrayOutputWithContext(context.Context) AzureMachineLearningStudioOutputColumnResponseArrayOutput
}

type AzureMachineLearningStudioOutputColumnResponseArray []AzureMachineLearningStudioOutputColumnResponseInput

func (AzureMachineLearningStudioOutputColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioOutputColumnResponse)(nil)).Elem()
}

func (i AzureMachineLearningStudioOutputColumnResponseArray) ToAzureMachineLearningStudioOutputColumnResponseArrayOutput() AzureMachineLearningStudioOutputColumnResponseArrayOutput {
	return i.ToAzureMachineLearningStudioOutputColumnResponseArrayOutputWithContext(context.Background())
}

func (i AzureMachineLearningStudioOutputColumnResponseArray) ToAzureMachineLearningStudioOutputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMachineLearningStudioOutputColumnResponseArrayOutput)
}

type AzureMachineLearningStudioOutputColumnResponseOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioOutputColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMachineLearningStudioOutputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioOutputColumnResponseOutput) ToAzureMachineLearningStudioOutputColumnResponseOutput() AzureMachineLearningStudioOutputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnResponseOutput) ToAzureMachineLearningStudioOutputColumnResponseOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnResponseOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioOutputColumnResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o AzureMachineLearningStudioOutputColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMachineLearningStudioOutputColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type AzureMachineLearningStudioOutputColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureMachineLearningStudioOutputColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureMachineLearningStudioOutputColumnResponse)(nil)).Elem()
}

func (o AzureMachineLearningStudioOutputColumnResponseArrayOutput) ToAzureMachineLearningStudioOutputColumnResponseArrayOutput() AzureMachineLearningStudioOutputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnResponseArrayOutput) ToAzureMachineLearningStudioOutputColumnResponseArrayOutputWithContext(ctx context.Context) AzureMachineLearningStudioOutputColumnResponseArrayOutput {
	return o
}

func (o AzureMachineLearningStudioOutputColumnResponseArrayOutput) Index(i pulumi.IntInput) AzureMachineLearningStudioOutputColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureMachineLearningStudioOutputColumnResponse {
		return vs[0].([]AzureMachineLearningStudioOutputColumnResponse)[vs[1].(int)]
	}).(AzureMachineLearningStudioOutputColumnResponseOutput)
}

type AzureSqlDatabaseOutputDataSource struct {
	AuthenticationMode *string  `pulumi:"authenticationMode"`
	Database           *string  `pulumi:"database"`
	MaxBatchCount      *float64 `pulumi:"maxBatchCount"`
	MaxWriterCount     *float64 `pulumi:"maxWriterCount"`
	Password           *string  `pulumi:"password"`
	Server             *string  `pulumi:"server"`
	Table              *string  `pulumi:"table"`
	Type               string   `pulumi:"type"`
	User               *string  `pulumi:"user"`
}





type AzureSqlDatabaseOutputDataSourceInput interface {
	pulumi.Input

	ToAzureSqlDatabaseOutputDataSourceOutput() AzureSqlDatabaseOutputDataSourceOutput
	ToAzureSqlDatabaseOutputDataSourceOutputWithContext(context.Context) AzureSqlDatabaseOutputDataSourceOutput
}

type AzureSqlDatabaseOutputDataSourceArgs struct {
	AuthenticationMode pulumi.StringPtrInput  `pulumi:"authenticationMode"`
	Database           pulumi.StringPtrInput  `pulumi:"database"`
	MaxBatchCount      pulumi.Float64PtrInput `pulumi:"maxBatchCount"`
	MaxWriterCount     pulumi.Float64PtrInput `pulumi:"maxWriterCount"`
	Password           pulumi.StringPtrInput  `pulumi:"password"`
	Server             pulumi.StringPtrInput  `pulumi:"server"`
	Table              pulumi.StringPtrInput  `pulumi:"table"`
	Type               pulumi.StringInput     `pulumi:"type"`
	User               pulumi.StringPtrInput  `pulumi:"user"`
}

func (AzureSqlDatabaseOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlDatabaseOutputDataSource)(nil)).Elem()
}

func (i AzureSqlDatabaseOutputDataSourceArgs) ToAzureSqlDatabaseOutputDataSourceOutput() AzureSqlDatabaseOutputDataSourceOutput {
	return i.ToAzureSqlDatabaseOutputDataSourceOutputWithContext(context.Background())
}

func (i AzureSqlDatabaseOutputDataSourceArgs) ToAzureSqlDatabaseOutputDataSourceOutputWithContext(ctx context.Context) AzureSqlDatabaseOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlDatabaseOutputDataSourceOutput)
}

type AzureSqlDatabaseOutputDataSourceOutput struct{ *pulumi.OutputState }

func (AzureSqlDatabaseOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlDatabaseOutputDataSource)(nil)).Elem()
}

func (o AzureSqlDatabaseOutputDataSourceOutput) ToAzureSqlDatabaseOutputDataSourceOutput() AzureSqlDatabaseOutputDataSourceOutput {
	return o
}

func (o AzureSqlDatabaseOutputDataSourceOutput) ToAzureSqlDatabaseOutputDataSourceOutputWithContext(ctx context.Context) AzureSqlDatabaseOutputDataSourceOutput {
	return o
}

func (o AzureSqlDatabaseOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) MaxBatchCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *float64 { return v.MaxBatchCount }).(pulumi.Float64PtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) MaxWriterCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *float64 { return v.MaxWriterCount }).(pulumi.Float64PtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

func (o AzureSqlDatabaseOutputDataSourceOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSource) *string { return v.User }).(pulumi.StringPtrOutput)
}

type AzureSqlDatabaseOutputDataSourceResponse struct {
	AuthenticationMode *string  `pulumi:"authenticationMode"`
	Database           *string  `pulumi:"database"`
	MaxBatchCount      *float64 `pulumi:"maxBatchCount"`
	MaxWriterCount     *float64 `pulumi:"maxWriterCount"`
	Password           *string  `pulumi:"password"`
	Server             *string  `pulumi:"server"`
	Table              *string  `pulumi:"table"`
	Type               string   `pulumi:"type"`
	User               *string  `pulumi:"user"`
}





type AzureSqlDatabaseOutputDataSourceResponseInput interface {
	pulumi.Input

	ToAzureSqlDatabaseOutputDataSourceResponseOutput() AzureSqlDatabaseOutputDataSourceResponseOutput
	ToAzureSqlDatabaseOutputDataSourceResponseOutputWithContext(context.Context) AzureSqlDatabaseOutputDataSourceResponseOutput
}

type AzureSqlDatabaseOutputDataSourceResponseArgs struct {
	AuthenticationMode pulumi.StringPtrInput  `pulumi:"authenticationMode"`
	Database           pulumi.StringPtrInput  `pulumi:"database"`
	MaxBatchCount      pulumi.Float64PtrInput `pulumi:"maxBatchCount"`
	MaxWriterCount     pulumi.Float64PtrInput `pulumi:"maxWriterCount"`
	Password           pulumi.StringPtrInput  `pulumi:"password"`
	Server             pulumi.StringPtrInput  `pulumi:"server"`
	Table              pulumi.StringPtrInput  `pulumi:"table"`
	Type               pulumi.StringInput     `pulumi:"type"`
	User               pulumi.StringPtrInput  `pulumi:"user"`
}

func (AzureSqlDatabaseOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlDatabaseOutputDataSourceResponse)(nil)).Elem()
}

func (i AzureSqlDatabaseOutputDataSourceResponseArgs) ToAzureSqlDatabaseOutputDataSourceResponseOutput() AzureSqlDatabaseOutputDataSourceResponseOutput {
	return i.ToAzureSqlDatabaseOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i AzureSqlDatabaseOutputDataSourceResponseArgs) ToAzureSqlDatabaseOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureSqlDatabaseOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlDatabaseOutputDataSourceResponseOutput)
}

type AzureSqlDatabaseOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlDatabaseOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlDatabaseOutputDataSourceResponse)(nil)).Elem()
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) ToAzureSqlDatabaseOutputDataSourceResponseOutput() AzureSqlDatabaseOutputDataSourceResponseOutput {
	return o
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) ToAzureSqlDatabaseOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureSqlDatabaseOutputDataSourceResponseOutput {
	return o
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) MaxBatchCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *float64 { return v.MaxBatchCount }).(pulumi.Float64PtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) MaxWriterCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *float64 { return v.MaxWriterCount }).(pulumi.Float64PtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o AzureSqlDatabaseOutputDataSourceResponseOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlDatabaseOutputDataSourceResponse) *string { return v.User }).(pulumi.StringPtrOutput)
}

type AzureSqlReferenceInputDataSource struct {
	Database           *string `pulumi:"database"`
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  *string `pulumi:"fullSnapshotQuery"`
	Password           *string `pulumi:"password"`
	RefreshRate        *string `pulumi:"refreshRate"`
	RefreshType        *string `pulumi:"refreshType"`
	Server             *string `pulumi:"server"`
	Table              *string `pulumi:"table"`
	Type               string  `pulumi:"type"`
	User               *string `pulumi:"user"`
}





type AzureSqlReferenceInputDataSourceInput interface {
	pulumi.Input

	ToAzureSqlReferenceInputDataSourceOutput() AzureSqlReferenceInputDataSourceOutput
	ToAzureSqlReferenceInputDataSourceOutputWithContext(context.Context) AzureSqlReferenceInputDataSourceOutput
}

type AzureSqlReferenceInputDataSourceArgs struct {
	Database           pulumi.StringPtrInput `pulumi:"database"`
	DeltaSnapshotQuery pulumi.StringPtrInput `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  pulumi.StringPtrInput `pulumi:"fullSnapshotQuery"`
	Password           pulumi.StringPtrInput `pulumi:"password"`
	RefreshRate        pulumi.StringPtrInput `pulumi:"refreshRate"`
	RefreshType        pulumi.StringPtrInput `pulumi:"refreshType"`
	Server             pulumi.StringPtrInput `pulumi:"server"`
	Table              pulumi.StringPtrInput `pulumi:"table"`
	Type               pulumi.StringInput    `pulumi:"type"`
	User               pulumi.StringPtrInput `pulumi:"user"`
}

func (AzureSqlReferenceInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlReferenceInputDataSource)(nil)).Elem()
}

func (i AzureSqlReferenceInputDataSourceArgs) ToAzureSqlReferenceInputDataSourceOutput() AzureSqlReferenceInputDataSourceOutput {
	return i.ToAzureSqlReferenceInputDataSourceOutputWithContext(context.Background())
}

func (i AzureSqlReferenceInputDataSourceArgs) ToAzureSqlReferenceInputDataSourceOutputWithContext(ctx context.Context) AzureSqlReferenceInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlReferenceInputDataSourceOutput)
}

type AzureSqlReferenceInputDataSourceOutput struct{ *pulumi.OutputState }

func (AzureSqlReferenceInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlReferenceInputDataSource)(nil)).Elem()
}

func (o AzureSqlReferenceInputDataSourceOutput) ToAzureSqlReferenceInputDataSourceOutput() AzureSqlReferenceInputDataSourceOutput {
	return o
}

func (o AzureSqlReferenceInputDataSourceOutput) ToAzureSqlReferenceInputDataSourceOutputWithContext(ctx context.Context) AzureSqlReferenceInputDataSourceOutput {
	return o
}

func (o AzureSqlReferenceInputDataSourceOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) DeltaSnapshotQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.DeltaSnapshotQuery }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) FullSnapshotQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.FullSnapshotQuery }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) RefreshRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.RefreshRate }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) RefreshType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.RefreshType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

func (o AzureSqlReferenceInputDataSourceOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSource) *string { return v.User }).(pulumi.StringPtrOutput)
}

type AzureSqlReferenceInputDataSourceResponse struct {
	Database           *string `pulumi:"database"`
	DeltaSnapshotQuery *string `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  *string `pulumi:"fullSnapshotQuery"`
	Password           *string `pulumi:"password"`
	RefreshRate        *string `pulumi:"refreshRate"`
	RefreshType        *string `pulumi:"refreshType"`
	Server             *string `pulumi:"server"`
	Table              *string `pulumi:"table"`
	Type               string  `pulumi:"type"`
	User               *string `pulumi:"user"`
}





type AzureSqlReferenceInputDataSourceResponseInput interface {
	pulumi.Input

	ToAzureSqlReferenceInputDataSourceResponseOutput() AzureSqlReferenceInputDataSourceResponseOutput
	ToAzureSqlReferenceInputDataSourceResponseOutputWithContext(context.Context) AzureSqlReferenceInputDataSourceResponseOutput
}

type AzureSqlReferenceInputDataSourceResponseArgs struct {
	Database           pulumi.StringPtrInput `pulumi:"database"`
	DeltaSnapshotQuery pulumi.StringPtrInput `pulumi:"deltaSnapshotQuery"`
	FullSnapshotQuery  pulumi.StringPtrInput `pulumi:"fullSnapshotQuery"`
	Password           pulumi.StringPtrInput `pulumi:"password"`
	RefreshRate        pulumi.StringPtrInput `pulumi:"refreshRate"`
	RefreshType        pulumi.StringPtrInput `pulumi:"refreshType"`
	Server             pulumi.StringPtrInput `pulumi:"server"`
	Table              pulumi.StringPtrInput `pulumi:"table"`
	Type               pulumi.StringInput    `pulumi:"type"`
	User               pulumi.StringPtrInput `pulumi:"user"`
}

func (AzureSqlReferenceInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlReferenceInputDataSourceResponse)(nil)).Elem()
}

func (i AzureSqlReferenceInputDataSourceResponseArgs) ToAzureSqlReferenceInputDataSourceResponseOutput() AzureSqlReferenceInputDataSourceResponseOutput {
	return i.ToAzureSqlReferenceInputDataSourceResponseOutputWithContext(context.Background())
}

func (i AzureSqlReferenceInputDataSourceResponseArgs) ToAzureSqlReferenceInputDataSourceResponseOutputWithContext(ctx context.Context) AzureSqlReferenceInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlReferenceInputDataSourceResponseOutput)
}

type AzureSqlReferenceInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlReferenceInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlReferenceInputDataSourceResponse)(nil)).Elem()
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) ToAzureSqlReferenceInputDataSourceResponseOutput() AzureSqlReferenceInputDataSourceResponseOutput {
	return o
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) ToAzureSqlReferenceInputDataSourceResponseOutputWithContext(ctx context.Context) AzureSqlReferenceInputDataSourceResponseOutput {
	return o
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) DeltaSnapshotQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.DeltaSnapshotQuery }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) FullSnapshotQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.FullSnapshotQuery }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) RefreshRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.RefreshRate }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) RefreshType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.RefreshType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o AzureSqlReferenceInputDataSourceResponseOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlReferenceInputDataSourceResponse) *string { return v.User }).(pulumi.StringPtrOutput)
}

type AzureSynapseOutputDataSource struct {
	Database *string `pulumi:"database"`
	Password *string `pulumi:"password"`
	Server   *string `pulumi:"server"`
	Table    *string `pulumi:"table"`
	Type     string  `pulumi:"type"`
	User     *string `pulumi:"user"`
}





type AzureSynapseOutputDataSourceInput interface {
	pulumi.Input

	ToAzureSynapseOutputDataSourceOutput() AzureSynapseOutputDataSourceOutput
	ToAzureSynapseOutputDataSourceOutputWithContext(context.Context) AzureSynapseOutputDataSourceOutput
}

type AzureSynapseOutputDataSourceArgs struct {
	Database pulumi.StringPtrInput `pulumi:"database"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Server   pulumi.StringPtrInput `pulumi:"server"`
	Table    pulumi.StringPtrInput `pulumi:"table"`
	Type     pulumi.StringInput    `pulumi:"type"`
	User     pulumi.StringPtrInput `pulumi:"user"`
}

func (AzureSynapseOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSynapseOutputDataSource)(nil)).Elem()
}

func (i AzureSynapseOutputDataSourceArgs) ToAzureSynapseOutputDataSourceOutput() AzureSynapseOutputDataSourceOutput {
	return i.ToAzureSynapseOutputDataSourceOutputWithContext(context.Background())
}

func (i AzureSynapseOutputDataSourceArgs) ToAzureSynapseOutputDataSourceOutputWithContext(ctx context.Context) AzureSynapseOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSynapseOutputDataSourceOutput)
}

type AzureSynapseOutputDataSourceOutput struct{ *pulumi.OutputState }

func (AzureSynapseOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSynapseOutputDataSource)(nil)).Elem()
}

func (o AzureSynapseOutputDataSourceOutput) ToAzureSynapseOutputDataSourceOutput() AzureSynapseOutputDataSourceOutput {
	return o
}

func (o AzureSynapseOutputDataSourceOutput) ToAzureSynapseOutputDataSourceOutputWithContext(ctx context.Context) AzureSynapseOutputDataSourceOutput {
	return o
}

func (o AzureSynapseOutputDataSourceOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSource) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSource) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSource) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSource) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

func (o AzureSynapseOutputDataSourceOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSource) *string { return v.User }).(pulumi.StringPtrOutput)
}

type AzureSynapseOutputDataSourceResponse struct {
	Database *string `pulumi:"database"`
	Password *string `pulumi:"password"`
	Server   *string `pulumi:"server"`
	Table    *string `pulumi:"table"`
	Type     string  `pulumi:"type"`
	User     *string `pulumi:"user"`
}





type AzureSynapseOutputDataSourceResponseInput interface {
	pulumi.Input

	ToAzureSynapseOutputDataSourceResponseOutput() AzureSynapseOutputDataSourceResponseOutput
	ToAzureSynapseOutputDataSourceResponseOutputWithContext(context.Context) AzureSynapseOutputDataSourceResponseOutput
}

type AzureSynapseOutputDataSourceResponseArgs struct {
	Database pulumi.StringPtrInput `pulumi:"database"`
	Password pulumi.StringPtrInput `pulumi:"password"`
	Server   pulumi.StringPtrInput `pulumi:"server"`
	Table    pulumi.StringPtrInput `pulumi:"table"`
	Type     pulumi.StringInput    `pulumi:"type"`
	User     pulumi.StringPtrInput `pulumi:"user"`
}

func (AzureSynapseOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSynapseOutputDataSourceResponse)(nil)).Elem()
}

func (i AzureSynapseOutputDataSourceResponseArgs) ToAzureSynapseOutputDataSourceResponseOutput() AzureSynapseOutputDataSourceResponseOutput {
	return i.ToAzureSynapseOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i AzureSynapseOutputDataSourceResponseArgs) ToAzureSynapseOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureSynapseOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSynapseOutputDataSourceResponseOutput)
}

type AzureSynapseOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (AzureSynapseOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSynapseOutputDataSourceResponse)(nil)).Elem()
}

func (o AzureSynapseOutputDataSourceResponseOutput) ToAzureSynapseOutputDataSourceResponseOutput() AzureSynapseOutputDataSourceResponseOutput {
	return o
}

func (o AzureSynapseOutputDataSourceResponseOutput) ToAzureSynapseOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureSynapseOutputDataSourceResponseOutput {
	return o
}

func (o AzureSynapseOutputDataSourceResponseOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSourceResponse) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSourceResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceResponseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSourceResponse) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSourceResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureSynapseOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o AzureSynapseOutputDataSourceResponseOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSynapseOutputDataSourceResponse) *string { return v.User }).(pulumi.StringPtrOutput)
}

type AzureTableOutputDataSource struct {
	AccountKey      *string  `pulumi:"accountKey"`
	AccountName     *string  `pulumi:"accountName"`
	BatchSize       *int     `pulumi:"batchSize"`
	ColumnsToRemove []string `pulumi:"columnsToRemove"`
	PartitionKey    *string  `pulumi:"partitionKey"`
	RowKey          *string  `pulumi:"rowKey"`
	Table           *string  `pulumi:"table"`
	Type            string   `pulumi:"type"`
}





type AzureTableOutputDataSourceInput interface {
	pulumi.Input

	ToAzureTableOutputDataSourceOutput() AzureTableOutputDataSourceOutput
	ToAzureTableOutputDataSourceOutputWithContext(context.Context) AzureTableOutputDataSourceOutput
}

type AzureTableOutputDataSourceArgs struct {
	AccountKey      pulumi.StringPtrInput   `pulumi:"accountKey"`
	AccountName     pulumi.StringPtrInput   `pulumi:"accountName"`
	BatchSize       pulumi.IntPtrInput      `pulumi:"batchSize"`
	ColumnsToRemove pulumi.StringArrayInput `pulumi:"columnsToRemove"`
	PartitionKey    pulumi.StringPtrInput   `pulumi:"partitionKey"`
	RowKey          pulumi.StringPtrInput   `pulumi:"rowKey"`
	Table           pulumi.StringPtrInput   `pulumi:"table"`
	Type            pulumi.StringInput      `pulumi:"type"`
}

func (AzureTableOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableOutputDataSource)(nil)).Elem()
}

func (i AzureTableOutputDataSourceArgs) ToAzureTableOutputDataSourceOutput() AzureTableOutputDataSourceOutput {
	return i.ToAzureTableOutputDataSourceOutputWithContext(context.Background())
}

func (i AzureTableOutputDataSourceArgs) ToAzureTableOutputDataSourceOutputWithContext(ctx context.Context) AzureTableOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableOutputDataSourceOutput)
}

type AzureTableOutputDataSourceOutput struct{ *pulumi.OutputState }

func (AzureTableOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableOutputDataSource)(nil)).Elem()
}

func (o AzureTableOutputDataSourceOutput) ToAzureTableOutputDataSourceOutput() AzureTableOutputDataSourceOutput {
	return o
}

func (o AzureTableOutputDataSourceOutput) ToAzureTableOutputDataSourceOutputWithContext(ctx context.Context) AzureTableOutputDataSourceOutput {
	return o
}

func (o AzureTableOutputDataSourceOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o AzureTableOutputDataSourceOutput) ColumnsToRemove() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) []string { return v.ColumnsToRemove }).(pulumi.StringArrayOutput)
}

func (o AzureTableOutputDataSourceOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceOutput) RowKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) *string { return v.RowKey }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureTableOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type AzureTableOutputDataSourceResponse struct {
	AccountKey      *string  `pulumi:"accountKey"`
	AccountName     *string  `pulumi:"accountName"`
	BatchSize       *int     `pulumi:"batchSize"`
	ColumnsToRemove []string `pulumi:"columnsToRemove"`
	PartitionKey    *string  `pulumi:"partitionKey"`
	RowKey          *string  `pulumi:"rowKey"`
	Table           *string  `pulumi:"table"`
	Type            string   `pulumi:"type"`
}





type AzureTableOutputDataSourceResponseInput interface {
	pulumi.Input

	ToAzureTableOutputDataSourceResponseOutput() AzureTableOutputDataSourceResponseOutput
	ToAzureTableOutputDataSourceResponseOutputWithContext(context.Context) AzureTableOutputDataSourceResponseOutput
}

type AzureTableOutputDataSourceResponseArgs struct {
	AccountKey      pulumi.StringPtrInput   `pulumi:"accountKey"`
	AccountName     pulumi.StringPtrInput   `pulumi:"accountName"`
	BatchSize       pulumi.IntPtrInput      `pulumi:"batchSize"`
	ColumnsToRemove pulumi.StringArrayInput `pulumi:"columnsToRemove"`
	PartitionKey    pulumi.StringPtrInput   `pulumi:"partitionKey"`
	RowKey          pulumi.StringPtrInput   `pulumi:"rowKey"`
	Table           pulumi.StringPtrInput   `pulumi:"table"`
	Type            pulumi.StringInput      `pulumi:"type"`
}

func (AzureTableOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableOutputDataSourceResponse)(nil)).Elem()
}

func (i AzureTableOutputDataSourceResponseArgs) ToAzureTableOutputDataSourceResponseOutput() AzureTableOutputDataSourceResponseOutput {
	return i.ToAzureTableOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i AzureTableOutputDataSourceResponseArgs) ToAzureTableOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureTableOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableOutputDataSourceResponseOutput)
}

type AzureTableOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (AzureTableOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableOutputDataSourceResponse)(nil)).Elem()
}

func (o AzureTableOutputDataSourceResponseOutput) ToAzureTableOutputDataSourceResponseOutput() AzureTableOutputDataSourceResponseOutput {
	return o
}

func (o AzureTableOutputDataSourceResponseOutput) ToAzureTableOutputDataSourceResponseOutputWithContext(ctx context.Context) AzureTableOutputDataSourceResponseOutput {
	return o
}

func (o AzureTableOutputDataSourceResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) ColumnsToRemove() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) []string { return v.ColumnsToRemove }).(pulumi.StringArrayOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) RowKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) *string { return v.RowKey }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o AzureTableOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AzureTableOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type BlobOutputDataSource struct {
	AuthenticationMode *string          `pulumi:"authenticationMode"`
	Container          *string          `pulumi:"container"`
	DateFormat         *string          `pulumi:"dateFormat"`
	PathPattern        *string          `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat         *string          `pulumi:"timeFormat"`
	Type               string           `pulumi:"type"`
}





type BlobOutputDataSourceInput interface {
	pulumi.Input

	ToBlobOutputDataSourceOutput() BlobOutputDataSourceOutput
	ToBlobOutputDataSourceOutputWithContext(context.Context) BlobOutputDataSourceOutput
}

type BlobOutputDataSourceArgs struct {
	AuthenticationMode pulumi.StringPtrInput    `pulumi:"authenticationMode"`
	Container          pulumi.StringPtrInput    `pulumi:"container"`
	DateFormat         pulumi.StringPtrInput    `pulumi:"dateFormat"`
	PathPattern        pulumi.StringPtrInput    `pulumi:"pathPattern"`
	StorageAccounts    StorageAccountArrayInput `pulumi:"storageAccounts"`
	TimeFormat         pulumi.StringPtrInput    `pulumi:"timeFormat"`
	Type               pulumi.StringInput       `pulumi:"type"`
}

func (BlobOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobOutputDataSource)(nil)).Elem()
}

func (i BlobOutputDataSourceArgs) ToBlobOutputDataSourceOutput() BlobOutputDataSourceOutput {
	return i.ToBlobOutputDataSourceOutputWithContext(context.Background())
}

func (i BlobOutputDataSourceArgs) ToBlobOutputDataSourceOutputWithContext(ctx context.Context) BlobOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobOutputDataSourceOutput)
}

type BlobOutputDataSourceOutput struct{ *pulumi.OutputState }

func (BlobOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobOutputDataSource)(nil)).Elem()
}

func (o BlobOutputDataSourceOutput) ToBlobOutputDataSourceOutput() BlobOutputDataSourceOutput {
	return o
}

func (o BlobOutputDataSourceOutput) ToBlobOutputDataSourceOutputWithContext(ctx context.Context) BlobOutputDataSourceOutput {
	return o
}

func (o BlobOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSource) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSource) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceOutput) PathPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSource) *string { return v.PathPattern }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceOutput) StorageAccounts() StorageAccountArrayOutput {
	return o.ApplyT(func(v BlobOutputDataSource) []StorageAccount { return v.StorageAccounts }).(StorageAccountArrayOutput)
}

func (o BlobOutputDataSourceOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSource) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type BlobOutputDataSourceResponse struct {
	AuthenticationMode *string                  `pulumi:"authenticationMode"`
	Container          *string                  `pulumi:"container"`
	DateFormat         *string                  `pulumi:"dateFormat"`
	PathPattern        *string                  `pulumi:"pathPattern"`
	StorageAccounts    []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat         *string                  `pulumi:"timeFormat"`
	Type               string                   `pulumi:"type"`
}





type BlobOutputDataSourceResponseInput interface {
	pulumi.Input

	ToBlobOutputDataSourceResponseOutput() BlobOutputDataSourceResponseOutput
	ToBlobOutputDataSourceResponseOutputWithContext(context.Context) BlobOutputDataSourceResponseOutput
}

type BlobOutputDataSourceResponseArgs struct {
	AuthenticationMode pulumi.StringPtrInput            `pulumi:"authenticationMode"`
	Container          pulumi.StringPtrInput            `pulumi:"container"`
	DateFormat         pulumi.StringPtrInput            `pulumi:"dateFormat"`
	PathPattern        pulumi.StringPtrInput            `pulumi:"pathPattern"`
	StorageAccounts    StorageAccountResponseArrayInput `pulumi:"storageAccounts"`
	TimeFormat         pulumi.StringPtrInput            `pulumi:"timeFormat"`
	Type               pulumi.StringInput               `pulumi:"type"`
}

func (BlobOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobOutputDataSourceResponse)(nil)).Elem()
}

func (i BlobOutputDataSourceResponseArgs) ToBlobOutputDataSourceResponseOutput() BlobOutputDataSourceResponseOutput {
	return i.ToBlobOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i BlobOutputDataSourceResponseArgs) ToBlobOutputDataSourceResponseOutputWithContext(ctx context.Context) BlobOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobOutputDataSourceResponseOutput)
}

type BlobOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (BlobOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobOutputDataSourceResponse)(nil)).Elem()
}

func (o BlobOutputDataSourceResponseOutput) ToBlobOutputDataSourceResponseOutput() BlobOutputDataSourceResponseOutput {
	return o
}

func (o BlobOutputDataSourceResponseOutput) ToBlobOutputDataSourceResponseOutputWithContext(ctx context.Context) BlobOutputDataSourceResponseOutput {
	return o
}

func (o BlobOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceResponseOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceResponseOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceResponseOutput) PathPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) *string { return v.PathPattern }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceResponseOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) []StorageAccountResponse { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o BlobOutputDataSourceResponseOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o BlobOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type BlobReferenceInputDataSource struct {
	Container       *string          `pulumi:"container"`
	DateFormat      *string          `pulumi:"dateFormat"`
	PathPattern     *string          `pulumi:"pathPattern"`
	StorageAccounts []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat      *string          `pulumi:"timeFormat"`
	Type            string           `pulumi:"type"`
}





type BlobReferenceInputDataSourceInput interface {
	pulumi.Input

	ToBlobReferenceInputDataSourceOutput() BlobReferenceInputDataSourceOutput
	ToBlobReferenceInputDataSourceOutputWithContext(context.Context) BlobReferenceInputDataSourceOutput
}

type BlobReferenceInputDataSourceArgs struct {
	Container       pulumi.StringPtrInput    `pulumi:"container"`
	DateFormat      pulumi.StringPtrInput    `pulumi:"dateFormat"`
	PathPattern     pulumi.StringPtrInput    `pulumi:"pathPattern"`
	StorageAccounts StorageAccountArrayInput `pulumi:"storageAccounts"`
	TimeFormat      pulumi.StringPtrInput    `pulumi:"timeFormat"`
	Type            pulumi.StringInput       `pulumi:"type"`
}

func (BlobReferenceInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobReferenceInputDataSource)(nil)).Elem()
}

func (i BlobReferenceInputDataSourceArgs) ToBlobReferenceInputDataSourceOutput() BlobReferenceInputDataSourceOutput {
	return i.ToBlobReferenceInputDataSourceOutputWithContext(context.Background())
}

func (i BlobReferenceInputDataSourceArgs) ToBlobReferenceInputDataSourceOutputWithContext(ctx context.Context) BlobReferenceInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobReferenceInputDataSourceOutput)
}

type BlobReferenceInputDataSourceOutput struct{ *pulumi.OutputState }

func (BlobReferenceInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobReferenceInputDataSource)(nil)).Elem()
}

func (o BlobReferenceInputDataSourceOutput) ToBlobReferenceInputDataSourceOutput() BlobReferenceInputDataSourceOutput {
	return o
}

func (o BlobReferenceInputDataSourceOutput) ToBlobReferenceInputDataSourceOutputWithContext(ctx context.Context) BlobReferenceInputDataSourceOutput {
	return o
}

func (o BlobReferenceInputDataSourceOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSource) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSource) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceOutput) PathPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSource) *string { return v.PathPattern }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceOutput) StorageAccounts() StorageAccountArrayOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSource) []StorageAccount { return v.StorageAccounts }).(StorageAccountArrayOutput)
}

func (o BlobReferenceInputDataSourceOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSource) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type BlobReferenceInputDataSourceResponse struct {
	Container       *string                  `pulumi:"container"`
	DateFormat      *string                  `pulumi:"dateFormat"`
	PathPattern     *string                  `pulumi:"pathPattern"`
	StorageAccounts []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat      *string                  `pulumi:"timeFormat"`
	Type            string                   `pulumi:"type"`
}





type BlobReferenceInputDataSourceResponseInput interface {
	pulumi.Input

	ToBlobReferenceInputDataSourceResponseOutput() BlobReferenceInputDataSourceResponseOutput
	ToBlobReferenceInputDataSourceResponseOutputWithContext(context.Context) BlobReferenceInputDataSourceResponseOutput
}

type BlobReferenceInputDataSourceResponseArgs struct {
	Container       pulumi.StringPtrInput            `pulumi:"container"`
	DateFormat      pulumi.StringPtrInput            `pulumi:"dateFormat"`
	PathPattern     pulumi.StringPtrInput            `pulumi:"pathPattern"`
	StorageAccounts StorageAccountResponseArrayInput `pulumi:"storageAccounts"`
	TimeFormat      pulumi.StringPtrInput            `pulumi:"timeFormat"`
	Type            pulumi.StringInput               `pulumi:"type"`
}

func (BlobReferenceInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobReferenceInputDataSourceResponse)(nil)).Elem()
}

func (i BlobReferenceInputDataSourceResponseArgs) ToBlobReferenceInputDataSourceResponseOutput() BlobReferenceInputDataSourceResponseOutput {
	return i.ToBlobReferenceInputDataSourceResponseOutputWithContext(context.Background())
}

func (i BlobReferenceInputDataSourceResponseArgs) ToBlobReferenceInputDataSourceResponseOutputWithContext(ctx context.Context) BlobReferenceInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobReferenceInputDataSourceResponseOutput)
}

type BlobReferenceInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (BlobReferenceInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobReferenceInputDataSourceResponse)(nil)).Elem()
}

func (o BlobReferenceInputDataSourceResponseOutput) ToBlobReferenceInputDataSourceResponseOutput() BlobReferenceInputDataSourceResponseOutput {
	return o
}

func (o BlobReferenceInputDataSourceResponseOutput) ToBlobReferenceInputDataSourceResponseOutputWithContext(ctx context.Context) BlobReferenceInputDataSourceResponseOutput {
	return o
}

func (o BlobReferenceInputDataSourceResponseOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSourceResponse) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceResponseOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSourceResponse) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceResponseOutput) PathPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSourceResponse) *string { return v.PathPattern }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceResponseOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSourceResponse) []StorageAccountResponse { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o BlobReferenceInputDataSourceResponseOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSourceResponse) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o BlobReferenceInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobReferenceInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type BlobStreamInputDataSource struct {
	Container            *string          `pulumi:"container"`
	DateFormat           *string          `pulumi:"dateFormat"`
	PathPattern          *string          `pulumi:"pathPattern"`
	SourcePartitionCount *int             `pulumi:"sourcePartitionCount"`
	StorageAccounts      []StorageAccount `pulumi:"storageAccounts"`
	TimeFormat           *string          `pulumi:"timeFormat"`
	Type                 string           `pulumi:"type"`
}





type BlobStreamInputDataSourceInput interface {
	pulumi.Input

	ToBlobStreamInputDataSourceOutput() BlobStreamInputDataSourceOutput
	ToBlobStreamInputDataSourceOutputWithContext(context.Context) BlobStreamInputDataSourceOutput
}

type BlobStreamInputDataSourceArgs struct {
	Container            pulumi.StringPtrInput    `pulumi:"container"`
	DateFormat           pulumi.StringPtrInput    `pulumi:"dateFormat"`
	PathPattern          pulumi.StringPtrInput    `pulumi:"pathPattern"`
	SourcePartitionCount pulumi.IntPtrInput       `pulumi:"sourcePartitionCount"`
	StorageAccounts      StorageAccountArrayInput `pulumi:"storageAccounts"`
	TimeFormat           pulumi.StringPtrInput    `pulumi:"timeFormat"`
	Type                 pulumi.StringInput       `pulumi:"type"`
}

func (BlobStreamInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStreamInputDataSource)(nil)).Elem()
}

func (i BlobStreamInputDataSourceArgs) ToBlobStreamInputDataSourceOutput() BlobStreamInputDataSourceOutput {
	return i.ToBlobStreamInputDataSourceOutputWithContext(context.Background())
}

func (i BlobStreamInputDataSourceArgs) ToBlobStreamInputDataSourceOutputWithContext(ctx context.Context) BlobStreamInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStreamInputDataSourceOutput)
}

type BlobStreamInputDataSourceOutput struct{ *pulumi.OutputState }

func (BlobStreamInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStreamInputDataSource)(nil)).Elem()
}

func (o BlobStreamInputDataSourceOutput) ToBlobStreamInputDataSourceOutput() BlobStreamInputDataSourceOutput {
	return o
}

func (o BlobStreamInputDataSourceOutput) ToBlobStreamInputDataSourceOutputWithContext(ctx context.Context) BlobStreamInputDataSourceOutput {
	return o
}

func (o BlobStreamInputDataSourceOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceOutput) PathPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) *string { return v.PathPattern }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceOutput) SourcePartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) *int { return v.SourcePartitionCount }).(pulumi.IntPtrOutput)
}

func (o BlobStreamInputDataSourceOutput) StorageAccounts() StorageAccountArrayOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) []StorageAccount { return v.StorageAccounts }).(StorageAccountArrayOutput)
}

func (o BlobStreamInputDataSourceOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobStreamInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type BlobStreamInputDataSourceResponse struct {
	Container            *string                  `pulumi:"container"`
	DateFormat           *string                  `pulumi:"dateFormat"`
	PathPattern          *string                  `pulumi:"pathPattern"`
	SourcePartitionCount *int                     `pulumi:"sourcePartitionCount"`
	StorageAccounts      []StorageAccountResponse `pulumi:"storageAccounts"`
	TimeFormat           *string                  `pulumi:"timeFormat"`
	Type                 string                   `pulumi:"type"`
}





type BlobStreamInputDataSourceResponseInput interface {
	pulumi.Input

	ToBlobStreamInputDataSourceResponseOutput() BlobStreamInputDataSourceResponseOutput
	ToBlobStreamInputDataSourceResponseOutputWithContext(context.Context) BlobStreamInputDataSourceResponseOutput
}

type BlobStreamInputDataSourceResponseArgs struct {
	Container            pulumi.StringPtrInput            `pulumi:"container"`
	DateFormat           pulumi.StringPtrInput            `pulumi:"dateFormat"`
	PathPattern          pulumi.StringPtrInput            `pulumi:"pathPattern"`
	SourcePartitionCount pulumi.IntPtrInput               `pulumi:"sourcePartitionCount"`
	StorageAccounts      StorageAccountResponseArrayInput `pulumi:"storageAccounts"`
	TimeFormat           pulumi.StringPtrInput            `pulumi:"timeFormat"`
	Type                 pulumi.StringInput               `pulumi:"type"`
}

func (BlobStreamInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStreamInputDataSourceResponse)(nil)).Elem()
}

func (i BlobStreamInputDataSourceResponseArgs) ToBlobStreamInputDataSourceResponseOutput() BlobStreamInputDataSourceResponseOutput {
	return i.ToBlobStreamInputDataSourceResponseOutputWithContext(context.Background())
}

func (i BlobStreamInputDataSourceResponseArgs) ToBlobStreamInputDataSourceResponseOutputWithContext(ctx context.Context) BlobStreamInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStreamInputDataSourceResponseOutput)
}

type BlobStreamInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (BlobStreamInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStreamInputDataSourceResponse)(nil)).Elem()
}

func (o BlobStreamInputDataSourceResponseOutput) ToBlobStreamInputDataSourceResponseOutput() BlobStreamInputDataSourceResponseOutput {
	return o
}

func (o BlobStreamInputDataSourceResponseOutput) ToBlobStreamInputDataSourceResponseOutputWithContext(ctx context.Context) BlobStreamInputDataSourceResponseOutput {
	return o
}

func (o BlobStreamInputDataSourceResponseOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceResponseOutput) DateFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) *string { return v.DateFormat }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceResponseOutput) PathPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) *string { return v.PathPattern }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceResponseOutput) SourcePartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) *int { return v.SourcePartitionCount }).(pulumi.IntPtrOutput)
}

func (o BlobStreamInputDataSourceResponseOutput) StorageAccounts() StorageAccountResponseArrayOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) []StorageAccountResponse { return v.StorageAccounts }).(StorageAccountResponseArrayOutput)
}

func (o BlobStreamInputDataSourceResponseOutput) TimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) *string { return v.TimeFormat }).(pulumi.StringPtrOutput)
}

func (o BlobStreamInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BlobStreamInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type CSharpFunctionBinding struct {
	Class   *string `pulumi:"class"`
	DllPath *string `pulumi:"dllPath"`
	Method  *string `pulumi:"method"`
	Script  *string `pulumi:"script"`
	Type    string  `pulumi:"type"`
}





type CSharpFunctionBindingInput interface {
	pulumi.Input

	ToCSharpFunctionBindingOutput() CSharpFunctionBindingOutput
	ToCSharpFunctionBindingOutputWithContext(context.Context) CSharpFunctionBindingOutput
}

type CSharpFunctionBindingArgs struct {
	Class   pulumi.StringPtrInput `pulumi:"class"`
	DllPath pulumi.StringPtrInput `pulumi:"dllPath"`
	Method  pulumi.StringPtrInput `pulumi:"method"`
	Script  pulumi.StringPtrInput `pulumi:"script"`
	Type    pulumi.StringInput    `pulumi:"type"`
}

func (CSharpFunctionBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CSharpFunctionBinding)(nil)).Elem()
}

func (i CSharpFunctionBindingArgs) ToCSharpFunctionBindingOutput() CSharpFunctionBindingOutput {
	return i.ToCSharpFunctionBindingOutputWithContext(context.Background())
}

func (i CSharpFunctionBindingArgs) ToCSharpFunctionBindingOutputWithContext(ctx context.Context) CSharpFunctionBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSharpFunctionBindingOutput)
}

type CSharpFunctionBindingOutput struct{ *pulumi.OutputState }

func (CSharpFunctionBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CSharpFunctionBinding)(nil)).Elem()
}

func (o CSharpFunctionBindingOutput) ToCSharpFunctionBindingOutput() CSharpFunctionBindingOutput {
	return o
}

func (o CSharpFunctionBindingOutput) ToCSharpFunctionBindingOutputWithContext(ctx context.Context) CSharpFunctionBindingOutput {
	return o
}

func (o CSharpFunctionBindingOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBinding) *string { return v.Class }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingOutput) DllPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBinding) *string { return v.DllPath }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBinding) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBinding) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CSharpFunctionBinding) string { return v.Type }).(pulumi.StringOutput)
}

type CSharpFunctionBindingResponse struct {
	Class   *string `pulumi:"class"`
	DllPath *string `pulumi:"dllPath"`
	Method  *string `pulumi:"method"`
	Script  *string `pulumi:"script"`
	Type    string  `pulumi:"type"`
}





type CSharpFunctionBindingResponseInput interface {
	pulumi.Input

	ToCSharpFunctionBindingResponseOutput() CSharpFunctionBindingResponseOutput
	ToCSharpFunctionBindingResponseOutputWithContext(context.Context) CSharpFunctionBindingResponseOutput
}

type CSharpFunctionBindingResponseArgs struct {
	Class   pulumi.StringPtrInput `pulumi:"class"`
	DllPath pulumi.StringPtrInput `pulumi:"dllPath"`
	Method  pulumi.StringPtrInput `pulumi:"method"`
	Script  pulumi.StringPtrInput `pulumi:"script"`
	Type    pulumi.StringInput    `pulumi:"type"`
}

func (CSharpFunctionBindingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CSharpFunctionBindingResponse)(nil)).Elem()
}

func (i CSharpFunctionBindingResponseArgs) ToCSharpFunctionBindingResponseOutput() CSharpFunctionBindingResponseOutput {
	return i.ToCSharpFunctionBindingResponseOutputWithContext(context.Background())
}

func (i CSharpFunctionBindingResponseArgs) ToCSharpFunctionBindingResponseOutputWithContext(ctx context.Context) CSharpFunctionBindingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CSharpFunctionBindingResponseOutput)
}

type CSharpFunctionBindingResponseOutput struct{ *pulumi.OutputState }

func (CSharpFunctionBindingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CSharpFunctionBindingResponse)(nil)).Elem()
}

func (o CSharpFunctionBindingResponseOutput) ToCSharpFunctionBindingResponseOutput() CSharpFunctionBindingResponseOutput {
	return o
}

func (o CSharpFunctionBindingResponseOutput) ToCSharpFunctionBindingResponseOutputWithContext(ctx context.Context) CSharpFunctionBindingResponseOutput {
	return o
}

func (o CSharpFunctionBindingResponseOutput) Class() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBindingResponse) *string { return v.Class }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingResponseOutput) DllPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBindingResponse) *string { return v.DllPath }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBindingResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingResponseOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CSharpFunctionBindingResponse) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o CSharpFunctionBindingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CSharpFunctionBindingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ClusterInfo struct {
	Id *string `pulumi:"id"`
}





type ClusterInfoInput interface {
	pulumi.Input

	ToClusterInfoOutput() ClusterInfoOutput
	ToClusterInfoOutputWithContext(context.Context) ClusterInfoOutput
}

type ClusterInfoArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ClusterInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfo)(nil)).Elem()
}

func (i ClusterInfoArgs) ToClusterInfoOutput() ClusterInfoOutput {
	return i.ToClusterInfoOutputWithContext(context.Background())
}

func (i ClusterInfoArgs) ToClusterInfoOutputWithContext(ctx context.Context) ClusterInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoOutput)
}

func (i ClusterInfoArgs) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return i.ToClusterInfoPtrOutputWithContext(context.Background())
}

func (i ClusterInfoArgs) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoOutput).ToClusterInfoPtrOutputWithContext(ctx)
}









type ClusterInfoPtrInput interface {
	pulumi.Input

	ToClusterInfoPtrOutput() ClusterInfoPtrOutput
	ToClusterInfoPtrOutputWithContext(context.Context) ClusterInfoPtrOutput
}

type clusterInfoPtrType ClusterInfoArgs

func ClusterInfoPtr(v *ClusterInfoArgs) ClusterInfoPtrInput {
	return (*clusterInfoPtrType)(v)
}

func (*clusterInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfo)(nil)).Elem()
}

func (i *clusterInfoPtrType) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return i.ToClusterInfoPtrOutputWithContext(context.Background())
}

func (i *clusterInfoPtrType) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoPtrOutput)
}

type ClusterInfoOutput struct{ *pulumi.OutputState }

func (ClusterInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfo)(nil)).Elem()
}

func (o ClusterInfoOutput) ToClusterInfoOutput() ClusterInfoOutput {
	return o
}

func (o ClusterInfoOutput) ToClusterInfoOutputWithContext(ctx context.Context) ClusterInfoOutput {
	return o
}

func (o ClusterInfoOutput) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return o.ToClusterInfoPtrOutputWithContext(context.Background())
}

func (o ClusterInfoOutput) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterInfo) *ClusterInfo {
		return &v
	}).(ClusterInfoPtrOutput)
}

func (o ClusterInfoOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInfo) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ClusterInfoPtrOutput struct{ *pulumi.OutputState }

func (ClusterInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfo)(nil)).Elem()
}

func (o ClusterInfoPtrOutput) ToClusterInfoPtrOutput() ClusterInfoPtrOutput {
	return o
}

func (o ClusterInfoPtrOutput) ToClusterInfoPtrOutputWithContext(ctx context.Context) ClusterInfoPtrOutput {
	return o
}

func (o ClusterInfoPtrOutput) Elem() ClusterInfoOutput {
	return o.ApplyT(func(v *ClusterInfo) ClusterInfo {
		if v != nil {
			return *v
		}
		var ret ClusterInfo
		return ret
	}).(ClusterInfoOutput)
}

func (o ClusterInfoPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterInfo) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ClusterInfoResponse struct {
	Id *string `pulumi:"id"`
}





type ClusterInfoResponseInput interface {
	pulumi.Input

	ToClusterInfoResponseOutput() ClusterInfoResponseOutput
	ToClusterInfoResponseOutputWithContext(context.Context) ClusterInfoResponseOutput
}

type ClusterInfoResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ClusterInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfoResponse)(nil)).Elem()
}

func (i ClusterInfoResponseArgs) ToClusterInfoResponseOutput() ClusterInfoResponseOutput {
	return i.ToClusterInfoResponseOutputWithContext(context.Background())
}

func (i ClusterInfoResponseArgs) ToClusterInfoResponseOutputWithContext(ctx context.Context) ClusterInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoResponseOutput)
}

func (i ClusterInfoResponseArgs) ToClusterInfoResponsePtrOutput() ClusterInfoResponsePtrOutput {
	return i.ToClusterInfoResponsePtrOutputWithContext(context.Background())
}

func (i ClusterInfoResponseArgs) ToClusterInfoResponsePtrOutputWithContext(ctx context.Context) ClusterInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoResponseOutput).ToClusterInfoResponsePtrOutputWithContext(ctx)
}









type ClusterInfoResponsePtrInput interface {
	pulumi.Input

	ToClusterInfoResponsePtrOutput() ClusterInfoResponsePtrOutput
	ToClusterInfoResponsePtrOutputWithContext(context.Context) ClusterInfoResponsePtrOutput
}

type clusterInfoResponsePtrType ClusterInfoResponseArgs

func ClusterInfoResponsePtr(v *ClusterInfoResponseArgs) ClusterInfoResponsePtrInput {
	return (*clusterInfoResponsePtrType)(v)
}

func (*clusterInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfoResponse)(nil)).Elem()
}

func (i *clusterInfoResponsePtrType) ToClusterInfoResponsePtrOutput() ClusterInfoResponsePtrOutput {
	return i.ToClusterInfoResponsePtrOutputWithContext(context.Background())
}

func (i *clusterInfoResponsePtrType) ToClusterInfoResponsePtrOutputWithContext(ctx context.Context) ClusterInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterInfoResponsePtrOutput)
}

type ClusterInfoResponseOutput struct{ *pulumi.OutputState }

func (ClusterInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterInfoResponse)(nil)).Elem()
}

func (o ClusterInfoResponseOutput) ToClusterInfoResponseOutput() ClusterInfoResponseOutput {
	return o
}

func (o ClusterInfoResponseOutput) ToClusterInfoResponseOutputWithContext(ctx context.Context) ClusterInfoResponseOutput {
	return o
}

func (o ClusterInfoResponseOutput) ToClusterInfoResponsePtrOutput() ClusterInfoResponsePtrOutput {
	return o.ToClusterInfoResponsePtrOutputWithContext(context.Background())
}

func (o ClusterInfoResponseOutput) ToClusterInfoResponsePtrOutputWithContext(ctx context.Context) ClusterInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterInfoResponse) *ClusterInfoResponse {
		return &v
	}).(ClusterInfoResponsePtrOutput)
}

func (o ClusterInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ClusterInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ClusterInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterInfoResponse)(nil)).Elem()
}

func (o ClusterInfoResponsePtrOutput) ToClusterInfoResponsePtrOutput() ClusterInfoResponsePtrOutput {
	return o
}

func (o ClusterInfoResponsePtrOutput) ToClusterInfoResponsePtrOutputWithContext(ctx context.Context) ClusterInfoResponsePtrOutput {
	return o
}

func (o ClusterInfoResponsePtrOutput) Elem() ClusterInfoResponseOutput {
	return o.ApplyT(func(v *ClusterInfoResponse) ClusterInfoResponse {
		if v != nil {
			return *v
		}
		var ret ClusterInfoResponse
		return ret
	}).(ClusterInfoResponseOutput)
}

func (o ClusterInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type Compression struct {
	Type string `pulumi:"type"`
}





type CompressionInput interface {
	pulumi.Input

	ToCompressionOutput() CompressionOutput
	ToCompressionOutputWithContext(context.Context) CompressionOutput
}

type CompressionArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (CompressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Compression)(nil)).Elem()
}

func (i CompressionArgs) ToCompressionOutput() CompressionOutput {
	return i.ToCompressionOutputWithContext(context.Background())
}

func (i CompressionArgs) ToCompressionOutputWithContext(ctx context.Context) CompressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionOutput)
}

func (i CompressionArgs) ToCompressionPtrOutput() CompressionPtrOutput {
	return i.ToCompressionPtrOutputWithContext(context.Background())
}

func (i CompressionArgs) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionOutput).ToCompressionPtrOutputWithContext(ctx)
}









type CompressionPtrInput interface {
	pulumi.Input

	ToCompressionPtrOutput() CompressionPtrOutput
	ToCompressionPtrOutputWithContext(context.Context) CompressionPtrOutput
}

type compressionPtrType CompressionArgs

func CompressionPtr(v *CompressionArgs) CompressionPtrInput {
	return (*compressionPtrType)(v)
}

func (*compressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Compression)(nil)).Elem()
}

func (i *compressionPtrType) ToCompressionPtrOutput() CompressionPtrOutput {
	return i.ToCompressionPtrOutputWithContext(context.Background())
}

func (i *compressionPtrType) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionPtrOutput)
}

type CompressionOutput struct{ *pulumi.OutputState }

func (CompressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Compression)(nil)).Elem()
}

func (o CompressionOutput) ToCompressionOutput() CompressionOutput {
	return o
}

func (o CompressionOutput) ToCompressionOutputWithContext(ctx context.Context) CompressionOutput {
	return o
}

func (o CompressionOutput) ToCompressionPtrOutput() CompressionPtrOutput {
	return o.ToCompressionPtrOutputWithContext(context.Background())
}

func (o CompressionOutput) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Compression) *Compression {
		return &v
	}).(CompressionPtrOutput)
}

func (o CompressionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Compression) string { return v.Type }).(pulumi.StringOutput)
}

type CompressionPtrOutput struct{ *pulumi.OutputState }

func (CompressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compression)(nil)).Elem()
}

func (o CompressionPtrOutput) ToCompressionPtrOutput() CompressionPtrOutput {
	return o
}

func (o CompressionPtrOutput) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return o
}

func (o CompressionPtrOutput) Elem() CompressionOutput {
	return o.ApplyT(func(v *Compression) Compression {
		if v != nil {
			return *v
		}
		var ret Compression
		return ret
	}).(CompressionOutput)
}

func (o CompressionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compression) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type CompressionResponse struct {
	Type string `pulumi:"type"`
}





type CompressionResponseInput interface {
	pulumi.Input

	ToCompressionResponseOutput() CompressionResponseOutput
	ToCompressionResponseOutputWithContext(context.Context) CompressionResponseOutput
}

type CompressionResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (CompressionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CompressionResponse)(nil)).Elem()
}

func (i CompressionResponseArgs) ToCompressionResponseOutput() CompressionResponseOutput {
	return i.ToCompressionResponseOutputWithContext(context.Background())
}

func (i CompressionResponseArgs) ToCompressionResponseOutputWithContext(ctx context.Context) CompressionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionResponseOutput)
}

func (i CompressionResponseArgs) ToCompressionResponsePtrOutput() CompressionResponsePtrOutput {
	return i.ToCompressionResponsePtrOutputWithContext(context.Background())
}

func (i CompressionResponseArgs) ToCompressionResponsePtrOutputWithContext(ctx context.Context) CompressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionResponseOutput).ToCompressionResponsePtrOutputWithContext(ctx)
}









type CompressionResponsePtrInput interface {
	pulumi.Input

	ToCompressionResponsePtrOutput() CompressionResponsePtrOutput
	ToCompressionResponsePtrOutputWithContext(context.Context) CompressionResponsePtrOutput
}

type compressionResponsePtrType CompressionResponseArgs

func CompressionResponsePtr(v *CompressionResponseArgs) CompressionResponsePtrInput {
	return (*compressionResponsePtrType)(v)
}

func (*compressionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CompressionResponse)(nil)).Elem()
}

func (i *compressionResponsePtrType) ToCompressionResponsePtrOutput() CompressionResponsePtrOutput {
	return i.ToCompressionResponsePtrOutputWithContext(context.Background())
}

func (i *compressionResponsePtrType) ToCompressionResponsePtrOutputWithContext(ctx context.Context) CompressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompressionResponsePtrOutput)
}

type CompressionResponseOutput struct{ *pulumi.OutputState }

func (CompressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompressionResponse)(nil)).Elem()
}

func (o CompressionResponseOutput) ToCompressionResponseOutput() CompressionResponseOutput {
	return o
}

func (o CompressionResponseOutput) ToCompressionResponseOutputWithContext(ctx context.Context) CompressionResponseOutput {
	return o
}

func (o CompressionResponseOutput) ToCompressionResponsePtrOutput() CompressionResponsePtrOutput {
	return o.ToCompressionResponsePtrOutputWithContext(context.Background())
}

func (o CompressionResponseOutput) ToCompressionResponsePtrOutputWithContext(ctx context.Context) CompressionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CompressionResponse) *CompressionResponse {
		return &v
	}).(CompressionResponsePtrOutput)
}

func (o CompressionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CompressionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type CompressionResponsePtrOutput struct{ *pulumi.OutputState }

func (CompressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompressionResponse)(nil)).Elem()
}

func (o CompressionResponsePtrOutput) ToCompressionResponsePtrOutput() CompressionResponsePtrOutput {
	return o
}

func (o CompressionResponsePtrOutput) ToCompressionResponsePtrOutputWithContext(ctx context.Context) CompressionResponsePtrOutput {
	return o
}

func (o CompressionResponsePtrOutput) Elem() CompressionResponseOutput {
	return o.ApplyT(func(v *CompressionResponse) CompressionResponse {
		if v != nil {
			return *v
		}
		var ret CompressionResponse
		return ret
	}).(CompressionResponseOutput)
}

func (o CompressionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type CsvSerialization struct {
	Encoding       *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Type           string  `pulumi:"type"`
}





type CsvSerializationInput interface {
	pulumi.Input

	ToCsvSerializationOutput() CsvSerializationOutput
	ToCsvSerializationOutputWithContext(context.Context) CsvSerializationOutput
}

type CsvSerializationArgs struct {
	Encoding       pulumi.StringPtrInput `pulumi:"encoding"`
	FieldDelimiter pulumi.StringPtrInput `pulumi:"fieldDelimiter"`
	Type           pulumi.StringInput    `pulumi:"type"`
}

func (CsvSerializationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CsvSerialization)(nil)).Elem()
}

func (i CsvSerializationArgs) ToCsvSerializationOutput() CsvSerializationOutput {
	return i.ToCsvSerializationOutputWithContext(context.Background())
}

func (i CsvSerializationArgs) ToCsvSerializationOutputWithContext(ctx context.Context) CsvSerializationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CsvSerializationOutput)
}

type CsvSerializationOutput struct{ *pulumi.OutputState }

func (CsvSerializationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CsvSerialization)(nil)).Elem()
}

func (o CsvSerializationOutput) ToCsvSerializationOutput() CsvSerializationOutput {
	return o
}

func (o CsvSerializationOutput) ToCsvSerializationOutputWithContext(ctx context.Context) CsvSerializationOutput {
	return o
}

func (o CsvSerializationOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CsvSerialization) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o CsvSerializationOutput) FieldDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CsvSerialization) *string { return v.FieldDelimiter }).(pulumi.StringPtrOutput)
}

func (o CsvSerializationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CsvSerialization) string { return v.Type }).(pulumi.StringOutput)
}

type CsvSerializationResponse struct {
	Encoding       *string `pulumi:"encoding"`
	FieldDelimiter *string `pulumi:"fieldDelimiter"`
	Type           string  `pulumi:"type"`
}





type CsvSerializationResponseInput interface {
	pulumi.Input

	ToCsvSerializationResponseOutput() CsvSerializationResponseOutput
	ToCsvSerializationResponseOutputWithContext(context.Context) CsvSerializationResponseOutput
}

type CsvSerializationResponseArgs struct {
	Encoding       pulumi.StringPtrInput `pulumi:"encoding"`
	FieldDelimiter pulumi.StringPtrInput `pulumi:"fieldDelimiter"`
	Type           pulumi.StringInput    `pulumi:"type"`
}

func (CsvSerializationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CsvSerializationResponse)(nil)).Elem()
}

func (i CsvSerializationResponseArgs) ToCsvSerializationResponseOutput() CsvSerializationResponseOutput {
	return i.ToCsvSerializationResponseOutputWithContext(context.Background())
}

func (i CsvSerializationResponseArgs) ToCsvSerializationResponseOutputWithContext(ctx context.Context) CsvSerializationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CsvSerializationResponseOutput)
}

type CsvSerializationResponseOutput struct{ *pulumi.OutputState }

func (CsvSerializationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CsvSerializationResponse)(nil)).Elem()
}

func (o CsvSerializationResponseOutput) ToCsvSerializationResponseOutput() CsvSerializationResponseOutput {
	return o
}

func (o CsvSerializationResponseOutput) ToCsvSerializationResponseOutputWithContext(ctx context.Context) CsvSerializationResponseOutput {
	return o
}

func (o CsvSerializationResponseOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CsvSerializationResponse) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o CsvSerializationResponseOutput) FieldDelimiter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CsvSerializationResponse) *string { return v.FieldDelimiter }).(pulumi.StringPtrOutput)
}

func (o CsvSerializationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CsvSerializationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type CustomClrSerialization struct {
	SerializationClassName *string `pulumi:"serializationClassName"`
	SerializationDllPath   *string `pulumi:"serializationDllPath"`
	Type                   string  `pulumi:"type"`
}





type CustomClrSerializationInput interface {
	pulumi.Input

	ToCustomClrSerializationOutput() CustomClrSerializationOutput
	ToCustomClrSerializationOutputWithContext(context.Context) CustomClrSerializationOutput
}

type CustomClrSerializationArgs struct {
	SerializationClassName pulumi.StringPtrInput `pulumi:"serializationClassName"`
	SerializationDllPath   pulumi.StringPtrInput `pulumi:"serializationDllPath"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (CustomClrSerializationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomClrSerialization)(nil)).Elem()
}

func (i CustomClrSerializationArgs) ToCustomClrSerializationOutput() CustomClrSerializationOutput {
	return i.ToCustomClrSerializationOutputWithContext(context.Background())
}

func (i CustomClrSerializationArgs) ToCustomClrSerializationOutputWithContext(ctx context.Context) CustomClrSerializationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomClrSerializationOutput)
}

type CustomClrSerializationOutput struct{ *pulumi.OutputState }

func (CustomClrSerializationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomClrSerialization)(nil)).Elem()
}

func (o CustomClrSerializationOutput) ToCustomClrSerializationOutput() CustomClrSerializationOutput {
	return o
}

func (o CustomClrSerializationOutput) ToCustomClrSerializationOutputWithContext(ctx context.Context) CustomClrSerializationOutput {
	return o
}

func (o CustomClrSerializationOutput) SerializationClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomClrSerialization) *string { return v.SerializationClassName }).(pulumi.StringPtrOutput)
}

func (o CustomClrSerializationOutput) SerializationDllPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomClrSerialization) *string { return v.SerializationDllPath }).(pulumi.StringPtrOutput)
}

func (o CustomClrSerializationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CustomClrSerialization) string { return v.Type }).(pulumi.StringOutput)
}

type CustomClrSerializationResponse struct {
	SerializationClassName *string `pulumi:"serializationClassName"`
	SerializationDllPath   *string `pulumi:"serializationDllPath"`
	Type                   string  `pulumi:"type"`
}





type CustomClrSerializationResponseInput interface {
	pulumi.Input

	ToCustomClrSerializationResponseOutput() CustomClrSerializationResponseOutput
	ToCustomClrSerializationResponseOutputWithContext(context.Context) CustomClrSerializationResponseOutput
}

type CustomClrSerializationResponseArgs struct {
	SerializationClassName pulumi.StringPtrInput `pulumi:"serializationClassName"`
	SerializationDllPath   pulumi.StringPtrInput `pulumi:"serializationDllPath"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (CustomClrSerializationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomClrSerializationResponse)(nil)).Elem()
}

func (i CustomClrSerializationResponseArgs) ToCustomClrSerializationResponseOutput() CustomClrSerializationResponseOutput {
	return i.ToCustomClrSerializationResponseOutputWithContext(context.Background())
}

func (i CustomClrSerializationResponseArgs) ToCustomClrSerializationResponseOutputWithContext(ctx context.Context) CustomClrSerializationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomClrSerializationResponseOutput)
}

type CustomClrSerializationResponseOutput struct{ *pulumi.OutputState }

func (CustomClrSerializationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomClrSerializationResponse)(nil)).Elem()
}

func (o CustomClrSerializationResponseOutput) ToCustomClrSerializationResponseOutput() CustomClrSerializationResponseOutput {
	return o
}

func (o CustomClrSerializationResponseOutput) ToCustomClrSerializationResponseOutputWithContext(ctx context.Context) CustomClrSerializationResponseOutput {
	return o
}

func (o CustomClrSerializationResponseOutput) SerializationClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomClrSerializationResponse) *string { return v.SerializationClassName }).(pulumi.StringPtrOutput)
}

func (o CustomClrSerializationResponseOutput) SerializationDllPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomClrSerializationResponse) *string { return v.SerializationDllPath }).(pulumi.StringPtrOutput)
}

func (o CustomClrSerializationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CustomClrSerializationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DiagnosticConditionResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
	Since   string `pulumi:"since"`
}





type DiagnosticConditionResponseInput interface {
	pulumi.Input

	ToDiagnosticConditionResponseOutput() DiagnosticConditionResponseOutput
	ToDiagnosticConditionResponseOutputWithContext(context.Context) DiagnosticConditionResponseOutput
}

type DiagnosticConditionResponseArgs struct {
	Code    pulumi.StringInput `pulumi:"code"`
	Message pulumi.StringInput `pulumi:"message"`
	Since   pulumi.StringInput `pulumi:"since"`
}

func (DiagnosticConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticConditionResponse)(nil)).Elem()
}

func (i DiagnosticConditionResponseArgs) ToDiagnosticConditionResponseOutput() DiagnosticConditionResponseOutput {
	return i.ToDiagnosticConditionResponseOutputWithContext(context.Background())
}

func (i DiagnosticConditionResponseArgs) ToDiagnosticConditionResponseOutputWithContext(ctx context.Context) DiagnosticConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticConditionResponseOutput)
}





type DiagnosticConditionResponseArrayInput interface {
	pulumi.Input

	ToDiagnosticConditionResponseArrayOutput() DiagnosticConditionResponseArrayOutput
	ToDiagnosticConditionResponseArrayOutputWithContext(context.Context) DiagnosticConditionResponseArrayOutput
}

type DiagnosticConditionResponseArray []DiagnosticConditionResponseInput

func (DiagnosticConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiagnosticConditionResponse)(nil)).Elem()
}

func (i DiagnosticConditionResponseArray) ToDiagnosticConditionResponseArrayOutput() DiagnosticConditionResponseArrayOutput {
	return i.ToDiagnosticConditionResponseArrayOutputWithContext(context.Background())
}

func (i DiagnosticConditionResponseArray) ToDiagnosticConditionResponseArrayOutputWithContext(ctx context.Context) DiagnosticConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticConditionResponseArrayOutput)
}

type DiagnosticConditionResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticConditionResponse)(nil)).Elem()
}

func (o DiagnosticConditionResponseOutput) ToDiagnosticConditionResponseOutput() DiagnosticConditionResponseOutput {
	return o
}

func (o DiagnosticConditionResponseOutput) ToDiagnosticConditionResponseOutputWithContext(ctx context.Context) DiagnosticConditionResponseOutput {
	return o
}

func (o DiagnosticConditionResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o DiagnosticConditionResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o DiagnosticConditionResponseOutput) Since() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticConditionResponse) string { return v.Since }).(pulumi.StringOutput)
}

type DiagnosticConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (DiagnosticConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiagnosticConditionResponse)(nil)).Elem()
}

func (o DiagnosticConditionResponseArrayOutput) ToDiagnosticConditionResponseArrayOutput() DiagnosticConditionResponseArrayOutput {
	return o
}

func (o DiagnosticConditionResponseArrayOutput) ToDiagnosticConditionResponseArrayOutputWithContext(ctx context.Context) DiagnosticConditionResponseArrayOutput {
	return o
}

func (o DiagnosticConditionResponseArrayOutput) Index(i pulumi.IntInput) DiagnosticConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiagnosticConditionResponse {
		return vs[0].([]DiagnosticConditionResponse)[vs[1].(int)]
	}).(DiagnosticConditionResponseOutput)
}

type DiagnosticsResponse struct {
	Conditions []DiagnosticConditionResponse `pulumi:"conditions"`
}





type DiagnosticsResponseInput interface {
	pulumi.Input

	ToDiagnosticsResponseOutput() DiagnosticsResponseOutput
	ToDiagnosticsResponseOutputWithContext(context.Context) DiagnosticsResponseOutput
}

type DiagnosticsResponseArgs struct {
	Conditions DiagnosticConditionResponseArrayInput `pulumi:"conditions"`
}

func (DiagnosticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsResponse)(nil)).Elem()
}

func (i DiagnosticsResponseArgs) ToDiagnosticsResponseOutput() DiagnosticsResponseOutput {
	return i.ToDiagnosticsResponseOutputWithContext(context.Background())
}

func (i DiagnosticsResponseArgs) ToDiagnosticsResponseOutputWithContext(ctx context.Context) DiagnosticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsResponseOutput)
}

func (i DiagnosticsResponseArgs) ToDiagnosticsResponsePtrOutput() DiagnosticsResponsePtrOutput {
	return i.ToDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (i DiagnosticsResponseArgs) ToDiagnosticsResponsePtrOutputWithContext(ctx context.Context) DiagnosticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsResponseOutput).ToDiagnosticsResponsePtrOutputWithContext(ctx)
}









type DiagnosticsResponsePtrInput interface {
	pulumi.Input

	ToDiagnosticsResponsePtrOutput() DiagnosticsResponsePtrOutput
	ToDiagnosticsResponsePtrOutputWithContext(context.Context) DiagnosticsResponsePtrOutput
}

type diagnosticsResponsePtrType DiagnosticsResponseArgs

func DiagnosticsResponsePtr(v *DiagnosticsResponseArgs) DiagnosticsResponsePtrInput {
	return (*diagnosticsResponsePtrType)(v)
}

func (*diagnosticsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsResponse)(nil)).Elem()
}

func (i *diagnosticsResponsePtrType) ToDiagnosticsResponsePtrOutput() DiagnosticsResponsePtrOutput {
	return i.ToDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (i *diagnosticsResponsePtrType) ToDiagnosticsResponsePtrOutputWithContext(ctx context.Context) DiagnosticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsResponsePtrOutput)
}

type DiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsResponse)(nil)).Elem()
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutput() DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutputWithContext(ctx context.Context) DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponsePtrOutput() DiagnosticsResponsePtrOutput {
	return o.ToDiagnosticsResponsePtrOutputWithContext(context.Background())
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponsePtrOutputWithContext(ctx context.Context) DiagnosticsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsResponse) *DiagnosticsResponse {
		return &v
	}).(DiagnosticsResponsePtrOutput)
}

func (o DiagnosticsResponseOutput) Conditions() DiagnosticConditionResponseArrayOutput {
	return o.ApplyT(func(v DiagnosticsResponse) []DiagnosticConditionResponse { return v.Conditions }).(DiagnosticConditionResponseArrayOutput)
}

type DiagnosticsResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsResponse)(nil)).Elem()
}

func (o DiagnosticsResponsePtrOutput) ToDiagnosticsResponsePtrOutput() DiagnosticsResponsePtrOutput {
	return o
}

func (o DiagnosticsResponsePtrOutput) ToDiagnosticsResponsePtrOutputWithContext(ctx context.Context) DiagnosticsResponsePtrOutput {
	return o
}

func (o DiagnosticsResponsePtrOutput) Elem() DiagnosticsResponseOutput {
	return o.ApplyT(func(v *DiagnosticsResponse) DiagnosticsResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsResponse
		return ret
	}).(DiagnosticsResponseOutput)
}

func (o DiagnosticsResponsePtrOutput) Conditions() DiagnosticConditionResponseArrayOutput {
	return o.ApplyT(func(v *DiagnosticsResponse) []DiagnosticConditionResponse {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(DiagnosticConditionResponseArrayOutput)
}

type DocumentDbOutputDataSource struct {
	AccountId             *string `pulumi:"accountId"`
	AccountKey            *string `pulumi:"accountKey"`
	CollectionNamePattern *string `pulumi:"collectionNamePattern"`
	Database              *string `pulumi:"database"`
	DocumentId            *string `pulumi:"documentId"`
	PartitionKey          *string `pulumi:"partitionKey"`
	Type                  string  `pulumi:"type"`
}





type DocumentDbOutputDataSourceInput interface {
	pulumi.Input

	ToDocumentDbOutputDataSourceOutput() DocumentDbOutputDataSourceOutput
	ToDocumentDbOutputDataSourceOutputWithContext(context.Context) DocumentDbOutputDataSourceOutput
}

type DocumentDbOutputDataSourceArgs struct {
	AccountId             pulumi.StringPtrInput `pulumi:"accountId"`
	AccountKey            pulumi.StringPtrInput `pulumi:"accountKey"`
	CollectionNamePattern pulumi.StringPtrInput `pulumi:"collectionNamePattern"`
	Database              pulumi.StringPtrInput `pulumi:"database"`
	DocumentId            pulumi.StringPtrInput `pulumi:"documentId"`
	PartitionKey          pulumi.StringPtrInput `pulumi:"partitionKey"`
	Type                  pulumi.StringInput    `pulumi:"type"`
}

func (DocumentDbOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentDbOutputDataSource)(nil)).Elem()
}

func (i DocumentDbOutputDataSourceArgs) ToDocumentDbOutputDataSourceOutput() DocumentDbOutputDataSourceOutput {
	return i.ToDocumentDbOutputDataSourceOutputWithContext(context.Background())
}

func (i DocumentDbOutputDataSourceArgs) ToDocumentDbOutputDataSourceOutputWithContext(ctx context.Context) DocumentDbOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDbOutputDataSourceOutput)
}

type DocumentDbOutputDataSourceOutput struct{ *pulumi.OutputState }

func (DocumentDbOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentDbOutputDataSource)(nil)).Elem()
}

func (o DocumentDbOutputDataSourceOutput) ToDocumentDbOutputDataSourceOutput() DocumentDbOutputDataSourceOutput {
	return o
}

func (o DocumentDbOutputDataSourceOutput) ToDocumentDbOutputDataSourceOutputWithContext(ctx context.Context) DocumentDbOutputDataSourceOutput {
	return o
}

func (o DocumentDbOutputDataSourceOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceOutput) CollectionNamePattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) *string { return v.CollectionNamePattern }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceOutput) DocumentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) *string { return v.DocumentId }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type DocumentDbOutputDataSourceResponse struct {
	AccountId             *string `pulumi:"accountId"`
	AccountKey            *string `pulumi:"accountKey"`
	CollectionNamePattern *string `pulumi:"collectionNamePattern"`
	Database              *string `pulumi:"database"`
	DocumentId            *string `pulumi:"documentId"`
	PartitionKey          *string `pulumi:"partitionKey"`
	Type                  string  `pulumi:"type"`
}





type DocumentDbOutputDataSourceResponseInput interface {
	pulumi.Input

	ToDocumentDbOutputDataSourceResponseOutput() DocumentDbOutputDataSourceResponseOutput
	ToDocumentDbOutputDataSourceResponseOutputWithContext(context.Context) DocumentDbOutputDataSourceResponseOutput
}

type DocumentDbOutputDataSourceResponseArgs struct {
	AccountId             pulumi.StringPtrInput `pulumi:"accountId"`
	AccountKey            pulumi.StringPtrInput `pulumi:"accountKey"`
	CollectionNamePattern pulumi.StringPtrInput `pulumi:"collectionNamePattern"`
	Database              pulumi.StringPtrInput `pulumi:"database"`
	DocumentId            pulumi.StringPtrInput `pulumi:"documentId"`
	PartitionKey          pulumi.StringPtrInput `pulumi:"partitionKey"`
	Type                  pulumi.StringInput    `pulumi:"type"`
}

func (DocumentDbOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentDbOutputDataSourceResponse)(nil)).Elem()
}

func (i DocumentDbOutputDataSourceResponseArgs) ToDocumentDbOutputDataSourceResponseOutput() DocumentDbOutputDataSourceResponseOutput {
	return i.ToDocumentDbOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i DocumentDbOutputDataSourceResponseArgs) ToDocumentDbOutputDataSourceResponseOutputWithContext(ctx context.Context) DocumentDbOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentDbOutputDataSourceResponseOutput)
}

type DocumentDbOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (DocumentDbOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DocumentDbOutputDataSourceResponse)(nil)).Elem()
}

func (o DocumentDbOutputDataSourceResponseOutput) ToDocumentDbOutputDataSourceResponseOutput() DocumentDbOutputDataSourceResponseOutput {
	return o
}

func (o DocumentDbOutputDataSourceResponseOutput) ToDocumentDbOutputDataSourceResponseOutputWithContext(ctx context.Context) DocumentDbOutputDataSourceResponseOutput {
	return o
}

func (o DocumentDbOutputDataSourceResponseOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceResponseOutput) CollectionNamePattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) *string { return v.CollectionNamePattern }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceResponseOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceResponseOutput) DocumentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) *string { return v.DocumentId }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o DocumentDbOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DocumentDbOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubOutputDataSource struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}





type EventHubOutputDataSourceInput interface {
	pulumi.Input

	ToEventHubOutputDataSourceOutput() EventHubOutputDataSourceOutput
	ToEventHubOutputDataSourceOutputWithContext(context.Context) EventHubOutputDataSourceOutput
}

type EventHubOutputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	EventHubName           pulumi.StringPtrInput   `pulumi:"eventHubName"`
	PartitionKey           pulumi.StringPtrInput   `pulumi:"partitionKey"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (EventHubOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubOutputDataSource)(nil)).Elem()
}

func (i EventHubOutputDataSourceArgs) ToEventHubOutputDataSourceOutput() EventHubOutputDataSourceOutput {
	return i.ToEventHubOutputDataSourceOutputWithContext(context.Background())
}

func (i EventHubOutputDataSourceArgs) ToEventHubOutputDataSourceOutputWithContext(ctx context.Context) EventHubOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutputDataSourceOutput)
}

type EventHubOutputDataSourceOutput struct{ *pulumi.OutputState }

func (EventHubOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubOutputDataSource)(nil)).Elem()
}

func (o EventHubOutputDataSourceOutput) ToEventHubOutputDataSourceOutput() EventHubOutputDataSourceOutput {
	return o
}

func (o EventHubOutputDataSourceOutput) ToEventHubOutputDataSourceOutputWithContext(ctx context.Context) EventHubOutputDataSourceOutput {
	return o
}

func (o EventHubOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o EventHubOutputDataSourceOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubOutputDataSourceResponse struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}





type EventHubOutputDataSourceResponseInput interface {
	pulumi.Input

	ToEventHubOutputDataSourceResponseOutput() EventHubOutputDataSourceResponseOutput
	ToEventHubOutputDataSourceResponseOutputWithContext(context.Context) EventHubOutputDataSourceResponseOutput
}

type EventHubOutputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	EventHubName           pulumi.StringPtrInput   `pulumi:"eventHubName"`
	PartitionKey           pulumi.StringPtrInput   `pulumi:"partitionKey"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (EventHubOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubOutputDataSourceResponse)(nil)).Elem()
}

func (i EventHubOutputDataSourceResponseArgs) ToEventHubOutputDataSourceResponseOutput() EventHubOutputDataSourceResponseOutput {
	return i.ToEventHubOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i EventHubOutputDataSourceResponseArgs) ToEventHubOutputDataSourceResponseOutputWithContext(ctx context.Context) EventHubOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutputDataSourceResponseOutput)
}

type EventHubOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (EventHubOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubOutputDataSourceResponse)(nil)).Elem()
}

func (o EventHubOutputDataSourceResponseOutput) ToEventHubOutputDataSourceResponseOutput() EventHubOutputDataSourceResponseOutput {
	return o
}

func (o EventHubOutputDataSourceResponseOutput) ToEventHubOutputDataSourceResponseOutputWithContext(ctx context.Context) EventHubOutputDataSourceResponseOutput {
	return o
}

func (o EventHubOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceResponseOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceResponseOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o EventHubOutputDataSourceResponseOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubStreamInputDataSource struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}





type EventHubStreamInputDataSourceInput interface {
	pulumi.Input

	ToEventHubStreamInputDataSourceOutput() EventHubStreamInputDataSourceOutput
	ToEventHubStreamInputDataSourceOutputWithContext(context.Context) EventHubStreamInputDataSourceOutput
}

type EventHubStreamInputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	ConsumerGroupName      pulumi.StringPtrInput `pulumi:"consumerGroupName"`
	EventHubName           pulumi.StringPtrInput `pulumi:"eventHubName"`
	ServiceBusNamespace    pulumi.StringPtrInput `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (EventHubStreamInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubStreamInputDataSource)(nil)).Elem()
}

func (i EventHubStreamInputDataSourceArgs) ToEventHubStreamInputDataSourceOutput() EventHubStreamInputDataSourceOutput {
	return i.ToEventHubStreamInputDataSourceOutputWithContext(context.Background())
}

func (i EventHubStreamInputDataSourceArgs) ToEventHubStreamInputDataSourceOutputWithContext(ctx context.Context) EventHubStreamInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubStreamInputDataSourceOutput)
}

type EventHubStreamInputDataSourceOutput struct{ *pulumi.OutputState }

func (EventHubStreamInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubStreamInputDataSource)(nil)).Elem()
}

func (o EventHubStreamInputDataSourceOutput) ToEventHubStreamInputDataSourceOutput() EventHubStreamInputDataSourceOutput {
	return o
}

func (o EventHubStreamInputDataSourceOutput) ToEventHubStreamInputDataSourceOutputWithContext(ctx context.Context) EventHubStreamInputDataSourceOutput {
	return o
}

func (o EventHubStreamInputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) *string { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubStreamInputDataSourceResponse struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}





type EventHubStreamInputDataSourceResponseInput interface {
	pulumi.Input

	ToEventHubStreamInputDataSourceResponseOutput() EventHubStreamInputDataSourceResponseOutput
	ToEventHubStreamInputDataSourceResponseOutputWithContext(context.Context) EventHubStreamInputDataSourceResponseOutput
}

type EventHubStreamInputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	ConsumerGroupName      pulumi.StringPtrInput `pulumi:"consumerGroupName"`
	EventHubName           pulumi.StringPtrInput `pulumi:"eventHubName"`
	ServiceBusNamespace    pulumi.StringPtrInput `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (EventHubStreamInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubStreamInputDataSourceResponse)(nil)).Elem()
}

func (i EventHubStreamInputDataSourceResponseArgs) ToEventHubStreamInputDataSourceResponseOutput() EventHubStreamInputDataSourceResponseOutput {
	return i.ToEventHubStreamInputDataSourceResponseOutputWithContext(context.Background())
}

func (i EventHubStreamInputDataSourceResponseArgs) ToEventHubStreamInputDataSourceResponseOutputWithContext(ctx context.Context) EventHubStreamInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubStreamInputDataSourceResponseOutput)
}

type EventHubStreamInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (EventHubStreamInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubStreamInputDataSourceResponse)(nil)).Elem()
}

func (o EventHubStreamInputDataSourceResponseOutput) ToEventHubStreamInputDataSourceResponseOutput() EventHubStreamInputDataSourceResponseOutput {
	return o
}

func (o EventHubStreamInputDataSourceResponseOutput) ToEventHubStreamInputDataSourceResponseOutputWithContext(ctx context.Context) EventHubStreamInputDataSourceResponseOutput {
	return o
}

func (o EventHubStreamInputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceResponseOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) *string { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceResponseOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceResponseOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubStreamInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubStreamInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubV2OutputDataSource struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}





type EventHubV2OutputDataSourceInput interface {
	pulumi.Input

	ToEventHubV2OutputDataSourceOutput() EventHubV2OutputDataSourceOutput
	ToEventHubV2OutputDataSourceOutputWithContext(context.Context) EventHubV2OutputDataSourceOutput
}

type EventHubV2OutputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	EventHubName           pulumi.StringPtrInput   `pulumi:"eventHubName"`
	PartitionKey           pulumi.StringPtrInput   `pulumi:"partitionKey"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (EventHubV2OutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2OutputDataSource)(nil)).Elem()
}

func (i EventHubV2OutputDataSourceArgs) ToEventHubV2OutputDataSourceOutput() EventHubV2OutputDataSourceOutput {
	return i.ToEventHubV2OutputDataSourceOutputWithContext(context.Background())
}

func (i EventHubV2OutputDataSourceArgs) ToEventHubV2OutputDataSourceOutputWithContext(ctx context.Context) EventHubV2OutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubV2OutputDataSourceOutput)
}

type EventHubV2OutputDataSourceOutput struct{ *pulumi.OutputState }

func (EventHubV2OutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2OutputDataSource)(nil)).Elem()
}

func (o EventHubV2OutputDataSourceOutput) ToEventHubV2OutputDataSourceOutput() EventHubV2OutputDataSourceOutput {
	return o
}

func (o EventHubV2OutputDataSourceOutput) ToEventHubV2OutputDataSourceOutputWithContext(ctx context.Context) EventHubV2OutputDataSourceOutput {
	return o
}

func (o EventHubV2OutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o EventHubV2OutputDataSourceOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubV2OutputDataSourceResponse struct {
	AuthenticationMode     *string  `pulumi:"authenticationMode"`
	EventHubName           *string  `pulumi:"eventHubName"`
	PartitionKey           *string  `pulumi:"partitionKey"`
	PropertyColumns        []string `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string  `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string  `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string  `pulumi:"sharedAccessPolicyName"`
	Type                   string   `pulumi:"type"`
}





type EventHubV2OutputDataSourceResponseInput interface {
	pulumi.Input

	ToEventHubV2OutputDataSourceResponseOutput() EventHubV2OutputDataSourceResponseOutput
	ToEventHubV2OutputDataSourceResponseOutputWithContext(context.Context) EventHubV2OutputDataSourceResponseOutput
}

type EventHubV2OutputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	EventHubName           pulumi.StringPtrInput   `pulumi:"eventHubName"`
	PartitionKey           pulumi.StringPtrInput   `pulumi:"partitionKey"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (EventHubV2OutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2OutputDataSourceResponse)(nil)).Elem()
}

func (i EventHubV2OutputDataSourceResponseArgs) ToEventHubV2OutputDataSourceResponseOutput() EventHubV2OutputDataSourceResponseOutput {
	return i.ToEventHubV2OutputDataSourceResponseOutputWithContext(context.Background())
}

func (i EventHubV2OutputDataSourceResponseArgs) ToEventHubV2OutputDataSourceResponseOutputWithContext(ctx context.Context) EventHubV2OutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubV2OutputDataSourceResponseOutput)
}

type EventHubV2OutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (EventHubV2OutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2OutputDataSourceResponse)(nil)).Elem()
}

func (o EventHubV2OutputDataSourceResponseOutput) ToEventHubV2OutputDataSourceResponseOutput() EventHubV2OutputDataSourceResponseOutput {
	return o
}

func (o EventHubV2OutputDataSourceResponseOutput) ToEventHubV2OutputDataSourceResponseOutputWithContext(ctx context.Context) EventHubV2OutputDataSourceResponseOutput {
	return o
}

func (o EventHubV2OutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2OutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubV2OutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubV2StreamInputDataSource struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}





type EventHubV2StreamInputDataSourceInput interface {
	pulumi.Input

	ToEventHubV2StreamInputDataSourceOutput() EventHubV2StreamInputDataSourceOutput
	ToEventHubV2StreamInputDataSourceOutputWithContext(context.Context) EventHubV2StreamInputDataSourceOutput
}

type EventHubV2StreamInputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	ConsumerGroupName      pulumi.StringPtrInput `pulumi:"consumerGroupName"`
	EventHubName           pulumi.StringPtrInput `pulumi:"eventHubName"`
	ServiceBusNamespace    pulumi.StringPtrInput `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (EventHubV2StreamInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2StreamInputDataSource)(nil)).Elem()
}

func (i EventHubV2StreamInputDataSourceArgs) ToEventHubV2StreamInputDataSourceOutput() EventHubV2StreamInputDataSourceOutput {
	return i.ToEventHubV2StreamInputDataSourceOutputWithContext(context.Background())
}

func (i EventHubV2StreamInputDataSourceArgs) ToEventHubV2StreamInputDataSourceOutputWithContext(ctx context.Context) EventHubV2StreamInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubV2StreamInputDataSourceOutput)
}

type EventHubV2StreamInputDataSourceOutput struct{ *pulumi.OutputState }

func (EventHubV2StreamInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2StreamInputDataSource)(nil)).Elem()
}

func (o EventHubV2StreamInputDataSourceOutput) ToEventHubV2StreamInputDataSourceOutput() EventHubV2StreamInputDataSourceOutput {
	return o
}

func (o EventHubV2StreamInputDataSourceOutput) ToEventHubV2StreamInputDataSourceOutputWithContext(ctx context.Context) EventHubV2StreamInputDataSourceOutput {
	return o
}

func (o EventHubV2StreamInputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) *string { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type EventHubV2StreamInputDataSourceResponse struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	EventHubName           *string `pulumi:"eventHubName"`
	ServiceBusNamespace    *string `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}





type EventHubV2StreamInputDataSourceResponseInput interface {
	pulumi.Input

	ToEventHubV2StreamInputDataSourceResponseOutput() EventHubV2StreamInputDataSourceResponseOutput
	ToEventHubV2StreamInputDataSourceResponseOutputWithContext(context.Context) EventHubV2StreamInputDataSourceResponseOutput
}

type EventHubV2StreamInputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	ConsumerGroupName      pulumi.StringPtrInput `pulumi:"consumerGroupName"`
	EventHubName           pulumi.StringPtrInput `pulumi:"eventHubName"`
	ServiceBusNamespace    pulumi.StringPtrInput `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (EventHubV2StreamInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2StreamInputDataSourceResponse)(nil)).Elem()
}

func (i EventHubV2StreamInputDataSourceResponseArgs) ToEventHubV2StreamInputDataSourceResponseOutput() EventHubV2StreamInputDataSourceResponseOutput {
	return i.ToEventHubV2StreamInputDataSourceResponseOutputWithContext(context.Background())
}

func (i EventHubV2StreamInputDataSourceResponseArgs) ToEventHubV2StreamInputDataSourceResponseOutputWithContext(ctx context.Context) EventHubV2StreamInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubV2StreamInputDataSourceResponseOutput)
}

type EventHubV2StreamInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (EventHubV2StreamInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubV2StreamInputDataSourceResponse)(nil)).Elem()
}

func (o EventHubV2StreamInputDataSourceResponseOutput) ToEventHubV2StreamInputDataSourceResponseOutput() EventHubV2StreamInputDataSourceResponseOutput {
	return o
}

func (o EventHubV2StreamInputDataSourceResponseOutput) ToEventHubV2StreamInputDataSourceResponseOutputWithContext(ctx context.Context) EventHubV2StreamInputDataSourceResponseOutput {
	return o
}

func (o EventHubV2StreamInputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceResponseOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) *string { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceResponseOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceResponseOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o EventHubV2StreamInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubV2StreamInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type External struct {
	Container      *string         `pulumi:"container"`
	Path           *string         `pulumi:"path"`
	StorageAccount *StorageAccount `pulumi:"storageAccount"`
}





type ExternalInput interface {
	pulumi.Input

	ToExternalOutput() ExternalOutput
	ToExternalOutputWithContext(context.Context) ExternalOutput
}

type ExternalArgs struct {
	Container      pulumi.StringPtrInput  `pulumi:"container"`
	Path           pulumi.StringPtrInput  `pulumi:"path"`
	StorageAccount StorageAccountPtrInput `pulumi:"storageAccount"`
}

func (ExternalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*External)(nil)).Elem()
}

func (i ExternalArgs) ToExternalOutput() ExternalOutput {
	return i.ToExternalOutputWithContext(context.Background())
}

func (i ExternalArgs) ToExternalOutputWithContext(ctx context.Context) ExternalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOutput)
}

func (i ExternalArgs) ToExternalPtrOutput() ExternalPtrOutput {
	return i.ToExternalPtrOutputWithContext(context.Background())
}

func (i ExternalArgs) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalOutput).ToExternalPtrOutputWithContext(ctx)
}









type ExternalPtrInput interface {
	pulumi.Input

	ToExternalPtrOutput() ExternalPtrOutput
	ToExternalPtrOutputWithContext(context.Context) ExternalPtrOutput
}

type externalPtrType ExternalArgs

func ExternalPtr(v *ExternalArgs) ExternalPtrInput {
	return (*externalPtrType)(v)
}

func (*externalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**External)(nil)).Elem()
}

func (i *externalPtrType) ToExternalPtrOutput() ExternalPtrOutput {
	return i.ToExternalPtrOutputWithContext(context.Background())
}

func (i *externalPtrType) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalPtrOutput)
}

type ExternalOutput struct{ *pulumi.OutputState }

func (ExternalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*External)(nil)).Elem()
}

func (o ExternalOutput) ToExternalOutput() ExternalOutput {
	return o
}

func (o ExternalOutput) ToExternalOutputWithContext(ctx context.Context) ExternalOutput {
	return o
}

func (o ExternalOutput) ToExternalPtrOutput() ExternalPtrOutput {
	return o.ToExternalPtrOutputWithContext(context.Background())
}

func (o ExternalOutput) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v External) *External {
		return &v
	}).(ExternalPtrOutput)
}

func (o ExternalOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v External) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o ExternalOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v External) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExternalOutput) StorageAccount() StorageAccountPtrOutput {
	return o.ApplyT(func(v External) *StorageAccount { return v.StorageAccount }).(StorageAccountPtrOutput)
}

type ExternalPtrOutput struct{ *pulumi.OutputState }

func (ExternalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**External)(nil)).Elem()
}

func (o ExternalPtrOutput) ToExternalPtrOutput() ExternalPtrOutput {
	return o
}

func (o ExternalPtrOutput) ToExternalPtrOutputWithContext(ctx context.Context) ExternalPtrOutput {
	return o
}

func (o ExternalPtrOutput) Elem() ExternalOutput {
	return o.ApplyT(func(v *External) External {
		if v != nil {
			return *v
		}
		var ret External
		return ret
	}).(ExternalOutput)
}

func (o ExternalPtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *External) *string {
		if v == nil {
			return nil
		}
		return v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExternalPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *External) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ExternalPtrOutput) StorageAccount() StorageAccountPtrOutput {
	return o.ApplyT(func(v *External) *StorageAccount {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(StorageAccountPtrOutput)
}

type ExternalResponse struct {
	Container      *string                 `pulumi:"container"`
	Path           *string                 `pulumi:"path"`
	StorageAccount *StorageAccountResponse `pulumi:"storageAccount"`
}





type ExternalResponseInput interface {
	pulumi.Input

	ToExternalResponseOutput() ExternalResponseOutput
	ToExternalResponseOutputWithContext(context.Context) ExternalResponseOutput
}

type ExternalResponseArgs struct {
	Container      pulumi.StringPtrInput          `pulumi:"container"`
	Path           pulumi.StringPtrInput          `pulumi:"path"`
	StorageAccount StorageAccountResponsePtrInput `pulumi:"storageAccount"`
}

func (ExternalResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalResponse)(nil)).Elem()
}

func (i ExternalResponseArgs) ToExternalResponseOutput() ExternalResponseOutput {
	return i.ToExternalResponseOutputWithContext(context.Background())
}

func (i ExternalResponseArgs) ToExternalResponseOutputWithContext(ctx context.Context) ExternalResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalResponseOutput)
}

func (i ExternalResponseArgs) ToExternalResponsePtrOutput() ExternalResponsePtrOutput {
	return i.ToExternalResponsePtrOutputWithContext(context.Background())
}

func (i ExternalResponseArgs) ToExternalResponsePtrOutputWithContext(ctx context.Context) ExternalResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalResponseOutput).ToExternalResponsePtrOutputWithContext(ctx)
}









type ExternalResponsePtrInput interface {
	pulumi.Input

	ToExternalResponsePtrOutput() ExternalResponsePtrOutput
	ToExternalResponsePtrOutputWithContext(context.Context) ExternalResponsePtrOutput
}

type externalResponsePtrType ExternalResponseArgs

func ExternalResponsePtr(v *ExternalResponseArgs) ExternalResponsePtrInput {
	return (*externalResponsePtrType)(v)
}

func (*externalResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalResponse)(nil)).Elem()
}

func (i *externalResponsePtrType) ToExternalResponsePtrOutput() ExternalResponsePtrOutput {
	return i.ToExternalResponsePtrOutputWithContext(context.Background())
}

func (i *externalResponsePtrType) ToExternalResponsePtrOutputWithContext(ctx context.Context) ExternalResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalResponsePtrOutput)
}

type ExternalResponseOutput struct{ *pulumi.OutputState }

func (ExternalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExternalResponse)(nil)).Elem()
}

func (o ExternalResponseOutput) ToExternalResponseOutput() ExternalResponseOutput {
	return o
}

func (o ExternalResponseOutput) ToExternalResponseOutputWithContext(ctx context.Context) ExternalResponseOutput {
	return o
}

func (o ExternalResponseOutput) ToExternalResponsePtrOutput() ExternalResponsePtrOutput {
	return o.ToExternalResponsePtrOutputWithContext(context.Background())
}

func (o ExternalResponseOutput) ToExternalResponsePtrOutputWithContext(ctx context.Context) ExternalResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExternalResponse) *ExternalResponse {
		return &v
	}).(ExternalResponsePtrOutput)
}

func (o ExternalResponseOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExternalResponse) *string { return v.Container }).(pulumi.StringPtrOutput)
}

func (o ExternalResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExternalResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExternalResponseOutput) StorageAccount() StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v ExternalResponse) *StorageAccountResponse { return v.StorageAccount }).(StorageAccountResponsePtrOutput)
}

type ExternalResponsePtrOutput struct{ *pulumi.OutputState }

func (ExternalResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalResponse)(nil)).Elem()
}

func (o ExternalResponsePtrOutput) ToExternalResponsePtrOutput() ExternalResponsePtrOutput {
	return o
}

func (o ExternalResponsePtrOutput) ToExternalResponsePtrOutputWithContext(ctx context.Context) ExternalResponsePtrOutput {
	return o
}

func (o ExternalResponsePtrOutput) Elem() ExternalResponseOutput {
	return o.ApplyT(func(v *ExternalResponse) ExternalResponse {
		if v != nil {
			return *v
		}
		var ret ExternalResponse
		return ret
	}).(ExternalResponseOutput)
}

func (o ExternalResponsePtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExternalResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o ExternalResponsePtrOutput) StorageAccount() StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v *ExternalResponse) *StorageAccountResponse {
		if v == nil {
			return nil
		}
		return v.StorageAccount
	}).(StorageAccountResponsePtrOutput)
}

type FunctionType struct {
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}





type FunctionTypeInput interface {
	pulumi.Input

	ToFunctionTypeOutput() FunctionTypeOutput
	ToFunctionTypeOutputWithContext(context.Context) FunctionTypeOutput
}

type FunctionTypeArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
}

func (FunctionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (i FunctionTypeArgs) ToFunctionTypeOutput() FunctionTypeOutput {
	return i.ToFunctionTypeOutputWithContext(context.Background())
}

func (i FunctionTypeArgs) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTypeOutput)
}





type FunctionTypeArrayInput interface {
	pulumi.Input

	ToFunctionTypeArrayOutput() FunctionTypeArrayOutput
	ToFunctionTypeArrayOutputWithContext(context.Context) FunctionTypeArrayOutput
}

type FunctionTypeArray []FunctionTypeInput

func (FunctionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionType)(nil)).Elem()
}

func (i FunctionTypeArray) ToFunctionTypeArrayOutput() FunctionTypeArrayOutput {
	return i.ToFunctionTypeArrayOutputWithContext(context.Background())
}

func (i FunctionTypeArray) ToFunctionTypeArrayOutputWithContext(ctx context.Context) FunctionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionTypeArrayOutput)
}

type FunctionTypeOutput struct{ *pulumi.OutputState }

func (FunctionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (o FunctionTypeOutput) ToFunctionTypeOutput() FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FunctionTypeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v FunctionType) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

type FunctionTypeArrayOutput struct{ *pulumi.OutputState }

func (FunctionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionType)(nil)).Elem()
}

func (o FunctionTypeArrayOutput) ToFunctionTypeArrayOutput() FunctionTypeArrayOutput {
	return o
}

func (o FunctionTypeArrayOutput) ToFunctionTypeArrayOutputWithContext(ctx context.Context) FunctionTypeArrayOutput {
	return o
}

func (o FunctionTypeArrayOutput) Index(i pulumi.IntInput) FunctionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionType {
		return vs[0].([]FunctionType)[vs[1].(int)]
	}).(FunctionTypeOutput)
}

type FunctionInputType struct {
	DataType                 *string `pulumi:"dataType"`
	IsConfigurationParameter *bool   `pulumi:"isConfigurationParameter"`
}





type FunctionInputTypeInput interface {
	pulumi.Input

	ToFunctionInputTypeOutput() FunctionInputTypeOutput
	ToFunctionInputTypeOutputWithContext(context.Context) FunctionInputTypeOutput
}

type FunctionInputTypeArgs struct {
	DataType                 pulumi.StringPtrInput `pulumi:"dataType"`
	IsConfigurationParameter pulumi.BoolPtrInput   `pulumi:"isConfigurationParameter"`
}

func (FunctionInputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputType)(nil)).Elem()
}

func (i FunctionInputTypeArgs) ToFunctionInputTypeOutput() FunctionInputTypeOutput {
	return i.ToFunctionInputTypeOutputWithContext(context.Background())
}

func (i FunctionInputTypeArgs) ToFunctionInputTypeOutputWithContext(ctx context.Context) FunctionInputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputTypeOutput)
}





type FunctionInputTypeArrayInput interface {
	pulumi.Input

	ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput
	ToFunctionInputTypeArrayOutputWithContext(context.Context) FunctionInputTypeArrayOutput
}

type FunctionInputTypeArray []FunctionInputTypeInput

func (FunctionInputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputType)(nil)).Elem()
}

func (i FunctionInputTypeArray) ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput {
	return i.ToFunctionInputTypeArrayOutputWithContext(context.Background())
}

func (i FunctionInputTypeArray) ToFunctionInputTypeArrayOutputWithContext(ctx context.Context) FunctionInputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputTypeArrayOutput)
}

type FunctionInputTypeOutput struct{ *pulumi.OutputState }

func (FunctionInputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputType)(nil)).Elem()
}

func (o FunctionInputTypeOutput) ToFunctionInputTypeOutput() FunctionInputTypeOutput {
	return o
}

func (o FunctionInputTypeOutput) ToFunctionInputTypeOutputWithContext(ctx context.Context) FunctionInputTypeOutput {
	return o
}

func (o FunctionInputTypeOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionInputType) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o FunctionInputTypeOutput) IsConfigurationParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FunctionInputType) *bool { return v.IsConfigurationParameter }).(pulumi.BoolPtrOutput)
}

type FunctionInputTypeArrayOutput struct{ *pulumi.OutputState }

func (FunctionInputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputType)(nil)).Elem()
}

func (o FunctionInputTypeArrayOutput) ToFunctionInputTypeArrayOutput() FunctionInputTypeArrayOutput {
	return o
}

func (o FunctionInputTypeArrayOutput) ToFunctionInputTypeArrayOutputWithContext(ctx context.Context) FunctionInputTypeArrayOutput {
	return o
}

func (o FunctionInputTypeArrayOutput) Index(i pulumi.IntInput) FunctionInputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionInputType {
		return vs[0].([]FunctionInputType)[vs[1].(int)]
	}).(FunctionInputTypeOutput)
}

type FunctionInputResponse struct {
	DataType                 *string `pulumi:"dataType"`
	IsConfigurationParameter *bool   `pulumi:"isConfigurationParameter"`
}





type FunctionInputResponseInput interface {
	pulumi.Input

	ToFunctionInputResponseOutput() FunctionInputResponseOutput
	ToFunctionInputResponseOutputWithContext(context.Context) FunctionInputResponseOutput
}

type FunctionInputResponseArgs struct {
	DataType                 pulumi.StringPtrInput `pulumi:"dataType"`
	IsConfigurationParameter pulumi.BoolPtrInput   `pulumi:"isConfigurationParameter"`
}

func (FunctionInputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputResponse)(nil)).Elem()
}

func (i FunctionInputResponseArgs) ToFunctionInputResponseOutput() FunctionInputResponseOutput {
	return i.ToFunctionInputResponseOutputWithContext(context.Background())
}

func (i FunctionInputResponseArgs) ToFunctionInputResponseOutputWithContext(ctx context.Context) FunctionInputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputResponseOutput)
}





type FunctionInputResponseArrayInput interface {
	pulumi.Input

	ToFunctionInputResponseArrayOutput() FunctionInputResponseArrayOutput
	ToFunctionInputResponseArrayOutputWithContext(context.Context) FunctionInputResponseArrayOutput
}

type FunctionInputResponseArray []FunctionInputResponseInput

func (FunctionInputResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputResponse)(nil)).Elem()
}

func (i FunctionInputResponseArray) ToFunctionInputResponseArrayOutput() FunctionInputResponseArrayOutput {
	return i.ToFunctionInputResponseArrayOutputWithContext(context.Background())
}

func (i FunctionInputResponseArray) ToFunctionInputResponseArrayOutputWithContext(ctx context.Context) FunctionInputResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionInputResponseArrayOutput)
}

type FunctionInputResponseOutput struct{ *pulumi.OutputState }

func (FunctionInputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionInputResponse)(nil)).Elem()
}

func (o FunctionInputResponseOutput) ToFunctionInputResponseOutput() FunctionInputResponseOutput {
	return o
}

func (o FunctionInputResponseOutput) ToFunctionInputResponseOutputWithContext(ctx context.Context) FunctionInputResponseOutput {
	return o
}

func (o FunctionInputResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionInputResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o FunctionInputResponseOutput) IsConfigurationParameter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FunctionInputResponse) *bool { return v.IsConfigurationParameter }).(pulumi.BoolPtrOutput)
}

type FunctionInputResponseArrayOutput struct{ *pulumi.OutputState }

func (FunctionInputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionInputResponse)(nil)).Elem()
}

func (o FunctionInputResponseArrayOutput) ToFunctionInputResponseArrayOutput() FunctionInputResponseArrayOutput {
	return o
}

func (o FunctionInputResponseArrayOutput) ToFunctionInputResponseArrayOutputWithContext(ctx context.Context) FunctionInputResponseArrayOutput {
	return o
}

func (o FunctionInputResponseArrayOutput) Index(i pulumi.IntInput) FunctionInputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionInputResponse {
		return vs[0].([]FunctionInputResponse)[vs[1].(int)]
	}).(FunctionInputResponseOutput)
}

type FunctionOutputType struct {
	DataType *string `pulumi:"dataType"`
}





type FunctionOutputTypeInput interface {
	pulumi.Input

	ToFunctionOutputTypeOutput() FunctionOutputTypeOutput
	ToFunctionOutputTypeOutputWithContext(context.Context) FunctionOutputTypeOutput
}

type FunctionOutputTypeArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
}

func (FunctionOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputType)(nil)).Elem()
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypeOutput() FunctionOutputTypeOutput {
	return i.ToFunctionOutputTypeOutputWithContext(context.Background())
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypeOutputWithContext(ctx context.Context) FunctionOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypeOutput)
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return i.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (i FunctionOutputTypeArgs) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypeOutput).ToFunctionOutputTypePtrOutputWithContext(ctx)
}









type FunctionOutputTypePtrInput interface {
	pulumi.Input

	ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput
	ToFunctionOutputTypePtrOutputWithContext(context.Context) FunctionOutputTypePtrOutput
}

type functionOutputTypePtrType FunctionOutputTypeArgs

func FunctionOutputTypePtr(v *FunctionOutputTypeArgs) FunctionOutputTypePtrInput {
	return (*functionOutputTypePtrType)(v)
}

func (*functionOutputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputType)(nil)).Elem()
}

func (i *functionOutputTypePtrType) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return i.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (i *functionOutputTypePtrType) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputTypePtrOutput)
}

type FunctionOutputTypeOutput struct{ *pulumi.OutputState }

func (FunctionOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputType)(nil)).Elem()
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypeOutput() FunctionOutputTypeOutput {
	return o
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypeOutputWithContext(ctx context.Context) FunctionOutputTypeOutput {
	return o
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return o.ToFunctionOutputTypePtrOutputWithContext(context.Background())
}

func (o FunctionOutputTypeOutput) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FunctionOutputType) *FunctionOutputType {
		return &v
	}).(FunctionOutputTypePtrOutput)
}

func (o FunctionOutputTypeOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionOutputType) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

type FunctionOutputTypePtrOutput struct{ *pulumi.OutputState }

func (FunctionOutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputType)(nil)).Elem()
}

func (o FunctionOutputTypePtrOutput) ToFunctionOutputTypePtrOutput() FunctionOutputTypePtrOutput {
	return o
}

func (o FunctionOutputTypePtrOutput) ToFunctionOutputTypePtrOutputWithContext(ctx context.Context) FunctionOutputTypePtrOutput {
	return o
}

func (o FunctionOutputTypePtrOutput) Elem() FunctionOutputTypeOutput {
	return o.ApplyT(func(v *FunctionOutputType) FunctionOutputType {
		if v != nil {
			return *v
		}
		var ret FunctionOutputType
		return ret
	}).(FunctionOutputTypeOutput)
}

func (o FunctionOutputTypePtrOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionOutputType) *string {
		if v == nil {
			return nil
		}
		return v.DataType
	}).(pulumi.StringPtrOutput)
}

type FunctionOutputResponse struct {
	DataType *string `pulumi:"dataType"`
}





type FunctionOutputResponseInput interface {
	pulumi.Input

	ToFunctionOutputResponseOutput() FunctionOutputResponseOutput
	ToFunctionOutputResponseOutputWithContext(context.Context) FunctionOutputResponseOutput
}

type FunctionOutputResponseArgs struct {
	DataType pulumi.StringPtrInput `pulumi:"dataType"`
}

func (FunctionOutputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputResponse)(nil)).Elem()
}

func (i FunctionOutputResponseArgs) ToFunctionOutputResponseOutput() FunctionOutputResponseOutput {
	return i.ToFunctionOutputResponseOutputWithContext(context.Background())
}

func (i FunctionOutputResponseArgs) ToFunctionOutputResponseOutputWithContext(ctx context.Context) FunctionOutputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputResponseOutput)
}

func (i FunctionOutputResponseArgs) ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput {
	return i.ToFunctionOutputResponsePtrOutputWithContext(context.Background())
}

func (i FunctionOutputResponseArgs) ToFunctionOutputResponsePtrOutputWithContext(ctx context.Context) FunctionOutputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputResponseOutput).ToFunctionOutputResponsePtrOutputWithContext(ctx)
}









type FunctionOutputResponsePtrInput interface {
	pulumi.Input

	ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput
	ToFunctionOutputResponsePtrOutputWithContext(context.Context) FunctionOutputResponsePtrOutput
}

type functionOutputResponsePtrType FunctionOutputResponseArgs

func FunctionOutputResponsePtr(v *FunctionOutputResponseArgs) FunctionOutputResponsePtrInput {
	return (*functionOutputResponsePtrType)(v)
}

func (*functionOutputResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputResponse)(nil)).Elem()
}

func (i *functionOutputResponsePtrType) ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput {
	return i.ToFunctionOutputResponsePtrOutputWithContext(context.Background())
}

func (i *functionOutputResponsePtrType) ToFunctionOutputResponsePtrOutputWithContext(ctx context.Context) FunctionOutputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionOutputResponsePtrOutput)
}

type FunctionOutputResponseOutput struct{ *pulumi.OutputState }

func (FunctionOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionOutputResponse)(nil)).Elem()
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponseOutput() FunctionOutputResponseOutput {
	return o
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponseOutputWithContext(ctx context.Context) FunctionOutputResponseOutput {
	return o
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput {
	return o.ToFunctionOutputResponsePtrOutputWithContext(context.Background())
}

func (o FunctionOutputResponseOutput) ToFunctionOutputResponsePtrOutputWithContext(ctx context.Context) FunctionOutputResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FunctionOutputResponse) *FunctionOutputResponse {
		return &v
	}).(FunctionOutputResponsePtrOutput)
}

func (o FunctionOutputResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionOutputResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

type FunctionOutputResponsePtrOutput struct{ *pulumi.OutputState }

func (FunctionOutputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionOutputResponse)(nil)).Elem()
}

func (o FunctionOutputResponsePtrOutput) ToFunctionOutputResponsePtrOutput() FunctionOutputResponsePtrOutput {
	return o
}

func (o FunctionOutputResponsePtrOutput) ToFunctionOutputResponsePtrOutputWithContext(ctx context.Context) FunctionOutputResponsePtrOutput {
	return o
}

func (o FunctionOutputResponsePtrOutput) Elem() FunctionOutputResponseOutput {
	return o.ApplyT(func(v *FunctionOutputResponse) FunctionOutputResponse {
		if v != nil {
			return *v
		}
		var ret FunctionOutputResponse
		return ret
	}).(FunctionOutputResponseOutput)
}

func (o FunctionOutputResponsePtrOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FunctionOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.DataType
	}).(pulumi.StringPtrOutput)
}

type FunctionResponse struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}





type FunctionResponseInput interface {
	pulumi.Input

	ToFunctionResponseOutput() FunctionResponseOutput
	ToFunctionResponseOutputWithContext(context.Context) FunctionResponseOutput
}

type FunctionResponseArgs struct {
	Id         pulumi.StringInput    `pulumi:"id"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (FunctionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionResponse)(nil)).Elem()
}

func (i FunctionResponseArgs) ToFunctionResponseOutput() FunctionResponseOutput {
	return i.ToFunctionResponseOutputWithContext(context.Background())
}

func (i FunctionResponseArgs) ToFunctionResponseOutputWithContext(ctx context.Context) FunctionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionResponseOutput)
}





type FunctionResponseArrayInput interface {
	pulumi.Input

	ToFunctionResponseArrayOutput() FunctionResponseArrayOutput
	ToFunctionResponseArrayOutputWithContext(context.Context) FunctionResponseArrayOutput
}

type FunctionResponseArray []FunctionResponseInput

func (FunctionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionResponse)(nil)).Elem()
}

func (i FunctionResponseArray) ToFunctionResponseArrayOutput() FunctionResponseArrayOutput {
	return i.ToFunctionResponseArrayOutputWithContext(context.Background())
}

func (i FunctionResponseArray) ToFunctionResponseArrayOutputWithContext(ctx context.Context) FunctionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionResponseArrayOutput)
}

type FunctionResponseOutput struct{ *pulumi.OutputState }

func (FunctionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionResponse)(nil)).Elem()
}

func (o FunctionResponseOutput) ToFunctionResponseOutput() FunctionResponseOutput {
	return o
}

func (o FunctionResponseOutput) ToFunctionResponseOutputWithContext(ctx context.Context) FunctionResponseOutput {
	return o
}

func (o FunctionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FunctionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FunctionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FunctionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FunctionResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v FunctionResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o FunctionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FunctionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FunctionResponseArrayOutput struct{ *pulumi.OutputState }

func (FunctionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FunctionResponse)(nil)).Elem()
}

func (o FunctionResponseArrayOutput) ToFunctionResponseArrayOutput() FunctionResponseArrayOutput {
	return o
}

func (o FunctionResponseArrayOutput) ToFunctionResponseArrayOutputWithContext(ctx context.Context) FunctionResponseArrayOutput {
	return o
}

func (o FunctionResponseArrayOutput) Index(i pulumi.IntInput) FunctionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FunctionResponse {
		return vs[0].([]FunctionResponse)[vs[1].(int)]
	}).(FunctionResponseOutput)
}

type Identity struct {
	PrincipalId            *string     `pulumi:"principalId"`
	TenantId               *string     `pulumi:"tenantId"`
	Type                   *string     `pulumi:"type"`
	UserAssignedIdentities interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	PrincipalId            pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.Input          `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.AnyOutput {
	return o.ApplyT(func(v Identity) interface{} { return v.UserAssignedIdentities }).(pulumi.AnyOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.AnyOutput {
	return o.ApplyT(func(v *Identity) interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.AnyOutput)
}

type IdentityResponse struct {
	PrincipalId            *string     `pulumi:"principalId"`
	TenantId               *string     `pulumi:"tenantId"`
	Type                   *string     `pulumi:"type"`
	UserAssignedIdentities interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId            pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.Input          `pulumi:"userAssignedIdentities"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() pulumi.AnyOutput {
	return o.ApplyT(func(v IdentityResponse) interface{} { return v.UserAssignedIdentities }).(pulumi.AnyOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() pulumi.AnyOutput {
	return o.ApplyT(func(v *IdentityResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.AnyOutput)
}

type InputType struct {
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}





type InputTypeInput interface {
	pulumi.Input

	ToInputTypeOutput() InputTypeOutput
	ToInputTypeOutputWithContext(context.Context) InputTypeOutput
}

type InputTypeArgs struct {
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
}

func (InputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputType)(nil)).Elem()
}

func (i InputTypeArgs) ToInputTypeOutput() InputTypeOutput {
	return i.ToInputTypeOutputWithContext(context.Background())
}

func (i InputTypeArgs) ToInputTypeOutputWithContext(ctx context.Context) InputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputTypeOutput)
}





type InputTypeArrayInput interface {
	pulumi.Input

	ToInputTypeArrayOutput() InputTypeArrayOutput
	ToInputTypeArrayOutputWithContext(context.Context) InputTypeArrayOutput
}

type InputTypeArray []InputTypeInput

func (InputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputType)(nil)).Elem()
}

func (i InputTypeArray) ToInputTypeArrayOutput() InputTypeArrayOutput {
	return i.ToInputTypeArrayOutputWithContext(context.Background())
}

func (i InputTypeArray) ToInputTypeArrayOutputWithContext(ctx context.Context) InputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputTypeArrayOutput)
}

type InputTypeOutput struct{ *pulumi.OutputState }

func (InputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputType)(nil)).Elem()
}

func (o InputTypeOutput) ToInputTypeOutput() InputTypeOutput {
	return o
}

func (o InputTypeOutput) ToInputTypeOutputWithContext(ctx context.Context) InputTypeOutput {
	return o
}

func (o InputTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InputTypeOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v InputType) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

type InputTypeArrayOutput struct{ *pulumi.OutputState }

func (InputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputType)(nil)).Elem()
}

func (o InputTypeArrayOutput) ToInputTypeArrayOutput() InputTypeArrayOutput {
	return o
}

func (o InputTypeArrayOutput) ToInputTypeArrayOutputWithContext(ctx context.Context) InputTypeArrayOutput {
	return o
}

func (o InputTypeArrayOutput) Index(i pulumi.IntInput) InputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputType {
		return vs[0].([]InputType)[vs[1].(int)]
	}).(InputTypeOutput)
}

type InputResponse struct {
	Id         string      `pulumi:"id"`
	Name       *string     `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	Type       string      `pulumi:"type"`
}





type InputResponseInput interface {
	pulumi.Input

	ToInputResponseOutput() InputResponseOutput
	ToInputResponseOutputWithContext(context.Context) InputResponseOutput
}

type InputResponseArgs struct {
	Id         pulumi.StringInput    `pulumi:"id"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	Properties pulumi.Input          `pulumi:"properties"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (InputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputResponse)(nil)).Elem()
}

func (i InputResponseArgs) ToInputResponseOutput() InputResponseOutput {
	return i.ToInputResponseOutputWithContext(context.Background())
}

func (i InputResponseArgs) ToInputResponseOutputWithContext(ctx context.Context) InputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputResponseOutput)
}





type InputResponseArrayInput interface {
	pulumi.Input

	ToInputResponseArrayOutput() InputResponseArrayOutput
	ToInputResponseArrayOutputWithContext(context.Context) InputResponseArrayOutput
}

type InputResponseArray []InputResponseInput

func (InputResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputResponse)(nil)).Elem()
}

func (i InputResponseArray) ToInputResponseArrayOutput() InputResponseArrayOutput {
	return i.ToInputResponseArrayOutputWithContext(context.Background())
}

func (i InputResponseArray) ToInputResponseArrayOutputWithContext(ctx context.Context) InputResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputResponseArrayOutput)
}

type InputResponseOutput struct{ *pulumi.OutputState }

func (InputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputResponse)(nil)).Elem()
}

func (o InputResponseOutput) ToInputResponseOutput() InputResponseOutput {
	return o
}

func (o InputResponseOutput) ToInputResponseOutputWithContext(ctx context.Context) InputResponseOutput {
	return o
}

func (o InputResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InputResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o InputResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InputResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v InputResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o InputResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InputResponse) string { return v.Type }).(pulumi.StringOutput)
}

type InputResponseArrayOutput struct{ *pulumi.OutputState }

func (InputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputResponse)(nil)).Elem()
}

func (o InputResponseArrayOutput) ToInputResponseArrayOutput() InputResponseArrayOutput {
	return o
}

func (o InputResponseArrayOutput) ToInputResponseArrayOutputWithContext(ctx context.Context) InputResponseArrayOutput {
	return o
}

func (o InputResponseArrayOutput) Index(i pulumi.IntInput) InputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputResponse {
		return vs[0].([]InputResponse)[vs[1].(int)]
	}).(InputResponseOutput)
}

type IoTHubStreamInputDataSource struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	Endpoint               *string `pulumi:"endpoint"`
	IotHubNamespace        *string `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}





type IoTHubStreamInputDataSourceInput interface {
	pulumi.Input

	ToIoTHubStreamInputDataSourceOutput() IoTHubStreamInputDataSourceOutput
	ToIoTHubStreamInputDataSourceOutputWithContext(context.Context) IoTHubStreamInputDataSourceOutput
}

type IoTHubStreamInputDataSourceArgs struct {
	ConsumerGroupName      pulumi.StringPtrInput `pulumi:"consumerGroupName"`
	Endpoint               pulumi.StringPtrInput `pulumi:"endpoint"`
	IotHubNamespace        pulumi.StringPtrInput `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (IoTHubStreamInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTHubStreamInputDataSource)(nil)).Elem()
}

func (i IoTHubStreamInputDataSourceArgs) ToIoTHubStreamInputDataSourceOutput() IoTHubStreamInputDataSourceOutput {
	return i.ToIoTHubStreamInputDataSourceOutputWithContext(context.Background())
}

func (i IoTHubStreamInputDataSourceArgs) ToIoTHubStreamInputDataSourceOutputWithContext(ctx context.Context) IoTHubStreamInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubStreamInputDataSourceOutput)
}

type IoTHubStreamInputDataSourceOutput struct{ *pulumi.OutputState }

func (IoTHubStreamInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTHubStreamInputDataSource)(nil)).Elem()
}

func (o IoTHubStreamInputDataSourceOutput) ToIoTHubStreamInputDataSourceOutput() IoTHubStreamInputDataSourceOutput {
	return o
}

func (o IoTHubStreamInputDataSourceOutput) ToIoTHubStreamInputDataSourceOutputWithContext(ctx context.Context) IoTHubStreamInputDataSourceOutput {
	return o
}

func (o IoTHubStreamInputDataSourceOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSource) *string { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSource) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceOutput) IotHubNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSource) *string { return v.IotHubNamespace }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type IoTHubStreamInputDataSourceResponse struct {
	ConsumerGroupName      *string `pulumi:"consumerGroupName"`
	Endpoint               *string `pulumi:"endpoint"`
	IotHubNamespace        *string `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  *string `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string `pulumi:"sharedAccessPolicyName"`
	Type                   string  `pulumi:"type"`
}





type IoTHubStreamInputDataSourceResponseInput interface {
	pulumi.Input

	ToIoTHubStreamInputDataSourceResponseOutput() IoTHubStreamInputDataSourceResponseOutput
	ToIoTHubStreamInputDataSourceResponseOutputWithContext(context.Context) IoTHubStreamInputDataSourceResponseOutput
}

type IoTHubStreamInputDataSourceResponseArgs struct {
	ConsumerGroupName      pulumi.StringPtrInput `pulumi:"consumerGroupName"`
	Endpoint               pulumi.StringPtrInput `pulumi:"endpoint"`
	IotHubNamespace        pulumi.StringPtrInput `pulumi:"iotHubNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput `pulumi:"sharedAccessPolicyName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (IoTHubStreamInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTHubStreamInputDataSourceResponse)(nil)).Elem()
}

func (i IoTHubStreamInputDataSourceResponseArgs) ToIoTHubStreamInputDataSourceResponseOutput() IoTHubStreamInputDataSourceResponseOutput {
	return i.ToIoTHubStreamInputDataSourceResponseOutputWithContext(context.Background())
}

func (i IoTHubStreamInputDataSourceResponseArgs) ToIoTHubStreamInputDataSourceResponseOutputWithContext(ctx context.Context) IoTHubStreamInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IoTHubStreamInputDataSourceResponseOutput)
}

type IoTHubStreamInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (IoTHubStreamInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IoTHubStreamInputDataSourceResponse)(nil)).Elem()
}

func (o IoTHubStreamInputDataSourceResponseOutput) ToIoTHubStreamInputDataSourceResponseOutput() IoTHubStreamInputDataSourceResponseOutput {
	return o
}

func (o IoTHubStreamInputDataSourceResponseOutput) ToIoTHubStreamInputDataSourceResponseOutputWithContext(ctx context.Context) IoTHubStreamInputDataSourceResponseOutput {
	return o
}

func (o IoTHubStreamInputDataSourceResponseOutput) ConsumerGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSourceResponse) *string { return v.ConsumerGroupName }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceResponseOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSourceResponse) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceResponseOutput) IotHubNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSourceResponse) *string { return v.IotHubNamespace }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o IoTHubStreamInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IoTHubStreamInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JavaScriptFunctionBinding struct {
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}





type JavaScriptFunctionBindingInput interface {
	pulumi.Input

	ToJavaScriptFunctionBindingOutput() JavaScriptFunctionBindingOutput
	ToJavaScriptFunctionBindingOutputWithContext(context.Context) JavaScriptFunctionBindingOutput
}

type JavaScriptFunctionBindingArgs struct {
	Script pulumi.StringPtrInput `pulumi:"script"`
	Type   pulumi.StringInput    `pulumi:"type"`
}

func (JavaScriptFunctionBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JavaScriptFunctionBinding)(nil)).Elem()
}

func (i JavaScriptFunctionBindingArgs) ToJavaScriptFunctionBindingOutput() JavaScriptFunctionBindingOutput {
	return i.ToJavaScriptFunctionBindingOutputWithContext(context.Background())
}

func (i JavaScriptFunctionBindingArgs) ToJavaScriptFunctionBindingOutputWithContext(ctx context.Context) JavaScriptFunctionBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaScriptFunctionBindingOutput)
}

type JavaScriptFunctionBindingOutput struct{ *pulumi.OutputState }

func (JavaScriptFunctionBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JavaScriptFunctionBinding)(nil)).Elem()
}

func (o JavaScriptFunctionBindingOutput) ToJavaScriptFunctionBindingOutput() JavaScriptFunctionBindingOutput {
	return o
}

func (o JavaScriptFunctionBindingOutput) ToJavaScriptFunctionBindingOutputWithContext(ctx context.Context) JavaScriptFunctionBindingOutput {
	return o
}

func (o JavaScriptFunctionBindingOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JavaScriptFunctionBinding) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o JavaScriptFunctionBindingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JavaScriptFunctionBinding) string { return v.Type }).(pulumi.StringOutput)
}

type JavaScriptFunctionBindingResponse struct {
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}





type JavaScriptFunctionBindingResponseInput interface {
	pulumi.Input

	ToJavaScriptFunctionBindingResponseOutput() JavaScriptFunctionBindingResponseOutput
	ToJavaScriptFunctionBindingResponseOutputWithContext(context.Context) JavaScriptFunctionBindingResponseOutput
}

type JavaScriptFunctionBindingResponseArgs struct {
	Script pulumi.StringPtrInput `pulumi:"script"`
	Type   pulumi.StringInput    `pulumi:"type"`
}

func (JavaScriptFunctionBindingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JavaScriptFunctionBindingResponse)(nil)).Elem()
}

func (i JavaScriptFunctionBindingResponseArgs) ToJavaScriptFunctionBindingResponseOutput() JavaScriptFunctionBindingResponseOutput {
	return i.ToJavaScriptFunctionBindingResponseOutputWithContext(context.Background())
}

func (i JavaScriptFunctionBindingResponseArgs) ToJavaScriptFunctionBindingResponseOutputWithContext(ctx context.Context) JavaScriptFunctionBindingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JavaScriptFunctionBindingResponseOutput)
}

type JavaScriptFunctionBindingResponseOutput struct{ *pulumi.OutputState }

func (JavaScriptFunctionBindingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JavaScriptFunctionBindingResponse)(nil)).Elem()
}

func (o JavaScriptFunctionBindingResponseOutput) ToJavaScriptFunctionBindingResponseOutput() JavaScriptFunctionBindingResponseOutput {
	return o
}

func (o JavaScriptFunctionBindingResponseOutput) ToJavaScriptFunctionBindingResponseOutputWithContext(ctx context.Context) JavaScriptFunctionBindingResponseOutput {
	return o
}

func (o JavaScriptFunctionBindingResponseOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JavaScriptFunctionBindingResponse) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o JavaScriptFunctionBindingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JavaScriptFunctionBindingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JobStorageAccount struct {
	AccountKey         *string `pulumi:"accountKey"`
	AccountName        *string `pulumi:"accountName"`
	AuthenticationMode *string `pulumi:"authenticationMode"`
}





type JobStorageAccountInput interface {
	pulumi.Input

	ToJobStorageAccountOutput() JobStorageAccountOutput
	ToJobStorageAccountOutputWithContext(context.Context) JobStorageAccountOutput
}

type JobStorageAccountArgs struct {
	AccountKey         pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName        pulumi.StringPtrInput `pulumi:"accountName"`
	AuthenticationMode pulumi.StringPtrInput `pulumi:"authenticationMode"`
}

func (JobStorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccount)(nil)).Elem()
}

func (i JobStorageAccountArgs) ToJobStorageAccountOutput() JobStorageAccountOutput {
	return i.ToJobStorageAccountOutputWithContext(context.Background())
}

func (i JobStorageAccountArgs) ToJobStorageAccountOutputWithContext(ctx context.Context) JobStorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountOutput)
}

func (i JobStorageAccountArgs) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return i.ToJobStorageAccountPtrOutputWithContext(context.Background())
}

func (i JobStorageAccountArgs) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountOutput).ToJobStorageAccountPtrOutputWithContext(ctx)
}









type JobStorageAccountPtrInput interface {
	pulumi.Input

	ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput
	ToJobStorageAccountPtrOutputWithContext(context.Context) JobStorageAccountPtrOutput
}

type jobStorageAccountPtrType JobStorageAccountArgs

func JobStorageAccountPtr(v *JobStorageAccountArgs) JobStorageAccountPtrInput {
	return (*jobStorageAccountPtrType)(v)
}

func (*jobStorageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccount)(nil)).Elem()
}

func (i *jobStorageAccountPtrType) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return i.ToJobStorageAccountPtrOutputWithContext(context.Background())
}

func (i *jobStorageAccountPtrType) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountPtrOutput)
}

type JobStorageAccountOutput struct{ *pulumi.OutputState }

func (JobStorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccount)(nil)).Elem()
}

func (o JobStorageAccountOutput) ToJobStorageAccountOutput() JobStorageAccountOutput {
	return o
}

func (o JobStorageAccountOutput) ToJobStorageAccountOutputWithContext(ctx context.Context) JobStorageAccountOutput {
	return o
}

func (o JobStorageAccountOutput) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return o.ToJobStorageAccountPtrOutputWithContext(context.Background())
}

func (o JobStorageAccountOutput) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStorageAccount) *JobStorageAccount {
		return &v
	}).(JobStorageAccountPtrOutput)
}

func (o JobStorageAccountOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccount) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccount) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccount) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

type JobStorageAccountPtrOutput struct{ *pulumi.OutputState }

func (JobStorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccount)(nil)).Elem()
}

func (o JobStorageAccountPtrOutput) ToJobStorageAccountPtrOutput() JobStorageAccountPtrOutput {
	return o
}

func (o JobStorageAccountPtrOutput) ToJobStorageAccountPtrOutputWithContext(ctx context.Context) JobStorageAccountPtrOutput {
	return o
}

func (o JobStorageAccountPtrOutput) Elem() JobStorageAccountOutput {
	return o.ApplyT(func(v *JobStorageAccount) JobStorageAccount {
		if v != nil {
			return *v
		}
		var ret JobStorageAccount
		return ret
	}).(JobStorageAccountOutput)
}

func (o JobStorageAccountPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountPtrOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AuthenticationMode
	}).(pulumi.StringPtrOutput)
}

type JobStorageAccountResponse struct {
	AccountKey         *string `pulumi:"accountKey"`
	AccountName        *string `pulumi:"accountName"`
	AuthenticationMode *string `pulumi:"authenticationMode"`
}





type JobStorageAccountResponseInput interface {
	pulumi.Input

	ToJobStorageAccountResponseOutput() JobStorageAccountResponseOutput
	ToJobStorageAccountResponseOutputWithContext(context.Context) JobStorageAccountResponseOutput
}

type JobStorageAccountResponseArgs struct {
	AccountKey         pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName        pulumi.StringPtrInput `pulumi:"accountName"`
	AuthenticationMode pulumi.StringPtrInput `pulumi:"authenticationMode"`
}

func (JobStorageAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccountResponse)(nil)).Elem()
}

func (i JobStorageAccountResponseArgs) ToJobStorageAccountResponseOutput() JobStorageAccountResponseOutput {
	return i.ToJobStorageAccountResponseOutputWithContext(context.Background())
}

func (i JobStorageAccountResponseArgs) ToJobStorageAccountResponseOutputWithContext(ctx context.Context) JobStorageAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountResponseOutput)
}

func (i JobStorageAccountResponseArgs) ToJobStorageAccountResponsePtrOutput() JobStorageAccountResponsePtrOutput {
	return i.ToJobStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (i JobStorageAccountResponseArgs) ToJobStorageAccountResponsePtrOutputWithContext(ctx context.Context) JobStorageAccountResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountResponseOutput).ToJobStorageAccountResponsePtrOutputWithContext(ctx)
}









type JobStorageAccountResponsePtrInput interface {
	pulumi.Input

	ToJobStorageAccountResponsePtrOutput() JobStorageAccountResponsePtrOutput
	ToJobStorageAccountResponsePtrOutputWithContext(context.Context) JobStorageAccountResponsePtrOutput
}

type jobStorageAccountResponsePtrType JobStorageAccountResponseArgs

func JobStorageAccountResponsePtr(v *JobStorageAccountResponseArgs) JobStorageAccountResponsePtrInput {
	return (*jobStorageAccountResponsePtrType)(v)
}

func (*jobStorageAccountResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccountResponse)(nil)).Elem()
}

func (i *jobStorageAccountResponsePtrType) ToJobStorageAccountResponsePtrOutput() JobStorageAccountResponsePtrOutput {
	return i.ToJobStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (i *jobStorageAccountResponsePtrType) ToJobStorageAccountResponsePtrOutputWithContext(ctx context.Context) JobStorageAccountResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStorageAccountResponsePtrOutput)
}

type JobStorageAccountResponseOutput struct{ *pulumi.OutputState }

func (JobStorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStorageAccountResponse)(nil)).Elem()
}

func (o JobStorageAccountResponseOutput) ToJobStorageAccountResponseOutput() JobStorageAccountResponseOutput {
	return o
}

func (o JobStorageAccountResponseOutput) ToJobStorageAccountResponseOutputWithContext(ctx context.Context) JobStorageAccountResponseOutput {
	return o
}

func (o JobStorageAccountResponseOutput) ToJobStorageAccountResponsePtrOutput() JobStorageAccountResponsePtrOutput {
	return o.ToJobStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (o JobStorageAccountResponseOutput) ToJobStorageAccountResponsePtrOutputWithContext(ctx context.Context) JobStorageAccountResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStorageAccountResponse) *JobStorageAccountResponse {
		return &v
	}).(JobStorageAccountResponsePtrOutput)
}

func (o JobStorageAccountResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccountResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccountResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStorageAccountResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

type JobStorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStorageAccountResponse)(nil)).Elem()
}

func (o JobStorageAccountResponsePtrOutput) ToJobStorageAccountResponsePtrOutput() JobStorageAccountResponsePtrOutput {
	return o
}

func (o JobStorageAccountResponsePtrOutput) ToJobStorageAccountResponsePtrOutputWithContext(ctx context.Context) JobStorageAccountResponsePtrOutput {
	return o
}

func (o JobStorageAccountResponsePtrOutput) Elem() JobStorageAccountResponseOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) JobStorageAccountResponse {
		if v != nil {
			return *v
		}
		var ret JobStorageAccountResponse
		return ret
	}).(JobStorageAccountResponseOutput)
}

func (o JobStorageAccountResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o JobStorageAccountResponsePtrOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthenticationMode
	}).(pulumi.StringPtrOutput)
}

type JsonSerialization struct {
	Encoding *string `pulumi:"encoding"`
	Format   *string `pulumi:"format"`
	Type     string  `pulumi:"type"`
}





type JsonSerializationInput interface {
	pulumi.Input

	ToJsonSerializationOutput() JsonSerializationOutput
	ToJsonSerializationOutputWithContext(context.Context) JsonSerializationOutput
}

type JsonSerializationArgs struct {
	Encoding pulumi.StringPtrInput `pulumi:"encoding"`
	Format   pulumi.StringPtrInput `pulumi:"format"`
	Type     pulumi.StringInput    `pulumi:"type"`
}

func (JsonSerializationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonSerialization)(nil)).Elem()
}

func (i JsonSerializationArgs) ToJsonSerializationOutput() JsonSerializationOutput {
	return i.ToJsonSerializationOutputWithContext(context.Background())
}

func (i JsonSerializationArgs) ToJsonSerializationOutputWithContext(ctx context.Context) JsonSerializationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonSerializationOutput)
}

type JsonSerializationOutput struct{ *pulumi.OutputState }

func (JsonSerializationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonSerialization)(nil)).Elem()
}

func (o JsonSerializationOutput) ToJsonSerializationOutput() JsonSerializationOutput {
	return o
}

func (o JsonSerializationOutput) ToJsonSerializationOutputWithContext(ctx context.Context) JsonSerializationOutput {
	return o
}

func (o JsonSerializationOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonSerialization) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o JsonSerializationOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonSerialization) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o JsonSerializationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JsonSerialization) string { return v.Type }).(pulumi.StringOutput)
}

type JsonSerializationResponse struct {
	Encoding *string `pulumi:"encoding"`
	Format   *string `pulumi:"format"`
	Type     string  `pulumi:"type"`
}





type JsonSerializationResponseInput interface {
	pulumi.Input

	ToJsonSerializationResponseOutput() JsonSerializationResponseOutput
	ToJsonSerializationResponseOutputWithContext(context.Context) JsonSerializationResponseOutput
}

type JsonSerializationResponseArgs struct {
	Encoding pulumi.StringPtrInput `pulumi:"encoding"`
	Format   pulumi.StringPtrInput `pulumi:"format"`
	Type     pulumi.StringInput    `pulumi:"type"`
}

func (JsonSerializationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonSerializationResponse)(nil)).Elem()
}

func (i JsonSerializationResponseArgs) ToJsonSerializationResponseOutput() JsonSerializationResponseOutput {
	return i.ToJsonSerializationResponseOutputWithContext(context.Background())
}

func (i JsonSerializationResponseArgs) ToJsonSerializationResponseOutputWithContext(ctx context.Context) JsonSerializationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonSerializationResponseOutput)
}

type JsonSerializationResponseOutput struct{ *pulumi.OutputState }

func (JsonSerializationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonSerializationResponse)(nil)).Elem()
}

func (o JsonSerializationResponseOutput) ToJsonSerializationResponseOutput() JsonSerializationResponseOutput {
	return o
}

func (o JsonSerializationResponseOutput) ToJsonSerializationResponseOutputWithContext(ctx context.Context) JsonSerializationResponseOutput {
	return o
}

func (o JsonSerializationResponseOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonSerializationResponse) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o JsonSerializationResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonSerializationResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o JsonSerializationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JsonSerializationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type OutputType struct {
	Datasource    interface{} `pulumi:"datasource"`
	Name          *string     `pulumi:"name"`
	Serialization interface{} `pulumi:"serialization"`
	SizeWindow    *float64    `pulumi:"sizeWindow"`
	TimeWindow    *string     `pulumi:"timeWindow"`
}





type OutputTypeInput interface {
	pulumi.Input

	ToOutputTypeOutput() OutputTypeOutput
	ToOutputTypeOutputWithContext(context.Context) OutputTypeOutput
}

type OutputTypeArgs struct {
	Datasource    pulumi.Input           `pulumi:"datasource"`
	Name          pulumi.StringPtrInput  `pulumi:"name"`
	Serialization pulumi.Input           `pulumi:"serialization"`
	SizeWindow    pulumi.Float64PtrInput `pulumi:"sizeWindow"`
	TimeWindow    pulumi.StringPtrInput  `pulumi:"timeWindow"`
}

func (OutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (i OutputTypeArgs) ToOutputTypeOutput() OutputTypeOutput {
	return i.ToOutputTypeOutputWithContext(context.Background())
}

func (i OutputTypeArgs) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputTypeOutput)
}





type OutputTypeArrayInput interface {
	pulumi.Input

	ToOutputTypeArrayOutput() OutputTypeArrayOutput
	ToOutputTypeArrayOutputWithContext(context.Context) OutputTypeArrayOutput
}

type OutputTypeArray []OutputTypeInput

func (OutputTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputType)(nil)).Elem()
}

func (i OutputTypeArray) ToOutputTypeArrayOutput() OutputTypeArrayOutput {
	return i.ToOutputTypeArrayOutputWithContext(context.Background())
}

func (i OutputTypeArray) ToOutputTypeArrayOutputWithContext(ctx context.Context) OutputTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputTypeArrayOutput)
}

type OutputTypeOutput struct{ *pulumi.OutputState }

func (OutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputType)(nil)).Elem()
}

func (o OutputTypeOutput) ToOutputTypeOutput() OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) ToOutputTypeOutputWithContext(ctx context.Context) OutputTypeOutput {
	return o
}

func (o OutputTypeOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputType) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputTypeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputType) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputTypeOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputType) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o OutputTypeOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v OutputType) *float64 { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

func (o OutputTypeOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputType) *string { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

type OutputTypeArrayOutput struct{ *pulumi.OutputState }

func (OutputTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputType)(nil)).Elem()
}

func (o OutputTypeArrayOutput) ToOutputTypeArrayOutput() OutputTypeArrayOutput {
	return o
}

func (o OutputTypeArrayOutput) ToOutputTypeArrayOutputWithContext(ctx context.Context) OutputTypeArrayOutput {
	return o
}

func (o OutputTypeArrayOutput) Index(i pulumi.IntInput) OutputTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputType {
		return vs[0].([]OutputType)[vs[1].(int)]
	}).(OutputTypeOutput)
}

type OutputResponse struct {
	Datasource    interface{}         `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse `pulumi:"diagnostics"`
	Etag          string              `pulumi:"etag"`
	Id            string              `pulumi:"id"`
	Name          *string             `pulumi:"name"`
	Serialization interface{}         `pulumi:"serialization"`
	SizeWindow    *float64            `pulumi:"sizeWindow"`
	TimeWindow    *string             `pulumi:"timeWindow"`
	Type          string              `pulumi:"type"`
}





type OutputResponseInput interface {
	pulumi.Input

	ToOutputResponseOutput() OutputResponseOutput
	ToOutputResponseOutputWithContext(context.Context) OutputResponseOutput
}

type OutputResponseArgs struct {
	Datasource    pulumi.Input             `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponseInput `pulumi:"diagnostics"`
	Etag          pulumi.StringInput       `pulumi:"etag"`
	Id            pulumi.StringInput       `pulumi:"id"`
	Name          pulumi.StringPtrInput    `pulumi:"name"`
	Serialization pulumi.Input             `pulumi:"serialization"`
	SizeWindow    pulumi.Float64PtrInput   `pulumi:"sizeWindow"`
	TimeWindow    pulumi.StringPtrInput    `pulumi:"timeWindow"`
	Type          pulumi.StringInput       `pulumi:"type"`
}

func (OutputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputResponse)(nil)).Elem()
}

func (i OutputResponseArgs) ToOutputResponseOutput() OutputResponseOutput {
	return i.ToOutputResponseOutputWithContext(context.Background())
}

func (i OutputResponseArgs) ToOutputResponseOutputWithContext(ctx context.Context) OutputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputResponseOutput)
}





type OutputResponseArrayInput interface {
	pulumi.Input

	ToOutputResponseArrayOutput() OutputResponseArrayOutput
	ToOutputResponseArrayOutputWithContext(context.Context) OutputResponseArrayOutput
}

type OutputResponseArray []OutputResponseInput

func (OutputResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputResponse)(nil)).Elem()
}

func (i OutputResponseArray) ToOutputResponseArrayOutput() OutputResponseArrayOutput {
	return i.ToOutputResponseArrayOutputWithContext(context.Background())
}

func (i OutputResponseArray) ToOutputResponseArrayOutputWithContext(ctx context.Context) OutputResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OutputResponseArrayOutput)
}

type OutputResponseOutput struct{ *pulumi.OutputState }

func (OutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputResponse)(nil)).Elem()
}

func (o OutputResponseOutput) ToOutputResponseOutput() OutputResponseOutput {
	return o
}

func (o OutputResponseOutput) ToOutputResponseOutputWithContext(ctx context.Context) OutputResponseOutput {
	return o
}

func (o OutputResponseOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputResponse) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o OutputResponseOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v OutputResponse) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o OutputResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o OutputResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o OutputResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OutputResponseOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v OutputResponse) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o OutputResponseOutput) SizeWindow() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v OutputResponse) *float64 { return v.SizeWindow }).(pulumi.Float64PtrOutput)
}

func (o OutputResponseOutput) TimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OutputResponse) *string { return v.TimeWindow }).(pulumi.StringPtrOutput)
}

func (o OutputResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v OutputResponse) string { return v.Type }).(pulumi.StringOutput)
}

type OutputResponseArrayOutput struct{ *pulumi.OutputState }

func (OutputResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OutputResponse)(nil)).Elem()
}

func (o OutputResponseArrayOutput) ToOutputResponseArrayOutput() OutputResponseArrayOutput {
	return o
}

func (o OutputResponseArrayOutput) ToOutputResponseArrayOutputWithContext(ctx context.Context) OutputResponseArrayOutput {
	return o
}

func (o OutputResponseArrayOutput) Index(i pulumi.IntInput) OutputResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OutputResponse {
		return vs[0].([]OutputResponse)[vs[1].(int)]
	}).(OutputResponseOutput)
}

type ParquetSerialization struct {
	Type string `pulumi:"type"`
}





type ParquetSerializationInput interface {
	pulumi.Input

	ToParquetSerializationOutput() ParquetSerializationOutput
	ToParquetSerializationOutputWithContext(context.Context) ParquetSerializationOutput
}

type ParquetSerializationArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ParquetSerializationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParquetSerialization)(nil)).Elem()
}

func (i ParquetSerializationArgs) ToParquetSerializationOutput() ParquetSerializationOutput {
	return i.ToParquetSerializationOutputWithContext(context.Background())
}

func (i ParquetSerializationArgs) ToParquetSerializationOutputWithContext(ctx context.Context) ParquetSerializationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParquetSerializationOutput)
}

type ParquetSerializationOutput struct{ *pulumi.OutputState }

func (ParquetSerializationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParquetSerialization)(nil)).Elem()
}

func (o ParquetSerializationOutput) ToParquetSerializationOutput() ParquetSerializationOutput {
	return o
}

func (o ParquetSerializationOutput) ToParquetSerializationOutputWithContext(ctx context.Context) ParquetSerializationOutput {
	return o
}

func (o ParquetSerializationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParquetSerialization) string { return v.Type }).(pulumi.StringOutput)
}

type ParquetSerializationResponse struct {
	Type string `pulumi:"type"`
}





type ParquetSerializationResponseInput interface {
	pulumi.Input

	ToParquetSerializationResponseOutput() ParquetSerializationResponseOutput
	ToParquetSerializationResponseOutputWithContext(context.Context) ParquetSerializationResponseOutput
}

type ParquetSerializationResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ParquetSerializationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParquetSerializationResponse)(nil)).Elem()
}

func (i ParquetSerializationResponseArgs) ToParquetSerializationResponseOutput() ParquetSerializationResponseOutput {
	return i.ToParquetSerializationResponseOutputWithContext(context.Background())
}

func (i ParquetSerializationResponseArgs) ToParquetSerializationResponseOutputWithContext(ctx context.Context) ParquetSerializationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParquetSerializationResponseOutput)
}

type ParquetSerializationResponseOutput struct{ *pulumi.OutputState }

func (ParquetSerializationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParquetSerializationResponse)(nil)).Elem()
}

func (o ParquetSerializationResponseOutput) ToParquetSerializationResponseOutput() ParquetSerializationResponseOutput {
	return o
}

func (o ParquetSerializationResponseOutput) ToParquetSerializationResponseOutputWithContext(ctx context.Context) ParquetSerializationResponseOutput {
	return o
}

func (o ParquetSerializationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ParquetSerializationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PowerBIOutputDataSource struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	Dataset                *string `pulumi:"dataset"`
	GroupId                *string `pulumi:"groupId"`
	GroupName              *string `pulumi:"groupName"`
	RefreshToken           *string `pulumi:"refreshToken"`
	Table                  *string `pulumi:"table"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}





type PowerBIOutputDataSourceInput interface {
	pulumi.Input

	ToPowerBIOutputDataSourceOutput() PowerBIOutputDataSourceOutput
	ToPowerBIOutputDataSourceOutputWithContext(context.Context) PowerBIOutputDataSourceOutput
}

type PowerBIOutputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	Dataset                pulumi.StringPtrInput `pulumi:"dataset"`
	GroupId                pulumi.StringPtrInput `pulumi:"groupId"`
	GroupName              pulumi.StringPtrInput `pulumi:"groupName"`
	RefreshToken           pulumi.StringPtrInput `pulumi:"refreshToken"`
	Table                  pulumi.StringPtrInput `pulumi:"table"`
	TokenUserDisplayName   pulumi.StringPtrInput `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName pulumi.StringPtrInput `pulumi:"tokenUserPrincipalName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (PowerBIOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerBIOutputDataSource)(nil)).Elem()
}

func (i PowerBIOutputDataSourceArgs) ToPowerBIOutputDataSourceOutput() PowerBIOutputDataSourceOutput {
	return i.ToPowerBIOutputDataSourceOutputWithContext(context.Background())
}

func (i PowerBIOutputDataSourceArgs) ToPowerBIOutputDataSourceOutputWithContext(ctx context.Context) PowerBIOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerBIOutputDataSourceOutput)
}

type PowerBIOutputDataSourceOutput struct{ *pulumi.OutputState }

func (PowerBIOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerBIOutputDataSource)(nil)).Elem()
}

func (o PowerBIOutputDataSourceOutput) ToPowerBIOutputDataSourceOutput() PowerBIOutputDataSourceOutput {
	return o
}

func (o PowerBIOutputDataSourceOutput) ToPowerBIOutputDataSourceOutputWithContext(ctx context.Context) PowerBIOutputDataSourceOutput {
	return o
}

func (o PowerBIOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) Dataset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.Dataset }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) TokenUserDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.TokenUserDisplayName }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) TokenUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) *string { return v.TokenUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PowerBIOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type PowerBIOutputDataSourceResponse struct {
	AuthenticationMode     *string `pulumi:"authenticationMode"`
	Dataset                *string `pulumi:"dataset"`
	GroupId                *string `pulumi:"groupId"`
	GroupName              *string `pulumi:"groupName"`
	RefreshToken           *string `pulumi:"refreshToken"`
	Table                  *string `pulumi:"table"`
	TokenUserDisplayName   *string `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName *string `pulumi:"tokenUserPrincipalName"`
	Type                   string  `pulumi:"type"`
}





type PowerBIOutputDataSourceResponseInput interface {
	pulumi.Input

	ToPowerBIOutputDataSourceResponseOutput() PowerBIOutputDataSourceResponseOutput
	ToPowerBIOutputDataSourceResponseOutputWithContext(context.Context) PowerBIOutputDataSourceResponseOutput
}

type PowerBIOutputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput `pulumi:"authenticationMode"`
	Dataset                pulumi.StringPtrInput `pulumi:"dataset"`
	GroupId                pulumi.StringPtrInput `pulumi:"groupId"`
	GroupName              pulumi.StringPtrInput `pulumi:"groupName"`
	RefreshToken           pulumi.StringPtrInput `pulumi:"refreshToken"`
	Table                  pulumi.StringPtrInput `pulumi:"table"`
	TokenUserDisplayName   pulumi.StringPtrInput `pulumi:"tokenUserDisplayName"`
	TokenUserPrincipalName pulumi.StringPtrInput `pulumi:"tokenUserPrincipalName"`
	Type                   pulumi.StringInput    `pulumi:"type"`
}

func (PowerBIOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerBIOutputDataSourceResponse)(nil)).Elem()
}

func (i PowerBIOutputDataSourceResponseArgs) ToPowerBIOutputDataSourceResponseOutput() PowerBIOutputDataSourceResponseOutput {
	return i.ToPowerBIOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i PowerBIOutputDataSourceResponseArgs) ToPowerBIOutputDataSourceResponseOutputWithContext(ctx context.Context) PowerBIOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PowerBIOutputDataSourceResponseOutput)
}

type PowerBIOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (PowerBIOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PowerBIOutputDataSourceResponse)(nil)).Elem()
}

func (o PowerBIOutputDataSourceResponseOutput) ToPowerBIOutputDataSourceResponseOutput() PowerBIOutputDataSourceResponseOutput {
	return o
}

func (o PowerBIOutputDataSourceResponseOutput) ToPowerBIOutputDataSourceResponseOutputWithContext(ctx context.Context) PowerBIOutputDataSourceResponseOutput {
	return o
}

func (o PowerBIOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) Dataset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.Dataset }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) GroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.GroupName }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.Table }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) TokenUserDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.TokenUserDisplayName }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) TokenUserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) *string { return v.TokenUserPrincipalName }).(pulumi.StringPtrOutput)
}

func (o PowerBIOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PowerBIOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RawOutputDatasource struct {
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}





type RawOutputDatasourceInput interface {
	pulumi.Input

	ToRawOutputDatasourceOutput() RawOutputDatasourceOutput
	ToRawOutputDatasourceOutputWithContext(context.Context) RawOutputDatasourceOutput
}

type RawOutputDatasourceArgs struct {
	PayloadUri pulumi.StringPtrInput `pulumi:"payloadUri"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (RawOutputDatasourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RawOutputDatasource)(nil)).Elem()
}

func (i RawOutputDatasourceArgs) ToRawOutputDatasourceOutput() RawOutputDatasourceOutput {
	return i.ToRawOutputDatasourceOutputWithContext(context.Background())
}

func (i RawOutputDatasourceArgs) ToRawOutputDatasourceOutputWithContext(ctx context.Context) RawOutputDatasourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RawOutputDatasourceOutput)
}

type RawOutputDatasourceOutput struct{ *pulumi.OutputState }

func (RawOutputDatasourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RawOutputDatasource)(nil)).Elem()
}

func (o RawOutputDatasourceOutput) ToRawOutputDatasourceOutput() RawOutputDatasourceOutput {
	return o
}

func (o RawOutputDatasourceOutput) ToRawOutputDatasourceOutputWithContext(ctx context.Context) RawOutputDatasourceOutput {
	return o
}

func (o RawOutputDatasourceOutput) PayloadUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawOutputDatasource) *string { return v.PayloadUri }).(pulumi.StringPtrOutput)
}

func (o RawOutputDatasourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RawOutputDatasource) string { return v.Type }).(pulumi.StringOutput)
}

type RawOutputDatasourceResponse struct {
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}





type RawOutputDatasourceResponseInput interface {
	pulumi.Input

	ToRawOutputDatasourceResponseOutput() RawOutputDatasourceResponseOutput
	ToRawOutputDatasourceResponseOutputWithContext(context.Context) RawOutputDatasourceResponseOutput
}

type RawOutputDatasourceResponseArgs struct {
	PayloadUri pulumi.StringPtrInput `pulumi:"payloadUri"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (RawOutputDatasourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RawOutputDatasourceResponse)(nil)).Elem()
}

func (i RawOutputDatasourceResponseArgs) ToRawOutputDatasourceResponseOutput() RawOutputDatasourceResponseOutput {
	return i.ToRawOutputDatasourceResponseOutputWithContext(context.Background())
}

func (i RawOutputDatasourceResponseArgs) ToRawOutputDatasourceResponseOutputWithContext(ctx context.Context) RawOutputDatasourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RawOutputDatasourceResponseOutput)
}

type RawOutputDatasourceResponseOutput struct{ *pulumi.OutputState }

func (RawOutputDatasourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RawOutputDatasourceResponse)(nil)).Elem()
}

func (o RawOutputDatasourceResponseOutput) ToRawOutputDatasourceResponseOutput() RawOutputDatasourceResponseOutput {
	return o
}

func (o RawOutputDatasourceResponseOutput) ToRawOutputDatasourceResponseOutputWithContext(ctx context.Context) RawOutputDatasourceResponseOutput {
	return o
}

func (o RawOutputDatasourceResponseOutput) PayloadUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawOutputDatasourceResponse) *string { return v.PayloadUri }).(pulumi.StringPtrOutput)
}

func (o RawOutputDatasourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RawOutputDatasourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RawReferenceInputDataSource struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}





type RawReferenceInputDataSourceInput interface {
	pulumi.Input

	ToRawReferenceInputDataSourceOutput() RawReferenceInputDataSourceOutput
	ToRawReferenceInputDataSourceOutputWithContext(context.Context) RawReferenceInputDataSourceOutput
}

type RawReferenceInputDataSourceArgs struct {
	Payload    pulumi.StringPtrInput `pulumi:"payload"`
	PayloadUri pulumi.StringPtrInput `pulumi:"payloadUri"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (RawReferenceInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RawReferenceInputDataSource)(nil)).Elem()
}

func (i RawReferenceInputDataSourceArgs) ToRawReferenceInputDataSourceOutput() RawReferenceInputDataSourceOutput {
	return i.ToRawReferenceInputDataSourceOutputWithContext(context.Background())
}

func (i RawReferenceInputDataSourceArgs) ToRawReferenceInputDataSourceOutputWithContext(ctx context.Context) RawReferenceInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RawReferenceInputDataSourceOutput)
}

type RawReferenceInputDataSourceOutput struct{ *pulumi.OutputState }

func (RawReferenceInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RawReferenceInputDataSource)(nil)).Elem()
}

func (o RawReferenceInputDataSourceOutput) ToRawReferenceInputDataSourceOutput() RawReferenceInputDataSourceOutput {
	return o
}

func (o RawReferenceInputDataSourceOutput) ToRawReferenceInputDataSourceOutputWithContext(ctx context.Context) RawReferenceInputDataSourceOutput {
	return o
}

func (o RawReferenceInputDataSourceOutput) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawReferenceInputDataSource) *string { return v.Payload }).(pulumi.StringPtrOutput)
}

func (o RawReferenceInputDataSourceOutput) PayloadUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawReferenceInputDataSource) *string { return v.PayloadUri }).(pulumi.StringPtrOutput)
}

func (o RawReferenceInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RawReferenceInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type RawReferenceInputDataSourceResponse struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}





type RawReferenceInputDataSourceResponseInput interface {
	pulumi.Input

	ToRawReferenceInputDataSourceResponseOutput() RawReferenceInputDataSourceResponseOutput
	ToRawReferenceInputDataSourceResponseOutputWithContext(context.Context) RawReferenceInputDataSourceResponseOutput
}

type RawReferenceInputDataSourceResponseArgs struct {
	Payload    pulumi.StringPtrInput `pulumi:"payload"`
	PayloadUri pulumi.StringPtrInput `pulumi:"payloadUri"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (RawReferenceInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RawReferenceInputDataSourceResponse)(nil)).Elem()
}

func (i RawReferenceInputDataSourceResponseArgs) ToRawReferenceInputDataSourceResponseOutput() RawReferenceInputDataSourceResponseOutput {
	return i.ToRawReferenceInputDataSourceResponseOutputWithContext(context.Background())
}

func (i RawReferenceInputDataSourceResponseArgs) ToRawReferenceInputDataSourceResponseOutputWithContext(ctx context.Context) RawReferenceInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RawReferenceInputDataSourceResponseOutput)
}

type RawReferenceInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RawReferenceInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RawReferenceInputDataSourceResponse)(nil)).Elem()
}

func (o RawReferenceInputDataSourceResponseOutput) ToRawReferenceInputDataSourceResponseOutput() RawReferenceInputDataSourceResponseOutput {
	return o
}

func (o RawReferenceInputDataSourceResponseOutput) ToRawReferenceInputDataSourceResponseOutputWithContext(ctx context.Context) RawReferenceInputDataSourceResponseOutput {
	return o
}

func (o RawReferenceInputDataSourceResponseOutput) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawReferenceInputDataSourceResponse) *string { return v.Payload }).(pulumi.StringPtrOutput)
}

func (o RawReferenceInputDataSourceResponseOutput) PayloadUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawReferenceInputDataSourceResponse) *string { return v.PayloadUri }).(pulumi.StringPtrOutput)
}

func (o RawReferenceInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RawReferenceInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RawStreamInputDataSource struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}





type RawStreamInputDataSourceInput interface {
	pulumi.Input

	ToRawStreamInputDataSourceOutput() RawStreamInputDataSourceOutput
	ToRawStreamInputDataSourceOutputWithContext(context.Context) RawStreamInputDataSourceOutput
}

type RawStreamInputDataSourceArgs struct {
	Payload    pulumi.StringPtrInput `pulumi:"payload"`
	PayloadUri pulumi.StringPtrInput `pulumi:"payloadUri"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (RawStreamInputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RawStreamInputDataSource)(nil)).Elem()
}

func (i RawStreamInputDataSourceArgs) ToRawStreamInputDataSourceOutput() RawStreamInputDataSourceOutput {
	return i.ToRawStreamInputDataSourceOutputWithContext(context.Background())
}

func (i RawStreamInputDataSourceArgs) ToRawStreamInputDataSourceOutputWithContext(ctx context.Context) RawStreamInputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RawStreamInputDataSourceOutput)
}

type RawStreamInputDataSourceOutput struct{ *pulumi.OutputState }

func (RawStreamInputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RawStreamInputDataSource)(nil)).Elem()
}

func (o RawStreamInputDataSourceOutput) ToRawStreamInputDataSourceOutput() RawStreamInputDataSourceOutput {
	return o
}

func (o RawStreamInputDataSourceOutput) ToRawStreamInputDataSourceOutputWithContext(ctx context.Context) RawStreamInputDataSourceOutput {
	return o
}

func (o RawStreamInputDataSourceOutput) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawStreamInputDataSource) *string { return v.Payload }).(pulumi.StringPtrOutput)
}

func (o RawStreamInputDataSourceOutput) PayloadUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawStreamInputDataSource) *string { return v.PayloadUri }).(pulumi.StringPtrOutput)
}

func (o RawStreamInputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RawStreamInputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type RawStreamInputDataSourceResponse struct {
	Payload    *string `pulumi:"payload"`
	PayloadUri *string `pulumi:"payloadUri"`
	Type       string  `pulumi:"type"`
}





type RawStreamInputDataSourceResponseInput interface {
	pulumi.Input

	ToRawStreamInputDataSourceResponseOutput() RawStreamInputDataSourceResponseOutput
	ToRawStreamInputDataSourceResponseOutputWithContext(context.Context) RawStreamInputDataSourceResponseOutput
}

type RawStreamInputDataSourceResponseArgs struct {
	Payload    pulumi.StringPtrInput `pulumi:"payload"`
	PayloadUri pulumi.StringPtrInput `pulumi:"payloadUri"`
	Type       pulumi.StringInput    `pulumi:"type"`
}

func (RawStreamInputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RawStreamInputDataSourceResponse)(nil)).Elem()
}

func (i RawStreamInputDataSourceResponseArgs) ToRawStreamInputDataSourceResponseOutput() RawStreamInputDataSourceResponseOutput {
	return i.ToRawStreamInputDataSourceResponseOutputWithContext(context.Background())
}

func (i RawStreamInputDataSourceResponseArgs) ToRawStreamInputDataSourceResponseOutputWithContext(ctx context.Context) RawStreamInputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RawStreamInputDataSourceResponseOutput)
}

type RawStreamInputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (RawStreamInputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RawStreamInputDataSourceResponse)(nil)).Elem()
}

func (o RawStreamInputDataSourceResponseOutput) ToRawStreamInputDataSourceResponseOutput() RawStreamInputDataSourceResponseOutput {
	return o
}

func (o RawStreamInputDataSourceResponseOutput) ToRawStreamInputDataSourceResponseOutputWithContext(ctx context.Context) RawStreamInputDataSourceResponseOutput {
	return o
}

func (o RawStreamInputDataSourceResponseOutput) Payload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawStreamInputDataSourceResponse) *string { return v.Payload }).(pulumi.StringPtrOutput)
}

func (o RawStreamInputDataSourceResponseOutput) PayloadUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RawStreamInputDataSourceResponse) *string { return v.PayloadUri }).(pulumi.StringPtrOutput)
}

func (o RawStreamInputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RawStreamInputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReferenceInputProperties struct {
	Compression   *Compression `pulumi:"compression"`
	Datasource    interface{}  `pulumi:"datasource"`
	PartitionKey  *string      `pulumi:"partitionKey"`
	Serialization interface{}  `pulumi:"serialization"`
	Type          string       `pulumi:"type"`
}





type ReferenceInputPropertiesInput interface {
	pulumi.Input

	ToReferenceInputPropertiesOutput() ReferenceInputPropertiesOutput
	ToReferenceInputPropertiesOutputWithContext(context.Context) ReferenceInputPropertiesOutput
}

type ReferenceInputPropertiesArgs struct {
	Compression   CompressionPtrInput   `pulumi:"compression"`
	Datasource    pulumi.Input          `pulumi:"datasource"`
	PartitionKey  pulumi.StringPtrInput `pulumi:"partitionKey"`
	Serialization pulumi.Input          `pulumi:"serialization"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ReferenceInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceInputProperties)(nil)).Elem()
}

func (i ReferenceInputPropertiesArgs) ToReferenceInputPropertiesOutput() ReferenceInputPropertiesOutput {
	return i.ToReferenceInputPropertiesOutputWithContext(context.Background())
}

func (i ReferenceInputPropertiesArgs) ToReferenceInputPropertiesOutputWithContext(ctx context.Context) ReferenceInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputPropertiesOutput)
}

type ReferenceInputPropertiesOutput struct{ *pulumi.OutputState }

func (ReferenceInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceInputProperties)(nil)).Elem()
}

func (o ReferenceInputPropertiesOutput) ToReferenceInputPropertiesOutput() ReferenceInputPropertiesOutput {
	return o
}

func (o ReferenceInputPropertiesOutput) ToReferenceInputPropertiesOutputWithContext(ctx context.Context) ReferenceInputPropertiesOutput {
	return o
}

func (o ReferenceInputPropertiesOutput) Compression() CompressionPtrOutput {
	return o.ApplyT(func(v ReferenceInputProperties) *Compression { return v.Compression }).(CompressionPtrOutput)
}

func (o ReferenceInputPropertiesOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v ReferenceInputProperties) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o ReferenceInputPropertiesOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceInputProperties) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o ReferenceInputPropertiesOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v ReferenceInputProperties) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o ReferenceInputPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceInputProperties) string { return v.Type }).(pulumi.StringOutput)
}

type ReferenceInputPropertiesResponse struct {
	Compression   *CompressionResponse `pulumi:"compression"`
	Datasource    interface{}          `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse  `pulumi:"diagnostics"`
	Etag          string               `pulumi:"etag"`
	PartitionKey  *string              `pulumi:"partitionKey"`
	Serialization interface{}          `pulumi:"serialization"`
	Type          string               `pulumi:"type"`
}





type ReferenceInputPropertiesResponseInput interface {
	pulumi.Input

	ToReferenceInputPropertiesResponseOutput() ReferenceInputPropertiesResponseOutput
	ToReferenceInputPropertiesResponseOutputWithContext(context.Context) ReferenceInputPropertiesResponseOutput
}

type ReferenceInputPropertiesResponseArgs struct {
	Compression   CompressionResponsePtrInput `pulumi:"compression"`
	Datasource    pulumi.Input                `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponseInput    `pulumi:"diagnostics"`
	Etag          pulumi.StringInput          `pulumi:"etag"`
	PartitionKey  pulumi.StringPtrInput       `pulumi:"partitionKey"`
	Serialization pulumi.Input                `pulumi:"serialization"`
	Type          pulumi.StringInput          `pulumi:"type"`
}

func (ReferenceInputPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceInputPropertiesResponse)(nil)).Elem()
}

func (i ReferenceInputPropertiesResponseArgs) ToReferenceInputPropertiesResponseOutput() ReferenceInputPropertiesResponseOutput {
	return i.ToReferenceInputPropertiesResponseOutputWithContext(context.Background())
}

func (i ReferenceInputPropertiesResponseArgs) ToReferenceInputPropertiesResponseOutputWithContext(ctx context.Context) ReferenceInputPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceInputPropertiesResponseOutput)
}

type ReferenceInputPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ReferenceInputPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceInputPropertiesResponse)(nil)).Elem()
}

func (o ReferenceInputPropertiesResponseOutput) ToReferenceInputPropertiesResponseOutput() ReferenceInputPropertiesResponseOutput {
	return o
}

func (o ReferenceInputPropertiesResponseOutput) ToReferenceInputPropertiesResponseOutputWithContext(ctx context.Context) ReferenceInputPropertiesResponseOutput {
	return o
}

func (o ReferenceInputPropertiesResponseOutput) Compression() CompressionResponsePtrOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) *CompressionResponse { return v.Compression }).(CompressionResponsePtrOutput)
}

func (o ReferenceInputPropertiesResponseOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o ReferenceInputPropertiesResponseOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o ReferenceInputPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ReferenceInputPropertiesResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o ReferenceInputPropertiesResponseOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o ReferenceInputPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceInputPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ScalarFunctionProperties struct {
	Binding interface{}         `pulumi:"binding"`
	Inputs  []FunctionInputType `pulumi:"inputs"`
	Output  *FunctionOutputType `pulumi:"output"`
	Type    string              `pulumi:"type"`
}





type ScalarFunctionPropertiesInput interface {
	pulumi.Input

	ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput
	ToScalarFunctionPropertiesOutputWithContext(context.Context) ScalarFunctionPropertiesOutput
}

type ScalarFunctionPropertiesArgs struct {
	Binding pulumi.Input                `pulumi:"binding"`
	Inputs  FunctionInputTypeArrayInput `pulumi:"inputs"`
	Output  FunctionOutputTypePtrInput  `pulumi:"output"`
	Type    pulumi.StringInput          `pulumi:"type"`
}

func (ScalarFunctionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionProperties)(nil)).Elem()
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput {
	return i.ToScalarFunctionPropertiesOutputWithContext(context.Background())
}

func (i ScalarFunctionPropertiesArgs) ToScalarFunctionPropertiesOutputWithContext(ctx context.Context) ScalarFunctionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesOutput)
}

type ScalarFunctionPropertiesOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionProperties)(nil)).Elem()
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesOutput() ScalarFunctionPropertiesOutput {
	return o
}

func (o ScalarFunctionPropertiesOutput) ToScalarFunctionPropertiesOutputWithContext(ctx context.Context) ScalarFunctionPropertiesOutput {
	return o
}

func (o ScalarFunctionPropertiesOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

func (o ScalarFunctionPropertiesOutput) Inputs() FunctionInputTypeArrayOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) []FunctionInputType { return v.Inputs }).(FunctionInputTypeArrayOutput)
}

func (o ScalarFunctionPropertiesOutput) Output() FunctionOutputTypePtrOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) *FunctionOutputType { return v.Output }).(FunctionOutputTypePtrOutput)
}

func (o ScalarFunctionPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionProperties) string { return v.Type }).(pulumi.StringOutput)
}

type ScalarFunctionPropertiesResponse struct {
	Binding interface{}             `pulumi:"binding"`
	Etag    string                  `pulumi:"etag"`
	Inputs  []FunctionInputResponse `pulumi:"inputs"`
	Output  *FunctionOutputResponse `pulumi:"output"`
	Type    string                  `pulumi:"type"`
}





type ScalarFunctionPropertiesResponseInput interface {
	pulumi.Input

	ToScalarFunctionPropertiesResponseOutput() ScalarFunctionPropertiesResponseOutput
	ToScalarFunctionPropertiesResponseOutputWithContext(context.Context) ScalarFunctionPropertiesResponseOutput
}

type ScalarFunctionPropertiesResponseArgs struct {
	Binding pulumi.Input                    `pulumi:"binding"`
	Etag    pulumi.StringInput              `pulumi:"etag"`
	Inputs  FunctionInputResponseArrayInput `pulumi:"inputs"`
	Output  FunctionOutputResponsePtrInput  `pulumi:"output"`
	Type    pulumi.StringInput              `pulumi:"type"`
}

func (ScalarFunctionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionPropertiesResponse)(nil)).Elem()
}

func (i ScalarFunctionPropertiesResponseArgs) ToScalarFunctionPropertiesResponseOutput() ScalarFunctionPropertiesResponseOutput {
	return i.ToScalarFunctionPropertiesResponseOutputWithContext(context.Background())
}

func (i ScalarFunctionPropertiesResponseArgs) ToScalarFunctionPropertiesResponseOutputWithContext(ctx context.Context) ScalarFunctionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalarFunctionPropertiesResponseOutput)
}

type ScalarFunctionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ScalarFunctionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalarFunctionPropertiesResponse)(nil)).Elem()
}

func (o ScalarFunctionPropertiesResponseOutput) ToScalarFunctionPropertiesResponseOutput() ScalarFunctionPropertiesResponseOutput {
	return o
}

func (o ScalarFunctionPropertiesResponseOutput) ToScalarFunctionPropertiesResponseOutputWithContext(ctx context.Context) ScalarFunctionPropertiesResponseOutput {
	return o
}

func (o ScalarFunctionPropertiesResponseOutput) Binding() pulumi.AnyOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) interface{} { return v.Binding }).(pulumi.AnyOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Inputs() FunctionInputResponseArrayOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) []FunctionInputResponse { return v.Inputs }).(FunctionInputResponseArrayOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Output() FunctionOutputResponsePtrOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) *FunctionOutputResponse { return v.Output }).(FunctionOutputResponsePtrOutput)
}

func (o ScalarFunctionPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ScalarFunctionPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceBusQueueOutputDataSource struct {
	AuthenticationMode     *string     `pulumi:"authenticationMode"`
	PropertyColumns        []string    `pulumi:"propertyColumns"`
	QueueName              *string     `pulumi:"queueName"`
	ServiceBusNamespace    *string     `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string     `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string     `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  interface{} `pulumi:"systemPropertyColumns"`
	Type                   string      `pulumi:"type"`
}





type ServiceBusQueueOutputDataSourceInput interface {
	pulumi.Input

	ToServiceBusQueueOutputDataSourceOutput() ServiceBusQueueOutputDataSourceOutput
	ToServiceBusQueueOutputDataSourceOutputWithContext(context.Context) ServiceBusQueueOutputDataSourceOutput
}

type ServiceBusQueueOutputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	QueueName              pulumi.StringPtrInput   `pulumi:"queueName"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  pulumi.Input            `pulumi:"systemPropertyColumns"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (ServiceBusQueueOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueOutputDataSource)(nil)).Elem()
}

func (i ServiceBusQueueOutputDataSourceArgs) ToServiceBusQueueOutputDataSourceOutput() ServiceBusQueueOutputDataSourceOutput {
	return i.ToServiceBusQueueOutputDataSourceOutputWithContext(context.Background())
}

func (i ServiceBusQueueOutputDataSourceArgs) ToServiceBusQueueOutputDataSourceOutputWithContext(ctx context.Context) ServiceBusQueueOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueOutputDataSourceOutput)
}

type ServiceBusQueueOutputDataSourceOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueOutputDataSource)(nil)).Elem()
}

func (o ServiceBusQueueOutputDataSourceOutput) ToServiceBusQueueOutputDataSourceOutput() ServiceBusQueueOutputDataSourceOutput {
	return o
}

func (o ServiceBusQueueOutputDataSourceOutput) ToServiceBusQueueOutputDataSourceOutputWithContext(ctx context.Context) ServiceBusQueueOutputDataSourceOutput {
	return o
}

func (o ServiceBusQueueOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) SystemPropertyColumns() pulumi.AnyOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) interface{} { return v.SystemPropertyColumns }).(pulumi.AnyOutput)
}

func (o ServiceBusQueueOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceBusQueueOutputDataSourceResponse struct {
	AuthenticationMode     *string     `pulumi:"authenticationMode"`
	PropertyColumns        []string    `pulumi:"propertyColumns"`
	QueueName              *string     `pulumi:"queueName"`
	ServiceBusNamespace    *string     `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string     `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string     `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  interface{} `pulumi:"systemPropertyColumns"`
	Type                   string      `pulumi:"type"`
}





type ServiceBusQueueOutputDataSourceResponseInput interface {
	pulumi.Input

	ToServiceBusQueueOutputDataSourceResponseOutput() ServiceBusQueueOutputDataSourceResponseOutput
	ToServiceBusQueueOutputDataSourceResponseOutputWithContext(context.Context) ServiceBusQueueOutputDataSourceResponseOutput
}

type ServiceBusQueueOutputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	QueueName              pulumi.StringPtrInput   `pulumi:"queueName"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  pulumi.Input            `pulumi:"systemPropertyColumns"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (ServiceBusQueueOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueOutputDataSourceResponse)(nil)).Elem()
}

func (i ServiceBusQueueOutputDataSourceResponseArgs) ToServiceBusQueueOutputDataSourceResponseOutput() ServiceBusQueueOutputDataSourceResponseOutput {
	return i.ToServiceBusQueueOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i ServiceBusQueueOutputDataSourceResponseArgs) ToServiceBusQueueOutputDataSourceResponseOutputWithContext(ctx context.Context) ServiceBusQueueOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusQueueOutputDataSourceResponseOutput)
}

type ServiceBusQueueOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusQueueOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusQueueOutputDataSourceResponse)(nil)).Elem()
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) ToServiceBusQueueOutputDataSourceResponseOutput() ServiceBusQueueOutputDataSourceResponseOutput {
	return o
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) ToServiceBusQueueOutputDataSourceResponseOutputWithContext(ctx context.Context) ServiceBusQueueOutputDataSourceResponseOutput {
	return o
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) SystemPropertyColumns() pulumi.AnyOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) interface{} { return v.SystemPropertyColumns }).(pulumi.AnyOutput)
}

func (o ServiceBusQueueOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusQueueOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceBusTopicOutputDataSource struct {
	AuthenticationMode     *string           `pulumi:"authenticationMode"`
	PropertyColumns        []string          `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string           `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string           `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string           `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  map[string]string `pulumi:"systemPropertyColumns"`
	TopicName              *string           `pulumi:"topicName"`
	Type                   string            `pulumi:"type"`
}





type ServiceBusTopicOutputDataSourceInput interface {
	pulumi.Input

	ToServiceBusTopicOutputDataSourceOutput() ServiceBusTopicOutputDataSourceOutput
	ToServiceBusTopicOutputDataSourceOutputWithContext(context.Context) ServiceBusTopicOutputDataSourceOutput
}

type ServiceBusTopicOutputDataSourceArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  pulumi.StringMapInput   `pulumi:"systemPropertyColumns"`
	TopicName              pulumi.StringPtrInput   `pulumi:"topicName"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (ServiceBusTopicOutputDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicOutputDataSource)(nil)).Elem()
}

func (i ServiceBusTopicOutputDataSourceArgs) ToServiceBusTopicOutputDataSourceOutput() ServiceBusTopicOutputDataSourceOutput {
	return i.ToServiceBusTopicOutputDataSourceOutputWithContext(context.Background())
}

func (i ServiceBusTopicOutputDataSourceArgs) ToServiceBusTopicOutputDataSourceOutputWithContext(ctx context.Context) ServiceBusTopicOutputDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicOutputDataSourceOutput)
}

type ServiceBusTopicOutputDataSourceOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicOutputDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicOutputDataSource)(nil)).Elem()
}

func (o ServiceBusTopicOutputDataSourceOutput) ToServiceBusTopicOutputDataSourceOutput() ServiceBusTopicOutputDataSourceOutput {
	return o
}

func (o ServiceBusTopicOutputDataSourceOutput) ToServiceBusTopicOutputDataSourceOutputWithContext(ctx context.Context) ServiceBusTopicOutputDataSourceOutput {
	return o
}

func (o ServiceBusTopicOutputDataSourceOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) SystemPropertyColumns() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) map[string]string { return v.SystemPropertyColumns }).(pulumi.StringMapOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) *string { return v.TopicName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSource) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceBusTopicOutputDataSourceResponse struct {
	AuthenticationMode     *string           `pulumi:"authenticationMode"`
	PropertyColumns        []string          `pulumi:"propertyColumns"`
	ServiceBusNamespace    *string           `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  *string           `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName *string           `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  map[string]string `pulumi:"systemPropertyColumns"`
	TopicName              *string           `pulumi:"topicName"`
	Type                   string            `pulumi:"type"`
}





type ServiceBusTopicOutputDataSourceResponseInput interface {
	pulumi.Input

	ToServiceBusTopicOutputDataSourceResponseOutput() ServiceBusTopicOutputDataSourceResponseOutput
	ToServiceBusTopicOutputDataSourceResponseOutputWithContext(context.Context) ServiceBusTopicOutputDataSourceResponseOutput
}

type ServiceBusTopicOutputDataSourceResponseArgs struct {
	AuthenticationMode     pulumi.StringPtrInput   `pulumi:"authenticationMode"`
	PropertyColumns        pulumi.StringArrayInput `pulumi:"propertyColumns"`
	ServiceBusNamespace    pulumi.StringPtrInput   `pulumi:"serviceBusNamespace"`
	SharedAccessPolicyKey  pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyKey"`
	SharedAccessPolicyName pulumi.StringPtrInput   `pulumi:"sharedAccessPolicyName"`
	SystemPropertyColumns  pulumi.StringMapInput   `pulumi:"systemPropertyColumns"`
	TopicName              pulumi.StringPtrInput   `pulumi:"topicName"`
	Type                   pulumi.StringInput      `pulumi:"type"`
}

func (ServiceBusTopicOutputDataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicOutputDataSourceResponse)(nil)).Elem()
}

func (i ServiceBusTopicOutputDataSourceResponseArgs) ToServiceBusTopicOutputDataSourceResponseOutput() ServiceBusTopicOutputDataSourceResponseOutput {
	return i.ToServiceBusTopicOutputDataSourceResponseOutputWithContext(context.Background())
}

func (i ServiceBusTopicOutputDataSourceResponseArgs) ToServiceBusTopicOutputDataSourceResponseOutputWithContext(ctx context.Context) ServiceBusTopicOutputDataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusTopicOutputDataSourceResponseOutput)
}

type ServiceBusTopicOutputDataSourceResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusTopicOutputDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusTopicOutputDataSourceResponse)(nil)).Elem()
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) ToServiceBusTopicOutputDataSourceResponseOutput() ServiceBusTopicOutputDataSourceResponseOutput {
	return o
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) ToServiceBusTopicOutputDataSourceResponseOutputWithContext(ctx context.Context) ServiceBusTopicOutputDataSourceResponseOutput {
	return o
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) AuthenticationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) *string { return v.AuthenticationMode }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) PropertyColumns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) []string { return v.PropertyColumns }).(pulumi.StringArrayOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) ServiceBusNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) *string { return v.ServiceBusNamespace }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) SharedAccessPolicyKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) *string { return v.SharedAccessPolicyKey }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) SharedAccessPolicyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) *string { return v.SharedAccessPolicyName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) SystemPropertyColumns() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) map[string]string { return v.SystemPropertyColumns }).(pulumi.StringMapOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) *string { return v.TopicName }).(pulumi.StringPtrOutput)
}

func (o ServiceBusTopicOutputDataSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusTopicOutputDataSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageAccount struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}





type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

type StorageAccountArgs struct {
	AccountKey  pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

func (i StorageAccountArgs) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput).ToStorageAccountPtrOutputWithContext(ctx)
}









type StorageAccountPtrInput interface {
	pulumi.Input

	ToStorageAccountPtrOutput() StorageAccountPtrOutput
	ToStorageAccountPtrOutputWithContext(context.Context) StorageAccountPtrOutput
}

type storageAccountPtrType StorageAccountArgs

func StorageAccountPtr(v *StorageAccountArgs) StorageAccountPtrInput {
	return (*storageAccountPtrType)(v)
}

func (*storageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPtrOutput)
}





type StorageAccountArrayInput interface {
	pulumi.Input

	ToStorageAccountArrayOutput() StorageAccountArrayOutput
	ToStorageAccountArrayOutputWithContext(context.Context) StorageAccountArrayOutput
}

type StorageAccountArray []StorageAccountInput

func (StorageAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (i StorageAccountArray) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return i.ToStorageAccountArrayOutputWithContext(context.Background())
}

func (i StorageAccountArray) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountArrayOutput)
}

type StorageAccountOutput struct{ *pulumi.OutputState }

func (StorageAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (o StorageAccountOutput) ToStorageAccountOutput() StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return o
}

func (o StorageAccountOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (o StorageAccountOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccount) *StorageAccount {
		return &v
	}).(StorageAccountPtrOutput)
}

func (o StorageAccountOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o StorageAccountOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccount) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

type StorageAccountPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return o
}

func (o StorageAccountPtrOutput) Elem() StorageAccountOutput {
	return o.ApplyT(func(v *StorageAccount) StorageAccount {
		if v != nil {
			return *v
		}
		var ret StorageAccount
		return ret
	}).(StorageAccountOutput)
}

func (o StorageAccountPtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountPtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccount) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

type StorageAccountArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccount)(nil)).Elem()
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutput() StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) ToStorageAccountArrayOutputWithContext(ctx context.Context) StorageAccountArrayOutput {
	return o
}

func (o StorageAccountArrayOutput) Index(i pulumi.IntInput) StorageAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccount {
		return vs[0].([]StorageAccount)[vs[1].(int)]
	}).(StorageAccountOutput)
}

type StorageAccountResponse struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName *string `pulumi:"accountName"`
}





type StorageAccountResponseInput interface {
	pulumi.Input

	ToStorageAccountResponseOutput() StorageAccountResponseOutput
	ToStorageAccountResponseOutputWithContext(context.Context) StorageAccountResponseOutput
}

type StorageAccountResponseArgs struct {
	AccountKey  pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
}

func (StorageAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return i.ToStorageAccountResponseOutputWithContext(context.Background())
}

func (i StorageAccountResponseArgs) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseOutput)
}

func (i StorageAccountResponseArgs) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return i.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountResponseArgs) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseOutput).ToStorageAccountResponsePtrOutputWithContext(ctx)
}









type StorageAccountResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput
	ToStorageAccountResponsePtrOutputWithContext(context.Context) StorageAccountResponsePtrOutput
}

type storageAccountResponsePtrType StorageAccountResponseArgs

func StorageAccountResponsePtr(v *StorageAccountResponseArgs) StorageAccountResponsePtrInput {
	return (*storageAccountResponsePtrType)(v)
}

func (*storageAccountResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountResponse)(nil)).Elem()
}

func (i *storageAccountResponsePtrType) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return i.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountResponsePtrType) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponsePtrOutput)
}





type StorageAccountResponseArrayInput interface {
	pulumi.Input

	ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput
	ToStorageAccountResponseArrayOutputWithContext(context.Context) StorageAccountResponseArrayOutput
}

type StorageAccountResponseArray []StorageAccountResponseInput

func (StorageAccountResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (i StorageAccountResponseArray) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return i.ToStorageAccountResponseArrayOutputWithContext(context.Background())
}

func (i StorageAccountResponseArray) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountResponseArrayOutput)
}

type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountResponseOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountResponse) *StorageAccountResponse {
		return &v
	}).(StorageAccountResponsePtrOutput)
}

func (o StorageAccountResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

type StorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) Elem() StorageAccountResponseOutput {
	return o.ApplyT(func(v *StorageAccountResponse) StorageAccountResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountResponse
		return ret
	}).(StorageAccountResponseOutput)
}

func (o StorageAccountResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o StorageAccountResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountName
	}).(pulumi.StringPtrOutput)
}

type StorageAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutput() StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) ToStorageAccountResponseArrayOutputWithContext(ctx context.Context) StorageAccountResponseArrayOutput {
	return o
}

func (o StorageAccountResponseArrayOutput) Index(i pulumi.IntInput) StorageAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageAccountResponse {
		return vs[0].([]StorageAccountResponse)[vs[1].(int)]
	}).(StorageAccountResponseOutput)
}

type StreamInputProperties struct {
	Compression   *Compression `pulumi:"compression"`
	Datasource    interface{}  `pulumi:"datasource"`
	PartitionKey  *string      `pulumi:"partitionKey"`
	Serialization interface{}  `pulumi:"serialization"`
	Type          string       `pulumi:"type"`
}





type StreamInputPropertiesInput interface {
	pulumi.Input

	ToStreamInputPropertiesOutput() StreamInputPropertiesOutput
	ToStreamInputPropertiesOutputWithContext(context.Context) StreamInputPropertiesOutput
}

type StreamInputPropertiesArgs struct {
	Compression   CompressionPtrInput   `pulumi:"compression"`
	Datasource    pulumi.Input          `pulumi:"datasource"`
	PartitionKey  pulumi.StringPtrInput `pulumi:"partitionKey"`
	Serialization pulumi.Input          `pulumi:"serialization"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (StreamInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputProperties)(nil)).Elem()
}

func (i StreamInputPropertiesArgs) ToStreamInputPropertiesOutput() StreamInputPropertiesOutput {
	return i.ToStreamInputPropertiesOutputWithContext(context.Background())
}

func (i StreamInputPropertiesArgs) ToStreamInputPropertiesOutputWithContext(ctx context.Context) StreamInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputPropertiesOutput)
}

type StreamInputPropertiesOutput struct{ *pulumi.OutputState }

func (StreamInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputProperties)(nil)).Elem()
}

func (o StreamInputPropertiesOutput) ToStreamInputPropertiesOutput() StreamInputPropertiesOutput {
	return o
}

func (o StreamInputPropertiesOutput) ToStreamInputPropertiesOutputWithContext(ctx context.Context) StreamInputPropertiesOutput {
	return o
}

func (o StreamInputPropertiesOutput) Compression() CompressionPtrOutput {
	return o.ApplyT(func(v StreamInputProperties) *Compression { return v.Compression }).(CompressionPtrOutput)
}

func (o StreamInputPropertiesOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v StreamInputProperties) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o StreamInputPropertiesOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamInputProperties) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o StreamInputPropertiesOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v StreamInputProperties) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o StreamInputPropertiesOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StreamInputProperties) string { return v.Type }).(pulumi.StringOutput)
}

type StreamInputPropertiesResponse struct {
	Compression   *CompressionResponse `pulumi:"compression"`
	Datasource    interface{}          `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponse  `pulumi:"diagnostics"`
	Etag          string               `pulumi:"etag"`
	PartitionKey  *string              `pulumi:"partitionKey"`
	Serialization interface{}          `pulumi:"serialization"`
	Type          string               `pulumi:"type"`
}





type StreamInputPropertiesResponseInput interface {
	pulumi.Input

	ToStreamInputPropertiesResponseOutput() StreamInputPropertiesResponseOutput
	ToStreamInputPropertiesResponseOutputWithContext(context.Context) StreamInputPropertiesResponseOutput
}

type StreamInputPropertiesResponseArgs struct {
	Compression   CompressionResponsePtrInput `pulumi:"compression"`
	Datasource    pulumi.Input                `pulumi:"datasource"`
	Diagnostics   DiagnosticsResponseInput    `pulumi:"diagnostics"`
	Etag          pulumi.StringInput          `pulumi:"etag"`
	PartitionKey  pulumi.StringPtrInput       `pulumi:"partitionKey"`
	Serialization pulumi.Input                `pulumi:"serialization"`
	Type          pulumi.StringInput          `pulumi:"type"`
}

func (StreamInputPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputPropertiesResponse)(nil)).Elem()
}

func (i StreamInputPropertiesResponseArgs) ToStreamInputPropertiesResponseOutput() StreamInputPropertiesResponseOutput {
	return i.ToStreamInputPropertiesResponseOutputWithContext(context.Background())
}

func (i StreamInputPropertiesResponseArgs) ToStreamInputPropertiesResponseOutputWithContext(ctx context.Context) StreamInputPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamInputPropertiesResponseOutput)
}

type StreamInputPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StreamInputPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StreamInputPropertiesResponse)(nil)).Elem()
}

func (o StreamInputPropertiesResponseOutput) ToStreamInputPropertiesResponseOutput() StreamInputPropertiesResponseOutput {
	return o
}

func (o StreamInputPropertiesResponseOutput) ToStreamInputPropertiesResponseOutputWithContext(ctx context.Context) StreamInputPropertiesResponseOutput {
	return o
}

func (o StreamInputPropertiesResponseOutput) Compression() CompressionResponsePtrOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) *CompressionResponse { return v.Compression }).(CompressionResponsePtrOutput)
}

func (o StreamInputPropertiesResponseOutput) Datasource() pulumi.AnyOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) interface{} { return v.Datasource }).(pulumi.AnyOutput)
}

func (o StreamInputPropertiesResponseOutput) Diagnostics() DiagnosticsResponseOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) DiagnosticsResponse { return v.Diagnostics }).(DiagnosticsResponseOutput)
}

func (o StreamInputPropertiesResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o StreamInputPropertiesResponseOutput) PartitionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) *string { return v.PartitionKey }).(pulumi.StringPtrOutput)
}

func (o StreamInputPropertiesResponseOutput) Serialization() pulumi.AnyOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) interface{} { return v.Serialization }).(pulumi.AnyOutput)
}

func (o StreamInputPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StreamInputPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type Transformation struct {
	Name                *string `pulumi:"name"`
	Query               *string `pulumi:"query"`
	StreamingUnits      *int    `pulumi:"streamingUnits"`
	ValidStreamingUnits []int   `pulumi:"validStreamingUnits"`
}





type TransformationInput interface {
	pulumi.Input

	ToTransformationOutput() TransformationOutput
	ToTransformationOutputWithContext(context.Context) TransformationOutput
}

type TransformationArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	Query               pulumi.StringPtrInput `pulumi:"query"`
	StreamingUnits      pulumi.IntPtrInput    `pulumi:"streamingUnits"`
	ValidStreamingUnits pulumi.IntArrayInput  `pulumi:"validStreamingUnits"`
}

func (TransformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil)).Elem()
}

func (i TransformationArgs) ToTransformationOutput() TransformationOutput {
	return i.ToTransformationOutputWithContext(context.Background())
}

func (i TransformationArgs) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput)
}

func (i TransformationArgs) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i TransformationArgs) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput).ToTransformationPtrOutputWithContext(ctx)
}









type TransformationPtrInput interface {
	pulumi.Input

	ToTransformationPtrOutput() TransformationPtrOutput
	ToTransformationPtrOutputWithContext(context.Context) TransformationPtrOutput
}

type transformationPtrType TransformationArgs

func TransformationPtr(v *TransformationArgs) TransformationPtrInput {
	return (*transformationPtrType)(v)
}

func (*transformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (i *transformationPtrType) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i *transformationPtrType) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationPtrOutput)
}

type TransformationOutput struct{ *pulumi.OutputState }

func (TransformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil)).Elem()
}

func (o TransformationOutput) ToTransformationOutput() TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o.ToTransformationPtrOutputWithContext(context.Background())
}

func (o TransformationOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Transformation) *Transformation {
		return &v
	}).(TransformationPtrOutput)
}

func (o TransformationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Transformation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TransformationOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Transformation) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o TransformationOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Transformation) *int { return v.StreamingUnits }).(pulumi.IntPtrOutput)
}

func (o TransformationOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v Transformation) []int { return v.ValidStreamingUnits }).(pulumi.IntArrayOutput)
}

type TransformationPtrOutput struct{ *pulumi.OutputState }

func (TransformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil)).Elem()
}

func (o TransformationPtrOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) Elem() TransformationOutput {
	return o.ApplyT(func(v *Transformation) Transformation {
		if v != nil {
			return *v
		}
		var ret Transformation
		return ret
	}).(TransformationOutput)
}

func (o TransformationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TransformationPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformation) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o TransformationPtrOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Transformation) *int {
		if v == nil {
			return nil
		}
		return v.StreamingUnits
	}).(pulumi.IntPtrOutput)
}

func (o TransformationPtrOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *Transformation) []int {
		if v == nil {
			return nil
		}
		return v.ValidStreamingUnits
	}).(pulumi.IntArrayOutput)
}

type TransformationResponse struct {
	Etag                string  `pulumi:"etag"`
	Id                  string  `pulumi:"id"`
	Name                *string `pulumi:"name"`
	Query               *string `pulumi:"query"`
	StreamingUnits      *int    `pulumi:"streamingUnits"`
	Type                string  `pulumi:"type"`
	ValidStreamingUnits []int   `pulumi:"validStreamingUnits"`
}





type TransformationResponseInput interface {
	pulumi.Input

	ToTransformationResponseOutput() TransformationResponseOutput
	ToTransformationResponseOutputWithContext(context.Context) TransformationResponseOutput
}

type TransformationResponseArgs struct {
	Etag                pulumi.StringInput    `pulumi:"etag"`
	Id                  pulumi.StringInput    `pulumi:"id"`
	Name                pulumi.StringPtrInput `pulumi:"name"`
	Query               pulumi.StringPtrInput `pulumi:"query"`
	StreamingUnits      pulumi.IntPtrInput    `pulumi:"streamingUnits"`
	Type                pulumi.StringInput    `pulumi:"type"`
	ValidStreamingUnits pulumi.IntArrayInput  `pulumi:"validStreamingUnits"`
}

func (TransformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformationResponse)(nil)).Elem()
}

func (i TransformationResponseArgs) ToTransformationResponseOutput() TransformationResponseOutput {
	return i.ToTransformationResponseOutputWithContext(context.Background())
}

func (i TransformationResponseArgs) ToTransformationResponseOutputWithContext(ctx context.Context) TransformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationResponseOutput)
}

func (i TransformationResponseArgs) ToTransformationResponsePtrOutput() TransformationResponsePtrOutput {
	return i.ToTransformationResponsePtrOutputWithContext(context.Background())
}

func (i TransformationResponseArgs) ToTransformationResponsePtrOutputWithContext(ctx context.Context) TransformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationResponseOutput).ToTransformationResponsePtrOutputWithContext(ctx)
}









type TransformationResponsePtrInput interface {
	pulumi.Input

	ToTransformationResponsePtrOutput() TransformationResponsePtrOutput
	ToTransformationResponsePtrOutputWithContext(context.Context) TransformationResponsePtrOutput
}

type transformationResponsePtrType TransformationResponseArgs

func TransformationResponsePtr(v *TransformationResponseArgs) TransformationResponsePtrInput {
	return (*transformationResponsePtrType)(v)
}

func (*transformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TransformationResponse)(nil)).Elem()
}

func (i *transformationResponsePtrType) ToTransformationResponsePtrOutput() TransformationResponsePtrOutput {
	return i.ToTransformationResponsePtrOutputWithContext(context.Background())
}

func (i *transformationResponsePtrType) ToTransformationResponsePtrOutputWithContext(ctx context.Context) TransformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationResponsePtrOutput)
}

type TransformationResponseOutput struct{ *pulumi.OutputState }

func (TransformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransformationResponse)(nil)).Elem()
}

func (o TransformationResponseOutput) ToTransformationResponseOutput() TransformationResponseOutput {
	return o
}

func (o TransformationResponseOutput) ToTransformationResponseOutputWithContext(ctx context.Context) TransformationResponseOutput {
	return o
}

func (o TransformationResponseOutput) ToTransformationResponsePtrOutput() TransformationResponsePtrOutput {
	return o.ToTransformationResponsePtrOutputWithContext(context.Background())
}

func (o TransformationResponseOutput) ToTransformationResponsePtrOutputWithContext(ctx context.Context) TransformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TransformationResponse) *TransformationResponse {
		return &v
	}).(TransformationResponsePtrOutput)
}

func (o TransformationResponseOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Etag }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TransformationResponseOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *string { return v.Query }).(pulumi.StringPtrOutput)
}

func (o TransformationResponseOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TransformationResponse) *int { return v.StreamingUnits }).(pulumi.IntPtrOutput)
}

func (o TransformationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TransformationResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o TransformationResponseOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v TransformationResponse) []int { return v.ValidStreamingUnits }).(pulumi.IntArrayOutput)
}

type TransformationResponsePtrOutput struct{ *pulumi.OutputState }

func (TransformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransformationResponse)(nil)).Elem()
}

func (o TransformationResponsePtrOutput) ToTransformationResponsePtrOutput() TransformationResponsePtrOutput {
	return o
}

func (o TransformationResponsePtrOutput) ToTransformationResponsePtrOutputWithContext(ctx context.Context) TransformationResponsePtrOutput {
	return o
}

func (o TransformationResponsePtrOutput) Elem() TransformationResponseOutput {
	return o.ApplyT(func(v *TransformationResponse) TransformationResponse {
		if v != nil {
			return *v
		}
		var ret TransformationResponse
		return ret
	}).(TransformationResponseOutput)
}

func (o TransformationResponsePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) StreamingUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *int {
		if v == nil {
			return nil
		}
		return v.StreamingUnits
	}).(pulumi.IntPtrOutput)
}

func (o TransformationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o TransformationResponsePtrOutput) ValidStreamingUnits() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *TransformationResponse) []int {
		if v == nil {
			return nil
		}
		return v.ValidStreamingUnits
	}).(pulumi.IntArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AggregateFunctionPropertiesOutput{})
	pulumi.RegisterOutputType(AggregateFunctionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AvroSerializationOutput{})
	pulumi.RegisterOutputType(AvroSerializationResponseOutput{})
	pulumi.RegisterOutputType(AzureDataLakeStoreOutputDataSourceOutput{})
	pulumi.RegisterOutputType(AzureDataLakeStoreOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(AzureFunctionOutputDataSourceOutput{})
	pulumi.RegisterOutputType(AzureFunctionOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceFunctionBindingOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceFunctionBindingResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceInputColumnOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceInputColumnArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceInputColumnResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceInputColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceOutputColumnOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceOutputColumnArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceOutputColumnResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningServiceOutputColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioFunctionBindingOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioFunctionBindingResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputColumnOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputColumnArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputColumnResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputsOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputsPtrOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputsResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioInputsResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioOutputColumnOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioOutputColumnArrayOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioOutputColumnResponseOutput{})
	pulumi.RegisterOutputType(AzureMachineLearningStudioOutputColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureSqlDatabaseOutputDataSourceOutput{})
	pulumi.RegisterOutputType(AzureSqlDatabaseOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlReferenceInputDataSourceOutput{})
	pulumi.RegisterOutputType(AzureSqlReferenceInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(AzureSynapseOutputDataSourceOutput{})
	pulumi.RegisterOutputType(AzureSynapseOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(AzureTableOutputDataSourceOutput{})
	pulumi.RegisterOutputType(AzureTableOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(BlobOutputDataSourceOutput{})
	pulumi.RegisterOutputType(BlobOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(BlobReferenceInputDataSourceOutput{})
	pulumi.RegisterOutputType(BlobReferenceInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(BlobStreamInputDataSourceOutput{})
	pulumi.RegisterOutputType(BlobStreamInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(CSharpFunctionBindingOutput{})
	pulumi.RegisterOutputType(CSharpFunctionBindingResponseOutput{})
	pulumi.RegisterOutputType(ClusterInfoOutput{})
	pulumi.RegisterOutputType(ClusterInfoPtrOutput{})
	pulumi.RegisterOutputType(ClusterInfoResponseOutput{})
	pulumi.RegisterOutputType(ClusterInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(CompressionOutput{})
	pulumi.RegisterOutputType(CompressionPtrOutput{})
	pulumi.RegisterOutputType(CompressionResponseOutput{})
	pulumi.RegisterOutputType(CompressionResponsePtrOutput{})
	pulumi.RegisterOutputType(CsvSerializationOutput{})
	pulumi.RegisterOutputType(CsvSerializationResponseOutput{})
	pulumi.RegisterOutputType(CustomClrSerializationOutput{})
	pulumi.RegisterOutputType(CustomClrSerializationResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponsePtrOutput{})
	pulumi.RegisterOutputType(DocumentDbOutputDataSourceOutput{})
	pulumi.RegisterOutputType(DocumentDbOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(EventHubOutputDataSourceOutput{})
	pulumi.RegisterOutputType(EventHubOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(EventHubStreamInputDataSourceOutput{})
	pulumi.RegisterOutputType(EventHubStreamInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(EventHubV2OutputDataSourceOutput{})
	pulumi.RegisterOutputType(EventHubV2OutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(EventHubV2StreamInputDataSourceOutput{})
	pulumi.RegisterOutputType(EventHubV2StreamInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(ExternalOutput{})
	pulumi.RegisterOutputType(ExternalPtrOutput{})
	pulumi.RegisterOutputType(ExternalResponseOutput{})
	pulumi.RegisterOutputType(ExternalResponsePtrOutput{})
	pulumi.RegisterOutputType(FunctionTypeOutput{})
	pulumi.RegisterOutputType(FunctionTypeArrayOutput{})
	pulumi.RegisterOutputType(FunctionInputTypeOutput{})
	pulumi.RegisterOutputType(FunctionInputTypeArrayOutput{})
	pulumi.RegisterOutputType(FunctionInputResponseOutput{})
	pulumi.RegisterOutputType(FunctionInputResponseArrayOutput{})
	pulumi.RegisterOutputType(FunctionOutputTypeOutput{})
	pulumi.RegisterOutputType(FunctionOutputTypePtrOutput{})
	pulumi.RegisterOutputType(FunctionOutputResponseOutput{})
	pulumi.RegisterOutputType(FunctionOutputResponsePtrOutput{})
	pulumi.RegisterOutputType(FunctionResponseOutput{})
	pulumi.RegisterOutputType(FunctionResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(InputTypeOutput{})
	pulumi.RegisterOutputType(InputTypeArrayOutput{})
	pulumi.RegisterOutputType(InputResponseOutput{})
	pulumi.RegisterOutputType(InputResponseArrayOutput{})
	pulumi.RegisterOutputType(IoTHubStreamInputDataSourceOutput{})
	pulumi.RegisterOutputType(IoTHubStreamInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(JavaScriptFunctionBindingOutput{})
	pulumi.RegisterOutputType(JavaScriptFunctionBindingResponseOutput{})
	pulumi.RegisterOutputType(JobStorageAccountOutput{})
	pulumi.RegisterOutputType(JobStorageAccountPtrOutput{})
	pulumi.RegisterOutputType(JobStorageAccountResponseOutput{})
	pulumi.RegisterOutputType(JobStorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(JsonSerializationOutput{})
	pulumi.RegisterOutputType(JsonSerializationResponseOutput{})
	pulumi.RegisterOutputType(OutputTypeOutput{})
	pulumi.RegisterOutputType(OutputTypeArrayOutput{})
	pulumi.RegisterOutputType(OutputResponseOutput{})
	pulumi.RegisterOutputType(OutputResponseArrayOutput{})
	pulumi.RegisterOutputType(ParquetSerializationOutput{})
	pulumi.RegisterOutputType(ParquetSerializationResponseOutput{})
	pulumi.RegisterOutputType(PowerBIOutputDataSourceOutput{})
	pulumi.RegisterOutputType(PowerBIOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RawOutputDatasourceOutput{})
	pulumi.RegisterOutputType(RawOutputDatasourceResponseOutput{})
	pulumi.RegisterOutputType(RawReferenceInputDataSourceOutput{})
	pulumi.RegisterOutputType(RawReferenceInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(RawStreamInputDataSourceOutput{})
	pulumi.RegisterOutputType(RawStreamInputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(ReferenceInputPropertiesOutput{})
	pulumi.RegisterOutputType(ReferenceInputPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesOutput{})
	pulumi.RegisterOutputType(ScalarFunctionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueOutputDataSourceOutput{})
	pulumi.RegisterOutputType(ServiceBusQueueOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicOutputDataSourceOutput{})
	pulumi.RegisterOutputType(ServiceBusTopicOutputDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountArrayOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(StreamInputPropertiesOutput{})
	pulumi.RegisterOutputType(StreamInputPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TransformationOutput{})
	pulumi.RegisterOutputType(TransformationPtrOutput{})
	pulumi.RegisterOutputType(TransformationResponseOutput{})
	pulumi.RegisterOutputType(TransformationResponsePtrOutput{})
}
