// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Static Site Database Connection resource.
 *
 * Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
 *
 * Other available API versions: 2022-09-01, 2023-01-01, 2023-12-01, 2024-11-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class StaticSiteDatabaseConnection extends pulumi.CustomResource {
    /**
     * Get an existing StaticSiteDatabaseConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): StaticSiteDatabaseConnection {
        return new StaticSiteDatabaseConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:web:StaticSiteDatabaseConnection';

    /**
     * Returns true if the given object is an instance of StaticSiteDatabaseConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StaticSiteDatabaseConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StaticSiteDatabaseConnection.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * A list of configuration files associated with this database connection.
     */
    public /*out*/ readonly configurationFiles!: pulumi.Output<outputs.web.StaticSiteDatabaseConnectionConfigurationFileOverviewResponse[]>;
    /**
     * If present, the identity is used in conjunction with connection string to connect to the database. Use of the system-assigned managed identity is indicated with the string 'SystemAssigned', while use of a user-assigned managed identity is indicated with the resource id of the managed identity resource.
     */
    public readonly connectionIdentity!: pulumi.Output<string | undefined>;
    /**
     * The connection string to use to connect to the database.
     */
    public readonly connectionString!: pulumi.Output<string | undefined>;
    /**
     * Kind of resource.
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * Resource Name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The region of the database resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The resource id of the database.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a StaticSiteDatabaseConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StaticSiteDatabaseConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            resourceInputs["connectionIdentity"] = args ? args.connectionIdentity : undefined;
            resourceInputs["connectionString"] = args ? args.connectionString : undefined;
            resourceInputs["databaseConnectionName"] = args ? args.databaseConnectionName : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["configurationFiles"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["configurationFiles"] = undefined /*out*/;
            resourceInputs["connectionIdentity"] = undefined /*out*/;
            resourceInputs["connectionString"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:web/v20220901:StaticSiteDatabaseConnection" }, { type: "azure-native:web/v20230101:StaticSiteDatabaseConnection" }, { type: "azure-native:web/v20231201:StaticSiteDatabaseConnection" }, { type: "azure-native:web/v20240401:StaticSiteDatabaseConnection" }, { type: "azure-native:web/v20241101:StaticSiteDatabaseConnection" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(StaticSiteDatabaseConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a StaticSiteDatabaseConnection resource.
 */
export interface StaticSiteDatabaseConnectionArgs {
    /**
     * If present, the identity is used in conjunction with connection string to connect to the database. Use of the system-assigned managed identity is indicated with the string 'SystemAssigned', while use of a user-assigned managed identity is indicated with the resource id of the managed identity resource.
     */
    connectionIdentity?: pulumi.Input<string>;
    /**
     * The connection string to use to connect to the database.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * Name of the database connection.
     */
    databaseConnectionName?: pulumi.Input<string>;
    /**
     * Kind of resource.
     */
    kind?: pulumi.Input<string>;
    /**
     * Name of the static site
     */
    name: pulumi.Input<string>;
    /**
     * The region of the database resource.
     */
    region: pulumi.Input<string>;
    /**
     * Name of the resource group to which the resource belongs.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * The resource id of the database.
     */
    resourceId: pulumi.Input<string>;
}
