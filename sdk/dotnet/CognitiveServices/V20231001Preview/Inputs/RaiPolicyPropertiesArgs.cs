// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CognitiveServices.V20231001Preview.Inputs
{

    /// <summary>
    /// Azure OpenAI Content Filters properties.
    /// </summary>
    public sealed class RaiPolicyPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the base Content Filters.
        /// </summary>
        [Input("basePolicyName")]
        public Input<string>? BasePolicyName { get; set; }

        [Input("completionBlocklists")]
        private InputList<Inputs.RaiBlocklistConfigArgs>? _completionBlocklists;

        /// <summary>
        /// The list of blocklists for completion.
        /// </summary>
        public InputList<Inputs.RaiBlocklistConfigArgs> CompletionBlocklists
        {
            get => _completionBlocklists ?? (_completionBlocklists = new InputList<Inputs.RaiBlocklistConfigArgs>());
            set => _completionBlocklists = value;
        }

        [Input("contentFilters")]
        private InputList<Inputs.RaiPolicyContentFilterArgs>? _contentFilters;

        /// <summary>
        /// The list of Content Filters.
        /// </summary>
        public InputList<Inputs.RaiPolicyContentFilterArgs> ContentFilters
        {
            get => _contentFilters ?? (_contentFilters = new InputList<Inputs.RaiPolicyContentFilterArgs>());
            set => _contentFilters = value;
        }

        /// <summary>
        /// Content Filters mode.
        /// </summary>
        [Input("mode")]
        public InputUnion<string, Pulumi.AzureNative.CognitiveServices.V20231001Preview.RaiPolicyMode>? Mode { get; set; }

        [Input("promptBlocklists")]
        private InputList<Inputs.RaiBlocklistConfigArgs>? _promptBlocklists;

        /// <summary>
        /// The list of blocklists for prompt.
        /// </summary>
        public InputList<Inputs.RaiBlocklistConfigArgs> PromptBlocklists
        {
            get => _promptBlocklists ?? (_promptBlocklists = new InputList<Inputs.RaiBlocklistConfigArgs>());
            set => _promptBlocklists = value;
        }

        public RaiPolicyPropertiesArgs()
        {
        }
        public static new RaiPolicyPropertiesArgs Empty => new RaiPolicyPropertiesArgs();
    }
}
