// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ApiManagement.Inputs
{

    /// <summary>
    /// Diagnostic settings for Large Language Models
    /// </summary>
    public sealed class LLMDiagnosticSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether default diagnostic should be enabled for Large Language Models or not.
        /// </summary>
        [Input("logs")]
        public InputUnion<string, Pulumi.AzureNative.ApiManagement.LlmDiagnosticSettings>? Logs { get; set; }

        /// <summary>
        /// Diagnostic settings for Large Language Models requests.
        /// </summary>
        [Input("requests")]
        public Input<Inputs.LLMMessageDiagnosticSettingsArgs>? Requests { get; set; }

        /// <summary>
        /// Diagnostic settings for Large Language Models responses.
        /// </summary>
        [Input("responses")]
        public Input<Inputs.LLMMessageDiagnosticSettingsArgs>? Responses { get; set; }

        public LLMDiagnosticSettingsArgs()
        {
        }
        public static new LLMDiagnosticSettingsArgs Empty => new LLMDiagnosticSettingsArgs();
    }
}
