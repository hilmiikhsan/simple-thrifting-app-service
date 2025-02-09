package helpers

import (
	"github.com/jmoiron/sqlx"
)

func SeedUserAccount(db *sqlx.DB) error {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE username = $1", "ikhsan")
	if err != nil {
		Logger.Error("Error checking if user exists:", err)
		return err
	}

	if count > 0 {
		Logger.Println("User already exists, skipping seed.")
		return nil
	}

	hashedPassword, err := HashPassword("admin123")
	if err != nil {
		Logger.Error("Error hashing the password:", err)
		return err
	}

	query := `
		INSERT INTO users (username, password, email, full_name, nim, phone_number)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	values := []interface{}{
		"ikhsan",
		hashedPassword,
		"ikhsanhilmimuhammad@gmail.com",
		"Muhammad Ikhsan Hilmi",
		"22552012003",
		"087785110345",
	}

	_, err = db.Exec(query, values...)
	if err != nil {
		Logger.Println("Error creating admin user:", err)
		return err
	}

	Logger.Println("Admin user created successfully.")
	return nil
}
