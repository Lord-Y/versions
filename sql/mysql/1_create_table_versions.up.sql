CREATE TABLE `versions` (
  `versions_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `workload` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `platform` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `environment` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `version` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `changelog_url` text COLLATE utf8_unicode_ci DEFAULT NULL,
  `raw` text COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`versions_id`),
  KEY `idx_workload` (`workload`),
  KEY `idx_environment` (`environment`),
  KEY `idx_version` (`version`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;