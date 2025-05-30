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
    /// Custom error of an application gateway.
    /// </summary>
    [OutputType]
    public sealed class ApplicationGatewayCustomErrorResponse
    {
        /// <summary>
        /// Error page URL of the application gateway custom error.
        /// </summary>
        public readonly string? CustomErrorPageUrl;
        /// <summary>
        /// Status code of the application gateway custom error.
        /// </summary>
        public readonly string? StatusCode;

        [OutputConstructor]
        private ApplicationGatewayCustomErrorResponse(
            string? customErrorPageUrl,

            string? statusCode)
        {
            CustomErrorPageUrl = customErrorPageUrl;
            StatusCode = statusCode;
        }
    }
}
