package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	jsonpatch "github.com/evanphx/json-patch"

	"gwiapi/app/auth"
	"gwiapi/app/models"
	"gwiapi/app/responses"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetAssets will bring asset of the user
func GetAssets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	// Get Pagination variables
	page, found := mux.Vars(r)["page"]
	if !found {
		page = "1"
	}
	pg, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	limit, found := mux.Vars(r)["limit"]
	if !found {
		limit = "10"
	}
	lmt, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	asset := models.Asset{}

	// Select only User Assets from Token
	uid, err := auth.ExtractTokenID(r)

	assets, err := asset.GetAssets(db, uid, uint32(pg), uint32(lmt))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, assets)
}

func GetAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id, found := mux.Vars(r)["id"]
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Asset's Id was not passed to Endpoint"))
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("Provider Asset's Id could not be parsed"))
		return
	}

	// Select only User Asset from Token
	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	a, err = a.GetAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	if a.ID == 0 {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	switch a.AssetType {
	case "chart":
		a.Chart = getAssetChartDetails(db)
	case "insight":
		a.Insight = getAssetInsightDetails(db)
	case "audience":
		a.Audience = getAssetAudienceDetails(db)
	default:
	}

	responses.JSON(w, http.StatusOK, a)
}

func getAssetChartDetails(db *gorm.DB) *models.ChartDetails {
	c := &models.ChartDetails{}
	// TODO: Incomplete
	return c
}

func getAssetInsightDetails(db *gorm.DB) *models.InsightDetails {
	c := &models.InsightDetails{}
	// TODO: Incomplete
	return c
}

func getAssetAudienceDetails(db *gorm.DB) *models.AudienceDetails {
	c := &models.AudienceDetails{}
	// TODO: Incomplete
	return c
}

func DeleteAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	id, found := mux.Vars(r)["id"]
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("No asset Id provided"))
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)

	// Select only User Asset from Token
	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	rowsDeleted, err := a.DeleteAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", assetId))
	w.Header().Set("Rows Affected", fmt.Sprintf("%d", rowsDeleted))
	responses.JSON(w, http.StatusNoContent, "")
}

func UpdateAsset(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	id, found := mux.Vars(r)["id"]

	// *Uneccesary based on Route definition
	if !found {
		responses.ERROR(w, http.StatusInternalServerError, errors.New("No asset Id provided"))
		return
	}
	assetId, err := strconv.ParseUint(id, 10, 32)

	// Get Original
	asset := models.Asset{}
	originalAsset, err := asset.GetAsset(db, uint32(assetId), uid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	originalJSON, err := json.Marshal(originalAsset)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// Get Patch
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	patchOperation, err := jsonpatch.DecodePatch(body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Apply Patch
	targetJSON, err := patchOperation.Apply(originalJSON)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	targetAsset := models.Asset{}
	err = json.Unmarshal(targetJSON, &targetAsset)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedAsset, err := targetAsset.UpdateAsset(db, uint32(assetId), uid)
	responses.JSON(w, http.StatusOK, updatedAsset)
}

func CreateAssetChart(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	c := &models.Chart{}

	// TODO: Parse Body

	a, c, err = c.CreateAssetChart(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s/assets/%d", r.Host, r.RequestURI, a.ID))
	responses.JSON(w, http.StatusOK, a)
}

func CreateAssetInsight(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	a := &models.Asset{}
	i := &models.Insight{}

	// TODO: Parse Body

	a, i, err = i.CreateAssetInsight(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s/assets/%d", r.Host, r.RequestURI, a.ID))
	responses.JSON(w, http.StatusOK, a)
}

func CreateAssetAudience(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	uid, err := auth.ExtractTokenID(r)

	ass := &models.Asset{}
	aud := &models.Audience{}

	// TODO: Parse Body

	ass, aud, err = aud.CreateAssetAudience(db, uid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s/assets/%d", r.Host, r.RequestURI, ass.ID))
	responses.JSON(w, http.StatusOK, ass)
}
