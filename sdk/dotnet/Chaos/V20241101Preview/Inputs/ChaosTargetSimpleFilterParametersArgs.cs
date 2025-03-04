// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Chaos.V20241101Preview.Inputs
{

    /// <summary>
    /// Model that represents the Simple filter parameters.
    /// </summary>
    public sealed class ChaosTargetSimpleFilterParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// List of Azure availability zones to filter targets by.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public ChaosTargetSimpleFilterParametersArgs()
        {
        }
        public static new ChaosTargetSimpleFilterParametersArgs Empty => new ChaosTargetSimpleFilterParametersArgs();
    }
}
