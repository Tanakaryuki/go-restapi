CREATE TABLE IF NOT EXISTS `tasks` (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  administrator_user VARCHAR(255) NOT NULL,
  detail VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);