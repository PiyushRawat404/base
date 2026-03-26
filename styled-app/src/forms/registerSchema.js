import { z } from "zod";

export const registerSchema = z.object({
  name: z.string().min(3, "Name should contain 3 characters"),
  email: z.string().email("Invalid email"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  age: z.number().min(18, "Must be at least 18"),
});