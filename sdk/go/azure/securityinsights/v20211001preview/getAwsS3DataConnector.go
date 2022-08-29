


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAwsS3DataConnector(ctx *pulumi.Context, args *LookupAwsS3DataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAwsS3DataConnectorResult, error) {
	var rv LookupAwsS3DataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20211001preview:getAwsS3DataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAwsS3DataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAwsS3DataConnectorResult struct {
	DataTypes        AwsS3DataConnectorDataTypesResponse `pulumi:"dataTypes"`
	DestinationTable string                              `pulumi:"destinationTable"`
	Etag             *string                             `pulumi:"etag"`
	Id               string                              `pulumi:"id"`
	Kind             string                              `pulumi:"kind"`
	Name             string                              `pulumi:"name"`
	RoleArn          string                              `pulumi:"roleArn"`
	SqsUrls          []string                            `pulumi:"sqsUrls"`
	SystemData       SystemDataResponse                  `pulumi:"systemData"`
	Type             string                              `pulumi:"type"`
}

func LookupAwsS3DataConnectorOutput(ctx *pulumi.Context, args LookupAwsS3DataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAwsS3DataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAwsS3DataConnectorResult, error) {
			args := v.(LookupAwsS3DataConnectorArgs)
			r, err := LookupAwsS3DataConnector(ctx, &args, opts...)
			var s LookupAwsS3DataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAwsS3DataConnectorResultOutput)
}

type LookupAwsS3DataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAwsS3DataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsS3DataConnectorArgs)(nil)).Elem()
}


type LookupAwsS3DataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAwsS3DataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsS3DataConnectorResult)(nil)).Elem()
}

func (o LookupAwsS3DataConnectorResultOutput) ToLookupAwsS3DataConnectorResultOutput() LookupAwsS3DataConnectorResultOutput {
	return o
}

func (o LookupAwsS3DataConnectorResultOutput) ToLookupAwsS3DataConnectorResultOutputWithContext(ctx context.Context) LookupAwsS3DataConnectorResultOutput {
	return o
}

func (o LookupAwsS3DataConnectorResultOutput) DataTypes() AwsS3DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) AwsS3DataConnectorDataTypesResponse { return v.DataTypes }).(AwsS3DataConnectorDataTypesResponseOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) DestinationTable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.DestinationTable }).(pulumi.StringOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.RoleArn }).(pulumi.StringOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) SqsUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) []string { return v.SqsUrls }).(pulumi.StringArrayOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAwsS3DataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsS3DataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsS3DataConnectorResultOutput{})
}
