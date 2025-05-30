// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Database Migration Resource for Mongo to CosmosDb.
 *
 * Uses Azure REST API version 2023-07-15-preview. In version 2.x of the Azure Native provider, it used API version 2023-07-15-preview.
 *
 * Other available API versions: 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
 */
export class DatabaseMigrationsMongoToCosmosDbvCoreMongo extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseMigrationsMongoToCosmosDbvCoreMongo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DatabaseMigrationsMongoToCosmosDbvCoreMongo {
        return new DatabaseMigrationsMongoToCosmosDbvCoreMongo(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:datamigration:DatabaseMigrationsMongoToCosmosDbvCoreMongo';

    /**
     * Returns true if the given object is an instance of DatabaseMigrationsMongoToCosmosDbvCoreMongo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseMigrationsMongoToCosmosDbvCoreMongo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseMigrationsMongoToCosmosDbvCoreMongo.__pulumiType;
    }

    /**
     * The Azure API version of the resource.
     */
    public /*out*/ readonly azureApiVersion!: pulumi.Output<string>;
    /**
     * List of Mongo Collections to be migrated.
     */
    public readonly collectionList!: pulumi.Output<outputs.datamigration.MongoMigrationCollectionResponse[] | undefined>;
    /**
     * Database migration end time.
     */
    public /*out*/ readonly endedOn!: pulumi.Output<string>;
    /**
     *
     * Expected value is 'MongoToCosmosDbMongo'.
     */
    public readonly kind!: pulumi.Output<"MongoToCosmosDbMongo">;
    /**
     * Error details in case of migration failure.
     */
    public /*out*/ readonly migrationFailureError!: pulumi.Output<outputs.datamigration.ErrorInfoResponse>;
    /**
     * ID for current migration operation.
     */
    public readonly migrationOperationId!: pulumi.Output<string | undefined>;
    /**
     * Resource Id of the Migration Service.
     */
    public readonly migrationService!: pulumi.Output<string | undefined>;
    /**
     * Migration status.
     */
    public /*out*/ readonly migrationStatus!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Error message for migration provisioning failure, if any.
     */
    public readonly provisioningError!: pulumi.Output<string | undefined>;
    /**
     * Provisioning State of migration. ProvisioningState as Succeeded implies that validations have been performed and migration has started.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Resource Id of the target resource.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Source Mongo connection details.
     */
    public readonly sourceMongoConnection!: pulumi.Output<outputs.datamigration.MongoConnectionInformationResponse | undefined>;
    /**
     * Database migration start time.
     */
    public /*out*/ readonly startedOn!: pulumi.Output<string>;
    /**
     * Azure Resource Manager metadata containing createdBy and modifiedBy information.
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.datamigration.SystemDataResponse>;
    /**
     * Target Cosmos DB Mongo connection details.
     */
    public readonly targetMongoConnection!: pulumi.Output<outputs.datamigration.MongoConnectionInformationResponse | undefined>;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DatabaseMigrationsMongoToCosmosDbvCoreMongo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.kind === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kind'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.targetResourceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetResourceName'");
            }
            resourceInputs["collectionList"] = args ? args.collectionList : undefined;
            resourceInputs["kind"] = "MongoToCosmosDbMongo";
            resourceInputs["migrationName"] = args ? args.migrationName : undefined;
            resourceInputs["migrationOperationId"] = args ? args.migrationOperationId : undefined;
            resourceInputs["migrationService"] = args ? args.migrationService : undefined;
            resourceInputs["provisioningError"] = args ? args.provisioningError : undefined;
            resourceInputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["sourceMongoConnection"] = args ? args.sourceMongoConnection : undefined;
            resourceInputs["targetMongoConnection"] = args ? args.targetMongoConnection : undefined;
            resourceInputs["targetResourceName"] = args ? args.targetResourceName : undefined;
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["endedOn"] = undefined /*out*/;
            resourceInputs["migrationFailureError"] = undefined /*out*/;
            resourceInputs["migrationStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["startedOn"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["azureApiVersion"] = undefined /*out*/;
            resourceInputs["collectionList"] = undefined /*out*/;
            resourceInputs["endedOn"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["migrationFailureError"] = undefined /*out*/;
            resourceInputs["migrationOperationId"] = undefined /*out*/;
            resourceInputs["migrationService"] = undefined /*out*/;
            resourceInputs["migrationStatus"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["provisioningError"] = undefined /*out*/;
            resourceInputs["provisioningState"] = undefined /*out*/;
            resourceInputs["scope"] = undefined /*out*/;
            resourceInputs["sourceMongoConnection"] = undefined /*out*/;
            resourceInputs["startedOn"] = undefined /*out*/;
            resourceInputs["systemData"] = undefined /*out*/;
            resourceInputs["targetMongoConnection"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const aliasOpts = { aliases: [{ type: "azure-native:datamigration/v20230715preview:DatabaseMigrationsMongoToCosmosDbvCoreMongo" }, { type: "azure-native:datamigration/v20250315preview:DatabaseMigrationsMongoToCosmosDbvCoreMongo" }] };
        opts = pulumi.mergeOptions(opts, aliasOpts);
        super(DatabaseMigrationsMongoToCosmosDbvCoreMongo.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DatabaseMigrationsMongoToCosmosDbvCoreMongo resource.
 */
export interface DatabaseMigrationsMongoToCosmosDbvCoreMongoArgs {
    /**
     * List of Mongo Collections to be migrated.
     */
    collectionList?: pulumi.Input<pulumi.Input<inputs.datamigration.MongoMigrationCollectionArgs>[]>;
    /**
     *
     * Expected value is 'MongoToCosmosDbMongo'.
     */
    kind: pulumi.Input<"MongoToCosmosDbMongo">;
    /**
     * Name of the migration.
     */
    migrationName?: pulumi.Input<string>;
    /**
     * ID for current migration operation.
     */
    migrationOperationId?: pulumi.Input<string>;
    /**
     * Resource Id of the Migration Service.
     */
    migrationService?: pulumi.Input<string>;
    /**
     * Error message for migration provisioning failure, if any.
     */
    provisioningError?: pulumi.Input<string>;
    /**
     * Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    resourceGroupName: pulumi.Input<string>;
    /**
     * Resource Id of the target resource.
     */
    scope?: pulumi.Input<string>;
    /**
     * Source Mongo connection details.
     */
    sourceMongoConnection?: pulumi.Input<inputs.datamigration.MongoConnectionInformationArgs>;
    /**
     * Target Cosmos DB Mongo connection details.
     */
    targetMongoConnection?: pulumi.Input<inputs.datamigration.MongoConnectionInformationArgs>;
    /**
     * The name of the target resource/account.
     */
    targetResourceName: pulumi.Input<string>;
}
