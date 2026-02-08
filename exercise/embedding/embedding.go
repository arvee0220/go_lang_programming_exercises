//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

type SystemDashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

/* func Calculate(comp *SystemDashboard) (int,int,int) {
	var bUResult Bytes = 0
	var tempResult Celcius = 0
	var mUResult Bytes = 0
	

	for _, v := range comp.BandwidthUsage.amount {
		bUResult += v

	}

	for _, v := range comp.CpuTemp.temp {
		tempResult +=v
	}

	for _, v := range comp.MemoryUsage.amount {
		mUResult += v
	}

	var bWidthAvg int = int(bUResult) / len(comp.BandwidthUsage.amount)
	var tempAvg int = int(tempResult) / len(comp.CpuTemp.temp)
	var mUAvg int = int(mUResult) / len(comp.MemoryUsage.amount)


	return bWidthAvg, tempAvg, mUAvg
} */


func (bw *BandwidthUsage) AverageBandwidth() Bytes {
	var bUResult Bytes

	for _, v := range bw.amount {
		bUResult += v
	}

	avg := bUResult / Bytes(len(bw.amount))

	return avg
}

func (tmp *CpuTemp) AverageCpuTemp() Celcius {
	var tempResult Celcius

	for _, v := range tmp.temp {
		tempResult += v
	}

	avg := tempResult / Celcius(len(tmp.temp))

	return avg
}

func (mU *MemoryUsage) AverageMemoryUsage() Bytes {
	var mUResult Bytes

	for _, v := range mU.amount {
		mUResult += v
	}

	avg := mUResult / Bytes(len(mU.amount))

	return avg
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}

	dashboard := SystemDashboard{ BandwidthUsage: bandwidth, CpuTemp: temp, MemoryUsage: memory}

	fmt.Printf("Average Bandwidth: %v\nAverage CPU Temp: %v\nAverage Memory Usage: %v", dashboard.AverageBandwidth(), dashboard.AverageCpuTemp(), dashboard.AverageMemoryUsage())
}
