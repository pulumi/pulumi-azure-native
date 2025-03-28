// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.CostManagement.V20241001Preview.Inputs
{

    /// <summary>
    /// Will contain the filter name and value to operate on. This is currently only supported for Export Definition type of ReservationRecommendations.
    /// </summary>
    public sealed class FilterItemsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the filter. This is currently only supported for Export Definition type of ReservationRecommendations. Supported names are ['ReservationScope', 'LookBackPeriod', 'ResourceType']
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNative.CostManagement.V20241001Preview.FilterItemNames>? Name { get; set; }

        /// <summary>
        /// Value to filter by. Currently values supported per name are, for 'ReservationScope' supported values are ['Single', 'Shared'], for 'LookBackPeriod' supported values are ['Last7Days', 'Last30Days', 'Last60Days'] and for 'ResourceType' supported values are ['VirtualMachines', 'SQLDatabases', 'PostgreSQL', 'ManagedDisk', 'MySQL', 'RedHat', 'MariaDB', 'RedisCache', 'CosmosDB', 'SqlDataWarehouse', 'SUSELinux', 'AppService', 'BlockBlob', 'AzureDataExplorer', 'VMwareCloudSimple'].
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public FilterItemsArgs()
        {
        }
        public static new FilterItemsArgs Empty => new FilterItemsArgs();
    }
}
