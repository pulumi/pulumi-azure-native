// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.IoTOperations.V20240915Preview
{
    /// <summary>
    /// Instance broker resource
    /// </summary>
    [AzureNativeResourceType("azure-native:iotoperations/v20240915preview:Broker")]
    public partial class Broker : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Edge location of the resource.
        /// </summary>
        [Output("extendedLocation")]
        public Output<Outputs.ExtendedLocationResponse> ExtendedLocation { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.BrokerPropertiesResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Broker resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Broker(string name, BrokerArgs args, CustomResourceOptions? options = null)
            : base("azure-native:iotoperations/v20240915preview:Broker", name, args ?? new BrokerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Broker(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:iotoperations/v20240915preview:Broker", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:iotoperations/v20240701preview:Broker" },
                    new global::Pulumi.Alias { Type = "azure-native:iotoperations/v20240815preview:Broker" },
                    new global::Pulumi.Alias { Type = "azure-native:iotoperations/v20241101:Broker" },
                    new global::Pulumi.Alias { Type = "azure-native:iotoperations/v20250401:Broker" },
                    new global::Pulumi.Alias { Type = "azure-native:iotoperations:Broker" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Broker resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Broker Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Broker(name, id, options);
        }
    }

    public sealed class BrokerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of broker.
        /// </summary>
        [Input("brokerName")]
        public Input<string>? BrokerName { get; set; }

        /// <summary>
        /// Edge location of the resource.
        /// </summary>
        [Input("extendedLocation", required: true)]
        public Input<Inputs.ExtendedLocationArgs> ExtendedLocation { get; set; } = null!;

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// The resource-specific properties for this resource.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.BrokerPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public BrokerArgs()
        {
        }
        public static new BrokerArgs Empty => new BrokerArgs();
    }
}
