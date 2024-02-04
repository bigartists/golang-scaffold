-- ----------------------------
-- Table structure for organization
-- ----------------------------


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20)   NOT NULL AUTO_INCREMENT,
    `email`       varchar(255) NOT NULL,
    `username`    varchar(255) NOT NULL,
    `password`    varchar(255) NOT NULL,
    `admin`       tinyint(1)   NOT NULL,
    `active`      tinyint(1)            DEFAULT NULL,
    `nickname`        varchar(255)          DEFAULT NULL,
    `description` varchar(255)          DEFAULT NULL,
    `avatar`      varchar(255)          DEFAULT NULL,
    `create_time` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by`   bigint(20)   NOT NULL DEFAULT '0',
    `update_time` timestamp    NULL,
    `update_by`   bigint(20)            DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_email` (`email`),
    KEY `idx_username` (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;



INSERT INTO `user` (email, username, password, admin, active, nickname, description, avatar, create_by, update_by)
VALUES
    ('alberteinstein@outlook.com', 'Albert Einstein', 'password1', 1, 1, 'Albert', 'Renowned theoretical physicist and Nobel laureate. Known for the theory of relativity and the equation E=mc². Born on March 14, 1879, in Ulm, Germany.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00238.jpg', 0, NULL),
    ('mariecurie@outlook.com', 'Marie Curie', 'password2', 0, 1, 'Marie', 'Pioneering physicist and chemist. First woman to win a Nobel Prize and the only person to win Nobel Prizes in two different sciences. Born on November 7, 1867, in Warsaw, Poland.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00071.jpg', 0, NULL),
    ('isaacnewton@outlook.com', 'Isaac Newton', 'password3', 0, NULL, 'Isaac', 'Renowned mathematician, physicist, and astronomer. Known for his laws of motion and universal gravitation. Born on December 25, 1642, in Woolsthorpe, England.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00258.jpg', 0, NULL),
    ('galileogalilei@outlook.com', 'Galileo Galilei', 'password4', 0, NULL, 'Galileo', 'Italian astronomer, physicist, and engineer. Known as the "father of observational astronomy" and the "father of modern physics." Born on February 15, 1564, in Pisa, Italy.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00219.jpg', 0, NULL),
    ('charlesdarwin@outlook.com', 'Charles Darwin', 'password5', 0, 1, 'Charles', 'English naturalist and biologist. Known for his theory of evolution through natural selection. Author of "On the Origin of Species." Born on February 12, 1809, in Shrewsbury, England.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00297.jpg', 0, NULL),
    ('nikolatesla@outlook.com', 'Nikola Tesla', 'password6', 0, NULL, 'Nikola', 'Inventor, electrical engineer, and physicist. Known for his contributions to alternating current (AC) power systems. Born on July 10, 1856, in Smiljan, Croatia.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00006.jpg', 0, NULL),
    ('adalovelace@outlook.com', 'Ada Lovelace', 'password7', 0, 1, 'Ada', 'English mathematician and writer. Recognized for her work on Charles Babbage is proposed mechanical general-purpose computer, the Analytical Engine. Born on December 10, 1815, in London, England.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00203.jpg', 0, NULL),
  ('stephenhawking@outlook.com', 'Stephen Hawking', 'password8', 0, 1, 'Stephen', 'Renowned theoretical physicist and cosmologist. Known for his work on black holes and his book "A Brief History of Time." Born on January 8, 1942, in Oxford, England.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00080.jpg', 0, NULL),
  ('alanturing@outlook.com', 'Alan Turing', 'password9', 0, NULL, 'Alan', 'English mathematician, logician, and computer scientist. Pioneering work on code-breaking during World War II and considered the father of theoretical computer science and artificial intelligence. Born on June 23, 1912, in London, England.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00260.jpg', 0, NULL),
  ('rosiefranklin@outlook.com', 'Rosie Franklin', 'password10', 0, 1, 'Rosie', 'British biophysicist and X-ray crystallographer. Contributed to the discovery of the structure of DNA. Born on July 25, 1920, in London, England.', 'https://zidongtaichu.obs.cn-central-221.ovaijisuan.com/picasso/res_data/res_data/00119.jpg', 0, NULL);


