package grpc_record

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/logger"
	"strconv"
)

type MyServer struct {
	UnimplementedRecordServiceServer
}

func init() {
	stdsql.Init("mysql", "clown:1844@/bsn")
}

const LOG_SOURCE = "grpc_record"

func (MyServer) BabaSays(s RecordService_BabaSaysServer) error {
	var record *flux.Record
	defer func() {
		logger.Info("closing flux.Record", LOG_SOURCE)
		if record != nil {
			record.Close()
		}
	}()

	for {
		request, err := s.Recv()
		if err != nil {
			logger.Err(err, LOG_SOURCE)
			return err
		}

		fmt.Println("baba received - " + request.Operation + ", " + request.Param)

		// stdsql.NewRecord(tablename)
		if request.Operation == "NewRecord()" {
			r, err := _NewRecord(s, request)
			if err != nil {
				return err
			}

			record = r
		} else if request.Operation == "Query()" {
			err := _Query(s, request, record)
			if err != nil {
				return err
			}
		} else {
			if err := s.Send(&Response{Answer: "baba says " + request.Operation + ", " + request.Param}); err != nil {
				fmt.Println("Error " + err.Error())
				return nil
			}
		}
	}

	return nil
}

func _Query(s RecordService_BabaSaysServer, request *Request, record *flux.Record) error {
	logger.Info("_Query() called", LOG_SOURCE)

	count, err := record.Query()
	if err != nil {
		s.Send(&Response{Answer: "failed Query", Fault: err.Error()})
		return logger.Err(err, LOG_SOURCE)
	}

	err = s.Send(&Response{Answer: strconv.Itoa(count)})
	if err != nil {
		return logger.Err(err, LOG_SOURCE)
	}

	return nil
}

func _NewRecord(s RecordService_BabaSaysServer, request *Request) (*flux.Record, error) {
	logger.Info("_NewRecord() called", LOG_SOURCE)

	record, err := stdsql.NewRecord(request.Param)
	if err != nil {
		s.Send(&Response{Answer: "failed NewRecord", Fault: err.Error()})
		return nil, logger.Err(err, LOG_SOURCE)
	}

	err = s.Send(&Response{Answer: "success NewRecord"})
	if err != nil {
		return nil, logger.Err(err, LOG_SOURCE)
	}

	return record, nil
}

func (MyServer) mustEmbedUnimplementedRecordServiceServer() {}
