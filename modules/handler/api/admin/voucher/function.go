package voucher

import (
	"math"
	"net/http"
	"strconv"
	"time"

	cs "github.com/berrylradianh/ecowave-go/helper/customstatus"
	ve "github.com/berrylradianh/ecowave-go/modules/entity/voucher"

	"github.com/labstack/echo/v4"
)

func (vh *VoucherHandler) GetAllVoucher(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	vouchers, total, err := vh.voucherUsecase.GetAllVoucher(offset, pageSize)
	if err != nil {
		code, msg := cs.CustomStatus(err.Error())
		return c.JSON(code, echo.Map{
			"Status":  code,
			"Message": msg,
		})
	}

	var voucherResponses []ve.VoucherResponse
	for _, voucher := range *vouchers {
		outputDateFormat := "02 January 2006"
		startDate := voucher.StartDate.Format(outputDateFormat)
		endDate := voucher.EndDate.Format(outputDateFormat)

		voucherResponse := ve.VoucherResponse{
			VoucherId:      voucher.VoucherId,
			Type:           voucher.VoucherType.Type,
			ClaimableCount: voucher.ClaimableCount,
			StartDate:      startDate,
			EndDate:        endDate,
		}

		voucherResponses = append(voucherResponses, voucherResponse)
	}

	if len(voucherResponses) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list voucher",
			"Status":  http.StatusNotFound,
		})
	} else {
		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": "Halaman Tidak Ditemukan",
				"Status":  http.StatusNotFound,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Vouchers":  voucherResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	}
}

func (vh *VoucherHandler) CreateVoucher(c echo.Context) error {
	voucherTypeIDstr := c.FormValue("voucherTypeID")
	voucherTypeID, err := strconv.ParseUint(voucherTypeIDstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "VoucherTypeID harus berupa angka",
		})
	}

	startDateStr := c.FormValue("startDate")
	startDate, err := time.Parse("02 January 2006", startDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Tanggal mulai tidak valid",
		})
	}

	endDateStr := c.FormValue("endDate")
	endDate, err := time.Parse("02 January 2006", endDateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Tanggal berakhir tidak valid",
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
			"Message": "Gagal membuat voucher baru",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menambahkan voucher",
	})
}

func (vh *VoucherHandler) UpdateVoucher(c echo.Context) error {
	voucherID := c.Param("id")
	voucherTypeIDstr := c.FormValue("voucherTypeID")
	voucherTypeID, err := strconv.ParseUint(voucherTypeIDstr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "VoucherTypeID harus berupa angka",
		})
	}

	switch voucherTypeID {
	case 1:
		startDateStr := c.FormValue("startDate")
		startDate, err := time.Parse("02 January 2006", startDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Tanggal mulai tidak valid",
			})
		}

		endDateStr := c.FormValue("endDate")
		endDate, err := time.Parse("02 January 2006", endDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Tanggal berakhir tidak valid",
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
			VoucherTypeID:  uint(voucherTypeID),
			StartDate:      startDate,
			EndDate:        endDate,
			ClaimableCount: uint(claimableCount),
			MaxClaimLimit:  uint(maxClaimLimit),
		}

		err = vh.voucherUsecase.UpdateVoucher(voucherID, &voucher)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengubah voucher",
			})
		}
	case 2:
		startDateStr := c.FormValue("startDate")
		startDate, err := time.Parse("02 January 2006", startDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Tanggal mulai tidak valid",
			})
		}

		endDateStr := c.FormValue("endDate")
		endDate, err := time.Parse("02 January 2006", endDateStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"Message": "Tanggal berakhir tidak valid",
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

		err = vh.voucherUsecase.UpdateVoucher(voucherID, &voucher)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"Message": "Gagal mengubah voucher",
			})
		}
	default:
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"Message": "Voucher type id tidak valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil mengubah voucher",
	})
}

func (vh *VoucherHandler) DeleteVoucher(c echo.Context) error {
	voucherID := c.Param("id")

	var voucher ve.Voucher
	err := vh.voucherUsecase.DeleteVoucher(voucherID, &voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": "Gagal menghapus voucher",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menghapus voucher",
	})
}

func (vh *VoucherHandler) FilterVouchersByType(c echo.Context) error {
	voucherType := c.QueryParam("type")

	var voucher []ve.Voucher
	vouchers, err := vh.voucherUsecase.FilterVouchersByType(voucherType, &voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": "Gagal memfilter voucher",
		})
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
		"Message":  "Berhasil memfilter data voucher",
		"Vouchers": voucherResponses,
	})
}
