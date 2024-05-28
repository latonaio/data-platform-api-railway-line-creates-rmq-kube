package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-railway-line-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		RailwayLine:							data.RailwayLine,
		RailwayLineType:						data.RailwayLineType,
		RailwayLineOwner:						data.RailwayLineOwner,
		RailwayLineOwnerBusinessPartnerRole:	data.RailwayLineOwnerBusinessPartnerRole,
		Brand:									data.Brand,
		PersonResponsible:						data.PersonResponsible,
		URL:									data.URL,
		ValidityStartDate:						data.ValidityStartDate,
		ValidityEndDate:						data.ValidityEndDate,
		DepartureStationOfOutboundLine:			data.DepartureStationOfOutboundLine,
		DestinationStationOfOutboundLine:		data.DestinationStationOfOutboundLine,
		Description:							data.Description,
		LongText:								data.LongText,
		Introduction:							data.Introduction,
		RailwayLineSymbol:						data.RailwayLineSymbol,
		Project:								data.Project,
		WBSElement:								data.WBSElement,
		Tag1:									data.Tag1,
		Tag2:									data.Tag2,
		Tag3:									data.Tag3,
		Tag4:									data.Tag4,
		LastChangeDate:							data.LastChangeDate,
		LastChangeTime:							data.LastChangeTime,
		LastChangeUser:							data.LastChangeUser,
	}
}

func ConvertToPartnerUpdates(header dpfm_api_input_reader.Header, partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	dataHeader := header
	data := partner

	return &PartnerUpdates{
		RailwayLine:             dataHeader.RailwayLine,
		PartnerFunction:         data.PartnerFunction,
		BusinessPartner:         data.BusinessPartner,
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
		AddressID:               data.AddressID,
		EmailAddress:            data.EmailAddress,
	}
}

func ConvertToExpressTypeUpdates(expressType dpfm_api_input_reader.ExpressType) *ExpressTypeUpdates {
	data := expressType

	return &ExpressTypeUpdates{
		RailwayLine:            dataHeader.RailwayLine,
		ExpressType:			data.ExpressType,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityEndDate:		data.ValidityEndDate,
		Description:			data.Description,
		LongText:				data.LongText,
		Introduction:			data.Introduction,
		Project:				data.Project,
		WBSElement:				data.WBSElement,
		Tag1:					data.Tag1,
		Tag2:					data.Tag2,
		Tag3:					data.Tag3,
		Tag4:					data.Tag4,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
		LastChangeUser:			data.LastChangeUser,
	}
}

func ConvertToStopStationUpdates(stopStation dpfm_api_input_reader.StopStation) *StopStationUpdates {
	data := stopStation

	return &StopStationUpdates{
		RailwayLine:					dataHeader.RailwayLine,
		ExpressType:					data.ExpressType,
		RailwayLineStationID:			data.RailwayLineStationID,
		StopStation:					data.StopStation,
		OrderNumberOnOutboundLine:		data.OrderNumberOnOutboundLine,
		OrderNumberOnInboundLine:		data.OrderNumberOnInboundLine,
		PlatformNumberForOutboundLine:	data.PlatformNumberForOutboundLine,
		PlatformNumberForInboundLine:	data.PlatformNumberForInboundLine,
		ValidityStartDate:				data.ValidityStartDate,
		ValidityEndDate:				data.ValidityEndDate,
		Description:					data.Description,
		LongText:						data.LongText,
		Introduction:					data.Introduction,
		RailwayLineStopStationSymbol:	data.RailwayLineStopStationSymbol,
		Project:						data.Project,
		WBSElement:						data.WBSElement,
		Tag1:							data.Tag1,
		Tag2:							data.Tag2,
		Tag3:							data.Tag3,
		Tag4:							data.Tag4,
		LastChangeDate:					data.LastChangeDate,
		LastChangeTime:					data.LastChangeTime,
		LastChangeUser:					data.LastChangeUser,
	}
}

func ConvertToDepartureStationUpdates(departureStation dpfm_api_input_reader.DepartureStation) *DepartureStationUpdates {
	data := departureStation

	return &DepartureStationUpdates{
		RailwayLine:			dataHeader.RailwayLine,
		ExpressType:			data.ExpressType,
		DepartureStation:		data.DepartureStation,
		RailwayLineStationID:	data.RailwayLineStationID,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityEndDate:		data.ValidityEndDate,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
		LastChangeUser:			data.LastChangeUser,
	}
}

func ConvertToDestinationStationUpdates(destinationStation dpfm_api_input_reader.DestinationStation) *DestinationStationUpdates {
	data := destinationStation

	return &DestinationStationUpdates{
		RailwayLine:			dataHeader.RailwayLine,
		ExpressType:			data.ExpressType,
		DestinationStation:		data.DestinationStation,
		RailwayLineStationID:	data.RailwayLineStationID,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityEndDate:		data.ValidityEndDate,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
		LastChangeUser:			data.LastChangeUser,
	}
}

func ConvertToDptDstStationUpdates(dptDstStation dpfm_api_input_reader.DptDstStation) *DptDstStationUpdates {
	data := dptDstStation

	return &DptDstStationUpdates{
		RailwayLine:			dataHeader.RailwayLine,
		ExpressType:			data.ExpressType,
		DepartureStation:		data.DepartureStation,
		DestinationStation:		data.DestinationStation,
		ValidityStartDate:		data.ValidityStartDate,
		ValidityEndDate:		data.ValidityEndDate,
		LastChangeDate:			data.LastChangeDate,
		LastChangeTime:			data.LastChangeTime,
		LastChangeUser:			data.LastChangeUser,
	}
}
