package user

const (
	queryFindUserByUsername = `
			SELECT 
				id, 
				username,
				password,
				full_name
			FROM users WHERE username = ?
	`

	queryFindUserByID = `
			SELECT
				id,
				username,
				full_name,
				email,
				nim,
				phone_number
			FROM users 
			WHERE id = ?
	`
)
