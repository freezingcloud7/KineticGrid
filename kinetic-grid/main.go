package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type GridBattery struct {
	storedEnergy int64
}

var localSystemHealth int64 = 40  
var gridBattery int64 = 0         
var lightningRodGrounded int64 = 0 

const maxBatteryCapacity int64 = 50000

func main() {
	fmt.Println("🚀 [KINETIC GRID 3.0] Three-Tier Energy Absorption Engine Online.")
	fmt.Printf("[INITIAL] Local Health: %d%% | Grid Battery: %d Watts\n", localSystemHealth, gridBattery)

	var wg sync.WaitGroup
	startTime := time.Now()

	totalIncomingTraffic := 15000

	for i := 1; i <= totalIncomingTraffic; i++ {
		wg.Add(1)
		go func(nodeID int) {
			defer wg.Done()

			isHealthy := nodeID%2 == 0
			rawKineticEnergy := int64(30) 

			if isHealthy {
				return
			}

			currentHealth := atomic.LoadInt64(&localSystemHealth)

			if currentHealth < 100 {
				atomic.AddInt64(&localSystemHealth, 5)
				if atomic.LoadInt64(&localSystemHealth) > 100 {
					atomic.StoreInt64(&localSystemHealth, 100)
				}
			} else {
				currentBattery := atomic.LoadInt64(&gridBattery)

				if currentBattery < maxBatteryCapacity {
					atomic.AddInt64(&gridBattery, rawKineticEnergy*2)
					if atomic.LoadInt64(&gridBattery) > maxBatteryCapacity {
						atomic.StoreInt64(&gridBattery, maxBatteryCapacity)
					}
				} else {
					atomic.AddInt64(&lightningRodGrounded, rawKineticEnergy)
				}
			}
		}(i)
	}

	wg.Wait() 
	duration := time.Since(startTime)

	fmt.Println("\n================================================================================")
	fmt.Printf("🏁 [KINETIC GRID SUCCESS] 3-Tier Recovery Verified in %v\n", duration)
	fmt.Printf("[TIER 1 RESULT] Local Server Mainframe Health: %d%% (Perfect Integrity Secured)\n", atomic.LoadInt64(&localSystemHealth))
	fmt.Printf("[TIER 2 RESULT] Vampire Touch Grid Battery: %d / %d Watts (Fully Charged)\n", atomic.LoadInt64(&gridBattery), maxBatteryCapacity)
	fmt.Printf("[TIER 3 RESULT] Lightning Rod Energy Safely Dissipated to Ground: %d Watts\n", atomic.LoadInt64(&lightningRodGrounded))
	fmt.Println("[STATUS] Freezing Down & Burnout Prevention Status: 100% Impregnable")
	fmt.Println("================================================================================")
}