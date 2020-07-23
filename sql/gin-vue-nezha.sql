/*
Navicat MySQL Data Transfer

Source Server         : localhost3306
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : gin-vue-nezha

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2020-07-23 21:31:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for nezha_application
-- ----------------------------
DROP TABLE IF EXISTS `nezha_application`;
CREATE TABLE `nezha_application` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT '1',
  `name` varchar(255) DEFAULT NULL,
  `cn_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `harbor` varchar(255) DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `project_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_application_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_application
-- ----------------------------
INSERT INTO `nezha_application` VALUES ('1', null, null, null, null, null, null, '0', 'application-1', 'app-cn-name-1', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('2', null, null, null, null, null, null, '0', 'application-2', 'app-cn-name-2', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('3', null, null, null, null, null, null, '0', 'application-3', 'app-cn-name-3', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('4', null, null, null, null, null, null, '0', 'application-4', 'app-cn-name-4', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('5', null, null, null, null, null, null, '0', 'application-5', 'app-cn-name-5', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('6', null, null, null, null, null, null, '0', 'application-6', 'app-cn-name-6', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('7', null, null, null, null, null, null, '0', 'application-7', 'app-cn-name-7', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('8', null, null, null, null, null, null, '0', 'application-8', 'app-cn-name-8', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('9', null, null, null, null, null, null, '0', 'application-9', 'app-cn-name-9', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('10', null, null, null, null, null, null, '0', 'application-10', 'app-cn-name-10', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('11', null, null, null, null, null, null, '0', 'application-11', 'app-cn-name-11', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('12', null, null, null, null, null, null, '0', 'application-12', 'app-cn-name-12', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('13', null, null, null, null, null, null, '0', 'application-13', 'app-cn-name-13', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('14', null, null, null, null, null, null, '0', 'application-14', 'app-cn-name-14', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('15', null, null, null, null, null, null, '0', 'application-15', 'app-cn-name-15', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('16', null, null, null, null, null, null, '0', 'application-16', 'app-cn-name-16', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('17', null, null, null, null, null, null, '0', 'application-17', 'app-cn-name-17', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('18', null, null, null, null, null, null, '0', 'application-18', 'app-cn-name-18', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('19', null, null, null, null, null, null, '0', 'application-19', 'app-cn-name-19', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('20', null, null, null, null, null, null, '0', 'application-20', 'app-cn-name-20', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('21', null, null, null, null, null, null, '0', 'application-21', 'app-cn-name-21', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('22', null, null, null, null, null, null, '0', 'application-22', 'app-cn-name-22', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('23', null, null, null, null, null, null, '0', 'application-23', 'app-cn-name-23', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('24', null, null, null, null, null, null, '0', 'application-24', 'app-cn-name-24', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('25', null, null, null, null, null, null, '0', 'application-25', 'app-cn-name-25', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('26', null, null, null, null, null, null, '0', 'application-26', 'app-cn-name-26', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('27', null, null, null, null, null, null, '0', 'application-27', 'app-cn-name-27', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('28', null, null, null, null, null, null, '0', 'application-28', 'app-cn-name-28', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('29', null, null, null, null, null, null, '0', 'application-29', 'app-cn-name-29', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('30', null, null, null, null, null, null, '0', 'application-30', 'app-cn-name-30', 'description', '192.168.1.111:8089/tmp/spring-helloworld', '1', '1');
INSERT INTO `nezha_application` VALUES ('31', '2020-07-21 22:26:41', '2020-07-22 21:17:59', '2020-07-22 22:28:47', 'admin', '', '0', '0', 'postman-application31', 'postman应用31', '描述project的示例项目31', '192.168.1.111:8089/postman/application31', '1', '2');
INSERT INTO `nezha_application` VALUES ('32', '2020-07-23 21:16:04', '2020-07-23 21:16:04', null, 'admin', '', '0', '1', 'hello-spring', '春天hello', 'nezha测试部署DEMO项目', '192.168.1.111:8089/abc/hello-spring', '1', '4');

-- ----------------------------
-- Table structure for nezha_deploy
-- ----------------------------
DROP TABLE IF EXISTS `nezha_deploy`;
CREATE TABLE `nezha_deploy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT '1',
  `name` varchar(255) DEFAULT NULL,
  `version` varchar(255) DEFAULT NULL,
  `reason` varchar(255) DEFAULT NULL,
  `harbor` varchar(255) DEFAULT NULL,
  `yaml` text,
  `step` int(10) unsigned DEFAULT NULL,
  `result` varchar(255) DEFAULT NULL,
  `log` text,
  `environment_id` int(10) unsigned DEFAULT NULL,
  `k8s_id` int(10) unsigned DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `project_id` int(10) unsigned DEFAULT NULL,
  `application_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_deploy_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_deploy
-- ----------------------------
INSERT INTO `nezha_deploy` VALUES ('1', '2020-07-23 21:23:59', '2020-07-23 21:24:05', null, 'admin', '', '0', '1', '20200723212358BS', 'v0.1', 'second', '192.168.1.111:8089/abc/hello-spring:v0.1', 'apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: hello-spring\nspec:\n  replicas: 1\n  selector:\n    matchLabels:\n      name: hello-spring\n  template:\n    metadata:\n      labels:\n        name: hello-spring\n    spec:\n      containers:\n        - name: hello-spring\n          image: 192.168.1.111:8089/abc/hello-spring:v0.1\n          imagePullPolicy: IfNotPresent\n          ports:\n            - containerPort: 8899\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: hello-spring-service-nodeport\nspec:\n  ports:\n    - port: 8899\n      targetPort: 8899\n      protocol: TCP\n      nodePort: 30080\n  type: NodePort\n  selector:\n    name: hello-spring', '3', 'success', 'deployments is success', '1', '4', '1', '0', '32');
INSERT INTO `nezha_deploy` VALUES ('2', '2020-07-23 21:29:54', '2020-07-23 21:30:00', null, 'admin', '', '0', '1', '20200723212953JE', 'v0.2', '测试v0.2版本', '192.168.1.111:8089/abc/hello-spring:v0.1', 'apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: hello-spring\nspec:\n  replicas: 1\n  selector:\n    matchLabels:\n      name: hello-spring\n  template:\n    metadata:\n      labels:\n        name: hello-spring\n    spec:\n      containers:\n        - name: hello-spring\n          image: 192.168.1.111:8089/abc/hello-spring:v0.1\n          imagePullPolicy: IfNotPresent\n          ports:\n            - containerPort: 8899\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: hello-spring-service-nodeport\nspec:\n  ports:\n    - port: 8899\n      targetPort: 8899\n      protocol: TCP\n      nodePort: 30080\n  type: NodePort\n  selector:\n    name: hello-spring', '3', 'success', 'deployments is success', '1', '4', '1', '0', '32');
INSERT INTO `nezha_deploy` VALUES ('3', '2020-07-23 21:30:25', '2020-07-23 21:30:31', null, 'admin', '', '0', '1', '20200723213025DX', 'v0.3', '测试v0.3版本', '192.168.1.111:8089/abc/hello-spring:v0.1', 'apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: hello-spring\nspec:\n  replicas: 1\n  selector:\n    matchLabels:\n      name: hello-spring\n  template:\n    metadata:\n      labels:\n        name: hello-spring\n    spec:\n      containers:\n        - name: hello-spring\n          image: 192.168.1.111:8089/abc/hello-spring:v0.1\n          imagePullPolicy: IfNotPresent\n          ports:\n            - containerPort: 8899\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: hello-spring-service-nodeport\nspec:\n  ports:\n    - port: 8899\n      targetPort: 8899\n      protocol: TCP\n      nodePort: 30080\n  type: NodePort\n  selector:\n    name: hello-spring', '3', 'success', 'deployments is success', '1', '4', '1', '0', '32');

-- ----------------------------
-- Table structure for nezha_environment
-- ----------------------------
DROP TABLE IF EXISTS `nezha_environment`;
CREATE TABLE `nezha_environment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_environment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_environment
-- ----------------------------
INSERT INTO `nezha_environment` VALUES ('1', null, null, null, null, null, null, null, 'DEV', '用于开发测试的环境');
INSERT INTO `nezha_environment` VALUES ('2', null, null, null, null, null, null, null, 'PRD', '用于线上运行的环境');

-- ----------------------------
-- Table structure for nezha_k8s
-- ----------------------------
DROP TABLE IF EXISTS `nezha_k8s`;
CREATE TABLE `nezha_k8s` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `environment_id` int(10) unsigned DEFAULT NULL,
  `conf` text,
  `terminal` varchar(255) DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_k8s_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_k8s
-- ----------------------------
INSERT INTO `nezha_k8s` VALUES ('1', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'k8sDev', '1', 'kubeconf k8s Dev', null, '1');
INSERT INTO `nezha_k8s` VALUES ('2', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'k8sPrd', '2', 'kubeconf k8s Prd', null, '2');
INSERT INTO `nezha_k8s` VALUES ('3', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'k8sStg', '1', 'kubeconf k8s Prd', null, '1');
INSERT INTO `nezha_k8s` VALUES ('4', '2020-07-23 21:19:18', '2020-07-23 21:19:18', null, 'admin', '', '0', '1', 'K8S-DEV-ABC', '1', 'apiVersion: v1\nclusters:\n- cluster:\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUJXRENCL3FBREFnRUNBZ0VBTUFvR0NDcUdTTTQ5QkFNQ01DTXhJVEFmQmdOVkJBTU1HR3N6Y3kxelpYSjIKWlhJdFkyRkFNVFU1TkRFek5EY3dOREFlRncweU1EQTNNRGN4TlRFeE5EUmFGdzB6TURBM01EVXhOVEV4TkRSYQpNQ014SVRBZkJnTlZCQU1NR0dzemN5MXpaWEoyWlhJdFkyRkFNVFU1TkRFek5EY3dOREJaTUJNR0J5cUdTTTQ5CkFnRUdDQ3FHU000OUF3RUhBMElBQk9HZkdLYWVjRTZZMEdPSUdZcExtOHB3RCswQ0JwSiszM0kwK041RlJZc2wKZzRyV0p0Rk84NE5TVnkyVjVhQUFvUkhnQ1diTndZdW90TnhkMVFObzEvQ2pJekFoTUE0R0ExVWREd0VCL3dRRQpBd0lDcERBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUFvR0NDcUdTTTQ5QkFNQ0Ewa0FNRVlDSVFEL0gxME5OTmJGCmNkSEdMVG10ZElkazJySm9lOUxHa01ONlE5Qnk0QjA3M0FJaEFLTVhwei9YcitEMmdUNXdwdVlqVDlVZzRseEYKMjA5cHlFaWdhKy93SDI2TgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\n    server: https://192.168.1.111:6443\n  name: default\ncontexts:\n- context:\n    cluster: default\n    user: default\n  name: default\ncurrent-context: default\nkind: Config\npreferences: {}\nusers:\n- name: default\n  user:\n    password: 418077e8c197f409bbec37f7963ee8b2\n    username: admin', 'http://192.168.1.111:8188', '1');

-- ----------------------------
-- Table structure for nezha_pod
-- ----------------------------
DROP TABLE IF EXISTS `nezha_pod`;
CREATE TABLE `nezha_pod` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT NULL,
  `pod_name` varchar(255) DEFAULT NULL,
  `name_space` varchar(255) DEFAULT NULL,
  `container` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `version` varchar(255) DEFAULT NULL,
  `environment_id` int(10) unsigned DEFAULT NULL,
  `k8s_id` int(10) unsigned DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `application_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_pod_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_pod
-- ----------------------------
INSERT INTO `nezha_pod` VALUES ('1', '2020-07-23 21:24:05', '2020-07-23 21:24:05', '2020-07-23 21:30:00', '', '', '0', '0', 'hello-spring-55cc7d5f49-2hjzg-0', 'default', 'hello-spring', '', '', '1', '4', '1', '32');
INSERT INTO `nezha_pod` VALUES ('2', '2020-07-23 21:30:00', '2020-07-23 21:30:00', '2020-07-23 21:30:32', '', '', '0', '0', 'hello-spring-55cc7d5f49-2hjzg-0', 'default', 'hello-spring', '', '', '1', '4', '1', '32');
INSERT INTO `nezha_pod` VALUES ('3', '2020-07-23 21:30:32', '2020-07-23 21:30:32', null, '', '', '0', '0', 'hello-spring-55cc7d5f49-2hjzg-0', 'default', 'hello-spring', '', '', '1', '4', '1', '32');

-- ----------------------------
-- Table structure for nezha_project
-- ----------------------------
DROP TABLE IF EXISTS `nezha_project`;
CREATE TABLE `nezha_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `cn_name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_project_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_project
-- ----------------------------
INSERT INTO `nezha_project` VALUES ('1', '2020-07-18 08:45:28', '2020-07-18 08:45:28', null, 'admin', null, '0', '1', 'AI-PROJECT', '算法项目', '用于算法', '1');
INSERT INTO `nezha_project` VALUES ('2', '2020-07-18 08:45:28', '2020-07-18 08:45:28', null, 'admin', null, '0', '1', 'BI-PROJECT', '商务智能项目', '用于BI智能', '1');
INSERT INTO `nezha_project` VALUES ('3', '2020-07-18 08:45:28', '2020-07-18 08:45:28', null, 'admin', null, '0', '1', 'CI-PROJECT', '持续集成项目', '用于JKINS', '1');
INSERT INTO `nezha_project` VALUES ('4', '2020-07-18 08:45:28', '2020-07-18 08:45:28', null, 'admin', '', '0', '1', 'ABC', 'HELLO机器 人', '用于测试NEZHA的项目', '1');
INSERT INTO `nezha_project` VALUES ('5', '2020-07-18 21:25:24', '2020-07-18 22:13:00', '2020-07-19 21:10:02', 'admin', '', '0', '0', 'proejct-demo2', '示例项目2', '描述project的示例项目2', '0');

-- ----------------------------
-- Table structure for nezha_release
-- ----------------------------
DROP TABLE IF EXISTS `nezha_release`;
CREATE TABLE `nezha_release` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT '1',
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `branch` varchar(255) DEFAULT NULL,
  `yaml` text,
  `user_id` int(10) unsigned DEFAULT NULL,
  `project_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_release_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_release
-- ----------------------------

-- ----------------------------
-- Table structure for nezha_user
-- ----------------------------
DROP TABLE IF EXISTS `nezha_user`;
CREATE TABLE `nezha_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT '1',
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `user_type` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_user
-- ----------------------------
INSERT INTO `nezha_user` VALUES ('1', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'admin', '$2a$10$RKrDOAm70HEC2nTauRfFzON469s0roEFSdcxUc0dSyxeQl4gofsXm', '1');
INSERT INTO `nezha_user` VALUES ('2', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'dev', '$2a$10$RKrDOAm70HEC2nTauRfFzON469s0roEFSdcxUc0dSyxeQl4gofsXm', '2');
INSERT INTO `nezha_user` VALUES ('3', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'chengang', '$2a$10$RKrDOAm70HEC2nTauRfFzON469s0roEFSdcxUc0dSyxeQl4gofsXm', '2');

-- ----------------------------
-- Table structure for nezha_yaml
-- ----------------------------
DROP TABLE IF EXISTS `nezha_yaml`;
CREATE TABLE `nezha_yaml` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_by` varchar(255) DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  `deleted` int(10) unsigned DEFAULT NULL,
  `state` int(10) unsigned DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `yaml` text,
  `k8s_id` int(10) unsigned DEFAULT NULL,
  `application_id` int(10) unsigned DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_nezha_yaml_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of nezha_yaml
-- ----------------------------
INSERT INTO `nezha_yaml` VALUES ('1', '2020-04-06 13:01:09', null, null, null, null, null, '1', 'YamlAppK8s', 'desc', 'app k8s yaml', '1', '1', '1');
INSERT INTO `nezha_yaml` VALUES ('2', '2020-07-23 21:23:33', '2020-07-23 21:23:33', null, 'admin', '', '0', '1', 'YAML-DEV-ABC', 'description', 'apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: hello-spring\nspec:\n  replicas: 1\n  selector:\n    matchLabels:\n      name: hello-spring\n  template:\n    metadata:\n      labels:\n        name: hello-spring\n    spec:\n      containers:\n        - name: hello-spring\n          image: 192.168.1.111:8089/abc/hello-spring:v0.1\n          imagePullPolicy: IfNotPresent\n          ports:\n            - containerPort: 8899\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: hello-spring-service-nodeport\nspec:\n  ports:\n    - port: 8899\n      targetPort: 8899\n      protocol: TCP\n      nodePort: 30080\n  type: NodePort\n  selector:\n    name: hello-spring', '4', '32', '1');
