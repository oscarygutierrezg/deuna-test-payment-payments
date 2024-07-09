package enums

import (
	"encoding/json"
	"errors"
)

type PaymentStatus string

const (
	Unknown    = "Unknown"
	Pending    = "Pending"
	InProgress = "InProgress"
	Approved   = "Approved"
	Cancelled  = "Cancelled"
	Failed     = "Failed"
)

var statusToString = map[PaymentStatus]string{
	Unknown:    "Unknown",
	Pending:    "Pending",
	InProgress: "InProgress",
	Approved:   "Approved",
	Cancelled:  "Cancelled",
	Failed:     "Failed",
}

var stringToStatus = map[string]PaymentStatus{
	"Unknown":    Unknown,
	"Pending":    Pending,
	"InProgress": InProgress,
	"Approved":   Approved,
	"Cancelled":  Cancelled,
	"Failed":     Failed,
}

func (s PaymentStatus) String() string {
	return statusToString[s]
}

func Parse(string2 string) PaymentStatus {
	return stringToStatus[string2]
}

func (s *PaymentStatus) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}

	status, ok := stringToStatus[statusStr]
	if !ok {
		return errors.New("invalid status value")
	}

	*s = status
	return nil
}
