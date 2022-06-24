CREATE TYPE colors AS ENUM (
    'red',
    'blue',
    'green'
);

-- Example queries for sqlc
CREATE TABLE boxs (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  color colors
);

-- name: GetBoxes :many
SELECT * FROM boxs
WHERE 
CASE
	WHEN sqlc.narg('color') :: colors != null THEN
		color = sqlc.narg('color')
	ELSE true
END;