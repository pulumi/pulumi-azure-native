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
    'GetHuntResult',
    'AwaitableGetHuntResult',
    'get_hunt',
    'get_hunt_output',
]

@pulumi.output_type
class GetHuntResult:
    """
    Represents a Hunt in Azure Security Insights.
    """
    def __init__(__self__, attack_tactics=None, attack_techniques=None, azure_api_version=None, description=None, display_name=None, etag=None, hypothesis_status=None, id=None, labels=None, name=None, owner=None, status=None, system_data=None, type=None):
        if attack_tactics and not isinstance(attack_tactics, list):
            raise TypeError("Expected argument 'attack_tactics' to be a list")
        pulumi.set(__self__, "attack_tactics", attack_tactics)
        if attack_techniques and not isinstance(attack_techniques, list):
            raise TypeError("Expected argument 'attack_techniques' to be a list")
        pulumi.set(__self__, "attack_techniques", attack_techniques)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if hypothesis_status and not isinstance(hypothesis_status, str):
            raise TypeError("Expected argument 'hypothesis_status' to be a str")
        pulumi.set(__self__, "hypothesis_status", hypothesis_status)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, dict):
            raise TypeError("Expected argument 'owner' to be a dict")
        pulumi.set(__self__, "owner", owner)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="attackTactics")
    def attack_tactics(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of mitre attack tactics the hunt is associated with
        """
        return pulumi.get(self, "attack_tactics")

    @property
    @pulumi.getter(name="attackTechniques")
    def attack_techniques(self) -> Optional[Sequence[builtins.str]]:
        """
        A list of a mitre attack techniques the hunt is associated with
        """
        return pulumi.get(self, "attack_techniques")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        The description of the hunt
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> builtins.str:
        """
        The display name of the hunt
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> Optional[builtins.str]:
        """
        Etag of the azure resource
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="hypothesisStatus")
    def hypothesis_status(self) -> Optional[builtins.str]:
        """
        The hypothesis status of the hunt.
        """
        return pulumi.get(self, "hypothesis_status")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def labels(self) -> Optional[Sequence[builtins.str]]:
        """
        List of labels relevant to this hunt 
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> Optional['outputs.HuntOwnerResponse']:
        """
        Describes a user that the hunt is assigned to
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def status(self) -> Optional[builtins.str]:
        """
        The status of the hunt.
        """
        return pulumi.get(self, "status")

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


class AwaitableGetHuntResult(GetHuntResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHuntResult(
            attack_tactics=self.attack_tactics,
            attack_techniques=self.attack_techniques,
            azure_api_version=self.azure_api_version,
            description=self.description,
            display_name=self.display_name,
            etag=self.etag,
            hypothesis_status=self.hypothesis_status,
            id=self.id,
            labels=self.labels,
            name=self.name,
            owner=self.owner,
            status=self.status,
            system_data=self.system_data,
            type=self.type)


def get_hunt(hunt_id: Optional[builtins.str] = None,
             resource_group_name: Optional[builtins.str] = None,
             workspace_name: Optional[builtins.str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHuntResult:
    """
    Gets a hunt, without relations and comments.

    Uses Azure REST API version 2025-01-01-preview.

    Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str hunt_id: The hunt id (GUID)
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['huntId'] = hunt_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:securityinsights:getHunt', __args__, opts=opts, typ=GetHuntResult).value

    return AwaitableGetHuntResult(
        attack_tactics=pulumi.get(__ret__, 'attack_tactics'),
        attack_techniques=pulumi.get(__ret__, 'attack_techniques'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        etag=pulumi.get(__ret__, 'etag'),
        hypothesis_status=pulumi.get(__ret__, 'hypothesis_status'),
        id=pulumi.get(__ret__, 'id'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        owner=pulumi.get(__ret__, 'owner'),
        status=pulumi.get(__ret__, 'status'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_hunt_output(hunt_id: Optional[pulumi.Input[builtins.str]] = None,
                    resource_group_name: Optional[pulumi.Input[builtins.str]] = None,
                    workspace_name: Optional[pulumi.Input[builtins.str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetHuntResult]:
    """
    Gets a hunt, without relations and comments.

    Uses Azure REST API version 2025-01-01-preview.

    Other available API versions: 2023-04-01-preview, 2023-05-01-preview, 2023-06-01-preview, 2023-07-01-preview, 2023-08-01-preview, 2023-09-01-preview, 2023-10-01-preview, 2023-12-01-preview, 2024-01-01-preview, 2024-04-01-preview, 2024-10-01-preview, 2025-04-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native securityinsights [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str hunt_id: The hunt id (GUID)
    :param builtins.str resource_group_name: The name of the resource group. The name is case insensitive.
    :param builtins.str workspace_name: The name of the workspace.
    """
    __args__ = dict()
    __args__['huntId'] = hunt_id
    __args__['resourceGroupName'] = resource_group_name
    __args__['workspaceName'] = workspace_name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:securityinsights:getHunt', __args__, opts=opts, typ=GetHuntResult)
    return __ret__.apply(lambda __response__: GetHuntResult(
        attack_tactics=pulumi.get(__response__, 'attack_tactics'),
        attack_techniques=pulumi.get(__response__, 'attack_techniques'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        description=pulumi.get(__response__, 'description'),
        display_name=pulumi.get(__response__, 'display_name'),
        etag=pulumi.get(__response__, 'etag'),
        hypothesis_status=pulumi.get(__response__, 'hypothesis_status'),
        id=pulumi.get(__response__, 'id'),
        labels=pulumi.get(__response__, 'labels'),
        name=pulumi.get(__response__, 'name'),
        owner=pulumi.get(__response__, 'owner'),
        status=pulumi.get(__response__, 'status'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
