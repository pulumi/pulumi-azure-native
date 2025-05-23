// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Outputs
{

    /// <summary>
    /// Control Plane Apis version constraint for the API Management service.
    /// </summary>
    [OutputType]
    public sealed class ApiVersionConstraintResponse
    {
        /// <summary>
        /// Limit control plane API calls to API Management service with version equal to or newer than this value.
        /// </summary>
        public readonly string? MinApiVersion;

        [OutputConstructor]
        private ApiVersionConstraintResponse(string? minApiVersion)
        {
            MinApiVersion = minApiVersion;
        }
    }
}
