-- Create "todos" table
CREATE TABLE `todos` (`id` bigint NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` char(36) NOT NULL COMMENT "用户唯一id", `user_name` varchar(255) NOT NULL COMMENT "用户名", `last_active` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "最后活跃时间戳", PRIMARY KEY (`id`), UNIQUE INDEX `idx_user_id` (`user_id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin COMMENT "User is the model entity for the User schema.";
