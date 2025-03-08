package infrastructure

import (
	"fmt"

	core "api-hexagonal-cars/src/Core"
	"api-hexagonal-cars/src/rents/domain"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateRent(rent domain.Rent) (uint, error) {
	query := "INSERT INTO rents (id_customer, id_car, return_date) VALUES (?, ?, ?)"

	res, err := mysql.conn.ExecutePreparedQuery(query, rent.Id_Customer, rent.Id_Car, rent.Return_date_rent)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()
	return uint(rowsAffected), nil
}

func (mysql *MySQL) GetAllRents() []domain.Rent {
	query := "SELECT * FROM rents"
	var rents []domain.Rent

	rows := mysql.conn.FetchRows(query)
	if rows == nil {
		fmt.Println("No se pudieron obtener los datos.")
		return rents
	}

defer rows.Close()

	for rows.Next() {
		var r domain.Rent
		rows.Scan(&r.Id_Customer, &r.Id_Car, &r.Return_date_rent)
		rents = append(rents, r)
	}

	return rents
}

func (mysql *MySQL) DeleteRent(id_customer int, id_car int) (uint, error) {
	query := "DELETE FROM rents WHERE id_customer = ? AND id_car = ?"

	res, err := mysql.conn.ExecutePreparedQuery(query, id_customer, id_car)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	rows, _ := res.RowsAffected()
	return uint(rows), nil
}
