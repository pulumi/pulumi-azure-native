// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Sql.Inputs
{

    /// <summary>
    /// Read-only endpoint of the failover group instance.
    /// </summary>
    public sealed class InstanceFailoverGroupReadOnlyEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Failover policy of the read-only endpoint for the failover group.
        /// </summary>
        [Input("failoverPolicy")]
        public InputUnion<string, Pulumi.AzureNative.Sql.ReadOnlyEndpointFailoverPolicy>? FailoverPolicy { get; set; }

        public InstanceFailoverGroupReadOnlyEndpointArgs()
        {
        }
        public static new InstanceFailoverGroupReadOnlyEndpointArgs Empty => new InstanceFailoverGroupReadOnlyEndpointArgs();
    }
}
