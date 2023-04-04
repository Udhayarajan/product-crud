BEGIN;
ALTER TABLE products
    DROP COLUMN createdAt,
    DROP COLUMN updatedAt;

COMMIT;
