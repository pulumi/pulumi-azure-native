// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Outputs
{

    /// <summary>
    /// Optional. The object replication policy metrics feature options.
    /// </summary>
    [OutputType]
    public sealed class ObjectReplicationPolicyPropertiesResponseMetrics
    {
        /// <summary>
        /// Indicates whether object replication metrics feature is enabled for the policy.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private ObjectReplicationPolicyPropertiesResponseMetrics(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
