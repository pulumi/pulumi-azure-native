// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    public sealed class SubscriberSettingArgs : global::Pulumi.ResourceArgs
    {
        [Input("filterRules")]
        private InputList<Inputs.FilterRuleArgs>? _filterRules;

        /// <summary>
        /// The filter rules.
        /// </summary>
        public InputList<Inputs.FilterRuleArgs> FilterRules
        {
            get => _filterRules ?? (_filterRules = new InputList<Inputs.FilterRuleArgs>());
            set => _filterRules = value;
        }

        public SubscriberSettingArgs()
        {
        }
        public static new SubscriberSettingArgs Empty => new SubscriberSettingArgs();
    }
}
