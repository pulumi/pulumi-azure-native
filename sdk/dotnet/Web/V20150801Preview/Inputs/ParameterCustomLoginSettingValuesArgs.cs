// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Web.V20150801Preview.Inputs
{

    /// <summary>
    /// Custom logging setting values
    /// </summary>
    public sealed class ParameterCustomLoginSettingValuesArgs : global::Pulumi.ResourceArgs
    {
        [Input("customParameters")]
        private InputMap<Inputs.CustomLoginSettingValueArgs>? _customParameters;

        /// <summary>
        /// Custom parameters.
        /// </summary>
        public InputMap<Inputs.CustomLoginSettingValueArgs> CustomParameters
        {
            get => _customParameters ?? (_customParameters = new InputMap<Inputs.CustomLoginSettingValueArgs>());
            set => _customParameters = value;
        }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Kind of resource
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Resource Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ParameterCustomLoginSettingValuesArgs()
        {
        }
        public static new ParameterCustomLoginSettingValuesArgs Empty => new ParameterCustomLoginSettingValuesArgs();
    }
}
