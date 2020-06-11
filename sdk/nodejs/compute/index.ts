import * as pulumi from "@pulumi/pulumi";

export class VirtualMachine extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:compute:VirtualMachine", name, args, opts);
    }
}
