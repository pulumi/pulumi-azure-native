// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.V20241001Preview.Outputs
{

    /// <summary>
    /// This represents the fabric specific settings for export destination when delivery info type is 'MicrosoftFabric'.
    /// </summary>
    [OutputType]
    public sealed class MicrosoftFabricDestinationSettingsResponse
    {
        /// <summary>
        /// The Microsoft Fabric Lakehouse Identifier.
        /// </summary>
        public readonly Outputs.FabricLakehouseResponse? Lakehouse;
        /// <summary>
        /// The name of the Microsoft Fabric Lakehouse Delta table where export data will be ingested to.
        /// </summary>
        public readonly string? TableName;
        /// <summary>
        /// The Microsoft Fabric Workspace Identifier.
        /// </summary>
        public readonly Outputs.FabricWorkspaceResponse? Workspace;

        [OutputConstructor]
        private MicrosoftFabricDestinationSettingsResponse(
            Outputs.FabricLakehouseResponse? lakehouse,

            string? tableName,

            Outputs.FabricWorkspaceResponse? workspace)
        {
            Lakehouse = lakehouse;
            TableName = tableName;
            Workspace = workspace;
        }
    }
}
