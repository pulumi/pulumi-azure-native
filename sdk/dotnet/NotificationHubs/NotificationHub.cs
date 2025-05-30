// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs
{
    /// <summary>
    /// Notification Hub Resource.
    /// 
    /// Uses Azure REST API version 2023-10-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-01-01-preview.
    /// 
    /// Other available API versions: 2023-01-01-preview, 2023-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native notificationhubs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:notificationhubs:NotificationHub")]
    public partial class NotificationHub : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description of a NotificationHub AdmCredential.
        /// </summary>
        [Output("admCredential")]
        public Output<Outputs.AdmCredentialResponse?> AdmCredential { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub ApnsCredential.
        /// </summary>
        [Output("apnsCredential")]
        public Output<Outputs.ApnsCredentialResponse?> ApnsCredential { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the AuthorizationRules of the created NotificationHub
        /// </summary>
        [Output("authorizationRules")]
        public Output<ImmutableArray<Outputs.SharedAccessAuthorizationRulePropertiesResponse>> AuthorizationRules { get; private set; } = null!;

        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub BaiduCredential.
        /// </summary>
        [Output("baiduCredential")]
        public Output<Outputs.BaiduCredentialResponse?> BaiduCredential { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub BrowserCredential.
        /// </summary>
        [Output("browserCredential")]
        public Output<Outputs.BrowserCredentialResponse?> BrowserCredential { get; private set; } = null!;

        [Output("dailyMaxActiveDevices")]
        public Output<double> DailyMaxActiveDevices { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub FcmV1Credential.
        /// </summary>
        [Output("fcmV1Credential")]
        public Output<Outputs.FcmV1CredentialResponse?> FcmV1Credential { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub GcmCredential.
        /// </summary>
        [Output("gcmCredential")]
        public Output<Outputs.GcmCredentialResponse?> GcmCredential { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub MpnsCredential.
        /// </summary>
        [Output("mpnsCredential")]
        public Output<Outputs.MpnsCredentialResponse?> MpnsCredential { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the RegistrationTtl of the created NotificationHub
        /// </summary>
        [Output("registrationTtl")]
        public Output<string?> RegistrationTtl { get; private set; } = null!;

        /// <summary>
        /// The Sku description for a namespace
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Azure Resource Manager metadata containing createdBy and modifiedBy information.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub WnsCredential.
        /// </summary>
        [Output("wnsCredential")]
        public Output<Outputs.WnsCredentialResponse?> WnsCredential { get; private set; } = null!;

        /// <summary>
        /// Description of a NotificationHub XiaomiCredential.
        /// </summary>
        [Output("xiaomiCredential")]
        public Output<Outputs.XiaomiCredentialResponse?> XiaomiCredential { get; private set; } = null!;


        /// <summary>
        /// Create a NotificationHub resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotificationHub(string name, NotificationHubArgs args, CustomResourceOptions? options = null)
            : base("azure-native:notificationhubs:NotificationHub", name, args ?? new NotificationHubArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotificationHub(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:notificationhubs:NotificationHub", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:notificationhubs/v20140901:NotificationHub" },
                    new global::Pulumi.Alias { Type = "azure-native:notificationhubs/v20160301:NotificationHub" },
                    new global::Pulumi.Alias { Type = "azure-native:notificationhubs/v20170401:NotificationHub" },
                    new global::Pulumi.Alias { Type = "azure-native:notificationhubs/v20230101preview:NotificationHub" },
                    new global::Pulumi.Alias { Type = "azure-native:notificationhubs/v20230901:NotificationHub" },
                    new global::Pulumi.Alias { Type = "azure-native:notificationhubs/v20231001preview:NotificationHub" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NotificationHub resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotificationHub Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NotificationHub(name, id, options);
        }
    }

    public sealed class NotificationHubArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of a NotificationHub AdmCredential.
        /// </summary>
        [Input("admCredential")]
        public Input<Inputs.AdmCredentialArgs>? AdmCredential { get; set; }

        /// <summary>
        /// Description of a NotificationHub ApnsCredential.
        /// </summary>
        [Input("apnsCredential")]
        public Input<Inputs.ApnsCredentialArgs>? ApnsCredential { get; set; }

        /// <summary>
        /// Description of a NotificationHub BaiduCredential.
        /// </summary>
        [Input("baiduCredential")]
        public Input<Inputs.BaiduCredentialArgs>? BaiduCredential { get; set; }

        /// <summary>
        /// Description of a NotificationHub BrowserCredential.
        /// </summary>
        [Input("browserCredential")]
        public Input<Inputs.BrowserCredentialArgs>? BrowserCredential { get; set; }

        /// <summary>
        /// Description of a NotificationHub FcmV1Credential.
        /// </summary>
        [Input("fcmV1Credential")]
        public Input<Inputs.FcmV1CredentialArgs>? FcmV1Credential { get; set; }

        /// <summary>
        /// Description of a NotificationHub GcmCredential.
        /// </summary>
        [Input("gcmCredential")]
        public Input<Inputs.GcmCredentialArgs>? GcmCredential { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Description of a NotificationHub MpnsCredential.
        /// </summary>
        [Input("mpnsCredential")]
        public Input<Inputs.MpnsCredentialArgs>? MpnsCredential { get; set; }

        /// <summary>
        /// Gets or sets the NotificationHub name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Notification Hub name
        /// </summary>
        [Input("notificationHubName")]
        public Input<string>? NotificationHubName { get; set; }

        /// <summary>
        /// Gets or sets the RegistrationTtl of the created NotificationHub
        /// </summary>
        [Input("registrationTtl")]
        public Input<string>? RegistrationTtl { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Sku description for a namespace
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Description of a NotificationHub WnsCredential.
        /// </summary>
        [Input("wnsCredential")]
        public Input<Inputs.WnsCredentialArgs>? WnsCredential { get; set; }

        /// <summary>
        /// Description of a NotificationHub XiaomiCredential.
        /// </summary>
        [Input("xiaomiCredential")]
        public Input<Inputs.XiaomiCredentialArgs>? XiaomiCredential { get; set; }

        public NotificationHubArgs()
        {
        }
        public static new NotificationHubArgs Empty => new NotificationHubArgs();
    }
}
