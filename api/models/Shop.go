package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User repesent the structure of each user
type Shop struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	AdminName   string    `gorm:"size:255;not null;unique" json:"adminname"`
	AdminEmail  string    `gorm:"size:100;not null;unique" json:"adminemail"`
	ShopName    string    `gorm:"size:100;not null;unique" json:"shopname"`
	Stack       string    `gorm:"size:100;not null;unique" json:"stack"`
	TeamMembers string    `gorm:"size:100;not null;unique" json:"teammembers"`
	Bio         string    `gorm:"size:100;not null;unique" json:"bio"`
	Password    string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Hash fuction returns the hash of the user password
func HashShop(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword function authenticated the user password
func VerifyPasswordShop(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave function creates a hash of the password
func (u *Shop) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// Prepare function takes the input
func (u *Shop) Prepare() {
	u.ID = 0
	u.AdminName = html.EscapeString(strings.TrimSpace(u.AdminName))
	u.AdminEmail = html.EscapeString(strings.TrimSpace(u.AdminEmail))
	u.ShopName = html.EscapeString(strings.TrimSpace(u.ShopName))
	u.Stack = html.EscapeString(strings.TrimSpace(u.Stack))
	u.TeamMembers = html.EscapeString(strings.TrimSpace(u.TeamMembers))
	u.Bio = html.EscapeString(strings.TrimSpace(u.Bio))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// Validate functions ensures there is no empty input
func (u *Shop) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.AdminName == "" {
			return errors.New("Required Admin name")
		}
		if u.AdminEmail == "" {
			return errors.New("Required Lastname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.ShopName == "" {
			return errors.New("Required Shop name")
		}
		if u.Stack == "" {
			return errors.New("Required stack")
		}
		if u.TeamMembers == "" {
			return errors.New("Required Team members")
		}
		if u.Bio == "" {
			return errors.New("Required bio")
		}
		if err := checkmail.ValidateFormat(u.AdminEmail); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.AdminEmail == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.AdminEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.AdminName == "" {
			return errors.New("Required Admin name")
		}
		if u.ShopName == "" {
			return errors.New("Required Shop name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.AdminEmail == "" {
			return errors.New("Required Email")
		}
		if u.Stack == "" {
			return errors.New("Required stack")
		}
		if u.TeamMembers == "" {
			return errors.New("Required team members")
		}
		if u.Bio == "" {
			return errors.New("Required bio")
		}
		if err := checkmail.ValidateFormat(u.AdminEmail); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// SaveUser creates a new user in db
func (u *Shop) SaveShop(db *gorm.DB) (*Shop, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Shop{}, err
	}
	return u, nil
}

// FindAllUsers function returns all users in db
func (u *Shop) FindAllShop(db *gorm.DB) (*[]Shop, error) {
	var err error
	shops := []Shop{}
	err = db.Debug().Model(&Shop{}).Limit(100).Find(&shops).Error
	if err != nil {
		return &[]Shop{}, err
	}
	return &shops, err
}

// FindShopByID function returns a user by ID
func (u *Shop) FindShopByID(db *gorm.DB, uid uint32) (*Shop, error) {
	var err error
	err = db.Debug().Model(Shop{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Shop{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Shop{}, errors.New("Shop Not Found")
	}
	return u, err
}

// UpdateAUser edits existing user in db
func (u *Shop) UpdateAShop(db *gorm.DB, uid uint32) (*Shop, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Shop{}).Where("id = ?", uid).Take(&Shop{}).UpdateColumns(
		map[string]interface{}{
			"password":    u.Password,
			"adminname":   u.AdminName,
			"adminemail":  u.AdminEmail,
			"shopname":    u.ShopName,
			"stack":       u.Stack,
			"teammembers": u.TeamMembers,
			"bio":         u.Bio,
			"update_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &Shop{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Shop{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Shop{}, err
	}
	return u, nil
}

// DeleteAUser function removes a user from the db
func (u *User) DeleteAShop(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Shop{}).Where("id = ?", uid).Take(&Shop{}).Delete(&Shop{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
