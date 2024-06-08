drop table books;
drop table genres;
drop table book_request;
drop table loans;

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
	genre_name varchar(200) not null,
	created_at timestamp default now()
);

create table "books"(
	id serial primary key,
	genre_id int not null,
	title varchar(200),
	author varchar(100),
	publisher varchar(100),
	publish_year int,
	foreign key (genre_id) references genres(id),
	created_at timestamp default now()
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
	created_at timestamp default now(),
	foreign key (user_id) references users(id),
	foreign key (approved_admin_id) references admins(id)
);

create table "loans"(
	id serial primary key,
	user_id int not null,
	book_id int not null,
	deadline_date date,
	date_loan date,
	date_return date,
	status_loan bool,
	created_at timestamp default now(),
	foreign key (user_id) references users(id),
	foreign key (book_id) references books(id)
);
-- Menambahkan 5 Data Genere
insert into genres (genre_name) 
values('Fiction'),
		('Non-Fiction'),
		('Drama'),
		('Classics'),
		('Anthologies');
--Menambahkan 5 data buku masing-masing 5
insert into books (title,author,publisher,publish_year,genre_id) 
values('Educated','Tara Westover','Random House',2018,2)
,('Persepolis','Marjane Satrapi','Pantheon Books',2000,1)
,('A Streetcar Named Desire','Tennessee Williams','New Directions',1947,3)
,('Moby-Dick','Herman Melville','Harper & Brothers',1851,4)
,('The Best American Short Stories 2023','Various Authors','Mariner Books',2023,5);

--Menambahkan 5 data user
insert into admins (name, email, password, phone) values(
	'admin', 'admin@library.com','admin123','+62890997722'
);

insert into users (name, address, password, phone, birth_date) values
('anggi', 'Gedong Kuning RT06/02,Banjarnegara','anggi123','+62890997711',	'20-Juni-2002')
,('eko', 'Gedong Kuning RT04/02,Banjarnegara','eko123','+62890997733',		'03-Mei-2002')
,('pam', 'Gedong Kuning RT04/02,Banjarnegara','pam123','+62890997744',		'12-April-2002')
,('budi', 'Gedong Kuning RT02/02,Banjarnegara','budi123','+62890997755',	'19-Juni-2002')
,('santoso','Kuta Banjar RT01/05,Banjarnegara','santoso123','+62890997766',	'15-SEptember-2001');

--Menampilkan hasil pencarian data buku berdasarkan judul buku
select title as judul
from books;

--Menampilkan hasil pencarian data buku berdasarkan genre
select b.title , g.genre_name 
from books b
join genres g on g.id = b.genre_id;

--User 1 meminjam buku dengan id 1
insert into loans (user_id,book_id,deadline_date,date_loan,status_loan)
values (1,1,'2024-06-20','2024-06-8',true);
--User 1 meminjam buku dengan id 2
insert into loans (user_id,book_id,deadline_date,date_loan,status_loan)
values (1,2,'2024-06-20','2024-06-8',true);
--User 1 meminjam salah satu buku yang bergenre id 2 
insert into loans (user_id,book_id,deadline_date,date_loan,status_loan)
values (1,1,'2024-06-20','2024-06-8',true);
--User 2 meminjam buku dengan id 3
insert into loans (user_id,book_id,deadline_date,date_loan,status_loan)
values (2,3,'2024-06-20','2024-06-8',true);
--User 3 meminjam buku dengan id 1
insert into loans (user_id,book_id,deadline_date,date_loan,status_loan)
values (3,1,'2024-06-20','2024-06-8',true);


--Mengubah status peminjaman buku yang dipinjam oleh user 1 dan buku dengan id 1 tadi menjadi “dikembalikan”.
update loans 
set date_return = '2024-06-15' , status_loan = false
where user_id = 1 and book_id =1;

--Tampilkan data user beserta buku yang masih dipinjam/belum dikembalikan.
select * from loans where status_loan = true;

--Tampilkan data buku yang statusnya telah dikembalikan oleh user.
select * from loans where status_loan = false;


-- --Tampilkan data buku yang belum pernah dipinjam oleh user.
-- select b.title, l.status_loan 
-- from books b
-- left join loans l on l.book_id  = b.id
-- where l.status_loan is null;

-- --Tampilkan data user beserta banyaknya buku yang pernah dia pinjam.
-- select u.name, l.book_id 
-- from users u
-- join loans l on l.user_id  = u.id;

-- --Tampilkan data user yang belum pernah meminjam buku.
-- select u.name  , l.book_id 
-- from users u
-- left join loans l on l.user_id  = u.id
-- where l.book_id is null;

-- --Tampilkan data user yang sudah pernah meminjam buku.
-- select u.name  , l.book_id 
-- from users u
-- join loans l on l.user_id  = u.id;

-- --Tampilkan data user yang paling banyak meminjam buku
-- select 
--     u.name, 
--     count(l.user_id) as loan_count
-- from users u
-- left join loans l on l.user_id = u.id
-- group by u.name
-- order by loan_count desc ;

-- --Tampilkan data genre beserta banyaknya buku di masing-masing genre
-- select 
-- --    b.title as title,
--     g.genre_name as genre,
--     count(b.title)
-- from books b
-- left join genres g on b.genre_id = g.id
-- group by genre;

-- --Tampilkan genre yang paling banyak dipinjam bukunya oleh user
-- select 
--     g.genre_name as genre,
--     count(b.title)
-- from books b
-- left join genres g on b.genre_id = g.id
-- left join loans l on l.book_id  = b.id 
-- group by genre;

-- --Tampilkan data user, beserta buku yang dipinjam dan sekaligus genre dari buku tersebut dalam satu query statement.

-- select 
--     u.name, 
--     b.title,
--     g.genre_name 
-- from users u
-- left join loans l on l.user_id = u.id
-- left join books b on b.id = l.book_id 
-- left join genres g on g.id = b.genre_id
-- group by u.name, title, genre_name
-- order by b.title asc;

