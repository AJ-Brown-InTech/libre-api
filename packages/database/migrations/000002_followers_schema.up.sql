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
);
