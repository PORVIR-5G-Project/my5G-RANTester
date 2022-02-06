package gnb

import (
	log "github.com/sirupsen/logrus"
	"my5G-RANTester/config"
	"my5G-RANTester/internal/control_test_engine/gnb/context"
	serviceNas "my5G-RANTester/internal/control_test_engine/gnb/nas/service"
	serviceNgap "my5G-RANTester/internal/control_test_engine/gnb/ngap/service"
	"my5G-RANTester/internal/control_test_engine/gnb/ngap/trigger"
	"my5G-RANTester/internal/monitoring"
	"os"
	"os/signal"
	"sync"
	"time"
)

func InitGnb(conf config.Config, wg *sync.WaitGroup) {

	// instance new gnb.
	gnb := &context.GNBContext{}

	// new gnb context.
	gnb.NewRanGnbContext(
		conf.GNodeB.PlmnList.GnbId,
		conf.GNodeB.PlmnList.Mcc,
		conf.GNodeB.PlmnList.Mnc,
		conf.GNodeB.PlmnList.Tac,
		conf.GNodeB.SliceSupportList.Sst,
		conf.GNodeB.SliceSupportList.Sd,
		conf.GNodeB.ControlIF.Ip,
		conf.GNodeB.DataIF.Ip,
		conf.GNodeB.ControlIF.Port,
		conf.GNodeB.DataIF.Port)

	// start communication with AMF (server SCTP).

	// new AMF context.
	amf := gnb.NewGnBAmf(conf.AMF.Ip, conf.AMF.Port)

	// start communication with AMF(SCTP).
	if err := serviceNgap.InitConn(amf, gnb); err != nil {
		log.Fatal("Error in", err)
	} else {
		log.Info("[GNB] SCTP/NGAP service is running")
		// wg.Add(1)
	}

	// start communication with UE (server UNIX sockets).
	if err := serviceNas.InitServer(gnb); err != nil {
		log.Fatal("Error in", err)
	} else {
		log.Info("[GNB] UNIX/NAS service is running")
	}

	trigger.SendNgSetupRequest(gnb, amf)

	// control the signals
	sigGnb := make(chan os.Signal, 1)
	signal.Notify(sigGnb, os.Interrupt)

	// Block until a signal is received.
	<-sigGnb
	gnb.Terminate()
	wg.Done()
	// os.Exit(0)

}

func InitGnbForLoad(conf config.Config, wg *sync.WaitGroup,
	monitor *monitoring.Monitor, interval time.Duration) {

	// instance new gnb.
	gnb := &context.GNBContext{}

	// new gnb context.
	gnb.NewRanGnbContext(
		conf.GNodeB.PlmnList.GnbId,
		conf.GNodeB.PlmnList.Mcc,
		conf.GNodeB.PlmnList.Mnc,
		conf.GNodeB.PlmnList.Tac,
		conf.GNodeB.SliceSupportList.Sst,
		conf.GNodeB.SliceSupportList.Sd,
		conf.GNodeB.ControlIF.Ip,
		conf.GNodeB.DataIF.Ip,
		conf.GNodeB.ControlIF.Port,
		conf.GNodeB.DataIF.Port)

	// start communication with AMF (server SCTP).

	// new AMF context.
	amf := gnb.NewGnBAmf(conf.AMF.Ip, conf.AMF.Port)

	// start communication with AMF(SCTP).
	if err := serviceNgap.InitConn(amf, gnb); err != nil {
		log.Info("Error in", err)
		wg.Done()
		return
	} else {
		log.Info("[GNB] SCTP/NGAP service is running")
		// wg.Add(1)
	}

	trigger.SendNgSetupRequest(gnb, amf)

	time.Sleep(interval * time.Millisecond)

	if amf.GetState() == 0x01 {
		monitor.IncRqs()
	}

	gnb.Terminate()
	wg.Done()
	// os.Exit(0)
}

func InitGnbForLoadSeconds(conf config.Config, wg *sync.WaitGroup,
	monitor *monitoring.Monitor) {

	// instance new gnb.
	gnb := &context.GNBContext{}

	// new gnb context.
	gnb.NewRanGnbContext(
		conf.GNodeB.PlmnList.GnbId,
		conf.GNodeB.PlmnList.Mcc,
		conf.GNodeB.PlmnList.Mnc,
		conf.GNodeB.PlmnList.Tac,
		conf.GNodeB.SliceSupportList.Sst,
		conf.GNodeB.SliceSupportList.Sd,
		conf.GNodeB.ControlIF.Ip,
		conf.GNodeB.DataIF.Ip,
		conf.GNodeB.ControlIF.Port,
		conf.GNodeB.DataIF.Port)

	// start communication with AMF (server SCTP).

	// new AMF context.
	amf := gnb.NewGnBAmf(conf.AMF.Ip, conf.AMF.Port)

	// start communication with AMF(SCTP).
	if err := serviceNgap.InitConn(amf, gnb); err != nil {
		log.Info("Error in", err)
		time.Sleep(1000 * time.Millisecond)
		wg.Done()
		return
	} else {
		log.Info("[GNB] SCTP/NGAP service is running")
		// wg.Add(1)
	}

	trigger.SendNgSetupRequest(gnb, amf)

	time.Sleep(1000 * time.Millisecond)

	if amf.GetState() == 0x01 {
		monitor.IncRqs()
	}

	gnb.Terminate()
	wg.Done()
	// os.Exit(0)
}

// carga não fica suspensa por 1s
func InitGnbWithoutInterruption(conf config.Config, wg *sync.WaitGroup,
	monitor *monitoring.Monitor) {

	// instance new gnb.
	gnb := &context.GNBContext{}

	// new gnb context.
	gnb.NewRanGnbContext(
		conf.GNodeB.PlmnList.GnbId,
		conf.GNodeB.PlmnList.Mcc,
		conf.GNodeB.PlmnList.Mnc,
		conf.GNodeB.PlmnList.Tac,
		conf.GNodeB.SliceSupportList.Sst,
		conf.GNodeB.SliceSupportList.Sd,
		conf.GNodeB.ControlIF.Ip,
		conf.GNodeB.DataIF.Ip,
		conf.GNodeB.ControlIF.Port,
		conf.GNodeB.DataIF.Port)

	// start communication with AMF (server SCTP).

	// new AMF context.
	amf := gnb.NewGnBAmf(conf.AMF.Ip, conf.AMF.Port)

	// start communication with AMF(SCTP).
	if err := serviceNgap.InitConn(amf, gnb); err != nil {
		log.Info("Error in", err)
		wg.Done()
		return
	} else {
		log.Info("[GNB] SCTP/NGAP service is running")
		// wg.Add(1)
	}

	trigger.SendNgSetupRequest(gnb, amf)

	// handle time for response of packet
	start := time.Now()

	for {

		if amf.GetState() == 0x01 {
			monitor.InitRqsLocal()
			break
		}

		elapsed := time.Since(start)
		if (1000 - elapsed.Milliseconds()) <= 0 {
			break
		}
	}

	gnb.Terminate()
	wg.Done()
	// os.Exit(0)
}
