package v201603

import (
	"encoding/xml"
)

const (
	DATE_RANGE_TODAY               = "TODAY"
	DATE_RANGE_YESTERDAY           = "YESTERDAY"
	DATE_RANGE_LAST_7_DAYS         = "LAST_7_DAYS"
	DATE_RANGE_LAST_WEEK           = "LAST_WEEK"
	DATE_RANGE_LAST_BUSINESS_WEEK  = "LAST_BUSINESS_WEEK"
	DATE_RANGE_THIS_MONTH          = "THIS_MONTH"
	DATE_RANGE_LAST_MONTH          = "LAST_MONTH"
	DATE_RANGE_ALL_TIME            = "ALL_TIME"
	DATE_RANGE_CUSTOM_DATE         = "CUSTOM_DATE"
	DATE_RANGE_LAST_14_DAYS        = "LAST_14_DAYS"
	DATE_RANGE_LAST_30_DAYS        = "LAST_30_DAYS"
	DATE_RANGE_THIS_WEEK_SUN_TODAY = "THIS_WEEK_SUN_TODAY"
	DATE_RANGE_THIS_WEEK_MON_TODAY = "THIS_WEEK_MON_TODAY"
	DATE_RANGE_LAST_WEEK_SUN_SAT   = "LAST_WEEK_SUN_SAT"
)

const (
	SEGMENT_DATE_DAY     = "Date"
	SEGMENT_DATE_WEEK    = "Week"
	SEGMENT_DATE_MONTH   = "Month"
	SEGMENT_DATE_QUARTER = "Quarter"
	SEGMENT_DATE_YEAR    = "Year"
)

const (
	DOWNLOAD_FORMAT_XML = "XML"
)

type CampaignReport struct {
	XMLName xml.Name     `xml:"report"`
	Rows    CampaignRows `xml:"table>row"`
}

type CampaignRow struct {
	XMLName                                                   xml.Name     `xml:"row"`
	AccountCurrencyCode                                       string       `xml:"currency,attr"`
	AccountDescriptiveName                                    string       `xml:"account,attr"`
	AccountTimeZoneId                                         string       `xml:"timeZone,attr"`
	ActiveViewCpm                                             string       `xml:"activeViewAvgCPM,attr"`
	ActiveViewCtr                                             string       `xml:"activeViewViewableCTR,attr"`
	ActiveViewImpressions                                     int64        `xml:"activeViewViewableImpressions,attr"`
	ActiveViewMeasurability                                   float64      `xml:"activeViewMeasurableImprImpr,attr"`
	ActiveViewMeasurableCost                                  int64        `xml:"activeViewMeasurableCost,attr"`
	ActiveViewMeasurableImpressions                           int64        `xml:"activeViewMeasurableImpr,attr"`
	ActiveViewViewability                                     float64      `xml:"activeViewViewableImprMeasurableImpr,attr"`
	AdNetworkType1                                            string       `xml:"network,attr"`
	AdNetworkType2                                            string       `xml:"networkWithSearchPartners,attr"`
	AdvertiserExperimentSegmentationBin                       string       `xml:"aCESplit,attr"`
	AdvertisingChannelSubType                                 string       `xml:"advertisingSubChannel,attr"`
	AdvertisingChannelType                                    string       `xml:"advertisingChannel,attr"`
	AllConversionRate                                         float64      `xml:"allConvRate,attr"`
	AllConversions                                            float64      `xml:"allConv,attr"`
	AllConversionValue                                        float64      `xml:"allConvValue,attr"`
	Amount                                                    string       `xml:"budget,attr"`
	AverageCost                                               string       `xml:"avgCost,attr"`
	AverageCpc                                                string       `xml:"avgCPC,attr"`
	AverageCpe                                                string       `xml:"avgCPE,attr"`
	AverageCpm                                                string       `xml:"avgCPM,attr"`
	AverageCpv                                                string       `xml:"avgCPV,attr"`
	AverageFrequency                                          int64        `xml:"avgImprFreqPerCookie,attr"`
	AveragePageviews                                          int64        `xml:"pagesSession,attr"`
	AveragePosition                                           int64        `xml:"avgPosition,attr"`
	AverageTimeOnSite                                         int64        `xml:"avgSessionDurationSeconds,attr"`
	BiddingStrategyId                                         int64        `xml:"bidStrategyID,attr"`
	BiddingStrategyName                                       string       `xml:"bidStrategyName,attr"`
	BiddingStrategyType                                       string       `xml:"bidStrategyType,attr"`
	BidType                                                   string       `xml:"conversionOptimizerBidType,attr"`
	BounceRate                                                float64      `xml:"bounceRate,attr"`
	BudgetId                                                  int64        `xml:"budgetID,attr"`
	CampaignId                                                string       `xml:"campaignID,attr"`
	CampaignName                                              string       `xml:"campaign,attr"`
	CampaignStatus                                            string       `xml:"campaignState,attr"`
	ClickAssistedConversions                                  int64        `xml:"clickAssistedConv,attr"`
	ClickAssistedConversionsOverLastClickConversions          float64      `xml:"clickAssistedConvLastClickConv,attr"`
	ClickAssistedConversionValue                              float64      `xml:"clickAssistedConvValue,attr"`
	ClickConversionRate                                       float64      `xml:"clickConversionRate,attr"`
	ClickConversionRateSignificance                           string       `xml:"clickConversionRateACEIndicator,attr"`
	Clicks                                                    int64        `xml:"clicks,attr"`
	ClickSignificance                                         string       `xml:"clicksACEIndicator,attr"`
	ClickType                                                 string       `xml:"clickType,attr"`
	ContentBudgetLostImpressionShare                          float64      `xml:"contentLostISBudget,attr"`
	ContentImpressionShare                                    float64      `xml:"contentImprShare,attr"`
	ContentRankLostImpressionShare                            float64      `xml:"contentLostISRank,attr"`
	ConversionCategoryName                                    string       `xml:"conversionCategory,attr"`
	ConversionRate                                            int64        `xml:"convRate,attr"`
	Conversions                                               int64        `xml:"conversions,attr"`
	ConversionTrackerId                                       int64        `xml:"conversionTrackerId,attr"`
	ConversionTypeName                                        string       `xml:"conversionName,attr"`
	ConversionValue                                           int64        `xml:"totalConvValue,attr"`
	ConvertedClicks                                           int64        `xml:"convertedClicks,attr"`
	ConvertedClicksSignificance                               string       `xml:"convertedClicksACEIndicator,attr"`
	Cost                                                      int64        `xml:"cost,attr"`
	CostPerAllConversion                                      float64      `xml:"costAllConv,attr"`
	CostPerConversion                                         float64      `xml:"costConv,attr"`
	CostPerConvertedClick                                     float64      `xml:"costConvertedClick,attr"`
	CostPerConvertedClickSignificance                         string       `xml:"costConvertedClickACEIndicator,attr"`
	CostSignificance                                          string       `xml:"costACEIndicator,attr"`
	CpcSignificance                                           string       `xml:"cPCACEIndicator,attr"`
	CpmSignificance                                           string       `xml:"cPMACEIndicator,attr"`
	CrossDeviceConversions                                    float64      `xml:"crossDeviceConv,attr"`
	Ctr                                                       float64      `xml:"ctr,attr"`
	CtrSignificance                                           string       `xml:"cTRACEIndicator,attr"`
	CustomerDescriptiveName                                   string       `xml:"clientName,attr"`
	Date                                                      string       `xml:"day,attr"`
	DayOfWeek                                                 string       `xml:"dayOfWeek,attr"`
	Device                                                    string       `xml:"device,attr"`
	EndDate                                                   string       `xml:"endDate,attr"`
	EngagementRate                                            float64      `xml:"engagementRate,attr"`
	Engagements                                               int64        `xml:"engagements,attr"`
	EnhancedCpcEnabled                                        string       `xml:"enhancedCPCEnabled,attr"`
	EnhancedCpvEnabled                                        string       `xml:"enhancedCPVEnabled,attr"`
	ExternalCustomerId                                        int64        `xml:"customerID,attr"`
	GmailForwards                                             int64        `xml:"gmailForwards,attr"`
	GmailSaves                                                int64        `xml:"gmailSaves,attr"`
	GmailSecondaryClicks                                      int64        `xml:"gmailClicksToWebsite,attr"`
	HourOfDay                                                 int64        `xml:"hourOfDay,attr"`
	ImpressionAssistedConversions                             int64        `xml:"imprAssistedConv,attr"`
	ImpressionAssistedConversionsOverLastClickConversions     float64      `xml:"imprAssistedConvLastClickConv,attr"`
	ImpressionAssistedConversionValue                         float64      `xml:"imprAssistedConvValue,attr"`
	ImpressionReach                                           int64        `xml:"uniqueCookies,attr"`
	Impressions                                               int64        `xml:"impressions,attr"`
	ImpressionSignificance                                    int64        `xml:"impressionsACEIndicator,attr"`
	InteractionRate                                           float64      `xml:"interactionRate,attr"`
	Interactions                                              int64        `xml:"interactions,attr"`
	InvalidClickRate                                          float64      `xml:"invalidClickRate,attr"`
	InvalidClicks                                             int64        `xml:"invalidClicks,attr"`
	IsBudgetExplicitlyShared                                  int64        `xml:"budgetExplicitlyShared,attr"`
	LabelIds                                                  string       `xml:"labelIDs,attr"`
	Labels                                                    string       `xml:"labels,attr"`
	Month                                                     string       `xml:"month,attr"`
	MonthOfYear                                               string       `xml:"monthOfYear,attr"`
	NumOfflineImpressions                                     int64        `xml:"phoneImpressions,attr"`
	NumOfflineInteractions                                    int64        `xml:"phoneCalls,attr"`
	OfflineInteractionRate                                    float64      `xml:"ptr,attr"`
	PercentNewVisitors                                        float64      `xml:"newSessions,attr"`
	Period                                                    string       `xml:"budgetPeriod,attr"`
	PositionSignificance                                      string       `xml:"positionACEIndicator,attr"`
	PrimaryCompanyName                                        string       `xml:"companyName,attr"`
	Quarter                                                   string       `xml:"quarter,attr"`
	RelativeCtr                                               float64      `xml:"relativeCTR,attr"`
	SearchBudgetLostImpressionShare                           float64      `xml:"searchLostISBudget,attr"`
	SearchExactMatchImpressionShare                           float64      `xml:"searchExactMatchIS,attr"`
	SearchImpressionShare                                     float64      `xml:"searchImprShare,attr"`
	SearchRankLostImpressionShare                             float64      `xml:"searchLostISRank,attr"`
	ServingStatus                                             string       `xml:"campaignServingStatus,attr"`
	Slot                                                      string       `xml:"topVsOther,attr"`
	StartDate                                                 string       `xml:"startDate,attr"`
	TrackingUrlTemplate                                       string       `xml:"trackingTemplate,attr"`
	UrlCustomParameters                                       string       `xml:"customParameter,attr"`
	ValuePerAllConversion                                     float64      `xml:"valueAllConv,attr"`
	ValuePerConversion                                        float64      `xml:"valueConv,attr"`
	ValuePerConvertedClick                                    int64        `xml:"valueConvertedClick,attr"`
	VideoQuartile25Rate                                       float64      `xml:"videoPlayedTo25,attr"`
	VideoQuartile50Rate                                       float64      `xml:"videoPlayedTo50,attr"`
	VideoQuartile75Rate                                       float64      `xml:"videoPlayedTo75,attr"`
	VideoQuartile100Rate                                      float64      `xml:"videoPlayedTo100,attr"`
	VideoViewRate                                             float64      `xml:"viewRate,attr"`
	VideoViews                                                int64        `xml:"views,attr"`
	ViewThroughConversions                                    int64        `xml:"viewThroughConv,attr"`
	ViewThroughConversionsSignificance                        string       `xml:"viewThroughConvACEIndicator,attr"`
	Week                                                      string       `xml:"week,attr"`
	Year                                                      int64        `xml:"year,attr"`
}

type CampaignRows []*CampaignRow

func (campaignRows *CampaignRows) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type Alias CampaignRow

	row := &Alias{}
	dec.DecodeElement(&row, &start)

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "week":
			fallthrough
		case "month":
			fallthrough
		case "day":
			fallthrough
		case "quarter":
			fallthrough
		case "year":
			row.Date = attr.Value
		}
	}

	*campaignRows = append(*campaignRows, (*CampaignRow)(row))
	return nil
}

type AdGroupReport struct {
	XMLName xml.Name     `xml:"report"`
	Rows    AdGroupRows `xml:"table>row"`
}

type AdGroupRow struct {
	XMLName                                                   xml.Name     `xml:"row"`
	AccountCurrencyCode                                       string       `xml:"currency,attr"`
	AccountDescriptiveName                                    string       `xml:"account,attr"`
	AccountTimeZoneId                                         string       `xml:"timeZone,attr"`
	ActiveViewCpm                                             string       `xml:"activeViewAvgCPM,attr"`
	ActiveViewCtr                                             string       `xml:"activeViewViewableCTR,attr"`
	ActiveViewImpressions                                     int64        `xml:"activeViewViewableImpressions,attr"`
	ActiveViewMeasurability                                   float64      `xml:"activeViewMeasurableImprImpr,attr"`
	ActiveViewMeasurableCost                                  int64        `xml:"activeViewMeasurableCost,attr"`
	ActiveViewMeasurableImpressions                           int64        `xml:"activeViewMeasurableImpr,attr"`
	ActiveViewViewability                                     float64      `xml:"activeViewViewableImprMeasurableImpr,attr"`
	AdGroupId                                                 int64        `xml:"adGroupID,attr"`
	AdGroupName                                               string       `xml:"adGroup,attr"`
	AdGroupStatus                                             string       `xml:"adGroupState,attr"`
	AdNetworkType1                                            string       `xml:"network,attr"`
	AdNetworkType2                                            string       `xml:"networkWithSearchPartners,attr"`
	AdvertiserExperimentSegmentationBin                       string       `xml:"aCESplit,attr"`
	AllConversionRate                                         float64      `xml:"allConvRate,attr"`
	AllConversions                                            float64      `xml:"allConv,attr"`
	AllConversionValue                                        float64      `xml:"allConvValue,attr"`
	AverageCost                                               string       `xml:"avgCost,attr"`
	AverageCpc                                                string       `xml:"avgCPC,attr"`
	AverageCpe                                                string       `xml:"avgCPE,attr"`
	AverageCpm                                                string       `xml:"avgCPM,attr"`
	AverageCpv                                                string       `xml:"avgCPV,attr"`
	AveragePageviews                                          int64        `xml:"pagesSession,attr"`
	AveragePosition                                           int64        `xml:"avgPosition,attr"`
	AverageTimeOnSite                                         int64        `xml:"avgSessionDurationSeconds,attr"`
	BiddingStrategyId                                         int64        `xml:"bidStrategyID,attr"`
	BiddingStrategyName                                       string       `xml:"bidStrategyName,attr"`
	BiddingStrategyType                                       string       `xml:"bidStrategyType,attr"`
	BidType                                                   string       `xml:"conversionOptimizerBidType,attr"`
	BounceRate                                                float64      `xml:"bounceRate,attr"`
	CampaignId                                                string       `xml:"campaignID,attr"`
	CampaignName                                              string       `xml:"campaign,attr"`
	CampaignStatus                                            string       `xml:"campaignState,attr"`
	ClickAssistedConversions                                  int64        `xml:"clickAssistedConv,attr"`
	ClickAssistedConversionsOverLastClickConversions          float64      `xml:"clickAssistedConvLastClickConv,attr"`
	ClickAssistedConversionValue                              float64      `xml:"clickAssistedConvValue,attr"`
	ClickConversionRate                                       float64      `xml:"clickConversionRate,attr"`
	ClickConversionRateSignificance                           string       `xml:"clickConversionRateACEIndicator,attr"`
	Clicks                                                    int64        `xml:"clicks,attr"`
	ClickSignificance                                         string       `xml:"clicksACEIndicator,attr"`
	ClickType                                                 string       `xml:"clickType,attr"`
	ContentBidCriterionTypeGroup                              string       `xml:"contentNetworkBidDimension,attr"`
	ContentImpressionShare                                    float64      `xml:"contentImprShare,attr"`
	ContentRankLostImpressionShare                            float64      `xml:"contentLostISRank,attr"`
	ConversionCategoryName                                    string       `xml:"conversionCategory,attr"`
	ConversionRate                                            int64        `xml:"convRate,attr"`
	Conversions                                               int64        `xml:"conversions,attr"`
	ConversionTrackerId                                       int64        `xml:"conversionTrackerId,attr"`
	ConversionTypeName                                        string       `xml:"conversionName,attr"`
	ConversionValue                                           int64        `xml:"totalConvValue,attr"`
	ConvertedClicks                                           int64        `xml:"convertedClicks,attr"`
	ConvertedClicksSignificance                               string       `xml:"convertedClicksACEIndicator,attr"`
	Cost                                                      int64        `xml:"cost,attr"`
	CostPerAllConversion                                      float64      `xml:"costAllConv,attr"`
	CostPerConversion                                         float64      `xml:"costConv,attr"`
	CostPerConvertedClick                                     float64      `xml:"costConvertedClick,attr"`
	CostPerConvertedClickSignificance                         string       `xml:"costConvertedClickACEIndicator,attr"`
	CostSignificance                                          string       `xml:"costACEIndicator,attr"`
	CpcBid                                                    float64      `xml:"defaultMaxCPC,attr"`
	CpcSignificance                                           string       `xml:"cPCACEIndicator,attr"`
	CpmBid                                                    float64      `xml:"maxCPM,attr"`
	CpmSignificance                                           string       `xml:"cPMACEIndicator,attr"`
	CpvBid                                                    float64      `xml:"maxCPV,attr"`
	CrossDeviceConversions                                    float64      `xml:"crossDeviceConv,attr"`
	Ctr                                                       float64      `xml:"ctr,attr"`
	CtrSignificance                                           string       `xml:"cTRACEIndicator,attr"`
	CustomerDescriptiveName                                   string       `xml:"clientName,attr"`
	Date                                                      string       `xml:"day,attr"`
	DayOfWeek                                                 string       `xml:"dayOfWeek,attr"`
	Device                                                    string       `xml:"device,attr"`
	EngagementRate                                            float64      `xml:"engagementRate,attr"`
	Engagements                                               int64        `xml:"engagements,attr"`
	EnhancedCpcEnabled                                        string       `xml:"enhancedCPCEnabled,attr"`
	EnhancedCpvEnabled                                        string       `xml:"enhancedCPVEnabled,attr"`
	ExternalCustomerId                                        int64        `xml:"customerID,attr"`
	GmailForwards                                             int64        `xml:"gmailForwards,attr"`
	GmailSaves                                                int64        `xml:"gmailSaves,attr"`
	GmailSecondaryClicks                                      int64        `xml:"gmailClicksToWebsite,attr"`
	HourOfDay                                                 int64        `xml:"hourOfDay,attr"`
	ImpressionAssistedConversions                             int64        `xml:"imprAssistedConv,attr"`
	ImpressionAssistedConversionsOverLastClickConversions     float64      `xml:"imprAssistedConvLastClickConv,attr"`
	ImpressionAssistedConversionValue                         float64      `xml:"imprAssistedConvValue,attr"`
	Impressions                                               int64        `xml:"impressions,attr"`
	ImpressionSignificance                                    int64        `xml:"impressionsACEIndicator,attr"`
	InteractionRate                                           float64      `xml:"interactionRate,attr"`
	Interactions                                              int64        `xml:"interactions,attr"`
	LabelIds                                                  string       `xml:"labelIDs,attr"`
	Labels                                                    string       `xml:"labels,attr"`
	Month                                                     string       `xml:"month,attr"`
	MonthOfYear                                               string       `xml:"monthOfYear,attr"`
	NumOfflineImpressions                                     int64        `xml:"phoneImpressions,attr"`
	NumOfflineInteractions                                    int64        `xml:"phoneCalls,attr"`
	OfflineInteractionRate                                    float64      `xml:"ptr,attr"`
	PercentNewVisitors                                        float64      `xml:"newSessions,attr"`
	PositionSignificance                                      string       `xml:"positionACEIndicator,attr"`
	PrimaryCompanyName                                        string       `xml:"companyName,attr"`
	Quarter                                                   string       `xml:"quarter,attr"`
	RelativeCtr                                               float64      `xml:"relativeCTR,attr"`
	SearchExactMatchImpressionShare                           float64      `xml:"searchExactMatchIS,attr"`
	SearchImpressionShare                                     float64      `xml:"searchImprShare,attr"`
	SearchRankLostImpressionShare                             float64      `xml:"searchLostISRank,attr"`
	Slot                                                      string       `xml:"topVsOther,attr"`
	TargetCpa                                                 float64      `xml:"maxCPAConvertedClicks,attr"`
	TrackingUrlTemplate                                       string       `xml:"trackingTemplate,attr"`
	UrlCustomParameters                                       string       `xml:"customParameter,attr"`
	ValuePerAllConversion                                     float64      `xml:"valueAllConv,attr"`
	ValuePerConversion                                        float64      `xml:"valueConv,attr"`
	ValuePerConvertedClick                                    int64        `xml:"valueConvertedClick,attr"`
	VideoQuartile100Rate                                      float64      `xml:"videoPlayedTo100,attr"`
	VideoQuartile25Rate                                       float64      `xml:"videoPlayedTo25,attr"`
	VideoQuartile50Rate                                       float64      `xml:"videoPlayedTo50,attr"`
	VideoQuartile75Rate                                       float64      `xml:"videoPlayedTo75,attr"`
	VideoViewRate                                             float64      `xml:"viewRate,attr"`
	VideoViews                                                int64        `xml:"views,attr"`
	ViewThroughConversions                                    int64        `xml:"viewThroughConv,attr"`
	ViewThroughConversionsSignificance                        string       `xml:"viewThroughConvACEIndicator,attr"`
	Week                                                      string       `xml:"week,attr"`
	Year                                                      int64        `xml:"year,attr"`
}

type AdPerformanceRows []*AdPerformanceRow

type AdPerformanceReport struct {
	XMLName xml.Name     `xml:"report"`
	Rows    AdGroupRows `xml:"table>row"`

}

type AdPerformanceRow struct {
	XMLName                    xml.Name     `xml:"row"`
	AccountCurrencyCode            string       `xml:"currency,attr"`
	AccountDescriptiveName         string       `xml:"account,attr"`
	AccountTimeZoneId              string       `xml:"timeZone,attr"`
	AdGroupId                      string       `xml:"adGroupID,attr"`
	AdGroupName                    string       `xml:"adGroup,attr"`
	AdGroupStatus                  string       `xml:"adGroupState,attr"`
	AdNetworkType1                 string       `xml:"network,attr"`
	AdNetworkType2                 string       `xml:"networkWithSearchPartners,attr"`
	AdType                         string       `xml:"adType,attr"`
	AllConversions                 float64      `xml:"allConv,attr"`
	AllConversionValue             float64      `xml:"allConvValue,attr"`
	CampaignId                     string       `xml:"campaignID,attr"`
	CampaignName                   string       `xml:"campaign,attr"`
	CampaignStatus                 string       `xml:"campaignState,attr"`
	BaseAdGroupId                  float64      `xml:"baseAdGroupID,attr"`
	BaseCampaignId                 string       `xml:"baseCampaignID,attr"`
	Clicks                         int64        `xml:"clicks,attr"`
	Conversions                    int64        `xml:"conversions,attr"`
	ConvertedClicks                int64        `xml:"convertedClicks,attr"`
	Cost                           int64        `xml:"cost,attr"`
	CreativeApprovalStatus         string       `xml:"adApprovalStatus,attr"`
	CreativeDestinationUrl         string       `xml:"destinationURL,attr"`
	Date                           string       `xml:"day,attr"`
	Description1                   string       `xml:"descriptionLine1,attr"`
	Description2                   string       `xml:"descriptionLine2,attr"`
	Device                         string       `xml:"device,attr"`
	DevicePreference               string       `xml:"devicePreference,attr"`
	DisplayUrl                     string       `xml:"displayURL,attr"`
	Engagements                    int64        `xml:"engagements,attr"`
	ExternalCustomerId             string       `xml:"customerID,attr"`
	GmailForwards                  int64        `xml:"gmailForwards,attr"`
	GmailSaves                     int64        `xml:"gmailSaves,attr"`
	GmailSecondaryClicks           int64        `xml:"gmailClicksToWebsite,attr"`
	Headline                       string       `xml:"ad,attr"`
	Id                             int64        `xml:"adID,attr"`
	ImageAdUrl                     string       `xml:"imageAdURL,attr"`
	ImageCreativeName              string       `xml:"imageAdName,attr"`
	Impressions                    int64        `xml:"impressions,attr"`
	Interactions                   int64        `xml:"interactions,attr"`
	IsNegative                     string       `xml:"isNegative,attr"`
	KeywordId                      string       `xml:"keywordID,attr"`
	LabelIds                       string       `xml:"labelIDs,attr"`
	Labels                         string       `xml:"labels,attr"`
	Status                         string       `xml:"adState,attr"`
	Trademarks                     string       `xml:"trademarks,attr"`
	VideoQuartile25Rate            float64      `xml:"videoPlayedTo25,attr"`
	VideoQuartile50Rate            float64      `xml:"videoPlayedTo50,attr"`
	VideoQuartile75Rate            float64      `xml:"videoPlayedTo75,attr"`
	VideoQuartile100Rate           float64      `xml:"videoPlayedTo100,attr"`
	VideoViews                     int64        `xml:"views,attr"`
}

type AdGroupRows []*AdGroupRow

func (adGroupRows *AdGroupRows) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type Alias AdGroupRow

	row := &Alias{}
	dec.DecodeElement(&row, &start)

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "week":
			fallthrough
		case "month":
			fallthrough
		case "day":
			fallthrough
		case "quarter":
			fallthrough
		case "year":
			row.Date = attr.Value
		}
	}

	*adGroupRows = append(*adGroupRows, (*AdGroupRow)(row))
	return nil
}

type BudgetReport struct {
	XMLName xml.Name     `xml:"report"`
	Rows    []*BudgetRow `xml:"table>row"`
}

type BudgetRow struct {
	XMLName     xml.Name `xml:"row"`
	AvgCPC      float64  `xml:"avgCPC,attr"`
	AvgCPM      float64  `xml:"avgCPM,attr"`
	CampaignID  int64    `xml:"campaignID,attr"`
	Clicks      int64    `xml:"clicks,attr"`
	Cost        float64  `xml:"cost,attr"`
	Impressions int64    `xml:"impressions,attr"`
	Conversions int64    `xml:"conversions,attr"`
}

type PlaceholderReport struct {
	XMLName xml.Name       `xml:"report"`
	Rows    PlaceholderRows  `xml:"table>row"`
}

type PlaceholderRow struct {
	XMLName                        xml.Name     `xml:"row"`
	AccountCurrencyCode            string       `xml:"currency,attr"`
	AccountDescriptiveName         string       `xml:"account,attr"`
	AccountTimeZoneId              string       `xml:"timeZone,attr"`
	AdGroupId                      string       `xml:"adGroupID,attr"`
	AdGroupName                    string       `xml:"adGroup,attr"`
	AdGroupStatus                  string       `xml:"adGroupState,attr"`
	AdNetworkType1                 string       `xml:"network,attr"`
	AdNetworkType2                 string       `xml:"networkWithSearchPartners,attr"`
	AllConversions                 float64      `xml:"allConv,attr"`
	AllConversionValue             float64      `xml:"allConvValue,attr"`
	CampaignId                     string       `xml:"campaignID,attr"`
	CampaignName                   string       `xml:"campaign,attr"`
	CampaignStatus                 string       `xml:"campaignState,attr"`
	Clicks                         int64        `xml:"clicks,attr"`
	Conversions                    int64        `xml:"conversions,attr"`
	ConversionValue                float64      `xml:"totalConvValue,attr"`
	Cost                           int64        `xml:"cost,attr"`
	Criteria                       string       `xml:"placement,attr"`
	CrossDeviceConversions         float64      `xml:"crossDeviceConv,attr"`
	CustomerDescriptiveName        string       `xml:"clientName,attr"`
	Date                           string       `xml:"day,attr"`
	Device                         string       `xml:"device,attr"`
	Engagements                    int64        `xml:"engagements,attr"`
	ExternalCustomerId             string       `xml:"customerID,attr"`
	Impressions                    int64        `xml:"impressions,attr"`
	Interactions                   int64        `xml:"interactions,attr"`
	VideoViews                     int64        `xml:"views,attr"`
}
type PlaceholderRows []*PlaceholderRow

type PlacementReport struct {
	XMLName xml.Name       `xml:"report"`
	Rows    PlacementRows  `xml:"table>row"`
}

type PlacementRow struct {
	XMLName                             xml.Name     `xml:"row"`
	AccountCurrencyCode                 string       `xml:"currency,attr"`
	AccountDescriptiveName              string       `xml:"account,attr"`
	ActiveViewImpressions               int64        `xml:"activeViewViewableImpressions,attr"`
	ActiveViewMeasurability             float64      `xml:"activeViewMeasurableImprImpr,attr"`
	ActiveViewMeasurableImpressions     int64        `xml:"activeViewMeasurableImpr,attr"`
	ActiveViewViewability               float64      `xml:"activeViewViewableImprMeasurableImpr,attr"`
	AdGroupId                           string       `xml:"adGroupID,attr"`
	AdGroupName                         string       `xml:"adGroup,attr"`
	AdGroupStatus                       string       `xml:"adGroupState,attr"`
	AdNetworkType1                      string       `xml:"network,attr"`
	AdNetworkType2                      string       `xml:"networkWithSearchPartners,attr"`
	AllConversions                      float64      `xml:"allConv,attr"`
	AllConversionValue                  float64      `xml:"allConvValue,attr"`
	BidModifier                         float64      `xml:"bidAdj,attr"`
	BidType                             string       `xml:"conversionOptimizerBidType,attr"`
	CampaignId                          string       `xml:"campaignID,attr"`
	CampaignName                        string       `xml:"campaign,attr"`
	CampaignStatus                      string       `xml:"campaignState,attr"`
	ClickConversionRate                 float64      `xml:"clickConversionRate,attr"`
	Clicks                              int64        `xml:"clicks,attr"`
	Cost                                int64        `xml:"cost,attr"`
	CpcBid                              float64      `xml:"defaultMaxCPC,attr"`
	CpcBidSource                        string       `xml:"maxCPCSource,attr"`
	CpmBid                              float64      `xml:"maxCPM,attr"`
	CpmBidSource                        string       `xml:"maxCPMSource,attr"`
	Criteria                            string       `xml:"placement,attr"`
	CriteriaDestinationUrl              string       `xml:"destinationURL,attr"`
	CrossDeviceConversions              float64      `xml:"crossDeviceConv,attr"`
	CustomerDescriptiveName             string       `xml:"clientName,attr"`
	Date                                string       `xml:"day,attr"`
	Device                              string       `xml:"device,attr"`
	DisplayName                         string       `xml:"criteriaDisplayName,attr"`
	Engagements                         int64        `xml:"engagements,attr"`
	ExternalCustomerId                  int64        `xml:"customerID,attr"`
	FinalAppUrls                        string       `xml:"appFinalURL,attr"`
	FinalMobileUrls                     string       `xml:"mobileFinalURL,attr"`
	FinalUrls                           string       `xml:"finalURL,attr"`
	GmailForwards                       int64        `xml:"gmailForwards,attr"`
	GmailSaves                          int64        `xml:"gmailSaves,attr"`
	GmailSecondaryClicks                int64        `xml:"gmailClicksToWebsite,attr"`
	Impressions                         int64        `xml:"impressions,attr"`
	InteractionRate                     float64      `xml:"interactionRate,attr"`
	Interactions                        int64        `xml:"interactions,attr"`
	IsNegative                          string       `xml:"isNegative,attr"`
	IsRestrict                          string       `xml:"isRestricting,attr"`
	PrimaryCompanyName                  string       `xml:"companyName,attr"`
	Status                              string       `xml:"placementState,attr"`
	VideoQuartile100Rate                float64      `xml:"videoPlayedTo100,attr"`
	VideoQuartile25Rate                 float64      `xml:"videoPlayedTo25,attr"`
	VideoQuartile50Rate                 float64      `xml:"videoPlayedTo50,attr"`
	VideoQuartile75Rate                 float64      `xml:"videoPlayedTo75,attr"`
	VideoViews                          int64        `xml:"views,attr"`
	ViewThroughConversions              int64        `xml:"viewThroughConv,attr"`
}

type PlacementRows []*PlacementRow

type VideoPerformanceReport struct {
	XMLName xml.Name       `xml:"report"`
	Rows    VideoPerformanceRows  `xml:"table>row"`
}

type VideoPerformanceRow struct {
	XMLName                  xml.Name   `xml:"row"`
	AdGroupId                string     `xml:"adGroupID,attr"`
	AdGroupName              string     `xml:"adGroup,attr"`
	AdGroupStatus            string     `xml:"adGroupState,attr"`
	AdNetworkType1           string     `xml:"network,attr"`
	AdNetworkType2           string     `xml:"networkWithSearchPartners,attr"`
	CampaignId               string     `xml:"campaignID,attr"`
	CampaignName             string     `xml:"campaign,attr"`
	CampaignStatus           string     `xml:"campaignState,attr"`
	Clicks                   int64      `xml:"clicks,attr"`
	Conversions              int64      `xml:"conversions,attr"`
	ConversionValue          float64    `xml:"totalConvValue,attr"`
	Cost                     int64      `xml:"cost,attr"`
	CreativeId               int64      `xml:"adID,attr"`
	CreativeStatus           string     `xml:"adState,attr"`
	Date                     string     `xml:"day,attr"`
	Device                   string     `xml:"device,attr"`
	Engagements              int64      `xml:"engagements,attr"`
	Impressions              int64      `xml:"impressions,attr"`
	VideoChannelId           string     `xml:"videoChannelId,attr"`
	VideoId                  string     `xml:"videoId,attr"`
	VideoQuartile100Rate     float64    `xml:"videoPlayedTo100,attr"`
	VideoQuartile25Rate      float64    `xml:"videoPlayedTo25,attr"`
	VideoQuartile50Rate      float64    `xml:"videoPlayedTo50,attr"`
	VideoQuartile75Rate      float64    `xml:"videoPlayedTo75,attr"`
	VideoTitle               string     `xml:"videoTitle,attr"`
	VideoViews               int64      `xml:"views,attr"`
	ViewThroughConversions   int64      `xml:"viewThroughConv,attr"`
}

type VideoPerformanceRows []*VideoPerformanceRow

type GeoPerformanceReport struct {
	XMLName xml.Name       `xml:"report"`
	Rows    GeoPerformanceRows  `xml:"table>row"`
}

type GeoPerformanceRow struct {
	XMLName                             xml.Name     `xml:"row"`
	AccountCurrencyCode                 string       `xml:"currency,attr"`
	AccountDescriptiveName              string       `xml:"account,attr"`
	AccountTimeZoneId                   string       `xml:"timeZone,attr"`
	AdFormat                            string       `xml:"adType,attr"`
	AdGroupId                           string       `xml:"adGroupID,attr"`
	AdGroupName                         string       `xml:"adGroup,attr"`
	AdGroupStatus                       string       `xml:"adGroupState,attr"`
	AdNetworkType1                      string       `xml:"network,attr"`
	AdNetworkType2                      string       `xml:"networkWithSearchPartners,attr"`
	AllConversions                      float64      `xml:"allConv,attr"`
	AllConversionValue                  float64      `xml:"allConvValue,attr"`
	AveragePosition                     int64        `xml:"avgPosition,attr"`
	CampaignId                          string       `xml:"campaignID,attr"`
	CampaignName                        string       `xml:"campaign,attr"`
	CampaignStatus                      string       `xml:"campaignState,attr"`
	CityCriteriaId                      string       `xml:"city,attr"`
	Clicks                              int64        `xml:"clicks,attr"`
	Conversions                         int64        `xml:"conversions,attr"`
	ConversionValue                     float64      `xml:"totalConvValue,attr"`
	ConvertedClicks                     int64        `xml:"convertedClicks,attr"`
	CountryCriteriaId                   string       `xml:"countryTerritory,attr"`
	CrossDeviceConversions              float64      `xml:"crossDeviceConv,attr"`
	CustomerDescriptiveName             string       `xml:"clientName,attr"`
	Date                                string       `xml:"day,attr"`
	Device                              string       `xml:"device,attr"`
	ExternalCustomerId                  string       `xml:"customerID,attr"`
	Impressions                         int64        `xml:"impressions,attr"`
	Interactions                        int64        `xml:"interactions,attr"`
	IsTargetingLocation                 string       `xml:"isTargetable,attr"`
	LocationType                        string       `xml:"locationType,attr"`
	MetroCriteriaId                     string       `xml:"metroArea,attr"`
	MostSpecificCriteriaId              float64      `xml:"mostSpecificLocation,attr"`
	PrimaryCompanyName                  string       `xml:"companyName,attr"`
	RegionCriteriaId                    string       `xml:"region,attr"`
	VideoViews                          int64        `xml:"views,attr"`
	ViewThroughConversions              int64        `xml:"viewThroughConv,attr"`
}

type GeoPerformanceRows []*GeoPerformanceRow

type KeywordsPerformanceReport struct {
	XMLName xml.Name       `xml:"report"`
	Rows    VideoPerformanceRows  `xml:"table>row"`
}

type KeywordsPerformanceRow struct {
	XMLName                  xml.Name   `xml:"row"`
	AccountCurrencyCode               string       `xml:"currency,attr"`
	AccountDescriptiveName            string       `xml:"account,attr"`
	AccountTimeZoneId                 string       `xml:"timeZone,attr"`
	ActiveViewImpressions             int64        `xml:"activeViewViewableImpressions,attr"`
	ActiveViewMeasurability           float64      `xml:"activeViewMeasurableImprImpr,attr"`
	ActiveViewMeasurableImpressions   int64        `xml:"activeViewMeasurableImpr,attr"`
	ActiveViewViewability             float64      `xml:"activeViewViewableImprMeasurableImpr,attr"`
	AdGroupId                         string       `xml:"adGroupID,attr"`
	AdGroupName                       string       `xml:"adGroup,attr"`
	AdGroupStatus                     string       `xml:"adGroupState,attr"`
	AdNetworkType1                    string       `xml:"network,attr"`
	AdNetworkType2                    string       `xml:"networkWithSearchPartners,attr"`
	AllConversions                    float64      `xml:"allConv,attr"`
	AllConversionValue                float64      `xml:"allConvValue,attr"`
	ApprovalStatus                    string       `xml:"approvalStatus,attr"`
	BaseAdGroupId                     float64      `xml:"baseAdGroupID,attr"`
	BaseCampaignId                    string       `xml:"baseCampaignID,attr"`
	BiddingStrategyId                 int64        `xml:"bidStrategyID,attr"`
	BiddingStrategyName               string       `xml:"bidStrategyName,attr"`
	BiddingStrategySource             string       `xml:"biddingStrategySource,attr"`
	BiddingStrategyType               string       `xml:"bidStrategyType,attr"`
	BidType                           float64      `xml:"conversionOptimizerBidType,attr"`
	CampaignId                        string       `xml:"campaignID,attr"`
	CampaignName                      string       `xml:"campaign,attr"`
	CampaignStatus                    string       `xml:"campaignState,attr"`
	Clicks                            int64        `xml:"clicks,attr"`
	Cost                              int64        `xml:"cost,attr"`
	Criteria                          string       `xml:"keyword,attr"`
	CriteriaDestinationUrl            string       `xml:"destinationURL,attr"`
	CustomerDescriptiveName           string       `xml:"clientName,attr"`
	Date                              string       `xml:"day,attr"`
	Engagements                       int64        `xml:"engagements,attr"`
	EnhancedCpcEnabled                string       `xml:"enhancedCPCEnabled,attr"`
	ExternalCustomerId                string       `xml:"customerID,attr"`
	FinalAppUrls                      string       `xml:"appFinalURL,attr"`
	FinalMobileUrls                   string       `xml:"mobileFinalURL,attr"`
	FinalUrls                         string       `xml:"finalURL,attr"`
	GmailForwards                     int64        `xml:"gmailForwards,attr"`
	GmailSaves                        int64        `xml:"gmailSaves,attr"`
	GmailSecondaryClicks              int64        `xml:"gmailClicksToWebsite,attr"`
	Id                                string       `xml:"keywordID,attr"`
	Impressions                       int64        `xml:"impressions,attr"`
	Interactions                      int64        `xml:"interactions,attr"`
	IsNegative                        string       `xml:"isNegative,attr"`
	LabelIds                          string       `xml:"labelIDs,attr"`
	Labels                            string       `xml:"labels,attr"`
	PrimaryCompanyName                string       `xml:"companyName,attr"`
	Status                            string       `xml:"keywordState,attr"`
	VideoQuartile100Rate              float64      `xml:"videoPlayedTo100,attr"`
	VideoQuartile25Rate               float64      `xml:"videoPlayedTo25,attr"`
	VideoQuartile50Rate               float64      `xml:"videoPlayedTo50,attr"`
	VideoQuartile75Rate               float64      `xml:"videoPlayedTo75,attr"`
	VideoViews                        int64        `xml:"views,attr"`
	ViewThroughConversions            int64        `xml:"viewThroughConv,attr"`
}

type KeywordsPerformanceRows []*KeywordsPerformanceRow

type AccountPerformanceReport struct {
	XMLName xml.Name       `xml:"report"`
	Rows    VideoPerformanceRows  `xml:"table>row"`
}

type AccountPerformanceRow struct {
	XMLName                        xml.Name `xml:"row"`
	AccountCurrencyCode            string   `xml:"currency,attr"`
	AccountDescriptiveName         string   `xml:"account,attr"`
	AccountTimeZoneId              string   `xml:"timeZone,attr"`
	AdNetworkType1                 string   `xml:"network,attr"`
	AdNetworkType2                 string   `xml:"networkWithSearchPartners,attr"`
	CanManageClients               string   `xml:"canManageClients,attr"`
	Clicks                         int64    `xml:"clicks,attr"`
	Conversions                    int64    `xml:"conversions,attr"`
	ConversionValue                float64  `xml:"totalConvValue,attr"`
	Cost                           int64    `xml:"cost,attr"`
	CustomerDescriptiveName        string   `xml:"clientName,attr"`
	Date                           string   `xml:"day,attr"`
	Device                         string   `xml:"device,attr"`
	Engagements                    int64    `xml:"engagements,attr"`
	ExternalCustomerId             string   `xml:"customerID,attr"`
	HourOfDay                      int64    `xml:"hourOfDay,attr"`
	Impressions                    int64    `xml:"impressions,attr"`
	Interactions                   int64    `xml:"interactions,attr"`
	IsTestAccount                  string   `xml:"testAccount,attr"`
	PrimaryCompanyName             string   `xml:"companyName,attr"`
	VideoViews                     int64    `xml:"views,attr"`
}

type AccountPerformanceRows []*AccountPerformanceRow

type ReportDefinition struct {
	XMLName                xml.Name `xml:"reportDefinition"`
	Id                     string   `xml:"id,omitempty"`
	Selector               Selector `xml:"selector"`
	ReportName             string   `xml:"reportName"`
	ReportType             string   `xml:"reportType"`
	HasAttachment          string   `xml:"hasAttachment,omitempty"`
	DateRangeType          string   `xml:"dateRangeType"`
	CreationTime           string   `xml:"creationTime,omitempty"`
	DownloadFormat         string   `xml:"downloadFormat"`
}

//Magic that sets downloadFormat automaticaly
func (c *ReportDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	c.DownloadFormat = DOWNLOAD_FORMAT_XML

	type Alias ReportDefinition

	start.Name = xml.Name{
		"", "reportDefinition",
	}

	e.EncodeElement((*Alias)(c), start)
	return nil
}

type ReportUtils struct {
	Auth
}

func NewReportUtils(auth *Auth) *ReportUtils {
	return &ReportUtils{Auth: *auth}
}

func (s *ReportUtils) DownloadCampaignPerformaceReport(reportDefinition *ReportDefinition) (report CampaignReport, err error) {
	reportDefinition.ReportType = "CAMPAIGN_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = CampaignReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadAdGroupPerformaceReport(reportDefinition *ReportDefinition) (report AdGroupReport, err error) {
	reportDefinition.ReportType = "ADGROUP_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = AdGroupReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadBudgetPerformanceReport(reportDefinition *ReportDefinition) (report BudgetReport, err error) {
	reportDefinition.ReportType = "BUDGET_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = BudgetReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadPlacementPerformaceReport(reportDefinition *ReportDefinition) (report PlacementReport, err error) {
	reportDefinition.ReportType = "PLACEMENT_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = PlacementReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadPlaceholderFeedReport(reportDefinition *ReportDefinition) (report PlaceholderReport, err error) {
	reportDefinition.ReportType = "PLACEHOLDER_FEED_ITEM_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = PlaceholderReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadGeoPerformanceReport(reportDefinition *ReportDefinition) (report GeoPerformanceReport, err error) {
	reportDefinition.ReportType = "GEO_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = GeoPerformanceReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadVideoPerformanceReport(reportDefinition *ReportDefinition) (report VideoPerformanceReport, err error) {
	reportDefinition.ReportType = "VIDEO_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = VideoPerformanceReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadKeywordsPerformanceReport(reportDefinition *ReportDefinition) (report KeywordsPerformanceReport, err error) {
	reportDefinition.ReportType = "KEYWORDS_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = KeywordsPerformanceReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadAdPerformanceReport(reportDefinition *ReportDefinition) (report AdPerformanceReport, err error) {
	reportDefinition.ReportType = "AD_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = AdPerformanceReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}

func (s *ReportUtils) DownloadAccountPerformanceReport(reportDefinition *ReportDefinition) (report AccountPerformanceReport, err error) {
	reportDefinition.ReportType = "ACCOUNT_PERFORMANCE_REPORT"

	respBody, err := s.Auth.downloadReportRequest(
		reportDefinition,
	)

	if err != nil {
		return report, err
	}

	report = AccountPerformanceReport{}
	err = xml.Unmarshal([]byte(respBody), &report)

	if err != nil {
		return report, err
	}

	return report, err
}
