package product

const (
	queryInsertNewProduct = `
		INSERT INTO products
		(
			name,
			description,
			price,
			stock
		) VALUES (?, ?, ?, ?)
	`

	queryFindAllProduct = `
		SELECT
			id,
			name,
			description,
			price,
			stock
		FROM products
	`

	queryFindProductByID = `
		SELECT
			id,
			name,
			description,
			price,
			stock
		FROM products
		WHERE id = ?
	`
)
