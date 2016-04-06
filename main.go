package main

import (
    "fmt"
    "time"

    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
)

const (
    maxCp = 55
    maxMp = 55
    maxDp = 55

    maxHit = 5
    delay = time.Second * 2
)

var (
    cpHit = 0
    mpHit = 0
    dpHit = 0
)

func main() {
    for {
        check()

        time.Sleep(delay)
    }
}

func check() {
    m, _ := mem.VirtualMemory()
    c, _ := cpu.CPUPercent(time.Second * 1, false)
    d, _ := disk.DiskUsage("/")

    cp := int(c[0])
    mp := int(m.UsedPercent)
    dp := int(d.UsedPercent)

    fmt.Printf("======================\n")
    fmt.Printf("Cpu: %d%%\n", cp)
    fmt.Printf("Mem: %d%%\n", mp)
    fmt.Printf("Disk: %d%%\n", dp)
    fmt.Printf("======================\n")

    // execute hook when max limit reaches
    if cp > maxCp {
        if cpHit >= maxHit {
            cpHit = 0
            fmt.Printf("Cpu: %d%% MAX!\n", cp)
            hook("cpu")
        } else {
            cpHit++
            fmt.Printf("Cpu: %d%% %d hit!\n", cp, cpHit)
        }
    }

    if mp > maxMp {
        if mpHit >= maxHit {
            mpHit = 0
            fmt.Printf("Mem: %d%% MAX!\n", mp)
            hook("mem")
        } else {
            mpHit++
            fmt.Printf("Mem: %d%% %d hit!\n", mp, mpHit)
        }
    }

    if dp > maxDp {
        if dpHit >= maxHit {
            dpHit = 0
            fmt.Printf("Disk: %d%% MAX!\n", dp)
            hook("disk")
        } else {
            dpHit++
            fmt.Printf("Disk: %d%% %d hit!\n", dp, dpHit)
        }
    }
}

func hook(t string) {
    switch t {
    case "cpu":

    case "mem":

    case "disk":

    }
}

func config() {

}
