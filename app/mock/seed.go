package mock

import (
	"gwiapi/app/helpers"
	"gwiapi/app/models"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

func Seed(db *gorm.DB) error {
	// Participants Seeding
	err := seedParticipants(db)
	if err != nil {
		return err
	}

	// Users Seeding
	err = seedUsers(db)
	if err != nil {
		return err
	}

	// Assets Seeding
	err = seedAssets(db)
	if err != nil {
		return err
	}
	return nil
}

func seedParticipants(db *gorm.DB) error {

	var err error
	numberOfParticipants := 1000

	// hoursSpendOnSocialMedia
	hoursMin := 0.0
	hoursMax := 12.0
	hoursPeak := 2.0

	// Age Range
	ageMin := 15
	ageMax := 85

	// Countries
	country := models.Country{}
	countries := country.GetAllCountries()
	countriesLen := len(countries)

	participants := []models.Participant{}
	for i := 1; i <= numberOfParticipants; i++ {
		participants = append(participants, models.Participant{
			Gender:                  helpers.StringTernary(GenerateRandomBool() == true, "m", "f"),
			HoursSpendOnSocialDaily: uint8(GenerateNormalDistribution(hoursPeak, hoursMin, hoursMax)),
			Age:                     uint8(rand.Intn(ageMax-ageMin) + ageMin),
			CountryCode:             countries[rand.Intn(countriesLen)].Alpha2Code,
			CreatedAt:               time.Now(),
		})
	}

	for _, p := range participants {
		err = db.Model(&models.Participant{}).Save(&p).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Email:    "user1@example.com",
			Password: "secretpassword1",
		},
		{
			Email:    "user2@example.com",
			Password: "secretpassword2",
		},
	}
	for _, u := range users {
		u.Prepare()
		u.Validate("")
		err := db.Model(&models.User{}).Create(&u).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func seedAssets(db *gorm.DB) error {

	charts := []models.Chart{
		models.Chart{
			GroupedMetric: "age",
		},
		models.Chart{
			GroupedMetric: "CountryCode",
		},
	}
	for _, c := range charts {
		uid := uint32(rand.Intn(2) + 1)
		_, _, err := c.CreateAssetChart(db, uid)
		if err != nil {
			return err
		}
	}

	insights := []models.Insight{
		models.Insight{
			Gender:          "f",
			BirthCountry:    "",
			HoursComparator: ">",
			HoursReference:  5,
		},
		models.Insight{
			Gender:          "m",
			BirthCountry:    "US",
			HoursComparator: "<=",
			HoursReference:  1,
		},
		models.Insight{
			Gender:          "f",
			BirthCountry:    "US",
			HoursComparator: "<=",
			HoursReference:  1,
		},
		models.Insight{
			Gender:          "f",
			BirthCountry:    "",
			HoursComparator: ">",
			HoursReference:  5,
		},
		models.Insight{
			Gender:          "m",
			BirthCountry:    "",
			HoursComparator: ">",
			HoursReference:  6,
		},
		models.Insight{
			Gender:          "",
			BirthCountry:    "GR",
			HoursComparator: ">",
			HoursReference:  6,
		},
	}
	for _, i := range insights {
		uid := uint32(rand.Intn(2) + 1)
		_, _, err := i.CreateAssetInsight(db, uid)
		if err != nil {
			return err
		}
	}

	audiences := []models.Audience{
		models.Audience{
			AgeFrom:     20,
			AgeTo:       25,
			CountryCode: "",
			Gender:      "m",
		},
		models.Audience{
			AgeFrom:     20,
			AgeTo:       25,
			CountryCode: "",
			Gender:      "f",
		},
		models.Audience{
			AgeFrom:     25,
			AgeTo:       30,
			CountryCode: "GR",
			Gender:      "",
		},
		models.Audience{
			AgeFrom:     25,
			AgeTo:       30,
			CountryCode: "US",
			Gender:      "",
		},
		models.Audience{
			AgeFrom:     30,
			AgeTo:       35,
			CountryCode: "GR",
			Gender:      "",
		},
		models.Audience{
			AgeFrom:     30,
			AgeTo:       35,
			CountryCode: "GR",
			Gender:      "f",
		},
		models.Audience{
			AgeFrom:     30,
			AgeTo:       40,
			CountryCode: "US",
			Gender:      "",
		},
	}
	for _, i := range audiences {
		uid := uint32(rand.Intn(2) + 1)
		_, _, err := i.CreateAssetAudience(db, uid)
		if err != nil {
			return err
		}
	}

	return nil

}
