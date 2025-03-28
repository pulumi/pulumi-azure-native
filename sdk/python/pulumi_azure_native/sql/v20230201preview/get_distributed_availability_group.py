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
from ... import _utilities
from . import outputs

__all__ = [
    'GetDistributedAvailabilityGroupResult',
    'AwaitableGetDistributedAvailabilityGroupResult',
    'get_distributed_availability_group',
    'get_distributed_availability_group_output',
]

@pulumi.output_type
class GetDistributedAvailabilityGroupResult:
    """
    Distributed availability group between box and Sql Managed Instance.
    """
    def __init__(__self__, databases=None, distributed_availability_group_id=None, distributed_availability_group_name=None, id=None, instance_availability_group_name=None, instance_link_role=None, name=None, partner_availability_group_name=None, partner_endpoint=None, partner_link_role=None, replication_mode=None, type=None):
        if databases and not isinstance(databases, list):
            raise TypeError("Expected argument 'databases' to be a list")
        pulumi.set(__self__, "databases", databases)
        if distributed_availability_group_id and not isinstance(distributed_availability_group_id, str):
            raise TypeError("Expected argument 'distributed_availability_group_id' to be a str")
        pulumi.set(__self__, "distributed_availability_group_id", distributed_availability_group_id)
        if distributed_availability_group_name and not isinstance(distributed_availability_group_name, str):
            raise TypeError("Expected argument 'distributed_availability_group_name' to be a str")
        pulumi.set(__self__, "distributed_availability_group_name", distributed_availability_group_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_availability_group_name and not isinstance(instance_availability_group_name, str):
            raise TypeError("Expected argument 'instance_availability_group_name' to be a str")
        pulumi.set(__self__, "instance_availability_group_name", instance_availability_group_name)
        if instance_link_role and not isinstance(instance_link_role, str):
            raise TypeError("Expected argument 'instance_link_role' to be a str")
        pulumi.set(__self__, "instance_link_role", instance_link_role)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if partner_availability_group_name and not isinstance(partner_availability_group_name, str):
            raise TypeError("Expected argument 'partner_availability_group_name' to be a str")
        pulumi.set(__self__, "partner_availability_group_name", partner_availability_group_name)
        if partner_endpoint and not isinstance(partner_endpoint, str):
            raise TypeError("Expected argument 'partner_endpoint' to be a str")
        pulumi.set(__self__, "partner_endpoint", partner_endpoint)
        if partner_link_role and not isinstance(partner_link_role, str):
            raise TypeError("Expected argument 'partner_link_role' to be a str")
        pulumi.set(__self__, "partner_link_role", partner_link_role)
        if replication_mode and not isinstance(replication_mode, str):
            raise TypeError("Expected argument 'replication_mode' to be a str")
        pulumi.set(__self__, "replication_mode", replication_mode)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def databases(self) -> Sequence['outputs.DistributedAvailabilityGroupDatabaseResponse']:
        """
        Databases in the distributed availability group
        """
        return pulumi.get(self, "databases")

    @property
    @pulumi.getter(name="distributedAvailabilityGroupId")
    def distributed_availability_group_id(self) -> str:
        """
        ID of the distributed availability group
        """
        return pulumi.get(self, "distributed_availability_group_id")

    @property
    @pulumi.getter(name="distributedAvailabilityGroupName")
    def distributed_availability_group_name(self) -> str:
        """
        Name of the distributed availability group
        """
        return pulumi.get(self, "distributed_availability_group_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceAvailabilityGroupName")
    def instance_availability_group_name(self) -> str:
        """
        Managed instance side availability group name
        """
        return pulumi.get(self, "instance_availability_group_name")

    @property
    @pulumi.getter(name="instanceLinkRole")
    def instance_link_role(self) -> str:
        """
        Managed instance side link role
        """
        return pulumi.get(self, "instance_link_role")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partnerAvailabilityGroupName")
    def partner_availability_group_name(self) -> str:
        """
        SQL server side availability group name
        """
        return pulumi.get(self, "partner_availability_group_name")

    @property
    @pulumi.getter(name="partnerEndpoint")
    def partner_endpoint(self) -> str:
        """
        SQL server side endpoint - IP or DNS resolvable name
        """
        return pulumi.get(self, "partner_endpoint")

    @property
    @pulumi.getter(name="partnerLinkRole")
    def partner_link_role(self) -> str:
        """
        SQL server side link role
        """
        return pulumi.get(self, "partner_link_role")

    @property
    @pulumi.getter(name="replicationMode")
    def replication_mode(self) -> Optional[str]:
        """
        Replication mode of the link
        """
        return pulumi.get(self, "replication_mode")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetDistributedAvailabilityGroupResult(GetDistributedAvailabilityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDistributedAvailabilityGroupResult(
            databases=self.databases,
            distributed_availability_group_id=self.distributed_availability_group_id,
            distributed_availability_group_name=self.distributed_availability_group_name,
            id=self.id,
            instance_availability_group_name=self.instance_availability_group_name,
            instance_link_role=self.instance_link_role,
            name=self.name,
            partner_availability_group_name=self.partner_availability_group_name,
            partner_endpoint=self.partner_endpoint,
            partner_link_role=self.partner_link_role,
            replication_mode=self.replication_mode,
            type=self.type)


def get_distributed_availability_group(distributed_availability_group_name: Optional[str] = None,
                                       managed_instance_name: Optional[str] = None,
                                       resource_group_name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDistributedAvailabilityGroupResult:
    """
    Gets a distributed availability group info.


    :param str distributed_availability_group_name: The distributed availability group name.
    :param str managed_instance_name: The name of the managed instance.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    """
    __args__ = dict()
    __args__['distributedAvailabilityGroupName'] = distributed_availability_group_name
    __args__['managedInstanceName'] = managed_instance_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:sql/v20230201preview:getDistributedAvailabilityGroup', __args__, opts=opts, typ=GetDistributedAvailabilityGroupResult).value

    return AwaitableGetDistributedAvailabilityGroupResult(
        databases=pulumi.get(__ret__, 'databases'),
        distributed_availability_group_id=pulumi.get(__ret__, 'distributed_availability_group_id'),
        distributed_availability_group_name=pulumi.get(__ret__, 'distributed_availability_group_name'),
        id=pulumi.get(__ret__, 'id'),
        instance_availability_group_name=pulumi.get(__ret__, 'instance_availability_group_name'),
        instance_link_role=pulumi.get(__ret__, 'instance_link_role'),
        name=pulumi.get(__ret__, 'name'),
        partner_availability_group_name=pulumi.get(__ret__, 'partner_availability_group_name'),
        partner_endpoint=pulumi.get(__ret__, 'partner_endpoint'),
        partner_link_role=pulumi.get(__ret__, 'partner_link_role'),
        replication_mode=pulumi.get(__ret__, 'replication_mode'),
        type=pulumi.get(__ret__, 'type'))
def get_distributed_availability_group_output(distributed_availability_group_name: Optional[pulumi.Input[str]] = None,
                                              managed_instance_name: Optional[pulumi.Input[str]] = None,
                                              resource_group_name: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDistributedAvailabilityGroupResult]:
    """
    Gets a distributed availability group info.


    :param str distributed_availability_group_name: The distributed availability group name.
    :param str managed_instance_name: The name of the managed instance.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    """
    __args__ = dict()
    __args__['distributedAvailabilityGroupName'] = distributed_availability_group_name
    __args__['managedInstanceName'] = managed_instance_name
    __args__['resourceGroupName'] = resource_group_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:sql/v20230201preview:getDistributedAvailabilityGroup', __args__, opts=opts, typ=GetDistributedAvailabilityGroupResult)
    return __ret__.apply(lambda __response__: GetDistributedAvailabilityGroupResult(
        databases=pulumi.get(__response__, 'databases'),
        distributed_availability_group_id=pulumi.get(__response__, 'distributed_availability_group_id'),
        distributed_availability_group_name=pulumi.get(__response__, 'distributed_availability_group_name'),
        id=pulumi.get(__response__, 'id'),
        instance_availability_group_name=pulumi.get(__response__, 'instance_availability_group_name'),
        instance_link_role=pulumi.get(__response__, 'instance_link_role'),
        name=pulumi.get(__response__, 'name'),
        partner_availability_group_name=pulumi.get(__response__, 'partner_availability_group_name'),
        partner_endpoint=pulumi.get(__response__, 'partner_endpoint'),
        partner_link_role=pulumi.get(__response__, 'partner_link_role'),
        replication_mode=pulumi.get(__response__, 'replication_mode'),
        type=pulumi.get(__response__, 'type')))
