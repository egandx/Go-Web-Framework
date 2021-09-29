/*
 Navicat Premium Data Transfer

 Source Server         : egan
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : gin

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 29/09/2021 08:29:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for topic_classes
-- ----------------------------
DROP TABLE IF EXISTS `topic_classes`;
CREATE TABLE `topic_classes` (
  `class_id` int NOT NULL AUTO_INCREMENT,
  `class_name` varchar(255) NOT NULL,
  `class_remark` text,
  `classtype` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`class_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of topic_classes
-- ----------------------------
BEGIN;
INSERT INTO `topic_classes` VALUES (1, '新闻类', '新闻说明', '文');
INSERT INTO `topic_classes` VALUES (2, '技术类', '技术说明', '理');
INSERT INTO `topic_classes` VALUES (3, '文学类', '文学理解', '文');
INSERT INTO `topic_classes` VALUES (4, '地理类', '地理解释', '文');
COMMIT;

-- ----------------------------
-- Table structure for topics
-- ----------------------------
DROP TABLE IF EXISTS `topics`;
CREATE TABLE `topics` (
  `topic_id` int NOT NULL AUTO_INCREMENT,
  `topic_title` varchar(200) DEFAULT NULL,
  `topic_short_title` varchar(50) DEFAULT NULL,
  `user_ip` varchar(20) DEFAULT NULL,
  `topic_url` varchar(255) DEFAULT NULL,
  `topic_score` int DEFAULT NULL,
  `topic_date` datetime DEFAULT NULL,
  PRIMARY KEY (`topic_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of topics
-- ----------------------------
BEGIN;
INSERT INTO `topics` VALUES (1, '第一个标题', 'one', NULL, NULL, NULL, '2021-08-12 15:32:41');
INSERT INTO `topics` VALUES (2, '第二个标题', NULL, NULL, NULL, NULL, '2021-08-12 15:32:55');
INSERT INTO `topics` VALUES (3, '第二个标题', NULL, NULL, NULL, NULL, '2021-08-12 15:33:14');
INSERT INTO `topics` VALUES (4, 'TopicTitle', 'TopicShortTitle', '127.0.0.1', 'testurl', 0, '2021-09-28 14:07:47');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
