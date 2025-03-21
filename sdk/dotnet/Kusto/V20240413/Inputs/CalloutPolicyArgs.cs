// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Kusto.V20240413.Inputs
{

    /// <summary>
    /// Configuration for external callout policies, including URI patterns, access types, and service types.
    /// </summary>
    public sealed class CalloutPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the callout service, specifying the kind of external resource or service being accessed.
        /// </summary>
        [Input("calloutType")]
        public InputUnion<string, Pulumi.AzureNative.Kusto.V20240413.CalloutType>? CalloutType { get; set; }

        /// <summary>
        /// Regular expression or FQDN pattern for the callout URI.
        /// </summary>
        [Input("calloutUriRegex")]
        public Input<string>? CalloutUriRegex { get; set; }

        /// <summary>
        /// Indicates whether outbound access is permitted for the specified URI pattern.
        /// </summary>
        [Input("outboundAccess")]
        public InputUnion<string, Pulumi.AzureNative.Kusto.V20240413.OutboundAccess>? OutboundAccess { get; set; }

        public CalloutPolicyArgs()
        {
        }
        public static new CalloutPolicyArgs Empty => new CalloutPolicyArgs();
    }
}
