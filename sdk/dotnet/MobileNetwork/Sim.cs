// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MobileNetwork
{
    /// <summary>
    /// SIM resource.
    /// 
    /// Uses Azure REST API version 2024-04-01. In version 2.x of the Azure Native provider, it used API version 2023-06-01.
    /// 
    /// Other available API versions: 2022-04-01-preview, 2022-11-01, 2023-06-01, 2023-09-01, 2024-02-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native mobilenetwork [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
    /// </summary>
    [AzureNativeResourceType("azure-native:mobilenetwork:Sim")]
    public partial class Sim : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Azure API version of the resource.
        /// </summary>
        [Output("azureApiVersion")]
        public Output<string> AzureApiVersion { get; private set; } = null!;

        /// <summary>
        /// An optional free-form text field that can be used to record the device type this SIM is associated with, for example 'Video camera'. The Azure portal allows SIMs to be grouped and filtered based on this value.
        /// </summary>
        [Output("deviceType")]
        public Output<string?> DeviceType { get; private set; } = null!;

        /// <summary>
        /// The integrated circuit card ID (ICCID) for the SIM.
        /// </summary>
        [Output("integratedCircuitCardIdentifier")]
        public Output<string?> IntegratedCircuitCardIdentifier { get; private set; } = null!;

        /// <summary>
        /// The international mobile subscriber identity (IMSI) for the SIM.
        /// </summary>
        [Output("internationalMobileSubscriberIdentity")]
        public Output<string> InternationalMobileSubscriberIdentity { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the SIM resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The SIM policy used by this SIM. The SIM policy must be in the same location as the SIM.
        /// </summary>
        [Output("simPolicy")]
        public Output<Outputs.SimPolicyResourceIdResponse?> SimPolicy { get; private set; } = null!;

        /// <summary>
        /// The state of the SIM resource.
        /// </summary>
        [Output("simState")]
        public Output<string> SimState { get; private set; } = null!;

        /// <summary>
        /// A dictionary of sites to the provisioning state of this SIM on that site.
        /// </summary>
        [Output("siteProvisioningState")]
        public Output<ImmutableDictionary<string, string>> SiteProvisioningState { get; private set; } = null!;

        /// <summary>
        /// A list of static IP addresses assigned to this SIM. Each address is assigned at a defined network scope, made up of {attached data network, slice}.
        /// </summary>
        [Output("staticIpConfiguration")]
        public Output<ImmutableArray<Outputs.SimStaticIpPropertiesResponse>> StaticIpConfiguration { get; private set; } = null!;

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
        /// The public key fingerprint of the SIM vendor who provided this SIM, if any.
        /// </summary>
        [Output("vendorKeyFingerprint")]
        public Output<string> VendorKeyFingerprint { get; private set; } = null!;

        /// <summary>
        /// The name of the SIM vendor who provided this SIM, if any.
        /// </summary>
        [Output("vendorName")]
        public Output<string> VendorName { get; private set; } = null!;


        /// <summary>
        /// Create a Sim resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Sim(string name, SimArgs args, CustomResourceOptions? options = null)
            : base("azure-native:mobilenetwork:Sim", name, args ?? new SimArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Sim(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-native:mobilenetwork:Sim", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20220301preview:Sim" },
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20220401preview:Sim" },
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20221101:Sim" },
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20230601:Sim" },
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20230901:Sim" },
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20240201:Sim" },
                    new global::Pulumi.Alias { Type = "azure-native:mobilenetwork/v20240401:Sim" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Sim resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Sim Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Sim(name, id, options);
        }
    }

    public sealed class SimArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Ki value for the SIM.
        /// </summary>
        [Input("authenticationKey")]
        public Input<string>? AuthenticationKey { get; set; }

        /// <summary>
        /// An optional free-form text field that can be used to record the device type this SIM is associated with, for example 'Video camera'. The Azure portal allows SIMs to be grouped and filtered based on this value.
        /// </summary>
        [Input("deviceType")]
        public Input<string>? DeviceType { get; set; }

        /// <summary>
        /// The integrated circuit card ID (ICCID) for the SIM.
        /// </summary>
        [Input("integratedCircuitCardIdentifier")]
        public Input<string>? IntegratedCircuitCardIdentifier { get; set; }

        /// <summary>
        /// The international mobile subscriber identity (IMSI) for the SIM.
        /// </summary>
        [Input("internationalMobileSubscriberIdentity", required: true)]
        public Input<string> InternationalMobileSubscriberIdentity { get; set; } = null!;

        /// <summary>
        /// The Opc value for the SIM.
        /// </summary>
        [Input("operatorKeyCode")]
        public Input<string>? OperatorKeyCode { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SIM Group.
        /// </summary>
        [Input("simGroupName", required: true)]
        public Input<string> SimGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SIM.
        /// </summary>
        [Input("simName")]
        public Input<string>? SimName { get; set; }

        /// <summary>
        /// The SIM policy used by this SIM. The SIM policy must be in the same location as the SIM.
        /// </summary>
        [Input("simPolicy")]
        public Input<Inputs.SimPolicyResourceIdArgs>? SimPolicy { get; set; }

        [Input("staticIpConfiguration")]
        private InputList<Inputs.SimStaticIpPropertiesArgs>? _staticIpConfiguration;

        /// <summary>
        /// A list of static IP addresses assigned to this SIM. Each address is assigned at a defined network scope, made up of {attached data network, slice}.
        /// </summary>
        public InputList<Inputs.SimStaticIpPropertiesArgs> StaticIpConfiguration
        {
            get => _staticIpConfiguration ?? (_staticIpConfiguration = new InputList<Inputs.SimStaticIpPropertiesArgs>());
            set => _staticIpConfiguration = value;
        }

        public SimArgs()
        {
        }
        public static new SimArgs Empty => new SimArgs();
    }
}
