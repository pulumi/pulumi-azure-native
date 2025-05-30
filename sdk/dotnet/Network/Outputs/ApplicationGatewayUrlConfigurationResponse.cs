// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network.Outputs
{

    /// <summary>
    /// Url configuration of the Actions set in Application Gateway.
    /// </summary>
    [OutputType]
    public sealed class ApplicationGatewayUrlConfigurationResponse
    {
        /// <summary>
        /// Url path which user has provided for url rewrite. Null means no path will be updated. Default value is null.
        /// </summary>
        public readonly string? ModifiedPath;
        /// <summary>
        /// Query string which user has provided for url rewrite. Null means no query string will be updated. Default value is null.
        /// </summary>
        public readonly string? ModifiedQueryString;
        /// <summary>
        /// If set as true, it will re-evaluate the url path map provided in path based request routing rules using modified path. Default value is false.
        /// </summary>
        public readonly bool? Reroute;

        [OutputConstructor]
        private ApplicationGatewayUrlConfigurationResponse(
            string? modifiedPath,

            string? modifiedQueryString,

            bool? reroute)
        {
            ModifiedPath = modifiedPath;
            ModifiedQueryString = modifiedQueryString;
            Reroute = reroute;
        }
    }
}
