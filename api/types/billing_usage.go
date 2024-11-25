package types

import "github.com/google/uuid"

// ListCreditGrantsResponse returns the total remaining and granted credits for a customer.
type ListCreditGrantsResponse struct {
	RemainingBalanceCents int `json:"remaining_credits"`
	GrantedBalanceCents   int `json:"granted_credits"`
}

// ListCustomerUsageRequest is the request to list usage for a customer
type ListCustomerUsageRequest struct {
	// PreviousPeriods is the number of previous periods to include in the response.
	PreviousPeriods int `json:"previous_periods,omitempty"`
	// CurrentPeriod is whether to return only usage for the current billing period.
	CurrentPeriod bool `json:"current_period,omitempty"`
}

// Subscription is the subscription for a customer
type Subscription struct {
	ExternalID         string `json:"external_id"`
	ExternalCustomerID string `json:"external_customer_id"`
	Status             string `json:"status"`
	SubscriptionAt     string `json:"subscription_at"`
	EndingAt           string `json:"ending_at"`
}

// Usage is the aggregated usage for a customer
type Usage struct {
	FromDatetime     string        `json:"from_datetime"`
	ToDatetime       string        `json:"to_datetime"`
	TotalAmountCents int64         `json:"total_amount_cents"`
	ChargesUsage     []ChargeUsage `json:"charges_usage"`
}

// ChargeUsage is the usage for a charge
type ChargeUsage struct {
	Units          string         `json:"units"`
	AmountCents    int64          `json:"amount_cents"`
	AmountCurrency string         `json:"amount_currency"`
	BillableMetric BillableMetric `json:"billable_metric"`
}

// BillableMetric is the metric collected for billing
type BillableMetric struct {
	Name string `json:"name"`
}

// Plan is the plan for a customer
type Plan struct {
	ID           string `json:"id"`
	CustomerID   string `json:"customer_id"`
	StartingOn   string `json:"starting_on"`
	EndingBefore string `json:"ending_before"`
	TrialInfo    Trial  `json:"trial_info,omitempty"`
}

// Trial contains the information for a trial period
type Trial struct {
	EndingBefore string `json:"ending_before"`
}

// BillingEvent represents a Lago billing event.
type BillingEvent struct {
	CustomerID    string                 `json:"customer_id"`
	EventType     string                 `json:"event_type"`
	Properties    map[string]interface{} `json:"properties"`
	TransactionID string                 `json:"transaction_id"`
	Timestamp     string                 `json:"timestamp"`
}

// AppliedCoupon represents an applied coupon in the billing system.
type AppliedCoupon struct {
	Status                     string `json:"status"`
	FrequencyDuration          int    `json:"frequency_duration"`
	FrequencyDurationRemaining int    `json:"frequency_duration_remaining"`
	CreatedAt                  string `json:"created_at"`
}

// Wallet represents a customer credits wallet
type Wallet struct {
	LagoID                   uuid.UUID `json:"lago_id,omitempty"`
	Status                   string    `json:"status"`
	BalanceCents             int       `json:"balance_cents,omitempty"`
	CreditsOngoingBalance    string    `json:"credits_ongoing_balance,omitempty"`
	OngoingBalanceCents      int       `json:"ongoing_balance_cents,omitempty"`
	OngoingUsageBalanceCents int       `json:"ongoing_usage_balance_cents,omitempty"`
}

// Invoice represents an invoice in the billing system.
type Invoice struct {
	// The URL to view the hosted invoice.
	HostedInvoiceURL string `json:"hosted_invoice_url"`
	// The status of the invoice.
	Status string `json:"status"`
	// RFC 3339 timestamp for when the invoice was created.
	Created string `json:"created"`
}
