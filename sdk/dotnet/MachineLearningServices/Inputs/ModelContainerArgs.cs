// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.Inputs
{

    public sealed class ModelContainerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The asset description text.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Is the asset archived?
        /// </summary>
        [Input("isArchived")]
        public Input<bool>? IsArchived { get; set; }

        [Input("properties")]
        private InputMap<string>? _properties;

        /// <summary>
        /// The asset property dictionary.
        /// </summary>
        public InputMap<string> Properties
        {
            get => _properties ?? (_properties = new InputMap<string>());
            set => _properties = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tag dictionary. Tags can be added, removed, and updated.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ModelContainerArgs()
        {
            IsArchived = false;
        }
        public static new ModelContainerArgs Empty => new ModelContainerArgs();
    }
}
