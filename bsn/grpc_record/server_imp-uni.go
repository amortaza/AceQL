package grpc_record

/*
type MyServer struct {
	UnimplementedRecordServiceServer
}

func init() {
	stdsql.Init("mysql", "clown:1844@/bsn")
}

const LOG_SOURCE = "grpc_record"

func (MyServer) BabaSays(ctx context.Context, request *Request) (*Response, error) {
	//var record *flux.Record
	//defer func() {
	//	logger.Info("closing flux.Record", LOG_SOURCE)
	//	if record != nil {
	//		record.Close()
	//	}
	//}()
	//
	//fmt.Println("baba received - " + request.Operation + ", " + request.Param)
	//
	//// stdsql.NewRecord(tablename)
	//if request.Operation == "NewRecord()" {
	//	r, resp, err := _NewRecord(request)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	record = r
	//} else if request.Operation == "Query()" {
	//	err := _Query(request, record)
	//	if err != nil {
	//		return nil, err
	//	}
	//} else {
	//	if err := s.Send(&Response{Answer: "baba says " + request.Operation + ", " + request.Param}); err != nil {
	//		fmt.Println("Error " + err.Error())
	//		return nil
	//	}
	//}

	return &Response{Answer: "default ans", Fault: "default fault"}, nil
}

func _Query(request *Request, record *flux.Record) (*Response, error) {
	logger.Info("_Query() called", LOG_SOURCE)

	count, err := record.Query()
	if err != nil {
		return &Response{Answer: "no answer cuz error", Fault: err.Error()}, logger.Err(err, LOG_SOURCE)
	}

	return &Response{Answer: strconv.Itoa(count), Fault: ""}, nil
}

func _NewRecord(request *Request) (*flux.Record, *Response, error) {
	logger.Info("_NewRecord() called", LOG_SOURCE)

	record, err := stdsql.NewRecord(request.Param)
	if err != nil {
		return nil, &Response{Answer: "failed NewRecord", Fault: err.Error()}, logger.Err(err, LOG_SOURCE)
	}

	return record, &Response{Answer: "success NewRecord"}, nil
}

func (MyServer) mustEmbedUnimplementedRecordServiceServer() {}
*/
