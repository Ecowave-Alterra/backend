package order

import (
	te "github.com/berrylradianh/ecowave-go/modules/entity/transaction"
	"github.com/labstack/echo/v4"
)

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

func (or *orderRepo) GetOrderByID(transactionId string, transaction *te.TransactionDetailResponse) (te.TransactionDetailResponse, error) {
	if err := or.db.Model(&te.Transaction{}).
		Select("user_addresses.address AS Address, voucher_types.type AS Voucher, user_details.name AS Name, user_Addresses.phone_number AS PhoneNumber, receipt_number AS ReceiptNumber, total_product_price AS TotalPrice, total_shipping_price AS TotalShippingPrice, transactions.point AS Point, payment_method AS PaymentMethod, payment_status AS PaymentStatus, expedition_status AS ExpeditionStatus, canceled_reason AS CanceledReason, expedition_rating AS ExpeditionRating, transactions.created_at AS CreatedAt, transactions.updated_at AS UpdatedAt").
		Joins("JOIN vouchers ON vouchers.id = transactions.voucher_id").
		Joins("JOIN voucher_types ON  voucher_types.id = vouchers.voucher_type_id").
		Joins("JOIN users ON  users.id = transactions.user_id").
		Joins("JOIN user_details ON users.id = user_details.user_id").
		Joins("JOIN user_addresses ON transactions.address_id = user_addresses.id").
		Where("transactions.id = ?", transactionId).
		Scan(&transaction).Error; err != nil {
		return *transaction, err
	}
	return *transaction, nil
}

func (or orderRepo) GetOrderProducts(transactionId string, products *[]te.TransactionProductDetailResponse) ([]te.TransactionProductDetailResponse, error) {
	if err := or.db.Model(&te.TransactionDetail{}).
		Select("products.name AS ProductName, (SELECT product_image_url FROM product_images WHERE product_id = transaction_details.product_id LIMIT 1) AS ProductImageUrl, transaction_details.qty AS Qty").
		Joins("JOIN products ON products.product_id = transaction_details.product_id").
		Where("transaction_details.transaction_id = ?", transactionId).Scan(&products).Error; err != nil {
		return nil, err
	}
	return *products, nil
}
