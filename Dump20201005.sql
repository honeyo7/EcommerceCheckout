CREATE DATABASE  IF NOT EXISTS `ecom_checkout` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `ecom_checkout`;
-- MySQL dump 10.13  Distrib 8.0.17, for Win64 (x86_64)
--
-- Host: localhost    Database: ecom_checkout
-- ------------------------------------------------------
-- Server version	8.0.17


--
-- Table structure for table `ecom_offer`
--

DROP TABLE IF EXISTS `ecom_offer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ecom_offer` (
  `offer_id` int(11) NOT NULL AUTO_INCREMENT,
  `pdt_sku` varchar(45) NOT NULL,
  `min_qty` int(11) NOT NULL,
  `offer_type` tinyint(1) NOT NULL COMMENT '1-On full amount; 2-On Product',
  `disc_per` double NOT NULL,
  `offer_pdt_sku` varchar(45) NOT NULL,
  `offer_pdt_qty` int(11) NOT NULL,
  PRIMARY KEY (`offer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ecom_offer`
--

LOCK TABLES `ecom_offer` WRITE;
/*!40000 ALTER TABLE `ecom_offer` DISABLE KEYS */;
INSERT INTO `ecom_offer` VALUES (1,'43N23P',1,2,100,'234234',1),(2,'120P90',2,2,100,'120P90',1),(3,'A304SD',3,1,10,'0',0);
/*!40000 ALTER TABLE `ecom_offer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ecom_pdt`
--

DROP TABLE IF EXISTS `ecom_pdt`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ecom_pdt` (
  `pdt_sku` varchar(20) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `quantity` varchar(45) DEFAULT NULL,
  `offer_flag` tinyint(1) DEFAULT '0' COMMENT '0-No offer;1-offer',
  PRIMARY KEY (`pdt_sku`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ecom_pdt`
--

LOCK TABLES `ecom_pdt` WRITE;
/*!40000 ALTER TABLE `ecom_pdt` DISABLE KEYS */;
INSERT INTO `ecom_pdt` VALUES ('120P90','Google Home',49.99,'10',1),('234234','Raspberry Pi B',30,'2',0),('43N23P','MacBook Pro',5399.99,'5',1),('A304SD','Alexa Speaker',109.5,'10',1);
/*!40000 ALTER TABLE `ecom_pdt` ENABLE KEYS */;
UNLOCK TABLES;


