-- forum -> thread -> post

CREATE EXTENSION IF NOT EXISTS citext;

-- user
DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    uid         BIGSERIAL,
    nickname    CITEXT UNIQUE,
    fullName    TEXT,
    about       TEXT, 
    email       CITEXT UNIQUE,

    PRIMARY KEY(uid)
);

-- forum
DROP TABLE IF EXISTS forums CASCADE;
CREATE TABLE forums (
    fid     	BIGSERIAL,
    owner   	BIGINT,

    title   	TEXT,
    slug    	CITEXT UNIQUE,
	postCount   BIGINT,
	threadCount BIGINT,
    msgCount    BIGINT,

    PRIMARY KEY(fid),
    FOREIGN KEY(owner) REFERENCES users(uid)
);

-- thread
DROP TABLE IF EXISTS threads CASCADE;
CREATE TABLE threads (
    tid         BIGSERIAL,
    author      BIGINT,
    forum       BIGINT,
    
    title       TEXT,
	slug    	CITEXT UNIQUE,
    message     TEXT,
	created     TIMESTAMP WITH TIME ZONE,
	votes		int,
    
    PRIMARY KEY(tid)
);

DROP TABLE IF EXISTS votes CASCADE;
CREATE TABLE votes (
	thread		BIGINT,
	author		BIGINT,
	voice		INTEGER,
	
	FOREIGN KEY(thread) REFERENCES threads(tid),
	FOREIGN KEY(author) REFERENCES users(uid)
);

-- post
DROP TABLE IF EXISTS posts CASCADE;
CREATE TABLE posts (
    pid         BIGSERIAL,
    author      BIGINT,     -- user id
    parent      BIGINT,     -- parent message id (the message is an answer)
    thread      BIGINT,     -- message thread id

    message     TEXT,
    isEdited    BOOLEAN,
    created     TIMESTAMP,

    PRIMARY KEY(pid),
    FOREIGN KEY(author) REFERENCES users(uid),
    FOREIGN KEY(thread) REFERENCES threads(tid)
);
