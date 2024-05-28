package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-railway-line-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-railway-line-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-railway-line-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-railway-line-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var partner *[]dpfm_api_output_formatter.Partner
	var expressType *[]dpfm_api_output_formatter.ExpressType
	var stopStation *[]dpfm_api_output_formatter.StopStation
	var departureStation *[]dpfm_api_output_formatter.DepartureStation
	var destinationStation *[]dpfm_api_output_formatter.DestinationStation
	var dptDstStation *[]dpfm_api_output_formatter.DptDstStation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Partner":
			partner = c.partnerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ExpressType":
			expressType = c.expressTypeCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "StopStation":
			stopStation = c.stopStationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "DepartureStation":
			departureStation = c.departureStationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "DestinationStation":
			destinationStation = c.destinationStationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "DptDstStation":
			dptDstStation = c.dptDstStationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)

		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Partner:            partner,
		ExpressType:		expressType,
		StopStation:		stopStation,
		DepartureStation:	departureStation,
		DestinationStation:	destinationStation,
		DptDstStation:		dptDstStation,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var partner *[]dpfm_api_output_formatter.Partner
	var expressType *[]dpfm_api_output_formatter.ExpressType
	var stopStation *[]dpfm_api_output_formatter.StopStation
	var departureStation *[]dpfm_api_output_formatter.DepartureStation
	var destinationStation *[]dpfm_api_output_formatter.DestinationStation
	var dptDstStation *[]dpfm_api_output_formatter.DptDstStation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Partner":
			partner = c.partnerUpdateSql(mtx, input, output, errs, log)
		case "ExpressType":
			expressType = c.expressTypeUpdateSql(mtx, input, output, errs, log)
		case "StopStation":
			stopStation = c.stopStationUpdateSql(mtx, input, output, errs, log)
		case "DepartureStation":
			departureStation = c.departureStationUpdateSql(mtx, input, output, errs, log)
		case "DestinationStation":
			destinationStation = c.destinationStationUpdateSql(mtx, input, output, errs, log)
		case "DptDstStation":
			dptDstStation = c.dptDstStationUpdateSql(mtx, input, output, errs, log)
			
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:             header,
		Partner:            partner,
		ExpressType:		expressType,
		StopStation:		stopStation,
		DepartureStation:	departureStation,
		DestinationStation:	destinationStation,
		DptDstStation:		dptDstStation,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "RailwayLineHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		*errs = append(*errs, err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, partnerData := range *subfuncSDC.Message.Partner {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "RailwayLinePartner", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Partner Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) expressTypeCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ExpressType {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, expressTypeData := range *subfuncSDC.Message.ExpressType {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": expressTypeData, "function": "RailwayLineExpressType", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ExpressType Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToExpressTypeCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) stopStationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StopStation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, stopStationData := range *subfuncSDC.Message.StopStation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": stopStationData, "function": "RailwayLineStopStation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "StopStation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStopStationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) departureStationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DepartureStation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, departureStationData := range *subfuncSDC.Message.DepartureStation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": departureStationData, "function": "RailwayLineDepartureStation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "DepartureStation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToDepartureStationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) destinationStationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DestinationStation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, destinationStationData := range *subfuncSDC.Message.DestinationStation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": destinationStationData, "function": "RailwayLineDestinationStation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "DestinationStation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToDestinationStationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) dptDstStationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DptDstStation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, dptDstStationData := range *subfuncSDC.Message.DptDstStation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": dptDstStationData, "function": "RailwayLineDptDstStation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "DptDstStation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToDptDstStationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "RailwayLineHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) partnerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Partner {
	req := make([]dpfm_api_processing_formatter.PartnerUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, partner := range header.Partner {
		partnerData := *dpfm_api_processing_formatter.ConvertToPartnerUpdates(header, partner)

		if partnerIsUpdate(&partnerData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": partnerData, "function": "RailwayLinePartner", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Partner Data cannot update"
				return nil
			}
		}
		req = append(req, partnerData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToPartnerUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) expressTypeUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ExpressType {
	req := make([]dpfm_api_processing_formatter.ExpressTypeUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, expressType := range header.ExpressType {
		expressTypeData := *dpfm_api_processing_formatter.ConvertToExpressTypeUpdates(header, expressType)

		if expressTypeIsUpdate(&expressTypeData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": expressTypeData, "function": "RailwayLineExpressType", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "ExpressType Data cannot update"
				return nil
			}
		}
		req = append(req, expressTypeData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToExpressTypeUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) stopStationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StopStation {
	req := make([]dpfm_api_processing_formatter.StopStationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, stopStation := range header.StopStation {
		stopStationData := *dpfm_api_processing_formatter.ConvertToStopStationUpdates(header, stopStation)

		if stopStationIsUpdate(&stopStationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": stopStationData, "function": "RailwayLineStopStation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "StopStation Data cannot update"
				return nil
			}
		}
		req = append(req, stopStationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStopStationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) departureStationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DepartureStation {
	req := make([]dpfm_api_processing_formatter.DepartureStationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, departureStation := range header.DepartureStation {
		departureStationData := *dpfm_api_processing_formatter.ConvertToDepartureStationUpdates(header, departureStation)

		if departureStationIsUpdate(&departureStationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": departureStationData, "function": "RailwayLineDepartureStation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "DepartureStation Data cannot update"
				return nil
			}
		}
		req = append(req, departureStationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToDepartureStationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) destinationStationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DestinationStation {
	req := make([]dpfm_api_processing_formatter.DestinationStationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, destinationStation := range header.DestinationStation {
		destinationStationData := *dpfm_api_processing_formatter.ConvertToDestinationStationUpdates(header, destinationStation)

		if destinationStationIsUpdate(&destinationStationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": destinationStationData, "function": "RailwayLineDestinationStation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "DestinationStation Data cannot update"
				return nil
			}
		}
		req = append(req, destinationStationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToDestinationStationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) dptDstSationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.DptDstSation {
	req := make([]dpfm_api_processing_formatter.DptDstSationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, dptDstSation := range header.DptDstSation {
		dptDstSationData := *dpfm_api_processing_formatter.ConvertToDptDstSationUpdates(header, dptDstSation)

		if dptDstSationIsUpdate(&dptDstSationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": dptDstSationData, "function": "RailwayLineDptDstSation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "DptDstSation Data cannot update"
				return nil
			}
		}
		req = append(req, dptDstSationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToDptDstSationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	railwayLine := header.RailwayLine

	return !(railwayLine == 0)
}

func partnerIsUpdate(partner *dpfm_api_processing_formatter.PartnerUpdates) bool {
	railwayLine := partner.RailwayLine
	partnerFunction := partner.PartnerFunction
	businessPartner := partner.BusinessPartner

	return !(railwayLine == 0 || partnerFunction == "" || businessPartner == 0)
}

func expressTypeIsUpdate(expressType *dpfm_api_processing_formatter.ExpressTypeUpdates) bool {
	railwayLine := expressType.RailwayLine
	expressType := expressType.ExpressType

	return !(railwayLine == 0 || expressType == "")
}

func stopStationIsUpdate(stopStation *dpfm_api_processing_formatter.StopStationUpdates) bool {
	railwayLine := stopStation.RailwayLine
	expressType := stopStation.ExpressType
	railwayLineStationID := stopStation.RailwayLineStationID
	stopStation := stopStation.StopStation

	return !(railwayLine == 0 || expressType == "" || railwayLineStationID == 0 || stopStation == 0)
}

func departureStationIsUpdate(departureStation *dpfm_api_processing_formatter.DepartureStationUpdates) bool {
	railwayLine := departureStation.RailwayLine
	expressType := departureStation.ExpressType
	departureStation := departureStation.DepartureStation

	return !(railwayLine == 0 || expressType == "" ||  departureStation == 0)
}

func destinationStationIsUpdate(destinationStation *dpfm_api_processing_formatter.DestinationStationUpdates) bool {
	railwayLine := destinationStation.RailwayLine
	expressType := destinationStation.ExpressType
	destinationStation := destinationStation.DestinationStation

	return !(railwayLine == 0 || expressType == "" ||  destinationStation == 0)
}

func dptDstStationIsUpdate(dptDstStation *dpfm_api_processing_formatter.DptDstStationUpdates) bool {
	railwayLine := dptDstStation.RailwayLine
	expressType := dptDstStation.ExpressType
	departureStation := dptDstStation.DepartureStation
	destinationStation := dptDstStation.DestinationStation

	return !(railwayLine == 0 || expressType == "" ||  departureStation == 0 ||  destinationStation == 0)
}
