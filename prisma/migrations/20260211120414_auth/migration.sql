/*
  Warnings:

  - A unique constraint covering the columns `[email]` on the table `doctor` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[email]` on the table `patients` will be added. If there are existing duplicate values, this will fail.
  - A unique constraint covering the columns `[password]` on the table `patients` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `email` to the `doctor` table without a default value. This is not possible if the table is not empty.
  - Added the required column `email` to the `patients` table without a default value. This is not possible if the table is not empty.
  - Added the required column `password` to the `patients` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "doctor" ADD COLUMN     "email" TEXT NOT NULL;

-- AlterTable
ALTER TABLE "patients" ADD COLUMN     "email" TEXT NOT NULL,
ADD COLUMN     "password" TEXT NOT NULL,
ALTER COLUMN "phoneno" DROP NOT NULL,
ALTER COLUMN "phoneno" DROP DEFAULT,
ALTER COLUMN "age" DROP NOT NULL;
DROP SEQUENCE "patients_phoneno_seq";

-- CreateIndex
CREATE UNIQUE INDEX "doctor_email_key" ON "doctor"("email");

-- CreateIndex
CREATE UNIQUE INDEX "patients_email_key" ON "patients"("email");

-- CreateIndex
CREATE UNIQUE INDEX "patients_password_key" ON "patients"("password");
