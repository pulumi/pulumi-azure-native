// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EventGrid.Outputs
{

    /// <summary>
    /// Properties of the Topics Configuration.
    /// </summary>
    [OutputType]
    public sealed class TopicsConfigurationResponse
    {
        /// <summary>
        /// List of custom domain configurations for the namespace.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomDomainConfigurationResponse> CustomDomains;
        /// <summary>
        /// The hostname for the topics configuration. This is a read-only property.
        /// </summary>
        public readonly string Hostname;

        [OutputConstructor]
        private TopicsConfigurationResponse(
            ImmutableArray<Outputs.CustomDomainConfigurationResponse> customDomains,

            string hostname)
        {
            CustomDomains = customDomains;
            Hostname = hostname;
        }
    }
}
