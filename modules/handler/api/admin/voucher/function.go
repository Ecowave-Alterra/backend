package voucher

import (
	"net/http"
	"strconv"
	"time"

	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

	"github.com/labstack/echo/v4"
)

func (vh *VoucherHandler) CreateVoucher(c echo.Context) error {
	voucherTypeIDstr := c.FormValue("voucherTypeID")
	voucherTypeID, err := strconv.ParseUint(voucherTypeIDstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "VoucherTypeID must be numeric",
		})
	}

	startDateStr := c.FormValue("startDate")
	startDate, err := time.Parse("02 January 2006", startDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "err",
		})
	}

	endDateStr := c.FormValue("endDate")
	endDate, err := time.Parse("02 January 2006", endDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "err",
		})
	}

	minimumPurchaseStr := c.FormValue("minimumPurchase")
	minimumPurchase, err := strconv.ParseFloat(minimumPurchaseStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Minimum belanja harus berupa angka. Contoh : 100500",
		})
	}

	maximumDiscountStr := c.FormValue("maximumDiscount")
	maximumDiscount, err := strconv.ParseFloat(maximumDiscountStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Maksimum potongan harga harus berupa angka. Contoh 100500",
		})
	}

	discountPercentStr := c.FormValue("discountPercent")
	discountPercent, err := strconv.ParseFloat(discountPercentStr, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Diskon harus berupa angka dari 5 - 100. Contoh : 50",
		})
	}

	claimableCountStr := c.FormValue("claimableCount")
	claimableCount, err := strconv.ParseUint(claimableCountStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Jumlah voucher harus berupa angka. Contoh 100500",
		})
	}

	maxClaimLimitStr := c.FormValue("maxClaimLimit")
	maxClaimLimit, err := strconv.ParseUint(maxClaimLimitStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Maksimum klaim voucher harus berupa angka. Contoh 100500",
		})
	}

	voucher := ve.Voucher{
		VoucherTypeID:   uint(voucherTypeID),
		StartDate:       startDate,
		EndDate:         endDate,
		MinimumPurchase: minimumPurchase,
		MaximumDiscount: maximumDiscount,
		DiscountPercent: discountPercent,
		ClaimableCount:  uint(claimableCount),
		MaxClaimLimit:   uint(maxClaimLimit),
	}

	err = vh.voucherUsecase.CreateVoucher(&voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Failed to create voucher",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menambahkan voucher",
	})
}

func (vh *VoucherHandler) GetAllVoucher(c echo.Context) error {
	var vouchers []ve.Voucher

	vouchers, err := vh.voucherUsecase.GetAllVoucher(&vouchers)
	if err != nil {
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengambil data voucher",
			})
		}
	}

	var voucherResponses []ve.VoucherResponse
	for _, voucher := range vouchers {
		outputDateFormat := "02 January 2006"
		startDate := voucher.StartDate.Format(outputDateFormat)
		endDate := voucher.EndDate.Format(outputDateFormat)

		voucherResponse := ve.VoucherResponse{
			Type:           voucher.VoucherType.Type,
			ClaimableCount: voucher.ClaimableCount,
			StartDate:      startDate,
			EndDate:        endDate,
		}

		voucherResponses = append(voucherResponses, voucherResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message":  "Berhasil mengambil data voucher",
		"Vouchers": voucherResponses,
	})
}
