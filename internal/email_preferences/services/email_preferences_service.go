package services

import (
	"errors"
	"notifications-system/internal/email_preferences/models"
	"gorm.io/gorm"
)

type EmailPreferenceService struct {
	db *gorm.DB
}

func NewEmailPreferenceService(db *gorm.DB) *EmailPreferenceService {
	return &EmailPreferenceService{db: db}
}

// Create a new email preference
func (s *EmailPreferenceService) Create(emailPreference *models.EmailPreference) error {
	// if emailPreference.UserID == 0 || emailPreference.Category == "" || emailPreference.Type == "" {
	// 	return errors.New("user ID, category, and type are required")
	// }
	return s.db.Create(emailPreference).Error
}

// Get email preference by ID
func (s *EmailPreferenceService) GetByID(id string) (*models.EmailPreference, error) {
	var emailPreference models.EmailPreference
	if err := s.db.First(&emailPreference, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &emailPreference, nil
}

// Update email preference
func (s *EmailPreferenceService) Update(id string, updatedPreference *models.EmailPreference) error {
	result := s.db.Model(&models.EmailPreference{}).Where("id = ?", id).Updates(updatedPreference)
	if result.RowsAffected == 0 {
		return errors.New("no record found to update")
	}
	return result.Error
}

// Delete email preference
func (s *EmailPreferenceService) Delete(id string) error {
	result := s.db.Delete(&models.EmailPreference{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return errors.New("no record found to delete")
	}
	return result.Error
}