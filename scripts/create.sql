
-- Create pictures table
DROP TABLE IF EXISTS `pictures`;
CREATE TABLE `pictures` (
  `id` varchar(255) NOT NULL, -- uuid of the picture
  `name` varchar(255) NOT NULL, -- name of the picture
  `filename` varchar(255) NOT NULL, -- name of the uploaded file
  `content_type` varchar(255) NOT NULL, -- content type of the uploaded file
  `size` int NOT NULL, -- size of the uploaded file
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;