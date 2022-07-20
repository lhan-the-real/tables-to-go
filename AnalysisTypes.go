package dto

import (
	"database/sql"
	"time"
)

type AnalysisTypes struct {
	ID                      int            `db:"id"`
	CreatedAt               time.Time      `db:"created_at"`
	UpdatedAt               time.Time      `db:"updated_at"`
	SegmentID               int            `db:"segment_id"`
	AlertTypeID             int            `db:"alert_type_id"`
	Code                    string         `db:"code"`
	Display                 int            `db:"display"`
	Active                  int            `db:"active"`
	AnalyseLive             int            `db:"analyse_live"`
	ProgressWeight          int            `db:"progress_weight"`
	QualityWeight           float64        `db:"quality_weight"`
	ReferenceTable          sql.NullString `db:"reference_table"`
	ReferenceColumn         sql.NullString `db:"reference_column"`
	ReferenceSegmentFieldID sql.NullInt64  `db:"reference_segment_field_id"`
	Name                    sql.NullString `db:"name"`
	Description             sql.NullString `db:"description"`
	DescriptionDynamic      sql.NullString `db:"description_dynamic"`
	XMLFormatValidation     int            `db:"xml_format_validation"`
	DisplayReceiver         int            `db:"display_receiver"`
	URLSupport              string         `db:"url_support"`
}
