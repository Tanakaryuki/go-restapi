ALTER TABLE `tasks`
DROP PRIMARY KEY,
MODIFY COLUMN `id` VARCHAR(255) NOT NULL,
ADD PRIMARY KEY (`id`);