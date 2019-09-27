package mock

import (
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

	// TODO: Users Seeding

	// TODO: Assets Seeding

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
			Gender:                  TernaryString(GenerateRandomBool() == true, "m", "f"),
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
	/*
		// Tried to implement Concurrent "Bulk Insert" 
		// Having Problems with duplicate primary key constrains
		gopherCount := 10
		runtime.GOMAXPROCS(2)
		db.DB().SetMaxOpenConns(gopherCount)
		channelInsert := make(chan int, gopherCount)
		go func(gocherCount int) {
			var wg sync.WaitGroup
			wg.Add(gocherCount)
			for _, p := range participants {
				go func() {
					defer wg.Done()
					db.Model(&models.Participant{}).Save(&p)
					channelInsert <- 1
				}()
				db.Model(&models.Participant{}).Save(&p)
			}
			wg.Wait()
			close(channelInsert)
		}(10)
	*/
	return nil
}
