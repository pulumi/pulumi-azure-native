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
    /// Facility settings.
    /// </summary>
    public sealed class FacilitySettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The facilities cost.
        /// </summary>
        [Input("facilitiesCostPerKwh")]
        public Input<double>? FacilitiesCostPerKwh { get; set; }

        public FacilitySettingsArgs()
        {
        }
        public static new FacilitySettingsArgs Empty => new FacilitySettingsArgs();
    }
}
