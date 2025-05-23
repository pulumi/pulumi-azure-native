// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeviceRegistry.Outputs
{

    /// <summary>
    /// Defines the asset endpoint profile status error properties.
    /// </summary>
    [OutputType]
    public sealed class AssetEndpointProfileStatusErrorResponse
    {
        /// <summary>
        /// Error code for classification of errors (ex: 400, 404, 500, etc.).
        /// </summary>
        public readonly int Code;
        /// <summary>
        /// Human readable helpful error message to provide additional context for error (ex: “targetAddress 'foo' is not a valid url”).
        /// </summary>
        public readonly string Message;

        [OutputConstructor]
        private AssetEndpointProfileStatusErrorResponse(
            int code,

            string message)
        {
            Code = code;
            Message = message;
        }
    }
}
