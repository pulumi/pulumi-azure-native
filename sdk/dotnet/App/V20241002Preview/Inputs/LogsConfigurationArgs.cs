// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.App.V20241002Preview.Inputs
{

    /// <summary>
    /// Configuration of Open Telemetry logs
    /// </summary>
    public sealed class LogsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<string>? _destinations;

        /// <summary>
        /// Open telemetry logs destinations
        /// </summary>
        public InputList<string> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<string>());
            set => _destinations = value;
        }

        public LogsConfigurationArgs()
        {
        }
        public static new LogsConfigurationArgs Empty => new LogsConfigurationArgs();
    }
}
