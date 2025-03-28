// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DocumentDB.V20240515Preview.Outputs
{

    /// <summary>
    /// Configuration values for periodic mode backup
    /// </summary>
    [OutputType]
    public sealed class ContinuousModePropertiesResponse
    {
        /// <summary>
        /// Enum to indicate type of Continuos backup mode
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private ContinuousModePropertiesResponse(string? tier)
        {
            Tier = tier;
        }
    }
}
