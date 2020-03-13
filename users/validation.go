package users

import (
	"regexp"
	"fmt"
	"github.com/globalsign/mgo/bson"

	"github.com/Haski007/BitMedia_labs/database"
	"github.com/Haski007/BitMedia_labs/config"
)

func validateUser(user user) error {
	err := validateEmail(user.Email)
	if err != nil {
		return err
	}

	err = validateDate(user.BirthDate)
	if err != nil {
		return err
	}

	err = validateGender(user.Gender)
	if err != nil {
		return err
	}
	return nil
}

func validateEmail(email string) error {
	var user user

	database.UsersCollection.Find(bson.M{ "email" : email }).One(&user)
	if user.Email != "" {
		return fmt.Errorf("User with email [%s] already exists!", email)
	} else if ok, _ := regexp.MatchString(config.EmailPattern, email); !ok {
		return fmt.Errorf("Email [%s] is not valid", email)
	}
	return nil
}

func validateDate(date string) error {
	if ok, _ := regexp.MatchString(config.DatePattern, date); !ok {
		return fmt.Errorf("Wrong date format [%s]", date)
	}
	return nil
}

func validateGender(gender string) error {
	if gender != "Male" && gender != "Female" {
		return fmt.Errorf(`Wrong format of gender [%s], it shold be "Male/Female".`, gender)
	}
	return nil
}