package repositories

import (
	"goapi2/cmd/models"
	"goapi2/cmd/storage"
	"time"
)

func CreateMeasurements(measurements models.Measurements) (models.Measurements, error) {
	db := storage.GetDB()

	measurements.CreatedAt = time.Now()

	sqlStatement := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := db.QueryRow(sqlStatement, measurements.UserId, measurements.Weight, measurements.Height, measurements.BodyFat, measurements.CreatedAt).Scan(&measurements.Id)

	if err != nil {
		return measurements, err
	}
	return measurements, nil

}

func UpdateMeasurement(meansurements models.Measurements, id int) (models.Measurements, error) {

	db := storage.GetDB()

	sqlStatement := `UPDATE measurements SET weight = $1, height = $2, body_fat = $3, created_at = $4 WHERE id = $5 RETURNING id`

	err := db.QueryRow(sqlStatement, meansurements.Weight, meansurements.Height, meansurements.BodyFat, time.Now(), id).Scan(&id)

	if err != nil {
		return meansurements, err
	}
	meansurements.Id = id
	return meansurements, nil

}
