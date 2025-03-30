package core

func SampleMigrationCode() string {
	return `-- +migrate Up
-- Table structure for table users
CREATE TABLE users (
    id  bigint unsigned not null primary key auto_increment,
    fullname varchar(255) unique null,
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=DYNAMIC;

-- +migrate Down
DROP TABLE users;`
}

func SampleProviderCode() string {
	return `package migrator

import (
	"embed"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var fsCoreMigrations embed.FS

func Provide() migrate.EmbedFileSystemMigrationSource {
	return migrate.EmbedFileSystemMigrationSource{
		FileSystem: fsCoreMigrations,
		Root:       "migrations",
	}
}`
}
