package handlers

import (
	"github.com/DJANGO-JANE/models"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

func (handler *Handler) RetrieveALlStaff(context *gin.Context) {
	span, ctx := apm.StartSpan(context.Request.Context(),
		"RetrieveAllStaffHandler",
		"request")
	defer span.End()

	var results []models.Staff

	if err := context.BindQuery(&results); err != nil {
		handler.response.Successful = false
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
			err.Error())

		handler.response.Code = 500
		handler.response.Dispatch(context)
		return
	}

	results, err := handler.StaffRepository.GetAllStaff
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("An error occurred when retrieving all staff from database")

		handler.response.Code = 500
		handler.response.ErrorMessages = append(handler.response.ErrorMessages,
			err.Error)
		handler.response.Successful = false

		return
	} else {
		handler.response.Code = 200
		handler.response.Data = results
		handler.response.Successful = true
		handler.response.Dispatch(context)
		return
	}

}
