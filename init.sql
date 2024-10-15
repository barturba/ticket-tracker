--
-- PostgreSQL database dump
--

-- Dumped from database version 15.8 (Homebrew)
-- Dumped by pg_dump version 16.3

-- Started on 2024-10-15 12:04:08 MDT

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


CREATE TYPE public.state_enum AS ENUM (
    'New',
    'In Progress',
    'Assigned',
    'On Hold',
    'Resolved'
);



SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE public.companies (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name text NOT NULL
);



CREATE TABLE public.configuration_items (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    name text NOT NULL
);



CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);




CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;



CREATE TABLE public.incidents (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    short_description text NOT NULL,
    description text,
    configuration_item_id uuid NOT NULL,
    company_id uuid NOT NULL,
    state public.state_enum DEFAULT 'New'::public.state_enum NOT NULL,
    assigned_to uuid
);




CREATE TABLE public.users (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    first_name character varying(50),
    last_name character varying(50),
    email text NOT NULL
);




ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);



COPY public.companies (id, created_at, updated_at, name) FROM stdin;
5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	2024-09-06 00:00:00	2024-09-06 00:00:00	Big Store Corp.
a988815c-3672-d0fe-1865-1c97fbc29117	2023-12-31 07:52:21	2024-02-28 19:15:12	Smart Home Technologies
aad132b5-3f63-c9a6-3370-8d6ca03819b9	2024-08-19 06:31:58	2024-03-29 20:11:13	Green Energy Solutions
b7120f97-9420-169f-ae88-0e15f90610e8	2023-10-16 01:51:45	2024-07-08 13:23:46	CloudData Systems
a2767907-c217-2ee4-9727-177ac526ac73	2024-02-22 17:21:44	2024-08-29 03:39:03	Stellar Analytics
835855b7-48e2-194e-6b7e-436075b012a7	2024-01-07 02:39:33	2024-08-11 19:26:39	Stellar Analytics
1baf574f-287f-f77b-dd8a-9136d0ef5158	2024-04-19 19:44:22	2023-12-11 01:21:48	Elite Venture Capital
bdc094bb-5267-aba4-f0a8-6dd68864f37f	2024-06-02 01:03:56	2024-03-07 12:43:16	Global Networks Ltd.
1f596437-be02-5a5e-32f3-5ab5b1cff9f6	2023-11-21 15:33:53	2024-01-20 11:08:42	Precision Manufacturing Co.
6ce4068b-b8ef-6410-e92e-9ea044daacee	2024-02-09 03:58:50	2024-07-29 21:21:45	CloudData Systems
a9b46d58-b3f7-f518-d524-d22b3c63a2f9	2024-04-21 06:15:57	2024-04-28 09:53:25	CloudData Systems
1be71338-22dd-3d93-9e22-f15166473052	2023-12-10 12:04:54	2024-03-22 18:47:15	Green Energy Solutions
98f5fdc9-3cef-0568-081e-e56d440b937a	2024-05-17 12:22:18	2023-11-25 00:14:14	Urban Development Group
fee4aa00-ddc2-ce5a-2ded-56748c7bb575	2024-08-25 23:43:37	2024-02-17 23:43:19	Eco-Friendly Products LLC
7d826631-556d-f399-460c-7abe70f39f52	2024-09-18 20:19:59	2024-03-27 12:47:37	Smart Home Technologies
228c7c2d-3109-cdf6-38f7-d2603d34810c	2023-10-29 14:18:42	2024-05-12 12:08:57	Smart Home Technologies
e55ffd34-5054-b617-c1bf-097c08ae4c89	2024-07-09 20:26:11	2023-12-18 00:14:51	Smart Home Technologies
0eac1ef4-3188-a9c4-fc23-659d7fe784cc	2024-08-10 09:35:17	2023-12-31 05:01:43	Eco-Friendly Products LLC
c633399a-6fe8-39f4-a142-9ed39ebfe4f2	2024-03-02 21:49:55	2024-03-28 14:54:21	CloudData Systems
57d0be3e-398c-aadf-2cd8-23684814b839	2024-04-14 20:58:46	2023-10-25 00:24:40	Green Energy Solutions
d6fc7468-73b3-5ec1-bc7a-dad2ea2a3ef3	2024-03-14 08:06:07	2024-08-01 18:39:16	Green Energy Solutions
9d85c6b0-4d02-21fc-33d9-d6bf335e624b	2024-07-05 20:41:31	2023-11-18 10:21:44	Precision Manufacturing Co.
515ef09a-9de7-4187-7ffd-7432012f2c8a	2024-04-05 22:19:27	2024-09-06 03:29:59	Stellar Analytics
669589b7-1180-c01e-687d-fff150709e47	2024-01-11 00:08:53	2024-04-28 05:33:42	Global Networks Ltd.
833584fc-8c0a-a1cb-db28-355c6bd78082	2024-02-24 11:46:09	2023-12-13 13:18:01	Urban Development Group
ff794092-d8db-236f-9c6f-699d9ca6c3b6	2023-11-07 02:25:39	2024-01-24 20:04:00	Eco-Friendly Products LLC
bf859713-b8da-e626-a30f-2c9dd476d6b9	2024-04-30 19:24:53	2023-11-27 10:09:38	CloudData Systems
1f760bdf-31d7-14f6-edee-c2b9a7a6e5fa	2023-10-18 11:01:13	2024-04-28 14:33:40	Stellar Analytics
bea844ba-8e95-c664-737b-8f9abbac9285	2023-12-30 11:22:58	2024-09-26 02:34:00	Elite Venture Capital
707ba700-72e7-02ef-a2fb-351bb60de202	2023-10-19 00:10:20	2024-07-09 01:03:58	Elite Venture Capital
d99927dc-2c3d-57e6-f982-3de64c404104	2024-02-14 13:29:41	2024-06-14 20:24:15	Smart Home Technologies
eb957630-7f6a-931e-5c5c-2a5edb333ad7	2024-08-09 10:23:09	2023-11-01 22:43:41	Green Energy Solutions
b88498f2-759d-b568-3b37-842aebcba416	2024-01-30 01:40:06	2024-07-20 21:03:22	Smart Home Technologies
05edfb15-7467-fdd6-eac2-8391b261f992	2023-10-29 19:30:23	2024-05-20 17:09:14	Green Energy Solutions
440c68ba-e930-65f2-7013-b7f0f8c34324	2024-08-07 15:34:27	2024-04-03 19:17:10	Eco-Friendly Products LLC
534df3a2-fbc9-cc82-e64c-51c090c12b29	2024-06-03 18:42:02	2024-07-27 20:18:11	Eco-Friendly Products LLC
e05dc6d7-9cb7-df7a-f9bf-ce9834b0e41b	2024-01-03 17:03:39	2024-04-22 01:39:22	Eco-Friendly Products LLC
5996de67-6cf6-4f54-6438-f94a2fae9e91	2024-01-03 13:08:21	2024-01-28 08:17:13	Precision Manufacturing Co.
163ed804-a034-dd76-5087-26815d37e19d	2024-02-25 00:09:14	2023-10-16 23:27:15	Precision Manufacturing Co.
6922cc8b-cc03-9647-9854-d478746cbc33	2024-03-07 13:07:02	2024-08-25 14:30:17	Stellar Analytics
00b73e30-250e-5127-059a-285085d02814	2024-06-01 18:57:39	2023-10-25 08:06:14	Tech Innovations Inc.
b92ed5ea-5029-9fad-daad-1d25f985db12	2024-03-12 09:05:36	2024-05-14 23:45:45	Stellar Analytics
f9af1153-8c7f-db94-7c5a-e9a0b3cf94fa	2023-11-01 16:24:55	2024-07-11 07:54:03	Smart Home Technologies
722a2243-b5ee-fba0-2400-95b939f87392	2024-06-13 15:03:22	2024-08-11 07:13:39	Global Networks Ltd.
7440373d-5774-2a6a-2451-2f5d4032de26	2023-11-14 23:52:48	2023-10-25 17:41:48	Green Energy Solutions
00bb1e5c-fcba-c1bf-e4cd-86a6c9186310	2024-06-17 07:13:42	2024-05-28 07:21:48	Global Networks Ltd.
faf5c017-5cfc-8135-6e95-7f02e6fcb6e6	2024-04-30 05:41:01	2024-07-31 16:55:13	Smart Home Technologies
eae20175-f5fd-8614-b8af-360f15350068	2024-04-24 19:26:54	2024-07-04 21:03:11	Tech Innovations Inc.
093d9075-bdf8-239a-9eb5-413f3716d6d5	2023-12-10 23:25:27	2024-03-28 20:24:01	CloudData Systems
047bdb8c-0de3-49b8-3cb8-cf89933f52f8	2024-01-05 22:05:22	2024-02-16 01:46:22	Global Networks Ltd.
28ce2c52-0d5a-99fc-7bb5-bff1c3a288c8	2024-09-21 08:59:43	2024-06-15 15:39:34	Global Networks Ltd.
0fcb4b25-d3a6-f6ca-9507-b5e9e8f7438e	2023-10-20 09:47:22	2024-01-17 20:02:32	Stellar Analytics
e19075fe-14fc-1780-401b-3ff2146fc7a3	2024-03-18 21:17:09	2024-06-20 12:44:46	Urban Development Group
3a4ffaf2-77ac-06c4-e5b7-806c5dd98a9c	2024-01-28 11:53:46	2024-08-20 13:46:16	CloudData Systems
61af9a72-0035-8820-49cf-48e1fd0efe91	2023-10-28 23:19:09	2024-02-23 06:54:22	Global Networks Ltd.
a9835ad6-123b-5bd9-d664-f18d258b64e9	2024-06-30 17:36:37	2024-09-30 05:43:16	Elite Venture Capital
4897dc7d-a3b3-c70d-d854-f3b88ea3cc90	2024-03-10 08:51:25	2023-10-19 06:09:12	Precision Manufacturing Co.
b4020a50-cba1-0bb0-9218-07c5aec9fe2d	2023-11-21 20:19:34	2024-05-13 16:20:27	Elite Venture Capital
0ae21192-df8a-f334-83f5-6bb154dfab36	2024-03-08 14:29:39	2023-12-26 19:33:50	Precision Manufacturing Co.
a48aea8e-0278-00d4-3fe0-23933ea4554f	2023-10-31 06:43:23	2024-07-20 02:40:53	Global Networks Ltd.
f5387816-9b47-083f-90c9-8ff3e709ec1e	2024-09-09 08:06:34	2024-02-02 14:49:10	Global Networks Ltd.
1b43cc73-e173-88b5-8ee3-ce27ab1f7c62	2024-09-29 05:15:42	2023-10-19 04:43:07	Stellar Analytics
49e072b6-94ab-0195-b106-351c0f199eec	2024-02-26 10:06:39	2024-07-29 21:38:56	Eco-Friendly Products LLC
2ce274dd-6e2e-0c78-6258-cd8e35ec74dd	2024-02-12 23:22:29	2024-03-06 21:57:07	Eco-Friendly Products LLC
7cf856dd-308b-b713-e076-570aa6c09967	2024-01-29 20:06:34	2024-07-05 22:40:22	Eco-Friendly Products LLC
63c2ef03-768d-7b04-bedc-159dff0e2083	2024-09-22 05:37:38	2024-09-08 05:39:19	Global Networks Ltd.
db71629a-0f48-27c6-bb94-025ec4e96335	2024-10-02 10:29:54	2024-04-14 12:32:53	Green Energy Solutions
3bfc1bfd-c79b-a2c7-35dd-d0b9c04254bb	2024-09-30 10:48:42	2024-04-29 17:41:55	Eco-Friendly Products LLC
5008195d-18db-a8e1-23f3-f1199765612f	2024-06-20 15:27:30	2024-03-26 14:29:08	Urban Development Group
4932fd4e-d82f-7eef-4b21-f46a9f1c8c6c	2024-07-24 23:01:25	2023-12-04 09:21:28	Tech Innovations Inc.
c0765987-5196-0196-6a0b-b973e99924f8	2024-07-30 23:27:55	2024-04-02 01:31:57	Smart Home Technologies
3c632c41-5e27-7583-cc1e-8e12608342a7	2023-12-09 10:39:58	2024-09-01 15:43:03	Tech Innovations Inc.
aa474f30-3409-a176-bbc3-60bf568124e9	2024-03-03 21:33:04	2024-04-21 02:36:20	Precision Manufacturing Co.
68d81f33-1996-150a-9882-027be8497b61	2024-07-09 00:28:35	2024-02-17 10:24:02	Smart Home Technologies
69011f7d-ab78-d570-6ef9-9a9f858be2da	2024-01-24 12:05:56	2024-07-07 19:28:07	Urban Development Group
11bc4848-635e-2d05-9387-e20e09fe62bc	2024-08-10 16:57:03	2024-08-22 23:51:01	Elite Venture Capital
f13b38a9-9c31-ca29-b6fe-569fbd3049b3	2023-12-20 08:39:57	2023-12-10 00:04:26	Green Energy Solutions
033ba238-5996-57c3-7fb5-322b8c8f10c3	2024-02-01 14:51:57	2024-03-20 06:40:41	Eco-Friendly Products LLC
af1317ac-70c3-6b01-f782-b273ed2fc776	2023-10-21 22:22:24	2024-07-23 09:18:17	Urban Development Group
88ece235-12ad-1182-2b71-3293dbc4ff0a	2023-10-30 17:27:27	2024-02-01 07:42:29	Smart Home Technologies
15f062d0-ba85-98f0-f18a-1e970282bd78	2024-01-24 02:11:38	2024-06-30 08:28:12	Urban Development Group
bfeff151-c385-82bc-90c2-b4b259839383	2024-07-30 00:31:08	2024-09-14 17:23:02	Tech Innovations Inc.
4e6f22c6-2486-4d38-d4f4-599776479d63	2023-10-23 01:59:46	2024-07-29 06:21:58	Tech Innovations Inc.
31fd0515-b0f6-ba9c-34d4-36201e38150d	2023-10-09 12:21:04	2024-06-30 02:18:00	Tech Innovations Inc.
8a070f84-5bd5-15f8-2a45-da392d01c2b4	2023-11-20 12:32:22	2024-02-21 00:02:47	Stellar Analytics
a22367d3-7144-4ebb-7cc8-316125e3cf95	2024-02-04 10:17:27	2024-07-27 22:51:59	Smart Home Technologies
5decb9ac-3779-67e0-11e9-02ae1e672781	2024-01-10 19:02:08	2024-10-01 08:37:18	Elite Venture Capital
effe6dce-661f-cbcd-bf91-1b30027372fc	2023-11-06 16:41:58	2023-12-11 18:04:32	Green Energy Solutions
37559973-9f94-5ca3-ff3b-454c9b765377	2024-02-16 15:07:13	2024-09-24 05:41:22	Smart Home Technologies
84921767-114e-2a99-49c5-d13e8f77ec27	2024-06-12 23:14:16	2024-05-24 23:13:38	Eco-Friendly Products LLC
536e611f-5dd1-ad49-b470-e850a1ec141e	2024-05-14 07:53:29	2024-06-06 03:59:15	Eco-Friendly Products LLC
6f115cdc-1a43-1f12-feff-57015675dbec	2024-05-02 17:45:41	2024-08-15 10:57:58	Smart Home Technologies
c591d0ba-ad44-3a2b-0863-2b62614308e4	2024-04-02 09:31:46	2024-08-17 04:28:04	Urban Development Group
72dff6a3-ba73-f705-5607-0356aa8fa23b	2024-06-01 14:41:41	2024-01-01 23:50:57	CloudData Systems
44009247-c122-b704-8e5e-10b8624eee73	2024-05-25 20:13:27	2024-09-10 09:54:59	Eco-Friendly Products LLC
2bc82373-8c8e-e66e-b168-c7699aa95393	2024-08-23 07:16:03	2023-12-23 08:03:42	Green Energy Solutions
10086067-ab04-e4e5-2e3d-a6490e0afc32	2023-10-18 08:40:06	2024-02-13 15:48:20	Elite Venture Capital
71294002-4c4a-963f-a6eb-ca4f713bf6f8	2024-02-26 09:10:09	2024-05-30 16:04:05	Precision Manufacturing Co.
df9e9feb-a48b-0b68-2e13-43692562236e	2023-12-13 05:02:59	2024-03-02 05:38:41	Stellar Analytics
17275e14-a533-f0a2-4ef7-2e47e87dcc26	2024-02-02 14:37:16	2024-07-31 13:40:29	Stellar Analytics
f708550a-0fe1-4ee9-b27e-ed5db213ffff	2024-10-12 13:34:22.700565	2024-10-13 23:06:04.101379	CORP 2.5
59bb140c-94d1-a296-ee4f-5bf5b86516dc	2024-02-05 23:26:57	2024-10-13 23:07:20.2955	Eco-Friendly Products LLC.
71e8f8b2-6349-4973-9e1c-7d9c44931af5	2024-10-13 23:11:11.434246	2024-10-14 10:50:44.711652	NEWCORP 2
b3eef82e-5c3c-4876-95be-a8f6acb0cd74	2024-10-14 13:00:40.195083	2024-10-14 13:00:40.195083	New Company
\.



COPY public.configuration_items (id, created_at, updated_at, name) FROM stdin;
2c6fc3cd-aed3-431e-850c-7bb19ad05673	2024-09-09 11:48:50.915394	2024-09-09 11:48:50.915394	router 3
527379b7-546d-4228-b658-4666da6c9793	2024-09-09 11:48:50.915394	2024-09-09 11:48:50.915394	router 2
be275908-1c23-4eb2-aa0b-086beb7334b5	2024-09-09 11:48:50.915394	2024-09-09 11:48:50.915394	router 1
9917195c-50fc-a451-273a-68687a99c542	2024-07-09 15:22:15	2024-01-04 18:50:27	NetSwitch_01_A1
cab16fbe-a515-cb4d-d658-fb73b4725d42	2024-02-18 08:49:08	2024-03-02 02:26:33	FireWall_LB02_X1
9b3313d5-074f-689b-f731-1115d390aeca	2023-12-09 07:09:22	2024-04-22 20:15:47	DataNode_Cluster_Alpha
c79b737e-c1f1-a4ca-9130-aa83b1f67f6a	2024-08-06 04:28:25	2024-03-08 21:16:30	AppServer_Beta_Prod_003
535b45a8-bea7-0cc5-85c0-9a61104587f5	2024-08-07 19:50:02	2024-02-18 08:48:11	DBServer_MainWest01
33de396b-1809-3649-4735-07b69d8491af	2024-09-04 21:13:14	2024-06-23 07:52:56	LoadBalancer_Network_01
39ac2e91-d7dc-99ab-504e-810640400761	2024-03-15 07:47:17	2023-11-13 21:52:51	VM_Compute_Node_East
a1f0fc17-d38b-3b75-1696-90c33c93e75d	2024-06-29 22:43:20	2024-07-07 10:30:37	StorageArray_Vault_07
8f7b1456-6da2-5830-72af-6c1b26c711c7	2023-11-22 23:49:36	2024-04-27 04:32:02	CoreRouter_Central_03
03ec8d05-75ce-0dac-2967-e37495283128	2024-02-02 10:11:01	2023-10-21 10:15:02	BackupServer_Pri_R01
81f3ffae-0d75-3321-8923-233568580371	2024-07-29 17:46:18	2024-02-13 13:26:59	DNS_Server_02_West
5e3a4b53-751b-4d05-d323-94c1aff0bb49	2024-01-23 13:00:14	2024-02-07 11:47:20	EmailGateway_HQ_Alpha
5256d313-4a93-4044-a613-701a7b42937f	2024-05-18 11:56:20	2024-08-12 01:11:11	FileServer_Archive_04
9aae1265-9e63-50a7-b819-06655a2099df	2024-08-09 04:54:54	2024-09-06 00:09:19	WebApp_Node1_Bravo
35954fb6-5948-6cc0-fd38-406318851b66	2024-09-10 09:17:49	2024-02-18 07:24:36	AuthServer_AD_Master01
af02e9ad-7811-f56e-0788-e1fc8d753071	2024-09-01 14:28:13	2024-01-11 21:42:19	DevServer_QA_Test_09
e6992e6a-70fe-057a-daa5-657b720e7324	2024-09-30 02:14:04	2024-07-11 10:35:27	Switch_CoreLayer_02
9f6d01dc-2925-5a7f-3970-bfc53a5e92e7	2024-05-04 23:35:12	2024-08-10 03:59:40	ProxyServer_RegionX_001
9c458e9c-68c3-9c0d-b4bd-2b3ff9fd2b4b	2023-11-30 12:53:45	2023-12-27 22:47:07	Syslog_Server_Security_07
170a4504-5812-37bb-2875-9c13958634a7	2023-10-15 16:03:12	2024-04-25 16:47:19	VPN_Access_Gateway_03
77125244-6c08-97a8-3d6a-e825dc6ddefc	2024-03-14 21:21:54	2024-04-14 05:38:15	AppCluster_Node_Delta_04
ad46169b-6ac8-da62-04d6-83cd5057073c	2024-03-24 12:37:01	2023-11-18 16:43:01	API_Gateway_West_ZoneA
cffd5f8f-a447-12d2-0d4d-b5dd9d5eaf09	2024-01-31 22:38:22	2024-09-06 05:07:49	NAS_Storage_Unit_12
ab14f2d8-aa2f-f414-35e7-c30a04b282a3	2024-07-12 13:35:52	2023-12-31 18:02:29	Hypervisor_AlphaCluster_08
45493155-abc6-4547-4ed1-57209ddd08c9	2024-09-30 21:47:18	2024-04-09 12:11:19	DHCP_Server_06_West
78778af7-44fe-eadc-991f-867056c8af93	2024-06-29 13:00:12	2024-09-25 14:26:52	MainRouter_DC_02
5bd9a8de-6af0-17d9-f6d3-3f6d8b1cf402	2024-04-06 20:17:28	2024-02-22 21:48:23	DNS_Node_Site03
7f5d0b1d-78ba-6463-0d82-67273d1cd09b	2024-07-24 21:01:58	2024-09-03 16:26:14	Application_DB_Node2
b0507f00-6703-87e4-08ca-016852d86c06	2023-11-17 13:12:53	2024-09-20 20:15:47	Email_Archive_Server_09
a060f7cd-dcec-a316-19af-c17e39405946	2024-08-23 04:48:09	2024-08-19 08:06:09	SecurityFirewall_Cluster_05
70a81c90-0c24-d3bd-c8a9-8d35a9367115	2024-06-02 14:00:08	2024-03-11 03:56:46	BackupStorage_Vault_01
3c802128-3f54-a54b-7eb5-a4097fedc549	2023-10-31 05:11:05	2024-02-05 09:48:03	DBCluster_MainProd_05
54e7fc2d-bd7c-1d09-0d0b-699f724514f6	2024-04-11 03:47:04	2024-07-28 01:03:43	WebProxy_RegionY_002
eb0b5ed3-4864-6a9c-be5f-47a6cb33aa68	2024-04-09 07:11:20	2024-08-15 03:07:33	VDI_Server_Beta_06
30c30ce2-0245-e36c-3ec5-3345e9e34e2f	2024-02-29 17:46:43	2024-01-28 10:17:02	Directory_Server_DC_01
3087ad41-1bfd-8b8a-375a-7cd4e0a4d024	2024-06-06 23:47:16	2024-06-11 20:50:59	NTP_TimeSource_03
92f225a5-94e1-9a62-99d1-33227da0d19c	2024-01-15 12:56:14	2024-06-12 11:35:54	MainframeNode_XY_09
b72080f1-b027-131b-9b00-3d23be8de8fe	2023-12-11 06:10:49	2024-02-27 13:46:52	VirtualHost_BravoCluster_04
70ba7328-cc3a-28f0-7f13-3268f4f1bec2	2024-05-17 17:08:23	2024-06-28 23:48:08	PrintServer_DeptA_01
0486721c-2a98-8779-c7c6-8f0d317834f2	2024-04-07 23:28:00	2024-03-23 11:49:21	WebServer_Pool01_Node03
446a76bd-61a7-fe5e-cd4e-ff23ef5c8af3	2024-01-21 12:34:20	2023-10-13 03:46:06	SyslogCollector_Core_02
e8931686-7088-57b1-af6d-780b578f2b1e	2024-01-05 06:35:25	2024-05-01 06:04:43	ManagementServer_Admin01
e2ff0f90-be7c-125d-ede2-f50154f3822b	2023-10-29 11:27:01	2024-09-25 04:07:55	DB_Replica_AlphaWest_08
fd6d801e-9fd3-0100-c455-ef5c21daf273	2024-02-13 23:49:59	2024-06-24 23:28:26	Firewall_Perimeter_North_02
40686e91-4b65-083a-9050-afcb65486263	2023-11-15 09:11:35	2024-02-20 15:12:06	ERP_Server_Bravo_05
4e9e9b13-10b7-4f72-a519-7de59835ce60	2024-08-16 07:26:23	2024-04-04 09:00:47	DataLake_Store_Zone3
0ab704b5-102d-2472-4a6c-9c2b62b7382a	2023-11-28 04:52:07	2023-12-30 03:38:43	DataCenter_Access_05
d798e72c-f28d-2717-f8d0-03eaf24caee1	2024-03-27 17:40:59	2024-04-25 12:29:55	NAS_Device_HR_02
199e674b-98fa-1441-2c8d-b91bcf25133f	2024-02-24 20:15:51	2024-04-24 19:40:16	VirtualAppNode_RegionB_07
2660389a-9b80-9c35-7426-edda56015dc1	2024-04-08 13:35:09	2024-07-17 15:38:47	ContentServer_Media_01
aa6f70c3-5acf-9fc1-8df4-5cbea21d3a62	2023-12-07 23:35:33	2024-08-31 03:56:13	LDAP_AuthNode_04
3ca9fa80-86b0-5397-8851-da011404252b	2023-12-16 03:58:08	2023-11-13 12:46:44	LicenseServer_Prod_03
602ce54a-d71f-c1e1-e6f8-37894313e9f6	2023-12-29 06:47:19	2024-08-04 00:50:15	LoadBalancer_Beta_Cluster
658f8c88-19ef-b7eb-5172-b7dcb9ca377a	2024-09-13 07:34:42	2024-06-12 22:25:26	PatchServer_Deploy_08
5f54bf33-d44b-8048-ee35-2aca1459b0f8	2024-04-13 17:18:19	2024-06-27 17:10:09	Cache_Server_Proxy04
08876f42-8e6b-bee6-f09b-c445d4edf545	2024-04-16 22:03:00	2024-02-23 02:14:35	SNMP_Monitor_Region02
aaa95dcf-32dc-f706-e9e2-96f994afc8f7	2024-02-27 03:46:16	2024-07-27 16:52:00	FileTransferNode_SFTP_06
dfabda78-c61a-ed6d-b7af-e0f8dbedf819	2024-08-23 05:19:30	2023-10-22 17:12:15	NetworkFirewall_ZoneAlpha
98c181d4-d008-8c33-2752-57c31af3a326	2024-08-14 02:42:04	2024-07-25 10:19:52	VoIP_Gateway_HQ_01
13faea9a-ee1b-c029-0167-082e41226ae7	2023-11-06 03:51:45	2024-07-19 17:57:39	Management_Node_SecOps_07
092d6f8b-8c53-bafb-a8d7-3c0ad508f053	2024-08-03 15:29:22	2024-04-27 08:00:32	StoragePool_Archive03
c389b95c-a580-d546-dd4e-04402a460ba4	2023-12-05 01:29:49	2024-06-07 08:13:44	Mainframe_Backup_SiteC
3e9b3285-2f74-10d9-c529-f46fb5c2024c	2024-03-20 02:16:49	2024-01-16 15:46:10	APICluster_Production_02
b04cd62f-d0a6-f0a4-ec7d-45e4586a680d	2024-03-20 18:59:45	2024-06-17 01:34:55	VPN_Tunnel_Server02
7147d009-6fec-1ed0-eada-28469ecb42c4	2024-09-05 22:40:52	2024-08-17 20:13:49	CoreSwitch_A1_Cluster
6b744ef1-4137-8e59-661a-b1def8a663fc	2023-11-25 12:45:35	2024-04-02 15:03:13	WebDBServer_NodeX_05
4c9b9ef2-eda0-5efa-0999-a33c0be7ed81	2023-10-25 03:09:31	2024-08-19 20:04:55	FileShare_Server_Finance
b6b61a56-7fd6-acbb-c062-f03624e1316c	2024-07-23 18:47:25	2024-09-28 01:30:29	PatchMgmt_Server02
7bbb864e-114e-58ae-3aa0-2bc7d3911dad	2024-01-07 07:52:09	2024-02-13 08:49:56	DataReplica_Node03
ee2c4616-1b55-10a2-5292-384038bbfa5b	2024-04-15 00:55:28	2024-09-05 01:31:00	EmailGateway_Spam_02
3a1084ec-b697-b7c0-64ea-4af4528c5e96	2023-12-30 15:18:46	2024-03-05 13:15:54	CloudGateway_Node06
682bac93-8d64-a74d-9261-8e45c6d849d9	2023-12-12 08:31:41	2024-09-26 09:50:39	API_Proxy_Beta_08
b141d90e-1e40-76ce-0418-1fc3a250dd5a	2024-07-19 05:59:47	2024-02-07 20:37:15	Router_NorthZone_01
f5e7f28a-7986-2314-6674-9b9a3b217e8b	2024-09-01 06:20:41	2024-02-23 02:11:48	DataAggregator_Node_Bravo
ce591c55-66ac-c6bb-405b-ff71e3349c4b	2024-01-10 00:41:40	2024-09-05 04:56:06	InventoryServer_DC_Alpha
d42d6e33-6e7f-dcb9-a843-1d26587a22df	2024-04-18 21:31:27	2023-12-23 01:21:40	VDI_AppNode_Central_03
a98a3d7d-d4c9-4013-5108-1d631dfa71ec	2023-12-12 20:04:11	2024-05-23 00:51:51	LogCollector_Sec02
eb734708-8dcb-2da7-ad9f-873228faf841	2024-06-11 21:03:47	2024-03-22 01:48:57	Compliance_Server_Reg02
033d0582-abf0-c0b8-c9fb-a4d455626356	2024-08-30 11:50:09	2024-09-29 13:52:29	NetworkAnalyzer_Alpha
ffb4e5cd-424e-e2a4-16f5-e4e2667ace5c	2023-12-31 02:44:06	2023-10-12 03:01:40	EdgeServer_Node07
119c851a-a2eb-db53-5b7a-232c8155eef1	2024-05-20 00:37:06	2023-11-03 00:07:30	ConfigurationNode_X09
752dd84f-49e2-0bdb-6b1d-cf4978a0f116	2023-10-26 22:33:09	2023-10-28 12:20:40	Router_Regional_B_04
5a746783-86f9-1d79-6de8-22778ccfd0e7	2024-08-15 14:56:17	2024-03-16 13:02:39	WebContent_Node_Beta_03
09138641-7752-8ebf-635a-ef68f426efcb	2024-06-06 00:57:17	2024-10-08 13:47:20	DataStream_Node_Flow_05
1d4ba58a-45ac-16cf-7e03-0d584086a8d5	2023-12-04 01:31:09	2024-07-15 01:57:16	AuthProxy_Region_West_03
3c437442-bfb7-4b4a-09e4-2ebc829de24f	2024-02-26 17:44:46	2023-10-12 20:10:24	Cluster_Compute_Beta02
6d6dda53-500d-a5c0-2472-f353bd8f11c1	2024-04-18 02:54:29	2024-08-20 19:07:27	FailoverNode_DC_01
66c9caa4-6b57-9d97-906f-07f5f28e287a	2023-11-30 23:48:11	2023-11-04 23:36:23	EventMgmt_Server_02
ee4dadfa-e43a-8f6c-4700-c3022174c61f	2024-08-17 14:02:13	2024-07-23 03:17:52	DMZ_Switch01
b4e34189-0cf7-219f-c783-538b96fb2740	2023-11-08 13:24:03	2024-02-07 14:19:29	BackupAgent_Site07
2963ea35-cec2-814c-bf09-b88484d768b7	2023-12-10 11:26:13	2024-07-15 12:39:38	Replication_NodeX_03
c1671738-69a7-049a-3c3c-e85e8313f932	2023-12-25 03:11:00	2024-01-10 02:21:07	API_Gateway_BravoRegion
f24e068b-bbbe-c98e-995d-4e7d8c952093	2024-02-03 03:14:39	2024-05-24 19:40:39	SecurityHub_Node01
3809a2e6-2e1d-43b7-edbc-12ebb09d6c80	2024-06-12 14:20:05	2024-04-09 20:03:39	SiteBackup_Node03
f6321a20-4cb0-2baa-f3b0-0957acdfa11c	2024-08-09 13:55:20	2024-04-08 10:58:49	NetworkStorage_Unit09
bda8505f-bf11-396a-91b1-478ef38bfb3a	2024-03-25 18:44:57	2023-12-06 05:10:04	WebApp_NodeZ_Beta
638d81c2-40a1-c4f4-b9e1-acf329f622bb	2024-01-29 17:10:27	2024-06-01 22:03:39	AuthServer_DC_04
9e3c2db9-b837-a7bf-4418-d6b41e800ace	2023-10-23 16:22:35	2023-11-17 00:08:28	UserMgmt_Server01
58a6aa4d-8000-c588-3ae3-568ab0c3b882	2024-04-01 20:02:47	2023-10-13 18:02:37	VirtualDisk_Array03
64a7ec4a-bd48-a5fd-d60f-5fc2894f585b	2023-10-28 03:49:04	2024-07-27 01:35:10	CloudCluster_Node01
e11b9751-91b4-430d-a55a-0c93d1267d09	2024-10-14 00:30:34.208044	2024-10-14 00:30:34.208044	CI ONE
42564003-3af3-48b7-8826-0f7fe1feac3b	2024-10-14 11:06:51.968741	2024-10-14 11:06:57.861018	CI TWO POINT FIVE
\.



COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
1	0	t	2024-10-11 15:32:45.897104
2	1	t	2024-10-11 15:32:49.681997
3	2	t	2024-10-11 15:32:49.685591
4	3	t	2024-10-11 15:32:49.693341
5	4	t	2024-10-11 15:32:49.695958
6	5	t	2024-10-11 15:32:49.696966
7	6	t	2024-10-11 15:32:49.699285
8	7	t	2024-10-11 15:32:49.701085
9	8	t	2024-10-11 15:32:49.703388
10	9	t	2024-10-11 15:32:49.704577
11	10	t	2024-10-11 15:32:49.707752
12	11	t	2024-10-11 15:32:49.709702
13	12	t	2024-10-11 15:32:49.710463
14	13	t	2024-10-11 15:32:49.711869
15	14	t	2024-10-11 15:32:49.712687
16	15	t	2024-10-11 15:32:49.713517
17	16	t	2024-10-11 15:32:49.714489
18	17	t	2024-10-11 15:32:49.715055
19	18	t	2024-10-11 15:32:49.716874
20	19	t	2024-10-11 15:32:49.719481
21	20	t	2024-10-11 15:32:49.721034
22	21	t	2024-10-11 15:32:49.72241
23	22	t	2024-10-14 00:30:20.333032
24	23	t	2024-10-14 10:52:50.646553
25	24	t	2024-10-14 11:41:33.172632
\.



COPY public.incidents (id, created_at, updated_at, short_description, description, configuration_item_id, company_id, state, assigned_to) FROM stdin;
34ea1efb-1892-40b7-8bc7-b484aba60d0e	2024-09-09 11:48:50.915394	2024-09-09 11:48:50.915394	router 3 host status	\N	2c6fc3cd-aed3-431e-850c-7bb19ad05673	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
f839d323-67ea-432d-ac8e-0a42774cc8bd	2024-09-09 11:48:50.915394	2024-09-09 11:48:50.915394	snmp settings	configure snmp for router 3	2c6fc3cd-aed3-431e-850c-7bb19ad05673	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
85215904-4bea-42a7-9382-a12b8565558f	2024-09-09 11:48:50.915394	2024-09-09 11:48:50.915394	VPN tunnel not working	users cant access corporate resources on router 2	be275908-1c23-4eb2-aa0b-086beb7334b5	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
4444f93c-a690-1fea-18bb-390cb12a3d5e	2024-04-19 12:05:37	2024-09-18 07:21:09	Esse fugiendam temperantiamque expetendam, non quia voluptates fugiat, sed quia non numquam eius modi tempora incidunt.	Vexetur, ob eumque errorem et voluptatibus maximis saepe priventur et durissimis animi doloribus torqueantur, sapientia est adhibenda, quae et terroribus cupiditatibusque detractis et omnium falsarum opinionum temeritate derepta certissimam se nobis ducem praebeat ad voluptatem. Sapientia enim est a.	64a7ec4a-bd48-a5fd-d60f-5fc2894f585b	2bc82373-8c8e-e66e-b168-c7699aa95393	New	bd08d4b7-3eb3-213e-274a-d389bbf11145
bdb038ce-2c7a-77c6-7b71-2a20ab8bc19c	2024-09-04 06:38:52	2024-07-26 16:23:11	A spe pariendarum voluptatum seiungi non potest. Atque ut odia, invidiae, despicationes adversantur voluptatibus, sic amicitiae non modo.	Etiam, ut te consule, ipsi se indicaverunt. Quodsi qui satis sibi contra hominum conscientiam saepti esse et voluptates et dolores animi quam corporis. Nam corpore nihil nisi praesens et quod adest sentire.	e6992e6a-70fe-057a-daa5-657b720e7324	1f760bdf-31d7-14f6-edee-c2b9a7a6e5fa	Resolved	33caaa90-0640-23a3-8742-0e05f97b8581
f220c247-e1e9-72fe-5feb-4400ba42e278	2024-07-29 12:37:06	2024-08-29 03:37:24	Regione ferrentur et, ut modo docui, cognitionis regula et iudicio ab eadem illa constituto veri a falso distinctio traditur. Restat locus huic disputationi vel maxime necessarius de amicitia, quam, si voluptas esset bonum.	Ego a philosopho, si afferat eloquentiam, non asperner, si non habeat, non admodum indoctis, totum hoc displicet philosophari. Quidam autem non tam solido quam splendido nomine, virtutem autem nixam hoc honesto nullam requirere voluptatem atque ad beate vivendum se ipsa allicit nec patientia nec assiduitas nec vigiliae nec ea ipsa, quae tibi.	ee4dadfa-e43a-8f6c-4700-c3022174c61f	49e072b6-94ab-0195-b106-351c0f199eec	Assigned	d7edfdd0-b755-e446-29cc-cc83b3cdcf66
dde723cb-5fc4-640c-5cbe-5773f533efba	2023-12-27 12:42:37	2024-05-04 10:32:59	Sint ea quamque iucunda, neque pendet ex futuris, sed expectat illa, fruitur praesentibus ab iisque vitiis, quae paulo ante collegi, abest plurimum et, cum.	Suapte natura, non potest esse dubium, quin id sit summum atque extremum.	c389b95c-a580-d546-dd4e-04402a460ba4	bfeff151-c385-82bc-90c2-b4b259839383	On Hold	3f405e3d-599f-cb3b-a9b2-e2b8f8985bb0
32a1ef75-5001-d916-b5da-5f20917da6e0	2023-12-06 04:50:05	2024-04-27 15:37:56	Apud se dicere iuberet reque ex utraque parte audita pronuntiaret eum non talem videri.	Dividendo ac partiendo docet, non quo ignorare vos arbitrer, sed ut ratione et via procedat oratio. Quaerimus igitur.	70ba7328-cc3a-28f0-7f13-3268f4f1bec2	1be71338-22dd-3d93-9e22-f15166473052	Resolved	33caaa90-0640-23a3-8742-0e05f97b8581
84bcb8bb-9ba0-6694-5066-eafcde150206	2023-12-26 03:09:39	2024-04-12 00:21:50	Quo enim maxime consuevit iactare vestra se oratio, tua praesertim, qui studiose antiqua persequeris.	Extremum sit, ita ferri, ut concursionibus inter se cohaerescant, ex quo vitam amarissimam necesse est aut in voluptate esse aut in liberos.	a060f7cd-dcec-a316-19af-c17e39405946	28ce2c52-0d5a-99fc-7bb5-bff1c3a288c8	Assigned	d36adaef-e572-2f65-793a-79cbe2973bcd
c1eb0ece-2587-b2bd-1871-64a26ad042c5	2023-12-30 09:49:53	2024-03-29 04:42:48	Extremum malorum? Qua de re cum sit inter doctissimos summa dissensio, quis alienum putet eius.	Si essent vera, nihil afferrent, quo iucundius, id est.	ee2c4616-1b55-10a2-5292-384038bbfa5b	a9835ad6-123b-5bd9-d664-f18d258b64e9	Assigned	ce59e370-3406-36d6-0d8d-c146f734d94b
a711de04-f023-2908-d62d-52a968b20772	2024-05-30 16:38:12	2024-04-07 19:31:13	Iudicabit nulla ad legendum his esse potiora. Quid est enim in vita tantopere quaerendum.	Quam omnia iudicia rerum in sensibus ponit, quibus si semel.	54e7fc2d-bd7c-1d09-0d0b-699f724514f6	6f115cdc-1a43-1f12-feff-57015675dbec	In Progress	3f405e3d-599f-cb3b-a9b2-e2b8f8985bb0
216b0421-e126-41d8-0d53-6bd6c99ffc8f	2024-02-06 03:57:53	2023-10-11 09:49:48	Cur Graeca anteponant iis, quae et a spe pariendarum.	Vel eum iure reprehenderit, qui in ea voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa, qui officia.	e6992e6a-70fe-057a-daa5-657b720e7324	228c7c2d-3109-cdf6-38f7-d2603d34810c	Resolved	af99cbab-f5ed-624d-4deb-f349765ff89d
8c93f39c-e141-84c1-0d80-0ddcec29bdd3	2024-10-05 07:25:04	2024-03-24 03:05:26	Placet: constituam, quid et quale sit id, de quo Lucilius: 'ferreum scriptorem', verum, opinor, scriptorem tamen, ut legendus sit. Rudem enim esse omnino in nostris poetis aut inertissimae segnitiae est aut fastidii delicatissimi.	'Haec enim ipsa mihi sunt voluptati, et erant illa Torquatis.' Numquam hoc ita defendit Epicurus neque Metrodorus aut quisquam eorum.	535b45a8-bea7-0cc5-85c0-9a61104587f5	31fd0515-b0f6-ba9c-34d4-36201e38150d	Assigned	588d5a01-bcb2-3b94-76d1-d80bf7cc5868
a3835250-7c75-95fd-c9dc-b7b4eb75e1c5	2024-10-05 14:45:37	2024-05-22 14:25:32	Te hortatore facimus, consumeret, in quibus nulla solida utilitas omnisque puerilis est delectatio, aut se, ut Plato, in.	Ut ne minus amicos quam se ipsos diligant. Quod et scripta multa sunt, sic inprobitas si cuius in.	b04cd62f-d0a6-f0a4-ec7d-45e4586a680d	effe6dce-661f-cbcd-bf91-1b30027372fc	New	1c5893a5-92cb-b792-c6b4-84461f749ed1
e9592215-5668-24a6-8c96-a8e6638a3a98	2024-05-14 11:14:41	2024-01-23 13:13:41	Adipisci velit, sed quia pacem animis afferat et eos quasi concordia quadam placet ac leniat. Temperantia est enim, quae in rebus aut expetendis aut.	Minus omnes mea legant. Utinam esset ille Persius, Scipio vero et Rutilius multo etiam magis, quod, cuius in animo versatur, numquam sinit eum respirare, numquam adquiescere. Quodsi ne ipsarum quidem virtutum laus, in qua maxime ceterorum.	0486721c-2a98-8779-c7c6-8f0d317834f2	c633399a-6fe8-39f4-a142-9ed39ebfe4f2	Resolved	bdd06817-bacf-8252-e8e4-bffe89f7670b
5f582e17-6da7-c75a-51c2-408f6164c586	2024-04-26 21:10:49	2024-04-08 10:46:31	Quis autem vel eum iure reprehenderit, qui in una virtute ponunt et.	Quidem exercitus. -- Quid ex eo est consecutus? -- Laudem.	6b744ef1-4137-8e59-661a-b1def8a663fc	835855b7-48e2-194e-6b7e-436075b012a7	Assigned	21ca2ec5-61b0-d8e1-dd1d-f334a4fa443b
4d75f336-3178-4172-cb11-fb526e60a12f	2023-11-23 17:03:17	2024-01-19 18:48:14	Titillaret sensus, ut ita ruant itaque turbent, ut earum motus et impetus quo pertineant non intellegamus, tu tam egregios viros censes tantas res gessisse sine causa? Quae fuerit causa, mox.	Aliquid ferre denique' -- non enim illum ab industria, sed ab inliberali labore deterret --, sic isti curiosi, quos offendit noster minime nobis iniucundus labor. Iis igitur est.	03ec8d05-75ce-0dac-2967-e37495283128	11bc4848-635e-2d05-9387-e20e09fe62bc	Assigned	eb9be367-874b-f3d2-7e15-7c1abb5b3f37
3f2defc8-746d-f3a3-68ba-1759a1d4764c	2024-06-03 10:35:45	2024-08-12 00:40:36	Legat? Nisi qui se Latina scripta dicunt contemnere. In quibus.	Vita, cum ea non placeat, tamquam e theatro exeamus. Quibus rebus intellegitur nec intemperantiam propter se esse iucunda, per se ipsam optabilem, sed quia non numquam eius modi.	9b3313d5-074f-689b-f731-1115d390aeca	e55ffd34-5054-b617-c1bf-097c08ae4c89	In Progress	8d9d5b42-4aa5-e93d-eb12-d1cc093265c7
6fd020f1-03e2-9913-0159-b454953b6229	2024-08-18 03:48:19	2024-04-17 17:57:56	Tibi non vera videantur. Vide, quantum, inquam, fallare, Torquate. Oratio me istius philosophi non offendit; nam et complectitur verbis, quod vult, et dicit plane, quod intellegam; et tamen ego a philosopho, si afferat eloquentiam, non asperner, si non habeat, non admodum indoctis, totum hoc.	'Eadem', inquit, 'scientia confirmavit animum, ne quod aut sempiternum aut diuturnum timeret malum, quae perspexit in hoc ipso vitae spatio amicitiae praesidium esse firmissimum.' Sunt autem.	3809a2e6-2e1d-43b7-edbc-12ebb09d6c80	534df3a2-fbc9-cc82-e64c-51c090c12b29	New	bd08d4b7-3eb3-213e-274a-d389bbf11145
2fd35e4a-cf02-2a98-7596-045f90db7901	2023-10-25 21:18:19	2024-08-22 22:04:02	Possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam.	Dicant foedus esse quoddam sapientium, ut ne minus amicos quam.	92f225a5-94e1-9a62-99d1-33227da0d19c	b92ed5ea-5029-9fad-daad-1d25f985db12	In Progress	4c59113d-9dc0-b754-4ee0-c76add90488e
b3a02de6-c636-c084-967d-ff0673876479	2024-07-02 00:48:05	2024-05-21 03:45:18	De materia disseruerunt, vim et causam efficiendi reliquerunt.	Satis acuti, qui verentur ne, si amicitiam propter nostram voluptatem expetendam putemus, tota amicitia quasi claudicare videatur. Itaque primos congressus copulationesque et consuetudinum instituendarum voluntates fieri propter voluptatem; cum autem usus progrediens familiaritatem.	e6992e6a-70fe-057a-daa5-657b720e7324	aad132b5-3f63-c9a6-3370-8d6ca03819b9	Resolved	4c59113d-9dc0-b754-4ee0-c76add90488e
94960f8d-f108-75a6-8c55-a2faeadcaab5	2024-04-26 18:19:43	2024-07-17 05:26:09	Voluptate conectitur. Nam et complectitur verbis, quod vult, et dicit plane, quod intellegam; et tamen in quibusdam neque pecuniae modus est.	Satis sibi contra hominum conscientiam saepti esse et voluptates et dolores nasci fatemur e corporis voluptatibus et doloribus -- itaque concedo, quod modo dicebas, cadere causa, si qui e nostris aliter existimant, quos quidem video minime esse deterritum. Quae cum tota res (est) ficta pueriliter, tum ne efficit quidem, quod vult. Nam et praeterita grate meminit.	30c30ce2-0245-e36c-3ec5-3345e9e34e2f	835855b7-48e2-194e-6b7e-436075b012a7	On Hold	cd07f4b1-eff6-5f39-f455-45fdb9d6f644
d9cb4ba7-8b68-c5ff-9d95-280d9a7d0369	2024-01-25 18:07:35	2024-03-04 12:26:40	Probarem, quae ille diceret? Cum praesertim illa perdiscere ludus esset. Quam ob rem voluptas expetenda, fugiendus dolor sit. Sentiri haec putat, ut calere ignem, nivem esse albam, dulce mel. Quorum nihil oportere exquisitis rationibus confirmare, tantum satis esse admonere. Interesse enim inter argumentum conclusionemque.	Semel aliquid falsi pro vero probatum.	535b45a8-bea7-0cc5-85c0-9a61104587f5	31fd0515-b0f6-ba9c-34d4-36201e38150d	Assigned	b28c5a1d-e325-1f0a-67b3-3239a9ede311
575ccabe-5ff0-2a63-02b1-26663a5c6559	2024-01-21 03:22:25	2024-05-15 07:05:30	Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut perferendis doloribus asperiores repellat. Hanc ego.	Liberamur mortis metu, non conturbamur ignoratione rerum, e qua ipsa horribiles existunt saepe formidines. Denique etiam morati melius erimus, cum didicerimus quid natura postulet.	2660389a-9b80-9c35-7426-edda56015dc1	5008195d-18db-a8e1-23f3-f1199765612f	New	14d8e1a9-1967-d96d-e2c3-6418d974c060
ec979de0-435a-a4ac-7e98-6df5ad40fedd	2024-05-20 18:52:49	2024-06-17 05:15:52	Causa, nollem me ab eo et gravissimas res consilio ipsius.	Cognitione et scientia tollitur omnis ratio et vitae degendae et rerum gerendarum. Sic e physicis et fortitudo sumitur contra mortis timorem et constantia contra metum religionis et sedatio animi omnium rerum occultarum.	0486721c-2a98-8779-c7c6-8f0d317834f2	669589b7-1180-c01e-687d-fff150709e47	On Hold	163a4480-1335-f3f0-2cd1-c5c76cd06cc1
42a2296f-4353-a2cc-4272-e0878270e97d	2024-01-08 06:31:18	2024-06-06 07:38:56	Etiam labefactant saepe rem publicam. Ex cupiditatibus odia, discidia, discordiae, seditiones, bella nascuntur, nec eae se foris solum iactant nec tantum.	Sed multitudinem haec maxime allicit, quod ita putant dici ab illo, recta et honesta quae sint, ea facere ipsa per se ipsas tam expetendas.	e8931686-7088-57b1-af6d-780b578f2b1e	0eac1ef4-3188-a9c4-fc23-659d7fe784cc	New	bd08d4b7-3eb3-213e-274a-d389bbf11145
ef67dbea-bdff-de12-6876-4afc994afcfc	2023-10-21 16:10:53	2023-12-14 09:48:32	Sit, ut tollatur error omnis imperitorum intellegaturque ea, quae voluptaria, delicata, mollis habeatur disciplina, quam gravis, quam continens, quam severa sit. Non enim hanc solam sequimur, quae suavitate aliqua naturam.	Successerit, eoque intellegi potest quanta voluptas sit non dolere. Sed ut omittam pericula, labores, dolorem etiam, quem optimus quisque pro patria et pro suis suscipit, ut non plus habeat sapiens, quod gaudeat, quam.	81f3ffae-0d75-3321-8923-233568580371	05edfb15-7467-fdd6-eac2-8391b261f992	On Hold	05c3dc04-a9c7-caea-51fa-ca9311b5f3a1
9ef39b5f-524a-aa1f-9d2a-248e7c31bfea	2024-05-16 20:21:55	2024-04-27 21:30:34	Ipse constituit, e regione ferrentur et, ut modo docui, cognitionis regula.	Non asperner, si non habeat, non admodum indoctis, totum hoc displicet philosophari. Quidam autem non tam id reprehendunt, si remissius agatur, sed tantum studium tamque multam operam ponendam in eo essent. Quae cum tota res (est.	03ec8d05-75ce-0dac-2967-e37495283128	0ae21192-df8a-f334-83f5-6bb154dfab36	In Progress	82c84588-62e2-6a05-3bd4-df8618e9b1f7
e7ba3dd9-b2a3-7462-3af3-faceadba71df	2024-09-16 13:50:56	2023-12-26 02:58:34	Tamen nec modus est ullus investigandi veri, nisi inveneris, et quaerendi defatigatio turpis est, cum id, quod propositum est, summum bonum esse vult, summumque malum dolorem.	Iudicatum. Plerique autem, quod tenere atque servare id, quod maxime placeat, facere.	1d4ba58a-45ac-16cf-7e03-0d584086a8d5	88ece235-12ad-1182-2b71-3293dbc4ff0a	On Hold	b01d292f-8c87-e144-5218-34d6a5e82539
f59483aa-235f-2a74-4482-bf6e3dcf0c33	2024-08-29 05:01:29	2023-11-05 09:38:39	Cum praesertim illa perdiscere ludus esset. Quam ob.	Quibus, quantum potuimus, non modo quid nobis probaretur, sed etiam praetereat omnes voluptates.	78778af7-44fe-eadc-991f-867056c8af93	4932fd4e-d82f-7eef-4b21-f46a9f1c8c6c	In Progress	8ed88161-56ed-751c-0295-9b63a81a9261
acb4c1fa-7d90-51ee-df49-3652784d6bad	2023-12-24 17:18:14	2024-06-26 13:48:25	Cum Latinis tertio consulatu conflixisse apud Veserim propter voluptatem; cum autem usus progrediens familiaritatem effecerit, tum amorem efflorescere tantum, ut, etiamsi.	Quod aut sempiternum aut diuturnum timeret malum, quae perspexit in hoc ipso.	e8931686-7088-57b1-af6d-780b578f2b1e	722a2243-b5ee-fba0-2400-95b939f87392	Resolved	76908e10-9e68-69c2-c623-d26192f75e6b
0bbaf8d8-e605-b24d-7245-4c86b7558f55	2023-11-23 16:06:35	2024-09-12 19:48:10	Novi, ea tamen, quae te ipsum probaturum esse confidam. Certe, inquam, pertinax.	Si bona. O praeclaram beate vivendi et apertam et simplicem et directam viam! Cum enim certe nihil homini possit melius esse quam Graecam. Quando enim nobis, vel dicam aut oratoribus bonis aut poetis, postea quidem quam fuit quem imitarentur.	5a746783-86f9-1d79-6de8-22778ccfd0e7	faf5c017-5cfc-8135-6e95-7f02e6fcb6e6	New	f7845a09-230d-6cda-e45c-bc47c59944cf
e705afe6-6f42-cdb1-26e2-af9e375c9eca	2024-05-17 03:23:31	2024-05-09 11:07:10	Equos, si ludicra exercendi aut venandi consuetudine adamare solemus, quanto id in hominum consuetudine facilius fieri poterit et iustius? Sunt autem, qui dicant foedus esse quoddam sapientium, ut.	Sic agam, ut ipsi auctori huius disciplinae placet: constituam, quid et quale sit id, de quo quaerimus, non quo ignorare vos arbitrer, sed ut ratione et via procedat oratio. Quaerimus igitur, quid sit extremum et ultimum.	f5e7f28a-7986-2314-6674-9b9a3b217e8b	b4020a50-cba1-0bb0-9218-07c5aec9fe2d	Resolved	a8855e8c-d01f-a593-ac43-9eb0922a66a0
7e05f964-7376-45ca-c2b3-695d30d83ac5	2023-12-03 14:12:36	2024-03-08 01:46:55	Allicit nec patientia nec assiduitas nec vigiliae nec ea ipsa, quae tibi probarentur; si qua in iis corrigere voluit, deteriora fecit. Disserendi artem nullam habuit. Voluptatem cum summum bonum in voluptate ponit, quod summum bonum consequamur? Clamat Epicurus, is quem vos nimis voluptatibus esse deditum dicitis; non posse.	Scribendi ordinem adiungimus, quid habent, cur Graeca anteponant iis, quae et a falsis initiis.	752dd84f-49e2-0bdb-6b1d-cf4978a0f116	7440373d-5774-2a6a-2451-2f5d4032de26	Resolved	33caaa90-0640-23a3-8742-0e05f97b8581
8cfffcce-7e4c-9758-d200-5c32aef602a9	2024-08-07 16:59:30	2024-04-25 12:23:30	Momenti quam eorum utrumvis, si aeque diu sit in corpore. Non placet.	Quam vacare omni dolore detracto, nam quoniam, cum privamur dolore, ipsa liberatione et.	c79b737e-c1f1-a4ca-9130-aa83b1f67f6a	a9b46d58-b3f7-f518-d524-d22b3c63a2f9	On Hold	21ca2ec5-61b0-d8e1-dd1d-f334a4fa443b
eb136f9c-047c-a467-c637-fc268569617e	2024-02-05 20:17:17	2024-07-27 05:20:50	Quisquam est, qui Ennii Medeam aut Antiopam Pacuvii spernat aut reiciat, quod se.	Exquisitis rationibus confirmare, tantum satis esse admonere. Interesse enim inter argumentum conclusionemque rationis et inter mediocrem animadversionem atque admonitionem. Altera.	81f3ffae-0d75-3321-8923-233568580371	d99927dc-2c3d-57e6-f982-3de64c404104	On Hold	197f65b3-76e8-4017-624b-21d733a776c1
754d4a28-1ab3-abf2-4af7-b555343d9211	2023-11-26 13:00:46	2024-05-29 07:55:20	Ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate aut in.	Tollunt et nihil posse percipi dicunt, ii remotis sensibus ne id ipsum quidem expedire possunt, quod disserunt. Praeterea sublata cognitione et scientia tollitur omnis ratio et vitae degendae et rerum gerendarum. Sic.	8f7b1456-6da2-5830-72af-6c1b26c711c7	db71629a-0f48-27c6-bb94-025ec4e96335	Resolved	1f7ec3f0-ba9a-6294-70dc-085776867f9c
9ee1f893-6f82-c113-aa0f-545d740f0126	2023-10-16 14:55:15	2023-12-16 04:56:23	Aut Zenonem putas, quorum utrumque audivi, cum miraretur ille quidem utrumque, Phaedrum autem.	Non paranda nobis solum ea, sed fruenda etiam sapientia est; sive hoc difficile est, tamen nec modus ullus nec finis inveniri potest. Quodsi corporis gravioribus morbis vitae iucunditas impeditur, quanto magis animi morbis impediri necesse est! Animi autem morbi sunt.	ee4dadfa-e43a-8f6c-4700-c3022174c61f	eae20175-f5fd-8614-b8af-360f15350068	Assigned	2f919ce3-fa2e-e147-3811-79c59d0d6905
26ee3c67-ef52-9efb-6eb4-204eeb2a667e	2024-07-30 11:22:16	2023-12-21 02:33:08	Et ratione administrari neque maiorem voluptatem ex hoc facillime perspici potest: Constituamus aliquem magnis, multis, perpetuis fruentem et animo et attento intuemur, tum fit ut aegritudo sequatur, si illa mala sint, laetitia, si bona. O praeclaram.	Quasi naturalem atque insitam in animis nostris inesse notionem, ut.	199e674b-98fa-1441-2c8d-b91bcf25133f	0fcb4b25-d3a6-f6ca-9507-b5e9e8f7438e	New	2f919ce3-fa2e-e147-3811-79c59d0d6905
65ef87e0-f903-b6a5-057b-bff6078a6c3e	2024-01-02 01:29:44	2023-10-25 16:36:42	Earum rerum defuturum, quas natura non depravata desiderat. Et quem.	Vivendum reperiri posse, quod coniunctione tali sit aptius. Quibus ex omnibus iudicari potest non modo non repugnantibus, verum etiam summam voluptatem. Quisquis enim sentit, quem ad modum sit affectus, eum necesse est effici, ut sapiens solum amputata circumcisaque inanitate omni et errore naturae finibus contentus sine aegritudine possit et sine metu degendae praesidia firmissima. -- Filium.	33de396b-1809-3649-4735-07b69d8491af	a9835ad6-123b-5bd9-d664-f18d258b64e9	Resolved	e1f81b4c-bab0-7359-7928-bb980ea5c495
5effcdb7-21f9-2fa4-44b7-87f976bf566b	2024-07-11 12:55:39	2024-04-03 20:54:02	Autem nusquam. Hoc Epicurus in voluptate aut a dolore. Quod cum ita esset affecta, secundum non.	Graecis legendis operam malle consumere. Postremo aliquos futuros suspicor, qui me ad alias litteras vocent, genus hoc scribendi, etsi sit elegans, personae tamen et.	0486721c-2a98-8779-c7c6-8f0d317834f2	a48aea8e-0278-00d4-3fe0-23933ea4554f	Assigned	f8dcce78-9997-fd82-026d-35e9668e6fed
e48d6c3b-ddd4-f337-e60b-e7933b9befe3	2023-10-18 23:17:05	2024-01-14 17:12:04	Utuntur, benivolentiam sibi conciliant et, quod aptissimum est ad quiete vivendum, caritatem, praesertim cum omnino nulla sit utilitas ex amicitia.	Quae logikh dicitur, iste vester plane, ut mihi quidem videtur. Ac fieri potest, ut errem, sed ita prorsus existimo, neque eum Torquatum, qui hoc primus cognomen invenerit, aut torquem illum hosti detraxisse, ut aliquam ex eo perciperet corpore voluptatem, aut cum Latinis tertio consulatu conflixisse apud Veserim propter voluptatem; cum autem usus progrediens familiaritatem effecerit, tum.	5f54bf33-d44b-8048-ee35-2aca1459b0f8	e55ffd34-5054-b617-c1bf-097c08ae4c89	In Progress	7a5f4d28-366e-6743-a993-a1fe1c554af0
78ed8c39-34bb-17e1-c0b7-cd380ddcffb9	2024-05-02 19:37:19	2024-08-27 04:34:36	Est res ulla, quae sua natura aut sollicitare possit aut angere. Praeterea et appetendi et refugiendi et omnino rerum gerendarum initia proficiscuntur aut a voluptate aut a dolore.	Et complectitur verbis, quod vult, et dicit plane, quod intellegam; et tamen in.	6d6dda53-500d-a5c0-2472-f353bd8f11c1	bfeff151-c385-82bc-90c2-b4b259839383	Resolved	07a7b9db-403f-9507-5c1a-74a0759d49af
c2b2eb56-1c9a-dc1f-3677-000d18018f22	2024-04-12 05:48:42	2024-03-19 12:00:39	Prosperum nisi voluptatem, nihil asperum nisi dolorem, de quibus neque depravate iudicant neque corrupte, nonne ei maximam gratiam habere debemus, qui hac exaudita quasi voce naturae sic eam.	Quae laudatur, industria, ne fortitudo quidem, sed.	c389b95c-a580-d546-dd4e-04402a460ba4	88ece235-12ad-1182-2b71-3293dbc4ff0a	Resolved	d7edfdd0-b755-e446-29cc-cc83b3cdcf66
44886b1b-c86e-4547-b562-c6b08d8a092a	2024-02-01 17:23:47	2024-02-09 17:44:30	Autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet, ut et adversa quasi perpetua oblivione obruamus et secunda iucunde ac suaviter meminerimus. Sed cum ea, quae.	Hostis mi Albucius, hinc inimicus. Sed iure Mucius. Ego autem mirari satis non queo unde hoc sit tam.	3e9b3285-2f74-10d9-c529-f46fb5c2024c	00b73e30-250e-5127-059a-285085d02814	Assigned	6b0884fb-f2c4-fcc8-bf51-0c906ffb5b05
ba8fef09-902f-ed42-ac4e-236f8fa8fdbb	2024-01-25 08:44:35	2024-01-08 21:04:28	Praesenti nec expectata voluptate, quid eo miserius dici aut fingi potest? Quodsi vita doloribus referta maxime fugienda est.	Praesidia firmissima. -- Filium morte multavit. -- Si sine causa, nollem me ab eo dissentiunt, sed certe non probes, eum quem ego arbitror unum vidisse verum maximisque erroribus animos hominum liberavisse et omnia tradidisse, quae pertinerent ad bene beateque vivendum.	dfabda78-c61a-ed6d-b7af-e0f8dbedf819	1b43cc73-e173-88b5-8ee3-ce27ab1f7c62	On Hold	163a4480-1335-f3f0-2cd1-c5c76cd06cc1
c2cf4ecf-38fa-79c8-7ca9-8f67ff7132a8	2024-10-05 11:45:55	2024-07-05 16:13:05	Alii autem etiam amatoriis levitatibus dediti, alii petulantes, alii audaces, protervi, idem intemperantes et ignavi, numquam in sententia permanentes, quas.	Autem, qui dicant foedus esse quoddam sapientium, ut ne minus amicos quam se ipsos amentur. Etenim si delectamur, cum scribimus, quis est tam invidus, qui ab eo dissentiunt.	7f5d0b1d-78ba-6463-0d82-67273d1cd09b	534df3a2-fbc9-cc82-e64c-51c090c12b29	Resolved	197f65b3-76e8-4017-624b-21d733a776c1
05e3ea8c-df20-25d0-b5ac-fdd8acebfc8b	2024-07-26 05:35:18	2024-10-02 03:27:18	Et moderatio natura cupiditatum generibusque earum explicatis, et, ut dixi.	Et vita sine amicis insidiarum et metus plena sit, ratio ipsa monet amicitias.	3c802128-3f54-a54b-7eb5-a4097fedc549	a9835ad6-123b-5bd9-d664-f18d258b64e9	Resolved	d7edfdd0-b755-e446-29cc-cc83b3cdcf66
1c8de53b-f4c4-8d42-a81a-c14a0e59925c	2024-03-24 00:01:50	2024-08-13 14:07:41	In quibusdam neque pecuniae modus est neque honoris neque imperii nec libidinum nec epularum nec reliquarum cupiditatum, quas nulla praeda umquam improbe parta minuit, sed potius inflammat, ut coercendi magis quam dedocendi esse videantur. Invitat igitur vera ratio bene sanos ad iustitiam, aequitatem, fidem, neque homini infanti aut inpotenti iniuste facta conducunt, qui nec facile efficere.	Placatae, tranquillae, quietae, beatae vitae disciplinam iuvaret. An ille tempus aut in voluptate est. Extremum autem esse bonorum eum voluptate vivere. Huic certae stabilique sententiae quae sint coniuncta explicabo brevi. Nullus in ipsis error est finibus.	170a4504-5812-37bb-2875-9c13958634a7	72dff6a3-ba73-f705-5607-0356aa8fa23b	Assigned	8d9d5b42-4aa5-e93d-eb12-d1cc093265c7
9a9ecf60-f737-8a8a-ac1b-21493f27b931	2023-10-30 09:22:42	2024-02-02 14:44:49	Tibi, Torquate, quid huic Triario litterae, quid historiae cognitioque rerum, quid poetarum evolutio, quid tanta tot versuum memoria voluptatis affert? Nec mihi tamen, ne faciam, interdictum puto. Locos quidem quosdam, si videbitur, transferam, et maxime ab iis, quos modo.	Physici, credere aliquid esse minimum, quod profecto numquam putavisset, si a Polyaeno, familiari suo, geometrica discere maluisset quam illum etiam ipsum dedocere. Sol Democrito magnus videtur, quippe homini.	08876f42-8e6b-bee6-f09b-c445d4edf545	ff794092-d8db-236f-9c6f-699d9ca6c3b6	On Hold	197f65b3-76e8-4017-624b-21d733a776c1
cddc6a07-d5c3-c2b0-5f08-84311037621c	2023-12-10 15:18:52	2024-06-07 22:28:17	Zenonem putas, quorum utrumque audivi, cum miraretur ille quidem utrumque, Phaedrum autem etiam amatoriis levitatibus dediti, alii petulantes, alii audaces, protervi, idem intemperantes et ignavi, numquam in sententia permanentes, quas ob causas in eorum vita nulla.	Maximam adipiscuntur praetermittenda voluptate. Idem etiam dolorem saepe perpetiuntur, ne, si amicitiam propter nostram voluptatem expetendam putemus, tota amicitia quasi claudicare videatur. Itaque primos congressus copulationesque.	35954fb6-5948-6cc0-fd38-406318851b66	440c68ba-e930-65f2-7013-b7f0f8c34324	In Progress	8ed88161-56ed-751c-0295-9b63a81a9261
7553fbde-0a24-6a3a-d10a-475c50bf8f02	2024-05-23 19:59:15	2024-01-23 12:22:51	Facile efficere possit, quod conetur, nec optinere, si effecerit, et opes vel fortunae vel ingenii liberalitati magis conveniunt, qua qui utuntur.	Sint opera, studio, labore meo doctiores cives mei, nec cum iracundia aut pertinacia recte disputari potest. Sed ad haec, nisi molestum est, habeo quae velim. An me, inquam.	08876f42-8e6b-bee6-f09b-c445d4edf545	aa474f30-3409-a176-bbc3-60bf568124e9	On Hold	788201bf-293d-cce9-f450-4587bee8c48d
4f9188ad-adcf-976e-1707-c0b4bc8280c6	2024-09-09 17:40:15	2024-03-27 20:05:08	Comparaverit, nihil esse prosperum nisi voluptatem, nihil asperum nisi dolorem, de quibus.	Nostris non legantur? Quamquam, si plane sic verterem Platonem aut Aristotelem, ut verterunt nostri poetae fabulas, male, credo, mererer de meis civibus, si ad eorum cognitionem divina illa ingenia transferrem.	cab16fbe-a515-cb4d-d658-fb73b4725d42	f13b38a9-9c31-ca29-b6fe-569fbd3049b3	In Progress	12efc06b-c25d-79c6-687b-c695935d6d7f
bab9746d-1dbb-013d-7621-cf3a0d41a383	2023-12-28 04:52:47	2024-08-21 03:43:49	An de una voluptate quaeri, de qua.	Ullus nec finis inveniri potest. Quodsi corporis gravioribus morbis vitae iucunditas impeditur, quanto magis animi morbis impediri necesse est! Animi autem.	78778af7-44fe-eadc-991f-867056c8af93	6922cc8b-cc03-9647-9854-d478746cbc33	In Progress	8717f9ba-0af0-f8fd-c470-3f70eeeaf481
1fb0efb7-12a1-b1c8-1db2-75fc6ddba9d1	2023-11-09 00:07:03	2024-03-12 21:44:14	Illa individua et solida corpora ferri deorsum suo pondere ad lineam, numquam fore ut hic noster labor in varias reprehensiones incurreret. Nam quibusdam, et iis servire, qui vel utrisque litteris uti velint vel, si suas habent.	Musicis, geometria, numeris, astris contereret, quae et splendide dicta sint neque sint conversa de Graecis? Nam si concederetur, etiamsi ad corpus referri, nec ob.	7147d009-6fec-1ed0-eada-28469ecb42c4	84921767-114e-2a99-49c5-d13e8f77ec27	New	7440acd6-a0bb-1f71-57a1-99e271efdbe8
fd947eac-86f0-8cd0-7f13-6dc299608829	2024-05-31 22:06:19	2024-03-09 21:10:25	Ad voluptatem, voluptas autem est sola, quae nos a libidinum impetu et a spe pariendarum voluptatum seiungi non potest. Atque ut odia, invidiae, despicationes adversantur voluptatibus.	Omnes veri erunt, ut Epicuri ratio docet, tum denique poterit aliquid cognosci et percipi.	602ce54a-d71f-c1e1-e6f8-37894313e9f6	88ece235-12ad-1182-2b71-3293dbc4ff0a	New	54e9f34a-9117-04aa-1fa5-947f77c88324
5993fd39-5c88-a751-34a6-3132f6d266ba	2024-06-02 06:46:05	2024-04-16 12:49:14	Sive aliae declinabunt, aliae suo nutu recte ferentur, primum erit hoc quasi, provincias atomis.	Tritani, centurionum, praeclarorum hominum ac primorum signiferumque, maluisti dici. Graece ergo.	54e7fc2d-bd7c-1d09-0d0b-699f724514f6	b92ed5ea-5029-9fad-daad-1d25f985db12	New	197f65b3-76e8-4017-624b-21d733a776c1
ecec745f-3c2a-34b1-a539-d86af75422e6	2024-05-26 23:16:46	2024-01-28 08:14:58	Laudem et caritatem, quae sunt vitae sine metu degendae praesidia firmissima. -- Filium morte multavit. -- Si sine causa, nollem me ab eo nos abducat? Sin laboramus, quis est, qui alienae modum statuat.	Significet illum in hae esse rogatiuncula delectatum: 'Numquidnam manus tua sic affecta, quem ad modum, quaeso, interpretaris? Sicine eos censes aut in liberos atque in sanguinem suum tam.	77125244-6c08-97a8-3d6a-e825dc6ddefc	534df3a2-fbc9-cc82-e64c-51c090c12b29	New	770c94e7-6081-ea4d-b3be-b06c69c9e3fc
e2379109-de30-a2ca-cc1f-c44de00f7570	2024-08-11 07:01:46	2024-08-08 17:57:58	Cum tractat locos ab Aristotele ante tractatos? Quid? Epicurei num desistunt de isdem, de quibus.	. . .' nihilo minus legimus quam hoc idem Graecum.	b72080f1-b027-131b-9b00-3d23be8de8fe	f5387816-9b47-083f-90c9-8ff3e709ec1e	In Progress	15d33f02-f026-3606-45d2-da7d906fd91e
0a10c3ca-532c-d260-2517-a6ff76d54e11	2023-10-31 04:44:15	2024-09-04 11:53:56	Quod propositum est, summum bonum esse vult, summumque malum dolorem, idque instituit docere sic: Omne animal, simul atque natum sit, voluptatem appetere.	Geometrica discere maluisset quam illum etiam ipsum dedocere. Sol Democrito magnus videtur, quippe homini erudito in geometriaque perfecto, huic pedalis fortasse; tantum enim esse censet, quantus videtur, vel paulo aut maiorem aut.	ab14f2d8-aa2f-f414-35e7-c30a04b282a3	4897dc7d-a3b3-c70d-d854-f3b88ea3cc90	New	7a5623f3-f15c-9bed-05fc-1308b1b28a26
ec5dab3d-98e2-641d-aa37-5d5e47070e46	2024-05-27 12:09:52	2024-08-03 18:12:02	Voluptatem. Quisquis enim sentit, quem ad modum sit affectus, et firmitatem animi nec mortem nec dolorem timentis, quod mors sensu careat, dolor in longinquitate levis, in gravitate brevis soleat esse, ut ad Orestem pervenias profectus a.	Accusantium doloremque laudantium, totam rem aperiam eaque ipsa, quae ab illo est, tum innumerabiles mundi, qui et oriantur.	119c851a-a2eb-db53-5b7a-232c8155eef1	a988815c-3672-d0fe-1865-1c97fbc29117	Resolved	6eca4019-a2f3-d0fe-cf94-c69ad8a68291
cc84778f-e38a-c0d3-7e7e-6b02f4208563	2024-08-09 19:35:25	2024-08-17 03:39:50	Sin laboramus, quis est, qui dolorem ipsum, quia dolor sit, amet, consectetur, adipisci velit, sed quia consequuntur magni dolores eos, qui ratione.	Putem, de quo quaerimus, non quo modo efficiatur concludaturque ratio tradit, non qua via captiosa solvantur ambigua distinguantur ostendit; iudicia rerum in sensibus ponit.	d42d6e33-6e7f-dcb9-a843-1d26587a22df	44009247-c122-b704-8e5e-10b8624eee73	On Hold	84104518-2dce-4552-e75a-62566a17d9b5
c5b8f5ff-4f98-e1a0-7eb3-d1fda1b7ead0	2024-05-22 08:55:51	2024-09-05 05:41:21	A libidinum impetu et a falsis initiis profecta vera esse non possunt, victi et debilitati.	Modo nullam captet, sed etiam effectrices sunt voluptatum tam amicis quam sibi, quibus non solum videamus, sed etiam spe eriguntur consequentis ac posteri temporis. Quod quia nullo modo sine amicitia firmam et perpetuam iucunditatem vitae tenere.	64a7ec4a-bd48-a5fd-d60f-5fc2894f585b	72dff6a3-ba73-f705-5607-0356aa8fa23b	Resolved	7440acd6-a0bb-1f71-57a1-99e271efdbe8
5d02ca31-e45b-4808-f5a6-fe201ea924fb	2024-02-25 15:03:37	2023-12-01 16:59:14	Percurri omnem Epicuri disciplinam placet an de una voluptate quaeri, de qua Epicurus quidem ita dicit, omnium rerum, quas ad beate vivendum se ipsa esse contentam. Sed possunt haec quadam ratione dici non modo quid nobis probaretur, sed etiam cogitemus; infinitio ipsa, quam apeirian vocant, tota ab illo est, tum innumerabiles mundi.	E nostris, qui haec subtilius velint tradere et negent satis esse, quid bonum sit aut quid malum, sensu iudicari, sed animo etiam ac ratione.	54e7fc2d-bd7c-1d09-0d0b-699f724514f6	17275e14-a533-f0a2-4ef7-2e47e87dcc26	New	3f405e3d-599f-cb3b-a9b2-e2b8f8985bb0
78dc875a-f109-c6e7-9364-29cfba9cf688	2023-10-10 09:49:59	2024-07-14 18:33:07	Sublata et moderatio natura cupiditatum generibusque earum explicatis, et, ut dixi, ad lineam, numquam fore ut atomus altera alteram posset attingere itaque ** attulit rem commenticiam: declinare dixit atomum perpaulum, quo nihil posset fieri minus; ita effici complexiones et copulationes.	Eius esse dignitatis, quam mihi quisque tribuat, quid in omni re doloris amotio successionem efficit voluptatis. Itaque non placuit Epicuro medium esse quiddam inter dolorem et voluptatem; illud enim ipsum, quod quibusdam medium videretur, cum omni dolore et molestia perfruique.	5a746783-86f9-1d79-6de8-22778ccfd0e7	15f062d0-ba85-98f0-f18a-1e970282bd78	On Hold	7440acd6-a0bb-1f71-57a1-99e271efdbe8
d8db3e23-2b76-ae4b-d420-253fb5002253	2023-11-10 05:29:02	2024-07-30 20:21:21	Voluptatum tam amicis quam sibi, quibus non solum videamus, sed etiam quid a singulis philosophiae disciplinis.	Faciunt, qui ab eo et gravissimas res consilio ipsius.	eb734708-8dcb-2da7-ad9f-873228faf841	4897dc7d-a3b3-c70d-d854-f3b88ea3cc90	Assigned	eb9be367-874b-f3d2-7e15-7c1abb5b3f37
f8bfd686-07ac-4c89-ce08-d31fc36fec30	2023-12-22 00:14:26	2024-06-18 10:48:44	Enim tempus est ullum, quo non plus voluptatum habeat quam dolorum. Nam et complectitur verbis, quod vult, et dicit plane, quod intellegam; et tamen ego a.	Gravissimis rebus non delectet eos sermo patrius, cum idem fabellas Latinas ad verbum e Graecis expressas non inviti legant. Quis enim.	b04cd62f-d0a6-f0a4-ec7d-45e4586a680d	2bc82373-8c8e-e66e-b168-c7699aa95393	On Hold	04144b42-aab1-883f-0823-d57e77e5c2bd
01ee5d46-cd1b-e616-91f8-45285a44347f	2024-05-22 00:21:47	2024-06-12 02:56:10	Quam ex hoc facillime perspici potest: Constituamus aliquem magnis.	Tamquam in extremo, omnesque et metus et aegritudines ad dolorem referuntur, nec praeterea est res ulla, quae sua natura aut sollicitare possit aut angere. Praeterea et.	c389b95c-a580-d546-dd4e-04402a460ba4	7440373d-5774-2a6a-2451-2f5d4032de26	On Hold	b5ccf8d1-c14c-d393-9cf2-f47bf1d068ed
5bc3bda0-bddd-39eb-0e61-199bd852826b	2024-08-10 02:48:08	2023-11-30 22:24:17	Captet, sed etiam cogitemus; infinitio ipsa, quam apeirian vocant, tota ab illo inventore veritatis et quasi architecto beatae vitae deduceret? Qui quod.	Quaerimus igitur, quid sit extremum.	13faea9a-ee1b-c029-0167-082e41226ae7	a2767907-c217-2ee4-9727-177ac526ac73	Assigned	8e17ee9f-82f0-09e0-5918-e828b31a9265
a72185a1-ada3-6e62-e059-797ca0d74932	2024-10-07 10:51:43	2024-03-27 17:33:29	Quicquam nisi nescio quam illam umbram, quod appellant honestum non tam.	Ab iis, quos ego posse iudicare arbitrarer, plura suscepi veritus ne movere hominum studia viderer, retinere non posse.	3c437442-bfb7-4b4a-09e4-2ebc829de24f	033ba238-5996-57c3-7fb5-322b8c8f10c3	New	8dbfdfdf-9ce9-123b-3718-bdb79bcac1b7
2e7f1463-8ce8-e002-23fe-410abffd7388	2024-02-28 06:35:07	2024-09-18 11:30:31	Beatus. Multoque hoc melius nos veriusque quam Stoici. Illi enim negant esse bonum iucunde vivere. Id qui in una virtute ponunt et splendore nominis capti quid natura desideret. Tum vero, si stabilem.	Enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut.	b72080f1-b027-131b-9b00-3d23be8de8fe	5008195d-18db-a8e1-23f3-f1199765612f	New	357e8baa-8b8e-1dee-5492-c7fdb2ca42ff
e638a71f-ab13-c5d7-a54e-9e7799e528dd	2024-05-26 09:01:46	2024-09-04 01:44:21	Graecum te, Albuci, quam Romanum atque Sabinum, municipem Ponti, Tritani, centurionum, praeclarorum hominum ac primorum signiferumque, maluisti dici. Graece ergo praetor Athenis, id quod maluisti, te, cum ad te ne.	Est aut in liberos atque in sanguinem suum tam crudelis fuisse, nihil ut de utilitatibus, nihil ut de.	119c851a-a2eb-db53-5b7a-232c8155eef1	59bb140c-94d1-a296-ee4f-5bf5b86516dc	New	f897f55c-4f77-1c57-f838-3981895ffaea
d62dda90-b9e1-a23c-abc5-623582d82928	2024-07-27 06:01:31	2024-03-24 15:47:25	Quod quam magnum sit fictae veterum fabulae declarant, in quibus tam multis tamque variis ab ultima antiquitate repetitis tria vix.	Tritani, centurionum, praeclarorum hominum ac primorum signiferumque, maluisti dici. Graece ergo praetor Athenis, id.	a060f7cd-dcec-a316-19af-c17e39405946	88ece235-12ad-1182-2b71-3293dbc4ff0a	Resolved	788201bf-293d-cce9-f450-4587bee8c48d
27626625-40a8-5b8b-a76e-6f20f764dcb9	2024-07-28 19:34:57	2023-10-13 08:02:33	Utriusque cum multa non probo, tum illud in primis, quod, cum in rerum natura duo.	Cumque nihil impedit, quo minus omnes mea legant. Utinam esset ille Persius, Scipio vero et Rutilius multo etiam magis, quod, cuius in animo versatur, numquam sinit.	08876f42-8e6b-bee6-f09b-c445d4edf545	59bb140c-94d1-a296-ee4f-5bf5b86516dc	New	42c168df-da9c-9817-8d62-f2a985b7cf03
09766ace-bf65-8f60-4ab8-6a3e6d4dcb0f	2024-01-09 20:00:33	2024-07-09 06:54:24	Me de virtute misisti. Sed ex eo est consecutus? -- Laudem et caritatem, quae sunt vitae sine metu vivere. Quae est enim contra Cyrenaicos satis acute, nihil ad iucunde vivendum reperiri posse.	Malit quam deserere ullam officii partem, ad ea, quae senserit ille, tibi non vera videantur. Vide, quantum, inquam, fallare, Torquate. Oratio me istius philosophi non offendit; nam et complectitur verbis, quod vult, et dicit plane, quod intellegam; et tamen in.	fd6d801e-9fd3-0100-c455-ef5c21daf273	b88498f2-759d-b568-3b37-842aebcba416	In Progress	6eca4019-a2f3-d0fe-cf94-c69ad8a68291
65c45803-d441-f9d8-bd2c-2af293a9f3f8	2024-05-12 15:05:17	2023-10-16 15:52:40	Si dicent ab illis has res esse tractatas, ne ipsos.	Ab eadem illa individua et solida corpora ferri deorsum suo pondere ad lineam, numquam fore ut hic noster labor in varias reprehensiones incurreret. Nam quibusdam, et iis servire, qui vel utrisque litteris uti velint vel, si suas habent, illas non.	602ce54a-d71f-c1e1-e6f8-37894313e9f6	e05dc6d7-9cb7-df7a-f9bf-ce9834b0e41b	Assigned	a8855e8c-d01f-a593-ac43-9eb0922a66a0
099483f9-2897-e0de-bd98-9667d5eb97ec	2024-03-28 14:47:48	2023-10-21 18:24:54	Orationis vel copiosae vel elegantis ornatus defuit?	Facillimis ordiamur, prima veniat in medium Epicuri ratio, quae plerisque notissima est. Quam a nobis sic intelleges eitam, ut ab ea nullo modo poterimus sensuum iudicia defendere. Quicquid porro animo cernimus, id omne oritur a sensibus; qui si omnes.	658f8c88-19ef-b7eb-5172-b7dcb9ca377a	1f760bdf-31d7-14f6-edee-c2b9a7a6e5fa	Assigned	6b0884fb-f2c4-fcc8-bf51-0c906ffb5b05
1d895332-bd11-947d-d8fc-21f5a2c98ea5	2024-07-28 06:12:02	2023-10-15 13:01:16	Esse deterritum. Quae cum dixissem, magis ut illum provocarem quam ut ipse constituit, e regione.	Cyrenaicisque melius liberiusque defenditur, tamen eius modi tempora incidunt, ut labore et dolore magna aliqua. Ut enim ad sapientiam perveniri potest, non paranda nobis solum ea, sed fruenda etiam sapientia est; sive.	b0507f00-6703-87e4-08ca-016852d86c06	b7120f97-9420-169f-ae88-0e15f90610e8	New	7a5f4d28-366e-6743-a993-a1fe1c554af0
191e63b7-8b2b-87ec-4dfc-6e829a838b84	2024-06-02 01:11:08	2024-09-29 05:15:32	Quam magnum sit fictae veterum fabulae declarant.	Confectum tantis animi corporisque doloribus, quanti in hominem maximi cadere possunt, nulla spe proposita fore levius aliquando, nulla praeterea neque praesenti nec expectata voluptate, quid eo miserius dici aut fingi potest? Quodsi vita doloribus referta maxime fugienda est, summum profecto malum est vivere cum.	d798e72c-f28d-2717-f8d0-03eaf24caee1	a9b46d58-b3f7-f518-d524-d22b3c63a2f9	Resolved	eb7a5cbf-ec24-e712-acc7-68cb4aa0468a
52566225-bfd9-79a8-157b-4b2fca7ebde2	2024-04-28 10:27:34	2024-09-08 13:48:33	Natura ipsa iudicari. Ea quid percipit aut quid iudicat, quo aut petat aut fugiat aliquid, praeter voluptatem et parvam et non necessariam et quae sequamur et quae fugiamus refert omnia. Quod quamquam Aristippi est a Cyrenaicisque melius liberiusque defenditur, tamen eius modi tempora incidunt.	Memoriter, tum etiam erga nos amice et benivole collegisti, nec me tamen laudandis maioribus meis corrupisti nec segniorem ad respondendum reddidisti. Quorum facta quem.	64a7ec4a-bd48-a5fd-d60f-5fc2894f585b	a9b46d58-b3f7-f518-d524-d22b3c63a2f9	In Progress	69e20fbc-06bb-ea8b-3f43-8d38f18953b3
40ee3d60-2218-6a51-663d-0b44749a546e	2024-01-27 18:00:12	2023-11-08 03:59:08	Explicabo, voluptas ipsa quae qualisque sit, ut tollatur error omnis imperitorum intellegaturque ea, quae audiebamus.	Dominorum domus; quo minus omnes mea legant. Utinam esset ille Persius, Scipio vero et Rutilius multo etiam magis.	ad46169b-6ac8-da62-04d6-83cd5057073c	f13b38a9-9c31-ca29-b6fe-569fbd3049b3	In Progress	a8603adf-76d4-fa50-dec5-5c796c68ea17
c5ac33d9-f97e-3a62-9dbf-15773b13e84c	2024-03-27 00:32:40	2024-09-20 06:30:56	Isti curiosi, quos offendit noster minime nobis iniucundus labor. Iis igitur est non miser. Accedit etiam mors, quae quasi titillaret sensus, ut ita ruant itaque turbent, ut earum motus et impetus quo pertineant non intellegamus, tu tam egregios viros censes tantas res gessisse sine causa? Quae fuerit causa, mox videro; interea hoc.	Rerum necessitatibus saepe eveniet, ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum defuturum, quas natura non depravata desiderat. Et quem ad modum, quaeso, interpretaris? Sicine eos censes aut in armatum hostem impetum fecisse aut in poetis.	b6b61a56-7fd6-acbb-c062-f03624e1316c	3bfc1bfd-c79b-a2c7-35dd-d0b9c04254bb	On Hold	8ed88161-56ed-751c-0295-9b63a81a9261
763bd84b-a4da-bbd1-1cf4-e178f133d700	2024-01-05 13:50:41	2024-05-07 08:07:11	Vendibiliora, haec uberiora certe sunt. Quamquam id quidem facio provocatus gratissimo mihi libro, quem ad modum temeritas et libido et.	Dixi, sole ipso illustriora et clariora sunt, si omnia deorsus e regione inferiorem locum petentium sine causa dicere, -- et illum motum naturalem omnium ponderum, ut ipse constituit, e regione ferrentur et, ut modo docui, cognitionis regula et iudicio.	78778af7-44fe-eadc-991f-867056c8af93	59bb140c-94d1-a296-ee4f-5bf5b86516dc	Resolved	6dc03670-c3ea-0dd0-95cd-cb5c7b1a939c
b7e6e800-396c-38eb-8259-4791b5fdf5ee	2024-04-17 03:34:37	2024-02-29 16:33:34	Esse et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum defuturum, quas natura non depravata desiderat. Et quem ad modum eae semper voluptatibus inhaererent, eadem de amicitia dicenda sunt. Praeclare enim Epicurus his paene verbis: 'Eadem', inquit, 'scientia confirmavit animum, ne quod aut sempiternum aut diuturnum timeret malum, quae perspexit in hoc ipso vitae.	Autem inanes sunt, iis parendum non est. Nihil enim desiderabile concupiscunt.	af02e9ad-7811-f56e-0788-e1fc8d753071	835855b7-48e2-194e-6b7e-436075b012a7	Assigned	8717f9ba-0af0-f8fd-c470-3f70eeeaf481
85ccb9f5-3d83-b19b-dad9-c24868bea6d6	2024-07-05 16:18:08	2024-04-02 11:28:45	Inanes divitiarum, gloriae, dominationis, libidinosarum etiam voluptatum. Accedunt aegritudines, molestiae, maerores, qui exedunt animos conficiuntque curis.	Invidia verbi labefactetur oratio mea --. Nam cum ignoratione rerum.	3ca9fa80-86b0-5397-8851-da011404252b	69011f7d-ab78-d570-6ef9-9a9f858be2da	In Progress	6dc03670-c3ea-0dd0-95cd-cb5c7b1a939c
08773d0e-24aa-72fd-22d9-2348518f3429	2024-05-01 08:17:14	2024-04-15 16:28:55	Maledici, monstruosi, alii autem etiam amatoriis levitatibus dediti, alii petulantes.	Nos veriusque quam Stoici. Illi enim negant esse bonum iucunde vivere. Id qui in una virtute ponunt et splendore nominis capti quid natura postulet.	78778af7-44fe-eadc-991f-867056c8af93	6922cc8b-cc03-9647-9854-d478746cbc33	New	1f7ec3f0-ba9a-6294-70dc-085776867f9c
04a77663-fc9e-e675-eedf-2018bdb3e120	2024-02-12 14:39:50	2024-02-23 00:48:23	Reliquaque eiusdem generis et legimus libenter et legemus --, haec, quae.	Quantaque amoris conspiratione consentientis tenuit amicorum greges! Quod fit etiam nunc ab Epicureis. Sed ad haec, nisi molestum est, habeo quae velim. An me.	aa6f70c3-5acf-9fc1-8df4-5cbea21d3a62	534df3a2-fbc9-cc82-e64c-51c090c12b29	On Hold	a367f358-29d7-2ea3-c025-fb12f78f7924
6ab51ccc-9378-2154-a9bd-c6eb72849a31	2023-11-05 22:16:03	2023-11-15 17:11:35	Ipsum autem nusquam. Hoc Epicurus in voluptate ponit, quod summum bonum consequamur? Clamat Epicurus, is quem vos nimis voluptatibus esse deditum dicitis; non posse iucunde vivi, nisi sapienter, honeste iusteque vivatur, nec sapienter, honeste, iuste, nisi iucunde. Neque enim civitas in seditione beata esse potest nec in malis dolor, non existimant oportere.	Voluptatem. Ut enim mortis metu omnis quietae vitae status perturbatur, et ut succumbere doloribus eosque humili animo inbecilloque ferre miserum est, ob eamque causam propter voluptatem et parvam.	033d0582-abf0-c0b8-c9fb-a4d455626356	722a2243-b5ee-fba0-2400-95b939f87392	Assigned	aea7f8e3-952e-86ef-59f1-bb06fa4b28ec
bad8b7b9-54b3-9f9a-2951-ff94080e6d65	2024-02-29 00:19:18	2024-07-23 06:27:40	Quam autem ego dicam voluptatem, iam videtis, ne invidia verbi labefactetur oratio mea --. Nam cum solitudo et.	Aiunt hanc quasi naturalem atque insitam in animis inclusae inter.	8f7b1456-6da2-5830-72af-6c1b26c711c7	3a4ffaf2-77ac-06c4-e5b7-806c5dd98a9c	New	ce59e370-3406-36d6-0d8d-c146f734d94b
5a506983-f287-8cdc-fb11-7acba71a7b49	2024-04-23 15:25:39	2023-11-14 13:49:10	Modo non impediri rationem amicitiae, si summum bonum in voluptate velit esse, quam nihil.	Quidem ita dicit, omnium rerum, quas ad beate vivendum se ipsa allicit nec patientia nec assiduitas nec vigiliae nec ea ipsa, quae tibi probarentur.	b4e34189-0cf7-219f-c783-538b96fb2740	faf5c017-5cfc-8135-6e95-7f02e6fcb6e6	Resolved	99c10813-c9bc-6f97-0337-6fabf4aae81e
4420e6ef-ef03-90f2-eddd-73db62badd89	2024-10-08 21:37:09	2023-11-17 20:51:46	Vitam omnem continent, neglegentur? Nam, ut sint illa vendibiliora, haec.	Latinis litteris mandaremus, fore ut hic noster labor.	199e674b-98fa-1441-2c8d-b91bcf25133f	6ce4068b-b8ef-6410-e92e-9ea044daacee	In Progress	f221fcd0-c52f-553e-bd89-64cf69c86ca4
23858573-6ead-3006-5d3c-4ea16695b143	2024-03-25 22:46:24	2023-12-09 03:41:09	Splendido nomine, virtutem autem nixam hoc honesto nullam requirere voluptatem atque ad beate vivendum se ipsa esse contentam. Sed possunt haec quadam ratione dici non necesse est. Tribus igitur modis.	De caelo est ad quiete vivendum, caritatem, praesertim cum omnino nulla sit causa peccandi.	30c30ce2-0245-e36c-3ec5-3345e9e34e2f	84921767-114e-2a99-49c5-d13e8f77ec27	Assigned	aa5e01f8-479d-99a8-ac3d-4b8ae043e2e4
849029de-a5f6-ac1e-8346-b6aeee60a35b	2024-07-14 13:31:43	2023-11-11 13:10:54	Est sola, quae nos exhorrescere metu non sinat. Qua.	Idcirco et hoc ipsum efficitur in amicitia, et amicitia cum voluptate vivatur. Quoniam autem id est vel summum bonorum vel ultimum vel extremum .	6d6dda53-500d-a5c0-2472-f353bd8f11c1	b92ed5ea-5029-9fad-daad-1d25f985db12	New	8d9d5b42-4aa5-e93d-eb12-d1cc093265c7
e061d98f-88f2-c875-ba8b-e3156a05da4e	2024-03-16 02:41:52	2024-05-19 19:27:59	Et contrariis studiis consiliisque semper utens nihil quieti videre, nihil tranquilli potest. Quodsi vitam omnem continent, neglegentur? Nam, ut sint opera, studio, labore meo.	Omnium rerum, quas ad beate vivendum se ipsa esse contentam. Sed possunt haec quadam ratione dici non modo fautrices fidelissimae, sed etiam cogitemus; infinitio ipsa, quam apeirian vocant, tota ab illo inventore veritatis et quasi architecto.	66c9caa4-6b57-9d97-906f-07f5f28e287a	b4020a50-cba1-0bb0-9218-07c5aec9fe2d	New	6dc03670-c3ea-0dd0-95cd-cb5c7b1a939c
2bb4c95c-5052-b93e-58eb-4fc815db7a9b	2024-05-07 11:00:12	2024-08-07 09:51:53	Si vita suppetet; et tamen, qui diligenter haec, quae vitam omnem continent, neglegentur? Nam, ut sint illa vendibiliora, haec uberiora certe.	Compluribus permulta dicantur, cur nec voluptas in bonis sit numeranda nec in discordia dominorum domus; quo minus animus.	92f225a5-94e1-9a62-99d1-33227da0d19c	ff794092-d8db-236f-9c6f-699d9ca6c3b6	Assigned	3f405e3d-599f-cb3b-a9b2-e2b8f8985bb0
d6b86bab-a228-b1cf-635a-010ade37e6de	2023-10-11 14:26:52	2024-06-27 09:49:04	Modis video esse a nostris non legantur? Quamquam, si plane sic verterem Platonem aut Aristotelem, ut verterunt nostri poetae fabulas, male, credo, mererer de meis civibus, si ad eorum cognitionem divina illa ingenia transferrem. Sed id.	Et solida corpora ferri deorsum suo pondere ad.	ffb4e5cd-424e-e2a4-16f5-e4e2667ace5c	4e6f22c6-2486-4d38-d4f4-599776479d63	In Progress	3f405e3d-599f-cb3b-a9b2-e2b8f8985bb0
07651a3f-57ff-8077-f759-3ea5bad31e88	2024-06-08 16:05:44	2024-09-27 13:34:56	Quia bene navigandi rationem habet, utilitate, non arte laudatur, sic sapientia, quae ars vivendi putanda est, non satis politus iis artibus, quas qui tenent, eruditi appellantur -- aut.	Nobis iniucundus labor. Iis igitur est difficilius satis facere, qui se Latina scripta.	3a1084ec-b697-b7c0-64ea-4af4528c5e96	669589b7-1180-c01e-687d-fff150709e47	In Progress	8462bab8-9eeb-5542-82c3-cb1d37c0a014
ea02d6f2-4cb2-12aa-c1c2-32ce66bd93a6	2024-02-15 00:17:35	2024-04-17 11:18:19	Regula et iudicio ab eadem illa individua et solida corpora ferri deorsum suo pondere ad lineam, numquam fore ut atomus altera alteram posset attingere.	Impetus quo pertineant non intellegamus, tu tam egregios viros censes tantas res gessisse sine causa? Quae fuerit causa, mox videro; interea hoc tenebo, si ob aliquam causam ista, quae sine dubio praeclara sunt, fecerint, virtutem iis per se ipsa esse.	b04cd62f-d0a6-f0a4-ec7d-45e4586a680d	effe6dce-661f-cbcd-bf91-1b30027372fc	New	5b988155-7469-6252-dab7-bcc41394920a
fb03793c-eece-a91d-c457-86cb1398c7ab	2024-02-29 16:49:19	2024-10-05 19:23:15	Quae essent et naturales et necessariae, alterum, quae vis sit, quae quidque efficiat, de materia disseruerunt, vim et causam efficiendi reliquerunt. Sed hoc commune.	Poetis evolvendis, ut ego et Triarius te hortatore facimus, consumeret, in quibus.	33de396b-1809-3649-4735-07b69d8491af	1f596437-be02-5a5e-32f3-5ab5b1cff9f6	New	163a4480-1335-f3f0-2cd1-c5c76cd06cc1
58a55c34-07c8-5760-0306-d3a3f3e459b3	2023-10-18 19:23:39	2024-06-30 22:57:00	Nos veriusque quam Stoici. Illi enim negant esse bonum iucunde vivere. Id qui in una virtute ponunt et splendore nominis.	Autem hanc omnem quaestionem de finibus.	70a81c90-0c24-d3bd-c8a9-8d35a9367115	6922cc8b-cc03-9647-9854-d478746cbc33	On Hold	2f7ff7be-e3f2-c495-2c9e-9228cf19f269
887328de-f701-422a-3c69-f5c306b5e637	2024-04-24 03:23:53	2024-10-01 12:28:53	Praetulerit ius maiestatis atque imperii. Quid? T. Torquatus, is qui consul cum Cn. Octavio.	Despicationes adversantur voluptatibus, sic amicitiae non modo fautrices fidelissimae, sed etiam cogitemus; infinitio ipsa, quam apeirian vocant, tota ab illo est.	a1f0fc17-d38b-3b75-1696-90c33c93e75d	515ef09a-9de7-4187-7ffd-7432012f2c8a	New	89b543df-5315-b9e6-2cc7-97c739d60921
862f82a8-4f48-d9c4-5cb5-5c0c17e64c58	2024-07-17 02:03:26	2024-05-08 02:10:57	Quo efficiantur ea, quae sensum moveat, nulla successerit, eoque intellegi potest quanta voluptas sit non dolere. Sed ut iis bonis erigimur, quae expectamus, sic.	In voluptatis locum dolor forte successerit, at contra gaudere nosmet omittendis doloribus, etiamsi voluptas ea, quae dicta sunt ab iis quos probamus, eisque nostrum iudicium et nostrum scribendi ordinem adiungimus, quid habent, cur Graeca anteponant iis.	5e3a4b53-751b-4d05-d323-94c1aff0bb49	9d85c6b0-4d02-21fc-33d9-d6bf335e624b	Assigned	84104518-2dce-4552-e75a-62566a17d9b5
c6045b3b-04e8-f357-46e6-74527aa0341e	2024-04-23 00:25:26	2024-07-24 07:40:27	Depravare videatur. Ille atomos quas appellat, id est incorruptis atque integris testibus, si infantes pueri, mutae etiam bestiae paene loquuntur magistra ac duce natura nihil esse prosperum nisi voluptatem, nihil asperum nisi.	Quid est, cur nostri a nostris non legantur? Quamquam, si plane sic verterem.	7147d009-6fec-1ed0-eada-28469ecb42c4	49e072b6-94ab-0195-b106-351c0f199eec	On Hold	23803d86-b4d2-a4b5-768a-0ea24dc2f6a4
1399a431-dde2-41c1-b1ac-471ee388ccec	2024-10-13 20:26:43.234625	2024-10-13 20:26:43.234625	a	a	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
690e15f4-4817-48d1-bed4-2680e6ca408a	2024-10-13 20:27:16.07659	2024-10-13 20:27:16.07659	aaa	aaaoeuoeueou cluster prod 2	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
20d09f10-a69e-4163-b673-35e55d26a265	2024-10-13 20:28:19.837731	2024-10-13 20:28:19.837731	aoeuaoeu	aoeuaoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
a68565d6-900d-4abb-96c2-361f78abd84e	2024-10-13 20:32:51.661958	2024-10-13 20:32:51.661958	a	aaoaeu	7f5d0b1d-78ba-6463-0d82-67273d1cd09b	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
febb248a-c781-4d01-b1d8-a7b70c6b27fc	2024-10-13 21:01:17.979179	2024-10-13 22:06:21.212102	oauoeu	aoeuaoeu eee aoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	f708550a-0fe1-4ee9-b27e-ed5db213ffff	Assigned	8dbfdfdf-9ce9-123b-3718-bdb79bcac1b7
35500a24-ed7d-49c9-8a8c-763a46068969	2024-10-13 22:56:07.480599	2024-10-14 13:00:28.862156	Fix issue with API cluster	more details here	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	On Hold	21ca2ec5-61b0-d8e1-dd1d-f334a4fa443b
54d676bb-f134-489a-9e2e-3b6ec9cd2c39	2024-10-13 20:38:05.784352	2024-10-13 20:47:16.919176	aaaoeuaoeu	aoeuaoeuaoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	In Progress	94daac1b-4f98-fa5a-2d31-c997b2859de4
3f72c1a6-d59a-428d-b001-f2592a25b2e0	2024-10-13 20:52:27.181647	2024-10-13 20:52:27.181647	aoeu	aoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
477bfcaf-bc47-4cd3-82e8-151b5c37b828	2024-10-13 20:53:42.139551	2024-10-13 20:53:42.139551	aoeuaoeu	aoeuaoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
7b99083a-3835-47d5-b65e-cf01bd4cf86e	2024-10-13 20:54:41.803317	2024-10-13 20:54:41.803317	aoeu	aoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
a102078d-9054-4276-9b90-27f91a776047	2024-10-13 20:57:09.765819	2024-10-13 20:57:09.765819	aoeu	aoeu	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	New	\N
bb2c194b-441a-40e2-8597-e8af23c1ac9e	2024-10-14 13:01:09.133178	2024-10-14 13:01:09.133178	Replica host status	Replica not communicating with gateway	e2ff0f90-be7c-125d-ede2-f50154f3822b	b3eef82e-5c3c-4876-95be-a8f6acb0cd74	New	\N
f71f4fcc-3a1a-4cee-b734-45fe536d0b1c	2024-10-13 20:59:58.356077	2024-10-13 22:54:58.53118	aaa	longer description. and now even longer	3e9b3285-2f74-10d9-c529-f46fb5c2024c	5fcdebd5-cc0d-402c-b9e1-06f94d72a5b7	Resolved	14d8e1a9-1967-d96d-e2c3-6418d974c060
ca09f8e3-6efd-424f-a197-ae6fc5ff139a	2024-10-13 21:00:41.041029	2024-10-13 22:06:11.047515	ooo	oo	3e9b3285-2f74-10d9-c529-f46fb5c2024c	00b73e30-250e-5127-059a-285085d02814	New	94daac1b-4f98-fa5a-2d31-c997b2859de4
87ab2aa0-b619-4b6c-83ad-38a0d365d5cd	2024-10-14 13:02:30.115194	2024-10-14 13:02:53.214233	Failover issue	Failover not always happening	6d6dda53-500d-a5c0-2472-f353bd8f11c1	b3eef82e-5c3c-4876-95be-a8f6acb0cd74	Assigned	8e17ee9f-82f0-09e0-5918-e828b31a9265
\.



COPY public.users (id, created_at, updated_at, first_name, last_name, email) FROM stdin;
5faba39f-64fe-4805-8365-0a91bb396477	2024-09-06 00:00:00	2024-09-06 00:00:00	admin	admin	john@gmail.com
7c132f12-40b8-0029-43de-f5d464122cc3	2024-08-24 15:37:47	2024-06-29 03:37:34	Vita	Faltin	vita.faltin@yahoo.com
bd08d4b7-3eb3-213e-274a-d389bbf11145	2023-12-21 11:56:24	2024-02-06 10:08:23	Waylon	Stoffers	waylon.stoffers@live.co.uk
04144b42-aab1-883f-0823-d57e77e5c2bd	2024-01-04 00:29:09	2024-03-16 14:28:33	Chickie	Hazlehurst	chickie.hazlehurst@yahoo.com
54e9f34a-9117-04aa-1fa5-947f77c88324	2024-05-03 00:03:41	2024-05-24 05:44:21	Taylor	Van Dale	taylor.van.dale@comcast.net
6eca4019-a2f3-d0fe-cf94-c69ad8a68291	2024-05-13 20:17:25	2024-01-13 23:45:16	Munroe	Sallenger	munroe.sallenger@yahoo.com
b28c5a1d-e325-1f0a-67b3-3239a9ede311	2023-12-25 23:54:30	2023-11-14 05:01:56	Mauricio	Wagenen	mauricio.wagenen@orange.fr
4c59113d-9dc0-b754-4ee0-c76add90488e	2024-01-10 10:34:43	2024-09-12 06:22:40	Dorey	Betts	dorey.betts@hotmail.com
14d8e1a9-1967-d96d-e2c3-6418d974c060	2024-06-21 20:41:06	2024-05-23 19:25:17	Hedvig	Blazewski	hedvig.blazewski@gmail.com
d36adaef-e572-2f65-793a-79cbe2973bcd	2024-01-04 00:09:51	2023-10-10 01:39:39	Yale	Lundberg	yale.lundberg@yahoo.fr
23e056e9-788d-155c-640e-625dff203b2a	2024-02-15 03:07:59	2023-12-16 01:27:51	Baldwin	Cordero	baldwin.cordero@aol.com
99c10813-c9bc-6f97-0337-6fabf4aae81e	2024-08-23 05:29:11	2024-06-26 17:18:18	Jerome	Redpath	jerome.redpath@gmail.com
8d9d5b42-4aa5-e93d-eb12-d1cc093265c7	2024-08-01 17:25:26	2024-03-16 00:31:34	Ferguson	Cockroft	ferguson.cockroft@gmail.com
197f65b3-76e8-4017-624b-21d733a776c1	2024-03-08 23:16:11	2024-08-19 20:50:55	Mannie	Trenholme	mannie.trenholme@yahoo.com.br
d7edfdd0-b755-e446-29cc-cc83b3cdcf66	2024-08-18 02:11:54	2024-05-15 00:28:15	Arliene	Longhorne	arliene.longhorne@yahoo.com
8dbfdfdf-9ce9-123b-3718-bdb79bcac1b7	2023-12-12 20:46:28	2024-09-26 15:50:50	Aaren	Barltrop	aaren.barltrop@ymail.com
9d93b06f-ea9f-906e-5448-baeb572b06ce	2024-01-09 18:36:15	2024-06-01 04:25:02	Olwen	Broster	olwen.broster@hotmail.es
23803d86-b4d2-a4b5-768a-0ea24dc2f6a4	2024-06-07 21:40:48	2024-07-22 14:32:46	Ferrel	Trotman	ferrel.trotman@hotmail.com
cd80e0e8-405c-6d96-5a39-35f6cc9e7e50	2024-06-22 14:41:39	2024-03-07 10:48:06	Araldo	Steagall	araldo.steagall@hotmail.com
0ce4a2dc-c9c0-f712-17f6-4791398f0933	2024-04-15 10:35:38	2024-05-08 13:59:38	Geri	Tumpane	geri.tumpane@hotmail.com
9c70b291-1f05-d368-862e-4ae6f3d12ca3	2023-10-13 22:46:06	2024-03-25 17:05:45	Rudd	McCollum	rudd.mccollum@msn.com
788201bf-293d-cce9-f450-4587bee8c48d	2024-01-31 22:30:27	2024-01-11 18:55:25	Webb	Boulds	webb.boulds@msn.com
8340ebb6-29e0-f7cb-373c-fc42de6e3d51	2024-04-23 01:47:59	2024-07-27 20:07:52	Corny	Benardet	corny.benardet@yahoo.com
33caaa90-0640-23a3-8742-0e05f97b8581	2024-04-08 20:49:08	2024-02-14 13:04:12	Reginald	Mulbery	reginald.mulbery@yahoo.com
1c5893a5-92cb-b792-c6b4-84461f749ed1	2023-12-09 13:24:36	2024-08-05 07:40:29	Zarah	Hebblethwaite	zarah.hebblethwaite@orange.fr
357e8baa-8b8e-1dee-5492-c7fdb2ca42ff	2024-02-15 23:17:33	2024-07-10 07:31:02	Jerrold	Rahill	jerrold.rahill@yahoo.com
a367f358-29d7-2ea3-c025-fb12f78f7924	2024-04-25 12:21:08	2024-09-01 07:38:23	Anabal	Willden	anabal.willden@gmail.com
82c84588-62e2-6a05-3bd4-df8618e9b1f7	2024-03-30 05:19:54	2024-07-28 15:29:48	Care	Savidge	care.savidge@hotmail.com
299d02c8-2d84-e32f-5466-03f80e76a35f	2023-11-29 16:40:28	2024-03-15 13:31:42	Kirsteni	Huot	kirsteni.huot@gmail.com
5b988155-7469-6252-dab7-bcc41394920a	2024-07-27 09:42:06	2023-12-15 07:34:09	Claire	Clery	claire.clery@live.fr
0a721c4d-e7cb-0525-0694-4268e1e731dd	2024-07-09 14:16:20	2024-06-19 00:17:41	Blondell	Sysland	blondell.sysland@yahoo.com
aa5e01f8-479d-99a8-ac3d-4b8ae043e2e4	2024-07-22 02:28:39	2024-01-05 02:21:06	Giulietta	Streather	giulietta.streather@yahoo.com
84104518-2dce-4552-e75a-62566a17d9b5	2024-08-26 20:41:17	2024-09-10 03:48:34	Alexei	Hargey	alexei.hargey@msn.com
c47aac56-0671-58ce-527d-1c4957d47d25	2024-03-04 23:20:44	2023-12-20 11:52:18	Joleen	Brushfield	joleen.brushfield@gmail.com
e08475a5-f09d-fa1e-1006-4dca3acc54b6	2023-11-22 16:05:23	2023-10-27 16:15:46	Susana	Keeri	susana.keeri@yahoo.com
f897f55c-4f77-1c57-f838-3981895ffaea	2024-05-15 03:50:21	2024-08-06 18:07:06	Eugenio	Gorton	eugenio.gorton@aol.com
12efc06b-c25d-79c6-687b-c695935d6d7f	2024-05-09 08:26:50	2024-07-31 00:38:36	Bruno	Ivermee	bruno.ivermee@gmail.com
1f7ec3f0-ba9a-6294-70dc-085776867f9c	2024-07-22 17:19:44	2024-10-05 11:10:48	Olivero	Charlotte	olivero.charlotte@yahoo.com
69e20fbc-06bb-ea8b-3f43-8d38f18953b3	2024-03-22 16:36:52	2023-12-22 22:11:32	Marchelle	Hoggan	marchelle.hoggan@yahoo.com
25225dbd-5912-1e4b-f1cd-c38441f35f37	2024-08-03 21:18:00	2024-09-13 15:21:13	Morey	Roubay	morey.roubay@rocketmail.com
f7845a09-230d-6cda-e45c-bc47c59944cf	2023-12-08 11:34:29	2023-12-23 13:44:16	Ransell	Ewen	ransell.ewen@hotmail.com
6fd47d5a-4ce1-79c8-6dbf-243414f6398b	2024-07-18 07:49:44	2024-01-31 10:35:22	Tina	Fitzsymon	tina.fitzsymon@yahoo.com
163a4480-1335-f3f0-2cd1-c5c76cd06cc1	2024-03-01 00:11:53	2024-05-23 14:07:00	Jere	Pesticcio	jere.pesticcio@gmail.com
3ba68195-b702-4196-40ed-db47409f79f5	2024-04-09 08:18:02	2023-12-23 06:05:52	Joannes	Kilbane	joannes.kilbane@wanadoo.fr
2f919ce3-fa2e-e147-3811-79c59d0d6905	2024-07-26 22:02:26	2024-07-25 07:06:28	Kerrill	Brolechan	kerrill.brolechan@gmail.com
d27e47ca-90a9-666d-f964-5ae5847d66f2	2023-11-13 20:22:39	2023-11-04 00:26:54	Cassi	Pashler	cassi.pashler@hotmail.co.uk
c1d5b91c-322e-2e05-c6fd-4cd1fdbf3a45	2024-05-31 21:14:12	2023-12-29 09:38:29	Fabian	Langer	fabian.langer@yahoo.com
0f7b6d31-e11d-5ed3-8d75-6a01cee95b24	2024-05-09 03:01:05	2024-07-04 11:40:12	Hieronymus	Nel	hieronymus.nel@yahoo.com
0116e874-1370-cca6-0283-06f4ec793d1b	2023-12-23 15:33:05	2023-12-19 06:01:23	Marcello	Burbudge	marcello.burbudge@yahoo.com
abc8b651-f42b-d696-95d9-e3c73699640a	2024-09-09 00:32:11	2023-12-02 23:06:23	Sheela	Gilliat	sheela.gilliat@planet.nl
1f2543bf-dad1-b947-ad9b-1e10912de3cc	2024-08-25 09:04:52	2024-01-09 22:06:21	Dillon	Twydell	dillon.twydell@wanadoo.fr
08710c84-8198-2463-c057-c6071906df0e	2024-01-08 16:41:58	2023-12-29 23:19:27	Barry	Romei	barry.romei@hotmail.com
42c168df-da9c-9817-8d62-f2a985b7cf03	2024-01-10 23:21:06	2024-04-17 03:10:21	Gabriele	Johnke	gabriele.johnke@bigpond.com
8ed88161-56ed-751c-0295-9b63a81a9261	2024-05-19 20:11:26	2024-04-26 17:55:28	Rossy	Lippingwell	rossy.lippingwell@yahoo.com
3f405e3d-599f-cb3b-a9b2-e2b8f8985bb0	2024-04-06 07:14:57	2024-04-07 03:46:33	Yves	Potbury	yves.potbury@gmail.com
e282eb6c-fa4e-ebe1-af16-b3a838b38f8c	2024-06-25 19:50:54	2023-12-02 21:42:32	Elsinore	Vanns	elsinore.vanns@yahoo.com
9809f6c4-a1bf-1b67-d499-df5bf7c135be	2024-09-13 15:49:12	2024-04-04 05:57:17	Lilias	Hitscher	lilias.hitscher@yahoo.com
8717f9ba-0af0-f8fd-c470-3f70eeeaf481	2024-04-25 17:16:32	2024-07-03 10:20:59	Natal	Mercy	natal.mercy@yahoo.com
a8855e8c-d01f-a593-ac43-9eb0922a66a0	2024-03-20 16:39:55	2024-09-18 13:48:36	Hercules	Thredder	hercules.thredder@comcast.net
6ad0f1d1-7e42-8f43-494b-1b9f53e59027	2024-04-29 05:25:23	2024-03-18 00:44:52	Catlin	Dhenin	catlin.dhenin@hotmail.co.uk
770c94e7-6081-ea4d-b3be-b06c69c9e3fc	2024-07-14 13:49:52	2024-04-12 10:48:50	Townie	Mant	townie.mant@yahoo.com
7ec75e48-3cca-4ab3-48d3-0d8c8d6233c4	2024-05-04 00:50:48	2023-10-10 18:34:03	Katalin	Lipscomb	katalin.lipscomb@gmail.com
b01d292f-8c87-e144-5218-34d6a5e82539	2024-04-05 00:26:01	2024-10-02 08:10:20	Annabel	Petrello	annabel.petrello@live.com
a0d866be-071e-776b-f869-7cc01375b73e	2023-12-02 12:18:51	2024-08-26 08:09:48	Tobe	Bullimore	tobe.bullimore@yahoo.com.ar
f221fcd0-c52f-553e-bd89-64cf69c86ca4	2024-05-17 10:19:43	2023-10-23 02:34:06	Casar	Bramwich	casar.bramwich@yahoo.com
7a5f4d28-366e-6743-a993-a1fe1c554af0	2024-08-12 08:22:49	2024-08-15 20:55:45	Ashien	Patise	ashien.patise@hotmail.com
aea7f8e3-952e-86ef-59f1-bb06fa4b28ec	2024-02-23 08:00:56	2024-07-26 20:48:42	Riordan	Spurman	riordan.spurman@hotmail.com
bdd06817-bacf-8252-e8e4-bffe89f7670b	2024-08-02 17:10:44	2024-04-26 13:05:26	Binnie	Redgrove	binnie.redgrove@gmail.com
e1f81b4c-bab0-7359-7928-bb980ea5c495	2024-08-10 23:53:16	2023-10-27 11:32:50	Alexina	Cullon	alexina.cullon@yahoo.com
344f0add-9a50-5a49-8998-21e36dcf0c58	2023-11-15 01:54:33	2024-08-30 03:52:27	Laughton	Chopy	laughton.chopy@hotmail.com
21ca2ec5-61b0-d8e1-dd1d-f334a4fa443b	2024-09-16 20:57:56	2024-01-20 14:55:21	Terrijo	Harland	terrijo.harland@yahoo.com
7440acd6-a0bb-1f71-57a1-99e271efdbe8	2024-04-12 14:39:47	2024-01-30 11:07:17	Tremain	Bulter	tremain.bulter@aol.com
7a5623f3-f15c-9bed-05fc-1308b1b28a26	2024-05-18 18:34:09	2024-02-12 07:36:29	Justine	Kollach	justine.kollach@cox.net
313d36bb-733e-9419-14a5-3219392948e5	2024-04-30 14:24:38	2024-05-17 12:50:19	Carroll	Hess	carroll.hess@gmail.com
eb9be367-874b-f3d2-7e15-7c1abb5b3f37	2024-04-14 04:38:56	2023-12-06 20:47:29	Anton	Harling	anton.harling@hotmail.com
1687710d-00e7-5a70-60f8-90c82a664516	2024-09-16 04:50:21	2023-10-30 13:01:08	Pate	Teenan	pate.teenan@yahoo.com
6dc03670-c3ea-0dd0-95cd-cb5c7b1a939c	2023-10-29 16:29:03	2023-10-24 15:57:59	Bonni	Sneller	bonni.sneller@yahoo.com
b5ccf8d1-c14c-d393-9cf2-f47bf1d068ed	2023-11-05 00:18:52	2024-01-12 10:29:24	Bliss	Plowright	bliss.plowright@gmail.com
8e17ee9f-82f0-09e0-5918-e828b31a9265	2024-01-14 08:06:58	2024-09-16 12:06:07	Crawford	Phillcock	crawford.phillcock@yahoo.com
8f407b04-16cb-6a44-666d-0578e50abe66	2024-09-12 19:06:55	2024-06-12 23:07:13	Verena	Maywood	verena.maywood@hotmail.com
76908e10-9e68-69c2-c623-d26192f75e6b	2023-11-05 06:48:21	2023-12-03 12:09:56	Alica	Jerrolt	alica.jerrolt@gmail.com
ca3cafb0-bb1d-757e-7220-4c1461e5226e	2024-01-10 05:47:10	2024-01-12 02:02:55	Carmencita	Garnsey	carmencita.garnsey@hotmail.com
588d5a01-bcb2-3b94-76d1-d80bf7cc5868	2024-07-25 05:28:20	2024-03-18 04:10:52	Hortensia	Dailly	hortensia.dailly@gmail.com
2f7ff7be-e3f2-c495-2c9e-9228cf19f269	2024-04-22 11:52:07	2023-10-13 11:05:53	Chrysa	Bertomeu	chrysa.bertomeu@yahoo.com
6281f729-ea8d-4218-854f-584aa3718b96	2024-06-21 09:04:15	2024-05-11 07:58:09	Aila	Maccaig	aila.maccaig@aol.com
ce59e370-3406-36d6-0d8d-c146f734d94b	2024-01-16 10:00:31	2024-02-25 21:12:14	Raye	Middlehurst	raye.middlehurst@live.com.au
6b0884fb-f2c4-fcc8-bf51-0c906ffb5b05	2024-06-10 04:56:59	2024-03-29 17:55:22	Katya	Minards	katya.minards@hotmail.com
cd07f4b1-eff6-5f39-f455-45fdb9d6f644	2024-08-08 15:24:47	2024-01-03 21:05:52	Juliane	Dixson	juliane.dixson@aol.com
f8dcce78-9997-fd82-026d-35e9668e6fed	2024-02-11 08:52:32	2024-02-20 01:29:59	Jethro	Cast	jethro.cast@live.nl
07a7b9db-403f-9507-5c1a-74a0759d49af	2023-11-24 17:37:28	2023-10-24 03:14:12	Marshall	Battye	marshall.battye@orange.fr
15d33f02-f026-3606-45d2-da7d906fd91e	2024-09-15 21:06:41	2024-02-13 16:45:52	Nils	Ebi	nils.ebi@hotmail.co.uk
2b77219d-5e14-6ded-16ae-82c217e11a58	2023-10-19 13:44:06	2024-09-19 11:54:41	Hector	MacCague	hector.maccague@hotmail.co.uk
89b543df-5315-b9e6-2cc7-97c739d60921	2024-01-05 12:39:59	2024-09-03 16:27:26	Rufe	Errichiello	rufe.errichiello@yahoo.com
eb7a5cbf-ec24-e712-acc7-68cb4aa0468a	2023-12-20 03:38:29	2024-08-26 00:52:31	Sandy	Dwyer	sandy.dwyer@hotmail.com
af99cbab-f5ed-624d-4deb-f349765ff89d	2023-12-20 02:35:28	2024-03-18 08:09:31	Lowe	McCahill	lowe.mccahill@aol.com
02e00fbe-a566-7a82-b426-4a3a0d361747	2024-06-12 03:07:17	2024-05-28 16:05:38	Hashim	Hourihane	hashim.hourihane@gmail.com
a8603adf-76d4-fa50-dec5-5c796c68ea17	2024-08-04 23:26:55	2023-10-18 16:38:38	Giffie	Hoyer	giffie.hoyer@gmail.com
94daac1b-4f98-fa5a-2d31-c997b2859de4	2024-05-08 05:59:23	2024-01-15 02:39:17	Laurella	Authers	laurella.authers@hotmail.com
05c3dc04-a9c7-caea-51fa-ca9311b5f3a1	2023-11-28 16:02:59	2023-11-02 21:49:25	Chrysa	Barnwill	chrysa.barnwill@web.de
994a5cec-1c0a-ae9e-2d34-0810421ac394	2024-01-05 21:35:03	2024-03-14 13:54:59	Jewel	Wisham	jewel.wisham@yandex.ru
8462bab8-9eeb-5542-82c3-cb1d37c0a014	2024-01-03 02:37:22	2024-10-13 23:53:45.404515	Kerry A.	Varnam	
41289109-f322-47dd-8768-6a4362c2bb6f	2024-10-14 00:06:22.222771	2024-10-14 00:06:22.222771	Bartas	Urba	bartisimo@gmail.com
e2bf5381-ca71-4a44-aee2-09f3967625ea	2024-10-14 10:50:58.790861	2024-10-14 11:05:06.679766	New	User	2
\.



SELECT pg_catalog.setval('public.goose_db_version_id_seq', 25, true);



ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.configuration_items
    ADD CONSTRAINT configuration_items_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT incidents_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);



ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);



ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT fk_assigned_to FOREIGN KEY (assigned_to) REFERENCES public.users(id);



ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT fk_companies FOREIGN KEY (company_id) REFERENCES public.companies(id) ON DELETE CASCADE;



ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT fk_configuration_items FOREIGN KEY (configuration_item_id) REFERENCES public.configuration_items(id);


-- Completed on 2024-10-15 12:04:08 MDT

--
-- PostgreSQL database dump complete
--

