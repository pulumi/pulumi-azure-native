// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.AzureArcData.Inputs
{

    /// <summary>
    /// A list of Availability Group Database Replicas.
    /// </summary>
    public sealed class SqlServerAvailabilityGroupResourcePropertiesDatabasesArgs : global::Pulumi.ResourceArgs
    {
        [Input("value")]
        private InputList<Inputs.SqlAvailabilityGroupDatabaseReplicaResourcePropertiesArgs>? _value;

        /// <summary>
        /// Array of Availability Group Database Replicas.
        /// </summary>
        public InputList<Inputs.SqlAvailabilityGroupDatabaseReplicaResourcePropertiesArgs> Value
        {
            get => _value ?? (_value = new InputList<Inputs.SqlAvailabilityGroupDatabaseReplicaResourcePropertiesArgs>());
            set => _value = value;
        }

        public SqlServerAvailabilityGroupResourcePropertiesDatabasesArgs()
        {
        }
        public static new SqlServerAvailabilityGroupResourcePropertiesDatabasesArgs Empty => new SqlServerAvailabilityGroupResourcePropertiesDatabasesArgs();
    }
}
