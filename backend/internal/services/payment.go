package services

import (
	"errors"
	"os"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
)

type PaymentService struct {
	stripeKey string
}

func NewPaymentService() *PaymentService {
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	stripe.Key = stripeKey

	return &PaymentService{
		stripeKey: stripeKey,
	}
}

type PaymentIntent struct {
	ID           string
	Amount       int64
	Currency     string
	Status       string
	ClientSecret string
}

func (s *PaymentService) CreatePaymentIntent(amount float64, currency string) (*PaymentIntent, error) {
	if s.stripeKey == "" {
		return nil, errors.New("Stripe key not configured")
	}

	// Convert amount to cents
	amountCents := int64(amount * 100)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountCents),
		Currency: stripe.String(currency),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return &PaymentIntent{
		ID:           pi.ID,
		Amount:       pi.Amount,
		Currency:     string(pi.Currency),
		Status:       string(pi.Status),
		ClientSecret: pi.ClientSecret,
	}, nil
}

func (s *PaymentService) ConfirmPayment(paymentIntentID string) error {
	pi, err := paymentintent.Get(paymentIntentID, nil)
	if err != nil {
		return err
	}

	if pi.Status != stripe.PaymentIntentStatusSucceeded {
		return errors.New("Payment not successful")
	}

	return nil
}

func (s *PaymentService) RefundPayment(paymentIntentID string) error {
	// TODO: Implement refund logic using Stripe refund API
	return nil
}

// PayPal integration would go here
type PayPalService struct {
	clientID     string
	clientSecret string
}

func NewPayPalService() *PayPalService {
	return &PayPalService{
		clientID:     os.Getenv("PAYPAL_CLIENT_ID"),
		clientSecret: os.Getenv("PAYPAL_CLIENT_SECRET"),
	}
}

func (s *PayPalService) CreateOrder(amount float64, currency string) (string, error) {
	// TODO: Implement PayPal order creation
	// This would use the PayPal SDK to create an order
	return "", errors.New("PayPal not implemented yet")
}

func (s *PayPalService) CaptureOrder(orderID string) error {
	// TODO: Implement PayPal order capture
	return errors.New("PayPal not implemented yet")
}
