package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storageactions/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := storageactions.NewStorageTask(ctx, "storageTask", &storageactions.StorageTaskArgs{
Action: storageactions.StorageTaskActionResponse{
Else: interface{}{
Operations: storageactions.StorageTaskOperationArray{
&storageactions.StorageTaskOperationArgs{
Name: pulumi.String("DeleteBlob"),
OnFailure: storageactions.OnFailureBreak,
OnSuccess: storageactions.OnSuccessContinue,
},
},
},
If: interface{}{
Condition: pulumi.String("[[equals(AccessTier, 'Cool')]]"),
Operations: storageactions.StorageTaskOperationArray{
&storageactions.StorageTaskOperationArgs{
Name: pulumi.String("SetBlobTier"),
OnFailure: storageactions.OnFailureBreak,
OnSuccess: storageactions.OnSuccessContinue,
Parameters: pulumi.StringMap{
"tier": pulumi.String("Hot"),
},
},
},
},
},
Description: pulumi.String("My Storage task"),
Enabled: pulumi.Bool(true),
Location: pulumi.String("westus"),
ResourceGroupName: pulumi.String("res4228"),
StorageTaskName: pulumi.String("mytask1"),
})
if err != nil {
return err
}
return nil
})
}
