
CREATE TABLE IF NOT EXISTS `ed2k` (
  `serial` int(11) NOT NULL auto_increment,
  `scheme` char(7) NOT NULL,
  `type` char(4) NOT NULL,
  `filename` varchar(256) NOT NULL,
  `filesize` int(11) NOT NULL,
  `hash` char(32) NOT NULL,
  `host` varchar(64) default NULL,
  `ori` varchar(512) NOT NULL,
  `rctime` datetime NOT NULL,
  PRIMARY KEY  (`serial`),
  UNIQUE KEY `hash` (`hash`)
) ENGINE=MyISAM  DEFAULT CHARSET=latin1 AUTO_INCREMENT=857 ;


