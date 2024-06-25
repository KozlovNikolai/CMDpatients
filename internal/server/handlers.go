package server

import (
	"context"
	"fmt"
	"gin-pg-crud/internal/model"
	"gin-pg-crud/internal/store"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmployer(c *gin.Context) {
	var employer model.Employer
	var person model.Person
	if err := c.ShouldBindJSON(&employer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("employer=%v\n", employer)
	query := `
	SELECT persons.id,persons.name,persons.email
	FROM persons
	WHERE persons.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, employer.Person.ID)
	err := row.Scan(&person.ID, &person.Name, &person.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	employer.Person.Email = person.Email
	employer.Person.Name = person.Name
	query = `
		INSERT INTO employers (company,person_id)
		VALUES ($1,$2)
		RETURNING id`
	err = store.DB.QueryRow(context.Background(), query, employer.Company, employer.Person.ID).Scan(&employer.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, employer)
}

func GetEmployer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employer model.Employer
	query := `
		SELECT employers.id, employers.company, persons.id,persons.name,persons.email
		FROM employers
		JOIN persons ON employers.person_id = persons.id
		WHERE employers.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(&employer.ID, &employer.Company, &employer.Person.ID, &employer.Person.Name, &employer.Person.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employer not found"})
		return
	}
	c.JSON(http.StatusOK, employer)
}

func GetEmployerList(c *gin.Context) {
	var employers []model.Employer
	query := `
		SELECT employers.id, employers.company, persons.id,persons.name,persons.email
		FROM employers
		JOIN persons ON employers.person_id = persons.id
		LIMIT 100`
	rows, err := store.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var employer model.Employer
		err := rows.Scan(&employer.ID, &employer.Company, &employer.Person.ID, &employer.Person.Name, &employer.Person.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		employers = append(employers, employer)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employers)
}

func UpdateEmployer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employer model.Employer
	if err := c.ShouldBindJSON(&employer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	query := `
		UPDATE employers
		SET company=$1,name=$2,email=$3
		WHERE id=$4`
	_, err := store.DB.Exec(context.Background(), query, employer.Company, employer.Person.Name, employer.Person.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Update successfully"})
}

func DeleteEmployer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	query := `DELETE FROM employers WHERE id=$1`
	_, err := store.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}

func DeletePerson(c *gin.Context) {
	var person model.Person
	id, _ := strconv.Atoi(c.Param("id"))

	query := `
	SELECT persons.id,persons.name,persons.email
	FROM persons
	WHERE persons.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(&person.ID, &person.Name, &person.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}

	query = `DELETE FROM persons WHERE id=$1`
	_, err = store.DB.Exec(context.Background(), query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}

func CreatePerson(c *gin.Context) {
	var person model.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := `
		INSERT INTO persons (name,email)
		VALUES ($1,$2)
		RETURNING id`
	err := store.DB.QueryRow(context.Background(), query, person.Name, person.Email).Scan(&person.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, person)
}

func GetPerson(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var person model.Person
	query := `
		SELECT persons.id,persons.name,persons.email
		FROM persons
		WHERE persons.id=$1 LIMIT 100`
	row := store.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(&person.ID, &person.Name, &person.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

func GetPersonList(c *gin.Context) {
	var persons []model.Person
	query := `
		SELECT persons.id,persons.name,persons.email
		FROM persons
		LIMIT 100`
	rows, err := store.DB.Query(context.Background(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var person model.Person
		err := rows.Scan(&person.ID, &person.Name, &person.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, person)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, persons)
}
