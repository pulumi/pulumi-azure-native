


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAwsCloudTrailDataConnector(ctx *pulumi.Context, args *LookupAwsCloudTrailDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupAwsCloudTrailDataConnectorResult, error) {
	var rv LookupAwsCloudTrailDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20220101preview:getAwsCloudTrailDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAwsCloudTrailDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupAwsCloudTrailDataConnectorResult struct {
	AwsRoleArn *string                                     `pulumi:"awsRoleArn"`
	DataTypes  AwsCloudTrailDataConnectorDataTypesResponse `pulumi:"dataTypes"`
	Etag       *string                                     `pulumi:"etag"`
	Id         string                                      `pulumi:"id"`
	Kind       string                                      `pulumi:"kind"`
	Name       string                                      `pulumi:"name"`
	SystemData SystemDataResponse                          `pulumi:"systemData"`
	Type       string                                      `pulumi:"type"`
}

func LookupAwsCloudTrailDataConnectorOutput(ctx *pulumi.Context, args LookupAwsCloudTrailDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupAwsCloudTrailDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAwsCloudTrailDataConnectorResult, error) {
			args := v.(LookupAwsCloudTrailDataConnectorArgs)
			r, err := LookupAwsCloudTrailDataConnector(ctx, &args, opts...)
			var s LookupAwsCloudTrailDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAwsCloudTrailDataConnectorResultOutput)
}

type LookupAwsCloudTrailDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupAwsCloudTrailDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsCloudTrailDataConnectorArgs)(nil)).Elem()
}


type LookupAwsCloudTrailDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupAwsCloudTrailDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAwsCloudTrailDataConnectorResult)(nil)).Elem()
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) ToLookupAwsCloudTrailDataConnectorResultOutput() LookupAwsCloudTrailDataConnectorResultOutput {
	return o
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) ToLookupAwsCloudTrailDataConnectorResultOutputWithContext(ctx context.Context) LookupAwsCloudTrailDataConnectorResultOutput {
	return o
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) AwsRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) *string { return v.AwsRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) DataTypes() AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) AwsCloudTrailDataConnectorDataTypesResponse {
		return v.DataTypes
	}).(AwsCloudTrailDataConnectorDataTypesResponseOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupAwsCloudTrailDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAwsCloudTrailDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAwsCloudTrailDataConnectorResultOutput{})
}
