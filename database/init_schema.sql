INSERT INTO users (name)
SELECT 
    CASE (random() * 10)::int
        WHEN 0 THEN 'Алексей'
        WHEN 1 THEN 'Мария'
        WHEN 2 THEN 'Дмитрий'
        WHEN 3 THEN 'Елена'
        WHEN 4 THEN 'Сергей'
        WHEN 5 THEN 'Анна'
        WHEN 6 THEN 'Владимир'
        WHEN 7 THEN 'Ольга'
        WHEN 8 THEN 'Александр'
        WHEN 9 THEN 'Татьяна'
        WHEN 10 THEN 'Николай'
    END || ' ' || generate_series
FROM generate_series(1, 100);

INSERT INTO labels (name) VALUES
('bug'),
('feature'),
('enhancement'),
('docs'),
('question'),
('duplicate'),
('wontfix'),
('prod'),
('stage'),
('dev'),
('high-priority'),
('low-priority'),
('in-progress'),
('review-needed');