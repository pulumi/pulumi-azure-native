// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Storage.Inputs
{

    /// <summary>
    /// Optional. The object replication policy metrics feature options.
    /// </summary>
    public sealed class ObjectReplicationPolicyPropertiesMetricsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether object replication metrics feature is enabled for the policy.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ObjectReplicationPolicyPropertiesMetricsArgs()
        {
        }
        public static new ObjectReplicationPolicyPropertiesMetricsArgs Empty => new ObjectReplicationPolicyPropertiesMetricsArgs();
    }
}
