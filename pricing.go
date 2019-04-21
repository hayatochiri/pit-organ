package pitOrgan

// Receivers

type ReceiverPricing struct {
	AccountID  string
	Connection *Connection
}

type ReceiverPricingStream struct {
	AccountID  string
	Connection *Connection
}

func (r *ReceiverAccountID) Pricing() *ReceiverPricing {
	return &ReceiverPricing{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// TODO: GET /v3/accounts/{accountID}/pricing

func (r *ReceiverPricing) Stream() *ReceiverPricingStream {
	return &ReceiverPricingStream{
		AccountID:  r.AccountID,
		Connection: r.Connection,
	}
}

// TODO: GET /v3/accounts/{accountID}/pricing/stream

