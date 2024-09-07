-- +goose Up
INSERT INTO ORGANIZATIONS (id, created_at, updated_at, name)
VALUES ('99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '2024-09-06 00:00:00', '2024-09-06 00:00:00', 'admin organization');
INSERT INTO USERS (id, created_at, updated_at, name, organization_id)
VALUES ('5faba39f-64fe-4805-8365-0a91bb396477', '2024-09-06 00:00:00', '2024-09-06 00:00:00', 'admin', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187');

-- +goose Down
DELETE FROM ORGANIZATIONS 
WHERE id = '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187';
DELETE FROM USERS
WHERE id = '5faba39f-64fe-4805-8365-0a91bb396477';