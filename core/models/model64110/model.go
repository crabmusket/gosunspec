// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model64110

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block64110 - OutBack AXS device -

const (
	ModelID = 64110
)

const (
	AXS_Error         = "AXS_Error"
	AXS_Spare         = "AXS_Spare"
	AXS_Status        = "AXS_Status"
	Alarm_email_addr1 = "Alarm_email_addr1"
	Alarm_email_addr2 = "Alarm_email_addr2"
	Alarm_email_en    = "Alarm_email_en"
	Alarm_email_sub   = "Alarm_email_sub"
	Ambient_temp      = "Ambient_temp"
	Battery_temp      = "Battery_temp"
	DNS1_address      = "DNS1_address"
	DNS2_address      = "DNS2_address"
	Date_Day          = "Date_Day"
	Date_month        = "Date_month"
	Date_year         = "Date_year"
	EnableDHCP        = "EnableDHCP"
	EncrypKey         = "EncrypKey"
	FTP_password      = "FTP_password"
	Gateway_address   = "Gateway_address"
	Log_mode          = "Log_mode"
	Log_retain        = "Log_retain"
	Log_write_int     = "Log_write_int"
	MAC_Address       = "MAC_Address"
	MajorFWRev        = "MajorFWRev"
	MidFWRev          = "MidFWRev"
	MinorFWRev        = "MinorFWRev"
	Modbus_port       = "Modbus_port"
	NTP_enable        = "NTP_enable"
	NTP_server_nm     = "NTP_server_nm"
	SMTP_account_nm   = "SMTP_account_nm"
	SMTP_enable_SSL   = "SMTP_enable_SSL"
	SMTP_password     = "SMTP_password"
	SMTP_server_nm    = "SMTP_server_nm"
	SMTP_user_nm      = "SMTP_user_nm"
	Stat_email_addr1  = "Stat_email_addr1"
	Stat_email_addr2  = "Stat_email_addr2"
	Stat_email_int    = "Stat_email_int"
	Stat_email_sub    = "Stat_email_sub"
	Stat_start_HR     = "Stat_start_HR"
	TCPIP_Netmask     = "TCPIP_Netmask"
	TCPIP_address     = "TCPIP_address"
	TELNET_password   = "TELNET_password"
	Temp_SF           = "Temp_SF"
	TimeZone          = "TimeZone"
	Time_hour         = "Time_hour"
	Time_minute       = "Time_minute"
	Time_second       = "Time_second"
	WritePassword     = "WritePassword"
)

type Block64110 struct {
	MajorFWRev        uint16           `sunspec:"offset=0"`
	MidFWRev          uint16           `sunspec:"offset=1"`
	MinorFWRev        uint16           `sunspec:"offset=2"`
	EncrypKey         uint16           `sunspec:"offset=3"`
	MAC_Address       core.String      `sunspec:"offset=4,len=7"`
	WritePassword     core.String      `sunspec:"offset=11,len=8"`
	EnableDHCP        core.Enum16      `sunspec:"offset=19"`
	TCPIP_address     core.Ipaddr      `sunspec:"offset=20"`
	Gateway_address   core.Ipaddr      `sunspec:"offset=22"`
	TCPIP_Netmask     core.Ipaddr      `sunspec:"offset=24"`
	DNS1_address      core.Ipaddr      `sunspec:"offset=26"`
	DNS2_address      core.Ipaddr      `sunspec:"offset=28"`
	Modbus_port       uint16           `sunspec:"offset=30"`
	SMTP_server_nm    core.String      `sunspec:"offset=31,len=20"`
	SMTP_account_nm   core.String      `sunspec:"offset=51,len=16"`
	SMTP_enable_SSL   core.Enum16      `sunspec:"offset=67"`
	SMTP_password     core.String      `sunspec:"offset=68,len=8"`
	SMTP_user_nm      core.String      `sunspec:"offset=76,len=20"`
	Stat_email_int    uint16           `sunspec:"offset=96"`
	Stat_start_HR     uint16           `sunspec:"offset=97"`
	Stat_email_sub    core.String      `sunspec:"offset=98,len=25"`
	Stat_email_addr1  core.String      `sunspec:"offset=123,len=20"`
	Stat_email_addr2  core.String      `sunspec:"offset=143,len=20"`
	Alarm_email_en    core.Enum16      `sunspec:"offset=163"`
	Alarm_email_sub   core.String      `sunspec:"offset=164,len=25"`
	Alarm_email_addr1 core.String      `sunspec:"offset=189,len=20"`
	Alarm_email_addr2 core.String      `sunspec:"offset=209,len=20"`
	FTP_password      core.String      `sunspec:"offset=229,len=8"`
	TELNET_password   core.String      `sunspec:"offset=237,len=8"`
	Log_write_int     uint16           `sunspec:"offset=245"`
	Log_retain        uint16           `sunspec:"offset=246"`
	Log_mode          core.Enum16      `sunspec:"offset=247"`
	NTP_server_nm     core.String      `sunspec:"offset=248,len=20"`
	NTP_enable        core.Enum16      `sunspec:"offset=268"`
	TimeZone          int16            `sunspec:"offset=269"`
	Date_year         uint16           `sunspec:"offset=270"`
	Date_month        uint16           `sunspec:"offset=271"`
	Date_Day          uint16           `sunspec:"offset=272"`
	Time_hour         uint16           `sunspec:"offset=273"`
	Time_minute       uint16           `sunspec:"offset=274"`
	Time_second       uint16           `sunspec:"offset=275"`
	Battery_temp      int16            `sunspec:"offset=276,sf=Temp_SF"`
	Ambient_temp      int16            `sunspec:"offset=277,sf=Temp_SF"`
	Temp_SF           core.ScaleFactor `sunspec:"offset=278"`
	AXS_Error         core.Bitfield16  `sunspec:"offset=279"`
	AXS_Status        core.Bitfield16  `sunspec:"offset=280"`
	AXS_Spare         uint16           `sunspec:"offset=281"`
}

func (self *Block64110) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 282,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 282,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: MajorFWRev, Offset: 0, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: MidFWRev, Offset: 1, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: MinorFWRev, Offset: 2, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: EncrypKey, Offset: 3, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: MAC_Address, Offset: 4, Type: "string", Length: 7, Mandatory: true},
					smdx.PointElement{Id: WritePassword, Offset: 11, Type: "string", Length: 8, Mandatory: true},
					smdx.PointElement{Id: EnableDHCP, Offset: 19, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: TCPIP_address, Offset: 20, Type: "ipaddr", Mandatory: true},
					smdx.PointElement{Id: Gateway_address, Offset: 22, Type: "ipaddr", Mandatory: true},
					smdx.PointElement{Id: TCPIP_Netmask, Offset: 24, Type: "ipaddr", Mandatory: true},
					smdx.PointElement{Id: DNS1_address, Offset: 26, Type: "ipaddr", Mandatory: true},
					smdx.PointElement{Id: DNS2_address, Offset: 28, Type: "ipaddr", Mandatory: true},
					smdx.PointElement{Id: Modbus_port, Offset: 30, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: SMTP_server_nm, Offset: 31, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: SMTP_account_nm, Offset: 51, Type: "string", Length: 16, Mandatory: true},
					smdx.PointElement{Id: SMTP_enable_SSL, Offset: 67, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: SMTP_password, Offset: 68, Type: "string", Length: 8, Mandatory: true},
					smdx.PointElement{Id: SMTP_user_nm, Offset: 76, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: Stat_email_int, Offset: 96, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Stat_start_HR, Offset: 97, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Stat_email_sub, Offset: 98, Type: "string", Length: 25, Mandatory: true},
					smdx.PointElement{Id: Stat_email_addr1, Offset: 123, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: Stat_email_addr2, Offset: 143, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: Alarm_email_en, Offset: 163, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: Alarm_email_sub, Offset: 164, Type: "string", Length: 25, Mandatory: true},
					smdx.PointElement{Id: Alarm_email_addr1, Offset: 189, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: Alarm_email_addr2, Offset: 209, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: FTP_password, Offset: 229, Type: "string", Length: 8, Mandatory: true},
					smdx.PointElement{Id: TELNET_password, Offset: 237, Type: "string", Length: 8, Mandatory: true},
					smdx.PointElement{Id: Log_write_int, Offset: 245, Type: "uint16", Units: "Tms", Mandatory: true},
					smdx.PointElement{Id: Log_retain, Offset: 246, Type: "uint16", Units: "Tmd", Mandatory: true},
					smdx.PointElement{Id: Log_mode, Offset: 247, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: NTP_server_nm, Offset: 248, Type: "string", Length: 20, Mandatory: true},
					smdx.PointElement{Id: NTP_enable, Offset: 268, Type: "enum16", Mandatory: true},
					smdx.PointElement{Id: TimeZone, Offset: 269, Type: "int16", Units: "Tmh", Mandatory: true},
					smdx.PointElement{Id: Date_year, Offset: 270, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Date_month, Offset: 271, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Date_Day, Offset: 272, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Time_hour, Offset: 273, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Time_minute, Offset: 274, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Time_second, Offset: 275, Type: "uint16", Mandatory: true},
					smdx.PointElement{Id: Battery_temp, Offset: 276, Type: "int16", ScaleFactor: "Temp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: Ambient_temp, Offset: 277, Type: "int16", ScaleFactor: "Temp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: Temp_SF, Offset: 278, Type: "sunssf", Mandatory: true},
					smdx.PointElement{Id: AXS_Error, Offset: 279, Type: "bitfield16", Mandatory: true},
					smdx.PointElement{Id: AXS_Status, Offset: 280, Type: "bitfield16", Mandatory: true},
					smdx.PointElement{Id: AXS_Spare, Offset: 281, Type: "uint16", Mandatory: true},
				},
			},
		}})
}
