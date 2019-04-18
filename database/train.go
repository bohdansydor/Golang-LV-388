package database

import (
	"team-project/services/data"

	"github.com/google/uuid"

	_ "github.com/lib/pq" // pq lib for using postgres
)

// ITrain is an interface for tests mock
type ITrain interface {
	GetAllTrains() ([]data.Train, error)
	GetTrain(string) (data.Train, error)
	AddTrain(data.Train) error
	UpdateTrain(uuid.UUID, string, string, string, string, string, string) error
	DeleteTrain(uuid.UUID) error
}

//GetAllTrains is a method which Gets Trains from table trains
func GetAllTrains() ([]data.Train, error) {
	rows, err := Db.Query("select * from public.trains")
	if err != nil {
		return nil, err
	}

	trains := []data.Train{}
	for rows.Next() {
		t := data.Train{}
		err := rows.Scan(&t.ID, &t.DepartureCity, &t.ArrivalCity, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate)
		if err != nil {
			trains = append(trains, t)
			continue
		}
		trains = append(trains, t)
	}
	return trains, nil
}

//GetTrain is a method which Gets Train from table trains
func GetTrain(id string) (data.Train, error) {
	idint, err := uuid.Parse(id)
	if err != nil {
		return data.Train{}, err
	}
	row := Db.QueryRow("select * from trains where id = $1", idint)
	t := data.Train{}
	err = row.Scan(&t.ID, &t.DepartureCity, &t.ArrivalCity, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate)
	if err != nil {
		return t, err
	}
	return t, nil
}

//AddTrain is a method which Adds Train to table trains
func AddTrain(t data.Train) error {
	_, err := Db.Exec("insert into trains (departure_city,arrival_city,departure_time,departure_date,arrival_time,arrival_date) values ($1,$2,$3,$4,$5,$6)", t.DepartureCity, t.ArrivalCity, t.DepartureTime, t.DepartureDate, t.ArrivalTime, t.ArrivalDate)

	if err != nil {
		return err
	}

	return nil
}

//UpdateTrain is a method which Updates Train in table trains
func UpdateTrain(id uuid.UUID, departureCity string, arrivalCity string, departureTime string, departureDate string, arrivalTime string, arrivalDate string) error {
	_, err := Db.Exec("update public.trains set departure_city = $1 , arrival_city = $2, departure_time = $3, departure_date = $4, arrival_time = $5, arrival_date = $6 where id = $7", departureCity, arrivalCity, departureTime, departureDate, arrivalTime, arrivalDate, id)

	if err != nil {
		return err
	}
	return nil
}

//DeleteTrain is a method which Deletes Train from table trains
func DeleteTrain(id uuid.UUID) error {
	_, err := Db.Exec("delete from trains where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

/*
//GetLastTrain is a method which returns last train from table trains
func GetLastTrain() (data.Train, error) {
	row := Db.QueryRow("select * from trains order by id desc limit 1")
	t := data.Train{}
	err := row.Scan(&t.ID, &t.DepartureCity, &t.ArrivalCity, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate)
	if err != nil {
		return t, err
	}
	return t, nil
}
*/
