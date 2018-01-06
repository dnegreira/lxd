package main

type firewallType int

const (
	firewallTypeBridge firewallType = iota
)

func initFirewallDriver