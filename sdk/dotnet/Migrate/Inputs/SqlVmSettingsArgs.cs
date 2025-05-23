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
    /// SQL VM assessment settings.
    /// </summary>
    public sealed class SqlVmSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceSeries")]
        private InputList<Union<string, Pulumi.AzureNative.Migrate.AzureVmFamily>>? _instanceSeries;

        /// <summary>
        /// Gets or sets the Azure VM families (calling instance series to keep it
        /// consistent with other targets).
        /// </summary>
        public InputList<Union<string, Pulumi.AzureNative.Migrate.AzureVmFamily>> InstanceSeries
        {
            get => _instanceSeries ?? (_instanceSeries = new InputList<Union<string, Pulumi.AzureNative.Migrate.AzureVmFamily>>());
            set => _instanceSeries = value;
        }

        public SqlVmSettingsArgs()
        {
        }
        public static new SqlVmSettingsArgs Empty => new SqlVmSettingsArgs();
    }
}
