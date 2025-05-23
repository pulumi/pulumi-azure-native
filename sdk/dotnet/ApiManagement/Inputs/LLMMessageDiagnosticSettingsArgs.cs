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
    /// Diagnostic settings for Large Language Models Messages
    /// </summary>
    public sealed class LLMMessageDiagnosticSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum size of message to logs in bytes. The default size is 32KB.
        /// </summary>
        [Input("maxSizeInBytes")]
        public Input<int>? MaxSizeInBytes { get; set; }

        /// <summary>
        /// Specifies which message should be logged. Currently there is only 'all' option.
        /// </summary>
        [Input("messages")]
        public InputUnion<string, Pulumi.AzureNative.ApiManagement.LlmMessageLogTypes>? Messages { get; set; }

        public LLMMessageDiagnosticSettingsArgs()
        {
        }
        public static new LLMMessageDiagnosticSettingsArgs Empty => new LLMMessageDiagnosticSettingsArgs();
    }
}
