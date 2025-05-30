# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

__all__ = [
    'GetNetworkToNetworkInterconnectResult',
    'AwaitableGetNetworkToNetworkInterconnectResult',
    'get_network_to_network_interconnect',
    'get_network_to_network_interconnect_output',
]

@pulumi.output_type
class GetNetworkToNetworkInterconnectResult:
    """
    The Network To Network Interconnect resource definition.
    """
    def __init__(__self__, administrative_state=None, azure_api_version=None, configuration_state=None, egress_acl_id=None, export_route_policy=None, id=None, import_route_policy=None, ingress_acl_id=None, is_management_type=None, layer2_configuration=None, name=None, nni_type=None, npb_static_route_configuration=None, option_b_layer3_configuration=None, provisioning_state=None, system_data=None, type=None, use_option_b=None):
        if administrative_state and not isinstance(administrative_state, str):
            raise TypeError("Expected argument 'administrative_state' to be a str")
        pulumi.set(__self__, "administrative_state", administrative_state)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if configuration_state and not isinstance(configuration_state, str):
            raise TypeError("Expected argument 'configuration_state' to be a str")
        pulumi.set(__self__, "configuration_state", configuration_state)
        if egress_acl_id and not isinstance(egress_acl_id, str):
            raise TypeError("Expected argument 'egress_acl_id' to be a str")
        pulumi.set(__self__, "egress_acl_id", egress_acl_id)
        if export_route_policy and not isinstance(export_route_policy, dict):
            raise TypeError("Expected argument 'export_route_policy' to be a dict")
        pulumi.set(__self__, "export_route_policy", export_route_policy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if import_route_policy and not isinstance(import_route_policy, dict):
            raise TypeError("Expected argument 'import_route_policy' to be a dict")
        pulumi.set(__self__, "import_route_policy", import_route_policy)
        if ingress_acl_id and not isinstance(ingress_acl_id, str):
            raise TypeError("Expected argument 'ingress_acl_id' to be a str")
        pulumi.set(__self__, "ingress_acl_id", ingress_acl_id)
        if is_management_type and not isinstance(is_management_type, str):
            raise TypeError("Expected argument 'is_management_type' to be a str")
        pulumi.set(__self__, "is_management_type", is_management_type)
        if layer2_configuration and not isinstance(layer2_configuration, dict):
            raise TypeError("Expected argument 'layer2_configuration' to be a dict")
        pulumi.set(__self__, "layer2_configuration", layer2_configuration)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nni_type and not isinstance(nni_type, str):
            raise TypeError("Expected argument 'nni_type' to be a str")
        pulumi.set(__self__, "nni_type", nni_type)
        if npb_static_route_configuration and not isinstance(npb_static_route_configuration, dict):
            raise TypeError("Expected argument 'npb_static_route_configuration' to be a dict")
        pulumi.set(__self__, "npb_static_route_configuration", npb_static_route_configuration)
        if option_b_layer3_configuration and not isinstance(option_b_layer3_configuration, dict):
            raise TypeError("Expected argument 'option_b_layer3_configuration' to be a dict")
        pulumi.set(__self__, "option_b_layer3_configuration", option_b_layer3_configuration)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if use_option_b and not isinstance(use_option_b, str):
            raise TypeError("Expected argument 'use_option_b' to be a str")
        pulumi.set(__self__, "use_option_b", use_option_b)

    @property
    @pulumi.getter(name="administrativeState")
    def administrative_state(self) -> builtins.str:
        """
        Administrative state of the resource.
        """
        return pulumi.get(self, "administrative_state")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter(name="configurationState")
    def configuration_state(self) -> builtins.str:
        """
        Configuration state of the resource.
        """
        return pulumi.get(self, "configuration_state")

    @property
    @pulumi.getter(name="egressAclId")
    def egress_acl_id(self) -> Optional[builtins.str]:
        """
        Egress Acl. ARM resource ID of Access Control Lists.
        """
        return pulumi.get(self, "egress_acl_id")

    @property
    @pulumi.getter(name="exportRoutePolicy")
    def export_route_policy(self) -> Optional['outputs.ExportRoutePolicyInformationResponse']:
        """
        Export Route Policy configuration.
        """
        return pulumi.get(self, "export_route_policy")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="importRoutePolicy")
    def import_route_policy(self) -> Optional['outputs.ImportRoutePolicyInformationResponse']:
        """
        Import Route Policy configuration.
        """
        return pulumi.get(self, "import_route_policy")

    @property
    @pulumi.getter(name="ingressAclId")
    def ingress_acl_id(self) -> Optional[builtins.str]:
        """
        Ingress Acl. ARM resource ID of Access Control Lists.
        """
        return pulumi.get(self, "ingress_acl_id")

    @property
    @pulumi.getter(name="isManagementType")
    def is_management_type(self) -> Optional[builtins.str]:
        """
        Configuration to use NNI for Infrastructure Management. Example: True/False.
        """
        return pulumi.get(self, "is_management_type")

    @property
    @pulumi.getter(name="layer2Configuration")
    def layer2_configuration(self) -> Optional['outputs.Layer2ConfigurationResponse']:
        """
        Common properties for Layer2 Configuration.
        """
        return pulumi.get(self, "layer2_configuration")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nniType")
    def nni_type(self) -> Optional[builtins.str]:
        """
        Type of NNI used. Example: CE | NPB
        """
        return pulumi.get(self, "nni_type")

    @property
    @pulumi.getter(name="npbStaticRouteConfiguration")
    def npb_static_route_configuration(self) -> Optional['outputs.NpbStaticRouteConfigurationResponse']:
        """
        NPB Static Route Configuration properties.
        """
        return pulumi.get(self, "npb_static_route_configuration")

    @property
    @pulumi.getter(name="optionBLayer3Configuration")
    def option_b_layer3_configuration(self) -> Optional['outputs.NetworkToNetworkInterconnectPropertiesResponseOptionBLayer3Configuration']:
        """
        Common properties for Layer3Configuration.
        """
        return pulumi.get(self, "option_b_layer3_configuration")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> builtins.str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useOptionB")
    def use_option_b(self) -> builtins.str:
        """
        Based on this option layer3 parameters are mandatory. Example: True/False
        """
        return pulumi.get(self, "use_option_b")


class AwaitableGetNetworkToNetworkInterconnectResult(GetNetworkToNetworkInterconnectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkToNetworkInterconnectResult(
            administrative_state=self.administrative_state,
            azure_api_version=self.azure_api_version,
            configuration_state=self.configuration_state,
            egress_acl_id=self.egress_acl_id,
            export_route_policy=self.export_route_policy,
            id=self.id,
            import_route_policy=self.import_route_policy,
            ingress_acl_id=self.ingress_acl_id,
            is_management_type=self.is_management_type,
            layer2_configuration=self.layer2_configuration,
            name=self.name,
            nni_type=self.nni_type,
            npb_static_route_configuration=self.npb_static_route_configuration,
            option_b_layer3_configuration=self.option_b_layer3_configuration,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            type=self.type,
            use_option_b=self.use_option_b)


def get_network_to_network_interconnect(network_fabric_name: Optional[builtins.str] = None,
                                        network_to_network_interconnect_name: Optional[builtins.str] = None,
                                        resource_group_name: Optional[builtins.str] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkToNetworkInterconnectResult:
    """
    Implements NetworkToNetworkInterconnects GET method.

    Uses Azure REST API version 2023-06-15.

    Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str network_fabric_name: Name of the Network Fabric.
    :param builtins.str network_to_network_interconnect_name: Name of the Network to Network Interconnect.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkFabricName'] = network_fabric_name
    __args__['networkToNetworkInterconnectName'] = network_to_network_interconnect_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:managednetworkfabric:getNetworkToNetworkInterconnect', __args__, opts=opts, typ=GetNetworkToNetworkInterconnectResult).value

    return AwaitableGetNetworkToNetworkInterconnectResult(
        administrative_state=pulumi.get(__ret__, 'administrative_state'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        configuration_state=pulumi.get(__ret__, 'configuration_state'),
        egress_acl_id=pulumi.get(__ret__, 'egress_acl_id'),
        export_route_policy=pulumi.get(__ret__, 'export_route_policy'),
        id=pulumi.get(__ret__, 'id'),
        import_route_policy=pulumi.get(__ret__, 'import_route_policy'),
        ingress_acl_id=pulumi.get(__ret__, 'ingress_acl_id'),
        is_management_type=pulumi.get(__ret__, 'is_management_type'),
        layer2_configuration=pulumi.get(__ret__, 'layer2_configuration'),
        name=pulumi.get(__ret__, 'name'),
        nni_type=pulumi.get(__ret__, 'nni_type'),
        npb_static_route_configuration=pulumi.get(__ret__, 'npb_static_route_configuration'),
        option_b_layer3_configuration=pulumi.get(__ret__, 'option_b_layer3_configuration'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'),
        use_option_b=pulumi.get(__ret__, 'use_option_b'))
def get_network_to_network_interconnect_output(network_fabric_name: Optional[pulumi.Input[builtins.str]] = None,
                                               network_to_network_interconnect_name: Optional[pulumi.Input[builtins.str]] = None,
                                               resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNetworkToNetworkInterconnectResult]:
    """
    Implements NetworkToNetworkInterconnects GET method.

    Uses Azure REST API version 2023-06-15.

    Other available API versions: 2023-02-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native managednetworkfabric [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str network_fabric_name: Name of the Network Fabric.
    :param builtins.str network_to_network_interconnect_name: Name of the Network to Network Interconnect.
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['networkFabricName'] = network_fabric_name
    __args__['networkToNetworkInterconnectName'] = network_to_network_interconnect_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:managednetworkfabric:getNetworkToNetworkInterconnect', __args__, opts=opts, typ=GetNetworkToNetworkInterconnectResult)
    return __ret__.apply(lambda __response__: GetNetworkToNetworkInterconnectResult(
        administrative_state=pulumi.get(__response__, 'administrative_state'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        configuration_state=pulumi.get(__response__, 'configuration_state'),
        egress_acl_id=pulumi.get(__response__, 'egress_acl_id'),
        export_route_policy=pulumi.get(__response__, 'export_route_policy'),
        id=pulumi.get(__response__, 'id'),
        import_route_policy=pulumi.get(__response__, 'import_route_policy'),
        ingress_acl_id=pulumi.get(__response__, 'ingress_acl_id'),
        is_management_type=pulumi.get(__response__, 'is_management_type'),
        layer2_configuration=pulumi.get(__response__, 'layer2_configuration'),
        name=pulumi.get(__response__, 'name'),
        nni_type=pulumi.get(__response__, 'nni_type'),
        npb_static_route_configuration=pulumi.get(__response__, 'npb_static_route_configuration'),
        option_b_layer3_configuration=pulumi.get(__response__, 'option_b_layer3_configuration'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type'),
        use_option_b=pulumi.get(__response__, 'use_option_b')))
