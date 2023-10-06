package structs

type Report struct {
	ReportId       string
	ImageUrls      []string
	Locations      string
	Offense        []string
	SubmittedByRTO bool
	Social         bool
	TotalFine      int
	RTOApproved    bool
	Comments       string
	ReportedBy     string
	VehicleNumber  string
}
