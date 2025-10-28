-- +goose Up
-- Przykładowi użytkownicy
INSERT INTO users (username, email, password_hash, date_of_birth, bio, followers_count, following_count, is_admin) VALUES
('jan_kowalski', 'jan@example.com', '$2a$10$hashedpassword1', '1995-03-15', 'Programista z pasją do technologii', 150, 75, 0),
('anna_nowak', 'anna@example.com', '$2a$10$hashedpassword2', '1992-07-22', 'Designer UX/UI', 200, 120, 0),
('admin_user', 'admin@borg.com', '$2a$10$hashedpassword3', '1990-01-01', 'Administrator systemu', 50, 10, 1),
('maria_wisniewska', 'maria@example.com', '$2a$10$hashedpassword4', '1998-11-08', 'Studentka informatyki', 80, 45, 0),
('piotr_zawadzki', 'piotr@example.com', '$2a$10$hashedpassword5', '1993-05-30', 'Freelancer developer', 90, 60, 0);

-- Przykładowe posty
INSERT INTO posts (user_id, content, like_count, share_count, comment_count) VALUES
(1, 'Dzisiaj skończyłem projekt w React! 🚀 #programming #react', 25, 5, 8),
(2, 'Nowy design system gotowy! Co myślicie? #design #ux', 45, 12, 15),
(1, 'Debugowanie przez 3 godziny... w końcu znalazłem błąd w jednej linii 😅', 18, 3, 6),
(4, 'Pierwszy dzień na nowym stanowisku! Jestem podekscytowana 💪', 35, 8, 12),
(3, 'Aktualizacja systemu zaplanowana na jutro 2:00 AM', 5, 1, 2),
(5, 'Freelancing daje mi wolność, ale czasem brakuje stabilności 🤔', 22, 4, 9),
(2, 'Prototyp nowej aplikacji mobilnej gotowy! #mobile #app', 30, 7, 11),
(1, 'Code review z zespołem - zawsze uczę się czegoś nowego', 15, 2, 5);

-- Przykładowe like'i
INSERT INTO likes (post_id, user_id) VALUES
(1, 2), (1, 3), (1, 4), (1, 5),
(2, 1), (2, 3), (2, 4), (2, 5),
(3, 2), (3, 4), (3, 5),
(4, 1), (4, 2), (4, 3), (4, 5),
(5, 1), (5, 2),
(6, 1), (6, 2), (6, 3), (6, 4),
(7, 1), (7, 3), (7, 4), (7, 5),
(8, 2), (8, 4), (8, 5);

-- Przykładowe share'y
INSERT INTO shares (post_id, user_id) VALUES
(1, 2), (1, 4),
(2, 1), (2, 3), (2, 5),
(3, 2),
(4, 1), (4, 3), (4, 5),
(5, 1),
(6, 2), (6, 4),
(7, 1), (7, 3), (7, 5),
(8, 2);

-- Przykładowe followery
INSERT INTO followers (follower_id, following_id) VALUES
(2, 1), (3, 1), (4, 1), (5, 1),  -- wszyscy śledzą Jana
(1, 2), (3, 2), (4, 2), (5, 2),  -- wszyscy śledzą Annę
(1, 3), (2, 3), (4, 3), (5, 3),  -- wszyscy śledzą admina
(1, 4), (2, 4), (3, 4), (5, 4),  -- wszyscy śledzą Marię
(1, 5), (2, 5), (3, 5), (4, 5);  -- wszyscy śledzą Piotra

-- +goose Down
DELETE FROM followers;
DELETE FROM shares;
DELETE FROM likes;
DELETE FROM posts;
DELETE FROM users;