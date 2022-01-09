create table users (
                       id uuid default uuid_generate_v4() primary key,
                       username character varying NOT NULL unique,
                       password character varying NOT NULL,
                       role int4 NOT NULL
);

create table folders
(
    id      uuid default uuid_generate_v4() primary key,
    name    character varying COLLATE pg_catalog."default" NOT NULL,
    user_id uuid                                           not null,
    foreign key (user_id) references users (id)
);

CREATE TABLE images (
                        id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
                        user_id uuid not null,
                        folder_id uuid NOT NULL,
                        blob bytea not null,
                        foreign key (user_id) references users(id),
                        FOREIGN KEY(folder_id) REFERENCES folders(id)
);
