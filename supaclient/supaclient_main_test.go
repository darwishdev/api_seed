package supaclient

import (
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/meloneg/mln_data_pool/config"
)

var service SupabaseServiceInterface

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	service, err = NewSupabaseService(config.SupaUrl, config.SupaKey)
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	os.Exit(m.Run())
}
