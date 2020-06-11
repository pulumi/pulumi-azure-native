import * as pulumi from "@pulumi/pulumi";

export class Server extends pulumi.CustomResource {
    public readonly serverName: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:sql:Server", name, args, opts);
    }
}

export class ServerDatabase extends pulumi.CustomResource {
    public readonly databaseName: pulumi.Output<string>;
    constructor(name: string, args: any, opts?: pulumi.CustomResourceOptions) {
        super("azurerm:sql:ServerDatabase", name, args, opts);
    }
}
