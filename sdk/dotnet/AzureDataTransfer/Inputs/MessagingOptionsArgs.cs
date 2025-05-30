// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureDataTransfer.Inputs
{

    /// <summary>
    /// The option associated with messaging flows.
    /// </summary>
    public sealed class MessagingOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Billing tier for this messaging flow
        /// </summary>
        [Input("billingTier")]
        public InputUnion<string, Pulumi.AzureNative.AzureDataTransfer.FlowBillingTier>? BillingTier { get; set; }

        public MessagingOptionsArgs()
        {
        }
        public static new MessagingOptionsArgs Empty => new MessagingOptionsArgs();
    }
}
