package server

import (
	"context"
	"fmt"
	"time"

	"net/http"
	"strconv"

	"github.com/KozlovNikolai/CMDpatients/internal/model"
	"github.com/KozlovNikolai/CMDpatients/internal/store"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var patient model.Patient
	id, _ := strconv.Atoi(c.Param("id"))

	query := `
	SELECT patients.id
	FROM patients
	WHERE patients.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)

	err := row.Scan(&patient.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	query = `DELETE FROM patients WHERE id=$1`
	_, err = store.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}

func Create(c *gin.Context) {
	var patient model.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(patient)
	patient.CreatedAt = time.Now().UTC()
	fmt.Println(patient)
	query := `
		INSERT INTO patients (created_at,surname,name,patronymic,gender,birthday)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id`
	err := store.DB.QueryRow(context.Background(), query, patient.CreatedAt, patient.Surname, patient.Name,
		patient.Patronymic, patient.Gender, patient.Birthday).Scan(&patient.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var patient model.Patient
	query := `
		SELECT id,created_at,surname,name,patronymic,gender,birthday
		FROM patients
		WHERE patients.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(&patient.ID, &patient.CreatedAt, &patient.Surname, &patient.Name,
		&patient.Patronymic, &patient.Gender, &patient.Birthday)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func GetList(c *gin.Context) {
	var patients []model.Patient
	query := `
		SELECT id,created_at,surname,name,patronymic,gender,birthday
		FROM patients
		LIMIT 100`
	rows, err := store.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var patient model.Patient
		err := rows.Scan(&patient.ID, &patient.CreatedAt, &patient.Surname, &patient.Name,
			&patient.Patronymic, &patient.Gender, &patient.Birthday)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}
