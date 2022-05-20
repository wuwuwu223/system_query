package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"runtime"
	"strconv"
	"time"
)

func main() {
	var sid string
	var ethx string
	flag.StringVar(&sid, "s", "我的mac", "服务器名字")
	flag.StringVar(&ethx, "i", "en0", "网卡名")
	flag.Parse()
	txlast, rxlast, txplast, rxplast := getTxRx(ethx)
	for {
		time.Sleep(1 * time.Second)
		go func() {
			v, _ := mem.VirtualMemory()
			info, _ := cpu.Info()
			c, _ := cpu.Percent(0*time.Second, false)
			uptime, _ := host.Uptime()
			arch, _ := host.KernelArch()
			version, _ := host.KernelVersion()
			d, _ := disk.Usage("/")
			tcpConns, _ := net.Connections("tcp")
			udpConns, _ := net.Connections("udp")
			swap, _ := mem.SwapMemory()
			txnow, rxnow, txpnow, rxpnow := getTxRx(ethx)
			server := &Server{
				Name:              sid,
				Uptime:            uptime,
				NetworkRx:         rxnow - rxlast,
				NetworkTx:         txnow - txlast,
				NetworkIn:         rxnow,
				NetworkOut:        txnow,
				RxpNow:            rxpnow - rxplast,
				TxpNow:            txpnow - txplast,
				Rxp:               rxpnow,
				Txp:               txpnow,
				Cpu:               strconv.Itoa(runtime.NumCPU()) + "*" + info[0].ModelName,
				CpuUsedPercent:    c[0],
				CpuVersion:        version,
				CpuArch:           arch,
				MemoryTotal:       v.Total,
				MemoryUsedPercent: v.UsedPercent,
				SwapTotal:         swap.Total,
				SwapUsedPercent:   swap.UsedPercent,
				HddTotal:          d.Total,
				HddUsedPercent:    d.UsedPercent,
				TcpNum:            len(tcpConns),
				UdpNum:            len(udpConns),
			}
			str, _ := json.Marshal(server)
			fmt.Println(string(str))
			rxlast = rxnow
			txlast = txnow
		}()
	}
}

func getTxRx(name string) (uint64, uint64, uint64, uint64) {
	counters, _ := net.IOCounters(true)
	for i := range counters {
		if counters[i].Name == name {
			return counters[i].BytesRecv, counters[i].BytesSent, counters[i].PacketsRecv, counters[i].PacketsSent
		}
	}
	return 0, 0, 0, 0
}

type Server struct {
	Name              string  `json:"name"`
	Uptime            uint64  `json:"uptime"`
	NetworkRx         uint64  `json:"network_rx"`
	NetworkTx         uint64  `json:"network_tx"`
	NetworkIn         uint64  `json:"network_in"`
	NetworkOut        uint64  `json:"network_out"`
	Cpu               string  `json:"cpu"`
	CpuUsedPercent    float64 `json:"cpu_used_percent"`
	MemoryTotal       uint64  `json:"memory_total"`
	MemoryUsedPercent float64 `json:"memory_used_percent"`
	SwapTotal         uint64  `json:"swap_total"`
	SwapUsedPercent   float64 `json:"swap_used_percent"`
	HddTotal          uint64  `json:"hdd_total"`
	HddUsedPercent    float64 `json:"hdd_used_percent"`
	RxpNow            uint64  `json:"rxp_now"`
	TxpNow            uint64  `json:"txp_now"`
	CpuVersion        string  `json:"cpu_version"`
	CpuArch           string  `json:"cpu_arch"`
	Rxp               uint64  `json:"rxp"`
	Txp               uint64  `json:"txp"`
	TcpNum            int     `json:"tcp_num"`
	UdpNum            int     `json:"udp_num"`
}
