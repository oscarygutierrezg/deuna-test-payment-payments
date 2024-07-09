package enums

import (
	"encoding/json"
	"errors"
)

type PaymentType string

const (
	Payment = "Payment"
	Refund  = "Refund"
)

var typeToString = map[PaymentType]string{
	Payment: "Payment",
	Refund:  "Refund",
}

var stringToType = map[string]PaymentType{
	"Payment": Payment,
	"Refund":  Refund,
}

func (s PaymentType) String() string {
	return typeToString[s]
}

func (s *PaymentType) UnmarshalJSON(data []byte) error {
	var typeStr string
	if err := json.Unmarshal(data, &typeStr); err != nil {
		return err
	}

	newType, ok := stringToType[typeStr]
	if !ok {
		return errors.New("invalid type value")
	}

	*s = newType
	return nil
}
