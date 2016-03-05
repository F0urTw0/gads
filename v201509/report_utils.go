package v201509

import "encoding/xml"

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
	EngagementRate                                            float64      `xml:"endDate,attr"`
	Engagements                                               int64        `xml:"engagements,attr"`
	EnhancedCpcEnabled                                        string       `xml:"enhancedCPCEnabled,attr"`
	EnhancedCpvEnabled                                        string       `xml:"enhancedCPVEnabled,attr"`
	ExternalCustomerId                                        int64        `xml:"engagements,attr"`
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
	IncludeZeroImpressions bool     `xml:"includeZeroImpressions"`
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

	//	respBody := `<?xml version='1.0' encoding='UTF-8' standalone='yes'?>
	//<report>
	//	<report-name name="Report #553f5265b3d84"/>
	//	<date-range date="Apr 11, 2015-Jun 21, 2015"/>
	//	<table>
	//		<columns>
	//			<column display="Campaign ID" name="campaignID"/>
	//			<column display="Avg. CPC" name="avgCPC"/>
	//			<column display="Avg. CPM" name="avgCPM"/>
	//			<column display="Cost" name="cost"/>
	//			<column display="Clicks" name="clicks"/>
	//			<column display="Impressions" name="impressions"/>
	//			<column display="Week" name="week"/>
	//		</columns>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-06"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-20"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-27"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-04"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-11"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-18"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-25"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-06-15"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-13"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-06-01"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-06-08"/>
	//	</table>
	//</report>`

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

	//	respBody := `<?xml version='1.0' encoding='UTF-8' standalone='yes'?>
	//<report>
	//	<report-name name="Report #553f5265b3d84"/>
	//	<date-range date="Apr 11, 2015-Jun 21, 2015"/>
	//	<table>
	//		<columns>
	//			<column display="Campaign ID" name="campaignID"/>
	//			<column display="Avg. CPC" name="avgCPC"/>
	//			<column display="Avg. CPM" name="avgCPM"/>
	//			<column display="Cost" name="cost"/>
	//			<column display="Clicks" name="clicks"/>
	//			<column display="Impressions" name="impressions"/>
	//			<column display="Week" name="week"/>
	//		</columns>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-06"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-20"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-27"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-04"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-11"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-18"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-05-25"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-06-15"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-04-13"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-06-01"/>
	//		<row avgCPC="0" avgCPM="0" campaignID="246257700" clicks="0" cost="0" impressions="0" week="2015-06-08"/>
	//	</table>
	//</report>`

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

	//	respBody := `<?xml version='1.0' encoding='UTF-8' standalone='yes'?>
	//	<report>
	//		<report-name name="Report #553f5265b3d84"/>
	//		<date-range date="All Time"/>
	//		<table>
	//			<columns>
	//				<column display="Campaign ID" name="campaignID"/>
	//				<column display="Avg. CPC" name="avgCPC"/>
	//				<column display="Avg. CPM" name="avgCPM"/>
	//				<column display="Cost" name="cost"/>
	//				<column display="Clicks" name="clicks"/>
	//				<column display="Impressions" name="impressions"/>
	//			</columns>
	//			<row campaignID='1111' avgCPC='1' avgCPM='2' cost='3' clicks='5' impressions='7' convertedClicks='3'/>
	//			<row campaignID='1111' avgCPC='2' avgCPM='2' cost='3' clicks='5' impressions='7' convertedClicks='3'/>
	//			<row campaignID='1111' avgCPC='3' avgCPM='2' cost='3' clicks='5' impressions='7' convertedClicks='3'/>
	//		</table>
	//	</report>`

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
