package factory

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/meloneg/mln_data_pool/common/file"
	"github.com/rs/zerolog/log"
)

var factory PublicFactoryInterface

func TestMain(m *testing.M) {
	usersFIle, err := file.LoadFile("../../../static/seeds/initial_seeder_book_store.xlsx")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load users file")
	}
	factory = NewPublicFactory(usersFIle)
	os.Exit(m.Run())
}
