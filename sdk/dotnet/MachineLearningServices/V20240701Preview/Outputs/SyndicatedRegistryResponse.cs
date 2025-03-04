// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20240701Preview.Outputs
{

    /// <summary>
    /// A registry that is syndicated
    /// </summary>
    [OutputType]
    public sealed class SyndicatedRegistryResponse
    {
        /// <summary>
        /// The Registry Id of the syndicated Registry
        /// </summary>
        public readonly string? RegistryId;

        [OutputConstructor]
        private SyndicatedRegistryResponse(string? registryId)
        {
            RegistryId = registryId;
        }
    }
}
