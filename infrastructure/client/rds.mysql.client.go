package client

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	rdsMySqlSingleton *RdsGormClient
	once              sync.Once
)

// Interfaz para el cliente RDS con GORM
type GormRdsClient interface {
	GetInstance() (*gorm.DB, error)
}

// Estructura que implementa la interfaz
type RdsGormClient struct {
	db    *gorm.DB
	dbErr error
}

// Método que retorna la instancia de la base de datos
func (c *RdsGormClient) GetInstance() (*gorm.DB, error) {
	return c.db, c.dbErr
}

// Implementación singleton para obtener la instancia actual
func NewRdsMySQLGormClient() (GormRdsClient, error) {
	var err error
	once.Do(func() {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASS")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&tls=skip-verify",
			dbUser, dbPassword, dbHost, dbPort, dbName,
		)

		db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			err = dbErr
			fmt.Println("Error al conectar con la base de datos:", dbErr)
			return
		}

		rdsMySqlSingleton = &RdsGormClient{db: db, dbErr: dbErr}
	})

	if rdsMySqlSingleton == nil {
		return nil, fmt.Errorf("error al inicializar la conexión con la base de datos: %v", err)
	}

	return rdsMySqlSingleton, nil
}

// ResetRdsGormClient reinicia el singleton para pruebas o reconexión.
func ResetRdsGormClient() {
	rdsMySqlSingleton = nil
	once = sync.Once{}
}
