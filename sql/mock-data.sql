INSERT INTO users(name, nickname, email, password) 
VALUES 
("User 1", "User_1", "user1@gmail.com", "$2a$10$RYEJ7WoTM1W8KXpTG9DvDOiMwKlYKFjm4ufN0i4isvy7.QdmidKpe"),
("User 2", "User_2", "user2@gmail.com", "$2a$10$P.wDcvf2q7jN1WB2eV4r7eJefOMdTSAyAwcRxFvICbHNuGBCq46lO"),
("User 3", "User_3", "user3@gmail.com", "$2a$10$f1gYXAoukI2lp5Ob0LE5D.zR80KSE.3YaP7pV2OwdqxHR.yp1smQ.");

INSERT INTO followers(user_id, follower_id)
VALUES 
(1, 2),
(3, 1),
(1, 3);