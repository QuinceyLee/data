/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : localhost:3306
 Source Schema         : base

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 27/10/2021 16:03:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_honeyed
-- ----------------------------
DROP TABLE IF EXISTS `tb_honeyed`;
CREATE TABLE `tb_honeyed`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `ip_addr` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `port` int(0) NULL DEFAULT -1,
  `attack_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `attack_type` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `contract` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_success` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NO',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4232 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_honeyed
-- ----------------------------
INSERT INTO `tb_honeyed` VALUES (32, '192.168.221.12', '中国·杭州', 22, '2021-10-27 13:38:43', '密码登录', 'ssh', '通过ssh进行密码尝试：root&admin,root&admin123,root@root,root&test,root&test123', 'NO');
INSERT INTO `tb_honeyed` VALUES (4232, '116.50.191.91', '菲律宾', 445, '2021-10-27 15:56:45', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4233, '104.248.6.52', '美国', -1, '2021-10-27 15:56:45', '端口扫描', 'TCP', '扫描端口：	5064  6346  11312  11383  14345  16765  18008  21435  51562', 'NO');
INSERT INTO `tb_honeyed` VALUES (4234, '206.221.176.229', '菲律宾', 1434, '2021-10-27 15:56:45', 'UDP连接', 'UDP', '端口扫描： 1434', 'NO');
INSERT INTO `tb_honeyed` VALUES (4235, '113.234.188.206', '中国·辽宁·大连', 22, '2021-10-27 15:56:45', 'ssh连接', 'TCP', 'ssh密码登录：root&admin,root&admin123,root@root,root&test,root&test123,admin&admin,admin&admin123,admin&root,admin&test,admin&test123', 'NO');
INSERT INTO `tb_honeyed` VALUES (4236, '116.50.191.91', '中国·香港', 445, '2021-10-27 15:56:45', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4237, '172.93.189.171', '墨西哥', 445, '2021-10-27 15:56:45', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4238, '116.50.131.91', '菲律宾', 6379, '2021-10-27 15:56:45', 'Redis连接', 'TCP', '116.50.191.131:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4239, '116.50.133.91', '泰国', 445, '2021-10-27 15:56:45', 'TCP连接', 'TCP', '116.50.133.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4240, '116.33.191.91', '菲律宾', 8, '2021-10-27 15:56:45', 'TCP连接', 'TCP', '116.33.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4241, '116.50.191.91', '菲律宾', 445, '2021-10-27 15:57:07', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4242, '104.248.6.52', '美国', -1, '2021-10-27 15:57:07', '端口扫描', 'TCP', '扫描端口：	5064  6346  11312  11383  14345  16765  18008  21435  51562', 'NO');
INSERT INTO `tb_honeyed` VALUES (4243, '206.221.176.229', '菲律宾', 1434, '2021-10-27 15:57:07', 'UDP连接', 'UDP', '端口扫描： 1434', 'NO');
INSERT INTO `tb_honeyed` VALUES (4244, '113.234.188.206', '中国·辽宁·大连', 22, '2021-10-27 15:57:07', 'ssh连接', 'TCP', 'ssh密码登录：root&admin,root&admin123,root@root,root&test,root&test123,admin&admin,admin&admin123,admin&root,admin&test,admin&test123', 'NO');
INSERT INTO `tb_honeyed` VALUES (4245, '116.50.191.91', '中国·香港', 445, '2021-10-27 15:57:07', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4246, '172.93.189.171', '墨西哥', 445, '2021-10-27 15:57:07', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4247, '116.50.131.91', '菲律宾', 6379, '2021-10-27 15:57:07', 'Redis连接', 'TCP', '116.50.191.131:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4248, '116.50.133.91', '泰国', 445, '2021-10-27 15:57:07', 'TCP连接', 'TCP', '116.50.133.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4249, '116.33.191.91', '菲律宾', 8, '2021-10-27 15:57:07', 'TCP连接', 'TCP', '116.33.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4250, '116.50.191.91', '菲律宾', 445, '2021-10-27 15:57:08', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4251, '104.248.6.52', '美国', -1, '2021-10-27 15:57:08', '端口扫描', 'TCP', '扫描端口：	5064  6346  11312  11383  14345  16765  18008  21435  51562', 'NO');
INSERT INTO `tb_honeyed` VALUES (4252, '206.221.176.229', '菲律宾', 1434, '2021-10-27 15:57:08', 'UDP连接', 'UDP', '端口扫描： 1434', 'NO');
INSERT INTO `tb_honeyed` VALUES (4253, '113.234.188.206', '中国·辽宁·大连', 22, '2021-10-27 15:57:08', 'ssh连接', 'TCP', 'ssh密码登录：root&admin,root&admin123,root@root,root&test,root&test123,admin&admin,admin&admin123,admin&root,admin&test,admin&test123', 'NO');
INSERT INTO `tb_honeyed` VALUES (4254, '116.50.191.91', '中国·香港', 445, '2021-10-27 15:57:08', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4255, '172.93.189.171', '墨西哥', 445, '2021-10-27 15:57:08', 'TCP连接', 'TCP', '116.50.191.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4256, '116.50.131.91', '菲律宾', 6379, '2021-10-27 15:57:08', 'Redis连接', 'TCP', '116.50.191.131:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4257, '116.50.133.91', '泰国', 445, '2021-10-27 15:57:08', 'TCP连接', 'TCP', '116.50.133.91:61017 already connected.', 'NO');
INSERT INTO `tb_honeyed` VALUES (4258, '116.33.191.91', '菲律宾', 8, '2021-10-27 15:57:08', 'TCP连接', 'TCP', '116.33.191.91:61017 already connected.', 'NO');

-- ----------------------------
-- Table structure for tb_resource
-- ----------------------------
DROP TABLE IF EXISTS `tb_resource`;
CREATE TABLE `tb_resource`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_hide` int(0) NULL DEFAULT NULL,
  `level` int(0) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int(0) NULL DEFAULT NULL,
  `source_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `source_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `type` int(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `parent_id` int(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `FKf5ra2gn0xedeida2op8097sr5`(`parent_id`) USING BTREE,
  CONSTRAINT `FKf5ra2gn0xedeida2op8097sr5` FOREIGN KEY (`parent_id`) REFERENCES `tb_resource` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_resource
-- ----------------------------
INSERT INTO `tb_resource` VALUES (1, '2017-01-10 13:56:51', '用户管理', NULL, 0, 2, '用户管理', 1, 'system:user:index', '/admin/user/index', 1, '2021-10-22 15:24:19', NULL);
INSERT INTO `tb_resource` VALUES (2, '2017-01-10 13:56:51', '用户编辑', NULL, 0, 3, '用户编辑', 1, 'system:user:edit', '/admin/user/edit*', 2, '2017-01-10 16:26:42', 1);
INSERT INTO `tb_resource` VALUES (3, '2017-01-11 16:48:48', '用户添加', NULL, 0, 3, '用户添加', 2, 'system:user:add', '/admin/user/add', 2, '2017-01-11 16:49:26', 1);
INSERT INTO `tb_resource` VALUES (4, '2017-01-11 16:48:48', '用户删除', NULL, 0, 3, '用户删除', 3, 'system:user:deleteBatch', '/admin/user/deleteBatch', 2, '2017-01-18 14:11:41', 1);
INSERT INTO `tb_resource` VALUES (5, '2017-01-11 16:48:48', '角色分配', NULL, 0, 3, '角色分配', 4, 'system:user:grant', '/admin/user/grant/**', 2, '2017-01-18 14:11:51', 1);
INSERT INTO `tb_resource` VALUES (6, '2017-01-12 16:45:10', '角色管理', NULL, 0, 2, '角色管理', 2, 'system:role:index', '/admin/role/index', 1, '2017-01-12 16:46:52', NULL);
INSERT INTO `tb_resource` VALUES (7, '2017-01-12 16:47:02', '角色编辑', NULL, 0, 3, '角色编辑', 1, 'system:role:edit', '/admin/role/edit*', 2, '2017-01-18 10:24:06', 1);
INSERT INTO `tb_resource` VALUES (8, '2017-01-12 16:47:23', '角色添加', NULL, 0, 3, '角色添加', 2, 'system:role:add', '/admin/role/add', 2, '2017-01-12 16:49:16', 6);
INSERT INTO `tb_resource` VALUES (9, '2017-01-12 16:47:23', '角色删除', NULL, 0, 3, '角色删除', 3, 'system:role:deleteBatch', '/admin/role/deleteBatch', 2, '2017-01-18 14:12:03', 6);
INSERT INTO `tb_resource` VALUES (10, '2017-01-12 16:47:23', '资源分配', NULL, 0, 3, '资源分配', 4, 'system:role:grant', '/admin/role/grant/**', 2, '2017-01-18 14:12:11', 6);
INSERT INTO `tb_resource` VALUES (11, '2017-01-17 11:21:12', '资源管理', NULL, 0, 2, '资源管理', 3, 'system:resource:index', '/admin/resource/index', 1, '2017-01-17 11:21:42', NULL);
INSERT INTO `tb_resource` VALUES (12, '2017-01-17 11:21:52', '资源编辑', NULL, 0, 3, '资源编辑', 1, 'system:resource:edit', '/admin/resource/edit*', 2, '2017-01-17 11:22:36', 11);
INSERT INTO `tb_resource` VALUES (13, '2017-01-17 11:21:54', '资源添加', NULL, 0, 3, '资源添加', 2, 'system:resource:add', '/admin/resource/add', 2, '2017-01-17 11:22:39', 11);
INSERT INTO `tb_resource` VALUES (14, '2017-01-17 11:21:54', '资源删除', NULL, 0, 3, '资源删除', 3, 'system:resource:deleteBatch', '/admin/resource/deleteBatch', 2, '2017-01-18 14:12:31', 11);

-- ----------------------------
-- Table structure for tb_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_role`;
CREATE TABLE `tb_role`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `role_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `status` int(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_role
-- ----------------------------
INSERT INTO `tb_role` VALUES (1, '2017-01-09 17:25:30', '超级管理员', 'administrator', 'administrator', 0, '2017-01-09 17:26:25');

-- ----------------------------
-- Table structure for tb_role_resource
-- ----------------------------
DROP TABLE IF EXISTS `tb_role_resource`;
CREATE TABLE `tb_role_resource`  (
  `role_id` int(0) NOT NULL,
  `resource_id` int(0) NOT NULL,
  PRIMARY KEY (`role_id`, `resource_id`) USING BTREE,
  INDEX `FK868kc8iic48ilv5npa80ut6qo`(`resource_id`) USING BTREE,
  CONSTRAINT `FK7ffc7h6obqxflhj1aq1mk20jk` FOREIGN KEY (`role_id`) REFERENCES `tb_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FK868kc8iic48ilv5npa80ut6qo` FOREIGN KEY (`resource_id`) REFERENCES `tb_resource` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_role_resource
-- ----------------------------
INSERT INTO `tb_role_resource` VALUES (1, 1);
INSERT INTO `tb_role_resource` VALUES (1, 2);
INSERT INTO `tb_role_resource` VALUES (1, 3);
INSERT INTO `tb_role_resource` VALUES (1, 4);
INSERT INTO `tb_role_resource` VALUES (1, 5);
INSERT INTO `tb_role_resource` VALUES (1, 6);
INSERT INTO `tb_role_resource` VALUES (1, 7);
INSERT INTO `tb_role_resource` VALUES (1, 8);
INSERT INTO `tb_role_resource` VALUES (1, 9);
INSERT INTO `tb_role_resource` VALUES (1, 10);
INSERT INTO `tb_role_resource` VALUES (1, 11);
INSERT INTO `tb_role_resource` VALUES (1, 12);
INSERT INTO `tb_role_resource` VALUES (1, 13);
INSERT INTO `tb_role_resource` VALUES (1, 14);

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user`  (
  `id` int(0) NOT NULL AUTO_INCREMENT,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `birthday` datetime(0) NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `delete_status` int(0) NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `locked` int(0) NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sex` int(0) NULL DEFAULT NULL,
  `telephone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
INSERT INTO `tb_user` VALUES (1, '成都', '2017-01-17 00:00:00', '2017-01-09 17:26:41', 0, '超级管理员', 'whoismy8023@163.com', 0, 'admin', '3931MUEQD1939MQMLM4AISPVNE', 1, '15923930000', '2021-10-22 11:58:17', 'admin');

-- ----------------------------
-- Table structure for tb_user_role
-- ----------------------------
DROP TABLE IF EXISTS `tb_user_role`;
CREATE TABLE `tb_user_role`  (
  `user_id` int(0) NOT NULL,
  `role_id` int(0) NOT NULL,
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE,
  INDEX `FKea2ootw6b6bb0xt3ptl28bymv`(`role_id`) USING BTREE,
  CONSTRAINT `FK7vn3h53d0tqdimm8cp45gc0kl` FOREIGN KEY (`user_id`) REFERENCES `tb_user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKea2ootw6b6bb0xt3ptl28bymv` FOREIGN KEY (`role_id`) REFERENCES `tb_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_user_role
-- ----------------------------
INSERT INTO `tb_user_role` VALUES (1, 1);

SET FOREIGN_KEY_CHECKS = 1;
