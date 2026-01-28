-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jan 28, 2026 at 10:22 AM
-- Server version: 8.0.44
-- PHP Version: 8.1.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pdf_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `pdf_files`
--

CREATE TABLE `pdf_files` (
  `id` bigint NOT NULL,
  `filename` varchar(255) NOT NULL,
  `original_name` varchar(255) DEFAULT NULL,
  `filepath` varchar(500) NOT NULL,
  `size` bigint DEFAULT NULL,
  `status` enum('CREATED','UPLOADED','DELETED') NOT NULL DEFAULT 'CREATED',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `pdf_files`
--

INSERT INTO `pdf_files` (`id`, `filename`, `original_name`, `filepath`, `size`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'upload_1769593166241838300.pdf', 'Fullstack Developer - Jennatul Macwe.pdf', 'uploads/pdf/upload_1769593166241838300.pdf', 1045831, 'UPLOADED', '2026-01-28 02:39:26', '2026-01-28 09:39:26', NULL),
(2, 'upload_1769593217649174200.pdf', 'khs smt akhir.pdf', 'uploads/pdf/upload_1769593217649174200.pdf', 61563, 'DELETED', '2026-01-28 02:40:18', '2026-01-28 09:40:56', '2026-01-28 02:40:57'),
(3, 'report_1769593790909159600.pdf', NULL, 'uploads/pdf/report_1769593790909159600.pdf', NULL, 'CREATED', '2026-01-28 02:49:51', '2026-01-28 09:49:51', NULL),
(4, 'report_1769595358413737700.pdf', NULL, 'uploads/pdf/report_1769595358413737700.pdf', NULL, 'CREATED', '2026-01-28 03:15:58', '2026-01-28 10:15:58', NULL),
(5, 'upload_1769595513502020500.pdf', 'CV - Jennatul Macwe.pdf', 'uploads/pdf/upload_1769595513502020500.pdf', 264568, 'UPLOADED', '2026-01-28 03:18:34', '2026-01-28 10:18:33', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `pdf_files`
--
ALTER TABLE `pdf_files`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `pdf_files`
--
ALTER TABLE `pdf_files`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
