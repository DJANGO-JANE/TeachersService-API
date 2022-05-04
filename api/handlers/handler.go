package handlers

import (
	"github.com/DJANGO-JANE/internal/interfaces"
	log "github.com/sirupsen/logrus"
	"github.com/DJANGO-JANE/utils/helpers"
	"github.com/gin-gonic/gin"
)

//Configure services into the handler
type Handler struct {
	response helpers.ResponsePayload
	//tokenMaker         token.Maker
	StaffRepository interfaces.IStaffRepository
}


type Config struct {
	Routing *gin.Engine
	StaffRepository interfaces.IStaffRepository
	BaseUrl string
}

func NewHandler(config *Config){
	handler := &Handler{
		StaffRepository: config.StaffRepository,
	}
}

router := config.Routing.DefaultRouter(config.BaseUrl)

router.GET("/index", handler.)