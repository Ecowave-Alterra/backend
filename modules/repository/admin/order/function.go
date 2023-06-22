package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

// type TransactionResponse struct {
// 	ReceiptNumber     string
// 	TransactionId     string
// 	Unit              uint
// 	TotalPrice        float64
// 	OrderDate         time.Time
// 	StatusTransaction string
// }

func (or *orderRepo) GetAllOrder(transactions *[]te.TransactionResponse, offset, pageSize int) ([]te.TransactionResponse, int64, error) {
	var count int64
	if err := or.db.Model(&te.Transaction{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := or.db.Model(&te.Transaction{}).
		Select("transactions.receipt_number AS ReceiptNumber, transactions.transaction_id AS TransactionId, user_details.name AS Name, (SELECT COUNT(*) FROM transaction_details WHERE transaction_details.transaction_id = transactions.id) AS Unit, total_price AS TotalPrice, transactions.created_at AS OrderDate, status_transaction AS StatusTransaction").
		Joins("JOIN transaction_details ON transaction_details.transaction_id = transactions.id").
		Joins("JOIN users ON transactions.user_id = users.id").
		Joins("JOIN user_details ON users.id = user_details.user_id").
		Offset(offset).
		Limit(pageSize).
		Scan(&transactions).Error; err != nil {
		return nil, 0, nil
	}

	return *transactions, count, nil
}
