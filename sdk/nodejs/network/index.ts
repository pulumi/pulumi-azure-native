import * as pulumi from "@pulumi/pulumi";

export class NetworkInterface extends pulumi.CustomResource {
    id: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:NetworkInterface", name, args, opts);
    }
}

export class VirtualNetwork extends pulumi.CustomResource {
    public readonly virtualNetworkName: pulumi.Output<string>;
    public readonly id: pulumi.Output<string>;

    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:VirtualNetwork", name, args, opts);
    }
}

export class Subnet extends pulumi.CustomResource {
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:network:VirtualNetworkSubnet", name, args, opts);
    }
}
