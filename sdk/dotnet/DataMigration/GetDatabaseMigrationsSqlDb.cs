// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataMigration
{
    public static class GetDatabaseMigrationsSqlDb
    {
        /// <summary>
        /// Retrieve the Database Migration resource.
        /// 
        /// Uses Azure REST API version 2023-07-15-preview.
        /// 
        /// Other available API versions: 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Task<GetDatabaseMigrationsSqlDbResult> InvokeAsync(GetDatabaseMigrationsSqlDbArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseMigrationsSqlDbResult>("azure-native:datamigration:getDatabaseMigrationsSqlDb", args ?? new GetDatabaseMigrationsSqlDbArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the Database Migration resource.
        /// 
        /// Uses Azure REST API version 2023-07-15-preview.
        /// 
        /// Other available API versions: 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDatabaseMigrationsSqlDbResult> Invoke(GetDatabaseMigrationsSqlDbInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseMigrationsSqlDbResult>("azure-native:datamigration:getDatabaseMigrationsSqlDb", args ?? new GetDatabaseMigrationsSqlDbInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieve the Database Migration resource.
        /// 
        /// Uses Azure REST API version 2023-07-15-preview.
        /// 
        /// Other available API versions: 2022-03-30-preview, 2025-03-15-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native datamigration [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
        /// </summary>
        public static Output<GetDatabaseMigrationsSqlDbResult> Invoke(GetDatabaseMigrationsSqlDbInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseMigrationsSqlDbResult>("azure-native:datamigration:getDatabaseMigrationsSqlDb", args ?? new GetDatabaseMigrationsSqlDbInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseMigrationsSqlDbArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Complete migration details be included in the response.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// Optional migration operation ID. If this is provided, then details of migration operation for that ID are retrieved. If not provided (default), then details related to most recent or current operation are retrieved.
        /// </summary>
        [Input("migrationOperationId")]
        public string? MigrationOperationId { get; set; }

        /// <summary>
        /// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        [Input("sqlDbInstanceName", required: true)]
        public string SqlDbInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the target database.
        /// </summary>
        [Input("targetDbName", required: true)]
        public string TargetDbName { get; set; } = null!;

        public GetDatabaseMigrationsSqlDbArgs()
        {
        }
        public static new GetDatabaseMigrationsSqlDbArgs Empty => new GetDatabaseMigrationsSqlDbArgs();
    }

    public sealed class GetDatabaseMigrationsSqlDbInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Complete migration details be included in the response.
        /// </summary>
        [Input("expand")]
        public Input<string>? Expand { get; set; }

        /// <summary>
        /// Optional migration operation ID. If this is provided, then details of migration operation for that ID are retrieved. If not provided (default), then details related to most recent or current operation are retrieved.
        /// </summary>
        [Input("migrationOperationId")]
        public Input<string>? MigrationOperationId { get; set; }

        /// <summary>
        /// Name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("sqlDbInstanceName", required: true)]
        public Input<string> SqlDbInstanceName { get; set; } = null!;

        /// <summary>
        /// The name of the target database.
        /// </summary>
        [Input("targetDbName", required: true)]
        public Input<string> TargetDbName { get; set; } = null!;

        public GetDatabaseMigrationsSqlDbInvokeArgs()
        {
        }
        public static new GetDatabaseMigrationsSqlDbInvokeArgs Empty => new GetDatabaseMigrationsSqlDbInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseMigrationsSqlDbResult
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        public readonly string AzureApiVersion;
        public readonly string Id;
        public readonly string Name;
        /// <summary>
        /// Database Migration Resource properties for SQL database.
        /// </summary>
        public readonly Outputs.DatabaseMigrationPropertiesSqlDbResponse Properties;
        /// <summary>
        /// Metadata pertaining to creation and last modification of the resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        public readonly string Type;

        [OutputConstructor]
        private GetDatabaseMigrationsSqlDbResult(
            string azureApiVersion,

            string id,

            string name,

            Outputs.DatabaseMigrationPropertiesSqlDbResponse properties,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            AzureApiVersion = azureApiVersion;
            Id = id;
            Name = name;
            Properties = properties;
            SystemData = systemData;
            Type = type;
        }
    }
}
