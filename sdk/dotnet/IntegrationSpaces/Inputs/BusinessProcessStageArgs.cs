// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IntegrationSpaces.Inputs
{

    /// <summary>
    /// The properties of business process stage.
    /// </summary>
    public sealed class BusinessProcessStageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the business stage.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// The properties within the properties of the business process stage.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("stagesBefore")]
        private InputList<string>? _stagesBefore;

        /// <summary>
        /// The property to keep track of stages before current in the business process stage.
        /// </summary>
        public InputList<string> StagesBefore
        {
            get => _stagesBefore ?? (_stagesBefore = new InputList<string>());
            set => _stagesBefore = value;
        }

        public BusinessProcessStageArgs()
        {
        }
        public static new BusinessProcessStageArgs Empty => new BusinessProcessStageArgs();
    }
}
