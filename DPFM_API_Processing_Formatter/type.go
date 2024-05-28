package dpfm_api_processing_formatter

type HeaderUpdates struct {
	RailwayLine							int		`json:"RailwayLine"`
	RailwayLineType						string	`json:"RailwayLineType"`
	RailwayLineOwner					int		`json:"RailwayLineOwner"`
	RailwayLineOwnerBusinessPartnerRole	string	`json:"RailwayLineOwnerBusinessPartnerRole"`
	Brand								*int	`json:"Brand"`
	PersonResponsible					*string	`json:"PersonResponsible"`
	URL									*string	`json:"URL"`
	ValidityStartDate					string	`json:"ValidityStartDate"`
	ValidityEndDate						string	`json:"ValidityEndDate"`
	DepartureStationOfOutboundLine		int		`json:"DepartureStationOfOutboundLine"`
	DestinationStationOfOutboundLine	int		`json:"DestinationStationOfOutboundLine"`
	Description							string	`json:"Description"`
	LongText							string	`json:"LongText"`
	Introduction						*string	`json:"Introduction"`
	RailwayLineSymbol					*string	`json:"RailwayLineSymbol"`
	Project								*int	`json:"Project"`
	WBSElement							*int	`json:"WBSElement"`
	Tag1								*string	`json:"Tag1"`
	Tag2								*string	`json:"Tag2"`
	Tag3								*string	`json:"Tag3"`
	Tag4								*string	`json:"Tag4"`
	LastChangeDate						string	`json:"LastChangeDate"`
	LastChangeTime						string	`json:"LastChangeTime"`
	LastChangeUser						int		`json:"LastChangeUser"`
}

type PartnerUpdates struct {
	RailwayLine             int     `json:"RailwayLine"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
	EmailAddress            *string `json:"EmailAddress"`
}

type ExpressTypeUpdates struct {
	RailwayLine			int		`json:"RailwayLine"`
	ExpressType			string	`json:"ExpressType"`
	ValidityStartDate	string	`json:"ValidityStartDate"`
	ValidityEndDate		string	`json:"ValidityEndDate"`
	Description			string	`json:"Description"`
	LongText			string	`json:"LongText"`
	Introduction		*string	`json:"Introduction"`
	Project				*int	`json:"Project"`
	WBSElement			*int	`json:"WBSElement"`
	Tag1				*string	`json:"Tag1"`
	Tag2				*string	`json:"Tag2"`
	Tag3				*string	`json:"Tag3"`
	Tag4				*string	`json:"Tag4"`
	LastChangeDate		string	`json:"LastChangeDate"`
	LastChangeTime		string	`json:"LastChangeTime"`
	LastChangeUser		int		`json:"LastChangeUser"`
}

type StopStationUpdates struct {
	RailwayLine						int		`json:"RailwayLine"`
	ExpressType						string	`json:"ExpressType"`
	RailwayLineStationID			int		`json:"RailwayLineStationID"`
	StopStation						int		`json:"StopStation"`
	OrderNumberOnOutboundLine		int		`json:"OrderNumberOnOutboundLine"`
	OrderNumberOnInboundLine		int		`json:"OrderNumberOnInboundLine"`
	PlatformNumberForOutboundLine	string	`json:"PlatformNumberForOutboundLine"`
	PlatformNumberForInboundLine	string	`json:"PlatformNumberForInboundLine"`
	ValidityStartDate				string	`json:"ValidityStartDate"`
	ValidityEndDate					string	`json:"ValidityEndDate"`
	Description						string	`json:"Description"`
	LongText						string	`json:"LongText"`
	Introduction					*string	`json:"Introduction"`
	RailwayLineStopStationSymbol	*string	`json:"RailwayLineStopStationSymbol"`
	Project							*int	`json:"Project"`
	WBSElement						*int	`json:"WBSElement"`
	Tag1							*string	`json:"Tag1"`
	Tag2							*string	`json:"Tag2"`
	Tag3							*string	`json:"Tag3"`
	Tag4							*string	`json:"Tag4"`
	LastChangeDate					string	`json:"LastChangeDate"`
	LastChangeTime					string	`json:"LastChangeTime"`
	LastChangeUser					int		`json:"LastChangeUser"`
}

type DepartureStationUpdates struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	DepartureStation			int		`json:"DepartureStation"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	LastChangeUser				int		`json:"LastChangeUser"`
}

type DestinationStationUpdates struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	DestinationStation			int		`json:"DestinationStation"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	LastChangeUser				int		`json:"LastChangeUser"`
}

type DptDstStationUpdates struct {
	RailwayLine					int		`json:"RailwayLine"`
	ExpressType					string	`json:"ExpressType"`
	DepartureStation			int		`json:"DepartureStation"`
	DestinationStation			int		`json:"DestinationStation"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	LastChangeUser				int		`json:"LastChangeUser"`
}
