package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-railway-line-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-railway-line-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-railway-line-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToExpressTypeCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ExpressType, error) {
	expressTypes := make([]ExpressType, 0)

	for _, data := range *subfuncSDC.Message.ExpressType {
		expressType, err := TypeConverter[*ExpressType](data)
		if err != nil {
			return nil, err
		}

		expressTypes = append(expressTypes, *expressType)
	}

	return &expressTypes, nil
}

func ConvertToStopStationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]StopStation, error) {
	stopStations := make([]StopStation, 0)

	for _, data := range *subfuncSDC.Message.StopStation {
		stopStation, err := TypeConverter[*StopStation](data)
		if err != nil {
			return nil, err
		}

		stopStations = append(stopStations, *stopStation)
	}

	return &stopStations, nil
}

func ConvertToDepartureStationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]DepartureStation, error) {
	departureStations := make([]DepartureStation, 0)

	for _, data := range *subfuncSDC.Message.DepartureStation {
		departureStation, err := TypeConverter[*DepartureStation](data)
		if err != nil {
			return nil, err
		}

		departureStations = append(departureStations, *departureStation)
	}

	return &departureStations, nil
}

func ConvertToDestinationStationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]DestinationStation, error) {
	destinationStations := make([]DestinationStation, 0)

	for _, data := range *subfuncSDC.Message.DestinationStation {
		destinationStation, err := TypeConverter[*DestinationStation](data)
		if err != nil {
			return nil, err
		}

		destinationStations = append(destinationStations, *destinationStation)
	}

	return &destinationStations, nil
}

func ConvertToDptDstStationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]DptDstStation, error) {
	dptDstStations := make([]DptDstStation, 0)

	for _, data := range *subfuncSDC.Message.DptDstStation {
		dptDstStation, err := TypeConverter[*DptDstStation](data)
		if err != nil {
			return nil, err
		}

		dptDstStations = append(dptDstStations, *dptDstStation)
	}

	return &dptDstStations, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) (*[]Partner, error) {
	partners := make([]Partner, 0)

	for _, data := range *partnerUpdates {
		partner, err := TypeConverter[*Partner](data)
		if err != nil {
			return nil, err
		}

		partners = append(partners, *partner)
	}

	return &partners, nil
}

func ConvertToExpressTypeUpdates(expressTypeUpdates *[]dpfm_api_processing_formatter.ExpressTypeUpdates) (*[]ExpressType, error) {
	expressTypes := make([]ExpressType, 0)

	for _, data := range *expressTypeUpdates {
		expressType, err := TypeConverter[*ExpressType](data)
		if err != nil {
			return nil, err
		}

		expressTypes = append(expressTypes, *expressType)
	}

	return &expressTypes, nil
}

func ConvertToStopStationUpdates(stopStationUpdates *[]dpfm_api_processing_formatter.StopStationUpdates) (*[]StopStation, error) {
	stopStations := make([]StopStation, 0)

	for _, data := range *stopStationUpdates {
		stopStation, err := TypeConverter[*StopStation](data)
		if err != nil {
			return nil, err
		}

		stopStations = append(stopStations, *stopStation)
	}

	return &stopStations, nil
}

func ConvertToDepartureStationUpdates(departureStationUpdates *[]dpfm_api_processing_formatter.DepartureStationUpdates) (*[]DepartureStation, error) {
	departureStations := make([]DepartureStation, 0)

	for _, data := range *departureStationUpdates {
		departureStation, err := TypeConverter[*DepartureStation](data)
		if err != nil {
			return nil, err
		}

		departureStations = append(departureStations, *departureStation)
	}

	return &departureStations, nil
}

func ConvertToDestinationStationUpdates(destinationStationUpdates *[]dpfm_api_processing_formatter.DestinationStationUpdates) (*[]DestinationStation, error) {
	destinationStations := make([]DestinationStation, 0)

	for _, data := range *destinationStationUpdates {
		destinationStation, err := TypeConverter[*DestinationStation](data)
		if err != nil {
			return nil, err
		}

		destinationStations = append(destinationStations, *destinationStation)
	}

	return &destinationStations, nil
}

func ConvertToDptDstStationUpdates(dptDstStationUpdates *[]dpfm_api_processing_formatter.DptDstStationUpdates) (*[]DptDstStation, error) {
	dptDstStations := make([]DptDstStation, 0)

	for _, data := range *dptDstStationUpdates {
		dptDstStation, err := TypeConverter[*DptDstStation](data)
		if err != nil {
			return nil, err
		}

		dptDstStations = append(dptDstStations, *dptDstStation)
	}

	return &dptDstStations, nil
}

func ConvertToHeader(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	subfuncSDC.Message.Header = &sub_func_complementer.Header{
		RailwayLine:                       		*input.Header.RailwayLine,
		RailwayLineType:						input.Header.RailwayLineType,
		RailwayLineOwner:						input.Header.RailwayLineOwner,
		RailwayLineOwnerBusinessPartnerRole:	input.Header.RailwayLineOwnerBusinessPartnerRole,
		Brand:									input.Header.Brand,
		PersonResponsible:						input.Header.PersonResponsible,
		URL:									input.Header.URL,
		ValidityStartDate:						input.Header.ValidityStartDate,
		ValidityEndDate:						input.Header.ValidityEndDate,
		DepartureStationOfOutboundLine:			input.Header.DepartureStationOfOutboundLine,
		DestinationStationOfOutboundLine:		input.Header.DestinationStationOfOutboundLine,
		Description:							input.Header.Description,
		LongText:								input.Header.LongText,
		Introduction:							input.Header.Introduction,
		RailwayLineSymbol:						input.Header.RailwayLineSymbol,
		Project:								input.Header.Project,
		WBSElement:								input.Header.WBSElement,
		Tag1:									input.Header.Tag1,
		Tag2:									input.Header.Tag2,
		Tag3:									input.Header.Tag3,
		Tag4:									input.Header.Tag4,
		CreationDate:							input.Header.CreationDate,
		CreationTime:							input.Header.CreationTime,
		LastChangeDate:							input.Header.LastChangeDate,
		LastChangeTime:							input.Header.LastChangeTime,
		CreateUser:								input.Header.CreateUser,
		LastChangeUser:							input.Header.LastChangeUser,
		IsReleased:								input.Header.IsReleased,
		IsMarkedForDeletion:					input.Header.IsMarkedForDeletion,
	}

	return subfuncSDC
}

func ConvertToPartner(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var partners []sub_func_complementer.Partner

	partners = append(
		partners,
		sub_func_complementer.Partner{
			RailwayLine:          	 *input.Header.RailwayLine,
			PartnerFunction:         input.Header.Partner[0].PartnerFunction,
			BusinessPartner:         input.Header.Partner[0].BusinessPartner,
			BusinessPartnerFullName: input.Header.Partner[0].BusinessPartnerFullName,
			BusinessPartnerName:     input.Header.Partner[0].BusinessPartnerName,
			Organization:            input.Header.Partner[0].Organization,
			Country:                 input.Header.Partner[0].Country,
			Language:                input.Header.Partner[0].Language,
			Currency:                input.Header.Partner[0].Currency,
			ExternalDocumentID:      input.Header.Partner[0].ExternalDocumentID,
			AddressID:               input.Header.Partner[0].AddressID,
			EmailAddress:            input.Header.Partner[0].EmailAddress,
		},
	)

	subfuncSDC.Message.Partner = &partners

	return subfuncSDC
}

func ConvertToExpressType(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var expressTypes []sub_func_complementer.ExpressType

	expressTypes = append(
		expressTypes,
		sub_func_complementer.ExpressType{
			RailwayLine:          	*input.Header.RailwayLine,
			ExpressType:			input.Header.ExpressType[0].ExpressType,
			ValidityStartDate:		input.Header.ExpressType[0].ValidityStartDate,
			ValidityEndDate:		input.Header.ExpressType[0].ValidityEndDate,
			Description:			input.Header.ExpressType[0].Description,
			LongText:				input.Header.ExpressType[0].LongText,
			Introduction:			input.Header.ExpressType[0].Introduction,
			Project:				input.Header.ExpressType[0].Project,
			WBSElement:				input.Header.ExpressType[0].WBSElement,
			Tag1:					input.Header.ExpressType[0].Tag1,
			Tag2:					input.Header.ExpressType[0].Tag2,
			Tag3:					input.Header.ExpressType[0].Tag3,
			Tag4:					input.Header.ExpressType[0].Tag4,
			CreationDate:			input.Header.ExpressType[0].CreationDate,
			CreationTime:			input.Header.ExpressType[0].CreationTime,
			LastChangeDate:			input.Header.ExpressType[0].LastChangeDate,
			LastChangeTime:			input.Header.ExpressType[0].LastChangeTime,
			CreateUser:				input.Header.ExpressType[0].CreateUser,
			LastChangeUser:			input.Header.ExpressType[0].LastChangeUser,
			IsReleased:				input.Header.ExpressType[0].IsReleased,
			IsMarkedForDeletion:	input.Header.ExpressType[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.ExpressType = &expressTypes

	return subfuncSDC
}

func ConvertToStopStation(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var stopStations []sub_func_complementer.StopStation

	stopStations = append(
		stopStations,
		sub_func_complementer.StopStation{
			RailwayLine:          			*input.Header.RailwayLine,
			ExpressType:					input.Header.StopStation[0].ExpressType,
			RailwayLineStationID:			input.Header.StopStation[0].RailwayLineStationID,
			StopStation:					input.Header.StopStation[0].StopStation,
			OrderNumberOnOutboundLine:		input.Header.StopStation[0].OrderNumberOnOutboundLine,
			OrderNumberOnInboundLine:		input.Header.StopStation[0].OrderNumberOnInboundLine,
			PlatformNumberForOutboundLine:	input.Header.StopStation[0].PlatformNumberForOutboundLine,
			PlatformNumberForInboundLine:	input.Header.StopStation[0].PlatformNumberForInboundLine,
			ValidityStartDate:				input.Header.StopStation[0].ValidityStartDate,
			ValidityEndDate:				input.Header.StopStation[0].ValidityEndDate,
			Description:					input.Header.StopStation[0].Description,
			LongText:						input.Header.StopStation[0].LongText,
			Introduction:					input.Header.StopStation[0].Introduction,
			RailwayLineStopStationSymbol:	input.Header.StopStation[0].RailwayLineStopStationSymbol,
			Project:						input.Header.StopStation[0].Project,
			WBSElement:						input.Header.StopStation[0].WBSElement,
			Tag1:							input.Header.StopStation[0].Tag1,
			Tag2:							input.Header.StopStation[0].Tag2,
			Tag3:							input.Header.StopStation[0].Tag3,
			Tag4:							input.Header.StopStation[0].Tag4,
			CreationDate:					input.Header.StopStation[0].CreationDate,
			CreationTime:					input.Header.StopStation[0].CreationTime,
			LastChangeDate:					input.Header.StopStation[0].LastChangeDate,
			LastChangeTime:					input.Header.StopStation[0].LastChangeTime,
			CreateUser:						input.Header.StopStation[0].CreateUser,
			LastChangeUser:					input.Header.StopStation[0].LastChangeUser,
			IsReleased:						input.Header.StopStation[0].IsReleased,
			IsMarkedForDeletion:			input.Header.StopStation[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.StopStation = &stopStations

	return subfuncSDC
}

func ConvertToDepartureStation(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var departureStations []sub_func_complementer.DepartureStation

	departureStations = append(
		departureStations,
		sub_func_complementer.DepartureStation{
			RailwayLine:          	*input.Header.RailwayLine,
			ExpressType:			input.Header.DepartureStation[0].ExpressType,
			DepartureStation:		input.Header.DepartureStation[0].DepartureStation,
			RailwayLineStationID:	input.Header.DepartureStation[0].RailwayLineStationID,
			ValidityStartDate:		input.Header.DepartureStation[0].ValidityStartDate,
			ValidityEndDate:		input.Header.DepartureStation[0].ValidityEndDate,
			CreationDate:			input.Header.DepartureStation[0].CreationDate,
			CreationTime:			input.Header.DepartureStation[0].CreationTime,
			LastChangeDate:			input.Header.DepartureStation[0].LastChangeDate,
			LastChangeTime:			input.Header.DepartureStation[0].LastChangeTime,
			CreateUser:				input.Header.DepartureStation[0].CreateUser,
			LastChangeUser:			input.Header.DepartureStation[0].LastChangeUser,
			IsReleased:				input.Header.DepartureStation[0].IsReleased,
			IsMarkedForDeletion:	input.Header.DepartureStation[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.DepartureStation = &departureStations

	return subfuncSDC
}

func ConvertToDestinationStation(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var destinationStations []sub_func_complementer.DestinationStation

	destinationStations = append(
		destinationStations,
		sub_func_complementer.DestinationStation{
			RailwayLine:          	*input.Header.RailwayLine,
			ExpressType:			input.Header.DestinationStation[0].ExpressType,
			DestinationStation:		input.Header.DestinationStation[0].DestinationStation,
			RailwayLineStationID:	input.Header.DestinationStation[0].RailwayLineStationID,
			ValidityStartDate:		input.Header.DestinationStation[0].ValidityStartDate,
			ValidityEndDate:		input.Header.DestinationStation[0].ValidityEndDate,
			CreationDate:			input.Header.DestinationStation[0].CreationDate,
			CreationTime:			input.Header.DestinationStation[0].CreationTime,
			LastChangeDate:			input.Header.DestinationStation[0].LastChangeDate,
			LastChangeTime:			input.Header.DestinationStation[0].LastChangeTime,
			CreateUser:				input.Header.DestinationStation[0].CreateUser,
			LastChangeUser:			input.Header.DestinationStation[0].LastChangeUser,
			IsReleased:				input.Header.DestinationStation[0].IsReleased,
			IsMarkedForDeletion:	input.Header.DestinationStation[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.DestinationStation = &destinationStations

	return subfuncSDC
}

func ConvertToDptDstStation(
	input *dpfm_api_input_reader.SDC,
	subfuncSDC *sub_func_complementer.SDC,
) *sub_func_complementer.SDC {
	var dptDstStations []sub_func_complementer.DptDstStation

	dptDstStations = append(
		dptDstStations,
		sub_func_complementer.DptDstStation{
			RailwayLine:          	*input.Header.RailwayLine,
			ExpressType:			input.Header.DptDstStation[0].ExpressType,
			DepartureStation:		input.Header.DptDstStation[0].DepartureStation,
			DestinationStation:		input.Header.DptDstStation[0].DestinationStation,
			ValidityStartDate:		input.Header.DptDstStation[0].ValidityStartDate,
			ValidityEndDate:		input.Header.DptDstStation[0].ValidityEndDate,
			CreationDate:			input.Header.DptDstStation[0].CreationDate,
			CreationTime:			input.Header.DptDstStation[0].CreationTime,
			LastChangeDate:			input.Header.DptDstStation[0].LastChangeDate,
			LastChangeTime:			input.Header.DptDstStation[0].LastChangeTime,
			CreateUser:				input.Header.DptDstStation[0].CreateUser,
			LastChangeUser:			input.Header.DptDstStation[0].LastChangeUser,
			IsReleased:				input.Header.DptDstStation[0].IsReleased,
			IsMarkedForDeletion:	input.Header.DptDstStation[0].IsMarkedForDeletion,
		},
	)

	subfuncSDC.Message.DptDstStation = &dptDstStations

	return subfuncSDC
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
