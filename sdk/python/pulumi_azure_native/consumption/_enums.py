# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import pulumi
from enum import Enum

__all__ = [
    'BudgetOperatorType',
    'CategoryType',
    'CultureCode',
    'OperatorType',
    'ThresholdType',
    'TimeGrainType',
]


@pulumi.type_token("azure-native:consumption:BudgetOperatorType")
class BudgetOperatorType(builtins.str, Enum):
    """
    The operator to use for comparison.
    """
    IN_ = "In"


@pulumi.type_token("azure-native:consumption:CategoryType")
class CategoryType(builtins.str, Enum):
    """
    The category of the budget, whether the budget tracks cost or usage.
    """
    COST = "Cost"


@pulumi.type_token("azure-native:consumption:CultureCode")
class CultureCode(builtins.str, Enum):
    """
    Language in which the recipient will receive the notification
    """
    EN_US = "en-us"
    JA_JP = "ja-jp"
    ZH_CN = "zh-cn"
    DE_DE = "de-de"
    ES_ES = "es-es"
    FR_FR = "fr-fr"
    IT_IT = "it-it"
    KO_KR = "ko-kr"
    PT_BR = "pt-br"
    RU_RU = "ru-ru"
    ZH_TW = "zh-tw"
    CS_CZ = "cs-cz"
    PL_PL = "pl-pl"
    TR_TR = "tr-tr"
    DA_DK = "da-dk"
    EN_GB = "en-gb"
    HU_HU = "hu-hu"
    NB_NO = "nb-no"
    NL_NL = "nl-nl"
    PT_PT = "pt-pt"
    SV_SE = "sv-se"


@pulumi.type_token("azure-native:consumption:OperatorType")
class OperatorType(builtins.str, Enum):
    """
    The comparison operator.
    """
    EQUAL_TO = "EqualTo"
    """
    Alert will be triggered if the evaluated cost is the same as threshold value. Note: It’s not recommended to use this OperatorType as there’s low chance of cost being exactly the same as threshold value, leading to missing of your alert. This OperatorType will be deprecated in future. 
    """
    GREATER_THAN = "GreaterThan"
    """
    Alert will be triggered if the evaluated cost is greater than the threshold value. Note: This is the recommended OperatorType while configuring Budget Alert.
    """
    GREATER_THAN_OR_EQUAL_TO = "GreaterThanOrEqualTo"
    """
    Alert will be triggered if the evaluated cost is greater than or equal to the threshold value.
    """


@pulumi.type_token("azure-native:consumption:ThresholdType")
class ThresholdType(builtins.str, Enum):
    """
    The type of threshold
    """
    ACTUAL = "Actual"
    """
    Actual costs budget alerts notify when the actual accrued cost exceeds the allocated budget .
    """
    FORECASTED = "Forecasted"
    """
    Forecasted costs budget alerts provide advanced notification that your spending trends are likely to exceed your allocated budget, as it relies on forecasted cost predictions.
    """


@pulumi.type_token("azure-native:consumption:TimeGrainType")
class TimeGrainType(builtins.str, Enum):
    """
    The time covered by a budget. Tracking of the amount will be reset based on the time grain. BillingMonth, BillingQuarter, and BillingAnnual are only supported by WD customers
    """
    MONTHLY = "Monthly"
    QUARTERLY = "Quarterly"
    ANNUALLY = "Annually"
    BILLING_MONTH = "BillingMonth"
    BILLING_QUARTER = "BillingQuarter"
    BILLING_ANNUAL = "BillingAnnual"
