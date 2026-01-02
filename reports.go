package trading212

import (
	"context"
	"net/http"
	"time"
)

// ReportDataIncluded represents what data to include in a report
type ReportDataIncluded struct {
	IncludeDividends    bool `json:"includeDividends"`
	IncludeInterest     bool `json:"includeInterest"`
	IncludeOrders       bool `json:"includeOrders"`
	IncludeTransactions bool `json:"includeTransactions"`
}

// PublicReportRequest represents a report request
type PublicReportRequest struct {
	DataIncluded ReportDataIncluded `json:"dataIncluded"`
	TimeFrom     time.Time          `json:"timeFrom"`
	TimeTo       time.Time          `json:"timeTo"`
}

// EnqueuedReportResponse represents a report request response
type EnqueuedReportResponse struct {
	ReportID int64 `json:"reportId"`
}

// ReportResponse represents a report status response
type ReportResponse struct {
	DataIncluded ReportDataIncluded `json:"dataIncluded"`
	DownloadLink *string            `json:"downloadLink"`
	ReportID     int64              `json:"reportId"`
	Status       ReportStatus       `json:"status"`
	TimeFrom     time.Time          `json:"timeFrom"`
	TimeTo       time.Time          `json:"timeTo"`
}

// ReportStatus represents report status
type ReportStatus string

const (
	ReportStatusQueued     ReportStatus = "Queued"
	ReportStatusProcessing ReportStatus = "Processing"
	ReportStatusRunning    ReportStatus = "Running"
	ReportStatusCanceled   ReportStatus = "Canceled"
	ReportStatusFailed     ReportStatus = "Failed"
	ReportStatusFinished   ReportStatus = "Finished"
)

// RequestReport requests a CSV report generation
func (c *Client) RequestReport(ctx context.Context, req PublicReportRequest) (*EnqueuedReportResponse, error) {
	resp, err := c.makeRequest(ctx, http.MethodPost, "/api/v0/equity/history/exports", req)
	if err != nil {
		return nil, err
	}

	var result EnqueuedReportResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetReports retrieves all requested reports and their status
func (c *Client) GetReports(ctx context.Context) ([]ReportResponse, error) {
	resp, err := c.makeRequest(ctx, http.MethodGet, "/api/v0/equity/history/exports", nil)
	if err != nil {
		return nil, err
	}

	var reports []ReportResponse
	if err := c.handleResponse(resp, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}