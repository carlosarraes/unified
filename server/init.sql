CREATE TABLE search_histories (
	id bigint unsigned NOT NULL AUTO_INCREMENT,
	created_at datetime(3),
	updated_at datetime(3),
	deleted_at datetime(3),
	web longtext,
	category longtext,
	search_results longblob,
	PRIMARY KEY (id),
	KEY idx_search_histories_deleted_at (deleted_at)
);
