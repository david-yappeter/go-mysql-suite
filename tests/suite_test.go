package tests

import (
	"myapp/config"
	"myapp/entity"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type SuiteTest struct {
	suite.Suite
	db *gorm.DB
}

func TestSuite(t *testing.T) {
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "root")
	os.Setenv("DB_DATABASE", "go_mysql_suite_test")
	defer os.Unsetenv("DB_HOST")
	defer os.Unsetenv("DB_PORT")
	defer os.Unsetenv("DB_USER")
	defer os.Unsetenv("DB_PASS")
	defer os.Unsetenv("DB_DATABASE")

	suite.Run(t, new(SuiteTest))
}

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

// Setup db value
func (t *SuiteTest) SetupSuite() {
	config.ConnectGorm()
	t.db = config.GetDB()

	// Migrate Table
	for _, val := range getModels() {
		t.db.AutoMigrate(val)
	}
}

// Run After All Test Done
func (t *SuiteTest) TearDownSuite() {
	sqlDB, _ := t.db.DB()
	defer sqlDB.Close()

	// Drop Table
	for _, val := range getModels() {
		t.db.Migrator().DropTable(val)
	}
}

// Run Before a Test
func (t *SuiteTest) SetupTest() {

}

// Run After a Test
func (t *SuiteTest) TearDownTest() {

}
