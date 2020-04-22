package controllers

import (
	"net/http"
	"strconv"

	"directorio-tap/database"
	"directorio-tap/utils"

	"github.com/labstack/echo"
)

// CreateContact crea un contacto
func CreateContact(c echo.Context) error {
	requestContact := new(database.RequestContact)
	if err := c.Bind(requestContact); err != nil { // Leer el JSON recibido y llenar el objeto con los datos
		return utils.ProcessError(err)
	}
	if err := c.Validate(requestContact); err != nil { // Validar el request según las condiciones definidas
		return utils.ProcessError(err)
	}
	contact := requestContact.GetContact()
	if err := contact.Create(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, contact.GetResponseContact()) // Regresar un response 200 OK con un JSON del contacto
}

// GetContacts regresa todos los contactos
func GetContacts(c echo.Context) error {
	contacts, err := database.GetContacts()
	if err != nil {
		return utils.ProcessError(err)
	}
	responseContacts := make([]*database.ResponseContact, 0)
	for _, con := range contacts { // Obtener los ResponseContact de los Contact obtenidos para el response
		responseContacts = append(responseContacts, con.GetResponseContact())
	}
	return c.JSON(http.StatusOK, responseContacts)
}

// GetContact regresa el contacto con un ID
func GetContact(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id")) // Obtener parámetro :id de la ruta y convertir a int
	if err != nil {
		return utils.ProcessError(err)
	}
	contact, err := database.GetContact(uint(id))
	if err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, contact.GetResponseContact())
}

// UpdateContact modifica el contacto con un ID
func UpdateContact(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ProcessError(err)
	}
	updateContact := new(database.UpdateContact)
	if err := c.Bind(updateContact); err != nil {
		return utils.ProcessError(err)
	}
	contact, err := database.GetContact(uint(id))
	if err != nil {
		return utils.ProcessError(err)
	}
	if updateContact.Name != "" { // Si se proporcionó el campo «name», actualizar el nombre del contacto con lo que tenga
		contact.Name = updateContact.Name
	}
	if updateContact.Phone != "" { // Si se proporcionó el campo «phone», actualizar el teléfono del contacto con lo que tenga
		contact.Phone = updateContact.Phone
	}
	if err := contact.Update(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, contact.GetResponseContact())
}

// DeleteContact elimina el contacto con un ID
func DeleteContact(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.ProcessError(err)
	}
	contact, err := database.GetContact(uint(id))
	if err != nil {
		return utils.ProcessError(err)
	}
	if err := contact.Delete(); err != nil {
		return utils.ProcessError(err)
	}
	return c.JSON(http.StatusOK, contact.GetResponseContact())
}
