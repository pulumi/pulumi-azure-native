import * as pulumi from "@pulumi/pulumi";

export class ResourceGroup extends pulumi.CustomResource {
    public readonly resourceGroupName: pulumi.Output<string>;

    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:core:ResourceGroup", name, args, opts);
    }
}
