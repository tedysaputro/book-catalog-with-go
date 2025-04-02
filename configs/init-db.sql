INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (1, 'Bentang Pustaka', 'Penerbit buku Indonesia yang berbasis di Yogyakarta dan dikenal menerbitkan karya-karya sastra Indonesia kontemporer', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (2, 'Truedee Books', 'Penerbit independen yang didirikan oleh Dewi Lestari untuk menerbitkan karya-karyanya sendiri', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (3, 'Gramedia Pustaka Utama', 'Salah satu penerbit terbesar di Indonesia yang menerbitkan berbagai genre buku, termasuk karya-karya bestseller nasional', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (4, 'Scribner', 'Penerbit Amerika Serikat yang merupakan bagian dari Simon & Schuster dan menerbitkan banyak karya Stephen King', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (5, 'Doubleday', 'Penerbit Amerika Serikat yang menerbitkan beberapa karya awal Stephen King termasuk The Shining', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (6, 'Gagas Media', 'Penerbit Indonesia yang fokus pada buku-buku fiksi populer dan non-fiksi untuk pembaca muda', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (7, 'Harper & Row', 'Penerbit Amerika Serikat yang menerbitkan edisi bahasa Inggris dari banyak karya Gabriel García Márquez sebelum bergabung menjadi HarperCollins', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (8, 'Editorial Sudamericana', 'Penerbit Argentina yang menerbitkan banyak karya Gabriel García Márquez dalam bahasa Spanyol aslinya', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (9, 'Media Kita', 'Penerbit Indonesia yang fokus pada buku-buku fiksi populer dan karya-karya penulis muda Indonesia', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (10, 'Grove Press', 'Penerbit Amerika Serikat yang berfokus pada karya-karya sastra terjemahan, termasuk banyak karya penulis Jepang seperti Banana Yoshimoto', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (11, 'Kadokawa Shoten', 'Penerbit Jepang yang menerbitkan banyak karya sastra Jepang kontemporer termasuk beberapa karya Banana Yoshimoto dalam bahasa aslinya', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (12, 'Doubleday', 'Penerbit Amerika Serikat yang merupakan bagian dari Penguin Random House, terkenal menerbitkan karya-karya Dan Brown termasuk The Da Vinci Code', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (13, 'Editum', 'Penerbit independen Indonesia yang fokus pada penerbitan karya sastra dan puisi Indonesia', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (14, 'Random House', 'Penerbit buku internasional yang berbasis di Amerika Serikat dan merupakan bagian dari Penguin Random House, menerbitkan berbagai genre buku termasuk karya sastra', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO public.publishers (id, name, description, created_at, updated_at)
VALUES (15, 'Hamish Hamilton', 'Penerbit Inggris yang merupakan bagian dari Penguin Books, terkenal menerbitkan karya-karya fiksi sastra berkualitas tinggi', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


-- update sequence of publisher to 16
ALTER SEQUENCE publishers_id_seq RESTART WITH 16;
