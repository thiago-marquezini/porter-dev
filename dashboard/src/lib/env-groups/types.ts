import { z } from "zod";

export const envGroupFormValidator = z.object({
  name: z
    .string()
    .min(1, { message: "A service name is required" })
    .max(30)
    .regex(/^[a-z0-9-]+$/, {
      message: 'Lowercase letters, numbers, and " - " only.',
    }),
  envVariables: z
    .array(
      z.object({
        key: z.string().min(1, { message: "Key cannot be empty" }),
        value: z.string().min(1, { message: "Value cannot be empty" }),
        deleted: z.boolean(),
        hidden: z.boolean(),
        locked: z.boolean(),
      })
    )
    .min(1, { message: "At least one environment variable is required" })
});

export type EnvGroupFormData = z.infer<typeof envGroupFormValidator>;