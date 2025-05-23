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
    /// Azure ML Update Resource management activity.
    /// </summary>
    public sealed class AzureMLUpdateResourceActivityArgs : global::Pulumi.ResourceArgs
    {
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
        /// Linked service reference.
        /// </summary>
        [Input("linkedServiceName")]
        public Input<Inputs.LinkedServiceReferenceArgs>? LinkedServiceName { get; set; }

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
        /// Activity policy.
        /// </summary>
        [Input("policy")]
        public Input<Inputs.ActivityPolicyArgs>? Policy { get; set; }

        /// <summary>
        /// Activity state. This is an optional property and if not provided, the state will be Active by default.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNative.DataFactory.ActivityState>? State { get; set; }

        /// <summary>
        /// The relative file path in trainedModelLinkedService to represent the .ilearner file that will be uploaded by the update operation.  Type: string (or Expression with resultType string).
        /// </summary>
        [Input("trainedModelFilePath", required: true)]
        public Input<object> TrainedModelFilePath { get; set; } = null!;

        /// <summary>
        /// Name of Azure Storage linked service holding the .ilearner file that will be uploaded by the update operation.
        /// </summary>
        [Input("trainedModelLinkedServiceName", required: true)]
        public Input<Inputs.LinkedServiceReferenceArgs> TrainedModelLinkedServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the Trained Model module in the Web Service experiment to be updated. Type: string (or Expression with resultType string).
        /// </summary>
        [Input("trainedModelName", required: true)]
        public Input<object> TrainedModelName { get; set; } = null!;

        /// <summary>
        /// Type of activity.
        /// Expected value is 'AzureMLUpdateResource'.
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

        public AzureMLUpdateResourceActivityArgs()
        {
        }
        public static new AzureMLUpdateResourceActivityArgs Empty => new AzureMLUpdateResourceActivityArgs();
    }
}
