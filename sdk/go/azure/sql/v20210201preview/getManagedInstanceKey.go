


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagedInstanceKey(ctx *pulumi.Context, args *LookupManagedInstanceKeyArgs, opts ...pulumi.InvokeOption) (*LookupManagedInstanceKeyResult, error) {
	var rv LookupManagedInstanceKeyResult
	err := ctx.Invoke("azure-native:sql/v20210201preview:getManagedInstanceKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagedInstanceKeyArgs struct {
	KeyName             string `pulumi:"keyName"`
	ManagedInstanceName string `pulumi:"managedInstanceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupManagedInstanceKeyResult struct {
	AutoRotationEnabled bool   `pulumi:"autoRotationEnabled"`
	CreationDate        string `pulumi:"creationDate"`
	Id                  string `pulumi:"id"`
	Kind                string `pulumi:"kind"`
	Name                string `pulumi:"name"`
	Thumbprint          string `pulumi:"thumbprint"`
	Type                string `pulumi:"type"`
}
