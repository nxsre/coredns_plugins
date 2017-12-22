package view

import (
	"fmt"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

type View struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (c View) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	fmt.Println(state.QClass(), state.QType())
	return 0, nil
}

// Name implements the Handler interface.
func (c View) Name() string { return "view" }
