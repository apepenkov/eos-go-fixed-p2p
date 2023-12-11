package rex

import (
	eos "github.com/apepenkov/eos-go-fixed-p2p"
)

func NewSellREX(from eos.AccountName, rex eos.Asset) *eos.Action {
	return &eos.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []eos.PermissionLevel{
			{Actor: from, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From eos.AccountName
	REX  eos.Asset
}
