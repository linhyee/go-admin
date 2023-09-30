-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: blog
-- ------------------------------------------------------
-- Server version	8.0.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `blog`
--

DROP TABLE IF EXISTS `blog`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blog` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) NOT NULL,
  `user_name` varchar(50) NOT NULL,
  `user_image` varchar(500) NOT NULL,
  `name` varchar(50) NOT NULL,
  `summary` varchar(200) NOT NULL,
  `content` mediumtext NOT NULL,
  `created_at` double NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blog`
--

LOCK TABLES `blog` WRITE;
/*!40000 ALTER TABLE `blog` DISABLE KEYS */;
INSERT INTO `blog` VALUES (1,'1','admin','http://www.gravatar.com/avatar/34be3c7c0655313619d9b91a7e6f1ee6?d=mm&s=120','英伟达RTX 40系列明后年发布！性能两倍于RTX 3090','目前，NVIDIA的GPU架构仍然是Ampere，今年不会有全新的RTX 40系列，下代架构要等明年，未来延续2年更新一次架构的节奏','目前，NVIDIA的GPU架构仍然是Ampere，今年不会有全新的RTX 40系列，下代架构要等明年，未来延续2年更新一次架构的节奏。\n\n\n\n\n近日有推特网友曝料，RTX 40系列要等到2022年第四季度至2023年第一季度才会发布，也就是要再等上大概一年半。\n\n据悉RTX 40将采用台积电N5 5nm工艺制造，而在性能方面最高可以达到如今卡皇RTX 3090的两倍之多！\n\n\n\n此前有消息称，NVIDIA下代游戏卡GPU架构代号Ada Lovelace——阿达·洛芙莱斯女士，第一位计算机科学家，编写了历史上第一个计算机程序，人称“数字女王”。\n\n新架构的大核心编号AD102，将有多达12个GPC(图形处理集群)、72个TPC(纹理处理集群)、144个SM(流式多处理器)，而如果每个SM继续128个流处理器(CUDA核心)，那么整体下来就有多达18432个流处理器！\n\n相比于目前安培架构GA102核心的10752个，这一下子就增加了超过71％！\n\n此外，RTX 30 Super系列升级版将在明年初发布。\n\nTuring RTX 20时代经历了一次Super版本升级，现在移动版的RTX 30 Super有消息了，将在明年初登场，几乎可以肯定就是CES 2022上发布。\n\n就在日前，联想ThinkPad X1 Extreme G4游戏本的部分规格泄露，其中显卡部分赫然写着RTX 3080 Super 16GB、RTX 3070 Super 8GB，但没有看到RTX 3060 Super，而是继续使用RTX 3060。\n\n\n\nRTX 30 Super移动版的具体规格暂无消息，不过比较特殊的是，RTX 3080现在已经用了满血版的GA104核心，RTX 3080 Super不太可能只是提升频率，也不可能用规模更大、功耗更高的GA102核心。\n\n之前曾经流传出过GA103的名字，规模介于GA102、GA104之间，可能就是为RTX 3080 Super量身打造的。\n\n至于RTX 3070 Super，有可能将GA104核心再开放一些流处理器，介于5120/6144个之间，也有可能同样使用GA103。',1625760158.54571),(2,'1','kitty','','star city','this is a summary','today is good',13485723822);
/*!40000 ALTER TABLE `blog` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int NOT NULL,
  `blog_id` int NOT NULL,
  `user_id` int NOT NULL,
  `user_name` varchar(50) NOT NULL,
  `user_image` varchar(500) NOT NULL,
  `content` mediumtext NOT NULL,
  `created_at` double NOT NULL,
  `is_deleted` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (1,0,1,1,'admin','http://www.gravatar.com/avatar/34be3c7c0655313619d9b91a7e6f1ee6?d=mm&s=120','awfaweaweawe',1626281236.82243,1),(2,0,1,1,'admin','http://www.gravatar.com/avatar/34be3c7c0655313619d9b91a7e6f1ee6?d=mm&s=120','agasfsdfg',1626281272.29431,0);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `passwd` varchar(50) NOT NULL,
  `admin` tinyint(1) NOT NULL,
  `name` varchar(50) NOT NULL,
  `image` varchar(500) NOT NULL,
  `created_at` double NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin@qq.com','41189a9bd78cebd551e3b3ee2f5dd581b998f54f',1,'admin','http://www.gravatar.com/avatar/34be3c7c0655313619d9b91a7e6f1ee6?d=mm&s=120',1625759675.90478);
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'blog'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-09-29 13:48:47
