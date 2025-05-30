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
    /// Specify the Datadog organization name. In the case of linking to existing organizations, Id, ApiKey, and Applicationkey is required as well.
    /// </summary>
    [OutputType]
    public sealed class DatadogOrganizationPropertiesResponse
    {
        /// <summary>
        /// The configuration which describes the state of cloud security posture management. This collects configuration information for all resources in a subscription and track conformance to industry benchmarks.
        /// </summary>
        public readonly bool? Cspm;
        /// <summary>
        /// Id of the Datadog organization.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Name of the Datadog organization.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private DatadogOrganizationPropertiesResponse(
            bool? cspm,

            string? id,

            string? name)
        {
            Cspm = cspm;
            Id = id;
            Name = name;
        }
    }
}
