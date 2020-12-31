/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80012
Source Host           : localhost:3306
Source Database       : gin_mvc

Target Server Type    : MYSQL
Target Server Version : 80012
File Encoding         : 65001

Date: 2021-01-01 05:45:25
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `go_user`
-- ----------------------------
DROP TABLE IF EXISTS `go_user`;
CREATE TABLE `go_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `create_time` int(10) DEFAULT '0',
  `update_time` int(10) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of go_user
-- ----------------------------
INSERT INTO `go_user` VALUES ('1', 'harden', '$2a$10$RTwJAzNyDgbpdVd.ljVGWeP5Wd8uw/Mojxcb9hL/0sCn4l1n2hYvu', '1609441361', '0');
INSERT INTO `go_user` VALUES ('2', 'harden0', '$2a$10$Tj/KIfTfGCDMtACwatcGse6/WvaHUBWtiIuqM1eAo8QR.phjxnxeO', '1609449492', '0');
INSERT INTO `go_user` VALUES ('3', 'harden1', '$2a$10$TV2E7C9U7vAtHzrQZGXwE.7yInJWU5Y3AON4GV1aEVAgaeTECtGfa', '1609449493', '0');
INSERT INTO `go_user` VALUES ('4', 'harden2', '$2a$10$GM/v0Il67bUqo3Qe3alpPuRckyYfAqolQ3aNMAnb99VKZqIywQLJm', '1609449493', '0');
INSERT INTO `go_user` VALUES ('5', 'harden3', '$2a$10$3RY6qOfpNQ.BtyupWESSoOPGq/aL.njTyicWQgfQNI07YnrY0C6Wa', '1609449493', '0');
INSERT INTO `go_user` VALUES ('6', 'harden4', '$2a$10$vDTDKAsixib5M5mA9exx5OqQFXCzX3fHstiNYym9ldm3DiRdp/ZKO', '1609449493', '0');
INSERT INTO `go_user` VALUES ('7', 'harden5', '$2a$10$ztyidhLdQmssxWh.paWRHeWyA/ROv/LiVR/hShfinmmYNHZ86rhkW', '1609449493', '0');
INSERT INTO `go_user` VALUES ('8', 'harden6', '$2a$10$WfkPujSM92Pi2fPERzjGkOwO6JlsoLST7ShyXWfy8/pRMgqwy934a', '1609449493', '0');
INSERT INTO `go_user` VALUES ('9', 'harden7', '$2a$10$bd5G2iJbj29KFJLFyACVsuHFb2jfvQgNkIXg3MtaJIJoL9dMEKzoC', '1609449493', '0');
INSERT INTO `go_user` VALUES ('10', 'harden8', '$2a$10$tgqohd8WSVa5xoGeIiLhFOy..L1ZrQycDF4uXNyavFM10tCFIrfEy', '1609449493', '0');
INSERT INTO `go_user` VALUES ('11', 'harden9', '$2a$10$4q7bb8Dnv4NAY2ilJteD8Onk/QxUITVzwU1FTmValZgHnRudEghva', '1609449493', '0');
