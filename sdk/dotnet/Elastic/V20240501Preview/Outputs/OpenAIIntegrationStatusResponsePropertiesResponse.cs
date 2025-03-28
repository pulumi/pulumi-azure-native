// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Elastic.V20240501Preview.Outputs
{

    /// <summary>
    /// Status of the OpenAI Integration
    /// </summary>
    [OutputType]
    public sealed class OpenAIIntegrationStatusResponsePropertiesResponse
    {
        /// <summary>
        /// Status of the OpenAI Integration
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private OpenAIIntegrationStatusResponsePropertiesResponse(string? status)
        {
            Status = status;
        }
    }
}
