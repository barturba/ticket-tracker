-- +goose Up
INSERT INTO incidents_sequence (id, created_at, updated_at, organization_id, sequence)
VALUES ('8bdb7562-1e79-4dd8-85f1-34693e28215a', '2024-09-06 00:00:00', '2024-09-06 00:00:00', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', 0);

-- +goose Down
DELETE FROM incidents_sequence
WHERE id = '8bdb7562-1e79-4dd8-85f1-34693e28215a';