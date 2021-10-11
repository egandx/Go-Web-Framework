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

 Date: 11/10/2021 18:02:21
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
INSERT INTO `topics` VALUES (1, '兴趣电商助力经济发展', '兴趣电商', '192.168.251.11', 'https://192.168.251.11/1', 12, '2021-08-12 15:32:41');
INSERT INTO `topics` VALUES (2, '阿里蚂蚁捐款7000万驰援山西', '阿里捐款', '192.168.251.12', 'https://192.168.251.11/2', 34, '2021-08-12 15:32:55');
INSERT INTO `topics` VALUES (3, '美团被罚，反垄断监管规则更加清晰', '美团被罚', '192.168.251.13', 'https://192.168.251.11/3', 56, '2021-08-12 15:33:14');
INSERT INTO `topics` VALUES (4, 'SpaceX准备安装星际飞船发射塔巨型机械臂', 'SpaceX开始星际飞船', '192.168.251.14', 'https://192.168.251.11/4', 78, '2021-09-28 14:07:47');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
