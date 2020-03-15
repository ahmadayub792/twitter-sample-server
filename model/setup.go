package model

import (
	"fmt"
	"strings"

	"github.com/ahmadayub792/twitter-sample-server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Setup() error {

	if err := config.DB.Transaction(migrate); err != nil {
		return err
	}

	return nil
}

func migrate(tx *gorm.DB) error {
	models := []interface{}{
		ClientUser{},
		ClientTarget{},
		User{},
		Client{},
		Target{},
	}

	if err := tx.DropTableIfExists(models...).Error; err != nil {
		return fmt.Errorf("drop table: %w", err)
	}

	if err := tx.Exec(`Drop TYPE "role"`).Error; err != nil {
		return fmt.Errorf("drop role enum: %w", err)
	}

	if err := tx.Exec(`DROP TYPE "client_role"`).Error; err != nil {
		return fmt.Errorf("drop client role enum: %w", err)
	}

	roles := []Role{RoleAdmin, RoleClient}
	strList := []string{}
	for _, role := range roles {
		strList = append(strList, fmt.Sprintf("'%s'", role))
	}
	query := fmt.Sprintf(`CREATE TYPE "role" AS ENUM (%s)`, strings.Join(strList, ","))
	if err := tx.Exec(query).Error; err != nil {
		return fmt.Errorf("create role enum: %w", err)
	}

	clientRoles := []ClientRole{ClientRoleAdmin, ClientRoleManager, ClientRoleAnalyst}
	strList = []string{}
	for _, role := range clientRoles {
		strList = append(strList, fmt.Sprintf("'%s'", role))
	}
	query = fmt.Sprintf(`CREATE TYPE "client_role" AS ENUM(%s)`, strings.Join(strList, ","))
	if err := tx.Exec(query).Error; err != nil {
		return fmt.Errorf("create client role enum: %w", err)
	}

	if err := tx.CreateTable(models...).Error; err != nil {
		return fmt.Errorf("create table: %w", err)
	}

	if err := tx.Model(ClientUser{}).AddForeignKey("client_id", "clients(id)", "CASCADE", "CASCADE").Error; err != nil {
		return fmt.Errorf("ClientUser constraint (client_id): %w", err)
	}

	if err := tx.Model(ClientUser{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error; err != nil {
		return fmt.Errorf("ClientUser constraint (user_id): %w", err)
	}

	if err := tx.Model(ClientTarget{}).AddForeignKey("client_id", "clients(id)", "CASCADE", "CASCADE").Error; err != nil {
		return fmt.Errorf("ClientTarget constraint (client_id): %w", err)
	}

	if err := tx.Model(ClientTarget{}).AddForeignKey("target_id", "targets(id)", "CASCADE", "CASCADE").Error; err != nil {
		return fmt.Errorf("ClientTarget constraint (target_id): %w", err)
	}

	return nil
}
