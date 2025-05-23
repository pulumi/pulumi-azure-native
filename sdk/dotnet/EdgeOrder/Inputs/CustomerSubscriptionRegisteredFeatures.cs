// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Inputs
{

    /// <summary>
    /// Represents subscription registered features.
    /// </summary>
    public sealed class CustomerSubscriptionRegisteredFeatures : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of subscription registered feature.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// State of subscription registered feature.
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public CustomerSubscriptionRegisteredFeatures()
        {
        }
        public static new CustomerSubscriptionRegisteredFeatures Empty => new CustomerSubscriptionRegisteredFeatures();
    }
}
