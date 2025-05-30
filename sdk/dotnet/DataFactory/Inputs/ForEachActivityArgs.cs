// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataFactory.Inputs
{

    /// <summary>
    /// This activity is used for iterating over a collection and execute given activities.
    /// </summary>
    public sealed class ForEachActivityArgs : global::Pulumi.ResourceArgs
    {
        [Input("activities", required: true)]
        private InputList<object>? _activities;

        /// <summary>
        /// List of activities to execute .
        /// </summary>
        public InputList<object> Activities
        {
            get => _activities ?? (_activities = new InputList<object>());
            set => _activities = value;
        }

        /// <summary>
        /// Batch count to be used for controlling the number of parallel execution (when isSequential is set to false).
        /// </summary>
        [Input("batchCount")]
        public Input<int>? BatchCount { get; set; }

        [Input("dependsOn")]
        private InputList<Inputs.ActivityDependencyArgs>? _dependsOn;

        /// <summary>
        /// Activity depends on condition.
        /// </summary>
        public InputList<Inputs.ActivityDependencyArgs> DependsOn
        {
            get => _dependsOn ?? (_dependsOn = new InputList<Inputs.ActivityDependencyArgs>());
            set => _dependsOn = value;
        }

        /// <summary>
        /// Activity description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Should the loop be executed in sequence or in parallel (max 50)
        /// </summary>
        [Input("isSequential")]
        public Input<bool>? IsSequential { get; set; }

        /// <summary>
        /// Collection to iterate.
        /// </summary>
        [Input("items", required: true)]
        public Input<Inputs.ExpressionArgs> Items { get; set; } = null!;

        /// <summary>
        /// Activity name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Status result of the activity when the state is set to Inactive. This is an optional property and if not provided when the activity is inactive, the status will be Succeeded by default.
        /// </summary>
        [Input("onInactiveMarkAs")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.ActivityOnInactiveMarkAs>? OnInactiveMarkAs { get; set; }

        /// <summary>
        /// Activity state. This is an optional property and if not provided, the state will be Active by default.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.ActivityState>? State { get; set; }

        /// <summary>
        /// Type of activity.
        /// Expected value is 'ForEach'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("userProperties")]
        private InputList<Inputs.UserPropertyArgs>? _userProperties;

        /// <summary>
        /// Activity user properties.
        /// </summary>
        public InputList<Inputs.UserPropertyArgs> UserProperties
        {
            get => _userProperties ?? (_userProperties = new InputList<Inputs.UserPropertyArgs>());
            set => _userProperties = value;
        }

        public ForEachActivityArgs()
        {
        }
        public static new ForEachActivityArgs Empty => new ForEachActivityArgs();
    }
}
