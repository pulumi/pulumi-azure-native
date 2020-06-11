import * as pulumi from "@pulumi/pulumi";

export class Component extends pulumi.CustomResource {
    public readonly properties: pulumi.Output<ComponentProperties>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:insights:Component", name, args, opts);
    }
}

export class ComponentProperties {
    public readonly InstrumentationKey: pulumi.Output<string>;
}