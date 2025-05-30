// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Class representing a database principal assignment.
 *
 * Uses Azure REST API version 2021-06-01-preview. In version 2.x of the Azure Native provider, it used API version 2021-06-01-preview.
 */
export class KustoPoolDatabasePrincipalAssignment extends pulumi.CustomResource {
    /**
     * Get an existing KustoPoolDatabasePrincipalAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KustoPoolDatabasePrincipalAssignment {
        return new KustoPoolDatabasePrincipalAssignment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:synapse:KustoPoolDatabasePrincipalAssignment';

    /**
     * Returns true if the given object is an instance of KustoPoolDatabasePrincipalAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KustoPoolDatabasePrincipalAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KustoPoolDatabasePrincipalAssignment.__pulumiType;
    }

    /**
     * The service principal object id in AAD (Azure active directory)
     */
    public /*out*/ readonly aadObjectId!: pulumi.Output<string>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
     */
    public readonly principalId!: pulumi.Output<string>;
    /**
     * The principal name
     */
    public /*out*/ readonly principalName!: pulumi.Output<string>;
    /**
     * Principal type.
     */
    public readonly principalType!: pulumi.Output<string>;
    /**
     * The provisioned state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Database principal role.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.synapse.SystemDataResponse>;
    /**
     * The tenant id of the principal
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The tenant name of the principal
     */
    public /*out*/ readonly tenantName!: pulumi.Output<string>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a KustoPoolDatabasePrincipalAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KustoPoolDatabasePrincipalAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.kustoPoolName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kustoPoolName'");
            }
            if ((!args || args.principalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalId'");
            }
            if ((!args || args.principalType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.workspaceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workspaceName'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["kustoPoolName"] = args ? args.kustoPoolName : undefined;
            resourceInputs["principalAssignmentName"] = args ? args.principalAssignmentName : undefined;
            resourceInputs["principalId"] = args ? args.principalId : undefined;
            resourceInputs["principalType"] = args ? args.principalType : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["workspaceName"] = args ? args.workspaceName : undefined;
            resourceInputs["aadObjectId"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["principalName"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tenantName"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["aadObjectId"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["principalId"] = undefined /*out*/;
            resourceInputs["principalName"] = undefined /*out*/;
            resourceInputs["principalType"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["role"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
            resourceInputs["tenantName"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:synapse/v20210401preview:DatabasePrincipalAssignment" }, { type: "azure-native:synapse/v20210401preview:KustoPoolDatabasePrincipalAssignment" }, { type: "azure-native:synapse/v20210601preview:KustoPoolDatabasePrincipalAssignment" }, { type: "azure-native:synapse:DatabasePrincipalAssignment" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(KustoPoolDatabasePrincipalAssignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KustoPoolDatabasePrincipalAssignment resource.
 */
export interface KustoPoolDatabasePrincipalAssignmentArgs {
    /**
     * The name of the database in the Kusto pool.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The name of the Kusto pool.
     */
    kustoPoolName: pulumi.Input<string>;
    /**
     * The name of the Kusto principalAssignment.
     */
    principalAssignmentName?: pulumi.Input<string>;
    /**
     * The principal ID assigned to the database principal. It can be a user email, application ID, or security group name.
     */
    principalId: pulumi.Input<string>;
    /**
     * Principal type.
     */
    principalType: pulumi.Input<string | enums.synapse.PrincipalType>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Database principal role.
     */
    role: pulumi.Input<string | enums.synapse.DatabasePrincipalRole>;
    /**
     * The tenant id of the principal
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    workspaceName: pulumi.Input<string>;
}
