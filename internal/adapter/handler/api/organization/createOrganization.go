package organization

import (
	"gestor_de_usuario/internal/adapter/handler/api/request"
	"gestor_de_usuario/internal/adapter/handler/api/response"
	"github.com/gin-gonic/gin"
)

func (organizationHandler *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var organizationRequest request.OrganizationRequest
	if err := c.BindJSON(&organizationRequest); err != nil {
		response.ValidationError(c, err)
		return
	}

	newOrganization, err := organizationHandler.OrganizationService.CreateOrganization(&organizationRequest)

	if err != nil {
		response.HandleError(c, err)
		return
	}

	resp := response.NewOrganizationResponse(newOrganization)

	response.HandleSuccess(c, resp)

}
