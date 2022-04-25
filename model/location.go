package model

import (
	"encoding/json"
	"strings"
)

type Location struct {
	ID        string    `json:"id"`
	IDOwner   string    `json:"id_owner"`
	OwnerKind OwnerKind `json:"owner_kind"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	CreateAt  int64     `json:"create_at"`
}

type OwnerKind uint8

const (
	OwnerKindUnknown OwnerKind = iota
	OwnerKindTaxi
	OwnerKindUser
)

func (ok *OwnerKind) UnmarshalJSON(b []byte) error {
	var ks string
	if err := json.Unmarshal(b, &ks); err != nil {
		return err
	}

	switch strings.ToLower(ks) {
	case "taxi":
		*ok = OwnerKindTaxi
	case "user":
		*ok = OwnerKindUser
	default:
		*ok = OwnerKindUnknown
	}

	return nil
}

func (ok OwnerKind) MarshalJSON() ([]byte, error) {
	var s string
	switch ok {
	case OwnerKindTaxi:
		s = "taxi"
	case OwnerKindUser:
		s = "user"
	default:
		s = "unknown"
	}

	return json.Marshal(s)
}
