// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Outputs
{

    /// <summary>
    /// Class for deployed resource properties.
    /// </summary>
    [OutputType]
    public sealed class DeployedResourcesPropertiesResponse
    {
        /// <summary>
        /// Gets or sets the context of deployed resources.
        /// </summary>
        public readonly string Context;
        /// <summary>
        /// Gets or sets the deployed resource id.
        /// </summary>
        public readonly string DeployedResourceId;
        /// <summary>
        /// Gets or sets the deployment timestamp.
        /// </summary>
        public readonly string DeploymentTimestamp;
        /// <summary>
        /// Gets or sets the name of deployed resources.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Gets or sets the ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Gets or sets a value indicating whether resources are cleaned up from target.
        /// </summary>
        public readonly bool IsCleanUpDone;
        /// <summary>
        /// Gets or sets a value indicating whether scenario is test migration.
        /// </summary>
        public readonly bool IsTestMigration;
        /// <summary>
        /// Gets or sets the status of deployed resources.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Gets or sets the type of deployed resources.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private DeployedResourcesPropertiesResponse(
            string context,

            string deployedResourceId,

            string deploymentTimestamp,

            string displayName,

            string id,

            bool isCleanUpDone,

            bool isTestMigration,

            string status,

            string type)
        {
            Context = context;
            DeployedResourceId = deployedResourceId;
            DeploymentTimestamp = deploymentTimestamp;
            DisplayName = displayName;
            Id = id;
            IsCleanUpDone = isCleanUpDone;
            IsTestMigration = isTestMigration;
            Status = status;
            Type = type;
        }
    }
}
