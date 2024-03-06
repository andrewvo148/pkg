-- name: GetCompany :one
SELECT * FROM jobs.companies
WHERE id = $1 LIMIT 1;

-- name: ListCompanies :many
SELECT * FROM jobs.companies;

