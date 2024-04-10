package common

// This file contains components from github.com/brlbil/wpaclient
//
// Copyright (c) 2017 Birol Bilgin
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// nd/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// UT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

const (
	CmdWpaStatus               = "STATUS"
	CmdWpaIfname               = "IFNAME"
	CmdWpaPing                 = "PING"
	CmdWpaRelog                = "RELOG"
	CmdWpaNote                 = "NOTE"
	CmdWpaMib                  = "MIB"
	CmdWpaHelp                 = "HELP"
	CmdWpaInterface            = "INTERFACE"
	CmdWpaLevel                = "LEVEL"
	CmdWpaLicense              = "LICENSE"
	CmdWpaQuit                 = "QUIT"
	CmdWpaSet                  = "SET"
	CmdWpaDump                 = "DUMP"
	CmdWpaGet                  = "GET"
	CmdWpaDriverFlags          = "DRIVER_FLAGS"
	CmdWpaLogon                = "LOGON"
	CmdWpaLogoff               = "LOGOFF"
	CmdWpaPmksa                = "PMKSA"
	CmdWpaPmksaFlush           = "PMKSA_FLUSH"
	CmdWpaReassociate          = "REASSOCIATE"
	CmdWpaReattach             = "REATTACH"
	CmdWpaPreauthenticate      = "PREAUTHENTICATE"
	CmdWpaIdentity             = "IDENTITY"
	CmdWpaPassword             = "PASSWORD"
	CmdWpaNewPassword          = "NEW_PASSWORD"
	CmdWpaPin                  = "PIN"
	CmdWpaOtp                  = "OTP"
	CmdWpaPassphrase           = "PASSPHRASE"
	CmdWpaSim                  = "SIM"
	CmdWpaBssid                = "BSSID"
	CmdWpaBlacklist            = "BLACKLIST"
	CmdWpaLogLevel             = "LOG_LEVEL"
	CmdWpaListNetworks         = "LIST_NETWORKS"
	CmdWpaSelectNetwork        = "SELECT_NETWORK"
	CmdWpaEnableNetwork        = "ENABLE_NETWORK"
	CmdWpaDisableNetwork       = "DISABLE_NETWORK"
	CmdWpaAddNetwork           = "ADD_NETWORK"
	CmdWpaRemoveNetwork        = "REMOVE_NETWORK"
	CmdWpaSetNetwork           = "SET_NETWORK"
	CmdWpaGetNetwork           = "GET_NETWORK"
	CmdWpaDupNetwork           = "DUP_NETWORK"
	CmdWpaListCreds            = "LIST_CREDS"
	CmdWpaAddCred              = "ADD_CRED"
	CmdWpaRemoveCred           = "REMOVE_CRED"
	CmdWpaSetCred              = "SET_CRED"
	CmdWpaGetCred              = "GET_CRED"
	CmdWpaSaveConfig           = "SAVE_CONFIG"
	CmdWpaDisconnect           = "DISCONNECT"
	CmdWpaReconnect            = "RECONNECT"
	CmdWpaScan                 = "SCAN"
	CmdWpaScanResults          = "SCAN_RESULTS"
	CmdWpaAbortScan            = "ABORT_SCAN"
	CmdWpaBss                  = "BSS"
	CmdWpaGetCapability        = "GET_CAPABILITY"
	CmdWpaReconfigure          = "RECONFIGURE"
	CmdWpaTerminate            = "TERMINATE"
	CmdWpaInterfaceAdd         = "InterfaceAdd"
	CmdWpaInterfaceRemove      = "INTERFACE_REMOVE"
	CmdWpaInterfaceList        = "INTERFACE_LIST"
	CmdWpaApScan               = "AP_SCAN"
	CmdWpaScanInterval         = "SCAN_INTERVAL"
	CmdWpaBssExpireAge         = "BSS_EXPIRE_AGE"
	CmdWpaBssExpireCount       = "BSS_EXPIRE_COUNT"
	CmdWpaBssFlush             = "BSS_FLUSH"
	CmdWpaStkstart             = "STKSTART"
	CmdWpaFtDs                 = "FT_DS"
	CmdWpaWpsPbc               = "WPS_PBC"
	CmdWpaWpsPin               = "WPS_PIN"
	CmdWpaWpsCheckPin          = "WPS_CHECK_PIN"
	CmdWpaWpsReg               = "WPS_REG"
	CmdWpaWpsApPin             = "WPS_AP_PIN"
	CmdWpaWpsErStart           = "WPS_ER_START"
	CmdWpaWpsErStop            = "WPS_ER_STOP"
	CmdWpaWpsErPin             = "WPS_ER_PIN"
	CmdWpaWpsErPbc             = "WPS_ER_PBC"
	CmdWpaWpsErLearn           = "WPS_ER_LEARN"
	CmdWpaWpsErSetConfig       = "WPS_ER_SET_CONFIG"
	CmdWpaWpsErConfig          = "WPS_ER_CONFIG"
	CmdWpaIbssRsn              = "IBSS_RSN"
	CmdWpaSta                  = "STA"
	CmdWpaAllSta               = "ALL_STA"
	CmdWpaDeauthenticate       = "DEAUTHENTICATE"
	CmdWpaDisassociate         = "DISASSOCIATE"
	CmdWpaChanSwitch           = "CHAN_SWITCH"
	CmdWpaSuspend              = "SUSPEND"
	CmdWpaResume               = "RESUME"
	CmdWpaRoam                 = "ROAM"
	CmdWpaP2pFind              = "P2P_FIND"
	CmdWpaP2pStopFind          = "P2P_STOP_FIND"
	CmdWpaP2pAspProvision      = "P2P_ASP_PROVISION"
	CmdWpaP2pAspProvisionResp  = "P2P_ASP_PROVISION_RESP"
	CmdWpaP2pConnect           = "P2P_CONNECT"
	CmdWpaP2pListen            = "P2P_LISTEN"
	CmdWpaP2pGroupRemove       = "P2P_GROUP_REMOVE"
	CmdWpaP2pGroupAdd          = "P2P_GROUP_ADD"
	CmdWpaP2pGroupMember       = "P2P_GROUP_MEMBER"
	CmdWpaP2pProvDisc          = "P2P_PROV_DISC"
	CmdWpaP2pGetPassphrase     = "P2P_GET_PASSPHRASE"
	CmdWpaP2pServDiscReq       = "P2P_SERV_DISC_REQ"
	CmdWpaP2pServDiscCancelReq = "P2P_SERV_DISC_CANCEL_REQ"
	CmdWpaP2pServDiscResp      = "P2P_SERV_DISC_RESP"
	CmdWpaP2pServiceUpdate     = "P2P_SERVICE_UPDATE"
	CmdWpaP2pServDiscExternal  = "P2P_SERV_DISC_EXTERNAL"
	CmdWpaP2pServiceFlush      = "P2P_SERVICE_FLUSH"
	CmdWpaP2pServiceAdd        = "P2P_SERVICE_ADD"
	CmdWpaP2pServiceRep        = "P2P_SERVICE_REP"
	CmdWpaP2pServiceDel        = "P2P_SERVICE_DEL"
	CmdWpaP2pReject            = "P2P_REJECT"
	CmdWpaP2pInvite            = "P2P_INVITE"
	CmdWpaP2pPeers             = "P2P_PEERS"
	CmdWpaP2pPeer              = "P2P_PEER"
	CmdWpaP2pSet               = "P2P_SET"
	CmdWpaP2pFlush             = "P2P_FLUSH"
	CmdWpaP2pCancel            = "P2P_CANCEL"
	CmdWpaP2pUnauthorize       = "P2P_UNAUTHORIZE"
	CmdWpaP2pPresenceReq       = "P2P_PRESENCE_REQ"
	CmdWpaP2pExtListen         = "P2P_EXT_LISTEN"
	CmdWpaP2pRemoveClient      = "P2P_REMOVE_CLIENT"
	CmdWpaVendorElemAdd        = "VENDOR_ELEM_ADD"
	CmdWpaVendorElemGet        = "VENDOR_ELEM_GET"
	CmdWpaVendorElemRemove     = "VENDOR_ELEM_REMOVE"
	CmdWpaStaAutoconnect       = "STA_AUTOCONNECT"
	CmdWpaTdlsDiscover         = "TDLS_DISCOVER"
	CmdWpaTdlsSetup            = "TDLS_SETUP"
	CmdWpaTdlsTeardown         = "TDLS_TEARDOWN"
	CmdWpaTdlsLinkStatus       = "TDLS_LINK_STATUS"
	CmdWpaWmmAcAddts           = "WMM_AC_ADDTS"
	CmdWpaWmmAcDelts           = "WMM_AC_DELTS"
	CmdWpaWmmAcStatus          = "WMM_AC_STATUS"
	CmdWpaTdlsChanSwitch       = "TDLS_CHAN_SWITCH"
	CmdWpaTdlsCancelChanSwitch = "TDLS_CANCEL_CHAN_SWITCH"
	CmdWpaSignalPoll           = "SIGNAL_POLL"
	CmdWpaSignalMonitor        = "SIGNAL_MONITOR"
	CmdWpaPktcntPoll           = "PKTCNT_POLL"
	CmdWpaReauthenticate       = "REAUTHENTICATE"
	CmdWpaRaw                  = "RAW"
	CmdWpaFlush                = "FLUSH"
	CmdWpaRadioWork            = "RADIO_WORK"
	CmdWpaVendor               = "VENDOR"
	CmdWpaNeighborRepRequest   = "NEIGHBOR_REP_REQUEST"
	CmdWpaErpFlush             = "ERP_FLUSH"
	CmdWpaMacRandScan          = "MAC_RAND_SCAN"
	CmdWpaGetPrefFreqList      = "GET_PREF_FREQ_LIST"
	CmdWpaP2pLoStart           = "P2P_LO_START"
	CmdWpaP2pLoStop            = "P2P_LO_STOP"
)

const (
	CmdHostapdEnable = "ENABLE"
	CmdHostapdDisable = "DISABLE"
	CmdHostapdReload = "RELOAD"
	CmdHostapdListStations = "LIST_STA"
	CmdHostapdStatus= "STATUS"
	CmdHostapdDeauthenticate = "DEAUTHENTICATE"
	CmdHostapdDisassociate = "DISASSOCIATE"
)