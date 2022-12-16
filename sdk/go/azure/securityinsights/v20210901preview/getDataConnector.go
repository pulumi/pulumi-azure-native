


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupDataConnector(ctx *pulumi.Context, args *LookupDataConnectorArgs, opts ...pulumi.InvokeOption) (*LookupDataConnectorResult, error) {
	var rv LookupDataConnectorResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getDataConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataConnectorArgs struct {
	DataConnectorId   string `pulumi:"dataConnectorId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataConnectorResult struct {
	Etag       *string            `pulumi:"etag"`
	Id         string             `pulumi:"id"`
	Kind       string             `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func LookupDataConnectorOutput(ctx *pulumi.Context, args LookupDataConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupDataConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataConnectorResult, error) {
			args := v.(LookupDataConnectorArgs)
			r, err := LookupDataConnector(ctx, &args, opts...)
			var s LookupDataConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataConnectorResultOutput)
}

type LookupDataConnectorOutputArgs struct {
	DataConnectorId   pulumi.StringInput `pulumi:"dataConnectorId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupDataConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataConnectorArgs)(nil)).Elem()
}


type LookupDataConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupDataConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataConnectorResult)(nil)).Elem()
}

func (o LookupDataConnectorResultOutput) ToLookupDataConnectorResultOutput() LookupDataConnectorResultOutput {
	return o
}

func (o LookupDataConnectorResultOutput) ToLookupDataConnectorResultOutputWithContext(ctx context.Context) LookupDataConnectorResultOutput {
	return o
}

func (o LookupDataConnectorResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupDataConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataConnectorResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupDataConnectorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDataConnectorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDataConnectorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataConnectorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataConnectorResultOutput{})
}
