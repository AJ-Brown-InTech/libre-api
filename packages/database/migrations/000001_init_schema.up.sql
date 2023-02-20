CREATE TABLE IF NOT EXISTS Account(
   user_id int UNIQUE INDEX PRIMARY KEY,
   uuid bigserial UNIQUE NOT NULL,
   dob DATE NOT NULL,
   username VARCHAR (25) UNIQUE NOT NULL,
   first_name VARCHAR (25)  NULL DEFAULT NULL,
   last_name VARCHAR (25)  NULL DEFAULT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR(50) UNIQUE NOT NULL,
   session DATETIME,
   rating INT NOT NULL,
   post []INT,
   active BOOLEAN,
   bio text NULL DEFAULT NULL,
   link VARCHAR(50) null DEFAULT NULL,
   verfied BOOLEAN,
   creadted_at DATETIME NOT NULL,
   updated_at DATETIME NOT NULL,
)

CREATE TABLE IF NOT EXISTS Followers (
user_id int NOT NULL
follower_id int NOT NULL
FOREIGN KEY (user_id) REFERENCES Account(user_id)
FOREIGN KEY (follower_id) REFERENCES Account(user_id)
)


