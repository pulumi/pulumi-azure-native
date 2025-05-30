// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    /// <summary>
    /// Information about packet capture session.
    /// 
    /// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2023-02-01.
    /// 
    /// Other available API versions: 2018-06-01, 2018-07-01, 2018-08-01, 2018-10-01, 2018-11-01, 2018-12-01, 2019-02-01, 2019-04-01, 2019-06-01, 2019-07-01, 2019-08-01, 2019-09-01, 2019-11-01, 2019-12-01, 2020-03-01, 2020-04-01, 2020-05-01, 2020-06-01, 2020-07-01, 2020-08-01, 2020-11-01, 2021-02-01, 2021-03-01, 2021-05-01, 2021-08-01, 2022-01-01, 2022-05-01, 2022-07-01, 2022-09-01, 2022-11-01, 2023-02-01, 2023-04-01, 2023-05-01, 2023-06-01, 2023-09-01, 2023-11-01, 2024-01-01, 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:network:PacketCapture")]
    public partial class PacketCapture : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Number of bytes captured per packet, the remaining bytes are truncated.
        /// </summary>
        [Output("bytesToCapturePerPacket")]
        public Output<double?> BytesToCapturePerPacket { get; private set; } = null!;

        /// <summary>
        /// The capture setting holds the 'FileCount', 'FileSizeInBytes', 'SessionTimeLimitInSeconds' values.
        /// </summary>
        [Output("captureSettings")]
        public Output<Outputs.PacketCaptureSettingsResponse?> CaptureSettings { get; private set; } = null!;

        /// <summary>
        /// This continuous capture is a nullable boolean, which can hold 'null', 'true' or 'false' value. If we do not pass this parameter, it would be consider as 'null', default value is 'null'.
        /// </summary>
        [Output("continuousCapture")]
        public Output<bool?> ContinuousCapture { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A list of packet capture filters.
        /// </summary>
        [Output("filters")]
        public Output<ImmutableArray<Outputs.PacketCaptureFilterResponse>> Filters { get; private set; } = null!;

        /// <summary>
        /// Name of the packet capture session.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the packet capture session.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// A list of AzureVMSS instances which can be included or excluded to run packet capture. If both included and excluded are empty, then the packet capture will run on all instances of AzureVMSS.
        /// </summary>
        [Output("scope")]
        public Output<Outputs.PacketCaptureMachineScopeResponse?> Scope { get; private set; } = null!;

        /// <summary>
        /// The storage location for a packet capture session.
        /// </summary>
        [Output("storageLocation")]
        public Output<Outputs.PacketCaptureStorageLocationResponse> StorageLocation { get; private set; } = null!;

        /// <summary>
        /// The ID of the targeted resource, only AzureVM and AzureVMSS as target type are currently supported.
        /// </summary>
        [Output("target")]
        public Output<string> Target { get; private set; } = null!;

        /// <summary>
        /// Target type of the resource provided.
        /// </summary>
        [Output("targetType")]
        public Output<string?> TargetType { get; private set; } = null!;

        /// <summary>
        /// Maximum duration of the capture session in seconds.
        /// </summary>
        [Output("timeLimitInSeconds")]
        public Output<int?> TimeLimitInSeconds { get; private set; } = null!;

        /// <summary>
        /// Maximum size of the capture output.
        /// </summary>
        [Output("totalBytesPerSession")]
        public Output<double?> TotalBytesPerSession { get; private set; } = null!;


        /// <summary>
        /// Create a PacketCapture resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PacketCapture(string name, PacketCaptureArgs args, CustomResourceOptions? options = null)
            : base("azure-native:network:PacketCapture", name, args ?? new PacketCaptureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PacketCapture(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:network:PacketCapture", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:network/v20160901:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20161201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170301:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170601:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170801:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20170901:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20171001:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20171101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180401:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180601:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180701:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20180801:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181001:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20181201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190401:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190601:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190701:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190801:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20190901:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20191201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200301:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200401:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200501:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200601:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200701:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20200801:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20201101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210301:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210501:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20210801:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220501:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220701:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20220901:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20221101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230201:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230401:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230501:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230601:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20230901:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20231101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240101:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240301:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240501:PacketCapture" },
                    new global::Pulumi.Alias { Type = "azure-native:network/v20240701:PacketCapture" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PacketCapture resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PacketCapture Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PacketCapture(name, id, options);
        }
    }

    public sealed class PacketCaptureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of bytes captured per packet, the remaining bytes are truncated.
        /// </summary>
        [Input("bytesToCapturePerPacket")]
        public Input<double>? BytesToCapturePerPacket { get; set; }

        /// <summary>
        /// The capture setting holds the 'FileCount', 'FileSizeInBytes', 'SessionTimeLimitInSeconds' values.
        /// </summary>
        [Input("captureSettings")]
        public Input<Inputs.PacketCaptureSettingsArgs>? CaptureSettings { get; set; }

        /// <summary>
        /// This continuous capture is a nullable boolean, which can hold 'null', 'true' or 'false' value. If we do not pass this parameter, it would be consider as 'null', default value is 'null'.
        /// </summary>
        [Input("continuousCapture")]
        public Input<bool>? ContinuousCapture { get; set; }

        [Input("filters")]
        private InputList<Inputs.PacketCaptureFilterArgs>? _filters;

        /// <summary>
        /// A list of packet capture filters.
        /// </summary>
        public InputList<Inputs.PacketCaptureFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.PacketCaptureFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the packet capture session.
        /// </summary>
        [Input("packetCaptureName")]
        public Input<string>? PacketCaptureName { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// A list of AzureVMSS instances which can be included or excluded to run packet capture. If both included and excluded are empty, then the packet capture will run on all instances of AzureVMSS.
        /// </summary>
        [Input("scope")]
        public Input<Inputs.PacketCaptureMachineScopeArgs>? Scope { get; set; }

        /// <summary>
        /// The storage location for a packet capture session.
        /// </summary>
        [Input("storageLocation", required: true)]
        public Input<Inputs.PacketCaptureStorageLocationArgs> StorageLocation { get; set; } = null!;

        /// <summary>
        /// The ID of the targeted resource, only AzureVM and AzureVMSS as target type are currently supported.
        /// </summary>
        [Input("target", required: true)]
        public Input<string> Target { get; set; } = null!;

        /// <summary>
        /// Target type of the resource provided.
        /// </summary>
        [Input("targetType")]
        public Input<Pulumi.AzureNative.Network.PacketCaptureTargetType>? TargetType { get; set; }

        /// <summary>
        /// Maximum duration of the capture session in seconds.
        /// </summary>
        [Input("timeLimitInSeconds")]
        public Input<int>? TimeLimitInSeconds { get; set; }

        /// <summary>
        /// Maximum size of the capture output.
        /// </summary>
        [Input("totalBytesPerSession")]
        public Input<double>? TotalBytesPerSession { get; set; }

        public PacketCaptureArgs()
        {
            BytesToCapturePerPacket = 0;
            TimeLimitInSeconds = 18000;
            TotalBytesPerSession = 1073741824;
        }
        public static new PacketCaptureArgs Empty => new PacketCaptureArgs();
    }
}
