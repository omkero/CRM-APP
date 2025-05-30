PGDMP      ,                }         
   crm_system    17.2    17.2 Q    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            �           1262    16397 
   crm_system    DATABASE     v   CREATE DATABASE crm_system WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';
    DROP DATABASE crm_system;
                     postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
                     pg_database_owner    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                        pg_database_owner    false    4            �            1259    16398    activity    TABLE     s  CREATE TABLE public.activity (
    activity_id integer NOT NULL,
    activity_employee_id integer NOT NULL,
    activity_action character varying(255) NOT NULL,
    activity_type character varying(255) NOT NULL,
    activity_log text DEFAULT 'No log saved'::text,
    activity_ipv4 character varying(15) NOT NULL,
    created_at timestamp with time zone DEFAULT now()
);
    DROP TABLE public.activity;
       public         heap r       postgres    false    4            �            1259    16404    activity_activity_id_seq    SEQUENCE     �   CREATE SEQUENCE public.activity_activity_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.activity_activity_id_seq;
       public               postgres    false    217    4            �           0    0    activity_activity_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.activity_activity_id_seq OWNED BY public.activity.activity_id;
          public               postgres    false    218            �            1259    16405    customer    TABLE       CREATE TABLE public.customer (
    customer_id bigint NOT NULL,
    customer_username character varying(255) NOT NULL,
    customer_uuid character varying(36) NOT NULL,
    customer_position character varying(255) NOT NULL,
    customer_full_name character varying(255) NOT NULL,
    customer_phone_number character varying(255) NOT NULL,
    customer_email_address character varying(255) NOT NULL,
    customer_created_by_employee_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.customer;
       public         heap r       postgres    false    4            �            1259    16411    customer_address    TABLE     �  CREATE TABLE public.customer_address (
    address_id integer NOT NULL,
    customer_id integer NOT NULL,
    country_name character varying(255) NOT NULL,
    country_code character varying(10) NOT NULL,
    city_name character varying(255) NOT NULL,
    zip_code character varying(255) DEFAULT NULL::character varying,
    street_address character varying(255) DEFAULT NULL::character varying,
    created_at timestamp without time zone DEFAULT now()
);
 $   DROP TABLE public.customer_address;
       public         heap r       postgres    false    4            �            1259    16419    customer_address_address_id_seq    SEQUENCE     �   CREATE SEQUENCE public.customer_address_address_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 6   DROP SEQUENCE public.customer_address_address_id_seq;
       public               postgres    false    220    4            �           0    0    customer_address_address_id_seq    SEQUENCE OWNED BY     c   ALTER SEQUENCE public.customer_address_address_id_seq OWNED BY public.customer_address.address_id;
          public               postgres    false    221            �            1259    16420    customer_customer_id_seq    SEQUENCE     �   CREATE SEQUENCE public.customer_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.customer_customer_id_seq;
       public               postgres    false    4    219            �           0    0    customer_customer_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.customer_customer_id_seq OWNED BY public.customer.customer_id;
          public               postgres    false    222            �            1259    16421    employee    TABLE     �  CREATE TABLE public.employee (
    employee_id integer NOT NULL,
    employee_username character varying(255) NOT NULL,
    employee_uuid character varying(36) NOT NULL,
    employee_position character varying(255) NOT NULL,
    employee_full_name character varying(255) NOT NULL,
    employee_phone_number character varying(255) NOT NULL,
    employee_email_address character varying(255) NOT NULL,
    employee_password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    created_by_employee_id bigint NOT NULL,
    employee_role text[] NOT NULL,
    is_banned boolean DEFAULT false,
    is_suspended boolean DEFAULT false,
    suspension_duration timestamp with time zone
);
    DROP TABLE public.employee;
       public         heap r       postgres    false    4            �            1259    16427    employee_employee_id_seq    SEQUENCE     �   CREATE SEQUENCE public.employee_employee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.employee_employee_id_seq;
       public               postgres    false    223    4            �           0    0    employee_employee_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.employee_employee_id_seq OWNED BY public.employee.employee_id;
          public               postgres    false    224            �            1259    16428    order_detail    TABLE     	  CREATE TABLE public.order_detail (
    order_detail_id integer NOT NULL,
    order_id integer NOT NULL,
    order_quantity integer NOT NULL,
    order_unit_price numeric(10,2),
    order_product_name character varying(255),
    order_product_id integer NOT NULL
);
     DROP TABLE public.order_detail;
       public         heap r       postgres    false    4            �            1259    16431     order_detail_order_detail_id_seq    SEQUENCE     �   CREATE SEQUENCE public.order_detail_order_detail_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 7   DROP SEQUENCE public.order_detail_order_detail_id_seq;
       public               postgres    false    225    4            �           0    0     order_detail_order_detail_id_seq    SEQUENCE OWNED BY     e   ALTER SEQUENCE public.order_detail_order_detail_id_seq OWNED BY public.order_detail.order_detail_id;
          public               postgres    false    226            �            1259    16432    orders    TABLE       CREATE TABLE public.orders (
    order_id integer NOT NULL,
    order_type character varying(255) NOT NULL,
    order_total_amount numeric(10,2) NOT NULL,
    created_by_employee_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.orders;
       public         heap r       postgres    false    4            �            1259    16436    orders_order_id_seq    SEQUENCE     �   CREATE SEQUENCE public.orders_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE public.orders_order_id_seq;
       public               postgres    false    227    4            �           0    0    orders_order_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE public.orders_order_id_seq OWNED BY public.orders.order_id;
          public               postgres    false    228            �            1259    16541    product    TABLE     J  CREATE TABLE public.product (
    product_id integer NOT NULL,
    product_title character varying(255) NOT NULL,
    product_description text NOT NULL,
    product_price numeric(10,2) NOT NULL,
    created_by_employee_id integer NOT NULL,
    product_cover text NOT NULL,
    created_at timestamp with time zone DEFAULT now()
);
    DROP TABLE public.product;
       public         heap r       postgres    false    4            �            1259    16540    product_product_id_seq    SEQUENCE     �   CREATE SEQUENCE public.product_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 -   DROP SEQUENCE public.product_product_id_seq;
       public               postgres    false    236    4            �           0    0    product_product_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE public.product_product_id_seq OWNED BY public.product.product_id;
          public               postgres    false    235            �            1259    16497    system_roles    TABLE       CREATE TABLE public.system_roles (
    role_id integer NOT NULL,
    role_title character varying(255) NOT NULL,
    role_permissions text[] NOT NULL,
    role_created_by_employee_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now()
);
     DROP TABLE public.system_roles;
       public         heap r       postgres    false    4            �            1259    16496    system_roles_role_id_seq    SEQUENCE     �   CREATE SEQUENCE public.system_roles_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE public.system_roles_role_id_seq;
       public               postgres    false    234    4            �           0    0    system_roles_role_id_seq    SEQUENCE OWNED BY     U   ALTER SEQUENCE public.system_roles_role_id_seq OWNED BY public.system_roles.role_id;
          public               postgres    false    233            �            1259    16444    tasks    TABLE     �  CREATE TABLE public.tasks (
    task_id integer NOT NULL,
    task_title character varying(255) NOT NULL,
    task_description text NOT NULL,
    task_start_from timestamp with time zone NOT NULL,
    task_end_in timestamp with time zone NOT NULL,
    task_created_by_employee_id integer NOT NULL,
    task_to_employee_uuid character varying(36) NOT NULL,
    is_task_finished boolean DEFAULT false,
    is_task_canceled boolean DEFAULT false,
    canceled_reason text,
    finished_at timestamp with time zone,
    canceled_at timestamp with time zone,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    priority character varying(255) DEFAULT 'low'::character varying
);
    DROP TABLE public.tasks;
       public         heap r       postgres    false    4            �            1259    16452    tasks_task_id_seq    SEQUENCE     �   CREATE SEQUENCE public.tasks_task_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.tasks_task_id_seq;
       public               postgres    false    4    229            �           0    0    tasks_task_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.tasks_task_id_seq OWNED BY public.tasks.task_id;
          public               postgres    false    230            �            1259    16453    test    TABLE       CREATE TABLE public.test (
    customer_id integer NOT NULL,
    customer_name character varying(255) NOT NULL,
    customer_role character varying(255) NOT NULL,
    prefered_product text NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.test;
       public         heap r       postgres    false    4            �            1259    16459    test_customer_id_seq    SEQUENCE     �   CREATE SEQUENCE public.test_customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 +   DROP SEQUENCE public.test_customer_id_seq;
       public               postgres    false    4    231            �           0    0    test_customer_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE public.test_customer_id_seq OWNED BY public.test.customer_id;
          public               postgres    false    232            �           2604    16460    activity activity_id    DEFAULT     |   ALTER TABLE ONLY public.activity ALTER COLUMN activity_id SET DEFAULT nextval('public.activity_activity_id_seq'::regclass);
 C   ALTER TABLE public.activity ALTER COLUMN activity_id DROP DEFAULT;
       public               postgres    false    218    217            �           2604    16461    customer customer_id    DEFAULT     |   ALTER TABLE ONLY public.customer ALTER COLUMN customer_id SET DEFAULT nextval('public.customer_customer_id_seq'::regclass);
 C   ALTER TABLE public.customer ALTER COLUMN customer_id DROP DEFAULT;
       public               postgres    false    222    219            �           2604    16462    customer_address address_id    DEFAULT     �   ALTER TABLE ONLY public.customer_address ALTER COLUMN address_id SET DEFAULT nextval('public.customer_address_address_id_seq'::regclass);
 J   ALTER TABLE public.customer_address ALTER COLUMN address_id DROP DEFAULT;
       public               postgres    false    221    220                        2604    16463    employee employee_id    DEFAULT     |   ALTER TABLE ONLY public.employee ALTER COLUMN employee_id SET DEFAULT nextval('public.employee_employee_id_seq'::regclass);
 C   ALTER TABLE public.employee ALTER COLUMN employee_id DROP DEFAULT;
       public               postgres    false    224    223                       2604    16464    order_detail order_detail_id    DEFAULT     �   ALTER TABLE ONLY public.order_detail ALTER COLUMN order_detail_id SET DEFAULT nextval('public.order_detail_order_detail_id_seq'::regclass);
 K   ALTER TABLE public.order_detail ALTER COLUMN order_detail_id DROP DEFAULT;
       public               postgres    false    226    225                       2604    16465    orders order_id    DEFAULT     r   ALTER TABLE ONLY public.orders ALTER COLUMN order_id SET DEFAULT nextval('public.orders_order_id_seq'::regclass);
 >   ALTER TABLE public.orders ALTER COLUMN order_id DROP DEFAULT;
       public               postgres    false    228    227                       2604    16544    product product_id    DEFAULT     x   ALTER TABLE ONLY public.product ALTER COLUMN product_id SET DEFAULT nextval('public.product_product_id_seq'::regclass);
 A   ALTER TABLE public.product ALTER COLUMN product_id DROP DEFAULT;
       public               postgres    false    235    236    236                       2604    16500    system_roles role_id    DEFAULT     |   ALTER TABLE ONLY public.system_roles ALTER COLUMN role_id SET DEFAULT nextval('public.system_roles_role_id_seq'::regclass);
 C   ALTER TABLE public.system_roles ALTER COLUMN role_id DROP DEFAULT;
       public               postgres    false    234    233    234                       2604    16467    tasks task_id    DEFAULT     n   ALTER TABLE ONLY public.tasks ALTER COLUMN task_id SET DEFAULT nextval('public.tasks_task_id_seq'::regclass);
 <   ALTER TABLE public.tasks ALTER COLUMN task_id DROP DEFAULT;
       public               postgres    false    230    229                       2604    16468    test customer_id    DEFAULT     t   ALTER TABLE ONLY public.test ALTER COLUMN customer_id SET DEFAULT nextval('public.test_customer_id_seq'::regclass);
 ?   ALTER TABLE public.test ALTER COLUMN customer_id DROP DEFAULT;
       public               postgres    false    232    231            �          0    16398    activity 
   TABLE DATA           �   COPY public.activity (activity_id, activity_employee_id, activity_action, activity_type, activity_log, activity_ipv4, created_at) FROM stdin;
    public               postgres    false    217   2j       �          0    16405    customer 
   TABLE DATA           �   COPY public.customer (customer_id, customer_username, customer_uuid, customer_position, customer_full_name, customer_phone_number, customer_email_address, customer_created_by_employee_id, created_at) FROM stdin;
    public               postgres    false    219   Oj       �          0    16411    customer_address 
   TABLE DATA           �   COPY public.customer_address (address_id, customer_id, country_name, country_code, city_name, zip_code, street_address, created_at) FROM stdin;
    public               postgres    false    220   lj       �          0    16421    employee 
   TABLE DATA             COPY public.employee (employee_id, employee_username, employee_uuid, employee_position, employee_full_name, employee_phone_number, employee_email_address, employee_password, created_at, created_by_employee_id, employee_role, is_banned, is_suspended, suspension_duration) FROM stdin;
    public               postgres    false    223   �j       �          0    16428    order_detail 
   TABLE DATA           �   COPY public.order_detail (order_detail_id, order_id, order_quantity, order_unit_price, order_product_name, order_product_id) FROM stdin;
    public               postgres    false    225   �j       �          0    16432    orders 
   TABLE DATA           n   COPY public.orders (order_id, order_type, order_total_amount, created_by_employee_id, created_at) FROM stdin;
    public               postgres    false    227   �j       �          0    16541    product 
   TABLE DATA           �   COPY public.product (product_id, product_title, product_description, product_price, created_by_employee_id, product_cover, created_at) FROM stdin;
    public               postgres    false    236   �j       �          0    16497    system_roles 
   TABLE DATA           v   COPY public.system_roles (role_id, role_title, role_permissions, role_created_by_employee_id, created_at) FROM stdin;
    public               postgres    false    234   �j       �          0    16444    tasks 
   TABLE DATA           �   COPY public.tasks (task_id, task_title, task_description, task_start_from, task_end_in, task_created_by_employee_id, task_to_employee_uuid, is_task_finished, is_task_canceled, canceled_reason, finished_at, canceled_at, created_at, priority) FROM stdin;
    public               postgres    false    229   k       �          0    16453    test 
   TABLE DATA           g   COPY public.test (customer_id, customer_name, customer_role, prefered_product, created_at) FROM stdin;
    public               postgres    false    231   7k       �           0    0    activity_activity_id_seq    SEQUENCE SET     I   SELECT pg_catalog.setval('public.activity_activity_id_seq', 1670, true);
          public               postgres    false    218            �           0    0    customer_address_address_id_seq    SEQUENCE SET     M   SELECT pg_catalog.setval('public.customer_address_address_id_seq', 1, true);
          public               postgres    false    221            �           0    0    customer_customer_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.customer_customer_id_seq', 24, true);
          public               postgres    false    222            �           0    0    employee_employee_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.employee_employee_id_seq', 43, true);
          public               postgres    false    224            �           0    0     order_detail_order_detail_id_seq    SEQUENCE SET     N   SELECT pg_catalog.setval('public.order_detail_order_detail_id_seq', 1, true);
          public               postgres    false    226            �           0    0    orders_order_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('public.orders_order_id_seq', 1, true);
          public               postgres    false    228            �           0    0    product_product_id_seq    SEQUENCE SET     E   SELECT pg_catalog.setval('public.product_product_id_seq', 60, true);
          public               postgres    false    235            �           0    0    system_roles_role_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('public.system_roles_role_id_seq', 19, true);
          public               postgres    false    233            �           0    0    tasks_task_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.tasks_task_id_seq', 31, true);
          public               postgres    false    230            �           0    0    test_customer_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('public.test_customer_id_seq', 9, true);
          public               postgres    false    232                       2606    16470    activity activity_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY public.activity
    ADD CONSTRAINT activity_pkey PRIMARY KEY (activity_id);
 @   ALTER TABLE ONLY public.activity DROP CONSTRAINT activity_pkey;
       public                 postgres    false    217                       2606    16472 &   customer_address customer_address_pkey 
   CONSTRAINT     l   ALTER TABLE ONLY public.customer_address
    ADD CONSTRAINT customer_address_pkey PRIMARY KEY (address_id);
 P   ALTER TABLE ONLY public.customer_address DROP CONSTRAINT customer_address_pkey;
       public                 postgres    false    220                       2606    16474 '   customer customer_customer_username_key 
   CONSTRAINT     o   ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_customer_username_key UNIQUE (customer_username);
 Q   ALTER TABLE ONLY public.customer DROP CONSTRAINT customer_customer_username_key;
       public                 postgres    false    219                       2606    16476 #   customer customer_customer_uuid_key 
   CONSTRAINT     g   ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_customer_uuid_key UNIQUE (customer_uuid);
 M   ALTER TABLE ONLY public.customer DROP CONSTRAINT customer_customer_uuid_key;
       public                 postgres    false    219                       2606    16478    customer customer_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (customer_id);
 @   ALTER TABLE ONLY public.customer DROP CONSTRAINT customer_pkey;
       public                 postgres    false    219                       2606    16480 ,   employee employee_employee_email_address_key 
   CONSTRAINT     y   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_employee_email_address_key UNIQUE (employee_email_address);
 V   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_employee_email_address_key;
       public                 postgres    false    223                       2606    16482 '   employee employee_employee_username_key 
   CONSTRAINT     o   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_employee_username_key UNIQUE (employee_username);
 Q   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_employee_username_key;
       public                 postgres    false    223            !           2606    16484    employee employee_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (employee_id);
 @   ALTER TABLE ONLY public.employee DROP CONSTRAINT employee_pkey;
       public                 postgres    false    223            #           2606    16486    order_detail order_detail_pkey 
   CONSTRAINT     i   ALTER TABLE ONLY public.order_detail
    ADD CONSTRAINT order_detail_pkey PRIMARY KEY (order_detail_id);
 H   ALTER TABLE ONLY public.order_detail DROP CONSTRAINT order_detail_pkey;
       public                 postgres    false    225            %           2606    16488    orders orders_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public                 postgres    false    227            /           2606    16549    product product_pkey 
   CONSTRAINT     Z   ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (product_id);
 >   ALTER TABLE ONLY public.product DROP CONSTRAINT product_pkey;
       public                 postgres    false    236            +           2606    16505    system_roles system_roles_pkey 
   CONSTRAINT     a   ALTER TABLE ONLY public.system_roles
    ADD CONSTRAINT system_roles_pkey PRIMARY KEY (role_id);
 H   ALTER TABLE ONLY public.system_roles DROP CONSTRAINT system_roles_pkey;
       public                 postgres    false    234            -           2606    16507 (   system_roles system_roles_role_title_key 
   CONSTRAINT     i   ALTER TABLE ONLY public.system_roles
    ADD CONSTRAINT system_roles_role_title_key UNIQUE (role_title);
 R   ALTER TABLE ONLY public.system_roles DROP CONSTRAINT system_roles_role_title_key;
       public                 postgres    false    234            '           2606    16492    tasks tasks_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (task_id);
 :   ALTER TABLE ONLY public.tasks DROP CONSTRAINT tasks_pkey;
       public                 postgres    false    229            )           2606    16494    test test_pkey 
   CONSTRAINT     U   ALTER TABLE ONLY public.test
    ADD CONSTRAINT test_pkey PRIMARY KEY (customer_id);
 8   ALTER TABLE ONLY public.test DROP CONSTRAINT test_pkey;
       public                 postgres    false    231            �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �      �      x������ � �     