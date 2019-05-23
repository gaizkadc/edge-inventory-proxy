package eip_sidecar

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
)

const (
	command = "vpncmd"
	cmdMode = "/Client"
	hub = "/HUB:DEFAULT"
	cmdCmd = "/cmd"
	nicCreateCmd = "NicCreate"
	nicName ="nicname"
	nicUser ="/NICNAME:nicname"
	accountCreateCmd = "AccountCreate"
	accountPasswordSetCmd = "AccountPasswordSet"
	vpnClientAddress = "localhost"
	user = "admin"
	password = "admin"
	dhcclientCmd = "/sbin/dhclient"
)

// VpnHelper structure to interface with the VPN client command to join the VPN.
type VpnHelper struct{
	vpnServerAddress string
}

func NewVpnHelper (address string) *VpnHelper{
	return &VpnHelper{address}
}

// execCmd executes a given command on the command line.
func (h * VpnHelper) execCmd(cmdName string, args ...string) error {
	cmd := exec.Command(cmdName, args...)
	output, err := cmd.CombinedOutput()
	log.Warn().Str("output", string(output)).Msg("Command output")
	if err != nil{
		log.Warn().Str("cmd", cmdName).Strs("args", args).Str("error", err.Error()).Msg("cannot execute command")
		return err
	}
	return nil
}

// ConfigureLocalVPN triggers the steps required to configure the VPN connection.
func (h * VpnHelper) ConfigureLocalVPN () error {

	log.Info().Str("user", user).Msg("Configuring Local VPN")

	// NicCreate
	err := h.execCmd(command, cmdMode, vpnClientAddress, cmdCmd, nicCreateCmd, nicName)
	if err != nil {
		log.Info().Str("error", err.Error()).Msg("error creating nicName")
	}
	vpnUserName := fmt.Sprintf("/USERNAME:%s", user)

	// Account Create
	vpnServer := fmt.Sprintf("/SERVER:%s", j.vpnServerAddress)
	err = h.execCmd(command, cmdMode, vpnClientAddress,cmdCmd, accountCreateCmd, user, vpnServer, hub, vpnUserName, nicUser)
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error creating account")
	}

	// Account PasswordSet
	pass := fmt.Sprintf("/PASSWORD:%s", password)
	err = h.execCmd(command, cmdMode, vpnClientAddress,cmdCmd, accountPasswordSetCmd, password, pass, "/TYPE:standard")
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error creating password")
	}

	// Connect
	err = h.execCmd(command, cmdMode, vpnClientAddress,cmdCmd, "accountConnect", user)
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error connecting account")
		return err
	}

	// Get VPN IP
	vpnNicName := fmt.Sprintf("vpn_%s", nicName)
	err = h.execCmd(dhcclientCmd, vpnNicName)
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error getting VPN IP")
		return err
	}

	// Show the address
	err = h.execCmd("/sbin/ip", "addr", "show", "dev", vpnNicName)
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error getting IP addr")
		return err
	}

	// Success!
	log.Info().Str("user", user).Msg("connected")

	return nil
}