


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetkustoPool(ctx *pulumi.Context, args *GetkustoPoolArgs, opts ...pulumi.InvokeOption) (*GetkustoPoolResult, error) {
	var rv GetkustoPoolResult
	err := ctx.Invoke("azure-native:synapse/v20210401preview:getkustoPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetkustoPoolArgs struct {
	KustoPoolName     string `pulumi:"kustoPoolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetkustoPoolResult struct {
	DataIngestionUri  string             `pulumi:"dataIngestionUri"`
	EngineType        *string            `pulumi:"engineType"`
	Etag              string             `pulumi:"etag"`
	Id                string             `pulumi:"id"`
	Location          string             `pulumi:"location"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	Sku               AzureSkuResponse   `pulumi:"sku"`
	State             string             `pulumi:"state"`
	StateReason       string             `pulumi:"stateReason"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Tags              map[string]string  `pulumi:"tags"`
	Type              string             `pulumi:"type"`
	Uri               string             `pulumi:"uri"`
	WorkspaceUid      *string            `pulumi:"workspaceUid"`
}

func GetkustoPoolOutput(ctx *pulumi.Context, args GetkustoPoolOutputArgs, opts ...pulumi.InvokeOption) GetkustoPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetkustoPoolResult, error) {
			args := v.(GetkustoPoolArgs)
			r, err := GetkustoPool(ctx, &args, opts...)
			var s GetkustoPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetkustoPoolResultOutput)
}

type GetkustoPoolOutputArgs struct {
	KustoPoolName     pulumi.StringInput `pulumi:"kustoPoolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetkustoPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetkustoPoolArgs)(nil)).Elem()
}


type GetkustoPoolResultOutput struct{ *pulumi.OutputState }

func (GetkustoPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetkustoPoolResult)(nil)).Elem()
}

func (o GetkustoPoolResultOutput) ToGetkustoPoolResultOutput() GetkustoPoolResultOutput {
	return o
}

func (o GetkustoPoolResultOutput) ToGetkustoPoolResultOutputWithContext(ctx context.Context) GetkustoPoolResultOutput {
	return o
}

func (o GetkustoPoolResultOutput) DataIngestionUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.DataIngestionUri }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) EngineType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetkustoPoolResult) *string { return v.EngineType }).(pulumi.StringPtrOutput)
}

func (o GetkustoPoolResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) Sku() AzureSkuResponseOutput {
	return o.ApplyT(func(v GetkustoPoolResult) AzureSkuResponse { return v.Sku }).(AzureSkuResponseOutput)
}

func (o GetkustoPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.StateReason }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v GetkustoPoolResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GetkustoPoolResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetkustoPoolResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetkustoPoolResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GetkustoPoolResult) string { return v.Uri }).(pulumi.StringOutput)
}

func (o GetkustoPoolResultOutput) WorkspaceUid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetkustoPoolResult) *string { return v.WorkspaceUid }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetkustoPoolResultOutput{})
}
