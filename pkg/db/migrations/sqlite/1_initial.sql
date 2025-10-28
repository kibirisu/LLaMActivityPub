-- +goose Up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                    
    username TEXT UNIQUE NOT NULL,     
    email TEXT UNIQUE NOT NULL,       
    password_hash TEXT NOT NULL,     
    date_of_birth TEXT,
    bio TEXT,
    followers_count INTEGER DEFAULT 0,        
    following_count INTEGER DEFAULT 0,        
    is_admin INTEGER DEFAULT 0,           
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                    
    user_id INTEGER NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    like_count INTEGER DEFAULT 0,
    share_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE likes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                    
    post_id INTEGER NOT NULL REFERENCES posts(id) ,
    user_id INTEGER NOT NULL REFERENCES users(id) ,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(post_id, user_id)
);

CREATE TABLE shares (
    id INTEGER PRIMARY KEY AUTOINCREMENT,                 
    post_id INTEGER NOT NULL REFERENCES posts(id) ,
    user_id INTEGER NOT NULL REFERENCES users(id) ,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(post_id, user_id)
);

CREATE TABLE followers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    follower_id INTEGER NOT NULL REFERENCES users(id) ,
    following_id INTEGER NOT NULL REFERENCES users(id) ,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(follower_id, following_id),
    CHECK(follower_id != following_id)
);

/*
CREATE TABLE comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER NOT NULL REFERENCES posts(id) ,
    user_id INTEGER NOT NULL REFERENCES users(id) ,
    content TEXT NOT NULL,
    parent_id INTEGER REFERENCES comments(id),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
*/

-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS likes;
DROP TABLE IF EXISTS shares;
DROP TABLE IF EXISTS followers;
--DROP TABLE IF EXISTS comments;
