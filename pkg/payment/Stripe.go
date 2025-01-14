package payment

import "github.com/stripe/stripe-go/v74"

type StripeService struct{}

func NewStripeService(apikey string) *StripeService {
	stripe.Key = apikey
	return &StripeService{}
}
func (s *StripeService) Charge(amount int64, currency, description, token string) error {
	return nil //incluir integração de pagamento posteriormente
}
