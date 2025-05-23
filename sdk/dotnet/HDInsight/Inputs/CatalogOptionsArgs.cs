// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HDInsight.Inputs
{

    /// <summary>
    /// Trino cluster catalog options.
    /// </summary>
    public sealed class CatalogOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("hive")]
        private InputList<Inputs.HiveCatalogOptionArgs>? _hive;

        /// <summary>
        /// hive catalog options.
        /// </summary>
        public InputList<Inputs.HiveCatalogOptionArgs> Hive
        {
            get => _hive ?? (_hive = new InputList<Inputs.HiveCatalogOptionArgs>());
            set => _hive = value;
        }

        public CatalogOptionsArgs()
        {
        }
        public static new CatalogOptionsArgs Empty => new CatalogOptionsArgs();
    }
}
