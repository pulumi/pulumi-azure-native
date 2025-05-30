// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    /// <summary>
    /// Inference pool configuration
    /// </summary>
    public sealed class InferencePoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("properties")]
        private InputList<Inputs.StringStringKeyValuePairArgs>? _properties;

        /// <summary>
        /// Property dictionary. Properties can be added, but not removed or altered.
        /// </summary>
        public InputList<Inputs.StringStringKeyValuePairArgs> Properties
        {
            get => _properties ?? (_properties = new InputList<Inputs.StringStringKeyValuePairArgs>());
            set => _properties = value;
        }

        /// <summary>
        /// Gets or sets ScaleUnitConfiguration for the inference pool. Used if PoolType=ScaleUnit.
        /// </summary>
        [Input("scaleUnitConfiguration")]
        public Input<Inputs.ScaleUnitConfigurationArgs>? ScaleUnitConfiguration { get; set; }

        public InferencePoolArgs()
        {
        }
        public static new InferencePoolArgs Empty => new InferencePoolArgs();
    }
}
