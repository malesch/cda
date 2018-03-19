--
-- PostgreSQL database dump
--

-- Dumped from database version 10.3
-- Dumped by pg_dump version 10.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: devices; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.devices (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.devices OWNER TO cda;

--
-- Name: events; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.events (
    id uuid NOT NULL,
    "sceneID" uuid NOT NULL,
    "deviceID" uuid NOT NULL,
    start integer NOT NULL,
    "end" integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.events OWNER TO cda;

--
-- Name: media; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.media (
    id uuid NOT NULL,
    name text,
    type character varying(255) NOT NULL,
    size integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.media OWNER TO cda;

--
-- Name: props; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.props (
    id uuid NOT NULL,
    "eventID" uuid NOT NULL,
    name character varying(255) NOT NULL,
    value character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.props OWNER TO cda;

--
-- Name: scenes; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.scenes (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    "mediumID" uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.scenes OWNER TO cda;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO cda;

--
-- Name: users; Type: TABLE; Schema: public; Owner: cda
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    password_hash character varying(255) NOT NULL,
    locale_id integer NOT NULL,
    role_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO cda;

--
-- Name: devices devices_pkey; Type: CONSTRAINT; Schema: public; Owner: cda
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pkey PRIMARY KEY (id);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: cda
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- Name: media media_pkey; Type: CONSTRAINT; Schema: public; Owner: cda
--

ALTER TABLE ONLY public.media
    ADD CONSTRAINT media_pkey PRIMARY KEY (id);


--
-- Name: props props_pkey; Type: CONSTRAINT; Schema: public; Owner: cda
--

ALTER TABLE ONLY public.props
    ADD CONSTRAINT props_pkey PRIMARY KEY (id);


--
-- Name: scenes scenes_pkey; Type: CONSTRAINT; Schema: public; Owner: cda
--

ALTER TABLE ONLY public.scenes
    ADD CONSTRAINT scenes_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: cda
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: cda
--

CREATE UNIQUE INDEX version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

