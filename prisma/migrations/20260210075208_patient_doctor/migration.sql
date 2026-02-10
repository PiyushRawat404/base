-- CreateEnum
CREATE TYPE "branch" AS ENUM ('GYANE', 'ORTHO', 'ENT', 'DENTIST', 'NEURO');

-- CreateEnum
CREATE TYPE "timing" AS ENUM ('MORNING', 'EVENING', 'NIGHT');

-- CreateTable
CREATE TABLE "doctor" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "specialist" "branch" NOT NULL,
    "shift" "timing" NOT NULL,

    CONSTRAINT "doctor_pkey" PRIMARY KEY ("id")
);
