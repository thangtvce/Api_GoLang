PGDMP  6                    |            demo    16.2    16.2 $               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16630    demo    DATABASE     ~   CREATE DATABASE demo WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.949';
    DROP DATABASE demo;
                postgres    false            �            1259    16671    notifications    TABLE     ^   CREATE TABLE public.notifications (
    notification_id integer NOT NULL,
    message text
);
 !   DROP TABLE public.notifications;
       public         heap    postgres    false            �            1259    16670 !   notifications_notification_id_seq    SEQUENCE     �   CREATE SEQUENCE public.notifications_notification_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 8   DROP SEQUENCE public.notifications_notification_id_seq;
       public          postgres    false    218                       0    0 !   notifications_notification_id_seq    SEQUENCE OWNED BY     g   ALTER SEQUENCE public.notifications_notification_id_seq OWNED BY public.notifications.notification_id;
          public          postgres    false    217            �            1259    16681    roles    TABLE     k   CREATE TABLE public.roles (
    role_id integer NOT NULL,
    role_name character varying(255) NOT NULL
);
    DROP TABLE public.roles;
       public         heap    postgres    false            �            1259    16680    roles_role_id_seq    SEQUENCE     �   CREATE SEQUENCE public.roles_role_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.roles_role_id_seq;
       public          postgres    false    220                       0    0    roles_role_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.roles_role_id_seq OWNED BY public.roles.role_id;
          public          postgres    false    219            �            1259    16714    usernotifications    TABLE     n   CREATE TABLE public.usernotifications (
    user_id integer NOT NULL,
    notification_id integer NOT NULL
);
 %   DROP TABLE public.usernotifications;
       public         heap    postgres    false            �            1259    16694 	   userroles    TABLE     ^   CREATE TABLE public.userroles (
    user_id integer NOT NULL,
    role_id integer NOT NULL
);
    DROP TABLE public.userroles;
       public         heap    postgres    false            �            1259    16632    users    TABLE     �   CREATE TABLE public.users (
    userid integer NOT NULL,
    username character varying(255) NOT NULL,
    role character varying(50) NOT NULL,
    password character varying(255) NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    16631    users_userid_seq    SEQUENCE     �   CREATE SEQUENCE public.users_userid_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.users_userid_seq;
       public          postgres    false    216                       0    0    users_userid_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.users_userid_seq OWNED BY public.users.userid;
          public          postgres    false    215            c           2604    16674    notifications notification_id    DEFAULT     �   ALTER TABLE ONLY public.notifications ALTER COLUMN notification_id SET DEFAULT nextval('public.notifications_notification_id_seq'::regclass);
 L   ALTER TABLE public.notifications ALTER COLUMN notification_id DROP DEFAULT;
       public          postgres    false    218    217    218            d           2604    16684    roles role_id    DEFAULT     n   ALTER TABLE ONLY public.roles ALTER COLUMN role_id SET DEFAULT nextval('public.roles_role_id_seq'::regclass);
 <   ALTER TABLE public.roles ALTER COLUMN role_id DROP DEFAULT;
       public          postgres    false    219    220    220            b           2604    16635    users userid    DEFAULT     l   ALTER TABLE ONLY public.users ALTER COLUMN userid SET DEFAULT nextval('public.users_userid_seq'::regclass);
 ;   ALTER TABLE public.users ALTER COLUMN userid DROP DEFAULT;
       public          postgres    false    215    216    216                      0    16671    notifications 
   TABLE DATA           A   COPY public.notifications (notification_id, message) FROM stdin;
    public          postgres    false    218   �(       	          0    16681    roles 
   TABLE DATA           3   COPY public.roles (role_id, role_name) FROM stdin;
    public          postgres    false    220   �(                 0    16714    usernotifications 
   TABLE DATA           E   COPY public.usernotifications (user_id, notification_id) FROM stdin;
    public          postgres    false    222   �(       
          0    16694 	   userroles 
   TABLE DATA           5   COPY public.userroles (user_id, role_id) FROM stdin;
    public          postgres    false    221   )                 0    16632    users 
   TABLE DATA           A   COPY public.users (userid, username, role, password) FROM stdin;
    public          postgres    false    216   ")                  0    0 !   notifications_notification_id_seq    SEQUENCE SET     P   SELECT pg_catalog.setval('public.notifications_notification_id_seq', 1, false);
          public          postgres    false    217                       0    0    roles_role_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.roles_role_id_seq', 1, false);
          public          postgres    false    219                       0    0    users_userid_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.users_userid_seq', 17, true);
          public          postgres    false    215            h           2606    16678     notifications notifications_pkey 
   CONSTRAINT     k   ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (notification_id);
 J   ALTER TABLE ONLY public.notifications DROP CONSTRAINT notifications_pkey;
       public            postgres    false    218            j           2606    16686    roles roles_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (role_id);
 :   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_pkey;
       public            postgres    false    220            l           2606    16688    roles roles_role_name_key 
   CONSTRAINT     Y   ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_role_name_key UNIQUE (role_name);
 C   ALTER TABLE ONLY public.roles DROP CONSTRAINT roles_role_name_key;
       public            postgres    false    220            p           2606    16718 (   usernotifications usernotifications_pkey 
   CONSTRAINT     |   ALTER TABLE ONLY public.usernotifications
    ADD CONSTRAINT usernotifications_pkey PRIMARY KEY (user_id, notification_id);
 R   ALTER TABLE ONLY public.usernotifications DROP CONSTRAINT usernotifications_pkey;
       public            postgres    false    222    222            n           2606    16698    userroles userroles_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.userroles
    ADD CONSTRAINT userroles_pkey PRIMARY KEY (user_id, role_id);
 B   ALTER TABLE ONLY public.userroles DROP CONSTRAINT userroles_pkey;
       public            postgres    false    221    221            f           2606    16637    users users_pkey 
   CONSTRAINT     R   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (userid);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    216            s           2606    16724 8   usernotifications usernotifications_notification_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.usernotifications
    ADD CONSTRAINT usernotifications_notification_id_fkey FOREIGN KEY (notification_id) REFERENCES public.notifications(notification_id);
 b   ALTER TABLE ONLY public.usernotifications DROP CONSTRAINT usernotifications_notification_id_fkey;
       public          postgres    false    222    218    4712            t           2606    16719 0   usernotifications usernotifications_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.usernotifications
    ADD CONSTRAINT usernotifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(userid);
 Z   ALTER TABLE ONLY public.usernotifications DROP CONSTRAINT usernotifications_user_id_fkey;
       public          postgres    false    216    4710    222            q           2606    16704     userroles userroles_role_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.userroles
    ADD CONSTRAINT userroles_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.roles(role_id);
 J   ALTER TABLE ONLY public.userroles DROP CONSTRAINT userroles_role_id_fkey;
       public          postgres    false    220    221    4714            r           2606    16699     userroles userroles_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.userroles
    ADD CONSTRAINT userroles_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(userid);
 J   ALTER TABLE ONLY public.userroles DROP CONSTRAINT userroles_user_id_fkey;
       public          postgres    false    221    216    4710                  x�3��H�������� �      	      x�3�LL�������� $�            x�3�4����� ]      
      x������ � �         0   x�3�LL� .Cs�Ԋ�܂������"N0)H,..�/J����� ԡb     