// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataProtection.V20240201Preview.Outputs
{

    /// <summary>
    /// Generic parameters to be used during configuration of backup
    /// </summary>
    [OutputType]
    public sealed class GenericBackupDatasourceParametersResponse
    {
        /// <summary>
        /// Type of the specific object - used for deserializing
        /// Expected value is 'GenericBackupDatasourceParameters'.
        /// </summary>
        public readonly string ObjectType;
        /// <summary>
        /// List of resource selectors to be backed up during configuration of backup
        /// </summary>
        public readonly ImmutableArray<string> ResourceSelectors;

        [OutputConstructor]
        private GenericBackupDatasourceParametersResponse(
            string objectType,

            ImmutableArray<string> resourceSelectors)
        {
            ObjectType = objectType;
            ResourceSelectors = resourceSelectors;
        }
    }
}
