CREATE TABLE `bandtogether`.`event` (
  `eventId` INT NOT NULL AUTO_INCREMENT,
  `ownerId` INT NOT NULL,
  `json` JSON NULL,
  PRIMARY KEY (`eventId`),
  FOREIGN KEY (`ownerId`) REFERENCES user(`UserId`)
  ON DELETE CASCADE
  ON UPDATE CASCADE,
  UNIQUE INDEX `eventId_UNIQUE` (`eventId` ASC) VISIBLE);