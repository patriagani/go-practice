package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetDailyReport - untuk report hari ini
func (repo *ReportRepository) GetDailyReport() (*models.DailyReport, error) {
	today := time.Now().Format("2006-01-02")
	return repo.GetReportByDateRange(today, today)
}

// GetReportByDateRange - untuk optional challenge dengan rentang tanggal
func (repo *ReportRepository) GetReportByDateRange(startDate, endDate string) (*models.DailyReport, error) {
	report := &models.DailyReport{}

	// Query total revenue dan total transaksi
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE DATE(created_at) BETWEEN $1 AND $2
	`

	err := repo.db.QueryRow(query, startDate, endDate).Scan(
		&report.TotalRevenue,
		&report.TotalTransaksi,
	)
	if err != nil {
		return nil, err
	}

	// Query produk terlaris
	productQuery := `
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) BETWEEN $1 AND $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`

	var productName string
	var qtyTerjual int

	err = repo.db.QueryRow(productQuery, startDate, endDate).Scan(&productName, &qtyTerjual)
	if err == sql.ErrNoRows {
		// Tidak ada transaksi, produk terlaris null
		report.ProdukTerlaris = nil
		return report, nil
	}
	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = &models.TopProduct{
		Nama:       productName,
		QtyTerjual: qtyTerjual,
	}

	return report, nil
}
