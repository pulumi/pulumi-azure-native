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
    /// Instance type schema.
    /// </summary>
    public sealed class InstanceTypeSchemaArgs : global::Pulumi.ResourceArgs
    {
        [Input("nodeSelector")]
        private InputMap<string>? _nodeSelector;

        /// <summary>
        /// Node Selector
        /// </summary>
        public InputMap<string> NodeSelector
        {
            get => _nodeSelector ?? (_nodeSelector = new InputMap<string>());
            set => _nodeSelector = value;
        }

        /// <summary>
        /// Resource requests/limits for this instance type
        /// </summary>
        [Input("resources")]
        public Input<Inputs.InstanceTypeSchemaResourcesArgs>? Resources { get; set; }

        public InstanceTypeSchemaArgs()
        {
        }
        public static new InstanceTypeSchemaArgs Empty => new InstanceTypeSchemaArgs();
    }
}
