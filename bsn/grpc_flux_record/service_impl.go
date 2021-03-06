package grpc_flux_record

import (
	"fmt"
	"github.com/amortaza/aceql/flux"
	"github.com/amortaza/aceql/flux-drivers/stdsql"
	"github.com/amortaza/aceql/flux/query"
	"github.com/amortaza/aceql/logger"
	"strconv"
)

type FluxRecordServiceImp struct {
	UnimplementedFluxRecordServiceServer
}

func init() {
	stdsql.Init("mysql", "clown:1844@/bsn")
}

const LOG_SOURCE = "grpc_record"

func (FluxRecordServiceImp) GetServiceStream(s FluxRecordService_GetServiceStreamServer) error {
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
			return logger.Err(err, LOG_SOURCE)
		}

		logger.Info("baba received - "+request.Operation+", "+request.Param1, LOG_SOURCE)

		// stdsql.NewRecord(tablename)
		if request.Operation == "Open()" {
			r, err := _Open(s, request)
			if err != nil {
				return err
			}

			record = r
		} else if request.Operation == "Query()" {
			if err := _Query(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Next()" {
			if err := _Next(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Initialize()" {
			if err := _Initialize(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "GetTable()" {
			if err := _GetTable(s, request, record); err != nil {
				return err
			}
			//} else if request.Operation == "GetMap()" {
			//	if err := _GetMap(s, request, record); err != nil {
			//		return err
			//	}
		} else if request.Operation == "SetOrderByDesc()" {
			if err := _SetOrderByDesc(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "SetOrderBy()" {
			if err := _SetOrderBy(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "GetFieldType()" {
			if err := _GetFieldType(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Set()" {
			if err := _Set(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Close()" {
			if err := _Close(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Insert()" {
			if err := _Insert(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Update()" {
			if err := _Update(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Delete()" {
			if err := _Delete(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Pagination()" {
			if err := _Pagination(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Get()" {
			if err := _Get(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "AddPK()" {
			if err := _AddPK(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "SetEncodedQuery()" {
			if err := _SetEncodedQuery(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Add()" {
			if err := _Add(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "AddEq()" {
			if err := _AddEq(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "AddOr()" {
			if err := _AddOr(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "AndGroup()" {
			if err := _AndGroup(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "OrGroup()" {
			if err := _OrGroup(s, request, record); err != nil {
				return err
			}
		} else if request.Operation == "Not()" {
			if err := _Not(s, request, record); err != nil {
				return err
			}
		} else {
			if err := s.Send(&Response{Fault: "Unrecognized Operation: " + request.Operation}); err != nil {
				return logger.Error(fmt.Sprintf("error %s", err.Error()), "?")
			}
		}
	}
}

// We distinguish between "send_err" and "flux_err"
// We report flux_err via Response.Fault, we do not return error because that will kill FluxRecord session.
// We do however send back "send_err" becasude this means the session is dead anyways.
func _Not(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Not() called", LOG_SOURCE)

	record.Not()

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _OrGroup(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_OrGroup() called", LOG_SOURCE)

	flux_err := record.OrGroup()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _AndGroup(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_AndGroup() called", LOG_SOURCE)

	flux_err := record.AndGroup()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _AddOr(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_AddOr() called", LOG_SOURCE)

	field := request.Param1
	opTypeName := request.Param2
	rhs := request.Param3

	opType, flux_err := query.GetOpTypeByName(opTypeName)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	flux_err = record.AddOr(field, opType, rhs)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _AddEq(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_AddEq() called", LOG_SOURCE)

	field := request.Param1
	rhs := request.Param2

	flux_err := record.AddEq(field, rhs)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _Add(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Add() called", LOG_SOURCE)

	field := request.Param1
	opTypeName := request.Param2
	rhs := request.Param3

	opType, flux_err := query.GetOpTypeByName(opTypeName)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	flux_err = record.Add(field, opType, rhs)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _SetEncodedQuery(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_SetEncodedQuery() called", LOG_SOURCE)

	encodedQuery := request.Param1

	record.SetEncodedQuery(encodedQuery)

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _AddPK(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_AddPK() called", LOG_SOURCE)

	id := request.Param1

	flux_err := record.AddPK(id)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Get(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Get() called", LOG_SOURCE)

	fieldname := request.Param1

	value, flux_err := record.Get(fieldname)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{Answer: value}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _Pagination(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Pagination() called", LOG_SOURCE)

	indexParam := request.Param1
	sizeParam := request.Param2

	var flux_err error

	if indexParam == "" || sizeParam == "" {
		flux_err = logger.Error(fmt.Sprintf("GRPC Pagination(index,size), but got \"%s\" \"%s\"", sizeParam, indexParam), LOG_SOURCE)
	}

	var index, size int

	if flux_err == nil {
		index, flux_err = strconv.Atoi(indexParam)

		if flux_err == nil {
			size, flux_err = strconv.Atoi(sizeParam)
		}
	}

	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	record.Pagination(index, size)

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _Delete(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Delete() called", LOG_SOURCE)

	flux_err := record.Delete()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Update(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Update() called", LOG_SOURCE)

	flux_err := record.Update()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Insert(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Insert() called", LOG_SOURCE)

	id, flux_err := record.Insert()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{Answer: id}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Close(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Close() called", LOG_SOURCE)

	flux_err := record.Close()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _GetFieldType(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_GetFieldType() called", LOG_SOURCE)

	fieldname := request.Param1

	fieldType, flux_err := record.GetFieldType(fieldname)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{Answer: string(fieldType)}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Set(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Set() called", LOG_SOURCE)

	fieldname := request.Param1
	value := request.Param2

	flux_err := record.Set(fieldname, value)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _SetOrderBy(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_SetOrderBy() called", LOG_SOURCE)

	fields := request.Param1

	record.SetOrderBy(fields)

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _SetOrderByDesc(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_SetOrderByDesc() called", LOG_SOURCE)

	fields := request.Param1

	if flux_err := record.SetOrderByDesc(fields); flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _GetTable(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_GetTable() called", LOG_SOURCE)

	table := record.GetTable()

	if send_err := s.Send(&Response{Answer: table}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

func _Initialize(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Initialize() called", LOG_SOURCE)

	record.Initialize()

	if send_err := s.Send(&Response{}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Next(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Next() called", LOG_SOURCE)

	has, flux_err := record.Next()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{Answer: strconv.FormatBool(has)}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Query(s FluxRecordService_GetServiceStreamServer, request *Request, record *flux.Record) error {
	logger.Info("_Query() called", LOG_SOURCE)

	count, flux_err := record.Query()
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return logger.Err(send_err, LOG_SOURCE)
		}

		return nil
	}

	if send_err := s.Send(&Response{Answer: strconv.Itoa(count)}); send_err != nil {
		return logger.Err(send_err, LOG_SOURCE)
	}

	return nil
}

// in client_flux_record
func _Open(s FluxRecordService_GetServiceStreamServer, request *Request) (*flux.Record, error) {
	logger.Info("_NewRecord() called", LOG_SOURCE)

	record, flux_err := stdsql.NewRecord(request.Param1)
	if flux_err != nil {
		if send_err := s.Send(&Response{Fault: flux_err.Error()}); send_err != nil {
			return nil, send_err
		}

		return nil, nil
	}

	if send_err := s.Send(&Response{}); send_err != nil {
		return nil, logger.Err(send_err, LOG_SOURCE)
	}

	return record, nil
}

func (FluxRecordServiceImp) mustEmbedUnimplementedRecordServiceServer() {}
