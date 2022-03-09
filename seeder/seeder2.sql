insert into roles values (default, 'Admin'), (default, 'User');

ALTER TABLE kategori_produk
DROP CONSTRAINT fk_produk_kategori_produk,
ADD CONSTRAINT fk_produk_kategori_produk
   FOREIGN KEY (id_produk)
   REFERENCES produk(id)
   ON DELETE CASCADE;
   
ALTER TABLE gambar_produk
DROP CONSTRAINT fk_gambar_produk_produk,
ADD CONSTRAINT fk_gambar_produk_produk
   FOREIGN KEY (id_produk)
   REFERENCES produk(id)
   ON DELETE CASCADE;
   
ALTER TABLE keranjang
DROP CONSTRAINT fk_keranjang_user,
ADD CONSTRAINT fk_keranjang_user
   FOREIGN KEY (id_user)
   REFERENCES users(id)
   ON DELETE CASCADE;
   
   
INSERT INTO kategori VALUES
(DEFAULT, '2022-03-08 11:33:13.389941+07', '2022-03-08 11:33:13.389941+07', NULL, 'Flowers', 'false'),
(DEFAULT, '2022-03-08 11:33:25.461755+07', '2022-03-08 11:33:25.461755+07', NULL, 'Fashion', 'false'),
(DEFAULT, '2022-03-08 11:33:38.291445+07', '2022-03-08 11:33:38.291445+07', NULL, 'Personalized', 'false'),
(DEFAULT, '2022-03-08 11:33:49.373664+07', '2022-03-08 11:33:49.373664+07', NULL, 'Couple Gifts', 'false'),
(DEFAULT, '2022-03-08 11:33:56.969176+07', '2022-03-08 11:33:56.969176+07', NULL, 'Toys', 'false'),
(DEFAULT, '2022-03-08 11:34:04.763077+07', '2022-03-08 11:34:04.763077+07', NULL, 'Candles', 'false'),
(DEFAULT, '2022-03-08 11:34:17.276079+07', '2022-03-08 11:34:17.276079+07', NULL, 'Perfumes', 'false'),
(DEFAULT, '2022-03-08 11:34:24.809121+07', '2022-03-08 11:34:24.809121+07', NULL, 'Watches', 'false'),
(DEFAULT, '2022-03-08 11:34:40.059449+07', '2022-03-08 11:34:40.059449+07', NULL, 'Home Decor', 'false'),
(DEFAULT, '2022-03-08 11:36:44.93788+07', '2022-03-08 11:36:44.93788+07', NULL, 'Anniversary', 'true'),
(DEFAULT, '2022-03-08 11:36:54.098296+07', '2022-03-08 11:36:54.098296+07', NULL, 'Congratulations', 'true'),
(DEFAULT, '2022-03-08 11:37:00.797176+07', '2022-03-08 11:37:00.797176+07', NULL, 'Mourning', 'true'),
(DEFAULT, '2022-03-08 11:37:22.392773+07', '2022-03-08 11:37:22.392773+07', NULL, 'Wedding', 'true'),
(DEFAULT, '2022-03-08 11:37:28.279351+07', '2022-03-08 11:37:28.279351+07', NULL, 'Valentine', 'true'),
(DEFAULT, '2022-03-08 11:37:36.184238+07', '2022-03-08 11:37:36.184238+07', NULL, 'Rose', 'true'),
(DEFAULT, '2022-03-08 11:37:42.73089+07', '2022-03-08 11:37:42.73089+07', NULL, 'Lily', 'true'),
(DEFAULT, '2022-03-08 11:37:50.171129+07', '2022-03-08 11:37:50.171129+07', NULL, 'Sunflower', 'true'),
(DEFAULT, '2022-03-08 11:39:01.022081+07', '2022-03-08 11:39:01.022081+07', NULL, 'Peony', 'true'),
(DEFAULT, '2022-03-08 11:39:21.310436+07', '2022-03-08 11:39:21.310436+07', NULL, 'Tulips', 'true'),
(DEFAULT, '2022-03-08 11:39:32.63021+07', '2022-03-08 11:39:32.63021+07', NULL, 'Daisy', 'true'),
(DEFAULT, '2022-03-08 11:39:41.096661+07', '2022-03-08 11:39:41.096661+07', NULL, 'Orchid', 'true'),
(DEFAULT, '2022-03-08 11:40:05.320173+07', '2022-03-08 11:40:05.320173+07', NULL, 'Crysanthemum', 'true'),
(DEFAULT, '2022-03-08 11:40:15.314807+07', '2022-03-08 11:40:15.314807+07', NULL, 'Cally Lily', 'true'),
(DEFAULT, '2022-03-08 11:40:29.341321+07', '2022-03-08 11:40:29.341321+07', NULL, 'Carnation', 'true'),
(DEFAULT, '2022-03-08 11:40:39.06492+07', '2022-03-08 11:40:39.06492+07', NULL, 'Classic', 'true'),
(DEFAULT, '2022-03-08 11:40:54.276171+07', '2022-03-08 11:40:54.276171+07', NULL, 'Luxury', 'true'),
(DEFAULT, '2022-03-08 11:41:03.628554+07', '2022-03-08 11:41:03.628554+07', NULL, 'Modern', 'true'),
(DEFAULT, '2022-03-08 11:41:13.518738+07', '2022-03-08 11:41:13.518738+07', NULL, 'Pastel', 'true'),
(DEFAULT, '2022-03-08 11:41:26.100209+07', '2022-03-08 11:41:26.100209+07', NULL, 'Pink', 'true'),
(DEFAULT, '2022-03-08 11:41:37.258479+07', '2022-03-08 11:41:37.258479+07', NULL, 'Blue', 'true'),
(DEFAULT, '2022-03-08 11:41:44.995529+07', '2022-03-08 11:41:44.995529+07', NULL, 'Gold', 'true'),
(DEFAULT, '2022-03-08 11:41:54.544638+07', '2022-03-08 11:41:54.544638+07', NULL, 'Orange', 'true'),
(DEFAULT, '2022-03-08 11:42:01.6606+07', '2022-03-08 11:42:01.6606+07', NULL, 'Red', 'true'),
(DEFAULT, '2022-03-08 11:42:08.351366+07', '2022-03-08 11:42:08.351366+07', NULL, 'White', 'true'),
(DEFAULT, '2022-03-08 11:33:30.717901+07', '2022-03-08 11:56:52.24705+07', NULL, 'Fashion', 'false'),
(DEFAULT, '2022-03-08 11:39:49.891421+07', '2022-03-08 11:58:22.217446+07', NULL, 'Aster', 'true');

insert into metode_pembayaran values
(default, 'BCA', 'Virtual Account', 1000, 60*24),
(default, 'BRI', 'Virtual Account', 1000, 60*24),
(default, 'Mandiri', 'Virtual Account', 1000, 60*24),
(default, 'VISA', 'Kartu Debit', 0, 60),
(default, 'MasterCard', 'Kartu Debit', 0, 60),
(default, 'LinkAja', 'E-Wallet', 1000, 60*24),
(default, 'OVO', 'E-Wallet', 1000, 60*24),
(default, 'GoPay', 'E-Wallet', 1000, 60*24),
(default, 'ShopeePay', 'E-Wallet', 1000, 60*24),
(default, 'DANA', 'E-Wallet', 1000, 60*24);
