package services

import (
	"github.com/Hyatus/myapi/config"
	"github.com/Hyatus/myapi/models"
)


func GetAllUsers()([]models.User, error){
	rows, err := config.DB.Query("SELECT id, name, age, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}


func CreateUser(user models.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, age, email) VALUES (?, ?, ?)", user.Name, user.Age, user.Email)
	if err != nil {
		// Write log
		return err
	}
	return nil 
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, name, age, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Age, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(id int, user models.User) (mensaje string, err error) {
	resultado, err := config.DB.Exec("UPDATE users SET name = ?, age = ?, email = ? WHERE id = ?", user.Name, user.Age, user.Email, id)
	if err != nil {
		return "", err
	}
	rowsAffected, err := resultado.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowsAffected == 0 {
		return "Usuario no pudo ser actualizado - no Existe", nil
	}
	return "Usuario actualizado con éxito", nil
}

func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}



