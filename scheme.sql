CREATE EXTENSION IF NOT EXISTS citext;

-----------------------------| user |-----------------------------

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE users (
    uid         BIGSERIAL       NOT NULL,   -- A/
    nickname    CITEXT UNIQUE   NOT NULL,   -- 
    fullName    TEXT            NOT NULL,   -- 
    about       TEXT            NOT NULL,   -- 
    email       CITEXT UNIQUE   NOT NULL,   -- 
    PRIMARY KEY(uid)
);

CREATE INDEX users_uid  ON users USING hash (uid);
CREATE INDEX users_nick ON users USING hash (nickname);

-----------------------------| forum |-----------------------------

DROP TABLE IF EXISTS forums CASCADE;
CREATE TABLE forums (
    fid     	BIGSERIAL,                  -- A/
    owner   	BIGINT          NOT NULL,   -- 
    title   	TEXT            NOT NULL,   -- 
    slug    	CITEXT UNIQUE   NOT NULL,   -- 
	threadCount BIGINT          NOT NULL,   -- A/
    postCount   BIGINT          NOT NULL,   -- A/
    PRIMARY KEY(fid),
    FOREIGN KEY(owner) REFERENCES users(uid)
);

CREATE INDEX forums_slug      ON forums USING hash (fid );
CREATE UNIQUE INDEX forums_id ON forums (slug);


DROP TABLE IF EXISTS forum_users CASCADE;
CREATE TABLE forum_users (      -- A/
    username BIGINT NOT NULL,   -- A/
    forum    BIGINT NOT NULL,   -- A/
    FOREIGN KEY(username) REFERENCES users (uid),
    FOREIGN KEY(forum   ) REFERENCES forums(fid),
    UNIQUE(username, forum)
);

CREATE INDEX forum_users_username ON forum_users USING hash (username);
CREATE INDEX forum_users_forum    ON forum_users (forum   );

-----------------------------| thread |-----------------------------

DROP TABLE IF EXISTS threads CASCADE;
CREATE TABLE threads (
    tid         BIGSERIAL                   NOT NULL,   -- A/
    author      BIGINT                      NOT NULL,   --
    forum       BIGINT                      NOT NULL,   --
    title       TEXT                        NOT NULL,   --
	slug    	CITEXT UNIQUE                   NULL,   --
    message     TEXT                        NOT NULL,   --
	created     TIMESTAMP WITH TIME ZONE    NOT NULL,   --
	votes		int                         NOT NULL,   -- A/
    PRIMARY KEY(tid)
);

CREATE INDEX threads_id           ON threads USING hash (tid);
CREATE INDEX threads_slug         ON threads USING hash (slug);
CREATE INDEX thread_created       ON threads (created);
CREATE INDEX thread_forum_created ON threads (forum, created);


DROP TABLE IF EXISTS votes CASCADE;
CREATE TABLE votes (
	thread		BIGINT  NOT NULL, -- 
	author		BIGINT  NOT NULL, -- 
	voice		INTEGER NOT NULL, -- 
	
	FOREIGN KEY(thread) REFERENCES threads(tid),
	FOREIGN KEY(author) REFERENCES users(uid)
);

CREATE UNIQUE INDEX votes_main ON votes (author, thread);

----
CREATE OR REPLACE FUNCTION on_thread() RETURNS TRIGGER AS $BODY$ BEGIN 
    -- add the thread's user into a forum user list
    INSERT INTO forum_users (forum, username) 
        VALUES (
            new.forum, 
            new.author
        ) 
        ON conflict do nothing;
    -- increase a thread count
    UPDATE forums
		SET threadCount=threadCount+1
		WHERE fid=new.forum;
    RETURN new;
END; 
$BODY$ LANGUAGE plpgsql;

CREATE TRIGGER on_thread BEFORE INSERT ON threads FOR EACH ROW EXECUTE PROCEDURE on_thread();

----
CREATE OR REPLACE FUNCTION on_vote_edit() RETURNS TRIGGER AS $BODY$ BEGIN 
    -- update votes
    IF (TG_OP = 'INSERT') THEN
        UPDATE threads
        SET votes = votes + new.voice
        WHERE tid = new.thread;
    ELSE 
        UPDATE threads
        SET votes = votes - old.voice + new.voice
        WHERE tid = new.thread;
    END IF;
    RETURN new;
END; 
$BODY$ LANGUAGE plpgsql;

CREATE TRIGGER on_vote_edit AFTER UPDATE OR INSERT ON votes FOR EACH ROW EXECUTE PROCEDURE on_vote_edit();

-----------------------------| post |-----------------------------

DROP TABLE IF EXISTS posts CASCADE;
CREATE TABLE posts (
    pid         BIGSERIAL,  -- 
    author      CITEXT,     -- 
    parent      BIGINT,     -- 
    thread      BIGINT,     -- 
    forum       CITEXT,     -- 
    message     TEXT,       -- 
    isEdited    BOOLEAN,    -- A/
    created     TIMESTAMP,  -- 
    root        BIGINT,     -- A/
	path		BIGINT[] NOT NULL,

    PRIMARY KEY(pid),
    FOREIGN KEY(author) REFERENCES users(nickname),
    FOREIGN KEY(forum)  REFERENCES forums(slug),
    FOREIGN KEY(thread) REFERENCES threads(tid)
);

CREATE INDEX posts_main   ON posts USING hash (pid);
CREATE INDEX posts_root   ON posts (root);
CREATE INDEX posts_path   ON posts (path);
CREATE INDEX posts_thread ON posts (thread);
CREATE INDEX posts_sort   ON posts (path, created);

----
CREATE OR REPLACE FUNCTION on_post() RETURNS TRIGGER AS $BODY$ BEGIN
    -- fix path
    new.path = new.path || new.pid;
    new.root = new.path[1];
    RETURN new;
END; 
$BODY$ LANGUAGE plpgsql;

CREATE TRIGGER on_post BEFORE INSERT ON posts FOR EACH ROW EXECUTE PROCEDURE on_post();

----
CREATE OR REPLACE FUNCTION on_post_edit() RETURNS TRIGGER AS $BODY$ BEGIN 
    -- set edited flag
    IF new.message != old.message THEN
        new.isEdited = TRUE;
    END IF;
    RETURN new;
END; 
$BODY$ LANGUAGE plpgsql;

CREATE TRIGGER on_post_edit BEFORE UPDATE ON posts FOR EACH ROW EXECUTE PROCEDURE on_post_edit();