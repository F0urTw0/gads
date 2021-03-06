package v201603

import (
	"testing"

	"github.com/kr/pretty"
)

func testReportUtils(t *testing.T) (service *ReportUtils) {
	return &ReportUtils{Auth: testAuthSetup(t)}
}

func TestDownloadBudgetPerformanceReportCustomDate(t *testing.T) {
	ru := testReportUtils(t)

	predicates := []Predicate{
		{"AssociatedCampaignId", "EQUALS", []string{"246257700"}},
	}

	reportDefinition := &ReportDefinition{
		Selector: Selector{
			Fields: []string{
				"AssociatedCampaignId",
				"AverageCpc",
				"AverageCpm",
				"Cost",
				"Clicks",
				"Impressions",
				"Conversions",
			},
			Predicates: predicates,
			DateRange: &DateRange{
				Min: "20150201", //YYYYMMDD
				Max: "20150601", //YYYYMMDD
			},
		},
		ReportName:             "Report #553f5265b3d84",
		DateRangeType:          DATE_RANGE_CUSTOM_DATE,
		IncludeZeroImpressions: true,
	}
	report, err := ru.DownloadBudgetPerformanceReport(reportDefinition)
	pretty.Println(report)
	pretty.Println(err)
}

func TestDownloadBudgetPerformanceReport(t *testing.T) {
	ru := testReportUtils(t)

	predicates := []Predicate{
		{"AssociatedCampaignId", "EQUALS", []string{"246257700"}},
	}

	reportDefinition := &ReportDefinition{
		Selector: Selector{
			Fields: []string{
				"AssociatedCampaignId",
				"AverageCpc",
				"AverageCpm",
				"Cost",
				"Clicks",
				"Impressions",
				"Conversions",
			},
			Predicates: predicates,
		},
		ReportName:             "Report #553f5265b3d84",
		DateRangeType:          DATE_RANGE_ALL_TIME,
		IncludeZeroImpressions: true,
	}
	report, err := ru.DownloadBudgetPerformanceReport(reportDefinition)
	pretty.Println(report)
	pretty.Println(err)
}

func TestDownloadCampaignPerformaceReport(t *testing.T) {
	ru := testReportUtils(t)

	predicates := []Predicate{
		{"CampaignId", "EQUALS", []string{"246257700"}},
	}

	reportDefinition := &ReportDefinition{
		Selector: Selector{
			Fields: []string{
				"CampaignId",
				"AverageCpc",
				"AverageCpm",
				"Cost",
				"Clicks",
				"Impressions",
				"Week", //Quarter, Month , Year, Week, Date
			},
			Predicates: predicates,
			DateRange: &DateRange{
				Min: "20150411", //YYYYMMDD
				Max: "20150621", //YYYYMMDD
			},
		},
		ReportName: "Report #553f5265b3d84",
		//		DateRangeType:          DATE_RANGE_ALL_TIME,
		DateRangeType:          DATE_RANGE_CUSTOM_DATE,
		IncludeZeroImpressions: true,
	}
	report, err := ru.DownloadCampaignPerformaceReport(reportDefinition)
	pretty.Println(report)
	pretty.Println(err)
}

func TestDownloadAdGroupPerformaceReport(t *testing.T) {
	ru := testReportUtils(t)

	predicates := []Predicate{
		{"AdGroupId", "EQUALS", []string{"246257700"}},
	}

	reportDefinition := &ReportDefinition{
		Selector: Selector{
			Fields: []string{
				"CampaignId",
				"AdGroup",
				"AverageCpc",
				"AverageCpm",
				"Cost",
				"Clicks",
				"Impressions",
				"Date",
			},
			Predicates: predicates,
			DateRange: &DateRange{
				Min: "20150411", //YYYYMMDD
				Max: "20150621", //YYYYMMDD
			},
		},
		ReportName: "Report #553f5265b3d84",
		//		DateRangeType:          DATE_RANGE_ALL_TIME,
		DateRangeType:          DATE_RANGE_CUSTOM_DATE,
		IncludeZeroImpressions: true,
	}
	report, err := ru.DownloadAdGroupPerformaceReport(reportDefinition)
	pretty.Println(report)
	pretty.Println(err)
}