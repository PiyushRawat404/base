package repository

import (
	"context"
	"errors"
	"fmt"
	"product/internal/models"
	"product/pkg/utils"
	"strings"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) CreateUser(ctx context.Context, email, passwordHash string) (*models.User, error) {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
		RETURNING id, email, password_hash, created_at
	`

	var user models.User
	err := r.DB.QueryRow(ctx, query, email, passwordHash).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`

	var user models.User
	err := r.DB.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateProduct(ctx context.Context, product models.Product) (*models.Product, error) {
	query := `
		INSERT INTO products (name, description, price, quantity)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, description, price, quantity, created_at
	`

	var created models.Product
	err := r.DB.QueryRow(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
	).Scan(
		&created.ID,
		&created.Name,
		&created.Description,
		&created.Price,
		&created.Quantity,
		&created.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *Repository) GetProducts(ctx context.Context, pagination utils.PaginationParams, filter utils.ProductFilter) ([]models.Product, int, error) {
	conditions := []string{"1=1"}
	args := []any{}

	if filter.Search != "" {
		args = append(args, "%"+filter.Search+"%")
		conditions = append(conditions, fmt.Sprintf("(name ILIKE $%d OR description ILIKE $%d)", len(args), len(args)))
	}

	whereClause := strings.Join(conditions, " AND ")

	countQuery := "SELECT COUNT(*) FROM products WHERE " + whereClause
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	args = append(args, pagination.Limit, pagination.Offset)
	listQuery := fmt.Sprintf(`
		SELECT id, name, description, price, quantity, created_at
		FROM products
		WHERE %s
		ORDER BY %s %s
		LIMIT $%d OFFSET $%d
	`, whereClause, filter.SortField, filter.SortOrder, len(args)-1, len(args))

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		products = append(products, product)
	}

	return products, total, rows.Err()
}

func (r *Repository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	query := `
		SELECT id, name, description, price, quantity, created_at
		FROM products
		WHERE id = $1
	`

	var product models.Product
	err := r.DB.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (r *Repository) UpdateProduct(ctx context.Context, id int, product models.Product) (*models.Product, error) {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, quantity = $4
		WHERE id = $5
		RETURNING id, name, description, price, quantity, created_at
	`

	var updated models.Product
	err := r.DB.QueryRow(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		id,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Description,
		&updated.Price,
		&updated.Quantity,
		&updated.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &updated, nil
}

func (r *Repository) DeleteProduct(ctx context.Context, id int) (bool, error) {
	commandTag, err := r.DB.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	return commandTag.RowsAffected() > 0, nil
}
