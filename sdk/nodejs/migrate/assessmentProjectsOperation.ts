// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * An Assessment project site resource.
 *
 * Uses Azure REST API version 2024-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-03-15.
 *
 * Other available API versions: 2023-03-15, 2023-04-01-preview, 2023-05-01-preview, 2023-09-09-preview, 2024-01-15, 2024-03-03-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native migrate [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class AssessmentProjectsOperation extends pulumi.CustomResource {
    /**
     * Get an existing AssessmentProjectsOperation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AssessmentProjectsOperation {
        return new AssessmentProjectsOperation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:migrate:AssessmentProjectsOperation';

    /**
     * Returns true if the given object is an instance of AssessmentProjectsOperation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AssessmentProjectsOperation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AssessmentProjectsOperation.__pulumiType;
    }

    /**
     * Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
     */
    public readonly assessmentSolutionId!: pulumi.Output<string | undefined>;
    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * Time when this project was created. Date-Time represented in ISO-8601 format.
     */
    public /*out*/ readonly createdTimestamp!: pulumi.Output<string>;
    /**
     * The ARM id of the storage account used for interactions when public access is
     * disabled.
     */
    public readonly customerStorageAccountArmId!: pulumi.Output<string | undefined>;
    /**
     * The ARM id of service map workspace created by customer.
     */
    public readonly customerWorkspaceId!: pulumi.Output<string | undefined>;
    /**
     * Location of service map workspace created by customer.
     */
    public readonly customerWorkspaceLocation!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of private endpoint connections to the project.
     */
    public /*out*/ readonly privateEndpointConnections!: pulumi.Output<outputs.migrate.PrivateEndpointConnectionResponse[]>;
    /**
     * Assessment project status.
     */
    public readonly projectStatus!: pulumi.Output<string | undefined>;
    /**
     * The status of the last operation.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * This value can be set to 'enabled' to avoid breaking changes on existing
     * customer resources and templates. If set to 'disabled', traffic over public
     * interface is not allowed, and private endpoint connections would be the
     * exclusive access method.
     */
    public readonly publicNetworkAccess!: pulumi.Output<string | undefined>;
    /**
     * Endpoint at which the collector agent can call agent REST API.
     */
    public /*out*/ readonly serviceEndpoint!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.migrate.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Time when this project was last updated. Date-Time represented in ISO-8601
     * format.
     */
    public /*out*/ readonly updatedTimestamp!: pulumi.Output<string>;

    /**
     * Create a AssessmentProjectsOperation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssessmentProjectsOperationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            resourceInputs["assessmentSolutionId"] = args ? args.assessmentSolutionId : undefined;
            resourceInputs["customerStorageAccountArmId"] = args ? args.customerStorageAccountArmId : undefined;
            resourceInputs["customerWorkspaceId"] = args ? args.customerWorkspaceId : undefined;
            resourceInputs["customerWorkspaceLocation"] = args ? args.customerWorkspaceLocation : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["projectStatus"] = args ? args.projectStatus : undefined;
            resourceInputs["provisioningState"] = args ? args.provisioningState : undefined;
            resourceInputs["publicNetworkAccess"] = args ? args.publicNetworkAccess : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["createdTimestamp"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["privateEndpointConnections"] = undefined /*out*/;
            resourceInputs["serviceEndpoint"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedTimestamp"] = undefined /*out*/;
        } else {
            resourceInputs["assessmentSolutionId"] = undefined /*out*/;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["createdTimestamp"] = undefined /*out*/;
            resourceInputs["customerStorageAccountArmId"] = undefined /*out*/;
            resourceInputs["customerWorkspaceId"] = undefined /*out*/;
            resourceInputs["customerWorkspaceLocation"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["privateEndpointConnections"] = undefined /*out*/;
            resourceInputs["projectStatus"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["publicNetworkAccess"] = undefined /*out*/;
            resourceInputs["serviceEndpoint"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updatedTimestamp"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:migrate/v20191001:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20191001:Project" }, { type: "azure-native:migrate/v20230315:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20230401preview:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20230501preview:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20230909preview:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20240101preview:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20240115:AssessmentProjectsOperation" }, { type: "azure-native:migrate/v20240303preview:AssessmentProjectsOperation" }, { type: "azure-native:migrate:Project" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(AssessmentProjectsOperation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AssessmentProjectsOperation resource.
 */
export interface AssessmentProjectsOperationArgs {
    /**
     * Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
     */
    assessmentSolutionId?: pulumi.Input<string>;
    /**
     * The ARM id of the storage account used for interactions when public access is
     * disabled.
     */
    customerStorageAccountArmId?: pulumi.Input<string>;
    /**
     * The ARM id of service map workspace created by customer.
     */
    customerWorkspaceId?: pulumi.Input<string>;
    /**
     * Location of service map workspace created by customer.
     */
    customerWorkspaceLocation?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    location?: pulumi.Input<string>;
    /**
     * Assessment Project Name
     */
    projectName?: pulumi.Input<string>;
    /**
     * Assessment project status.
     */
    projectStatus?: pulumi.Input<string | enums.migrate.ProjectStatus>;
    /**
     * The status of the last operation.
     */
    provisioningState?: pulumi.Input<string | enums.migrate.ProvisioningState>;
    /**
     * This value can be set to 'enabled' to avoid breaking changes on existing
     * customer resources and templates. If set to 'disabled', traffic over public
     * interface is not allowed, and private endpoint connections would be the
     * exclusive access method.
     */
    publicNetworkAccess?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
