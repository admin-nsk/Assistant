package shutdown

import (
	"io"
	"note_service/app/pkg/logging"
	"os"
	"os/signal"
)

func Graceful(signals []os.Signal, closeItems ...io.Closer) {
	logger := logging.Getlogger()

	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, signals...)
	sig := <-signalC
	logger.Infof("Cought signal %s. Shutting down....", sig)
	for _, closer := range closeItems {
		if err := closer.Close(); err != nil {
			logger.Errorf("filed to close %v: %v", closer, err)
		}

	}

}
