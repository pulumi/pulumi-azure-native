// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBoxEdge
{
    /// <summary>
    /// The limited preview of Kubernetes Cluster Management from the Azure supports:
    /// 1. Using a simple turn-key option in Azure Portal, deploy a Kubernetes cluster on your Azure Stack Edge device.
    /// 2. Configure Kubernetes cluster running on your device with Arc enabled Kubernetes with a click of a button in the Azure Portal.
    ///     Azure Arc enables organizations to view, manage, and govern their on-premises Kubernetes clusters using the Azure Portal, command line tools, and APIs.
    /// 3. Easily configure Persistent Volumes using SMB and NFS shares for storing container data.
    ///     For more information, refer to the document here: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8-Cloud-Management-20210323.pdf
    ///     Or Demo: https://databoxupdatepackages.blob.core.windows.net/documentation/Microsoft-Azure-Stack-Edge-K8S-Cloud-Management-20210323.mp4
    ///     By using this feature, you agree to the preview legal terms. See the https://azure.microsoft.com/en-us/support/legal/preview-supplemental-terms/
    /// 
    /// Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.
    /// </summary>
    [AzureNativeResourceType("azure-native:databoxedge:KubernetesRole")]
    public partial class KubernetesRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Host OS supported by the Kubernetes role.
        /// </summary>
        [Output("hostPlatform")]
        public Output<string> HostPlatform { get; private set; } = null!;

        /// <summary>
        /// Platform where the runtime is hosted.
        /// </summary>
        [Output("hostPlatformType")]
        public Output<string> HostPlatformType { get; private set; } = null!;

        /// <summary>
        /// Role type.
        /// Expected value is 'Kubernetes'.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Kubernetes cluster configuration
        /// </summary>
        [Output("kubernetesClusterInfo")]
        public Output<Outputs.KubernetesClusterInfoResponse> KubernetesClusterInfo { get; private set; } = null!;

        /// <summary>
        /// Kubernetes role resources
        /// </summary>
        [Output("kubernetesRoleResources")]
        public Output<Outputs.KubernetesRoleResourcesResponse> KubernetesRoleResources { get; private set; } = null!;

        /// <summary>
        /// The object name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// State of Kubernetes deployment
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Role status.
        /// </summary>
        [Output("roleStatus")]
        public Output<string> RoleStatus { get; private set; } = null!;

        /// <summary>
        /// Metadata pertaining to creation and last modification of Role
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesRole(string name, KubernetesRoleArgs args, CustomResourceOptions? options = null)
            : base("azure-native:databoxedge:KubernetesRole", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesRole(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:databoxedge:KubernetesRole", name, null, MakeResourceOptions(options, id))
        {
        }

        private static KubernetesRoleArgs MakeArgs(KubernetesRoleArgs args)
        {
            args ??= new KubernetesRoleArgs();
            args.Kind = "Kubernetes";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20190301:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20190701:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20190801:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20200501preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20200901:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20200901preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20201201:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20210201:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20210201preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20210601:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20210601preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20220301:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20220401preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20221201preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230101preview:CloudEdgeManagementRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230101preview:IoTRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230101preview:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230101preview:MECRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230701:CloudEdgeManagementRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230701:IoTRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230701:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20230701:MECRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20231201:CloudEdgeManagementRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20231201:IoTRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20231201:KubernetesRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge/v20231201:MECRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge:CloudEdgeManagementRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge:IoTRole" },
                    new global::Pulumi.Alias { Type = "azure-native:databoxedge:MECRole" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KubernetesRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesRole Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KubernetesRole(name, id, options);
        }
    }

    public sealed class KubernetesRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device name.
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// Host OS supported by the Kubernetes role.
        /// </summary>
        [Input("hostPlatform", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataBoxEdge.PlatformType> HostPlatform { get; set; } = null!;

        /// <summary>
        /// Role type.
        /// Expected value is 'Kubernetes'.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Kubernetes cluster configuration
        /// </summary>
        [Input("kubernetesClusterInfo", required: true)]
        public Input<Inputs.KubernetesClusterInfoArgs> KubernetesClusterInfo { get; set; } = null!;

        /// <summary>
        /// Kubernetes role resources
        /// </summary>
        [Input("kubernetesRoleResources", required: true)]
        public Input<Inputs.KubernetesRoleResourcesArgs> KubernetesRoleResources { get; set; } = null!;

        /// <summary>
        /// The role name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Role status.
        /// </summary>
        [Input("roleStatus", required: true)]
        public InputUnion<string, Pulumi.AzureNative.DataBoxEdge.RoleStatus> RoleStatus { get; set; } = null!;

        public KubernetesRoleArgs()
        {
        }
        public static new KubernetesRoleArgs Empty => new KubernetesRoleArgs();
    }
}
