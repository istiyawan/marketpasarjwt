package controllertests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/istiyawan/marketpasarjwt/api/controllers"
	"github.com/istiyawan/marketpasarjwt/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var userInstance = models.User{}
var postInstance = models.Post{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}
