// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Migrate.Inputs
{

    /// <summary>
    /// Defines the move resource properties.
    /// </summary>
    public sealed class MoveResourcePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("dependsOnOverrides")]
        private InputList<Inputs.MoveResourceDependencyOverrideArgs>? _dependsOnOverrides;

        /// <summary>
        /// Gets or sets the move resource dependencies overrides.
        /// </summary>
        public InputList<Inputs.MoveResourceDependencyOverrideArgs> DependsOnOverrides
        {
            get => _dependsOnOverrides ?? (_dependsOnOverrides = new InputList<Inputs.MoveResourceDependencyOverrideArgs>());
            set => _dependsOnOverrides = value;
        }

        /// <summary>
        /// Gets or sets the existing target ARM Id of the resource.
        /// </summary>
        [Input("existingTargetId")]
        public Input<string>? ExistingTargetId { get; set; }

        /// <summary>
        /// Gets or sets the resource settings.
        /// </summary>
        [Input("resourceSettings")]
        public object? ResourceSettings { get; set; }

        /// <summary>
        /// Gets or sets the Source ARM Id of the resource.
        /// </summary>
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public MoveResourcePropertiesArgs()
        {
        }
        public static new MoveResourcePropertiesArgs Empty => new MoveResourcePropertiesArgs();
    }
}
