/*
  Warnings:

  - A unique constraint covering the columns `[password]` on the table `doctor` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `password` to the `doctor` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "doctor" ADD COLUMN     "password" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "doctor_password_key" ON "doctor"("password");
