package sudo

import eos "github.com/apepenkov/eos-go-fixed-p2p"

func init() {
	eos.RegisterAction(AN("eosio.wrap"), ActN("exec"), Exec{})
}

var AN = eos.AN
var ActN = eos.ActN
