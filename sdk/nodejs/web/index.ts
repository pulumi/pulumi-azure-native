import * as pulumi from "@pulumi/pulumi";

export class StaticSite extends pulumi.CustomResource {
    public readonly properties: pulumi.Output<StaticSiteProperties>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:web:StaticSite", name, args, opts);
    }
}

export class StaticSiteProperties {
    public readonly defaultHostname: pulumi.Output<string>;
}

export class AppServicePlan extends pulumi.CustomResource {
    id: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:web:AppServicePlan", name, args, opts);
    }
}

export class AppService extends pulumi.CustomResource {
    public readonly properties: pulumi.Output<AppServiceProperties>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:web:AppService", name, args, opts);
    }
}

export class AppServiceProperties {
    public readonly defaultHostName: pulumi.Output<string>;
}