-- +goose Up
INSERT INTO configuration_items (id, created_at, updated_at, name, organization_id, company_id)
VALUES ('2c6fc3cd-aed3-431e-850c-7bb19ad05673', '2024-09-09 11:48:50.915394', '2024-09-09 11:48:50.915394', 'router 3', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7');
INSERT INTO configuration_items (id, created_at, updated_at, name, organization_id, company_id)
VALUES ('527379b7-546d-4228-b658-4666da6c9793', '2024-09-09 11:48:50.915394', '2024-09-09 11:48:50.915394', 'router 2', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7');
INSERT INTO configuration_items (id, created_at, updated_at, name, organization_id, company_id)
VALUES ('be275908-1c23-4eb2-aa0b-086beb7334b5', '2024-09-09 11:48:50.915394', '2024-09-09 11:48:50.915394', 'router 1', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7');

-- +goose Down
DELETE FROM configuration_items
WHERE id = '2c6fc3cd-aed3-431e-850c-7bb19ad05673';
DELETE FROM configuration_items
WHERE id = '527379b7-546d-4228-b658-4666da6c9793';
DELETE FROM configuration_items
WHERE id = 'be275908-1c23-4eb2-aa0b-086beb7334b5';