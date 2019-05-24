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
	nicName ="nicName"
	nicUser ="/NICNAME:nicname"
	accountCreateCmd = "AccountCreate"
	accountPasswordSetCmd = "AccountPasswordSet"
	vpnClientAddress = "localhost"
	user = "admin"
	password = "admin"
)

type VpnHelper struct{
	vpnServerAddress string
}

func NewVpnHelper (dir string) *VpnHelper{

	return &VpnHelper{dir}
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
	vpnServer := fmt.Sprintf("/SERVER:%s", j.vpnServerAddress)
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
		//return err
	}

	log.Info().Str("user", user).Msg("connected")

	return nil
}