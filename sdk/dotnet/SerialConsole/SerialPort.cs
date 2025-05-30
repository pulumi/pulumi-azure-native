// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.SerialConsole
{
    /// <summary>
    /// Represents the serial port of the parent resource.
    /// 
    /// Uses Azure REST API version 2018-05-01. In version 2.x of the Azure Native provider, it used API version 2018-05-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:serialconsole:SerialPort")]
    public partial class SerialPort : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the port is enabled for a serial console connection.
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SerialPort resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SerialPort(string name, SerialPortArgs args, CustomResourceOptions? options = null)
            : base("azure-native:serialconsole:SerialPort", name, args ?? new SerialPortArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SerialPort(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:serialconsole:SerialPort", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:serialconsole/v20180501:SerialPort" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SerialPort resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SerialPort Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SerialPort(name, id, options);
        }
    }

    public sealed class SerialPortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource name, or subordinate path, for the parent of the serial port. For example: the name of the virtual machine.
        /// </summary>
        [Input("parentResource", required: true)]
        public Input<string> ParentResource { get; set; } = null!;

        /// <summary>
        /// The resource type of the parent resource.  For example: 'virtualMachines' or 'virtualMachineScaleSets'
        /// </summary>
        [Input("parentResourceType", required: true)]
        public Input<string> ParentResourceType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The namespace of the resource provider.
        /// </summary>
        [Input("resourceProviderNamespace", required: true)]
        public Input<string> ResourceProviderNamespace { get; set; } = null!;

        /// <summary>
        /// The name of the serial port to create.
        /// </summary>
        [Input("serialPort")]
        public Input<string>? SerialPort { get; set; }

        /// <summary>
        /// Specifies whether the port is enabled for a serial console connection.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.AzureNative.SerialConsole.SerialPortState>? State { get; set; }

        public SerialPortArgs()
        {
        }
        public static new SerialPortArgs Empty => new SerialPortArgs();
    }
}
