--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    category character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.categories
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: listings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.listings (
    id integer NOT NULL,
    title character varying(512),
    description text,
    picurl character varying(255),
    price integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);

--
-- Name: listings_categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.listings_categories (
    id integer NOT NULL,
    category_id integer,
    category_id integer
);


--
-- Name: listings_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.listings_categories ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.listings_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: listings_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.listings ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.listings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.categories (id, category, created_at, updated_at) FROM stdin;
1	Comedy	2022-09-23 00:00:00	2022-09-23 00:00:00
2	Sci-Fi	2022-09-23 00:00:00	2022-09-23 00:00:00
3	Horror	2022-09-23 00:00:00	2022-09-23 00:00:00
4	Romance	2022-09-23 00:00:00	2022-09-23 00:00:00
5	Action	2022-09-23 00:00:00	2022-09-23 00:00:00
6	Thriller	2022-09-23 00:00:00	2022-09-23 00:00:00
7	Drama	2022-09-23 00:00:00	2022-09-23 00:00:00
8	Mystery	2022-09-23 00:00:00	2022-09-23 00:00:00
9	Crime	2022-09-23 00:00:00	2022-09-23 00:00:00
10	Animation	2022-09-23 00:00:00	2022-09-23 00:00:00
11	Adventure	2022-09-23 00:00:00	2022-09-23 00:00:00
12	Fantasy	2022-09-23 00:00:00	2022-09-23 00:00:00
13	Superhero	2022-09-23 00:00:00	2022-09-23 00:00:00
\.


--
-- Data for Name: listings; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.listings (id, title, description, picurl, price, created_at, updated_at) FROM stdin;
1	Bicycle	Nice Bike for sale.	/8Z8dptJEypuLoOQro1WugD855YE.jpg  12500	2021-09-23 00:00:00	2022-09-23 00:00:00
2	Jacke	Nice Bike for sale.	/8Z8dptJEypuLoOQro1WugD855YE.jpg  12500	2019-09-23 00:00:00	2022-09-23 00:00:00
3	PS5	Nice Bike for sale.	/8Z8dptJEypuLoOQro1WugD855YE.jpg  12500	2020-09-23 00:00:00	2022-09-23 00:00:00

\.


--
-- Data for Name: listings_categories; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.listings_categories (id, category_id, category_id) FROM stdin;
1	1	5
2	1	12
3	2	5
4	2	11
5	3	9
6	3	7
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, first_name, last_name, email, password, created_at, updated_at) FROM stdin;
1	Admin	User	admin@example.com	$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy	2022-09-23 00:00:00	2022-09-23 00:00:00
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.categories_id_seq', 13, true);


--
-- Name: listings_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.listings_categories_id_seq', 6, true);


--
-- Name: listings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.listings_id_seq', 3, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: listings_categories listings_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.listings_categories
    ADD CONSTRAINT listings_categories_pkey PRIMARY KEY (id);


--
-- Name: listings listings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.listings
    ADD CONSTRAINT listings_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: listings_categories listings_categories_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.listings_categories
    ADD CONSTRAINT listings_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: listings_categories listings_categories_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.listings_categories
    ADD CONSTRAINT listings_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.listings(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

