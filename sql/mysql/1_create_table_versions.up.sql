CREATE TABLE `versions` (
  `versions_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `workload` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `platform` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `environment` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `version` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `changelog_url` text COLLATE utf8_unicode_ci DEFAULT NULL,
  `raw` text COLLATE utf8_unicode_ci DEFAULT NULL,
  `date` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`versions_id`),
  KEY `idx_workload` (`workload`),
  KEY `idx_environment` (`environment`),
  KEY `idx_version` (`version`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;