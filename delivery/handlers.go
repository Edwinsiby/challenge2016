package delivery

import (
	"fmt"
	"net/http"
	"realimage/repository"
	"realimage/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

var repo *repository.LocationRepository

func init() {
	repo = repository.NewLocationRepository()
}

// AddContributor  godoc
//
// @Summary		adding contributer
// @Description	adding contributor with specific locations
// @Tags			Contributor
// @Accept			multipart/form-data
// @Produce		json
// @Success		200	{string}	success	massage
// @Router			/addcontributer [post]
func AddContributor(c *gin.Context) {
	contributorIdStr := c.Request.FormValue("contributorId")
	contributorId, err := strconv.Atoi(contributorIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}

	err = usecase.CreateContributor(contributorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "Contributer created"})
	}

}

// GetLocations  godoc
//
//	@Summary		reading locations form .csv file
//	@Description	reading locations with respect to the limit
//	@Tags			Locations
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{string}	success	massage
//	@Router			/locations [get]
func GetLocations(c *gin.Context) {
	// using limit to read data from .csv file
	limitStr := c.Request.FormValue("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}

	// as per the limit the retrived data is stored in a slice and returned

	locations, err := repo.GetLocationsByLimit(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusFound, locations)
	}
}

// SetExcludedLocations  godoc
//
//	@Summary		reading and setting locations form .csv file
//	@Description	reading and setting excluded locations to the contributor
//	@Tags			Locations
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{string}	success	massage
//	@Router			/setexludedlocations [post]
func SetExludedLocations(c *gin.Context) {
	// using start and end to specify the no.of locations assigned for the particular distributer as exluded locations
	startStr := c.Request.FormValue("start")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}
	endStr := c.Request.FormValue("end")
	end, err := strconv.Atoi(endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}

	contributorIdStr := c.Request.FormValue("contributorId")
	contributorId, err := strconv.Atoi(contributorIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}
	fmt.Println(contributorId)
	err = usecase.SetExludedContributorLocations(start, end, contributorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location assigning failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "Locations assigned successfuly"})
	}

}

// SetIncludedLocations  godoc
//
//	@Summary		reading and setting locations form .csv file
//	@Description	reading and setting included locations to the contributor
//	@Tags			Locations
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{string}	success	massage
//	@Router			/setinludedlocations [post]
func SetIncludedLocations(c *gin.Context) {
	startStr := c.Request.FormValue("start")
	start, err := strconv.Atoi(startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	endStr := c.Request.FormValue("end")
	end, err := strconv.Atoi(endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	contributorIdStr := c.Request.FormValue("contributorId")
	contributorId, err := strconv.Atoi(contributorIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}
	err = usecase.SetIncludedContributorLocations(start, end, contributorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location assigning failed"})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "Locations assigned successfuly"})
	}

}

// CheckPermission  godoc
//
//	@Summary		checking persmission
//	@Description	reading and checking included and exluded locations to the contributor
//	@Tags			Locations
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{string}	success	massage
//	@Router			/checkpermission [get]
func CheckPermission(c *gin.Context) {
	contributorIdStr := c.Request.FormValue("contributorId")
	contributorId, err := strconv.Atoi(contributorIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "string conversion failed"})
	}
	countryName := c.Request.FormValue("country")
	provinceName := c.Request.FormValue("province")
	cityName := c.Request.FormValue("city")

	err, permission := usecase.CheckContributorPermission(contributorId, countryName, provinceName, cityName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if permission {
		c.JSON(http.StatusOK, gin.H{"success": "Permission allowed"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission Denied"})
	}
}
