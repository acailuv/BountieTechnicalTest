package bountie

import (
	"main/database/models"
	"main/utils"
	"net/http"
)

func (h *bountie) DecimalToBinary(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req models.DecimalToBinary
	err := utils.ReadBody(r.Body, &req)
	if err != nil {
		utils.InternalServerError(w, "Read Body", err)
		return
	}

	err = h.validator.Struct(req)
	if err != nil {
		utils.BadRequest(w, err.Error())
		return
	}

	decimalString := utils.ConvertDecimalToBinary(req.Decimal)

	utils.OK(w, decimalString)
}
