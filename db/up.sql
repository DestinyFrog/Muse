CREATE TABLE `autor`(
    `id` INTEGER PRIMARY KEY,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `name` VARCHAR(255) NOT NULL,
    `profile_path` VARCHAR(255),
    `nationality` VARCHAR(255)
);

CREATE TABLE `song`(
	`id` CHAR(36) PRIMARY KEY,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `name` VARCHAR(255) NOT NULL,
    `year_launch` DATE,
    `file_path` BIGINT,
    `autor_id` BIGINT
);

INSERT INTO `autor` (name, profile_path, nationality)
VALUES 
	('Queen', 'https://www.arkade.com.br/wp-content/uploads/2023/05/b040846ceba13c3e9c125d68389491094e7f2982.jpeg', 'British'),
	('Madonna', 'https://upload.wikimedia.org/wikipedia/pt/2/28/Madonna%2C_%C3%A1lbum.png', 'North-American');

INSERT INTO `song` (id, name, year_launch, file_path, autor_id)
VALUES
	('a7489581-30d4-4235-8a83-121968307518', 'Sheer Heart Attack', '1978-02-10', 'a7489581-30d4-4235-8a83-121968307518.mp3', 1),
	('08329e0c-74f8-4ba6-8d05-e3c92b83edc9', 'Material Girl', '1985-01-30', '08329e0c-74f8-4ba6-8d05-e3c92b83edc9.mp3', 2);