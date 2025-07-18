// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.ProviderHub.Inputs
{

    /// <summary>
    /// The status.
    /// </summary>
    public sealed class CustomRolloutPropertiesStatusArgs : global::Pulumi.ResourceArgs
    {
        [Input("completedRegions")]
        private InputList<string>? _completedRegions;

        /// <summary>
        /// The completed regions.
        /// </summary>
        public InputList<string> CompletedRegions
        {
            get => _completedRegions ?? (_completedRegions = new InputList<string>());
            set => _completedRegions = value;
        }

        [Input("failedOrSkippedRegions")]
        private InputMap<Inputs.ExtendedErrorInfoArgs>? _failedOrSkippedRegions;

        /// <summary>
        /// The failed or skipped regions.
        /// </summary>
        public InputMap<Inputs.ExtendedErrorInfoArgs> FailedOrSkippedRegions
        {
            get => _failedOrSkippedRegions ?? (_failedOrSkippedRegions = new InputMap<Inputs.ExtendedErrorInfoArgs>());
            set => _failedOrSkippedRegions = value;
        }

        /// <summary>
        /// The manifest checkin status.
        /// </summary>
        [Input("manifestCheckinStatus")]
        public Input<Inputs.CustomRolloutStatusManifestCheckinStatusArgs>? ManifestCheckinStatus { get; set; }

        public CustomRolloutPropertiesStatusArgs()
        {
        }
        public static new CustomRolloutPropertiesStatusArgs Empty => new CustomRolloutPropertiesStatusArgs();
    }
}
