import * as pulumi from "@pulumi/pulumi";

export class ContainerGroup extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:containerinstance:ContainerGroup", name, args, opts);
    }
}
