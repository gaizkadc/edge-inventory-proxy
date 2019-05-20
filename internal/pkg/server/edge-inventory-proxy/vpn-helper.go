package edge_inventory_proxy

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
)

const (
	command = "/usr/bin/vpnclient/vpncmd"
	cmdMode = "/Client"
	hub = "/HUB:DEFAULT"
	cmdCmd = "/cmd"
	nicCreateCmd = "NicCreate"
	nicName ="nicName"
	nicUser ="/NICNAME:nicname"
	accountCreateCmd = "AccountCreate"
	accountPasswordSetCmd = "AccountPasswordSet"
	vpnClientAddress = "localhost"
	user = "admin"
	password = "admin"
	vpnServer = "/SERVER:vpn-server.svc.cluster.local:5555"
)

type VpnHelper struct{

}

func NewVpnHelper () *VpnHelper{

	return &VpnHelper{}
}


func (j * VpnHelper) ConfigureLocalVPN () error {

	log.Info().Str("user", user).Msg("Configuring Local VPN")

	// NicCreate
	cmd := exec.Command(command, cmdMode, vpnClientAddress, cmdCmd, nicCreateCmd, nicName)
	err := cmd.Run()
	if err != nil {
		log.Info().Str("error", err.Error()).Msg("error creating nicName")
	}
	vpnUserName := fmt.Sprintf("/USERNAME:%s", user)

	// Account Create
	cmd = exec.Command(command, cmdMode, vpnClientAddress,cmdCmd, accountCreateCmd, user, vpnServer, hub, vpnUserName, nicUser)
	err = cmd.Run()
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error creating account")
	}

	// Account PasswordSet
	pass := fmt.Sprintf("/PASSWORD:%s", password)
	cmd = exec.Command(command, cmdMode, vpnClientAddress,cmdCmd, accountPasswordSetCmd, password, pass, "/TYPE:standard")
	err = cmd.Run()
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error creating password")
	}

	// Connect
	cmd = exec.Command(command, cmdMode, vpnClientAddress,cmdCmd, "accountConnect", user)
	err = cmd.Run()
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error connecting account")
		return err
	}

	log.Info().Str("user", user).Msg("connected")

	return nil
}