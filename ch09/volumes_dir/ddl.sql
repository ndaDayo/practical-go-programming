CREATE TABLE users (
  user_id varchar,
  user_name varchar,
  created_at timestamptz,
  PRIMARY key(user_id)
);

INSERT INTO users (user_id, user_name, created_at)
VALUES
('user001', 'Alice', '2021-10-14 09:00:00+00'),
('user002', 'Bob', '2021-10-14 09:10:00+00'),
('user003', 'Charlie', '2021-10-14 09:20:00+00'),
('user004', 'David', '2021-10-14 09:30:00+00')
ON CONFLICT (user_id) DO NOTHING;

