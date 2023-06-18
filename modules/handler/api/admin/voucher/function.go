package voucher

import (
	"math"
	"net/http"
	"strconv"

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
			VoucherId:          voucher.VoucherId,
			Type:               voucher.VoucherType.Type,
			ClaimableUserCount: voucher.ClaimableUserCount,
			StartDate:          startDate,
			EndDate:            endDate,
		}

		voucherResponses = append(voucherResponses, voucherResponse)
	}

	if len(voucherResponses) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list voucher",
			"Status":  http.StatusNotFound,
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Vouchers":  voucherResponses,
			"Page":      page,
			"TotalPage": int(math.Ceil(float64(total) / float64(pageSize))),
			"Status":    http.StatusOK,
		})
	}
}

func (vh *VoucherHandler) GetVoucherById(c echo.Context) error {
	var voucherResponse ve.VoucherResponse
	id := c.Param("id")

	voucher, err := vh.voucherUsecase.GetVoucherById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	outputDateFormat := "02 January 2006"
	startDate := voucher.StartDate.Format(outputDateFormat)
	endDate := voucher.EndDate.Format(outputDateFormat)

	voucherResponse = ve.VoucherResponse{
		VoucherId:          voucher.VoucherId,
		Type:               voucher.VoucherType.Type,
		StartDate:          startDate,
		EndDate:            endDate,
		MinimumPurchase:    voucher.MinimumPurchase,
		MaximumDiscount:    voucher.MaximumDiscount,
		DiscountPercent:    voucher.DiscountPercent,
		ClaimableUserCount: voucher.ClaimableUserCount,
		MaxClaimLimit:      voucher.MaxClaimLimit,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Voucher": voucherResponse,
		"Status":  http.StatusOK,
	})
}

func (vh *VoucherHandler) CreateVoucher(c echo.Context) error {
	var voucher *ve.VoucherRequest
	if err := c.Bind(&voucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusBadRequest,
		})
	}

	validVoucherTypeId := map[int]bool{1: true, 2: true}
	if voucher.VoucherId != "" || !validVoucherTypeId[int(voucher.VoucherTypeID)] {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Message": "Anda gagal membuat voucher",
			"Status":  http.StatusBadRequest,
		})
	}

	err := vh.voucherUsecase.CreateVoucher(voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"Message": "Anda berhasil membuat voucher",
		"Status":  http.StatusCreated,
	})
}

func (vh *VoucherHandler) UpdateVoucher(c echo.Context) error {
	var voucher ve.Voucher
	voucherID := c.Param("id")

	voucherBefore, err := vh.voucherUsecase.GetVoucherById(voucherID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": err.Error(),
			"Status":  http.StatusNotFound,
		})
	}

	if err := c.Bind(&voucher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusBadRequest,
		})
	}

	validVoucherTypeId := map[int]bool{1: true, 2: true}
	if voucher.VoucherId != "" && (voucher.VoucherTypeID != voucherBefore.VoucherTypeID || (voucher.VoucherId != "" && !validVoucherTypeId[int(voucher.VoucherTypeID)])) {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"Message": "Anda gagal mengubah voucher",
			"Status":  http.StatusBadRequest,
		})
	}

	err = vh.voucherUsecase.UpdateVoucher(voucherID, &voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err.Error(),
			"Status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil mengubah voucher",
		"Status":  http.StatusOK,
	})
}

func (vh *VoucherHandler) DeleteVoucher(c echo.Context) error {
	voucherID := c.Param("id")

	var voucher ve.Voucher
	err := vh.voucherUsecase.DeleteVoucher(voucherID, &voucher)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"Message": err,
			"Status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Message": "Anda berhasil menghapus voucher",
		"Status":  http.StatusOK,
	})
}

func (vh *VoucherHandler) FilterVoucher(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize := 10
	offset := (page - 1) * pageSize

	voucherType := c.QueryParam("type")

	var vouchers *[]ve.Voucher
	vouchers, total, err := vh.voucherUsecase.FilterVoucher(voucherType, offset, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Message": err,
			"Status":  http.StatusInternalServerError,
		})
	}

	if len(*vouchers) == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{
			"Message": "Belum ada list voucher",
			"Status":  http.StatusNotFound,
		})
	} else {
		var voucherResponses []ve.VoucherResponse
		for _, voucher := range *vouchers {
			outputDateFormat := "02 January 2006"
			startDate := voucher.StartDate.Format(outputDateFormat)
			endDate := voucher.EndDate.Format(outputDateFormat)

			voucherResponse := ve.VoucherResponse{
				VoucherId:          voucher.VoucherId,
				Type:               voucher.VoucherType.Type,
				ClaimableUserCount: voucher.ClaimableUserCount,
				StartDate:          startDate,
				EndDate:            endDate,
			}

			voucherResponses = append(voucherResponses, voucherResponse)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Vouchers":  voucherResponses,
			"Page":      page,
			"TotalPage": int(math.Ceil(float64(total) / float64(pageSize))),
			"Status":    http.StatusOK,
		})
	}
}
