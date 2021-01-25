CREATE TABLE IF NOT EXISTS syu_kan.routines(
  id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100),
  started_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL
) DEFAULT CHARSET = utf8mb4;

CREATE TABLE IF NOT EXISTS syu_kan.progress(
  id BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  routine_id BIGINT UNSIGNED NOT NULL,
  date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL,
  INDEX routine_index(routine_id),
  FOREIGN KEY fk_routine_id(routine_id) REFERENCES routines(id)
) DEFAULT CHARSET = utf8mb4;