package msig

import (
	"github.com/apepenkov/eos-go-fixed-p2p"
)

func init() {
	eos.RegisterAction(AN("eosio.msig"), ActN("propose"), &Propose{})
	eos.RegisterAction(AN("eosio.msig"), ActN("approve"), &Approve{})
	eos.RegisterAction(AN("eosio.msig"), ActN("unapprove"), &Unapprove{})
	eos.RegisterAction(AN("eosio.msig"), ActN("cancel"), &Cancel{})
	eos.RegisterAction(AN("eosio.msig"), ActN("exec"), &Exec{})
}

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN
