-- +goose Up
INSERT INTO incidents (id, created_at, updated_at, short_description, organization_id, company_id, configuration_item_id)
VALUES ('34ea1efb-1892-40b7-8bc7-b484aba60d0e', '2024-09-09 11:48:50.915394', '2024-09-09 11:48:50.915394', 'router 3 host status', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7', '2c6fc3cd-aed3-431e-850c-7bb19ad05673');
INSERT INTO incidents (id, created_at, updated_at, short_description, description, organization_id, company_id, configuration_item_id)
VALUES ('f839d323-67ea-432d-ac8e-0a42774cc8bd', '2024-09-09 11:48:50.915394', '2024-09-09 11:48:50.915394', 'snmp settings', 'configure snmp for router 3', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7', '2c6fc3cd-aed3-431e-850c-7bb19ad05673');
INSERT INTO incidents (id, created_at, updated_at, short_description, description, organization_id, company_id, configuration_item_id)
VALUES ('85215904-4bea-42a7-9382-a12b8565558f', '2024-09-09 11:48:50.915394', '2024-09-09 11:48:50.915394', 'VPN tunnel not working', 'users cant access corporate resources on router 2', '99ae48fb-89e8-4f7b-bd4c-b947b8fe8187', '5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7', 'be275908-1c23-4eb2-aa0b-086beb7334b5');

-- +goose Down
DELETE FROM incidents
WHERE id = '34ea1efb-1892-40b7-8bc7-b484aba60d0e';
DELETE FROM incidents
WHERE id = 'f839d323-67ea-432d-ac8e-0a42774cc8bd';
DELETE FROM incidents
WHERE id = '85215904-4bea-42a7-9382-a12b8565558f';