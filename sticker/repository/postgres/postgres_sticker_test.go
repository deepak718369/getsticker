package postgres_test

import (
	"database/sql"
	"testing"
	"time"

	domain "github.com/getsticker/domain"
	_stickerRepo "github.com/getsticker/sticker/repository/postgres"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB         *gorm.DB
	mock       sqlmock.Sqlmock
	repository domain.StickerRepository
}

func (s *Suite) SetupSuite() {

	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)
	s.repository = s.repository
}

func (s *Suite) TestGetStickers(t *testing.T) {

	db, mock, err := sqlmock.New()
	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "sticker", "priority", "time", "updated_at", "created_at"}).
		AddRow(1, "sticker", "1", time.Now(), time.Now(), time.Now())

	query := "SELECT * FROM sticker order by time DESC"
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WillReturnRows(rows)

	value, er := _stickerRepo.NewpostgresStickerRepository(s.DB).GetStickers()
	assert.NoError(t, er)
	assert.NotNil(t, value)
}
