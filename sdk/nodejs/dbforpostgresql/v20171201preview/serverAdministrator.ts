// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Represents a and external administrator to be created.
 */
export class ServerAdministrator extends pulumi.CustomResource {
    /**
     * Get an existing ServerAdministrator resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServerAdministrator {
        return new ServerAdministrator(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:dbforpostgresql/v20171201preview:ServerAdministrator';

    /**
     * Returns true if the given object is an instance of ServerAdministrator.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServerAdministrator {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServerAdministrator.__pulumiType;
    }

    /**
     * The type of administrator.
     */
    public readonly administratorType!: pulumi.Output<string>;
    /**
     * The server administrator login account name.
     */
    public readonly login!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The server administrator Sid (Secure ID).
     */
    public readonly sid!: pulumi.Output<string>;
    /**
     * The server Active Directory Administrator tenant id.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ServerAdministrator resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerAdministratorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.administratorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'administratorType'");
            }
            if ((!args || args.login === undefined) && !opts.urn) {
                throw new Error("Missing required property 'login'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.sid === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sid'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["administratorType"] = args ? args.administratorType : undefined;
            resourceInputs["login"] = args ? args.login : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["sid"] = args ? args.sid : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["administratorType"] = undefined /*out*/;
            resourceInputs["login"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["sid"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:dbforpostgresql/v20171201:ServerAdministrator" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(ServerAdministrator.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServerAdministrator resource.
 */
export interface ServerAdministratorArgs {
    /**
     * The type of administrator.
     */
    administratorType: pulumi.Input<string | enums.dbforpostgresql.v20171201preview.AdministratorType>;
    /**
     * The server administrator login account name.
     */
    login: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the server.
     */
    serverName: pulumi.Input<string>;
    /**
     * The server administrator Sid (Secure ID).
     */
    sid: pulumi.Input<string>;
    /**
     * The server Active Directory Administrator tenant id.
     */
    tenantId: pulumi.Input<string>;
}
