package repository

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"test-backend/internal/config"
	"test-backend/internal/entity"
	"test-backend/internal/util"
	error_wrapper "test-backend/internal/util/error_wrapper"

	"github.com/gofiber/fiber/v3"
)

type UserRepository struct {
	config *config.Configuration
}

type IUserRepository interface {
	GetUser(ctx fiber.Ctx, id string) ([]entity.Employee, error)
}

func NewUserRepository(
	config *config.Configuration,
) IUserRepository {
	return &UserRepository{
		config: config,
	}
}

// GetUser...
func (s *UserRepository) GetUser(ctx fiber.Ctx, id string) ([]entity.Employee, error) {
	log.Printf("[Service:GetUser] Request Id: %s", id)

	baseDir := util.ProcessFilesDirectory(s.config.FilesDirectory)
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		log.Printf("[Service:GetUser] Error reading directory %s: %v", baseDir, err)
		return nil, errors.New(error_wrapper.INTERNAL_ERROR.String())
	}

	var targetFilePath string
	foundFile := false
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileName := entry.Name()
		fileExt := filepath.Ext(fileName)
		baseName := strings.TrimSuffix(fileName, fileExt)

		//TODO: Check if the base name of the file matches the requested ID
		if baseName == id && strings.ToLower(fileExt) == ".csv" {
			targetFilePath = filepath.Join(baseDir, fileName)
			foundFile = true
			break
		}
	}

	if !foundFile {
		log.Printf("[Service:GetUser] No matching CSV file found for ID: %s in %s", id, baseDir)
		return []entity.Employee{}, errors.New(error_wrapper.NOT_FOUND.String())
	}

	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Printf("[Service:GetUser] Error opening file %s: %v", targetFilePath, err)
		return nil, errors.New(error_wrapper.INTERNAL_ERROR.String())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("[Service:GetUser] Error reading CSV records from %s: %v", targetFilePath, err)
		return nil, errors.New(error_wrapper.INTERNAL_ERROR.String())
	}

	if len(records) == 0 {
		log.Printf("[Service:GetUser] CSV file %s is empty or only header.", targetFilePath)
		return []entity.Employee{}, nil
	}

	employees := []entity.Employee{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		employee := entity.Employee{
			Id:   util.ConvertStringToInt(record[0]),
			Name: record[1],
			Age:  util.ConvertStringToInt(record[2]),
			Team: record[3],
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
