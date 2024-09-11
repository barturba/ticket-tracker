-- +goose Up
INSERT INTO companies (id, created_at, updated_at, name, organization_id)
VALUES ('5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7', '2024-09-06 00:00:00', '2024-09-06 00:00:00', 'Big Store Corp.', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187');

-- +goose Down
DELETE FROM companies
WHERE id = '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7';