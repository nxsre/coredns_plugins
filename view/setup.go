package view

import (
	"github.com/coredns/coredns/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/mholt/caddy"
)

// init 中注册插件
func init() {
	caddy.RegisterPlugin("view", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

// setup 函数
func setup(c *caddy.Controller) error {
	var err error

	for c.Next() { // Skip the directive name. 跳过指令名称
		if !c.NextArg() { // Expect at least one value. 至少一个预期的值
			return c.ArgErr() // Otherwise it's an error. 否则返回错误
		}
		value := c.Val() // Use the value.
	}

	cfg := dnsserver.GetConfig(c)
	mid := func(next plugin.Handler) plugin.Handler {
		return View{Next: next}
	}
	cfg.AddPlugin(mid)

	if err != nil {
		return plugin.Error("view", err)
	}
	return nil
}
