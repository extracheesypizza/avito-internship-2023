CREATE TABLE segments (
	seg_name varchar(255)  NOT NULL,
	seg_id serial PRIMARY KEY
);
	
CREATE TABLE user_segments (
	user_id int  NOT NULL,
	seg_id int  REFERENCES segments (seg_id) ON DELETE CASCADE
);
	
CREATE TABLE operations (
	operation_id serial PRIMARY KEY,
	user_id int  NOT NULL,
	seg_id int  NOT NULL,
	operation varchar(3)  NOT NULL,
	at_timestamp timestamp  NOT NULL,
	TTL int  NOT NULL
);