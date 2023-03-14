CREATE TABLE IF NOT EXISTS Accounts(
    user_id SERIAL PRIMARY KEY,
    uuid uuid  UNIQUE NOT NULL,
    dob DATE NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    first_name VARCHAR(25) NULL DEFAULT NULL,
    last_name VARCHAR(25) NULL DEFAULT NULL,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    rating INTEGER DEFAULT 0,
    active BOOLEAN,
    bio TEXT NULL DEFAULT NULL,
    link VARCHAR(50) NULL DEFAULT NULL,
    verified BOOLEAN,
    created_at DATE NOT NULL,
    updated_at DATE NOT NULL,
    post bigint[]
)

CREATE TABLE Posts(
post_id int(20) NOT NULL AUTO_INCREMENT,
user_id SERIAL NOT NULL
caption  VARCHAR(255),
latitude FLOAT DEFAULT 0.0,
longitude FLOAT  DEFAULT 0.0,
'type' ENUM('image', 'video'),
post_url VARCHAR(255) NOT NULL,
date_created DATE NOT NULL,
date_updated DATE,
PRIMARY KEY (post_id),
FOREIGN KEY (user_id) REFERENCES Accounts(user_id)
)

 CREATE TABLE Followers
(
    record_id serial PRIMARY KEY,
    follower_id uuid NOT NULL,
    following_id uuid NOT NULL,
    CONSTRAINT fk_follower FOREIGN KEY (follower_id)
        REFERENCES accounts (uuid) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_following FOREIGN KEY (following_id)
        REFERENCES accounts (uuid) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

CREATE TABLE Comments(
comment_id INT(20) NOT NULL AUTO_INCREMENT,
post_id INT(20) NOT NULL,
user_id SERIAL NOT NULL,
content TEXT NOT NULL,
date_created DATE NOT NULL,
date_updated DATE,
PRIMARY KEY (comment_id),
FOREIGN KEY (post_id) REFERENCES Posts(post_id)
FOREIGN KEY (user_id) REFERENCES Accounts(user_id)
)

CREATE TABLE Messages(
message_id INT(20) PRIMARY KEY AUTO_INCREMENT,
user_id_from SERIAL NOT NULL,
user_id_to SERIAL NOT NULL,
content text NOT NULL,
date_created date NOT NULL,
FOREIGN KEY (user_id_from) REFERENCES Accounts(user_id),
FOREIGN KEY (user_id_to) REFERENCES Accounts(user_id)
)

CREATE TABLE Likes(
user_id SERIAL NOT NULL AUTO_INCREMENT,
post_id int(20) NOT NULL,
date_created DATE NOT NULL,
PRIMARY KEY (user_id, post_id),
UNIQUE INDEX (post_id, user_id),
FOREIGN KEY (post_id) REFERENCES Posts(post_id),
FOREIGN KEY (user_id) REFERENCES Accounts(user_id)
);