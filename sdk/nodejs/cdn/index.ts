import * as pulumi from "@pulumi/pulumi";

export class Profile extends pulumi.CustomResource {
    public readonly profileName: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:cdn:Profile", name, args, opts);
    }
}

export class ProfileEndpoint extends pulumi.CustomResource {
    public readonly properties: pulumi.Output<ProfileEndpointProperties>;

    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:cdn:ProfileEndpoint", name, args, opts);
    }
}

export class ProfileEndpointProperties {
    public readonly hostName: pulumi.Output<string>;
}
