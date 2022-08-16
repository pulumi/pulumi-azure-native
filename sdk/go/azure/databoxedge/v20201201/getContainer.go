


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupContainer(ctx *pulumi.Context, args *LookupContainerArgs, opts ...pulumi.InvokeOption) (*LookupContainerResult, error) {
	var rv LookupContainerResult
	err := ctx.Invoke("azure-native:databoxedge/v20201201:getContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerArgs struct {
	ContainerName      string `pulumi:"containerName"`
	DeviceName         string `pulumi:"deviceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	StorageAccountName string `pulumi:"storageAccountName"`
}


type LookupContainerResult struct {
	ContainerStatus string                 `pulumi:"containerStatus"`
	CreatedDateTime string                 `pulumi:"createdDateTime"`
	DataFormat      string                 `pulumi:"dataFormat"`
	Id              string                 `pulumi:"id"`
	Name            string                 `pulumi:"name"`
	RefreshDetails  RefreshDetailsResponse `pulumi:"refreshDetails"`
	SystemData      SystemDataResponse     `pulumi:"systemData"`
	Type            string                 `pulumi:"type"`
}

func LookupContainerOutput(ctx *pulumi.Context, args LookupContainerOutputArgs, opts ...pulumi.InvokeOption) LookupContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerResult, error) {
			args := v.(LookupContainerArgs)
			r, err := LookupContainer(ctx, &args, opts...)
			var s LookupContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerResultOutput)
}

type LookupContainerOutputArgs struct {
	ContainerName      pulumi.StringInput `pulumi:"containerName"`
	DeviceName         pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	StorageAccountName pulumi.StringInput `pulumi:"storageAccountName"`
}

func (LookupContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerArgs)(nil)).Elem()
}


type LookupContainerResultOutput struct{ *pulumi.OutputState }

func (LookupContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerResult)(nil)).Elem()
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutput() LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutputWithContext(ctx context.Context) LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) ContainerStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.ContainerStatus }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) DataFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.DataFormat }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupContainerResultOutput) RefreshDetails() RefreshDetailsResponseOutput {
	return o.ApplyT(func(v LookupContainerResult) RefreshDetailsResponse { return v.RefreshDetails }).(RefreshDetailsResponseOutput)
}

func (o LookupContainerResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupContainerResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupContainerResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerResultOutput{})
}
