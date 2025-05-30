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
    /// Diagnostic settings for Large Language Models
    /// </summary>
    [OutputType]
    public sealed class LLMDiagnosticSettingsResponse
    {
        /// <summary>
        /// Specifies whether default diagnostic should be enabled for Large Language Models or not.
        /// </summary>
        public readonly string? Logs;
        /// <summary>
        /// Diagnostic settings for Large Language Models requests.
        /// </summary>
        public readonly Outputs.LLMMessageDiagnosticSettingsResponse? Requests;
        /// <summary>
        /// Diagnostic settings for Large Language Models responses.
        /// </summary>
        public readonly Outputs.LLMMessageDiagnosticSettingsResponse? Responses;

        [OutputConstructor]
        private LLMDiagnosticSettingsResponse(
            string? logs,

            Outputs.LLMMessageDiagnosticSettingsResponse? requests,

            Outputs.LLMMessageDiagnosticSettingsResponse? responses)
        {
            Logs = logs;
            Requests = requests;
            Responses = responses;
        }
    }
}
