


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSshPublicKey(ctx *pulumi.Context, args *LookupSshPublicKeyArgs, opts ...pulumi.InvokeOption) (*LookupSshPublicKeyResult, error) {
	var rv LookupSshPublicKeyResult
	err := ctx.Invoke("azure-native:compute/v20201201:getSshPublicKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSshPublicKeyArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SshPublicKeyName  string `pulumi:"sshPublicKeyName"`
}


type LookupSshPublicKeyResult struct {
	Id        string            `pulumi:"id"`
	Location  string            `pulumi:"location"`
	Name      string            `pulumi:"name"`
	PublicKey *string           `pulumi:"publicKey"`
	Tags      map[string]string `pulumi:"tags"`
	Type      string            `pulumi:"type"`
}
