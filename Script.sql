create table "users" (
	id serial primary key,
	name varchar(200) not null,
	email varchar(200) not null,
	password varchar(200) not null,
	address varchar(200),
	phone varchar(18),
	birth_date varchar(20)
);

create table "admins"(
	id serial primary key,
	name varchar(200) not null,
	email varchar(200) not null,
	password varchar(200) not null,
	phone varchar(18)
);

create table "genres"(
	id serial primary key,
	genre_name varchar(200) not null
);

create table "books"(
	id serial primary key,
	genre_id int not null,
	title varchar(200),
	author varchar(100),
	publisher varchar(100),
	publish_year int,
	foreign key (genre_id) references genres(id)
);

create table "book_request"(
	id serial primary key,
	user_id int not null,
	approved_admin_id int not null,
	title varchar(200),
	author varchar(100),
	publisher varchar(100),
	publish_year int,
	status_request int,
	foreign key (user_id) references users(id),
	foreign key (approved_admin_id) references admins(id)
);

create table "loans"(
	id serial primary key,
	user_id int not null,
	book_id int not null,
	deadline_date varchar(200),
	date_loan varchar(100),
	date_return varchar(100),
	status_loan int,
	foreign key (user_id) references users(id),
	foreign key (book_id) references books(id)
);

insert into genres (genre_name) values('Non-Fiction'),
('Children''s and Young Adult'),
('Poetry'),
('Drama'),
('Graphic Novels and Comics'),
('Speculative Fiction'),
('Erotica'),
('Classics'),
('Anthologies');

-- butuh 5 saja
delete from genres where genre_name = 'Children''s and Young Adult' or
genre_name = 'Poetry' or
genre_name = 'Graphic Novels and Comics'or
genre_name = 'Speculative Fiction' or
genre_name = 'Erotica' or
genre_name = 'Anthologies';

insert into genres (genre_name) values
('Anthologies');

insert into books (title,author,publisher,publish_year,genre_id) 
values('Educated','Tara Westover','Random House',2018,2)
,('Becoming','Michelle Obama','Crown Publishing',2018,2)
,('A Streetcar Named Desire','Tennessee Williams','New Directions',1947,5)
,('Moby-Dick','Herman Melville','Harper & Brothers',1851,9)
,('The Best American Short Stories 2023','Various Authors','Mariner Books',2023,11);

insert into books (title,author,publisher,publish_year,genre_id) 
values('Persepolis','Marjane Satrapi','Pantheon Books',2000,1);

insert into admins (name, email, password, phone) values(
	'admin', 'admin@library.com','admin123','+62890997722'
);

insert into users (name, address, password, phone, birth_date) values
('anggi', 'Gedong Kuning RT06/02,Banjarnegara','anggi123','+62890997711',	'20-Juni-2002')
,('eko', 'Gedong Kuning RT04/02,Banjarnegara','eko123','+62890997733',		'03-Mei-2002')
,('pam', 'Gedong Kuning RT04/02,Banjarnegara','pam123','+62890997744',		'12-April-2002')
,('budi', 'Gedong Kuning RT02/02,Banjarnegara','budi123','+62890997755',	'19-Juni-2002')
,('santoso','Kuta Banjar RT01/05,Banjarnegara','santoso123','+62890997766',	'15-SEptember-2001');
