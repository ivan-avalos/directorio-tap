package database

import (
	"github.com/jinzhu/gorm"
)

type (
	// Contact es el modelo del contacto
	Contact struct {
		gorm.Model
		Name  string `gorm:"not null"` // No son comentarios, son parámetros «reflect» y proporcionan metadatos importantes al struct
		Phone string `gorm:"not null"`
	}

	// RequestContact almacena los datos del contacto del request
	RequestContact struct {
		Name  string `json:"name" validate:"required"`
		Phone string `json:"phone" validate:"required"`
	}

	// UpdateContact almacena los datos del contacto del update request
	UpdateContact struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	// ResponseContact regresa los datos del contacto para el response
	ResponseContact struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}
)

// GetContact returns *Contact from *RequestContact
func (rc *RequestContact) GetContact() *Contact {
	return &Contact{
		Name:  rc.Name,
		Phone: rc.Phone,
	}
}

// GetResponseContact returns *ResponseContact from *Contact
func (contact *Contact) GetResponseContact() *ResponseContact {
	return &ResponseContact{
		ID:    contact.ID,
		Name:  contact.Name,
		Phone: contact.Phone,
	}
}

// Create inserta un contacto en la BD
func (contact *Contact) Create() error {
	return DB().Create(contact).Error
}

// GetContacts regresa todos los contactos en la BD
func GetContacts() ([]*Contact, error) {
	contacts := make([]*Contact, 0)
	err := DB().Find(&contacts).Error
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

// GetContact regresa un contacto de la BD
func GetContact(id uint) (*Contact, error) {
	contact := new(Contact)
	err := DB().Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil, err
	}
	return contact, nil
}

// Update guarda los cambios del contacto en la BD
func (contact *Contact) Update() error {
	return DB().Save(contact).Error
}

// Delete elimina al contacto de la DB
func (contact *Contact) Delete() error {
	return DB().Delete(contact).Error
}
