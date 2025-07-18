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
    'GetFqdnListGlobalRulestackResult',
    'AwaitableGetFqdnListGlobalRulestackResult',
    'get_fqdn_list_global_rulestack',
    'get_fqdn_list_global_rulestack_output',
]

@pulumi.output_type
class GetFqdnListGlobalRulestackResult:
    """
    GlobalRulestack fqdnList
    """
    def __init__(__self__, audit_comment=None, azure_api_version=None, description=None, etag=None, fqdn_list=None, id=None, name=None, provisioning_state=None, system_data=None, type=None):
        if audit_comment and not isinstance(audit_comment, str):
            raise TypeError("Expected argument 'audit_comment' to be a str")
        pulumi.set(__self__, "audit_comment", audit_comment)
        if azure_api_version and not isinstance(azure_api_version, str):
            raise TypeError("Expected argument 'azure_api_version' to be a str")
        pulumi.set(__self__, "azure_api_version", azure_api_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if fqdn_list and not isinstance(fqdn_list, list):
            raise TypeError("Expected argument 'fqdn_list' to be a list")
        pulumi.set(__self__, "fqdn_list", fqdn_list)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if system_data and not isinstance(system_data, dict):
            raise TypeError("Expected argument 'system_data' to be a dict")
        pulumi.set(__self__, "system_data", system_data)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="auditComment")
    def audit_comment(self) -> Optional[builtins.str]:
        """
        comment for this object
        """
        return pulumi.get(self, "audit_comment")

    @property
    @pulumi.getter(name="azureApiVersion")
    def azure_api_version(self) -> builtins.str:
        """
        The Azure API version of the resource.
        """
        return pulumi.get(self, "azure_api_version")

    @property
    @pulumi.getter
    def description(self) -> Optional[builtins.str]:
        """
        fqdn object description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> Optional[builtins.str]:
        """
        etag info
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="fqdnList")
    def fqdn_list(self) -> Sequence[builtins.str]:
        """
        fqdn list
        """
        return pulumi.get(self, "fqdn_list")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

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


class AwaitableGetFqdnListGlobalRulestackResult(GetFqdnListGlobalRulestackResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFqdnListGlobalRulestackResult(
            audit_comment=self.audit_comment,
            azure_api_version=self.azure_api_version,
            description=self.description,
            etag=self.etag,
            fqdn_list=self.fqdn_list,
            id=self.id,
            name=self.name,
            provisioning_state=self.provisioning_state,
            system_data=self.system_data,
            type=self.type)


def get_fqdn_list_global_rulestack(global_rulestack_name: Optional[builtins.str] = None,
                                   name: Optional[builtins.str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFqdnListGlobalRulestackResult:
    """
    Get a FqdnListGlobalRulestackResource

    Uses Azure REST API version 2025-02-06-preview.

    Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str global_rulestack_name: GlobalRulestack resource name
    :param builtins.str name: fqdn list name
    """
    __args__ = dict()
    __args__['globalRulestackName'] = global_rulestack_name
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('azure-native:cloudngfw:getFqdnListGlobalRulestack', __args__, opts=opts, typ=GetFqdnListGlobalRulestackResult).value

    return AwaitableGetFqdnListGlobalRulestackResult(
        audit_comment=pulumi.get(__ret__, 'audit_comment'),
        azure_api_version=pulumi.get(__ret__, 'azure_api_version'),
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        fqdn_list=pulumi.get(__ret__, 'fqdn_list'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        provisioning_state=pulumi.get(__ret__, 'provisioning_state'),
        system_data=pulumi.get(__ret__, 'system_data'),
        type=pulumi.get(__ret__, 'type'))
def get_fqdn_list_global_rulestack_output(global_rulestack_name: Optional[pulumi.Input[builtins.str]] = None,
                                          name: Optional[pulumi.Input[builtins.str]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetFqdnListGlobalRulestackResult]:
    """
    Get a FqdnListGlobalRulestackResource

    Uses Azure REST API version 2025-02-06-preview.

    Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview, 2025-05-23. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.


    :param builtins.str global_rulestack_name: GlobalRulestack resource name
    :param builtins.str name: fqdn list name
    """
    __args__ = dict()
    __args__['globalRulestackName'] = global_rulestack_name
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('azure-native:cloudngfw:getFqdnListGlobalRulestack', __args__, opts=opts, typ=GetFqdnListGlobalRulestackResult)
    return __ret__.apply(lambda __response__: GetFqdnListGlobalRulestackResult(
        audit_comment=pulumi.get(__response__, 'audit_comment'),
        azure_api_version=pulumi.get(__response__, 'azure_api_version'),
        description=pulumi.get(__response__, 'description'),
        etag=pulumi.get(__response__, 'etag'),
        fqdn_list=pulumi.get(__response__, 'fqdn_list'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        provisioning_state=pulumi.get(__response__, 'provisioning_state'),
        system_data=pulumi.get(__response__, 'system_data'),
        type=pulumi.get(__response__, 'type')))
