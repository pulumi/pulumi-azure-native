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
    /// Settings for a personal compute instance.
    /// </summary>
    public sealed class PersonalComputeInstanceSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A user explicitly assigned to a personal compute instance.
        /// </summary>
        [Input("assignedUser")]
        public Input<Inputs.AssignedUserArgs>? AssignedUser { get; set; }

        public PersonalComputeInstanceSettingsArgs()
        {
        }
        public static new PersonalComputeInstanceSettingsArgs Empty => new PersonalComputeInstanceSettingsArgs();
    }
}
