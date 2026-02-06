/*
  Warnings:

  - You are about to drop the `User` table. If the table is not empty, all the data it contains will be lost.

*/
-- DropTable
DROP TABLE "User";

-- CreateTable
CREATE TABLE "patients" (
    "clientno" SERIAL NOT NULL,
    "fullname" VARCHAR(20) NOT NULL,
    "phoneno" SERIAL NOT NULL,
    "gender" VARCHAR(10),
    "problem" VARCHAR(50),
    "age" INTEGER NOT NULL,

    CONSTRAINT "patients_pkey" PRIMARY KEY ("clientno")
);

-- CreateIndex
CREATE UNIQUE INDEX "patients_phoneno_key" ON "patients"("phoneno");
