// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Maintenance.Inputs
{

    /// <summary>
    /// Azure query for the update configuration.
    /// </summary>
    public sealed class ConfigurationAssignmentFilterPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("locations")]
        private InputList<string>? _locations;

        /// <summary>
        /// List of locations to scope the query to.
        /// </summary>
        public InputList<string> Locations
        {
            get => _locations ?? (_locations = new InputList<string>());
            set => _locations = value;
        }

        [Input("osTypes")]
        private InputList<string>? _osTypes;

        /// <summary>
        /// List of allowed operating systems.
        /// </summary>
        public InputList<string> OsTypes
        {
            get => _osTypes ?? (_osTypes = new InputList<string>());
            set => _osTypes = value;
        }

        [Input("resourceGroups")]
        private InputList<string>? _resourceGroups;

        /// <summary>
        /// List of allowed resource groups.
        /// </summary>
        public InputList<string> ResourceGroups
        {
            get => _resourceGroups ?? (_resourceGroups = new InputList<string>());
            set => _resourceGroups = value;
        }

        [Input("resourceTypes")]
        private InputList<string>? _resourceTypes;

        /// <summary>
        /// List of allowed resources.
        /// </summary>
        public InputList<string> ResourceTypes
        {
            get => _resourceTypes ?? (_resourceTypes = new InputList<string>());
            set => _resourceTypes = value;
        }

        /// <summary>
        /// Tag settings for the VM.
        /// </summary>
        [Input("tagSettings")]
        public Input<Inputs.TagSettingsPropertiesArgs>? TagSettings { get; set; }

        public ConfigurationAssignmentFilterPropertiesArgs()
        {
        }
        public static new ConfigurationAssignmentFilterPropertiesArgs Empty => new ConfigurationAssignmentFilterPropertiesArgs();
    }
}
