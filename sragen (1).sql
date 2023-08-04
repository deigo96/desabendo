-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 03, 2023 at 05:51 PM
-- Server version: 10.1.38-MariaDB
-- PHP Version: 5.6.40

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sragen`
--

-- --------------------------------------------------------

--
-- Table structure for table `auth`
--

CREATE TABLE `auth` (
  `auth_id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int(1) NOT NULL,
  `is_active` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `auth`
--

INSERT INTO `auth` (`auth_id`, `username`, `email`, `password`, `role_id`, `is_active`, `created_at`, `updated_at`) VALUES
(1, 'deigo', 'deigojo@gmail.com', '$2a$04$fY9mIdXHiQZpdFEqFpIaw.uWl0nFaxIYt/Ip4XvENML20HCLKPgKa', 1, 1, '2023-07-22 08:21:35', '2023-07-21 17:00:00'),
(5, 'siahaan', 'siahaan@gmail.com', '$2a$04$fY9mIdXHiQZpdFEqFpIaw.uWl0nFaxIYt/Ip4XvENML20HCLKPgKa', 2, 1, '2023-08-02 13:09:16', '2023-07-22 11:28:49'),
(6, 'jonathan', 'jonathan@gmail.com', '$2a$04$fY9mIdXHiQZpdFEqFpIaw.uWl0nFaxIYt/Ip4XvENML20HCLKPgKa', 2, 0, '2023-08-02 13:09:02', '2023-07-22 11:30:56'),
(7, 'dri', 'a@gmail.com', '$2a$04$Bt97CDU/141HvgPhXdhRxe.dliSxRdLIFunpf0ccNJyAVnrp919ga', 2, 0, '2023-08-02 13:09:09', '2023-07-23 01:22:50'),
(8, 'sobu', 'sobu@gmail.com', '$2a$04$NQD0X5KoSc.3gwaftnUtH.6VmuzDsJ6XOUTQ9EBRHV1CRUI2avzHu', 3, 1, '2023-08-03 15:35:57', '2023-07-23 01:54:17'),
(9, 'deigost', 'brengsek96', '$2a$04$tf1PfjVRSKHKlSE/nh5Hhuof8kfpIMtP/dE.0I3uon.AXxS7LAPHK', 3, 1, '2023-08-03 15:36:07', '2023-08-02 13:54:08');

-- --------------------------------------------------------

--
-- Table structure for table `master_pengajuan`
--

CREATE TABLE `master_pengajuan` (
  `id_jenis_pengajuan` int(11) NOT NULL,
  `nama_pengajuan` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_active` int(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `master_pengajuan`
--

INSERT INTO `master_pengajuan` (`id_jenis_pengajuan`, `nama_pengajuan`, `created_at`, `is_active`) VALUES
(1, 'Surat Domisili', '2023-07-30 01:57:13', 1),
(2, 'Surat Keterangan Usaha', '2023-07-30 01:57:13', 1),
(3, 'Surat Kartu Keluarga', '2023-07-30 01:57:13', 1),
(4, 'Surat Kartu Tanda Penduduk', '2023-07-30 01:57:13', 1),
(5, 'Surat Kematian', '2023-07-30 01:57:13', 1),
(6, 'Surat Kelahiran', '2023-07-30 01:57:13', 1),
(7, 'Surat kehilangan', '2023-07-30 01:57:56', 1);

-- --------------------------------------------------------

--
-- Table structure for table `pengajuan_dokumen`
--

CREATE TABLE `pengajuan_dokumen` (
  `id_pengajuan` int(11) NOT NULL,
  `auth_id` int(11) NOT NULL,
  `nama` varchar(100) DEFAULT NULL,
  `nik` varchar(20) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `jenis_kelamin` varchar(20) DEFAULT NULL,
  `agama` varchar(20) DEFAULT NULL,
  `pekerjaan` varchar(50) DEFAULT NULL,
  `tempat_lahir` varchar(50) DEFAULT NULL,
  `tanggal_lahir` date DEFAULT NULL,
  `alamat` varchar(100) DEFAULT NULL,
  `status_perkawinan` varchar(30) DEFAULT NULL,
  `foto` varchar(100) DEFAULT NULL,
  `ktp` varchar(100) DEFAULT NULL,
  `alasan` varchar(255) DEFAULT NULL,
  `keperluan` varchar(255) DEFAULT NULL,
  `status` int(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NULL DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  `jenis_pengajuan` int(11) NOT NULL,
  `kepala_keluarga` varchar(50) DEFAULT NULL,
  `nomor_kk` varchar(30) DEFAULT NULL,
  `tanda_tangan` varchar(255) DEFAULT NULL,
  `jenis_usaha` varchar(255) DEFAULT NULL,
  `nama_anak` varchar(50) DEFAULT NULL,
  `nama_ibu` varchar(50) DEFAULT NULL,
  `nama_ayah` varchar(50) DEFAULT NULL,
  `jenis_kelamin_anak` varchar(20) DEFAULT NULL,
  `nama_pengaju` varchar(50) DEFAULT NULL,
  `hubungan_pengaju` varchar(50) DEFAULT NULL,
  `alasan_penolakan` varchar(255) DEFAULT NULL,
  `dokumen` varchar(100) DEFAULT NULL,
  `hari_meninggal` varchar(20) DEFAULT NULL,
  `benda_hilang` varchar(100) DEFAULT NULL,
  `umur` int(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pengajuan_dokumen`
--

INSERT INTO `pengajuan_dokumen` (`id_pengajuan`, `auth_id`, `nama`, `nik`, `phone`, `jenis_kelamin`, `agama`, `pekerjaan`, `tempat_lahir`, `tanggal_lahir`, `alamat`, `status_perkawinan`, `foto`, `ktp`, `alasan`, `keperluan`, `status`, `created_at`, `updated_at`, `updated_by`, `jenis_pengajuan`, `kepala_keluarga`, `nomor_kk`, `tanda_tangan`, `jenis_usaha`, `nama_anak`, `nama_ibu`, `nama_ayah`, `jenis_kelamin_anak`, `nama_pengaju`, `hubungan_pengaju`, `alasan_penolakan`, `dokumen`, `hari_meninggal`, `benda_hilang`, `umur`) VALUES
(1, 8, 'Deigo', '1151352352352', '0812412414', 'laki-laki', 'kristen', 'developer', 'balige', '1996-11-14', 'grogol', NULL, '23243261-7a68-4889-9736-fd794efbbbb6.jpg', '7d85d9c4-8f33-40ff-b041-c8e9cd991848.jpg', 'testing', '', 2, '2023-07-24 15:55:23', '2023-08-01 14:32:49', 1, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL, 'asasddsgsdg', '', NULL, NULL, NULL),
(2, 8, 'Jonathan', '1151352352352', '0812412414', 'laki-laki', 'kristen', 'developer', 'balige', '1996-11-14', 'grogol asjfaposjf asfka [ksf[askf [aksf[ aksf[pkas[fkp', NULL, 'df686752-dc7f-47ca-8218-35b4b2d8f803.jpg', 'eaf4844d-6151-474e-9608-ee7d420d6910.jpg', 'testing', NULL, 0, '2023-07-30 03:07:49', '2023-07-30 03:07:49', NULL, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '', NULL, NULL, NULL),
(3, 0, NULL, '', '', '', 'Kristen Protestan', 'asd', 'Balige Toba Samosir', '2023-07-06', 'asd', NULL, '414ffb9a-bbc4-4ee1-99da-59c4abccbfc9.jpg', '711ba75c-16b4-4db1-80c8-acdfa1c23b76.jpg', '', NULL, 0, '2023-07-30 03:35:21', '2023-07-30 03:35:21', NULL, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '', NULL, NULL, NULL),
(4, 8, 'Siahaan', '', '', '', 'Kristen Protestan', 'asd', 'Balige Toba Samosir', '2023-07-06', 'asd', 'asd', '6993a000-b2ec-4148-9a3e-abb8af1751ce.jpg', 'b5fad330-9d41-4ad5-ba7b-3d3128a4b894.jpg', '', 'asd', 1, '2023-07-30 03:39:59', '2023-08-01 14:41:41', 1, 1, NULL, NULL, NULL, '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, '3ae5724c-d036-4a52-8ea2-cc9820aec068.pdf', NULL, NULL, NULL),
(5, 8, 'Deigo', '12412412412', NULL, '', 'Kristen Protestan', NULL, NULL, NULL, 'asd', NULL, NULL, NULL, 'asdfaegka', NULL, 1, '2023-07-30 14:02:17', '2023-08-01 15:28:13', 1, 3, 'beae', '12515151551', '04b2c871-c8f9-463e-b93f-00835e97b158.jpg', '', NULL, NULL, NULL, NULL, NULL, NULL, 'asd', '89dc3f0b-2a3e-416a-9ebd-392fa0436978.pdf', NULL, NULL, NULL),
(6, 8, 'Deigo', NULL, NULL, NULL, NULL, 'asd', NULL, NULL, 'sdgsdgsdg', NULL, NULL, '1b52594a-3ec4-475a-b0a9-2ba93991f761.jpg', NULL, NULL, 1, '2023-07-30 14:32:09', '2023-08-01 15:43:53', 1, 2, NULL, NULL, NULL, 'sdgsdgsdg', NULL, NULL, NULL, NULL, NULL, NULL, 'adgadgadg', 'b6d3b2e9-6061-4be7-ba42-cae3397af7c5.pdf', NULL, NULL, NULL),
(7, 8, '', NULL, NULL, NULL, NULL, NULL, 'London', '2023-06-28', 'asda', NULL, NULL, '3ff13709-f2cd-434d-9daa-33e7dd4a3e0d.jpg', NULL, NULL, 1, '2023-07-30 14:57:07', '2023-08-01 16:05:05', 1, 6, NULL, NULL, NULL, NULL, 'asd', 'dasd', 'dfjfgj', 'Laki-Laki', 'wywry', 'sdgsdg', NULL, 'a5d7a7c2-5cf9-47b0-a360-9111cddae88e.pdf', NULL, NULL, NULL),
(8, 8, 'doni', NULL, NULL, 'Laki-Laki', NULL, NULL, 'sorong', '2023-07-15', NULL, NULL, NULL, '7c0f8f10-4f36-467a-b488-e88fa8b80f2b.jpg', 'jantun', NULL, 0, '2023-08-02 00:46:15', '2023-08-02 00:46:15', NULL, 5, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'sabtu', NULL, NULL),
(9, 8, 'jono', '1513513513513', NULL, NULL, NULL, 'buruh', NULL, '2023-07-30', NULL, NULL, NULL, '55442a01-dac9-48e5-b4c4-3b1bd6eb90a1.jpg', 'hilang', NULL, 0, '2023-08-02 00:53:13', '2023-08-02 00:53:13', NULL, 7, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'tas', NULL),
(10, 8, 'sri', '3526246', NULL, NULL, NULL, NULL, NULL, '2023-07-31', 'asd', NULL, NULL, NULL, NULL, 'Penggantian', 0, '2023-08-02 01:11:35', '2023-08-02 01:11:35', NULL, 4, NULL, NULL, 'eaad243f-0ba2-4e13-bea4-507c0bf08805.jpg', NULL, NULL, 'ani', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL),
(11, 9, 'deigo jonathan', NULL, NULL, NULL, 'Kristen Protestan', 'developer', 'balige', '1996-11-14', 'grogol', 'belum menikah', '62af8851-d9cc-423f-b160-bfee6358aaf0.jpg', '763a2e20-6417-47a7-b1f9-76b949cef20f.jpg', NULL, 'surat domisili', 1, '2023-08-03 15:47:12', '2023-08-03 15:48:20', 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, 'dbbf6e71-6e54-4dff-b6dd-04a30d94e198.pdf', NULL, NULL, 0);

-- --------------------------------------------------------

--
-- Table structure for table `role`
--

CREATE TABLE `role` (
  `role_id` int(1) NOT NULL,
  `role_name` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `role`
--

INSERT INTO `role` (`role_id`, `role_name`) VALUES
(1, 'super'),
(2, 'admin'),
(3, 'user');

-- --------------------------------------------------------

--
-- Table structure for table `teams`
--

CREATE TABLE `teams` (
  `id` int(3) NOT NULL,
  `nama` varchar(50) DEFAULT NULL,
  `jabatan` varchar(50) DEFAULT NULL,
  `tentang` varchar(255) DEFAULT NULL,
  `foto` varchar(255) DEFAULT NULL,
  `twitter` varchar(100) DEFAULT NULL,
  `facebook` varchar(100) DEFAULT NULL,
  `instagram` varchar(100) DEFAULT NULL,
  `linkedin` varchar(100) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `created_by` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `auth`
--
ALTER TABLE `auth`
  ADD PRIMARY KEY (`auth_id`);

--
-- Indexes for table `master_pengajuan`
--
ALTER TABLE `master_pengajuan`
  ADD PRIMARY KEY (`id_jenis_pengajuan`);

--
-- Indexes for table `pengajuan_dokumen`
--
ALTER TABLE `pengajuan_dokumen`
  ADD PRIMARY KEY (`id_pengajuan`),
  ADD KEY `auth_id` (`auth_id`);

--
-- Indexes for table `teams`
--
ALTER TABLE `teams`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `auth`
--
ALTER TABLE `auth`
  MODIFY `auth_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `master_pengajuan`
--
ALTER TABLE `master_pengajuan`
  MODIFY `id_jenis_pengajuan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `pengajuan_dokumen`
--
ALTER TABLE `pengajuan_dokumen`
  MODIFY `id_pengajuan` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `teams`
--
ALTER TABLE `teams`
  MODIFY `id` int(3) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
