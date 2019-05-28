package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/nalej/derrors"
	"github.com/nalej/edge-inventory-proxy/internal/app/sidecar/config"
	"github.com/nalej/grpc-network-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os/exec"
	"time"
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
	dhcclientCmd = "/sbin/dhclient"
)

// VpnHelper structure to interface with the VPN client command to join the VPN.
type VpnHelper struct{
	config config.Config
}

func NewVpnHelper (config config.Config) *VpnHelper{
	return &VpnHelper{config}
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

	log.Info().Str("user", h.config.Username).Msg("Configuring Local VPN")

	// NicCreate
	err := h.execCmd(command, cmdMode, vpnClientAddress, cmdCmd, nicCreateCmd, nicName)
	if err != nil {
		log.Info().Str("error", err.Error()).Msg("error creating nicName")
	}
	vpnUserName := fmt.Sprintf("/USERNAME:%s", h.config.Username)

	// Account Create
	vpnServer := fmt.Sprintf("/SERVER:%s", h.config.VPNAddress)
	err = h.execCmd(command, cmdMode, vpnClientAddress,cmdCmd, accountCreateCmd, h.config.Username, vpnServer, hub, vpnUserName, nicUser)
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error creating account")
	}

	// Account PasswordSet
	pass := fmt.Sprintf("/PASSWORD:%s", h.config.Password)
	err = h.execCmd(command, cmdMode, vpnClientAddress,cmdCmd, accountPasswordSetCmd, h.config.Username, pass, "/TYPE:standard")
	if err != nil {
		log.Warn().Str("error", err.Error()).Msg("error creating password")
	}

	// Connect
	err = h.execCmd(command, cmdMode, vpnClientAddress,cmdCmd, "accountConnect", h.config.Username)
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
	log.Info().Str("user", h.config.Username).Msg("connected")

	return nil
}

func (h * VpnHelper) getVPNNicName() string{
	return fmt.Sprintf("vpn_%s", nicName)
}

func (h * VpnHelper) getVPNAddress() (*string, error){
	iface, err := net.InterfaceByName(h.getVPNNicName())
	if err != nil{
		return nil, err
	}

	addresses, err := iface.Addrs()
	if err != nil{
		return nil, err
	}
	for _, addr := range addresses{
		netIP, ok := addr.(*net.IPNet)
		if ok && !netIP.IP.IsLoopback() && netIP.IP.To4() != nil{
			ip := netIP.IP.String()
			return &ip, nil
		}
	}
	return nil, errors.New("cannot retrieve address list")
}

func (h * VpnHelper) RegisterVPNAddress() error{
	ip, err := h.getVPNAddress()
	if err != nil{
		return err
	}

	conn, err := grpc.Dial(h.config.NetworkManagerAddress, grpc.WithInsecure())
	if err != nil {
		return derrors.AsError(err, "cannot create connection with infrastructure monitor coordinator")
	}
	netManagerClient := grpc_network_go.NewServiceDNSClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	addRequest := &grpc_network_go.AddServiceDNSEntryRequest{
		OrganizationId:       "management",
		Fqdn:                 fmt.Sprintf("%s-vpn", h.config.ProxyName),
		Ip:                   *ip,
		Tags:                 []string{"EIC_PROXY"},
	}
	log.Debug().Interface("request", addRequest).Msg("registering entry")
	_, err = netManagerClient.AddEntry(ctx, addRequest)

	if err != nil{
		dErr := conversions.ToDerror(err)
		log.Error().Str("trace", dErr.DebugReport()).Msg("cannot register edge inventory proxy IP on the DNS")
		return err
	}
	return nil
}