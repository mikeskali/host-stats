package main

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/process"
	"github.com/caarlos0/env/v6"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
)



func main() {

	cfg := config{Engine: &engine.Config{}}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}
	logger := newLogger(cfg.Env)

	hi, _ := host.Info()
	printJson(hi, "Host Info")


	v, _ := mem.VirtualMemory()

	fmt.Println("====Virtual Memory====")
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	printJson(v, "Virtual memory")

	sm, _ := mem.SwapMemory()
	fmt.Println(sm, "swap memory")

	cput, _ := cpu.Times(true)
	printJson(cput, "CPU times")

	cpui,_ := cpu.Info()
	printJson(cpui, "CPU INFO")

	disku, _ := disk.Usage("/")
	printJson(disku, "Disk Usage")

	fmt.Println("====load avg====")
	la,_ := load.Avg()
	printJson(la, "Load Average")

	ni,_ := net.Interfaces()
	printJson(ni, "Network interfaces")

	nioc,_ := net.IOCounters(true)
	printJson(nioc, "net io counters")

	npc,_ := net.ProtoCounters([]string{})
	printJson(npc, "net protocl counters")

	ppi, _ := process.Pids()
	printJson(ppi, "process pids")

	pps,_ := process.Processes()
	printJson(pps, "process processes")
}

func printJson(data interface{}, title string){
	pretty, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("====================%s======================\n %s", title, pretty)
}