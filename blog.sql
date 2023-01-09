/*
 Navicat Premium Data Transfer
 
 Source Server         : docker-mysql
 Source Server Type    : MySQL
 Source Server Version : 80031
 Source Host           : 192.168.101.10:3306
 Source Schema         : blog
 
 Target Server Type    : MySQL
 Target Server Version : 80031
 File Encoding         : 65001
 
 Date: 09/01/2023 12:05:20
 */
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  `cover_image_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb3 COMMENT = '文章管理';
-- ----------------------------
-- Records of blog_article
-- ----------------------------
BEGIN;
INSERT INTO `blog_article` (
    `id`,
    `tag_id`,
    `title`,
    `desc`,
    `content`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`,
    `cover_image_url`
  )
VALUES (
    1,
    1,
    'test1',
    'test-desc',
    'test-content',
    1671632515,
    'test-created',
    1671632515,
    '',
    0,
    1,
    'http://127.0.0.1:8000/upload/images/671a0da0ba061c98de801409dbc57d7e.png'
  );
INSERT INTO `blog_article` (
    `id`,
    `tag_id`,
    `title`,
    `desc`,
    `content`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`,
    `cover_image_url`
  )
VALUES (
    2,
    1,
    'test1',
    'test-desc',
    'test-content',
    1671632586,
    'test-created',
    1671632586,
    '',
    0,
    1,
    'http://127.0.0.1:8000/upload/images/671a0da0ba061c98de801409dbc57d7e.png'
  );
INSERT INTO `blog_article` (
    `id`,
    `tag_id`,
    `title`,
    `desc`,
    `content`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`,
    `cover_image_url`
  )
VALUES (
    3,
    3,
    '文章3 tag3',
    '文章3 tag3',
    '文章3 tag3',
    1673107493,
    'admin',
    1673107493,
    '',
    0,
    1,
    'https://www.baidu.com'
  );
INSERT INTO `blog_article` (
    `id`,
    `tag_id`,
    `title`,
    `desc`,
    `content`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`,
    `cover_image_url`
  )
VALUES (
    4,
    3,
    '-4',
    '-4',
    '-4',
    1673107627,
    'admin',
    1673109418,
    'admin',
    1673109815,
    1,
    '-4'
  );
COMMIT;
-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '账号',
  `password` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb3;
-- ----------------------------
-- Records of blog_auth
-- ----------------------------
BEGIN;
INSERT INTO `blog_auth` (`id`, `username`, `password`)
VALUES (
    2,
    'admin',
    '$2a$04$bWATxaZowNCQtqfzyoQm/eOw4t33hUg1G5heImWMT7KjnfVcZMqZ2'
  );
COMMIT;
-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int unsigned DEFAULT '0',
  `state` tinyint unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8mb3 COMMENT = '文章标签管理';
-- ----------------------------
-- Records of blog_tag
-- ----------------------------
BEGIN;
INSERT INTO `blog_tag` (
    `id`,
    `name`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`
  )
VALUES (
    1,
    'name-x',
    1671632213,
    'test',
    1671634419,
    'edit1',
    0,
    0
  );
INSERT INTO `blog_tag` (
    `id`,
    `name`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`
  )
VALUES (
    2,
    'TAG-2',
    1671632262,
    'test',
    1671632262,
    '',
    1671634523,
    0
  );
INSERT INTO `blog_tag` (
    `id`,
    `name`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`
  )
VALUES (
    3,
    'TAG-3',
    1672756911,
    'test',
    1672756911,
    '',
    0,
    0
  );
INSERT INTO `blog_tag` (
    `id`,
    `name`,
    `created_on`,
    `created_by`,
    `modified_on`,
    `modified_by`,
    `deleted_on`,
    `state`
  )
VALUES (
    4,
    'Tag等地修改222',
    1672935740,
    'test',
    1672936795,
    'test',
    1672936814,
    1
  );
COMMIT;
SET FOREIGN_KEY_CHECKS = 1;