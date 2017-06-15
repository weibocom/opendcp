CREATE database image_build;

use image_build;

CREATE TABLE `t_build_history` (
    `id` int(32) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `project` varchar(50) NOT NULL DEFAULT '' COMMENT 'project',
    `operator` varchar(50) NOT NULL DEFAULT '' COMMENT 'people that start the building task',
    `time` datetime NOT NULL COMMENT 'build time' COMMENT 'time',
    `state` tinyint NOT NULL COMMENT '0: building 1: success 2: fail',
    `logs` text NOT NULL COMMENT 'logs in json',
    PRIMARY KEY (id)
)  COMMENT='table record build history';

CREATE TABLE `t_latest_image` (
    `id` int(32) NOT NULL AUTO_INCREMENT COMMENT 'id',
    `image` varchar(100) NOT NULL DEFAULT '' COMMENT 'image name',
    `tag` varchar(100) NOT NULL DEFAULT '' COMMENT 'image tag',
    PRIMARY KEY (id)
)  COMMENT='table record latest tag of each image';