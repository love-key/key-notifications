package validations

import (
    "github.com/go-playground/validator/v10" // External package for validation
    "notifications-system/internal/email_preferences/models"
    "notifications-system/utils" // Import your utils package
)

// ValidateEmailPreference validates the EmailPreference struct
func ValidateEmailPreference(input *models.EmailPreference) error {
    // Initialize the validator
    validate := validator.New()

    // Perform struct-level validations using the validator package
    err := validate.Struct(input)
    if err != nil {
        return utils.NewAppError(utils.ErrCodeBadRequest, "Validation failed for email preference")
    }

    // Custom field-specific validations (Ensure proper capitalization of fields)
    if input.UserID == "" {
        return utils.NewAppError(utils.ErrCodeBadRequest, "UserID is required")
    }

    if input.Category == "" {
        return utils.NewAppError(utils.ErrCodeBadRequest, "Category is required")
    }

    if input.Type == "" {
        return utils.NewAppError(utils.ErrCodeBadRequest, "Type is required")
    }

	if !input.IsEnabled {
        return utils.NewAppError(utils.ErrCodeBadRequest, "IsEnabled is required")
    }

    // If all validations pass, return nil (no error)
    return nil
}
