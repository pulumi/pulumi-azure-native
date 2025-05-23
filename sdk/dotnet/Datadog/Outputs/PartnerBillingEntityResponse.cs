// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Datadog.Outputs
{

    /// <summary>
    /// Partner Billing details associated with the resource.
    /// </summary>
    [OutputType]
    public sealed class PartnerBillingEntityResponse
    {
        /// <summary>
        /// The Datadog Organization Id.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The Datadog Organization Name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Link to the datadog organization page
        /// </summary>
        public readonly string? PartnerEntityUri;

        [OutputConstructor]
        private PartnerBillingEntityResponse(
            string? id,

            string? name,

            string? partnerEntityUri)
        {
            Id = id;
            Name = name;
            PartnerEntityUri = partnerEntityUri;
        }
    }
}
