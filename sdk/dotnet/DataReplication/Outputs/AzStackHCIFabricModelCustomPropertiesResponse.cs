// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataReplication.Outputs
{

    /// <summary>
    /// AzStackHCI fabric model custom properties.
    /// </summary>
    [OutputType]
    public sealed class AzStackHCIFabricModelCustomPropertiesResponse
    {
        /// <summary>
        /// Gets or sets the Appliance name.
        /// </summary>
        public readonly ImmutableArray<string> ApplianceName;
        /// <summary>
        /// Gets or sets the ARM Id of the AzStackHCI site.
        /// </summary>
        public readonly string AzStackHciSiteId;
        /// <summary>
        /// AzStackHCI cluster properties.
        /// </summary>
        public readonly Outputs.AzStackHCIClusterPropertiesResponse Cluster;
        /// <summary>
        /// Gets or sets the fabric container Id.
        /// </summary>
        public readonly string FabricContainerId;
        /// <summary>
        /// Gets or sets the fabric resource Id.
        /// </summary>
        public readonly string FabricResourceId;
        /// <summary>
        /// Gets or sets the instance type.
        /// Expected value is 'AzStackHCI'.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// Gets or sets the migration hub Uri.
        /// </summary>
        public readonly string MigrationHubUri;
        /// <summary>
        /// Gets or sets the Migration solution ARM Id.
        /// </summary>
        public readonly string MigrationSolutionId;

        [OutputConstructor]
        private AzStackHCIFabricModelCustomPropertiesResponse(
            ImmutableArray<string> applianceName,

            string azStackHciSiteId,

            Outputs.AzStackHCIClusterPropertiesResponse cluster,

            string fabricContainerId,

            string fabricResourceId,

            string instanceType,

            string migrationHubUri,

            string migrationSolutionId)
        {
            ApplianceName = applianceName;
            AzStackHciSiteId = azStackHciSiteId;
            Cluster = cluster;
            FabricContainerId = fabricContainerId;
            FabricResourceId = fabricResourceId;
            InstanceType = instanceType;
            MigrationHubUri = migrationHubUri;
            MigrationSolutionId = migrationSolutionId;
        }
    }
}
