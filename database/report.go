package database

import (
	"log"
	"rto/structs"
)

func SubmitReport(report structs.Report) {
	var inputArgs []interface{}
	var reportId string
	query := `insert into report.public_report (report_id, image_urls, "location", offense, is_submitted_by_RTO, social, total_fine,
	"comment", reported_by, vehicle_number, submitted_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, current_timestamp)`
	inputArgs = append(inputArgs, report.ReportId)
	inputArgs = append(inputArgs, report.ImageUrls)
	inputArgs = append(inputArgs, report.Locations)
	inputArgs = append(inputArgs, report.Offense)
	inputArgs = append(inputArgs, report.SubmittedByRTO)
	inputArgs = append(inputArgs, report.Social)
	inputArgs = append(inputArgs, report.TotalFine)
	inputArgs = append(inputArgs, report.Comments)
	inputArgs = append(inputArgs, report.ReportedBy)
	inputArgs = append(inputArgs, report.VehicleNumber)

	err2 := DB.QueryRow(query, inputArgs...).Scan(&reportId)
	if err2 != nil {
		log.Println("Error from DB")
	}

}
