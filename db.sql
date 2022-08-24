-- Copiando estrutura do banco de dados para procupet
CREATE DATABASE IF NOT EXISTS `procupet` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `procupet`;

-- Copiando estrutura para tabela procupet.animalper
CREATE TABLE IF NOT EXISTS `animalper` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `nome` varchar(30) NOT NULL,
  `raca` varchar(30) NOT NULL,
  `localperd` varchar(30) NOT NULL,
  `tipani` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- Exportação de dados foi desmarcado.

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;