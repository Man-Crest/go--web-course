package handlers

import (
	"log"
	"time"

	"github.com/Man-Crest/GO-Projects/01_bookings/pkg/models"
)

func (m *Repository) Availability(roomType *models.RoomType, desiredStartDate time.Time, desiredEndDate time.Time) ([]models.Rooms, error) {

	rows, err := m.DB.SQL.Query(`
	SELECT rooms.id, rooms.roomnumber, roomtype.name
	FROM rooms
	LEFT JOIN reservations ON rooms.id = reservations.room_id
	LEFT JOIN roomtype ON rooms.roomtypeid = roomtype.id
	WHERE reservations.room_id IS NULL
		OR NOT (
			($1 BETWEEN reservations.start_date AND reservations.end_date)
			OR ($2 BETWEEN reservations.start_date AND reservations.end_date)
			OR (reservations.start_date BETWEEN $1 AND $2)
			OR (reservations.end_date BETWEEN $1 AND $2)
			OR ($1 < reservations.start_date AND $2 > reservations.end_date)
		);`, desiredStartDate, desiredEndDate)

	if err != nil {
		log.Fatal(err, "after query parsing of search availability")
	}

	defer rows.Close()

	var rooms []models.Rooms

	// Iterate through the rows to fetch available rooms
	for rows.Next() {
		var room models.Rooms
		var RoomTypeName string
		err := rows.Scan(&room.ID, &room.RoomNumber, &RoomTypeName)

		if err != nil {
			log.Fatal(err, "after query parsing rows of search availability")
			return rooms, err
		}
		room.RoomType = &models.RoomType{Name: RoomTypeName}
		if room.RoomType.Name == roomType.Name {
			rooms = append(rooms, room)
		}
	}
	return rooms, nil
}
