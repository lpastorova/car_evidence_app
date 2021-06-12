package main

import "database/sql"

type Store interface {
	CreateCar(car *Car) error
	GetCar() ([]*Car, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateCar(car *Car) error {
	_, err := store.db.Query("INSERT INTO cars(spz, description) VALUES($1,$2)", car.Spz, car.Description)
	return err
}

func (store *dbStore) GetCar() ([]*Car, error) {
	rows, err := store.db.Query("SELECT spz, description FROM cars")

	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	cars := []*Car{}

	for rows.Next() {
		car := &Car{}
		if err := rows.Scan(&car.Spz, &car.Description); err != nil {
			return nil, err
		}
		cars = append(cars, car)

	}
	return cars, nil
}

var store Store

func InitStore(s Store) {
	store = s
}
