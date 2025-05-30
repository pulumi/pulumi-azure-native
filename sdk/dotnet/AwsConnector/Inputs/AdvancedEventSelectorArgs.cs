// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AwsConnector.Inputs
{

    /// <summary>
    /// Definition of AdvancedEventSelector
    /// </summary>
    public sealed class AdvancedEventSelectorArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldSelectors")]
        private InputList<Inputs.AdvancedFieldSelectorArgs>? _fieldSelectors;

        /// <summary>
        /// Contains all selector statements in an advanced event selector.
        /// </summary>
        public InputList<Inputs.AdvancedFieldSelectorArgs> FieldSelectors
        {
            get => _fieldSelectors ?? (_fieldSelectors = new InputList<Inputs.AdvancedFieldSelectorArgs>());
            set => _fieldSelectors = value;
        }

        /// <summary>
        /// An optional, descriptive name for an advanced event selector, such as 'Log data events for only two S3 buckets'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public AdvancedEventSelectorArgs()
        {
        }
        public static new AdvancedEventSelectorArgs Empty => new AdvancedEventSelectorArgs();
    }
}
