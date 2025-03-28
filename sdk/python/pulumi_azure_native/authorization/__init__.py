# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing
# Export this package's modules as members:
from ._enums import *
from .access_review_history_definition_by_id import *
from .access_review_schedule_definition_by_id import *
from .get_access_review_history_definition_by_id import *
from .get_access_review_schedule_definition_by_id import *
from .get_client_config import *
from .get_client_token import *
from .get_management_lock_at_resource_group_level import *
from .get_management_lock_at_resource_level import *
from .get_management_lock_at_subscription_level import *
from .get_management_lock_by_scope import *
from .get_pim_role_eligibility_schedule import *
from .get_policy_assignment import *
from .get_policy_definition import *
from .get_policy_definition_at_management_group import *
from .get_policy_definition_version import *
from .get_policy_definition_version_at_management_group import *
from .get_policy_exemption import *
from .get_policy_set_definition import *
from .get_policy_set_definition_at_management_group import *
from .get_policy_set_definition_version import *
from .get_policy_set_definition_version_at_management_group import *
from .get_private_link_association import *
from .get_resource_management_private_link import *
from .get_role_assignment import *
from .get_role_definition import *
from .get_role_management_policy import *
from .get_role_management_policy_assignment import *
from .get_scope_access_review_history_definition_by_id import *
from .get_scope_access_review_schedule_definition_by_id import *
from .get_variable import *
from .get_variable_at_management_group import *
from .get_variable_value import *
from .get_variable_value_at_management_group import *
from .list_policy_definition_version_all import *
from .list_policy_definition_version_all_at_management_group import *
from .list_policy_definition_version_all_builtins import *
from .list_policy_set_definition_version_all import *
from .list_policy_set_definition_version_all_at_management_group import *
from .list_policy_set_definition_version_all_builtins import *
from .management_lock_at_resource_group_level import *
from .management_lock_at_resource_level import *
from .management_lock_at_subscription_level import *
from .management_lock_by_scope import *
from .pim_role_eligibility_schedule import *
from .policy_assignment import *
from .policy_definition import *
from .policy_definition_at_management_group import *
from .policy_definition_version import *
from .policy_definition_version_at_management_group import *
from .policy_exemption import *
from .policy_set_definition import *
from .policy_set_definition_at_management_group import *
from .policy_set_definition_version import *
from .policy_set_definition_version_at_management_group import *
from .private_link_association import *
from .resource_management_private_link import *
from .role_assignment import *
from .role_definition import *
from .role_management_policy import *
from .role_management_policy_assignment import *
from .scope_access_review_history_definition_by_id import *
from .scope_access_review_schedule_definition_by_id import *
from .variable import *
from .variable_at_management_group import *
from .variable_value import *
from .variable_value_at_management_group import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_azure_native.authorization.v20171001preview as __v20171001preview
    v20171001preview = __v20171001preview
    import pulumi_azure_native.authorization.v20180501 as __v20180501
    v20180501 = __v20180501
    import pulumi_azure_native.authorization.v20190601 as __v20190601
    v20190601 = __v20190601
    import pulumi_azure_native.authorization.v20200301 as __v20200301
    v20200301 = __v20200301
    import pulumi_azure_native.authorization.v20200301preview as __v20200301preview
    v20200301preview = __v20200301preview
    import pulumi_azure_native.authorization.v20200401preview as __v20200401preview
    v20200401preview = __v20200401preview
    import pulumi_azure_native.authorization.v20200501 as __v20200501
    v20200501 = __v20200501
    import pulumi_azure_native.authorization.v20201001 as __v20201001
    v20201001 = __v20201001
    import pulumi_azure_native.authorization.v20201001preview as __v20201001preview
    v20201001preview = __v20201001preview
    import pulumi_azure_native.authorization.v20210601 as __v20210601
    v20210601 = __v20210601
    import pulumi_azure_native.authorization.v20211201preview as __v20211201preview
    v20211201preview = __v20211201preview
    import pulumi_azure_native.authorization.v20220401 as __v20220401
    v20220401 = __v20220401
    import pulumi_azure_native.authorization.v20220501preview as __v20220501preview
    v20220501preview = __v20220501preview
    import pulumi_azure_native.authorization.v20220601 as __v20220601
    v20220601 = __v20220601
    import pulumi_azure_native.authorization.v20220701preview as __v20220701preview
    v20220701preview = __v20220701preview
    import pulumi_azure_native.authorization.v20220801preview as __v20220801preview
    v20220801preview = __v20220801preview
    import pulumi_azure_native.authorization.v20230401 as __v20230401
    v20230401 = __v20230401
    import pulumi_azure_native.authorization.v20240201preview as __v20240201preview
    v20240201preview = __v20240201preview
    import pulumi_azure_native.authorization.v20240401 as __v20240401
    v20240401 = __v20240401
    import pulumi_azure_native.authorization.v20240501 as __v20240501
    v20240501 = __v20240501
    import pulumi_azure_native.authorization.v20240901preview as __v20240901preview
    v20240901preview = __v20240901preview
    import pulumi_azure_native.authorization.v20241201preview as __v20241201preview
    v20241201preview = __v20241201preview
    import pulumi_azure_native.authorization.v20250101 as __v20250101
    v20250101 = __v20250101
    import pulumi_azure_native.authorization.v20250301 as __v20250301
    v20250301 = __v20250301
else:
    v20171001preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20171001preview')
    v20180501 = _utilities.lazy_import('pulumi_azure_native.authorization.v20180501')
    v20190601 = _utilities.lazy_import('pulumi_azure_native.authorization.v20190601')
    v20200301 = _utilities.lazy_import('pulumi_azure_native.authorization.v20200301')
    v20200301preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20200301preview')
    v20200401preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20200401preview')
    v20200501 = _utilities.lazy_import('pulumi_azure_native.authorization.v20200501')
    v20201001 = _utilities.lazy_import('pulumi_azure_native.authorization.v20201001')
    v20201001preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20201001preview')
    v20210601 = _utilities.lazy_import('pulumi_azure_native.authorization.v20210601')
    v20211201preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20211201preview')
    v20220401 = _utilities.lazy_import('pulumi_azure_native.authorization.v20220401')
    v20220501preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20220501preview')
    v20220601 = _utilities.lazy_import('pulumi_azure_native.authorization.v20220601')
    v20220701preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20220701preview')
    v20220801preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20220801preview')
    v20230401 = _utilities.lazy_import('pulumi_azure_native.authorization.v20230401')
    v20240201preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20240201preview')
    v20240401 = _utilities.lazy_import('pulumi_azure_native.authorization.v20240401')
    v20240501 = _utilities.lazy_import('pulumi_azure_native.authorization.v20240501')
    v20240901preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20240901preview')
    v20241201preview = _utilities.lazy_import('pulumi_azure_native.authorization.v20241201preview')
    v20250101 = _utilities.lazy_import('pulumi_azure_native.authorization.v20250101')
    v20250301 = _utilities.lazy_import('pulumi_azure_native.authorization.v20250301')

