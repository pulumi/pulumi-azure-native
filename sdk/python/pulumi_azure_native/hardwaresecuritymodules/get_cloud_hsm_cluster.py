# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
    'GetCloudHsmClusterResult',
    'AwaitableGetCloudHsmClusterResult',
    'get_cloud_hsm_cluster',
    'get_cloud_hsm_cluster_output',
]

@pulumi.output_type
class GetCloudHsmClusterResult:
    """
    Resource information with extended details.
    """
    def __init__(__self__, auto_generated_domain_name_label_scope=None, hsms=None, id=None, location=None, name=None, private_endpoint_connections=None, provisioning_state=None, public_network_access=None, security_domain=None, sku=None, system_data=None, tags=None, type=None):
        if auto_generated_domain_name_label_scope and not isinstance(auto_generated_domain_name_label_scope, str):
            raise TypeError("Expected argument 'auto_generated_domain_name_label_scope' to be a str")
        pulumi.set(__self__, "auto_generated_domain_name_label_scope", auto_generated_domain_name_label_scope)
        if hsms and not isinstance(hsms, list):
            raise TypeError("Expected argument 'hsms' to be a list")
        pulumi.set(__self__, "hsms", hsms)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint_connections and not isinstance(private_endpoint_connections, list):
            raise TypeError("Expected argument 'private_endpoint_connections' to be a list")
        pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if public_network_access and not isinstance(public_network_access, str):
            raise TypeError("Expected argument 'public_network_access' to be a str")
        pulumi.set(__self__, "public_network_access", public_network_access)
        if security_domain and not isinstance(security_domain, dict):
            raise TypeError("Expected argument 'security_domain' to be a dict")
        pulumi.set(__self__, "security_domain", security_domain)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="autoGeneratedDomainNameLabelScope")
    def auto_generated_domain_name_label_scope(self) -> Optional[str]:
        """
        The Cloud HSM Cluster's auto-generated Domain Name Label Scope
        """
        return pulumi.get(self, "auto_generated_domain_name_label_scope")

    @property
    @pulumi.getter
    def hsms(self) -> Optional[Sequence['outputs.CloudHsmPropertiesResponse']]:
        """
        An array of Cloud HSM Cluster's HSMs
        """
        return pulumi.get(self, "hsms")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Optional[Sequence['outputs.PrivateEndpointConnectionResponse']]:
        """
        List of private endpoint connection resources
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        The Cloud HSM Cluster's provisioningState
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicNetworkAccess")
    def public_network_access(self) -> Optional[str]:
        """
        The Cloud HSM Cluster public network access
        """
        return pulumi.get(self, "public_network_access")

    @property
    @pulumi.getter(name="securityDomain")
    def security_domain(self) -> Optional['outputs.CloudHsmClusterSecurityDomainPropertiesResponse']:
        """
        Security domain properties information for Cloud HSM cluster
        """
        return pulumi.get(self, "security_domain")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.CloudHsmClusterSkuResponse']:
        """
        SKU details
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> 'outputs.SystemDataResponse':
        """
        Azure Resource Manager metadata containing createdBy and modifiedBy information.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetCloudHsmClusterResult(GetCloudHsmClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCloudHsmClusterResult(
            auto_generated_domain_name_label_scope=self.auto_generated_domain_name_label_scope,
            hsms=self.hsms,
            id=self.id,
            location=self.location,
            name=self.name,
            private_endpoint_connections=self.private_endpoint_connections,
            provisioning_state=self.provisioning_state,
            public_network_access=self.public_network_access,
            security_domain=self.security_domain,
            sku=self.sku,
            system_data=self.system_data,
            tags=self.tags,
            type=self.type)


def get_cloud_hsm_cluster(cloud_hsm_cluster_name: Optional[str] = None,
                          resource_group_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCloudHsmClusterResult:
    """
    Gets the specified Cloud HSM Cluster

    Uses Azure REST API version 2022-08-31-preview.

    Other available API versions: 2023-12-10-preview, 2024-06-30-preview.


    :param str cloud_hsm_cluster_name: The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must be between 3 and 24 characters in length.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['cloudHsmClusterName'] = cloud_hsm_cluster_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:hardwaresecuritymodules:getCloudHsmCluster', __args__, opts=opts, typ=GetCloudHsmClusterResult).value

    return AwaitableGetCloudHsmClusterResult(
        auto_generated_domain_name_label_scope=pulumi.get(__ret__, 'auto_generated_domain_name_label_scope'),
        hsms=pulumi.get(__ret__, 'hsms'),
        id=pulumi.get(__ret__, 'id'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        private_endpoint_connections=pulumi.get(__ret__, 'private_endpoint_connections'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        public_network_access=pulumi.get(__ret__, 'public_network_access'),
        security_domain=pulumi.get(__ret__, 'security_domain'),
        sku=pulumi.get(__ret__, 'sku'),
        system_data=pulumi.get(__ret__, 'system_data'),
        tags=pulumi.get(__ret__, 'tags'),
        type=pulumi.get(__ret__, 'type'))
def get_cloud_hsm_cluster_output(cloud_hsm_cluster_name: Optional[pulumi.Input[str]] = None,
                                 resource_group_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetCloudHsmClusterResult]:
    """
    Gets the specified Cloud HSM Cluster

    Uses Azure REST API version 2022-08-31-preview.

    Other available API versions: 2023-12-10-preview, 2024-06-30-preview.


    :param str cloud_hsm_cluster_name: The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must be between 3 and 24 characters in length.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['cloudHsmClusterName'] = cloud_hsm_cluster_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:hardwaresecuritymodules:getCloudHsmCluster', __args__, opts=opts, typ=GetCloudHsmClusterResult)
    return __ret__.apply(lambda __response__: GetCloudHsmClusterResult(
        auto_generated_domain_name_label_scope=pulumi.get(__response__, 'auto_generated_domain_name_label_scope'),
        hsms=pulumi.get(__response__, 'hsms'),
        id=pulumi.get(__response__, 'id'),
        location=pulumi.get(__response__, 'location'),
        name=pulumi.get(__response__, 'name'),
        private_endpoint_connections=pulumi.get(__response__, 'private_endpoint_connections'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        public_network_access=pulumi.get(__response__, 'public_network_access'),
        security_domain=pulumi.get(__response__, 'security_domain'),
        sku=pulumi.get(__response__, 'sku'),
        system_data=pulumi.get(__response__, 'system_data'),
        tags=pulumi.get(__response__, 'tags'),
        type=pulumi.get(__response__, 'type')))
