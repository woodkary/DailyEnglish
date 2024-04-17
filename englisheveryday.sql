/*
 Navicat Premium Data Transfer

 Source Server         : EnglishEveryday
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : englisheveryday

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 25/03/2024 09:29:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `title` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '单词本书名',
  `learner_num` int(0) NOT NULL COMMENT '正在学习人数',
  `finish_num` int(0) NOT NULL COMMENT '完成学习人数',
  
  `describe` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '描述',
  `id` int(0) NOT NULL COMMENT '书籍id',
  `words_num` int(0) NOT NULL COMMENT '单词数量',
  `grade` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '年级',
  `difficulty` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '难度',
  `date` date NOT NULL COMMENT '日期',
  PRIMARY KEY (`title`, `id`) USING BTREE,
  INDEX `title`(`title`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '词本库' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cet4_dictionary
-- ----------------------------
DROP TABLE IF EXISTS `cet4_dictionary`;
CREATE TABLE `cet4_dictionary`  (
  `words` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '单词',
  `word_id` int(0) NOT NULL COMMENT '单词序号',
  `soundmark` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '音标',
  `describe` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '单词描述，包括若干词性及其词义',
  `question_1` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '该单词的习题1及答案',
  `question_2` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '该单词的习题2及答案',
  PRIMARY KEY (`words`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '某本单词书' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for cet6_dictionary
-- ----------------------------
DROP TABLE IF EXISTS `cet6_dictionary`;
CREATE TABLE `cet6_dictionary`  (
  `words` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '单词',
  `word_id` int(0) NOT NULL COMMENT '单词序号',
  `soundmark` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '音标',
  `describe` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '单词描述，包括若干词性及其词义',
  `question_1` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '该单词的习题1及答案',
  `question_2` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '该单词的习题2及答案',
  PRIMARY KEY (`words`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '某本单词书' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for exam
-- ----------------------------
DROP TABLE IF EXISTS `exam`;
CREATE TABLE `exam`  (
  `examname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '考试名字',
  `exam_id` int(0) NOT NULL COMMENT 'id',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '试题内容',
  `question_num` int(0) DEFAULT NULL COMMENT '试题数量',
  PRIMARY KEY (`examname`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for mistakes
-- ----------------------------
DROP TABLE IF EXISTS `mistakes`;
CREATE TABLE `mistakes`  (
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户',
  `question` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '错题，用_隔开',
  INDEX `username`(`username`) USING BTREE,
  CONSTRAINT `mistakes_ibfk_1` FOREIGN KEY (`username`) REFERENCES `user_info` (`username`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '错题本' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for notebook
-- ----------------------------
DROP TABLE IF EXISTS `notebook`;
CREATE TABLE `notebook`  (
  `words` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '生词,用_连接',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  INDEX `username`(`username`) USING BTREE,
  CONSTRAINT `notebook_ibfk_2` FOREIGN KEY (`username`) REFERENCES `user_info` (`username`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '单词收藏本' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for system_manager
-- ----------------------------
DROP TABLE IF EXISTS `system_manager`;
CREATE TABLE `system_manager`  (
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `id` int(0) NOT NULL COMMENT 'id',
  PRIMARY KEY (`username`) USING BTREE,
  CONSTRAINT `system_manager_ibfk_1` FOREIGN KEY (`username`) REFERENCES `user_info` (`username`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统管理员名单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for team
-- ----------------------------
DROP TABLE IF EXISTS `team`;
CREATE TABLE `team`  (
  `teamname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '团队名',
  `team_id` int(0) NOT NULL COMMENT 'id唯一',
  `teammate` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '团队成员名单 用_连接',
  `team_manager` varchar(250) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '团队管理员名单',
  PRIMARY KEY (`teamname`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '团队成员及管理员名单，普通成员可以直接退出团队，管理员需要在管理员界面先申请取消管理员身份，再在普通成员界面申请退出团队。' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `id` int(0) NOT NULL COMMENT 'id',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '手机号',
  `pwd` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `email` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱',
  `age` int(0) NOT NULL COMMENT '年龄',
  `sex` int(0) NOT NULL COMMENT '性别0女1男',
  `register_date` datetime(0) NOT NULL COMMENT '注册时间',
  PRIMARY KEY (`username`) USING BTREE,
  INDEX `username`(`username`) USING BTREE,
  INDEX `id`(`id`, `username`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户基础信息表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_study
-- ----------------------------
DROP TABLE IF EXISTS `user_study`;
CREATE TABLE `user_study`  (
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户',
  `learn_book` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户正在学习的书，仅限一本',
  `finish_book` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户完成学习的书，若干，用大字符串保存',
  `words_num` int(0) DEFAULT NULL COMMENT '用户计划每天学习的单词数量(例如一本书会被分成6000/40=150天）',
  `words_index` int(0) DEFAULT NULL COMMENT '用户学到第几个单词',
  PRIMARY KEY (`username`) USING BTREE,
  INDEX `learn_book`(`learn_book`) USING BTREE,
  CONSTRAINT `user_study_ibfk_1` FOREIGN KEY (`learn_book`) REFERENCES `books` (`title`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `user_study_ibfk_2` FOREIGN KEY (`username`) REFERENCES `user_info` (`username`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户学习信息表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
