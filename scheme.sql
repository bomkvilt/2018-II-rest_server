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

CREATE UNIQUE INDEX users_main ON users (nickname);

-- forum
DROP TABLE IF EXISTS forums CASCADE;
CREATE TABLE forums (
    fid     	BIGSERIAL,
    owner   	BIGINT,

    title   	TEXT,
    slug    	CITEXT UNIQUE NULL,
	postCount   BIGINT,
	threadCount BIGINT,
    msgCount    BIGINT,

    PRIMARY KEY(fid),
    FOREIGN KEY(owner) REFERENCES users(uid)
);

-- CREATE UNIQUE INDEX forums_slug ON forums (fid );
-- CREATE UNIQUE INDEX forums_id   ON forums (slug);

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

CREATE UNIQUE INDEX threads_slug ON threads (tid );
CREATE UNIQUE INDEX threads_id   ON threads (slug);

DROP TABLE IF EXISTS votes CASCADE;
CREATE TABLE votes (
	thread		BIGINT,
	author		BIGINT,
	voice		INTEGER,
	
	FOREIGN KEY(thread) REFERENCES threads(tid),
	FOREIGN KEY(author) REFERENCES users(uid)
);

CREATE UNIQUE INDEX votes_main ON votes (author, thread);

-- post
DROP TABLE IF EXISTS posts CASCADE;
CREATE TABLE posts (
    pid         BIGSERIAL,
    author      CITEXT,     -- user id
    parent      BIGINT,     -- parent message id (the message is an answer)
    thread      BIGINT,     -- message thread id
    forum       CITEXT,

    message     TEXT,
    isEdited    BOOLEAN,
    created     TIMESTAMP,
	path		BIGINT[] NOT NULL,

    PRIMARY KEY(pid)
    -- FOREIGN KEY(author) REFERENCES users(uid),
    -- FOREIGN KEY(thread) REFERENCES threads(tid)
);

CREATE INDEX posts_main ON posts USING hash (pid);

-------------------| triggers |-------------------

CREATE OR REPLACE FUNCTION fix_path() RETURNS TRIGGER AS $BODY$ BEGIN 
    new.path = new.path || new.pid;
    RETURN new;
END; $BODY$ LANGUAGE plpgsql;

CREATE TRIGGER fix_path BEFORE INSERT ON posts FOR EACH ROW EXECUTE PROCEDURE fix_path();