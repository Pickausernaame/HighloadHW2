
-- Удаляем таблицу generations, если она есть
DROP TABLE IF EXISTS generations;

CREATE TABLE IF NOT EXISTS generations (
		id              BIGSERIAL         	PRIMARY KEY,
		generation		int          UNIQUE );

