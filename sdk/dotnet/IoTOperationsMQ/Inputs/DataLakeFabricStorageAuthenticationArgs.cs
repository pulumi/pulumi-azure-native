// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperationsMQ.Inputs
{

    /// <summary>
    /// DataLake Fabric Storage authentication details.
    /// </summary>
    public sealed class DataLakeFabricStorageAuthenticationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for managed identity authentication.
        /// </summary>
        [Input("systemAssignedManagedIdentity", required: true)]
        public Input<Inputs.ManagedIdentityAuthenticationArgs> SystemAssignedManagedIdentity { get; set; } = null!;

        public DataLakeFabricStorageAuthenticationArgs()
        {
        }
        public static new DataLakeFabricStorageAuthenticationArgs Empty => new DataLakeFabricStorageAuthenticationArgs();
    }
}
