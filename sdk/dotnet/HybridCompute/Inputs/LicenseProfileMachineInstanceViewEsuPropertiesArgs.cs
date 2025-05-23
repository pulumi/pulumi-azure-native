// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HybridCompute.Inputs
{

    /// <summary>
    /// Properties for the Machine ESU profile.
    /// </summary>
    public sealed class LicenseProfileMachineInstanceViewEsuPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The assigned license resource.
        /// </summary>
        [Input("assignedLicense")]
        public Input<Inputs.LicenseArgs>? AssignedLicense { get; set; }

        /// <summary>
        /// Describes the license assignment state (Assigned or NotAssigned).
        /// </summary>
        [Input("licenseAssignmentState")]
        public InputUnion<string, Pulumi.AzureNative.HybridCompute.LicenseAssignmentState>? LicenseAssignmentState { get; set; }

        public LicenseProfileMachineInstanceViewEsuPropertiesArgs()
        {
        }
        public static new LicenseProfileMachineInstanceViewEsuPropertiesArgs Empty => new LicenseProfileMachineInstanceViewEsuPropertiesArgs();
    }
}
