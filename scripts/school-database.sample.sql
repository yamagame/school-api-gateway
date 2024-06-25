-- データベースの作成
CREATE DATABASE IF NOT EXISTS `school-database`;

-- データベースの削除
DROP DATABASE IF EXISTS `school-database`;

-- テーブルを一覧表示
SHOW tables;

-- テーブルの作成
CREATE TABLE `labos` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主キーの標準フィールド',
  `name` varchar(255) DEFAULT NULL COMMENT '研究室の名前',
  `group` varchar(255) DEFAULT NULL COMMENT '研究室のグループ',
  `program` varchar(255) DEFAULT NULL COMMENT '研究室のプログラム',
  `building` varchar(255) DEFAULT NULL COMMENT '研究室の建物名',
  `created_at` datetime(3) DEFAULT NULL COMMENT 'GORMによって自動的に管理される作成時間',
  `updated_at` datetime(3) DEFAULT NULL COMMENT 'GORMによって自動的に管理される更新時間',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- テーブルの削除
DROP TABLE IF EXISTS labos;

-- カラムの追加
ALTER TABLE labos
ADD `group` varchar(255) COMMENT '研究室のグループ'
AFTER `name`;

-- カラムの削除
ALTER TABLE labos DROP COLUMN `group`;

-- テーブルのカラムを表示
DESC labos;

-- テーブルのカラムを詳細表示
SHOW FULL COLUMNS from labos;

-- レコードの作成
INSERT INTO labos (name)
VALUES ("研究室-001"),
  ("研究室-002"),
  ("研究室-003");

-- レコードの更新
UPDATE labos
SET `group` = "Ⅰ類（情報系）"
WHERE id = 3;

-- レコードの削除
DELETE FROM labos
WHERE id = 2;

-- レコードの表示
SELECT *
FROM `labos`;

-- レコードの表示
SELECT *
FROM `groups`;

-- レコードの表示
SELECT *
FROM `programs`;

-- レコードの表示
SELECT labos.id, labos.name, groups.name, programs.name
FROM `labos`
JOIN `groups` ON groups.id = labos.group_id
JOIN `programs` ON programs.id = labos.program_id
ORDER BY id;

SELECT LAST_INSERT_ID();

-- EOF