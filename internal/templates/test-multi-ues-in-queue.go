package templates

import (
	log "github.com/sirupsen/logrus"
	"my5G-RANTester/config"
	"my5G-RANTester/internal/control_test_engine/gnb"
	"my5G-RANTester/internal/control_test_engine/ue"
	"my5G-RANTester/internal/monitoring"
	"strconv"
	"sync"
	"time"
)

func TestMultiUesInQueue(numUes int) {

	wg := sync.WaitGroup{}

	cfg, err := config.GetConfig()
	if err != nil {
		//return nil
		log.Fatal("Error in get configuration")
	}

	go gnb.InitGnb(cfg, &wg)

	wg.Add(1)

	time.Sleep(1 * time.Second)

	for i := 1; i <= numUes; i++ {

		imsi := imsiGenerator(i)
		log.Info("[TESTER] TESTING REGISTRATION USING IMSI ", imsi, " UE")
		cfg.Ue.Msin = imsi
		go ue.RegistrationUe(cfg, uint8(i), &wg)
		wg.Add(1)

		time.Sleep(5 * time.Second)
	}

	wg.Wait()
}

// gera UE registration e mede a latência por segundos
func TestUesLatencyInInterval(interval int) int64 {

	wg := sync.WaitGroup{}

	monitor := monitoring.Monitor{
		LtRegisterGlobal: 0,
	}

	cfg, err := config.GetConfig()
	if err != nil {
		//return nil
		log.Fatal("Error in get configuration")
	}

	for i := 1; i <= interval; i++ {

		// sinal da gnb para manter a execução
		sigGnb := make(chan bool, 1)
		synch := make(chan bool, 1)

		// usado para sincronizar se a thread da gnb gerou um erro
		go gnb.InitGnb2(cfg, sigGnb, synch)

		// não houve erro na gnb
		if <-synch {

			time.Sleep(400 * time.Millisecond)

			log.Warn("[TESTER][UE] Test UE REGISTRATION:")
			go ue.RegistrationUeMonitor(cfg, uint8(i), &monitor, &wg)

			wg.Add(1)

			wg.Wait()

			// increment the latency global for the mean
			monitor.SetLtGlobal(monitor.LtRegisterLocal)
		} else {
			log.Warn("[TESTER][UE] UE LATENCY IN REGISTRATION: WITHOUT CONNECTION")
		}

		// ue registration per second
		time.Sleep(600 * time.Millisecond)

		// seta o sinal e termina a gnb
		sigGnb <- true

		time.Sleep(40 * time.Millisecond)
	}

	return monitor.LtRegisterGlobal
}

func imsiGenerator(i int) string {

	var base string
	switch true {
	case i < 10:
		base = "0000000"
	case i < 100:
		base = "000000"
	case i >= 100:
		base = "00000"
	}

	imsi := base + strconv.Itoa(i)
	return imsi
}
