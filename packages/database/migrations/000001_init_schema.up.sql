CREATE TABLE IF NOT EXISTS Accounts(
   user_id INTEGER NOT NULL,
   uuid VARCHAR(50) UNIQUE NOT NULL,
   dob DATE NOT NULL,
   username VARCHAR (25) UNIQUE NOT NULL,
   first_name VARCHAR (25)  NULL DEFAULT NULL,
   last_name VARCHAR (25)  NULL DEFAULT NULL,
   password VARCHAR (50) NOT NULL,
   email VARCHAR(50) UNIQUE NOT NULL,
   rating INTEGER DEFAULT 0,
   active BOOLEAN,
   bio text NULL DEFAULT NULL,
   link VARCHAR(50) null DEFAULT NULL,
   verified BOOLEAN,
   created_at DATE NOT NULL,
   updated_at DATE NOT NULL,
   PRIMARY KEY(user_id),
   FOREIGN KEY(uuid)
)