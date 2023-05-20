package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulteary/sparrow/internal/define"
)

type NewTempProcessorTypes struct {
	HasCustomerObject     bool `json:"has_customer_object,omitempty"`
	HasTransactionHistory bool `json:"has_transaction_history,omitempty"`
}

type NewTempProcessor struct {
	A001 NewTempProcessorTypes `json:"a001"`
	B001 NewTempProcessorTypes `json:"b001"`
}

type NewTempAccount struct {
	AccountUserRole                       string           `json:"account_user_role"`
	AccountUserID                         string           `json:"account_user_id"`
	Processor                             NewTempProcessor `json:"processor"`
	AccountID                             string           `json:"account_id"`
	IsMostRecentExpiredSubscriptionGratis bool             `json:"is_most_recent_expired_subscription_gratis"`
	HasPreviouslyPaidSubscription         bool             `json:"has_previously_paid_subscription"`
}

type NewTempEntitlement struct {
	SubscriptionID        string `json:"subscription_id"`
	HasActiveSubscription bool   `json:"has_active_subscription"`
	SubscriptionPlan      string `json:"subscription_plan"`
	ExpiresAt             string `json:"expires_at"`
}

type NewTempLastActiveSubscription struct {
	SubscriptionID         string `json:"subscription_id"`
	PurchaseOriginPlatform string `json:"purchase_origin_platform"`
	WillRenew              bool   `json:"will_renew"`
}

type NewTempDefault struct {
	Account                NewTempAccount                `json:"account"`
	Features               []string                      `json:"features"`
	Entitlement            NewTempEntitlement            `json:"entitlement"`
	LastActiveSubscription NewTempLastActiveSubscription `json:"last_active_subscription"`
}

type NewTempAccounts struct {
	Default NewTempDefault `json:"default"`
}

type NewTempCheck struct {
	Accounts          NewTempAccounts `json:"accounts"`
	TempApAvailableAt string          `json:"temp_ap_available_at"`
}

func AccountTempCheck(c *gin.Context) {
	data := NewTempCheck{
		Accounts: NewTempAccounts{
			Default: NewTempDefault{
				Account: NewTempAccount{
					AccountUserRole: "account-owner",
					AccountUserID:   define.GenerateUUID(),
					Processor: NewTempProcessor{
						A001: NewTempProcessorTypes{
							HasCustomerObject: true,
						},
						B001: NewTempProcessorTypes{
							HasTransactionHistory: false,
						},
					},
					AccountID:                             define.GenerateUUID(),
					IsMostRecentExpiredSubscriptionGratis: false,
					HasPreviouslyPaidSubscription:         true,
				},
				Features: GetFeatures(),
				Entitlement: NewTempEntitlement{
					SubscriptionID:        define.GenerateUUID(),
					HasActiveSubscription: true,
					SubscriptionPlan:      "chatgptplusplan",
					ExpiresAt:             "2099-12-31T23:59:00+00:00",
				},
				LastActiveSubscription: NewTempLastActiveSubscription{
					SubscriptionID:         define.GenerateUUID(),
					PurchaseOriginPlatform: "chatgpt_web",
					WillRenew:              true,
				},
			},
		},
		TempApAvailableAt: "2099-12-31T23:59:00+00:00",
	}

	if !define.ENABLE_PAID_SUBSCRIPTION {
		data.Accounts.Default.Entitlement.HasActiveSubscription = true
		data.Accounts.Default.Entitlement.SubscriptionPlan = "chatgptplusplan" // or: "chatgptplusfreeplan"
		data.Accounts.Default.Account.Processor.A001.HasCustomerObject = true
	} else {
		data.Accounts.Default.Entitlement.HasActiveSubscription = false
		data.Accounts.Default.Entitlement.SubscriptionPlan = ""
		data.Accounts.Default.Account.Processor.A001.HasCustomerObject = false
	}

	c.JSON(http.StatusOK, data)
}
