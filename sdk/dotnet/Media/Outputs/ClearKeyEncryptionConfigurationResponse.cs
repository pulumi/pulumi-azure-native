// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Media.Outputs
{

    /// <summary>
    /// Class to specify ClearKey configuration of common encryption schemes in Streaming Policy
    /// </summary>
    [OutputType]
    public sealed class ClearKeyEncryptionConfigurationResponse
    {
        /// <summary>
        /// Template for the URL of the custom service delivering content keys to end user players. Not required when using Azure Media Services for issuing licenses. The template supports replaceable tokens that the service will update at runtime with the value specific to the request.  The currently supported token value is {AlternativeMediaId}, which is replaced with the value of StreamingLocatorId.AlternativeMediaId.
        /// </summary>
        public readonly string? CustomKeysAcquisitionUrlTemplate;

        [OutputConstructor]
        private ClearKeyEncryptionConfigurationResponse(string? customKeysAcquisitionUrlTemplate)
        {
            CustomKeysAcquisitionUrlTemplate = customKeysAcquisitionUrlTemplate;
        }
    }
}
