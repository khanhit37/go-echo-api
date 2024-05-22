package repositories

import (
	"goapi2/cmd/models"
	"goapi2/cmd/storage"
	"time"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)

	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()

	user.UpdatedAt = time.Now()

	sqlStatement := `UPDATE users
	 SET name = $1, email = $2, password = $3, updated_at = $4 WHERE id = $5`

	_, err := db.Exec(sqlStatement, user.Name, user.Email, user.Password, user.UpdatedAt, id)

	if err != nil {
		return user, err
	}
	return user, nil

}

func Profile(user models.User, id int) (models.User, error) {

	db := storage.GetDB()

	sqlStatement := `SELECT * FROM users WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return user, err
	}
	return user, nil
}
